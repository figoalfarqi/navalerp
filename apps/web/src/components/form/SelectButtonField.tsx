"use client";

import { useState } from "react";
import { FaCheck } from "@/components/icons";

interface SelectOption {
  label: string;
  value: string | number;
}

interface SelectButtonFieldProps {
  id: string;
  label?: string;
  value: string | number;
  options: SelectOption[];
  onChange: (value: string | number | null) => void;
  required?: boolean;
  disabled?: boolean;
  error?: string;
  className?: string;
  labelClassName?: string;
  layout?: "horizontal" | "vertical"; // bentuk layout pilihan
}

export default function SelectButtonField({
  id,
  label,
  value,
  options,
  onChange,
  required = false,
  disabled = false,
  error = "",
  className = "",
  labelClassName = "",
  layout = "horizontal",
}: SelectButtonFieldProps) {
  const [selected, setSelected] = useState<string | number | null>(value);

  const handleSelect = (val: string | number) => {
    if (disabled) return;
    const newValue = val === selected ? null : val; // klik lagi untuk deselect
    setSelected(newValue);
    onChange(newValue);
  };

  const baseWrapper = "flex flex-col";
  const baseLabel = "mb-2 font-medium text-gray-700";
  const baseButton =
    "py-2 text-sm rounded-md border transition-all duration-200";

  return (
    <div className={`${baseWrapper} ${className}`}>
      {label && (
        <label htmlFor={id} className={`${baseLabel} ${labelClassName}`}>
          {label} {required && <span className="text-red-500">*</span>}
        </label>
      )}

      <div
        className={`flex ${
          layout === "horizontal"
            ? "flex-row flex-wrap gap-2"
            : "flex-col gap-2"
        }`}
      >
        {options.map((opt) => {
          const isActive = selected === opt.value;
          return (
            <button
              key={opt.value}
              type="button"
              onClick={() => handleSelect(opt.value)}
              disabled={disabled}
              className={`${baseButton} ${
                isActive
                  ? "pl-2 pr-4 bg-blue-500 text-white border-blue-500"
                  : "px-[18px]  bg-white text-gray-700 border-gray-300 hover:bg-blue-50"
              } ${disabled && "opacity-50 cursor-not-allowed"}`}
            >
              <div className="flex items-center gap-[2px]">
                <FaCheck
                  size={12}
                  className={`${
                    isActive ? "max-w-3" : "max-w-0"
                  } transition-all duration-200`}
                />
                {opt.label}
              </div>
            </button>
          );
        })}
      </div>

      {error && <p className="text-red-500 text-sm mt-1">{error}</p>}
    </div>
  );
}
