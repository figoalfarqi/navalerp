"use client";

import { useEffect, useMemo, useRef, useState } from "react";
import { createPortal } from "react-dom";
import {
  FaCalendarDays,
  FaChevronDown,
  FaChevronLeft,
  FaChevronRight,
  FaXmark,
} from "@/components/icons";
import { useCloseOnScrollDistance } from "@/hooks/useCloseOnScrollDistance";
import { usePopoverPosition } from "@/hooks/usePopoverPosition";

const monthNames = [
  "Januari",
  "Februari",
  "Maret",
  "April",
  "Mei",
  "Juni",
  "Juli",
  "Agustus",
  "September",
  "Oktober",
  "November",
  "Desember",
];

interface ParsedMonth {
  year: number;
  month: number;
}

export interface MonthFieldProps {
  id: string;
  label: string;
  value?: string | null;
  onChange: (value: string | null) => void;
  required?: boolean;
  disabled?: boolean;
  placeholder?: string;
  error?: string;
  className?: string;
  wrapperClassName?: string;
  labelClassName?: string;
  minMonth?: string;
  maxMonth?: string;
}

function parseMonth(value?: string | null): ParsedMonth | null {
  const match = value?.match(/^(\d{4})-(\d{2})$/);
  if (!match) return null;

  const year = Number(match[1]);
  const month = Number(match[2]) - 1;
  if (!Number.isInteger(year) || month < 0 || month > 11) return null;

  return { year, month };
}

function formatMonthValue(year: number, month: number) {
  return `${year}-${String(month + 1).padStart(2, "0")}`;
}

function monthOrder({ year, month }: ParsedMonth) {
  return year * 12 + month;
}

