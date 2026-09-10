"use client";

import Image from "next/image";

interface AdminRecordStateProps {
  error?: string;
  message?: string;
}

export default function AdminRecordState({
  error,
  message = "Memuat Data Formulir...",
}: AdminRecordStateProps) {
  if (error) {
    return (
      <div className="rounded-2xl border border-rose-800/40 bg-rose-950/20 p-6 text-sm text-rose-300 shadow-lg">
        <div className="flex items-center gap-3">
          <span className="flex h-8 w-8 items-center justify-center rounded-lg bg-rose-900/60 border border-rose-700 text-rose-200 font-bold text-base">
            ⚠
          </span>
          <div>
            <h4 className="font-semibold text-rose-200 font-mono text-sm">Gagal Memuat Data</h4>
            <p className="mt-0.5 text-xs text-rose-300/80">{error}</p>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="grid min-h-[360px] place-items-center rounded-2xl border border-blue-900/40 bg-gradient-to-b from-[#081d38]/80 to-[#040e1c]/80 backdrop-blur-sm p-8 shadow-xl text-center">
      <div className="flex flex-col items-center max-w-sm w-full">
        {/* Radar Spinner Hub */}
        <div className="relative w-20 h-20 flex items-center justify-center mb-5">
          <div className="absolute inset-0 rounded-full border border-cyan-400/20 animate-ping" />
          <div className="absolute inset-0 rounded-full border-2 border-dashed border-cyan-500/30 animate-[spin_8s_linear_infinite]" />
          <div className="absolute inset-1 rounded-full border-2 border-transparent border-t-cyan-400 border-r-sky-400 animate-[spin_1.5s_linear_infinite]" />
          <div className="w-10 h-10 rounded-full bg-[#030b14] border border-cyan-500/50 flex items-center justify-center shadow p-2">
            <Image
              src="/logo.webp"
              alt="Naval ERP"
              width={26}
              height={26}
              className="object-contain animate-pulse"
            />
          </div>
        </div>

        {/* Badge */}
        <span className="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full bg-cyan-950/70 border border-cyan-500/40 text-[9px] font-mono font-bold tracking-widest text-cyan-300 uppercase mb-2.5">
          <span className="w-1.5 h-1.5 rounded-full bg-cyan-400 animate-pulse" />
          TNI AL DATA RECORD
        </span>

        {/* Status Text */}
        <h4 className="text-sm font-semibold text-white font-mono tracking-wide">
          {message}
        </h4>
        <p className="mt-1 text-xs text-slate-400 font-sans">
          Mengambil data dari pangkalan data pusat...
        </p>

        {/* Shimmer Bar */}
        <div className="w-48 bg-[#030a14] rounded-full h-1 mt-4 overflow-hidden border border-blue-900/60 relative">
          <div className="shimmer-bar-record h-full rounded-full bg-gradient-to-r from-transparent via-cyan-400 to-transparent w-1/2" />
        </div>
      </div>

      <style jsx>{`
        @keyframes shimmer-move-rec {
          0% {
            transform: translateX(-100%);
          }
          100% {
            transform: translateX(250%);
          }
        }
        .shimmer-bar-record {
          animation: shimmer-move-rec 1.5s ease-in-out infinite;
        }
      `}</style>
    </div>
  );
}
