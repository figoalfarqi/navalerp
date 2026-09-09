"use client";

import { useState, useRef, useEffect } from "react";
import { createPortal } from "react-dom";
import DatePicker from "./picker/DatePicker";
import DateTimeActions from "./picker/ActionPicker";
import {
  durationToMinutes,
  formatDate,
  formatDuration,
  minutesToDuration,
  toLocalISOString,
} from "@/utils/dateTime";
import { usePopoverPosition } from "@/hooks/usePopoverPosition";
import { FaX } from "@/components/icons";
import DurationPicker from "./picker/DurationPicker";
import { useCloseOnScrollDistance } from "@/hooks/useCloseOnScrollDistance";

interface DurationFieldProps {
  id: string;
  label: string;
  value: number;
  onChange: (value: number | null) => void;
  required?: boolean;
  disabled?: boolean;
  placeholder?: string;
  error?: string;
  className?: string;
  wrapperClassName?: string;
  labelClassName?: string;
}

export default function DurationField({
  id,
  label,
  value,
  onChange,
  required = false,
  disabled = false,
  placeholder = "pilih durasi",
  error = "",
  className = "",
  wrapperClassName = "",
  labelClassName = "",
}: DurationFieldProps) {
  const baseWrapperClass = "flex flex-col";
  const baseLabelClass = "mb-1 font-medium text-gray-700";
  const baseInputClass =
    "px-3 py-2 border rounded-sm focus:outline-none focus:border-blue-500 focus:shadow-[0_0_6px_rgba(59,130,246,0.3)] disabled:bg-gray-100 disabled:text-gray-400 cursor-pointer flex items-center justify-start gap-2";

  const initial = minutesToDuration(value);

  const [tempDays, setTempDays] = useState(initial.days);
  const [tempHours, setTempHours] = useState(initial.hours);
  const [tempMinutes, setTempMinutes] = useState(initial.minutes);

  const [valDays, setValDays] = useState(initial.days);
  const [valHours, setValHours] = useState(initial.hours);
  const [valMinutes, setValMinutes] = useState(initial.minutes);
  const [showPopover, setShowPopover] = useState(false);

  const inputRef = useRef<HTMLButtonElement>(null);
  const popoverRef = useRef<HTMLDivElement>(null);
  const position = usePopoverPosition({
    inputRef,
    popoverRef,
    visible: showPopover,
    popoverHeight: 230,
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

  const applyChange = () => {
    onChange(durationToMinutes(tempDays, tempHours, tempMinutes));
    setValDays(tempDays);
    setValHours(tempHours);
    setValMinutes(tempMinutes);
    setShowPopover(false);
  };

  const cancelChange = () => {
    const reset = minutesToDuration(value);
    setTempDays(reset.days);
    setTempHours(reset.hours);
    setTempMinutes(reset.minutes);
    setShowPopover(false);
    setShowPopover(false);
  };

  return (
    <div className={`${baseWrapperClass} ${wrapperClassName}`}>
      <label htmlFor={id} className={`${baseLabelClass} ${labelClassName}`}>
        {label} {required && <span className="text-red-500">*</span>}
      </label>
      <button
        type="button"
        id={id}
        ref={inputRef}
        // value={displayValue}
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
        className={`${baseInputClass} ${className}`}
      >
        {formatDuration(valDays, valHours, valMinutes)}
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
            <DurationPicker
              days={tempDays}
              hours={tempHours}
              minutes={tempMinutes}
              onChange={({ days, hours, minutes }) => {
                setTempDays(days);
                setTempHours(hours);
                setTempMinutes(minutes);
              }}
            />
            <DateTimeActions
              tempDate={tempDays}
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
