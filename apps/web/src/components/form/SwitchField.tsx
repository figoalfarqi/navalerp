"use client";

import { useEffect } from "react";

interface SwitchFieldProps {
  id: string;
  label: string;
  value: number | null | undefined;
  className?: string;
  wrapperClassName?: string;
  labelClassName?: string;
  switchText?: { trueText: string; falseText: string };
  onChange: (value: number) => void;
  required?: boolean;
  disabled?: boolean;
  error?: string;
}

export default function SwitchField({
  id,
  label,
  value = 0,
  className = "",
  wrapperClassName = "",
  labelClassName = "",
  switchText = { trueText: "Aktif", falseText: "Tidak Aktif" },
  onChange,
  required = false,
  disabled = false,
  error = "",
}: SwitchFieldProps) {
  const baseWrapperClass = "flex flex-col";
  const baseLabelClass = "mb-1 font-medium text-gray-700";
  const baseSwitchWrapperClass = "flex items-center gap-3";
  const baseSwitchClass =
    "relative inline-flex h-8 w-18 items-center rounded-full transition-colors duration-300 focus:outline-none cursor-pointer";
  const baseCircleClass =
    "inline-block h-6 w-9 transform rounded-full bg-white transition-transform duration-300";

  useEffect(() => {
    if (value == undefined || value == null) {
      setTimeout(() => {
        onChange(0);
      }, 100);
    }
  }, [value]);


  return (
    <div className={`${baseWrapperClass} ${wrapperClassName}`}>
      <label htmlFor={id} className={`${baseLabelClass} ${labelClassName}`}>
        {label} {required && <span className="text-red-500">*</span>}
      </label>

      <div className={baseSwitchWrapperClass}>
        <button
          id={id}
          type="button"
          role="switch"
          aria-checked={value == 1 ? true : false}
          disabled={disabled}
          onClick={() => {
            value == 1 ? onChange(0) : onChange(1);
          }}
          className={`${baseSwitchClass} ${
            value ? "bg-blue-500" : "bg-gray-300"
          } ${disabled ? "opacity-50 cursor-not-allowed" : ""} ${className}`}
        >
          <span
            className={`${baseCircleClass} ${
              value ? "translate-x-8" : "translate-x-1.25"
            }`}
          />
        </button>
        <span>{value ? switchText.trueText : switchText.falseText}</span>
      </div>

      {error && <p className="text-red-500 text-sm mt-1">{error}</p>}
    </div>
  );
}
