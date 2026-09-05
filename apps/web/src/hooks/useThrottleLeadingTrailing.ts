/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useRef, useCallback } from "react";

export function useThrottleLeadingTrailing<T extends (...args: any[]) => void>(
  func: T,
  delay: number = 500
) {
  const lastCall = useRef(0);
  const timeout = useRef<NodeJS.Timeout | null>(null);
  const lastArgs = useRef<any[] | null>(null);

  const throttledFn = useCallback(
    (...args: Parameters<T>) => {
      const now = Date.now();

      // ---- LEADING: jalankan pertama langsung ----
      const shouldRunNow = now - lastCall.current >= delay;

      if (shouldRunNow) {
        lastCall.current = now;
        func(...args);
      } else {
        // ---- TRAILING: simpan args terakhir ----
        lastArgs.current = args;

        // jika belum ada timeout, buat timeout
        if (!timeout.current) {
          const remaining = delay - (now - lastCall.current);

          timeout.current = setTimeout(() => {
            timeout.current = null;
            lastCall.current = Date.now();

            if (lastArgs.current) {
              func(...lastArgs.current);
              lastArgs.current = null;
            }
          }, remaining);
        }
      }
    },
    [func, delay]
  );

  return throttledFn;
}
