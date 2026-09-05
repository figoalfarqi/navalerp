"use client";

interface TextFieldSkeletonProps {
  className?: string;
  wrapperClassName?: string;
  labelClassName?: string;
  isLabel?: boolean;
}

export default function TextFieldSkeleton({
  className = "",
  wrapperClassName = "",
  labelClassName = "",
  isLabel = false,
}: TextFieldSkeletonProps) {
  const baseWrapperClass = "flex flex-col animate-pulse";
  const baseLabelClass = "mb-[5px] h-6 bg-gray-300 rounded w-24";
  const baseInputClass = "h-10 bg-gray-300 rounded";

  return (
    <div className={`${baseWrapperClass} ${wrapperClassName}`}>
      {/* Skeleton Label */}
      {isLabel && <div className={`${baseLabelClass} ${labelClassName}`} />}
      {/* Skeleton Input */}
      <div className={`${baseInputClass} ${className}`} />
    </div>
  );
}
