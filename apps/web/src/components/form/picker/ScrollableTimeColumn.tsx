"use client";
import { useEffect, useRef } from "react";

interface ScrollableTimeColumnProps {
  values: number[];
  selected: number | null;
  onSelect: (value: number) => void;
}

export function ScrollableTimeColumn({
  values,
  selected,
  onSelect,
}: ScrollableTimeColumnProps) {
  const containerRef = useRef<HTMLDivElement>(null);
  const itemRefs = useRef<Map<number, HTMLButtonElement>>(new Map());

  // =========================
  // AUTO SCROLL TO SELECTED
  // =========================
  useEffect(() => {
    if (selected === null) return;

    const el = itemRefs.current.get(selected);
    el?.scrollIntoView({
      block: "center",
      behavior: "smooth",
    });
  }, [selected]);

  // =========================
  // KEYBOARD NAVIGATION
  // =========================
  const handleKeyDown = (e: React.KeyboardEvent<HTMLDivElement>) => {
    if (selected === null) return;

    const currentIndex = values.indexOf(selected);
    if (currentIndex === -1) return;

    if (e.key === "ArrowUp") {
      e.preventDefault();
      const prev = values[Math.max(0, currentIndex - 1)];
      onSelect(prev);
    }

    if (e.key === "ArrowDown") {
      e.preventDefault();
      const next = values[Math.min(values.length - 1, currentIndex + 1)];
      onSelect(next);
    }
  };

  return (
    <div
      ref={containerRef}
      tabIndex={0}
      onKeyDown={handleKeyDown}
      className="h-32 w-14 overflow-y-auto text-center snap-y snap-mandatory 
                 bg-linear-to-b from-gray-200 via-transparent to-gray-200 
                 scrollbar-none outline-none focus:ring-2 focus:ring-blue-400"
    >
      {values.map((val) => (
        <button
          key={val}
          ref={(el) => {
            if (el) itemRefs.current.set(val, el);
          }}
          type="button"
          className={`block w-full py-1 cursor-pointer ${
            selected === val
              ? "bg-blue-500 text-white font-semibold"
              : "hover:bg-blue-500/30"
          }`}
          onClick={() => onSelect(val)}
        >
          {val.toString().padStart(2, "0")}
        </button>
      ))}
    </div>
  );
}
