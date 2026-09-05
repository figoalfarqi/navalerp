"use client";

import { ChangeEvent, ReactNode, useEffect, useState } from "react";
import { TfiEmail } from "react-icons/tfi";
import { HiEye, HiEyeOff } from "react-icons/hi";
interface TextFieldProps {
  id: string;
  label?: string;
  type?: "text" | "number" | "email" | "password";
  value?: string | number;
  className?: string;
  wrapperClassName?: string;
  labelClassName?: string;
  columnLength?: number;
  onChange: (value: string | number) => void;
  required?: boolean;
  disabled?: boolean;
  placeholder?: string;
  error?: string;
  icon?: ReactNode;
  border?: string;
  isPasswordShowable?: boolean;
  onKeyDown?: React.KeyboardEventHandler<
    HTMLInputElement | HTMLTextAreaElement
  >;
}

export default function TextField({
  id,
  label,
  type = "text",
  value,
  className = "",
  wrapperClassName = "",
  labelClassName = "",
  columnLength,
  onChange,
  required = false,
  disabled = false,
  placeholder = "",
  error = "",
  icon = "",
  border = "",
  isPasswordShowable = false,
  onKeyDown,
}: TextFieldProps) {
  const baseWrapperClass = "flex flex-col";
  const baseLabelClass = "mb-1 font-medium text-gray-700";
  const baseInputClass =
    "px-3 py-2 border rounded-sm focus:outline-none focus:border-blue-500 focus:shadow-[0_0_6px_rgba(59,130,246,0.3)] disabled:bg-gray-100 disabled:text-gray-400";
  const [displayValue, setDisplayValue] = useState<string>(
    value?.toString() ?? "",
  );

  // console.log("valueeeeeeeeeeeeeeeeeeeeeeeeeee", value)

  const [showPassword, setShowPassword] = useState(false);

  const changeVal = (val: string) => {
    if (type === "number") {
      // Terima koma maupun titik sebagai pemisah desimal. Jika keduanya ada,
      // pemisah yang paling kanan dianggap desimal dan yang lain ribuan.
      const cleaned = val.replace(/[^0-9,.-]/g, "");
      const lastComma = cleaned.lastIndexOf(",");
      const lastDot = cleaned.lastIndexOf(".");
      const decimalIndex = Math.max(lastComma, lastDot);
      const hasBothSeparators = lastComma >= 0 && lastDot >= 0;
      const decimalSeparator =
        decimalIndex >= 0 ? cleaned[decimalIndex] : undefined;
      let numeric = cleaned;

      if (hasBothSeparators && decimalSeparator) {
        const thousandsSeparator = decimalSeparator === "," ? "." : ",";
        numeric = numeric.replaceAll(thousandsSeparator, "");
      }
      const normalizedDecimalIndex = decimalSeparator
        ? numeric.lastIndexOf(decimalSeparator)
        : -1;
      if (decimalSeparator === ",") {
        numeric =
          numeric.slice(0, normalizedDecimalIndex).replaceAll(",", "") +
          "." +
          numeric.slice(normalizedDecimalIndex + 1).replaceAll(",", "");
      } else if (decimalSeparator === ".") {
        numeric =
          numeric.slice(0, normalizedDecimalIndex).replaceAll(".", "") +
          "." +
          numeric.slice(normalizedDecimalIndex + 1).replaceAll(".", "");
      }

      // Konversi ke number
      const numberVal = numeric === "" ? "" : parseFloat(numeric);

      // Format string untuk ditampilkan
      let formatted = "";
      if (numeric === "-" || numeric === "--") {
        onChange("");
        setDisplayValue("-");
        return;
      } else if (
        numeric !== "" &&
        typeof numberVal == "number" &&
        !isNaN(numberVal)
      ) {
        const [intPart, decPart] = numeric.split(".");
        formatted = parseInt(intPart).toLocaleString("id-ID");
        if (decPart !== undefined) {
          formatted += "," + decPart;
        }
        // Kirim number murni ke parent
        onChange(numberVal);
        // Ganti value input ke formatted string
        setDisplayValue(formatted);
      } else {
        onChange(numberVal);
        setDisplayValue(formatted);
      }
    } else {
      onChange(val);
      setDisplayValue(val);
    }
  };
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const val = e.target.value;
    changeVal(val);
  };

  // useEffect(() => {
  //   const val = value.toString();
  //   changeVal(val);
  // }, [value]);

  useEffect(() => {
    if (type !== "number") {
      setDisplayValue(value?.toString() ?? "");
      return;
    }

    if (value === null || value === undefined || value === "") {
      setDisplayValue("");
      return;
    }

    const [intPart, decPart] = value.toString().split(".");
    let formatted = Number(intPart).toLocaleString("id-ID");

    if (decPart) {
      formatted += "," + decPart;
    }

    setDisplayValue(formatted);
  }, [value, type]);

  const inputType =
    type === "password" && isPasswordShowable
      ? showPassword
        ? "text"
        : "password"
      : type === "password"
        ? "password"
        : "text";
  return (
    <div
      className={`${baseWrapperClass} ${wrapperClassName}`}
      style={{ width: columnLength }}
    >
      {label && (
        <label htmlFor={id} className={`${baseLabelClass} ${labelClassName}`}>
          {label} {required && <span className="text-red-500">*</span>}
        </label>
      )}
      <div className="relative">
        {icon}
        <input
          id={id}
          type={inputType}
          inputMode={type === "number" ? "decimal" : undefined}
          value={displayValue}
          onChange={handleChange}
          placeholder={placeholder}
          disabled={disabled}
          className={`${baseInputClass} ${border} ${className} ${
            type === "password" && isPasswordShowable ? "pr-10" : ""
          }`}
          onKeyDown={onKeyDown}
        />
        {type === "email" && (
          <div className="absolute right-3 top-1/2 transform -translate-y-1/2">
            <TfiEmail size={24} />
          </div>
        )}
        {type === "password" && isPasswordShowable && (
          <button
            type="button"
            aria-label={
              showPassword ? "Sembunyikan password" : "Tampilkan password"
            }
            title={showPassword ? "Sembunyikan password" : "Tampilkan password"}
            className="absolute right-3 top-1/2 transform -translate-y-1/2 cursor-pointer text-gray-500"
            onClick={() => setShowPassword(!showPassword)}
          >
            {showPassword ? <HiEyeOff size={24} /> : <HiEye size={24} />}
          </button>
        )}
      </div>
      {error && <p className="text-red-500 text-sm mt-1">{error}</p>}
    </div>
  );
}
