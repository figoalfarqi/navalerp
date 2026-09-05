/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useEffect, useRef, useState } from "react";
import { Html5Qrcode } from "html5-qrcode";
import { LuScanLine, LuX, LuFlashlight } from "react-icons/lu";
import { BsArrowRepeat } from "react-icons/bs";
import Button from "../form/Button";
import { usePathname } from "next/navigation";

interface BarcodeScannerProps {
  onScan: (value: string) => void;
  buttonText?: string;
}

export default function BarcodeScanner({
  onScan,
  buttonText = "Buka Scanner",
}: BarcodeScannerProps) {
  const scannerRef = useRef<Html5Qrcode | null>(null);
  const containerRef = useRef<HTMLDivElement>(null);
  const audioRef = useRef<HTMLAudioElement | null>(null);

  const [scannerActive, setScannerActive] = useState(false);
  const [facingMode, setFacingMode] = useState<"environment" | "user">(
    "environment",
  );
  const [torchOn, setTorchOn] = useState(false);
  const pathname = usePathname();

  const startScanner = async () => {
    setScannerActive(true);
  };

  const stopScanner = async () => {
    if (!scannerRef.current) return;

    try {
      await scannerRef.current.stop();
      await scannerRef.current.clear();
    } catch (err) {}

    scannerRef.current = null;
    setScannerActive(false);
    setTorchOn(false);
  };

  useEffect(() => {
    if (!scannerActive) return;
    if (!containerRef.current) return;

    const html5QrCode = new Html5Qrcode("barcode-reader");
    scannerRef.current = html5QrCode;

    const runScanner = async () => {
      try {
        await html5QrCode.start(
          { facingMode },
          {
            fps: 15,
            qrbox: { width: 260, height: 260 },
            aspectRatio: 1,
          },
          async (decodedText) => {
            // 🔊 Play Sound
            if (audioRef.current) {
              audioRef.current.currentTime = 0;
              audioRef.current.play();
            }

            // 📳 Vibrate
            if (navigator.vibrate) {
              navigator.vibrate(200);
            }

            onScan(decodedText);
            await stopScanner();
          },
          () => {},
        );
      } catch (err) {
        console.error("Scanner error:", err);
        setScannerActive(false);
      }
    };

    runScanner();

    return () => {
      stopScanner();
    };
  }, [scannerActive, facingMode]);

  const switchCamera = async () => {
    await stopScanner();
    setFacingMode((prev) => (prev === "environment" ? "user" : "environment"));
  };

  const toggleTorch = async () => {
    if (!scannerRef.current) return;

    try {
      await scannerRef.current.applyVideoConstraints({
        advanced: [{ torch: !torchOn } as any],
      });

      setTorchOn(!torchOn);
    } catch (err) {
    }
  };

  useEffect(() => {
    if (scannerActive) {
      startScanner();
    }
  }, [facingMode]);

  useEffect(() => {
    stopScanner();
  }, [pathname]);

  useEffect(() => {
    return () => {
      stopScanner(); // panggil saja, jangan return Promise
    };
  }, []);

  return (
    <div className="relative rounded-sm overflow-hidden">
      <audio ref={audioRef} src="/audios/scanner.mpeg" preload="auto" />

      {scannerActive ? (
        <>
          <div id="barcode-reader" ref={containerRef} className="w-full" />

          {/* 🔲 Guide Frame */}
          <div className="absolute inset-0 flex items-center justify-center pointer-events-none">
            <div className="relative w-64 h-64 border-4 border-white/80 rounded-lg">
              <div className="absolute top-0 left-0 w-full h-1 bg-green-400 animate-scanLine" />
            </div>
          </div>

          {/* Controls */}
          <div className="absolute bottom-5 left-0 right-0 flex justify-around">
            <div
              onClick={stopScanner}
              className="w-14 h-14 rounded-full bg-white/20 backdrop-blur-md flex items-center justify-center"
            >
              <LuX size={30} className="text-white" />
            </div>

            <div
              onClick={toggleTorch}
              className={`w-14 h-14 rounded-full flex items-center justify-center backdrop-blur-md ${
                torchOn ? "bg-yellow-400/80" : "bg-white/20"
              }`}
            >
              <LuFlashlight size={28} className="text-white" />
            </div>

            <div
              onClick={switchCamera}
              className="w-14 h-14 rounded-full bg-white/20 backdrop-blur-md flex items-center justify-center"
            >
              <BsArrowRepeat size={28} className="text-white" />
            </div>
          </div>
        </>
      ) : (
        <Button
          id="open-scanner"
          type="button"
          onClick={startScanner}
          className="w-full bg-blue-500 hover:bg-blue-600 text-white"
        >
          <LuScanLine size={20} className="inline mr-2" />
          {buttonText}
        </Button>
      )}

      {/* Animation Style */}
      <style jsx>{`
        @keyframes scanLine {
          0% {
            transform: translateY(0);
          }
          100% {
            transform: translateY(250px);
          }
        }
        .animate-scanLine {
          animation: scanLine 2s linear infinite;
        }
      `}</style>
    </div>
  );
}
