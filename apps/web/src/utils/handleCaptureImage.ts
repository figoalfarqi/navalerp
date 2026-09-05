// /* eslint-disable @typescript-eslint/no-explicit-any */
// "use client";

// import { useFetchAPI } from "@/hooks/useFetchAPI";

// type CaptureOptions = {
//   canvas: HTMLCanvasElement;
//   location?: { lat: number; lng: number };
//   address?: { kecamatan?: string; kota?: string; provinsi?: string };
//   showOverlay?: boolean;
//   folder?: string;
//   filename?: string;
//   authToken?: "none" | "driver" | "admin";
//   onSuccess?: (url: string, tempUrl: string) => void;
//   onError?: (err: any, tempUrl?: string) => void;
// };

// export function useCaptureImage() {
//   const { postAPI, uploadAPI } = useFetchAPI();

//   const captureImage = async ({
//     canvas,
//     location,
//     address,
//     showOverlay = true,
//     folder = "irregular_cost",
//     filename = "foto.jpg",
//     authToken = "driver",
//     onSuccess,
//     onError,
//   }: CaptureOptions) => {
//     try {
//       const ctx = canvas.getContext("2d");
//       if (!ctx) throw new Error("Canvas context not found");

//       // 🖋️ Tambahkan teks overlay opsional
//       if (showOverlay) {
//         const timestamp = new Date().toLocaleString();
//         ctx.font = "24px Arial";
//         ctx.fillStyle = "#155dfc";
//         ctx.fillText(`Lat: ${location?.lat?.toFixed(5) ?? "-"}`, 10, 30);
//         ctx.fillText(`Lng: ${location?.lng?.toFixed(5) ?? "-"}`, 10, 60);
//         ctx.fillText(`Time: ${timestamp}`, 10, 90);
//         ctx.fillText(`Kecamatan: ${address?.kecamatan ?? "-"}`, 10, 120);
//         ctx.fillText(`Kota: ${address?.kota ?? "-"}`, 10, 150);
//         ctx.fillText(`Provinsi: ${address?.provinsi ?? "-"}`, 10, 180);
//       }

//       // 📸 Konversi canvas ke blob
//       const blob: Blob | null = await new Promise((resolve) =>
//         canvas.toBlob(resolve, "image/jpeg", 1)
//       );
//       if (!blob) throw new Error("Failed to convert canvas to blob");

//       const tempUrl = URL.createObjectURL(blob);

//       // 🪣 Dapatkan presigned URL
//       const presignedRes = await postAPI<any>(
//         `${process.env.NEXT_PUBLIC_API_BASE_URL}/allrole/presigned`,
//         { authToken },
//         { folder, filename, action: "upload" }
//       );

//       if (!presignedRes?.data) throw new Error("Presigned URL not received");
//       const { presigned, expiry, folder: remoteFolder, filename: remoteFilename } =
//         presignedRes.data;

//       // 📤 Upload file ke file-service
//       const formUpload = new FormData();
//       formUpload.append("presigned", presigned);
//       formUpload.append("expiry", expiry);
//       formUpload.append("folder", remoteFolder);
//       formUpload.append("image", blob, remoteFilename);
//       formUpload.append("filename", remoteFilename);

//       const uploadRes = await uploadAPI<any>(
//         `${process.env.NEXT_PUBLIC_API_FILESERVICE_URL}/upload/image`,
//         { authToken },
//         formUpload
//       );

//       if (uploadRes?.code !== 200) throw new Error("Upload failed");
//       onSuccess?.(uploadRes.data.url, tempUrl);
//       return uploadRes.data.url;
//     } catch (err) {
//       console.error("useCaptureImage error:", err);
//       onError?.(err);
//       throw err;
//     }
//   };

//   return { captureImage };
// }
