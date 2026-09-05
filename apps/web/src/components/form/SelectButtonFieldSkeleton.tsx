"use client";

interface SelectButtonFieldSkeletonProps {
  label?: string;
  count?: number; // jumlah tombol skeleton
  layout?: "horizontal" | "vertical";
  required?: boolean;
  labelClassName?: string;
}

export default function SelectButtonFieldSkeleton({
  label,
  count = 3,
  layout = "horizontal",
  labelClassName = "",
  required = false,
}: SelectButtonFieldSkeletonProps) {
  const baseLabel = "mb-2 font-medium text-gray-700";
  return (
    <div className="flex flex-col animate-pulse">
      {/* Label skeleton */}
      {/* <div className="h-4 w-24 bg-gray-200 rounded mb-3" /> */}
      {label && (
        <label className={`${baseLabel} ${labelClassName}`}>
          {label} {required && <span className="text-red-500">*</span>}
        </label>
      )}
      {/* Tombol skeleton */}
      <div
        className={`flex ${
          layout === "horizontal"
            ? "flex-row flex-wrap gap-2"
            : "flex-col gap-2"
        }`}
      >
        {Array.from({ length: count }).map((_, i) => (
          <div
            key={i}
            className="h-9 w-24 bg-gray-200 rounded-md border border-gray-300"
          />
        ))}
      </div>
    </div>
  );
}
