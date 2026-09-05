"use client";

import { ScrollableTimeColumn } from "./ScrollableTimeColumn";

interface TimePickerProps {
  tempDate: Date | null;
  onChange: (date: Date | null) => void;
}

export default function TimePicker({ tempDate, onChange }: TimePickerProps) {
  const safeDate = tempDate ?? new Date();

  const handleTimeChange = (hour: number, minute: number) => {
    const newDate = new Date(
      safeDate.getFullYear(),
      safeDate.getMonth(),
      safeDate.getDate(),
      hour,
      minute
    );
    onChange(newDate);
  };

  return (
    <div className="w-41">
      <div className="text-center mb-2 font-semibold">
        {tempDate
          ? `${safeDate.getHours().toString().padStart(2, "0")} : ${safeDate
              .getMinutes()
              .toString()
              .padStart(2, "0")}`
          : "-- : --"}
      </div>

      <div className="flex space-x-4 mb-3 justify-center items-center">
        {/* Hours */}
        <ScrollableTimeColumn
          values={Array.from({ length: 24 }, (_, i) => i)}
          selected={tempDate ? safeDate.getHours() : null}
          onSelect={(hour) => handleTimeChange(hour, safeDate.getMinutes())}
        />

        <span className="flex items-center font-semibold">:</span>

        {/* Minutes */}
        <ScrollableTimeColumn
          values={Array.from({ length: 60 }, (_, i) => i)}
          selected={tempDate ? safeDate.getMinutes() : null}
          onSelect={(minute) => handleTimeChange(safeDate.getHours(), minute)}
        />
      </div>
    </div>
  );
}
