"use client";

import { useEffect, useRef } from "react";

export function useSSE(
  url: string | null,
  event: string,
  onMessage: (data: any) => void
) {
  const callbackRef = useRef(onMessage);

  // selalu update callback terbaru tanpa re-init SSE
  useEffect(() => {
    callbackRef.current = onMessage;
  }, [onMessage]);

  useEffect(() => {
    if (!url) return;

    const es = new EventSource(url);

    const handler = (e: MessageEvent) => {
      try {
        const parsed = JSON.parse(e.data);
        callbackRef.current(parsed);
      } catch (err) {
        console.error("SSE parse error:", err);
      }
    };

    es.addEventListener(event, handler);

    es.onerror = (err) => {
      console.error("SSE error:", err);
      es.close();
    };

    return () => {
      es.removeEventListener(event, handler);
      es.close();
    };
  }, [url, event]);
}