export default function MonthField({
  id,
  label,
  value,
  onChange,
  required = false,
  disabled = false,
  placeholder = "Pilih bulan",
  error = "",
  className = "",
  wrapperClassName = "",
  labelClassName = "",
  minMonth,
  maxMonth,
}: MonthFieldProps) {
  const now = useMemo(() => new Date(), []);
  const selectedMonth = parseMonth(value);
  const parsedMinMonth = parseMonth(minMonth);
  const parsedMaxMonth = parseMonth(maxMonth);
  const [showPopover, setShowPopover] = useState(false);
  const [viewYear, setViewYear] = useState(
    selectedMonth?.year ?? now.getFullYear(),
  );
  const inputRef = useRef<HTMLButtonElement>(null);
  const popoverRef = useRef<HTMLDivElement>(null);
  const position = usePopoverPosition({
    inputRef,
    popoverRef,
    visible: showPopover,
    popoverHeight: 330,
  });

  useEffect(() => {
    if (!showPopover) return;

    const handlePointerDown = (event: MouseEvent) => {
      const target = event.target as Node;
      if (
        !inputRef.current?.contains(target) &&
        !popoverRef.current?.contains(target)
      ) {
        setShowPopover(false);
      }
    };
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === "Escape") {
        setShowPopover(false);
        inputRef.current?.focus();
      }
    };

    document.addEventListener("mousedown", handlePointerDown);
    window.addEventListener("keydown", handleKeyDown);
    return () => {
      document.removeEventListener("mousedown", handlePointerDown);
      window.removeEventListener("keydown", handleKeyDown);
    };
  }, [showPopover]);

  useCloseOnScrollDistance({
    triggerRef: inputRef,
    active: showPopover,
    onClose: () => setShowPopover(false),
    distance: 20,
  });

  const isMonthDisabled = (year: number, month: number) => {
    const candidate = monthOrder({ year, month });
    if (parsedMinMonth && candidate < monthOrder(parsedMinMonth)) return true;
    if (parsedMaxMonth && candidate > monthOrder(parsedMaxMonth)) return true;
    return false;
  };

  const yearHasAvailableMonth = (year: number) =>
    monthNames.some((_, month) => !isMonthDisabled(year, month));

  const selectMonth = (month: number) => {
    if (isMonthDisabled(viewYear, month)) return;
    onChange(formatMonthValue(viewYear, month));
    setShowPopover(false);
    requestAnimationFrame(() => inputRef.current?.focus());
  };

  const selectCurrentMonth = () => {
    const currentYear = now.getFullYear();
    const currentMonth = now.getMonth();
    if (isMonthDisabled(currentYear, currentMonth)) return;
    onChange(formatMonthValue(currentYear, currentMonth));
    setViewYear(currentYear);
    setShowPopover(false);
    requestAnimationFrame(() => inputRef.current?.focus());
  };

  const displayValue = selectedMonth
    ? `${monthNames[selectedMonth.month]} ${selectedMonth.year}`
    : "";

  return (
    <div className={`flex flex-col ${wrapperClassName}`}>
      <label
        htmlFor={id}
        className={`mb-1 font-medium text-gray-700 ${labelClassName}`}
      >
        {label} {required && <span className="text-red-500">*</span>}
      </label>

      <div className="relative">
        <button
          id={id}
          ref={inputRef}
          type="button"
          aria-expanded={showPopover}
          aria-haspopup="dialog"
          disabled={disabled}
          onClick={() => {
            if (showPopover) {
              setShowPopover(false);
              return;
            }
            setViewYear(selectedMonth?.year ?? now.getFullYear());
            setShowPopover(true);
          }}
          className={`group flex h-11 w-full items-center justify-between gap-3 rounded-xl border border-slate-200 bg-white px-3 text-left text-sm text-slate-700 shadow-sm outline-none transition hover:border-blue-300 focus:border-blue-500 focus:ring-4 focus:ring-blue-100 disabled:cursor-not-allowed disabled:bg-slate-100 disabled:text-slate-400 ${className}`}
        >
          <span className="flex min-w-0 items-center gap-2.5">
            <FaCalendarDays className="shrink-0 text-blue-500" />
            <span
              className={
                displayValue ? "truncate font-medium" : "truncate text-slate-400"
              }
            >
              {displayValue || placeholder}
            </span>
          </span>
          <FaChevronDown
            className={`shrink-0 text-xs text-slate-400 transition-transform ${
              showPopover ? "rotate-180" : ""
            } ${displayValue && !required ? "mr-7" : ""}`}
          />
        </button>

        {displayValue && !required && !disabled && (
          <button
            type="button"
            aria-label={`Kosongkan ${label}`}
            onClick={() => {
              onChange(null);
              setShowPopover(false);
            }}
            className="absolute right-8 top-1/2 -translate-y-1/2 rounded-md p-1 text-slate-400 transition hover:bg-red-50 hover:text-red-500 focus:outline-none focus:ring-2 focus:ring-red-100 cursor-pointer"
          >
            <FaXmark />
          </button>
        )}
      </div>

      {showPopover &&
        createPortal(
          <div
            ref={popoverRef}
            role="dialog"
            aria-label={`Pilih ${label}`}
            className="absolute z-50 w-80 rounded-2xl border border-slate-200 bg-white p-3 shadow-2xl shadow-slate-900/15"
            style={{
              top: position.top,
              left: position.left,
            }}
          >
            <div className="mb-3 flex items-center justify-between border-b border-slate-100 pb-2">
              <button
                type="button"
                aria-label="Tahun sebelumnya"
                disabled={!yearHasAvailableMonth(viewYear - 1)}
                onClick={() => setViewYear((year) => year - 1)}
                className="inline-flex h-9 w-9 items-center justify-center rounded-xl text-slate-500 transition hover:bg-blue-50 hover:text-blue-600 disabled:cursor-not-allowed disabled:opacity-30 cursor-pointer"
              >
                <FaChevronLeft />
              </button>
              <div className="text-center">
                <p className="text-xs font-medium uppercase tracking-wider text-slate-400">
                  Pilih bulan
                </p>
                <p className="font-semibold text-slate-800">{viewYear}</p>
              </div>
              <button
                type="button"
                aria-label="Tahun berikutnya"
                disabled={!yearHasAvailableMonth(viewYear + 1)}
                onClick={() => setViewYear((year) => year + 1)}
                className="inline-flex h-9 w-9 items-center justify-center rounded-xl text-slate-500 transition hover:bg-blue-50 hover:text-blue-600 disabled:cursor-not-allowed disabled:opacity-30 cursor-pointer"
              >
                <FaChevronRight />
              </button>
            </div>

            <div className="grid grid-cols-3 gap-2">
              {monthNames.map((monthName, month) => {
                const isSelected =
                  selectedMonth?.year === viewYear &&
                  selectedMonth.month === month;
                const isCurrent =
                  now.getFullYear() === viewYear && now.getMonth() === month;
                const isDisabled = isMonthDisabled(viewYear, month);

                return (
                  <button
                    key={monthName}
                    type="button"
                    disabled={isDisabled}
                    onClick={() => selectMonth(month)}
                    className={`rounded-xl px-2 py-2.5 text-sm font-medium transition cursor-pointer ${
                      isSelected
                        ? "bg-blue-600 text-white shadow-md shadow-blue-200"
                        : isCurrent
                          ? "bg-cyan-50 text-cyan-700 ring-1 ring-cyan-200"
                          : "text-slate-600 hover:bg-blue-50 hover:text-blue-700"
                    } disabled:cursor-not-allowed disabled:bg-slate-50 disabled:text-slate-300 disabled:shadow-none disabled:ring-0`}
                  >
                    {monthName.slice(0, 3)}
                  </button>
                );
              })}
            </div>

            <div className="mt-3 border-t border-slate-100 pt-3">
              <button
                type="button"
                disabled={isMonthDisabled(now.getFullYear(), now.getMonth())}
                onClick={selectCurrentMonth}
                className="w-full rounded-xl bg-slate-50 px-3 py-2 text-sm font-semibold text-blue-600 transition hover:bg-blue-50 disabled:cursor-not-allowed disabled:text-slate-300 cursor-pointer"
              >
                Bulan ini
              </button>
            </div>
          </div>,
          document.body,
        )}

      {error && <p className="mt-1 text-sm text-red-500">{error}</p>}
    </div>
  );
}
