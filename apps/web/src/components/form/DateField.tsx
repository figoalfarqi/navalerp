"use client";

import { useState, useRef, useEffect } from "react";
import { createPortal } from "react-dom";
import DatePicker from "./picker/DatePicker";
import DateTimeActions from "./picker/ActionPicker";
import { formatDate } from "@/utils/dateTime";
import { usePopoverPosition } from "@/hooks/usePopoverPosition";
import { FaX, FaCalendarDays } from "@/components/icons";
import { useCloseOnScrollDistance } from "@/hooks/useCloseOnScrollDistance";

interface DateFieldProps {
  id: string;
  label: string;
  value: string | number; // ISO string
  onChange: (value: string | null | number) => void;
  required?: boolean;
  disabled?: boolean;
  uncloseable?: boolean;
  placeholder?: string;
  error?: string;
  className?: string;
  wrapperClassName?: string;
  labelClassName?: string;
  minDate?: string;
  maxDate?: string;
  isYearOnly?: boolean;
}

export default function DateField({
  id,
  label,
  value,
  onChange,
  required = false,
  disabled = false,
  uncloseable = false,
  placeholder = "Pilih Tanggal",
  error = "",
  className = "",
  wrapperClassName = "",
  labelClassName = "",
  minDate,
  maxDate,
  isYearOnly = false,
}: DateFieldProps) {
  const baseWrapperClass = "flex flex-col w-full";
  const baseLabelClass = "mb-1 font-medium text-gray-700";
  const baseInputClass =
    "w-full px-3 py-2 border rounded-sm focus:outline-none focus:border-blue-500 focus:shadow-[0_0_6px_rgba(59,130,246,0.3)] disabled:bg-gray-100 disabled:text-gray-400 cursor-pointer flex items-center justify-between gap-2";

  const parsedValue =
    value && !isNaN(new Date(value).getTime()) ? new Date(value) : null;
  const [tempDate, setTempDate] = useState<Date | null>(parsedValue);
  const [showPopover, setShowPopover] = useState(false);

  useEffect(() => {
    setTempDate(parsedValue);
  }, [value]);

  const inputRef = useRef<HTMLButtonElement>(null);
  const popoverRef = useRef<HTMLDivElement>(null);
  const position = usePopoverPosition({
    inputRef,
    popoverRef,
    visible: showPopover,
    popoverHeight: 300,
  });

  // Tutup saat klik di luar
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

  const applyChangeYear = (year: number) => {
    const date = new Date(year, 0, 1); // Jan = 0
    onChange(year);
    setShowPopover(false);
    setTempDate(date);
  };

  const applyChange = () => {
    if (tempDate) {
      // onChange(tempDate.toISOString());
      const dateStr = new Date(tempDate).toISOString();
      onChange(dateStr);
    } else {
      onChange(null);
    }
    setShowPopover(false);
  };

  const cancelChange = () => {
    setTempDate(parsedValue);
    setShowPopover(false);
  };

  const displayValue =
    value &&
    (isYearOnly
      ? value
      : !isNaN(new Date(value).getTime())
        ? formatDate(value as string)
        : "");

  return (
    <div className={`${baseWrapperClass} ${wrapperClassName}`}>
      <label htmlFor={id} className={`${baseLabelClass} ${labelClassName}`}>
        {label} {required && <span className="text-red-500">*</span>}
      </label>
      <button
        id={id}
        ref={inputRef}
        value={displayValue}
        type="button"
        onClick={() => {
          if (!disabled) {
            setTempDate(parsedValue);
            setShowPopover(true);
          }
        }}
        onFocus={() => {
          if (!disabled) {
            setTempDate(parsedValue);
            setShowPopover(true);
          }
        }}
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
        className={`${baseInputClass} ${className}`}
      >
        <div className="flex items-center gap-2 overflow-hidden">
          {displayValue && !uncloseable && (
            <div
              onClick={(e) => {
                e.stopPropagation();
                onChange(null);
              }}
              className="cursor-pointer text-red-500 hover:bg-red-100 hover:text-red-600 rounded-xs p-0.5 shrink-0"
            >
              <FaX size={10} />
            </div>
          )}
          {displayValue ? (
            <span className="truncate">{displayValue}</span>
          ) : (
            <span className="text-gray-400 truncate">{placeholder}</span>
          )}
        </div>
        <FaCalendarDays className="text-gray-400 ml-auto shrink-0" size={14} />
      </button>

      {showPopover &&
        createPortal(
          <div
            ref={popoverRef}
            tabIndex={-1}
            className="absolute z-50 bg-white border rounded-sm shadow-lg p-3 w-75"
            style={{
              position: "absolute",
              top: position.top,
              left: position.left,
            }}
          >
            <DatePicker
              tempDate={tempDate}
              onChange={setTempDate}
              minDate={minDate ? new Date(minDate) : undefined}
              maxDate={maxDate ? new Date(maxDate) : undefined}
              {...(isYearOnly
                ? { onSelectYearImmediate: applyChangeYear }
                : {})}
            />
            {!isYearOnly && (
              <DateTimeActions
                tempDate={tempDate}
                onApply={applyChange}
                onCancel={cancelChange}
              />
            )}
          </div>,
          document.body,
        )}

      {error && <p className="text-red-500 text-sm mt-1">{error}</p>}
    </div>
  );
}
