"use client";

import { useState, useRef, useEffect } from "react";
import { createPortal } from "react-dom";
import TimePicker from "./picker/TimePicker";
import DateTimeActions from "./picker/ActionPicker";
import { formatTime, toLocalTimeString } from "@/utils/dateTime";
import { usePopoverPosition } from "@/hooks/usePopoverPosition";
import { FaX } from "react-icons/fa6";
import { useCloseOnScrollDistance } from "@/hooks/useCloseOnScrollDistance";

interface TimeFieldProps {
  id: string;
  label: string;
  value?: string | null; // bisa null atau undefined
  onChange: (value: string | null) => void;
  required?: boolean;
  disabled?: boolean;
  placeholder?: string;
  error?: string;
  className?: string;
  wrapperClassName?: string;
  labelClassName?: string;
}

export default function TimeField({
  id,
  label,
  value,
  onChange,
  required = false,
  disabled = false,
  placeholder = "Select time",
  error = "",
  className = "",
  wrapperClassName = "",
  labelClassName = "",
}: TimeFieldProps) {
  const baseWrapperClass = "flex flex-col";
  const baseLabelClass = "mb-1 font-medium text-gray-700";
  const baseInputClass =
    "px-3 py-2 border rounded-sm focus:outline-none focus:border-blue-500 focus:shadow-[0_0_6px_rgba(59,130,246,0.3)] disabled:bg-gray-100 disabled:text-gray-400 cursor-pointer flex items-center justify-start gap-2";

  const parsedValue =
    value && !isNaN(new Date(value).getTime()) ? new Date(value) : null;

  const [tempDate, setTempDate] = useState<Date | null>(parsedValue);
  const [showPopover, setShowPopover] = useState(false);

  const inputRef = useRef<HTMLButtonElement>(null);
  const popoverRef = useRef<HTMLDivElement>(null);
  const position = usePopoverPosition({ inputRef, visible: showPopover });

  // Tutup popover jika klik di luar
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

  const applyChange = () => {
    if (tempDate) onChange(toLocalTimeString(tempDate));
    else onChange(""); // kosongkan kalau null
    setShowPopover(false);
  };

  const cancelChange = () => {
    setTempDate(parsedValue);
    setShowPopover(false);
  };

  const displayValue =
    value && !isNaN(new Date(value).getTime()) ? formatTime(value) : "";

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
            className="absolute z-50 bg-white border rounded-sm shadow-lg p-3 w-52"
            style={{
              position: "absolute",
              top: position.top,
              left: position.left,
            }}
          >
            <TimePicker tempDate={tempDate} onChange={setTempDate} />
            <DateTimeActions
              tempDate={tempDate}
              onApply={applyChange}
              onCancel={cancelChange}
            />
          </div>,
          document.body,
        )}

      {error && <p className="text-red-500 text-sm mt-1">{error}</p>}
    </div>
  );
}
