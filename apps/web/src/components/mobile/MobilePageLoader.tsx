"use client";
import Image from "next/image";
import { useEffect, useState } from "react";

export default function MobilePageLoader() {
  const [dots, setDots] = useState(1);

  useEffect(() => {
    const id = setInterval(() => {
      setDots((d) => (d % 3) + 1); // 1 -> 2 -> 3 -> 1 ...
    }, 500);
    return () => clearInterval(id);
  }, []);

  return (
    <div className="w-full h-[calc(100vh-60px)] flex items-center justify-center">
      <div className="flex flex-col items-center gap-4">
        {/* Bungkus gambar agar animasi bekerja */}
        <div className="animate-scale-pulse transform">
          <Image
            src={"/logo-pml.png"}
            alt="loading"
            width={180}
            height={120}
            className="w-[180px] h-[120px] object-cover"
          />
        </div>

        {/* <div className="flex items-center justify-center w-full h-full">
          <div className="w-12 h-12 border-4 border-blue-500 border-t-transparent rounded-full animate-spin">
            
          </div>
        </div> */}

        {/* Label "loading ..." dengan titik 1..3 */}
        <div className="font-medium text-gray-600 tracking-wide">
          Loading{".".repeat(dots)}
        </div>
      </div>

      {/* Keyframes khusus */}
      <style jsx>{`
        @keyframes scale-pulse {
          0% {
            transform: scale(0.95);
          }
          50% {
            transform: scale(1.1);
          }
          100% {
            transform: scale(0.95);
          }
        }
        .animate-scale-pulse {
          animation: scale-pulse 1.2s ease-in-out infinite;
        }
      `}</style>
    </div>
  );
}
