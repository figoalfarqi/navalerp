/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useState } from "react";
import { authTokenType, useFetchAPI } from "@/hooks/useFetchAPI";
import { useAtom } from "jotai";
import { uploadFileAtom } from "@/atoms/UploadFile";
import { drawCameraOverlay, loadMapTile } from "@/components/camera/cameraUtils";
import { STORES } from "@/types/tableIDB";
import { useIndexedDB } from "./useIndexedDB";
import { useOnlineStatus } from "./useOnlineStatus";
import { useAuth } from "@/context/AuthContext";

export type FileType =
    | "image"
    | "video"
    | "audio"
    | "document"
    | "any";

export const FileTypeMap: Record<FileType, string[]> = {
    image: ["image/*"],
    video: ["video/*"],
    audio: ["audio/*"],
    document: [
        "application/pdf",
        "application/msword",
        "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
        "application/vnd.ms-excel",
        "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
        "text/plain",
        "text/csv",
        "application/vnd.ms-powerpoint",
        "application/vnd.openxmlformats-officedocument.presentationml.presentation",
    ],
    any: ["*/*"],
};

interface BaseUploadOptions {
    filename: string;
    folder?: string;
    authToken: authTokenType;
    onSuccess?: (url: string, tempUrl?: string) => void;
    onError?: (err: any) => void;
}

interface UploadFileOptions extends BaseUploadOptions {
    file: File | Blob;
    allowedTypes?: FileType[];
    maxSizeMB?: number;
}

interface CaptureOptions extends BaseUploadOptions {
    canvas: HTMLCanvasElement;
    location?: { lat: number; lng: number };
    address?: { kecamatan?: string; kota?: string; provinsi?: string };
    showOverlay?: boolean;
}

