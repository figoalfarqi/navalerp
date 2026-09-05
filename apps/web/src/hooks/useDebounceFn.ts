/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useRef, useCallback } from "react";

export function useDebounceFn<T extends (...args: any[]) => void>(
  fn: T,
  delay: number = 500
) {
  const timer = useRef<NodeJS.Timeout | null>(null);

  const debounced = useCallback(
    (...args: Parameters<T>) => {
      // Hapus timer sebelumnya
      if (timer.current) {
        clearTimeout(timer.current);
      }

      // Set timer baru
      timer.current = setTimeout(() => {
        fn(...args);
      }, delay);
    },
    [fn, delay]
  );

  return debounced;
}
