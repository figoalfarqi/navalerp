"use client";

import React from "react";
import Image from "next/image";

interface DesktopPageLoaderProps {
  message?: string;
  subMessage?: string;
  fullScreen?: boolean;
}

const DesktopPageLoader: React.FC<DesktopPageLoaderProps> = ({
  message = "Memuat Modul & Formulir Sistem...",
  subMessage = "Menginisialisasi sesi terenkripsi • Sinkronisasi pangkalan data",
  fullScreen = false,
}) => {
  const containerClasses = fullScreen
    ? "fixed inset-0 z-50 flex items-center justify-center bg-[#040e1c]/90 backdrop-blur-md p-4"
    : "flex items-center justify-center min-h-[70vh] w-full p-6";

  return (
    <div className={containerClasses}>
      <div className="relative overflow-hidden rounded-2xl bg-gradient-to-b from-[#081d38] via-[#06162d] to-[#040e1c] border border-cyan-500/30 p-8 sm:p-10 shadow-[0_20px_50px_rgba(0,15,35,0.7)] text-center max-w-md w-full mx-auto">
        {/* Ambient Top Glow */}
        <div className="absolute -top-16 left-1/2 -translate-x-1/2 w-48 h-48 bg-cyan-500/15 rounded-full blur-2xl pointer-events-none" />

        {/* Tactical Radar / Sonar Container */}
        <div className="relative mx-auto w-28 h-28 flex items-center justify-center mb-6">
          {/* Sonar wave pulse */}
          <div className="absolute inset-0 rounded-full border border-cyan-400/25 animate-ping pointer-events-none" />

          {/* Outer dashed radar ring */}
          <div className="absolute inset-0 rounded-full border-2 border-dashed border-cyan-500/30 animate-[spin_10s_linear_infinite]" />

          {/* Middle clockwise scanning ring */}
          <div className="absolute inset-1.5 rounded-full border-2 border-transparent border-t-cyan-400 border-r-sky-400 animate-[spin_1.5s_linear_infinite]" />

          {/* Inner counter-clockwise ring */}
          <div className="absolute inset-3.5 rounded-full border-2 border-transparent border-b-blue-400 border-l-cyan-300 animate-[spin_2s_linear_infinite_reverse]" />

          {/* Center Logo Hub */}
          <div className="relative w-14 h-14 rounded-full bg-[#030b14] border border-cyan-500/50 flex items-center justify-center shadow-lg p-2.5">
            <Image
              src="/logo.png"
              alt="Naval ERP"
              width={40}
              height={40}
              priority
              className="object-contain animate-pulse"
            />
          </div>
        </div>

        {/* Tactical Badge */}
        <div className="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-cyan-950/70 border border-cyan-500/40 text-[10px] font-mono font-bold tracking-widest text-cyan-300 uppercase shadow-inner mb-3">
          <span className="relative flex h-2 w-2">
            <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-cyan-400 opacity-75" />
            <span className="relative inline-flex rounded-full h-2 w-2 bg-emerald-400" />
          </span>
          TNI ANGKATAN LAUT • NAVAL ERP
        </div>

        {/* Headline / Message */}
        <h3 className="text-base sm:text-lg font-bold text-white tracking-wide font-mono">
          {message}
        </h3>

        {/* Subtitle */}
        <p className="mt-1.5 text-xs text-slate-300 font-sans leading-relaxed">
          {subMessage}
        </p>

        {/* Sleek Progress Shimmer Track */}
        <div className="w-full bg-[#030a14] rounded-full h-1.5 mt-5 overflow-hidden border border-blue-900/60 relative">
          <div className="shimmer-bar h-full rounded-full bg-gradient-to-r from-transparent via-cyan-400 to-transparent w-1/2" />
        </div>

        {/* Footer Security Tag */}
        <div className="mt-6 pt-4 border-t border-blue-900/40 flex items-center justify-between text-[10px] font-mono text-slate-400">
          <span className="flex items-center gap-1">
            <span className="w-1.5 h-1.5 rounded-full bg-cyan-400" />
            SECURE TLS 1.3 • AES-256
          </span>
          <span className="text-cyan-400 font-semibold tracking-wider">
            JALESVEVA JAYAMAHE
          </span>
        </div>
      </div>

      <style jsx>{`
        @keyframes shimmer-move {
          0% {
            transform: translateX(-100%);
          }
          100% {
            transform: translateX(250%);
          }
        }
        .shimmer-bar {
          animation: shimmer-move 1.6s ease-in-out infinite;
        }
      `}</style>
    </div>
  );
};

export default DesktopPageLoader;
