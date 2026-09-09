"use client";

import { useState, useRef, DragEvent, ChangeEvent } from "react";
import { FileType, useUploadFile } from "@/hooks/useUploadFile";
import { IoCloudUploadOutline } from "@/components/icons";
import Image from "next/image";
import { authTokenType } from "@/hooks/useFetchAPI";
import ImageViewer from "../ImageViewer";
import { LuX } from "@/components/icons";

interface UploadFileFieldProps {
  id: string;
  label?: string;
  folder: string;
  authToken: authTokenType;
  allowedTypes?: FileType[];
  maxSizeMB?: number;
  value: string | string[]; // ❗ bisa single atau array
  isMultiple?: boolean; // ❗ fitur baru
  onUploaded?: (url: string | string[] | null) => void;
  error?: string;
}

export default function UploadFileField({
  id,
  label = "Upload File",
  folder,
  authToken,
  allowedTypes = ["image"],
  maxSizeMB = 5,
  value,
  isMultiple = false,
  onUploaded,
  error = "",
}: UploadFileFieldProps) {
  const {
    uploadFile,
    isUploading,
    progress,
    error: error_upload,
  } = useUploadFile();

  const fileInputRef = useRef<HTMLInputElement>(null);
  const [photoViewIndex, setPhotoViewIndex] = useState<number | null>(null);

  const handleFiles = async (files: FileList | File[]) => {
    const fileArray = Array.from(files);

    for (const file of fileArray) {
      const url = await uploadFile({
        file,
        filename: file.name,
        authToken,
        folder,
        allowedTypes,
        maxSizeMB,
      });

      // SINGLE MODE
      if (!isMultiple) {
        onUploaded?.(url);
        return;
      }

      // MULTIPLE MODE
      onUploaded?.([...(Array.isArray(value) ? value : []), url]);
    }
  };

  const onFileSelected = (e: ChangeEvent<HTMLInputElement>) => {
    if (!e.target.files) return;
    handleFiles(e.target.files);
  };

  const onDrop = (e: DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    handleFiles(e.dataTransfer.files);
  };

  const removeFile = (index: number) => {
    if (isMultiple && Array.isArray(value)) {
      const newList = value.filter((_, i) => i !== index);
      onUploaded?.(newList);
    } else {
      onUploaded?.(null);
    }
  };
  let mergedImages: string[] = [];
  if (allowedTypes.includes("image")) {
    mergedImages = [...(Array.isArray(value) ? value : value ? [value] : [])];
  }

  return (
    <div className="w-full flex flex-col gap-2">
      {/* LABEL */}
      {label && (
        <label htmlFor={id} className="font-medium">
          {label}
        </label>
      )}

      {/* DROPZONE */}
      <div
        className="border-2 border-dashed rounded-md p-3 text-center cursor-pointer bg-gray-50"
        onClick={() => fileInputRef.current?.click()}
        onDrop={onDrop}
        onDragOver={(e) => e.preventDefault()}
      >
        <input
          id={id}
          ref={fileInputRef}
          type="file"
          className="hidden"
          onChange={onFileSelected}
          multiple={isMultiple} // ❗ penting
        />

        <div className="flex flex-col items-center gap-2">
          <IoCloudUploadOutline size={48} className="text-gray-500" />
          <p className="text-gray-600">
            Drag & Drop file atau{" "}
            <span className="font-semibold underline">Pilih File</span>
          </p>
        </div>
      </div>

      {mergedImages.length > 0 && (
        <div className="mt-3">
          <div
            className={`flex gap-3 ${isMultiple ? "overflow-x-auto pb-2" : ""}`}
          >
            {mergedImages.map((img, i) => (
              <div
                key={i}
                className="relative flex-shrink-0 w-28 h-28 rounded-md"
              >
                <div
                  className="absolute right-1 top-1 bg-white/80 hover:bg-white rounded-full cursor-pointer"
                  onClick={() => removeFile(i)}
                >
                  <LuX size={20} className="text-red-600" />
                </div>

                <Image
                  src={img}
                  alt="Preview"
                  width={120}
                  height={120}
                  className="w-28 h-28 object-cover rounded-md cursor-zoom-in"
                  unoptimized
                  onClick={() => {
                    setPhotoViewIndex(i);
                  }}
                />
              </div>
            ))}
          </div>
        </div>
      )}

      {/* PROGRESS BAR */}
      {isUploading && (
        <div className="w-full bg-gray-200 h-2 rounded overflow-hidden">
          <div
            className="bg-blue-500 h-2 transition-all"
            style={{ width: `${progress}%` }}
          ></div>
        </div>
      )}

      {/* ERROR */}
      {error_upload && (
        <p className="text-red-500 text-sm">{String(error_upload)}</p>
      )}
      {error && <p className="text-red-500 text-sm">{error}</p>}
      {photoViewIndex !== null && (
        <ImageViewer
          images={mergedImages}
          activeIndex={photoViewIndex}
          onClose={() => setPhotoViewIndex(null)}
        />
      )}
    </div>
  );
}
