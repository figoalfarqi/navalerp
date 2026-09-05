"use client";

import { useEffect } from "react";
import DateField from "./DateField";

interface dateAfterBeforeProps {
  afterName: string;
  beforeName: string;
  afterLabel: string;
  beforeLabel: string;
  afterValue: string | null;
  beforeValue: string | null;
  onChange: (name: string, value: string | null) => void;
  className?: string;
  required?: boolean;
}

export default function DateAfterBeforeField({
  afterName,
  beforeName,
  afterLabel,
  beforeLabel,
  afterValue,
  beforeValue,
  onChange,
  className = "",
  required = false,
}: dateAfterBeforeProps) {
  // Konversi ke Date object
  const afterDate = afterValue ? new Date(afterValue) : null;
  const beforeDate = beforeValue ? new Date(beforeValue) : null;

  // Jika sebelum < sesudah (invalid), kita perbaiki otomatis
  useEffect(() => {
    if (afterDate && beforeDate && beforeDate < afterDate) {
      // Jika before lebih kecil dari after, samakan ke after
      onChange(beforeName, afterValue);
    }
  }, [afterValue, beforeValue]);

  return (
    <>
      {/* Created After */}
      <DateField
        id={afterName}
        label={afterLabel}
        value={afterValue || ""}
        onChange={(val) => onChange(afterName, val as string | null)}
        required={required}
        maxDate={beforeValue ?? undefined}
        className={className}
      />

      {/* Created Before */}
      <DateField
        id={beforeName}
        label={beforeLabel}
        value={beforeValue || ""}
        onChange={(val) => onChange(beforeName, val as string | null)}
        required={required}
        minDate={afterValue ?? undefined}
        className={className}
      />
    </>
  );
}
