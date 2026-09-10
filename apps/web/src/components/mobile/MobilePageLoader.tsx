"use client";
import Image from "next/image";
import { useEffect, useState } from "react";

export default function MobilePageLoader() {
  const [dots, setDots] = useState(1);

  useEffect(() => {
    const id = setInterval(() => {
      setDots((d) => (d % 3) + 1);
    }, 450);
    return () => clearInterval(id);
  }, []);

  return (
    <div className="w-full min-h-[calc(100vh-80px)] flex items-center justify-center p-4">
      <div className="relative overflow-hidden rounded-2xl bg-gradient-to-b from-[#081d38] to-[#040e1c] border border-cyan-500/30 p-6 shadow-2xl text-center max-w-xs w-full mx-auto flex flex-col items-center">
        {/* Radar Hub */}
        <div className="relative w-24 h-24 flex items-center justify-center mb-4">
          <div className="absolute inset-0 rounded-full border border-cyan-400/25 animate-ping pointer-events-none" />
          <div className="absolute inset-0 rounded-full border-2 border-dashed border-cyan-500/30 animate-[spin_8s_linear_infinite]" />
          <div className="absolute inset-1 rounded-full border-2 border-transparent border-t-cyan-400 border-r-sky-400 animate-[spin_1.5s_linear_infinite]" />
          <div className="relative w-12 h-12 rounded-full bg-[#030b14] border border-cyan-500/50 flex items-center justify-center shadow-lg p-2">
            <Image
              src="/logo.webp"
              alt="Naval ERP"
              width={34}
              height={34}
              priority
              className="object-contain animate-pulse"
            />
          </div>
        </div>

        {/* Badge */}
        <span className="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full bg-cyan-950/70 border border-cyan-500/40 text-[9px] font-mono font-bold tracking-widest text-cyan-300 uppercase mb-2">
          <span className="w-1.5 h-1.5 rounded-full bg-cyan-400 animate-pulse" />
          TNI AL MOBILE
        </span>

        {/* Label */}
        <div className="font-mono text-sm font-semibold text-white tracking-wide">
          Memuat Sistem{".".repeat(dots)}
        </div>
        <p className="text-[11px] text-slate-400 mt-1">
          Sinkronisasi data terenkripsi
        </p>

        {/* Shimmer Bar */}
        <div className="w-full bg-[#030a14] rounded-full h-1 mt-4 overflow-hidden border border-blue-900/60 relative">
          <div className="shimmer-mobile h-full rounded-full bg-gradient-to-r from-transparent via-cyan-400 to-transparent w-1/2" />
        </div>
      </div>

      <style jsx>{`
        @keyframes shimmer-m {
          0% {
            transform: translateX(-100%);
          }
          100% {
            transform: translateX(250%);
          }
        }
        .shimmer-mobile {
          animation: shimmer-m 1.5s ease-in-out infinite;
        }
      `}</style>
    </div>
  );
}
