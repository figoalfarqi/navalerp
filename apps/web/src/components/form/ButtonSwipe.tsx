"use client";

import React, { useEffect, useRef, useState } from "react";
import { usePointerGesture } from "@/hooks/usePointerGesture";

interface ButtonSwipeProps {
  id: string;
  type?: "button" | "submit" | "reset";
  variant?:
    | "blue-solid"
    | "red-solid"
    | "yellow-solid"
    | "green-solid"
    | "purple-solid"
    | "gray-solid"
    | "blue-dkl";
  size?: "2xs" | "xs" | "sm" | "md" | "lg" | "xl" | "2xl";
  className?: string;
  wrapperClassName?: string;
  disabled?: boolean;
  children?: React.ReactNode;
  /** dipanggil saat swipe sukses penuh */
  onSwipeSuccess?: () => void;
  /** fallback jika hanya diklik */
  onClick?: () => void;
  /** Optional: Parent bisa akses fungsi internal seperti setSwipeProgress */
  onReady?: (controls: { setSwipeProgress: (value: number) => void }) => void;
}

export default function ButtonSwipe({
  id,
  type = "button",
  variant = "blue-solid",
  size = "md",
  className = "",
  wrapperClassName = "",
  disabled = false,
  onSwipeSuccess,
  onClick,
  onReady,
  children,
}: ButtonSwipeProps) {
  const containerRef = useRef<HTMLDivElement>(null);
  const [knobTravelWidth, setKnobTravelWidth] = useState(250);

  // 🧩 Hitung otomatis width container
  useEffect(() => {
    const measureWidth = () => {
      if (containerRef.current) {
        const width = containerRef.current.offsetWidth;
        // kurangi sedikit supaya knob tidak keluar sisi kanan
        setKnobTravelWidth(width - 40);
      }
    };

    measureWidth();
    window.addEventListener("resize", measureWidth);
    return () => window.removeEventListener("resize", measureWidth);
  }, []);

  const swipeTh = knobTravelWidth / 100;

  const {
    swipeProgress,
    setSwipeProgress,
    handlePressStart,
    handleMove,
    handlePressEnd,
  } = usePointerGesture({
    onSwipeSuccess,
    onClick,
    swipeThreshold: knobTravelWidth,
  });

  // 🧠 Kirim kontrol ke parent saat mount
  useEffect(() => {
    if (onReady) onReady({ setSwipeProgress });
  }, [onReady, setSwipeProgress]);

  // 🔹 Ukuran tinggi per size (menyesuaikan Button)
  const sizeClasses: Record<string, string> = {
    "2xs": "h-6 text-xs",
    xs: "h-7 text-xs",
    sm: "h-8 text-sm",
    md: "h-10 text-base",
    lg: "h-11 text-lg",
    xl: "h-12 text-xl",
    "2xl": "h-14 text-2xl",
  };

  // 🔹 Warna background & teks — hanya solid variants
  const variantClasses: Record<string, string> = {
    "blue-solid": "bg-blue-500 text-white",
    "red-solid": "bg-red-500 text-white",
    "yellow-solid": "bg-yellow-500 text-white",
    "green-solid": "bg-green-500 text-white",
    "purple-solid": "bg-purple-500 text-white",
    "gray-solid": "bg-gray-500 text-white",
    "blue-dkl": "bg-[#0a2540] text-white",
  };

  const baseButtonClass =
    "relative overflow-hidden select-none cursor-pointer inline-flex items-center justify-center font-medium rounded-md focus:outline-none disabled:opacity-50 disabled:cursor-not-allowed transition duration-200 w-full";

  const getClientX = (e: React.PointerEvent | React.TouchEvent): number => {
    if ("clientX" in e) return e.clientX; // PointerEvent
    if (e.touches && e.touches.length > 0) return e.touches[0].clientX; // TouchEvent
    return 0;
  };

  // 🔹 Pointer handlers
  const handlePointerDown = (e: React.PointerEvent | React.TouchEvent) => {
    if (disabled) return;
    const clientX = getClientX(e);
    handlePressStart(0, clientX);
  };

  const handlePointerMove = (e: React.PointerEvent | React.TouchEvent) => {
    if (disabled) return;
    const clientX = getClientX(e);
    handleMove(0, clientX);
  };

  const handlePointerUp = () => {
    if (disabled) return;
    handlePressEnd(0);
  };
  
  useEffect(() => {
    if (disabled) {
      setSwipeProgress(0);
    }
  }, [disabled]);

  return (
    <div className={`inline-flex w-full ${wrapperClassName}`}>
      <div
        id={id}
        ref={containerRef}
        role={type}
        onPointerDown={handlePointerDown}
        onPointerMove={handlePointerMove}
        onPointerUp={handlePointerUp}
        onTouchStart={handlePointerDown}
        onTouchMove={handlePointerMove}
        onTouchEnd={handlePointerUp}
        className={`${baseButtonClass} ${variantClasses[variant]} ${sizeClasses[size]} ${className}`}
      >
        {/* 🔹 Progress background */}
        <div
          className="absolute left-0 top-0 h-full w-full ease-linear rounded-md"
          style={{ background: `rgba(255,255,255,${swipeProgress / 200})` }}
        ></div>

        {/* 🔹 Label (di tengah) */}
        <div className="relative z-2 pointer-events-none">{children}</div>

        {/* 🔹 Knob geser */}
        <div
          className={`absolute top-1 left-3 z-3 bg-white shadow-sm ${
            size === "2xs"
              ? "h-4 w-4 rounded-sm"
              : size === "xs"
              ? "h-5 w-5 rounded-sm"
              : size === "sm"
              ? "h-6 w-6 rounded"
              : size === "md"
              ? "h-8 w-8 rounded"
              : size === "lg"
              ? "h-9 w-9 rounded-md"
              : size === "xl"
              ? "h-10 w-10 rounded-lg"
              : "h-12 w-12 rounded-lg"
          }`}
          style={{
            transform: `translateX(calc(${
              swipeProgress * swipeTh
            }px - 0.5rem))`,
          }}
        ></div>
      </div>
    </div>
  );
}
