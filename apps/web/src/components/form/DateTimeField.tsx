"use client";

import { useState, useRef, useEffect } from "react";
import { createPortal } from "react-dom";
import { formatDateTime, toLocalISOString } from "@/utils/dateTime";
import DatePicker from "./picker/DatePicker";
import TimePicker from "./picker/TimePicker";
import DateTimeActions from "./picker/ActionPicker";
import { useIsMobile } from "@/hooks/useIsMobile";
import { usePopoverPosition } from "@/hooks/usePopoverPosition";
import { FaX } from "react-icons/fa6";
import { useCloseOnScrollDistance } from "@/hooks/useCloseOnScrollDistance";

interface DateTimeFieldProps {
  id: string;
  label: string;
  value?: string | null; // bisa null atau undefined
  className?: string;
  wrapperClassName?: string;
  labelClassName?: string;
  onChange: (value: string | null) => void;
  required?: boolean;
  disabled?: boolean;
  placeholder?: string;
  error?: string;
}

export default function DateTimeField({
  id,
  label,
  value,
  className = "",
  wrapperClassName = "",
  labelClassName = "",
  onChange,
  required = false,
  disabled = false,
  placeholder = "Select date & time",
  error = "",
}: DateTimeFieldProps) {
  const baseWrapperClass = "flex flex-col";
  const baseLabelClass = "mb-1 font-medium text-gray-700";
  const baseInputClass =
    "px-3 py-2 border rounded-sm focus:outline-none focus:border-blue-500 focus:shadow-[0_0_6px_rgba(59,130,246,0.3)] disabled:bg-gray-100 disabled:text-gray-400 cursor-pointer flex items-center justify-start gap-2";

  // ✅ Pastikan hanya parse tanggal valid
  const parsedValue =
    value && !isNaN(new Date(value).getTime()) ? new Date(value) : null;

  const [tempDate, setTempDate] = useState<Date | null>(parsedValue);
  const [showPopover, setShowPopover] = useState(false);

  const { isMobile } = useIsMobile();

  const inputRef = useRef<HTMLButtonElement>(null);
  const popoverRef = useRef<HTMLDivElement>(null);
  const position = usePopoverPosition({
    inputRef,
    popoverRef,
    visible: showPopover,
    popoverHeight: isMobile ? 480 : 300,
  });

  // ✅ Tutup popover jika klik di luar
  useEffect(() => {
    const handleClick = (e: MouseEvent) => {
      const target = e.target as Node;

      if (
        inputRef.current &&
        popoverRef.current &&
        !inputRef.current.contains(target) &&
        !popoverRef.current.contains(target)
      ) {
        setShowPopover(false);
      }
    };
    document.addEventListener("mousedown", handleClick);
    return () => document.removeEventListener("mousedown", handleClick);
  }, []);

  // Tutup saat scroll
  useCloseOnScrollDistance({
    triggerRef: inputRef,
    active: showPopover,
    onClose: () => setShowPopover(false),
    distance: 20,
  });

  // ✅ Apply perubahan hanya jika valid
  const applyChange = () => {
    if (tempDate) {
      // onChange(toLocalISOString(tempDate));
      const dateStr = new Date(tempDate).toISOString();
      // console.log(dateStr)
      onChange(dateStr);
    } else onChange("");
    setShowPopover(false);
  };

  const cancelChange = () => {
    setTempDate(parsedValue);
    setShowPopover(false);
  };

  // ✅ Tampilkan value hanya kalau valid
  const displayValue =
    value && !isNaN(new Date(value).getTime()) ? formatDateTime(value) : "";

  return (
    <div className={`${baseWrapperClass} ${wrapperClassName}`}>
      <label htmlFor={id} className={`${baseLabelClass} ${labelClassName}`}>
        {label} {required && <span className="text-red-500">*</span>}
      </label>

      <button
        id={id}
        ref={inputRef}
        type="button"
        value={displayValue}
        onClick={() => !disabled && setShowPopover(true)}
        onFocus={() => !disabled && setShowPopover(true)}
        onBlur={(e) => {
          const nextFocused = e.relatedTarget as Node | null;
          if (
            popoverRef.current &&
            nextFocused &&
            popoverRef.current.contains(nextFocused)
          ) {
            return;
          }
          setShowPopover(false);
        }}
        disabled={disabled}
        className={`${baseInputClass} ${className} cursor-pointer`}
      >
        {displayValue && (
          <div
            onClick={(e) => {
              e.stopPropagation();
              onChange(null);
            }}
            className="text-red-500 hover:bg-red-100 hover:text-red-600 rounded-xs p-0.5"
          >
            <FaX size={10} />
          </div>
        )}
        {displayValue ? (
          displayValue
        ) : (
          <div className="text-gray-400">{placeholder}</div>
        )}
      </button>

      {showPopover &&
        createPortal(
          <div
            ref={popoverRef}
            tabIndex={-1}
            className={`absolute z-50 bg-white border rounded-sm shadow-lg p-3 ${
              isMobile ? "w-72" : "w-[29rem]"
            }`}
            style={{
              position: "absolute",
              top: position.top,
              left: position.left,
            }}
          >
            <div
              className={`flex ${isMobile ? "flex-col" : "flex-row"} gap-x-4`}
            >
              <DatePicker
                tempDate={tempDate ?? new Date()}
                onChange={(d) => setTempDate(d)}
              />
              <div className="flex flex-col justify-between gap-y-2">
                <div className="mx-auto">
                  <TimePicker
                    tempDate={tempDate ?? new Date()}
                    onChange={(d) => setTempDate(d)}
                  />
                </div>
                <DateTimeActions
                  tempDate={tempDate}
                  onApply={applyChange}
                  onCancel={cancelChange}
                />
              </div>
            </div>
          </div>,
          document.body,
        )}

      {error && <p className="text-red-500 text-sm mt-1">{error}</p>}
    </div>
  );
}
