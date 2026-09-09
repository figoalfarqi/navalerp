"use client";

import { useEffect, useRef, useState } from "react";
import { CgChevronLeft, CgChevronRight } from "@/components/icons";

interface DatePickerProps {
  tempDate: Date | null;
  onChange: (date: Date | null) => void;
  minDate?: Date; // ✅ new
  maxDate?: Date; // ✅ new
  onSelectYearImmediate?: (year: number) => void; //
}

type ViewMode = "day" | "month" | "year";

function rangeArray(start: number, end: number): number[] {
  return Array.from({ length: end - start + 1 }, (_, i) => start + i);
}

export default function DatePicker({
  tempDate,
  onChange,
  minDate,
  maxDate,
  onSelectYearImmediate,
}: DatePickerProps) {
  const today = new Date();
  const safeDate = tempDate ?? today;
  const [currentMonth, setCurrentMonth] = useState(
    new Date(safeDate.getFullYear(), safeDate.getMonth(), 1),
  );
  const [viewMode, setViewMode] = useState<ViewMode>(
    onSelectYearImmediate ? "year" : "day",
  );
  const [viewYear, setViewYear] = useState(currentMonth.getFullYear());
  const [years, setYears] = useState(rangeArray(viewYear - 50, viewYear + 50));

  useEffect(() => {
    setYears(rangeArray(viewYear - 50, viewYear + 50));
  }, [viewYear]);

  const startOfMonth = new Date(
    currentMonth.getFullYear(),
    currentMonth.getMonth(),
    1,
  );
  const endOfMonth = new Date(
    currentMonth.getFullYear(),
    currentMonth.getMonth() + 1,
    0,
  );

  const yearRefs = useRef<Record<number, HTMLButtonElement | null>>({});

  useEffect(() => {
    if (viewMode !== "year") return;

    const currentYear = currentMonth.getFullYear();
    const el = yearRefs.current[currentYear];

    if (el) {
      el.scrollIntoView({
        behavior: "smooth",
        block: "center",
      });
    }
  }, [viewMode, viewYear, currentMonth]);

  const startDay = startOfMonth.getDay();
  const daysInMonth = endOfMonth.getDate();

  const days: (Date | null)[] = [];
  for (let i = 0; i < startDay; i++) days.push(null);
  for (let d = 1; d <= daysInMonth; d++) {
    days.push(new Date(currentMonth.getFullYear(), currentMonth.getMonth(), d));
  }

  const handleSelectDate = (date: Date) => {
    onChange(
      new Date(
        date.getFullYear(),
        date.getMonth(),
        date.getDate(),
        safeDate.getHours(),
        safeDate.getMinutes(),
      ),
    );
  };

  const handleSelectMonth = (month: number) => {
    setCurrentMonth(new Date(currentMonth.getFullYear(), month, 1));
    setViewMode("day");
  };

  const handleSelectYear = (year: number) => {
    setCurrentMonth(new Date(year, currentMonth.getMonth(), 1));
    setViewMode("month");
  };

  const isOutOfRange = (date: Date): boolean => {
    const normalize = (d: Date) =>
      new Date(d.getFullYear(), d.getMonth(), d.getDate());

    const normalizedDate = normalize(date);
    const normalizedMin = minDate ? normalize(minDate) : null;
    const normalizedMax = maxDate ? normalize(maxDate) : null;

    if (normalizedMin && normalizedDate < normalizedMin) return true;
    if (normalizedMax && normalizedDate > normalizedMax) return true;
    return false;
  };

  return (
    <div className="w-65 h-62">
      {/* Header */}
      <div className="flex justify-between items-center mb-2">
        <button
          type="button"
          className="px-2 py-1 hover:bg-blue-500/30 rounded cursor-pointer"
          onClick={() => {
            if (viewMode === "day")
              setCurrentMonth(
                new Date(
                  currentMonth.getFullYear(),
                  currentMonth.getMonth() - 1,
                  1,
                ),
              );
            else if (viewMode === "month")
              setCurrentMonth(
                new Date(
                  currentMonth.getFullYear() - 1,
                  currentMonth.getMonth(),
                  1,
                ),
              );
            else if (viewMode === "year") setViewYear(viewYear - 100);
          }}
        >
          <CgChevronLeft />
        </button>

        <span className="font-semibold select-none">
          {viewMode === "day" && (
            <span
              onClick={() => setViewMode("month")}
              className="cursor-pointer hover:text-blue-600 transition"
            >
              {currentMonth.toLocaleString("default", { month: "long" })}{" "}
              {currentMonth.getFullYear()}
            </span>
          )}
          {viewMode === "month" && (
            <span
              onClick={() => setViewMode("year")}
              className="cursor-pointer hover:text-blue-600 transition"
            >
              {currentMonth.getFullYear()}
            </span>
          )}
          {viewMode === "year" && `${years[0]} - ${years[years.length - 1]}`}
        </span>

        <button
          type="button"
          className="px-2 py-1 hover:bg-blue-500/30 rounded cursor-pointer"
          onClick={() => {
            if (viewMode === "day")
              setCurrentMonth(
                new Date(
                  currentMonth.getFullYear(),
                  currentMonth.getMonth() + 1,
                  1,
                ),
              );
            else if (viewMode === "month")
              setCurrentMonth(
                new Date(
                  currentMonth.getFullYear() + 1,
                  currentMonth.getMonth(),
                  1,
                ),
              );
            else if (viewMode === "year") setViewYear(viewYear + 100);
          }}
        >
          <CgChevronRight />
        </button>
      </div>

      {/* Body */}
      {viewMode === "day" && (
        <div className="grid grid-cols-7 gap-[2px] text-center text-sm font-medium mb-1">
          {["Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"].map((d) => (
            <div key={d} className="text-gray-500">
              {d}
            </div>
          ))}
          {days.map((date, i) => {
            if (!date) return <div key={i} className="w-10 h-8" />;
            const isSelected =
              tempDate && date.toDateString() === tempDate.toDateString();
            const isToday = date.toDateString() === today.toDateString();
            const outOfRange = isOutOfRange(date);

            return (
              <button
                key={i}
                type="button"
                disabled={outOfRange}
                className={`w-10 h-8 rounded-full transition cursor-pointer ${
                  outOfRange
                    ? "text-gray-300 cursor-not-allowed"
                    : isSelected
                      ? "bg-blue-500 text-white"
                      : isToday
                        ? "bg-cyan-500 text-white"
                        : "hover:bg-blue-500/30"
                }`}
                onClick={() => !outOfRange && handleSelectDate(date)}
              >
                {date.getDate()}
              </button>
            );
          })}
        </div>
      )}

      {viewMode === "month" && (
        <div className="grid grid-cols-3 gap-2 text-center">
          {Array.from({ length: 12 }, (_, m) => (
            <button
              key={m}
              className={`py-2 px-3 rounded hover:bg-blue-500/30 cursor-pointer ${
                currentMonth.getMonth() === m ? "bg-blue-500 text-white" : ""
              }`}
              onClick={() => handleSelectMonth(m)}
            >
              {new Date(0, m).toLocaleString("default", { month: "short" })}
            </button>
          ))}
        </div>
      )}

      {viewMode === "year" && (
        <div
          className={`overflow-y-auto grid grid-cols-3 gap-2 text-center ${onSelectYearImmediate ? "h-50" : "h-40"}`}
        >
          {years.map((y) => (
            <button
              key={y}
              ref={(el) => {
                yearRefs.current[y] = el;
              }}
              className={`py-2 rounded hover:bg-blue-500/30 cursor-pointer ${
                currentMonth.getFullYear() === y ? "bg-blue-500 text-white" : ""
              }`}
              onClick={() => {
                if (onSelectYearImmediate) {
                  onSelectYearImmediate(y);
                } else {
                  handleSelectYear(y);
                }
              }}
            >
              {y}
            </button>
          ))}
        </div>
      )}
    </div>
  );
}
