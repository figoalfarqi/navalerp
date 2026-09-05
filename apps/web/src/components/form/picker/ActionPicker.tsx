"use client";

import Button from "../Button";

interface DateTimeActionsProps {
  tempDate: Date | number | null;
  onApply: () => void;
  onCancel: () => void;
}

export default function DateTimeActions({
  tempDate,
  onApply,
  onCancel,
}: DateTimeActionsProps) {
  return (
    <div className="flex justify-end space-x-2">
      <Button
        variant="gray-solid"
        onClick={onCancel}
        id="cancel-button"
        size="sm"
        disabled={tempDate == null}
      >
        Batal
      </Button>
      <Button
        variant="blue-solid"
        onClick={onApply}
        id="apply-button"
        size="sm"
        disabled={tempDate == null}
      >
        Apply
      </Button>
    </div>
  );
}
