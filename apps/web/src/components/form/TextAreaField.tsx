"use client";

import { Ref, useState } from "react";

interface TextAreaFieldProps {
  id: string;
  label?: string;
  ref?: Ref<HTMLTextAreaElement>;
  value: string;
  className?: string;
  wrapperClassName?: string;
  labelClassName?: string;
  onChange: (value: string) => void;
  required?: boolean;
  disabled?: boolean;
  placeholder?: string;
  error?: string;
  rows?: number;
  onKeyDown?: React.KeyboardEventHandler<
    HTMLInputElement | HTMLTextAreaElement
  >;
}

export default function TextAreaField({
  id,
  label,
  ref,
  value,
  className = "",
  wrapperClassName = "",
  labelClassName = "",
  onChange,
  required = false,
  disabled = false,
  placeholder = "",
  error = "",
  rows = 1, // default kecil
}: TextAreaFieldProps) {
  const [isFocused, setIsFocused] = useState(false);

  const baseWrapperClass = "flex flex-col mb-4";
  const baseLabelClass = "mb-1 font-medium text-gray-700";
  const baseTextareaClass =
    "app-scrollbar px-3 py-2 border rounded-sm focus:outline-none focus:border-blue-500 focus:shadow-[0_0_6px_rgba(59,130,246,0.3)] disabled:bg-gray-100 disabled:text-gray-400 resize-none transition-all duration-300 ease-in-out";

  return (
    <div className={`${baseWrapperClass} ${wrapperClassName}`}>
      {label && (
        <label htmlFor={id} className={`${baseLabelClass} ${labelClassName}`}>
          {label} {required && <span className="text-red-500">*</span>}
        </label>
      )}
      <textarea
        ref={ref}
        id={id}
        value={value}
        onChange={(e) => onChange(e.target.value)}
        onFocus={() => setIsFocused(true)}
        onBlur={() => setIsFocused(false)}
        disabled={disabled}
        placeholder={placeholder}
        rows={isFocused ? 4 : rows} // membesar saat fokus
        className={`${baseTextareaClass} ${className}`}
        style={{
          height: isFocused ? "100px" : "41.6px", // default tinggi textfield ~38px
        }}
      />
      {error && <p className="text-red-500 text-sm mt-1">{error}</p>}
    </div>
  );
}
