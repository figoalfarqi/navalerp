"use client";

import { RadioVariant, VariantClassesRadioMap, VariantSelectedBgRadioMap } from "@/consta/VariantClassesRadio";

export interface RadioOption {
  label: string;
  value: string | number;
  variant?: RadioVariant;
}

interface RadioFieldProps {
  id: string;
  label: string;
  value: string | number;
  options: RadioOption[];
  onChange: (value: string | number) => void;
  className?: string;
  labelClassName?: string;
  required?: boolean;
  error?: string;
}

export default function RadioField({
  id,
  label,
  value,
  options,
  onChange,
  className = "",
  labelClassName = "",
  required = false,
  error = "",
}: RadioFieldProps) {
  return (
    <div id={id} className={`flex flex-col gap-1 ${className}`}>
      {label && (
        <label className={`font-medium text-gray-700 ${labelClassName}`}>
          {label} {required && <span className="text-red-500">*</span>}
        </label>
      )}
      <div className="flex flex-wrap gap-3">
        {options.map((opt) => {
          const variant = opt.variant ?? "gray-outline";
          const baseStyle = VariantClassesRadioMap[variant];

          const selectedStyle =
            value === opt.value ? `${VariantSelectedBgRadioMap[variant]}` : "";

          return (
            <div
              key={opt.value}
              onClick={() => onChange(opt.value)}
              className={`
                px-4 py-2 rounded-md cursor-pointer select-none
                flex items-center gap-2 transition-all
                ${baseStyle} 
                ${selectedStyle}
              `}
            >
              {/* Custom radio circle */}
              <div
                className={`
                  w-4 h-4 rounded-full border 
                  flex items-center justify-center
                  ${
                    value === opt.value
                      ? "bg-current border-current"
                      : "border-current"
                  }
                `}
              />

              <span>{opt.label}</span>
            </div>
          );
        })}
      </div>

      {error && <p className="text-red-500 text-sm">{error}</p>}
    </div>
  );
}
