"use client";

import { useEffect } from "react";
import useOnlineSync from "@/hooks/useOnlineSync";

export default function ClientSync() {
  useOnlineSync();

  /* ===============================
     REGISTER + UPDATE SERVICE WORKER
     =============================== */
  useEffect(() => {
    if (!("serviceWorker" in navigator)) return;

    (async () => {
      try {
        const reg = await navigator.serviceWorker.register("/service-worker.js");

        await reg.update();
      } catch (err) {
        console.error("❌ SW registration failed:", err);
      }
    })();
  }, []);

  /* ===============================
     STORAGE PERSISTENCE (OPTIONAL)
     =============================== */
  useEffect(() => {
    if ("storage" in navigator && "persist" in navigator.storage) {
      void navigator.storage.persist();
    }
  }, []);

  return null;
}
