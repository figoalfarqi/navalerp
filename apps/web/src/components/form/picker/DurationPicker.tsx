"use client";
import { formatDuration } from "@/utils/dateTime";
import { ScrollableTimeColumn } from "./ScrollableTimeColumn";

interface DurationPickerProps {
  days: number;
  hours: number;
  minutes: number;
  onChange: (d: { days: number; hours: number; minutes: number }) => void;
}

export default function DurationPicker({
  days,
  hours,
  minutes,
  onChange,
}: DurationPickerProps) {
  return (
    <div className="w-65">
      <div className="text-center mb-2 font-semibold">
        {formatDuration(days, hours, minutes)}
      </div>
      <div className="flex space-x-4 mb-3 justify-center items-center">
        <ScrollableTimeColumn
          values={Array.from({ length: 31 }, (_, i) => i)}
          selected={days}
          onSelect={(v) => onChange({ days: v, hours, minutes })}
        />
        <span className="flex items-center font-semibold">:</span>
        <ScrollableTimeColumn
          values={Array.from({ length: 24 }, (_, i) => i)}
          selected={hours}
          onSelect={(v) => onChange({ days, hours: v, minutes })}
        />
        <span className="flex items-center font-semibold">:</span>
        <ScrollableTimeColumn
          values={Array.from({ length: 60 }, (_, i) => i)}
          selected={minutes}
          onSelect={(v) => onChange({ days, hours, minutes: v })}
        />
      </div>
    </div>
  );
}
