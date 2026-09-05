"use client";

import { useState, useEffect } from "react";
import Camera from "@/components/camera";
import { useIndexedDB } from "@/hooks/useIndexedDB";
import { STORES } from "@/types/tableIDB";

interface PhotoSectionProps {
  title: string;
  folderPath: string;
  photoUrls?: string[];
  progressMap?: Record<number, number>;
  onCapture: (canvas: HTMLCanvasElement, folder: string) => void;
  onPressStart: (index: number) => void;
  onPressEnd: (index: number) => void;
  resolveFileUrl: (url: string) => string;
  captureMode: "add" | "edit";
  children?: React.ReactNode; // untuk overlay
}

export function PhotoSection({
  title,
  folderPath,
  photoUrls,
  progressMap = {},
  onCapture,
  onPressStart,
  onPressEnd,
  resolveFileUrl,
  captureMode,
  children,
}: PhotoSectionProps) {
  const { dbReady, getIDB } = useIndexedDB();
  const [resolvedUrls, setResolvedUrls] = useState<string[]>([]);
  useEffect(() => {
    if (!photoUrls) return;

    const resolveUrls = async () => {
      const results = await Promise.all(
        photoUrls.map(async (url) => {
          if (url.startsWith("blob")) {
            const res = await getIDB(STORES.APP_FILES, url);
            if (res?.data?.file) {
              return URL.createObjectURL(res.data.file);
            }
          }
          return resolveFileUrl(url);
        }),
      );

      setResolvedUrls(results);
    };

    if (dbReady) resolveUrls();
  }, [photoUrls, dbReady]);
  return (
    <div className="space-y-2">
      <div>{title}</div>

      <Camera
        onCapture={(canvas) => onCapture(canvas, folderPath)}
        buttonText={title}
        isBlurCheck={true}
      >
        {children}
      </Camera>

      <div className="flex flex-wrap gap-2 mt-2">
        {resolvedUrls.map((url, i) => {
          return (
            <div className="relative" key={url}>
              <img
                alt={`preview-${url}`}
                src={url}
                className="w-24 h-24 object-cover rounded-md"
                onMouseDown={() => onPressStart(i)}
                onMouseUp={() => onPressEnd(i)}
                onTouchStart={() => onPressStart(i)}
                onTouchEnd={() => onPressEnd(i)}
                style={{
                  opacity: 1 - (progressMap[i] ?? 0) / 100,
                }}
              />
            </div>
          );
        })}
      </div>
    </div>
  );
}