export function useUploadFile() {
    const { postAPI } = useFetchAPI();
    const { adminToken, checkerToken, driverToken } = useAuth();
    const [uploadFileData, setUploadFileData] = useAtom(uploadFileAtom);

    const { putIDB } = useIndexedDB();
    const { isOnline } = useOnlineStatus();
    const [isUploading, setIsUploading] = useState(false);
    const [progress, setProgress] = useState(0);
    const [error, setError] = useState<any>(null);


    // =========================================================
    // 🔥 CORE UPLOAD LOGIC (Dipakai oleh File & Capture)
    // =========================================================
    const uploadToServer = async (
        blob: Blob,
        { filename, folder, authToken, onSuccess, onError }: BaseUploadOptions,
    ) => {
        try {
            setIsUploading(true);
            setProgress(0);
            setError(null);
            setUploadFileData({ folderName: folder ?? "", isUploading: true });

            // 1️⃣ Presigned
            const presigned = await postAPI<any>(
                `${process.env.NEXT_PUBLIC_API_BASE_URL}/allrole/presigned`,
                { authToken },
                { filename, folder, action: "upload" },
            );

            if (!presigned?.data)
                throw new Error("Presigned URL not received");

            const {
                presigned: signed,
                expiry,
                folder: remoteFolder,
                filename: remoteFilename,
            } = presigned.data;

            // 2️⃣ Upload
            const form = new FormData();
            form.append("presigned", signed);
            form.append("expiry", expiry);
            form.append("folder", remoteFolder);
            form.append("image", blob, remoteFilename);
            form.append("filename", remoteFilename);

            const uploadPromise = new Promise<any>((resolve, reject) => {
                const xhr = new XMLHttpRequest();
                const accessToken =
                    authToken === "checker"
                        ? checkerToken
                        : authToken === "driver"
                          ? driverToken
                          : authToken === "admin"
                            ? adminToken
                            : null;

                xhr.upload.onprogress = (e) => {
                    if (e.lengthComputable) {
                        setProgress((e.loaded / e.total) * 100);
                    }
                };

                xhr.onload = () => {
                    if (xhr.status >= 200 && xhr.status < 300) {
                        resolve(JSON.parse(xhr.responseText));
                    } else {
                        reject(new Error("Upload failed"));
                    }
                };

                xhr.onerror = () => reject(new Error("Upload error"));
                xhr.open(
                    "POST",
                    `${process.env.NEXT_PUBLIC_API_FILESERVICE_URL}/upload/image`,
                );
                if (accessToken) {
                    xhr.setRequestHeader("Authorization", `Bearer ${accessToken}`);
                }
                xhr.send(form);
            });

            const uploaded = await uploadPromise;
            if (uploaded.code !== 200)
                throw new Error(uploaded.message || "Upload failed");

            const finalUrl = uploaded.data.url;

            setProgress(100);
            setIsUploading(false);
            setUploadFileData({ folderName: folder ?? "", isUploading: false });

            onSuccess?.(finalUrl);
            return finalUrl;
        } catch (err) {
            setIsUploading(false);
            setUploadFileData({ folderName: folder ?? "", isUploading: false });
            setError(err);
            onError?.(err);
            throw err;
        }
    };

    // =========================================================
    // 📂 Upload File Normal
    // =========================================================
    const uploadFile = async ({
        file,
        allowedTypes = ["image"],
        maxSizeMB = 5,
        filename,
        folder,
        authToken,
        onSuccess,
        onError,
    }: UploadFileOptions) => {
        const maxBytes = maxSizeMB * 1024 * 1024;
        if (file.size > maxBytes) {
            throw new Error(`Ukuran file melebihi ${maxSizeMB}MB`);
        }

        const allowedMimeList = allowedTypes.flatMap((t) => FileTypeMap[t]);
        if (!allowedMimeList.includes("*/*")) {
            const match = allowedMimeList.some((mime) =>
                mime.endsWith("/*")
                    ? file.type.startsWith(mime.replace("/*", ""))
                    : file.type === mime,
            );
            if (!match) throw new Error("Tipe file tidak diizinkan");
        }

        if (!isOnline) {
            const temp_url = URL.createObjectURL(file);
            await putIDB(STORES.APP_FILES, {
                [`${STORES.APP_FILES}_id`]: temp_url,
                file, // blob disimpan langsung
                filename,
                folder,
                authToken,
                status: "pending",
                created_at: Date.now(),
            });
            return "stored-offline";
        }

        return uploadToServer(file, {
            filename,
            folder,
            authToken,
            onSuccess,
            onError,
        });
    };

    // =========================================================
    // 📸 Capture Image (Gabung Disini)
    // =========================================================
    const captureImage = async ({
        canvas,
        location,
        address,
        showOverlay = true,
        filename,
        folder,
        authToken,
        onSuccess,
        onError,
    }: CaptureOptions) => {
        const ctx = canvas.getContext("2d");
        if (!ctx) throw new Error("Canvas context not found");

        if (showOverlay) {
            let mapTileImage: HTMLImageElement | undefined;

            if (location) {
                mapTileImage = await loadMapTile(location.lat, location.lng);
            }

            drawCameraOverlay({
                ctx,
                canvas,
                location,
                address,
                mapTileImage,
            });
        }

        const blob: Blob | null = await new Promise((resolve) =>
            canvas.toBlob(resolve, "image/jpeg", 1),
        );

        if (!blob) throw new Error("Failed to convert canvas to blob");

        const tempUrl = URL.createObjectURL(blob);
        if (!isOnline) {
            await putIDB(STORES.APP_FILES, {
                [`${STORES.APP_FILES}_id`]: tempUrl,
                file: blob,
                filename,
                folder,
                authToken,
                status: "pending",
                created_at: Date.now(),
            });
            return onSuccess?.(tempUrl);
        }

        return uploadToServer(
            blob,
            { filename, folder, authToken, onSuccess, onError },
        );
    };

    return {
        uploadFileData,
        uploadFile,
        captureImage,
        isUploading,
        progress,
        error,
    };
}
