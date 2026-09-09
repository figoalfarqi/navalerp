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

const yearsPerPage = 12;

export interface YearFieldProps {
  id: string;
  label: string;
  value?: string | number | null;
  onChange: (value: number | null) => void;
  required?: boolean;
  disabled?: boolean;
  placeholder?: string;
  error?: string;
  className?: string;
  wrapperClassName?: string;
  labelClassName?: string;
  minYear?: number;
  maxYear?: number;
}

function parseYear(value?: string | number | null) {
  if (value === null || value === undefined || value === "") return null;
  const year = Number(value);
  return Number.isInteger(year) ? year : null;
}

function pageStartFor(year: number) {
  return Math.floor(year / yearsPerPage) * yearsPerPage;
}

export default function YearField({
  id,
  label,
  value,
  onChange,
  required = false,
  disabled = false,
  placeholder = "Pilih tahun",
  error = "",
  className = "",
  wrapperClassName = "",
  labelClassName = "",
  minYear,
  maxYear,
}: YearFieldProps) {
  const currentYear = useMemo(() => new Date().getFullYear(), []);
  const selectedYear = parseYear(value);
  const [showPopover, setShowPopover] = useState(false);
  const [pageStart, setPageStart] = useState(
    pageStartFor(selectedYear ?? currentYear),
  );
  const inputRef = useRef<HTMLButtonElement>(null);
  const popoverRef = useRef<HTMLDivElement>(null);
  const position = usePopoverPosition({
    inputRef,
    popoverRef,
    visible: showPopover,
    popoverHeight: 330,
  });
  const years = useMemo(
    () =>
      Array.from({ length: yearsPerPage }, (_, index) => pageStart + index),
    [pageStart],
  );

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

  const isYearDisabled = (year: number) =>
    (minYear !== undefined && year < minYear) ||
    (maxYear !== undefined && year > maxYear);

  const pageHasAvailableYear = (start: number) =>
    Array.from({ length: yearsPerPage }, (_, index) => start + index).some(
      (year) => !isYearDisabled(year),
    );

  const selectYear = (year: number) => {
    if (isYearDisabled(year)) return;
    onChange(year);
    setShowPopover(false);
    requestAnimationFrame(() => inputRef.current?.focus());
  };

  const selectCurrentYear = () => {
    if (isYearDisabled(currentYear)) return;
    onChange(currentYear);
    setPageStart(pageStartFor(currentYear));
    setShowPopover(false);
    requestAnimationFrame(() => inputRef.current?.focus());
  };

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
            setPageStart(pageStartFor(selectedYear ?? currentYear));
            setShowPopover(true);
          }}
          className={`group flex h-11 w-full items-center justify-between gap-3 rounded-xl border border-slate-200 bg-white px-3 text-left text-sm text-slate-700 shadow-sm outline-none transition hover:border-blue-300 focus:border-blue-500 focus:ring-4 focus:ring-blue-100 disabled:cursor-not-allowed disabled:bg-slate-100 disabled:text-slate-400 ${className}`}
        >
          <span className="flex min-w-0 items-center gap-2.5">
            <FaCalendarDays className="shrink-0 text-blue-500" />
            <span
              className={
                selectedYear
                  ? "truncate font-medium"
                  : "truncate text-slate-400"
              }
            >
              {selectedYear ?? placeholder}
            </span>
          </span>
          <FaChevronDown
            className={`shrink-0 text-xs text-slate-400 transition-transform ${
              showPopover ? "rotate-180" : ""
            } ${selectedYear && !required ? "mr-7" : ""}`}
          />
        </button>

        {selectedYear && !required && !disabled && (
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
            <div className="mb-3 flex items-center justify-between">
              <button
                type="button"
                aria-label="Rentang tahun sebelumnya"
                disabled={!pageHasAvailableYear(pageStart - yearsPerPage)}
                onClick={() => setPageStart((year) => year - yearsPerPage)}
                className="inline-flex h-9 w-9 items-center justify-center rounded-xl text-slate-500 transition hover:bg-blue-50 hover:text-blue-600 disabled:cursor-not-allowed disabled:opacity-30 cursor-pointer"
              >
                <FaChevronLeft />
              </button>
              <div className="text-center">
                <p className="text-xs font-medium uppercase tracking-wider text-slate-400">
                  Pilih tahun
                </p>
                <p className="font-semibold text-slate-800">
                  {years[0]} — {years[years.length - 1]}
                </p>
              </div>
              <button
                type="button"
                aria-label="Rentang tahun berikutnya"
                disabled={!pageHasAvailableYear(pageStart + yearsPerPage)}
                onClick={() => setPageStart((year) => year + yearsPerPage)}
                className="inline-flex h-9 w-9 items-center justify-center rounded-xl text-slate-500 transition hover:bg-blue-50 hover:text-blue-600 disabled:cursor-not-allowed disabled:opacity-30 cursor-pointer"
              >
                <FaChevronRight />
              </button>
            </div>

            <div className="grid grid-cols-3 gap-2">
              {years.map((year) => {
                const isSelected = selectedYear === year;
                const isCurrent = currentYear === year;

                return (
                  <button
                    key={year}
                    type="button"
                    disabled={isYearDisabled(year)}
                    onClick={() => selectYear(year)}
                    className={`rounded-xl px-2 py-2.5 text-sm font-medium transition cursor-pointer ${
                      isSelected
                        ? "bg-blue-600 text-white shadow-md shadow-blue-200"
                        : isCurrent
                          ? "bg-cyan-50 text-cyan-700 ring-1 ring-cyan-200"
                          : "text-slate-600 hover:bg-blue-50 hover:text-blue-700"
                    } disabled:cursor-not-allowed disabled:bg-slate-50 disabled:text-slate-300 disabled:shadow-none disabled:ring-0`}
                  >
                    {year}
                  </button>
                );
              })}
            </div>

            <div className="mt-3 border-t border-slate-100 pt-3">
              <button
                type="button"
                disabled={isYearDisabled(currentYear)}
                onClick={selectCurrentYear}
                className="w-full rounded-xl bg-slate-50 px-3 py-2 text-sm font-semibold text-blue-600 transition hover:bg-blue-50 disabled:cursor-not-allowed disabled:text-slate-300 cursor-pointer"
              >
                Tahun ini
              </button>
            </div>
          </div>,
          document.body,
        )}

      {error && <p className="mt-1 text-sm text-red-500">{error}</p>}
    </div>
  );
}
