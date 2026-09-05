"use client";

import { useEffect, useRef, useState } from "react";
import Button from "../form/Button";
import { LuCamera, LuX } from "react-icons/lu";
import { BsArrowRepeat } from "react-icons/bs";
import { usePathname } from "next/navigation";
import { setCameraStream } from "./cameraUtils";
import Modal from "../Modal";
import Image from "next/image";

interface CameraProps {
  onCapture: (photo: HTMLCanvasElement) => void;
  buttonText?: string;
  isBlurCheck?: boolean;
  children: React.ReactNode;
}

export default function Camera({
  onCapture,
  buttonText = "Ambil Foto",
  isBlurCheck = false,
  children,
}: CameraProps) {
  const videoRef = useRef<HTMLVideoElement>(null);
  const [stream, setStream] = useState<MediaStream | null>(null);
  const [cameraActive, setCameraActive] = useState(false);
  const [facingMode, setFacingMode] = useState<"user" | "environment">(
    "environment"
  );

  const [blurPreview, setBlurPreview] = useState<string | null>(null);
  const [showBlurModal, setShowBlurModal] = useState(false);
  const [varianceBlur, setVarianceBlur] = useState(0);

  const stopCamera = () => {
    if (!stream) return;
    stream.getTracks().forEach((track) => track.stop());
    setStream(null);
    setCameraActive(false);
  };

  const startCamera = async () => {
    try {
      const mediaStream = await navigator.mediaDevices.getUserMedia({
        video: { facingMode },
      });
      setCameraStream(mediaStream);
      setStream(mediaStream);
      if (videoRef.current) {
        videoRef.current.srcObject = mediaStream;
        videoRef.current.onloadedmetadata = () => videoRef.current?.play();
      }
      setCameraActive(true);
    } catch (err) {
      console.error("Error accessing camera:", err);
    }
  };

  const switchCamera = () => {
    setFacingMode((prev) => (prev === "environment" ? "user" : "environment"));
  };

  useEffect(() => {
    if (cameraActive) {
      stopCamera();
      startCamera();
    }
  }, [facingMode]);

  const pathname = usePathname();
  useEffect(() => {
    stopCamera();
  }, [pathname]);

  useEffect(() => {
    return () => stopCamera();
  }, []);

  function isImageBlurred(canvas: HTMLCanvasElement, threshold = 70): boolean {
    const ctx = canvas.getContext("2d")!;
    const { width, height } = canvas;
    const imageData = ctx.getImageData(0, 0, width, height);
    const data = imageData.data;

    const gray = new Float32Array(width * height);
    for (let i = 0; i < data.length; i += 4) {
      gray[i / 4] = 0.299 * data[i] + 0.587 * data[i + 1] + 0.114 * data[i + 2];
    }

    let sum = 0;
    let sumSq = 0;
    let count = 0;

    for (let y = 1; y < height - 1; y++) {
      for (let x = 1; x < width - 1; x++) {
        const i = y * width + x;
        const lap =
          gray[i - width] +
          gray[i + width] +
          gray[i - 1] +
          gray[i + 1] -
          4 * gray[i];

        sum += lap;
        sumSq += lap * lap;
        count++;
      }
    }

    const variance = sumSq / count - (sum / count) ** 2;
    setVarianceBlur(variance);
    return variance < threshold;
  }

  const capturePhoto = () => {
    if (!videoRef.current) return;

    const video = videoRef.current;
    const canvas = document.createElement("canvas");
    const ctx = canvas.getContext("2d")!;
    canvas.width = video.videoWidth;
    canvas.height = video.videoHeight;
    ctx.drawImage(video, 0, 0, canvas.width, canvas.height);
    if (isBlurCheck && isImageBlurred(canvas, 5)) {
      setBlurPreview(canvas.toDataURL("image/jpeg"));
      setShowBlurModal(true);
      return; // ⛔ stop di sini
    }

    onCapture(canvas);
    stopCamera();
  };

  return (
    <div className="relative rounded-sm">
      <video
        ref={videoRef}
        autoPlay
        playsInline
        className={`w-full rounded-sm ${!cameraActive && "hidden"}`}
      ></video>

      {cameraActive ? (
        <>
          {children}
          <div className="absolute bottom-2 right-2 left-2 flex justify-around">
            <div
              onClick={stopCamera}
              className="w-14 h-14 rounded-full bg-gray-100/30 flex items-center justify-center"
            >
              <LuX size={36} opacity={0.3} />
            </div>
            <div
              onClick={capturePhoto}
              className="w-14 h-14 rounded-full bg-gray-100/30 flex items-center justify-center"
            >
              <LuCamera size={36} opacity={0.3} />
            </div>
            <div
              onClick={switchCamera}
              className="w-14 h-14 rounded-full bg-gray-100/30 flex items-center justify-center"
            >
              <BsArrowRepeat size={36} opacity={0.3} />
            </div>
          </div>
        </>
      ) : (
        <Button
          id="aktifkan"
          type="button"
          onClick={startCamera}
          className="w-full bg-blue-500 hover:bg-blue-600 text-white"
        >
          <LuCamera size={20} className="inline " />
          {buttonText}
        </Button>
      )}
      <Modal
        isOpen={showBlurModal}
        title="Foto Terlalu Blur"
        confirmText="Ambil Ulang"
        cancelText=""
        confirmVariant="blue-solid"
        onCancel={() => setShowBlurModal(false)}
        onConfirm={() => setShowBlurModal(false)}
      >
        <div className="flex flex-col gap-2">
          {blurPreview && (
            <Image
              src={blurPreview}
              alt={"foto-contoh"}
              width={240}
              height={200}
              className="rounded-lg object-cover w-60 h-50"
            />
          )}
        </div>
        <div>Silakan ambil ulang foto dengan posisi lebih stabil.</div>
        <div>Level Blur : {Math.floor(varianceBlur)}</div>
        
      </Modal>
    </div>
  );
}
