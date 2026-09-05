/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useState, useCallback } from "react";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { useIndexedDB } from "@/hooks/useIndexedDB";
import { STORES } from "@/types/tableIDB";
import { APP_OBJECT_STORE, AppObjectStoreEndpointMap } from "@/consta/AppObjectsStore";
import { useOnlineStatus } from "./useOnlineStatus";

export function useAppObjectStore() {
  const { getAPI } = useFetchAPI();
  const { isOnline } = useOnlineStatus();
  const { dbReady, getIDB, putIDB, deleteIDB } = useIndexedDB();

  const storeName = STORES.APP_OBJECTS;

  // 🔹 state dinamis per key
  const [appObjectStoreData, setAppObjectStoreData] = useState<Partial<Record<APP_OBJECT_STORE, any>>>({});

  // =====================================================
  // 🔹 GENERIC FETCH + CACHE
  // =====================================================
  const fetchAppObjectStore = useCallback(
    async (key: APP_OBJECT_STORE) => {
      if (!dbReady) { return };
      const endpoint = AppObjectStoreEndpointMap[key];
      const url = `${process.env.NEXT_PUBLIC_API_BASE_URL}${endpoint}`;

      try {
        if(!isOnline){throw new Error("anda offline")}
        const res = await getAPI<any>(url, { authToken: "driver" });

        if (res.code === 200 && res.data) {
          const wrapped = { [`${storeName}_id`]: key, data: res.data };
          await putIDB(storeName, wrapped);

          setAppObjectStoreData((prev) => ({ ...prev, [key]: res.data }));
          return res.data;
        }

        await deleteIDB(storeName, key);
        setAppObjectStoreData((prev) => ({ ...prev, [key]: null }));
        return null;
      } catch (apiErr) {
        console.warn(`⚠️ API failed for ${key}, fallback to IDB`);

        try {
          const local = await getIDB(storeName, key);
          if (local?.data?.data) {
            setAppObjectStoreData((prev) => ({ ...prev, [key]: local.data.data }));
            return local.data.data;
          }
        } catch (idbErr) {
          console.error(`❌ IDB failed for ${key}`, idbErr);
        }

        return null;
      }
    },
    [dbReady, isOnline],
  );

  // =====================================================
  // 🔹 PATCH (merge data)
  // =====================================================
  const patchAppObjectStore = useCallback(
    async <T extends Record<string, any>>(
      key: APP_OBJECT_STORE,
      payload: Partial<T>,
    ) => {
      const currentData = appObjectStoreData[key] ?? {};

      // optimistic update
      const merged = { ...currentData, ...payload };
      setAppObjectStoreData((p) => ({ ...p, [key]: merged }));
      await putIDB(storeName, { [`${storeName}_id`]: key, data: merged });
    },
    [appObjectStoreData],
  );

  // =====================================================
  // 🔹 UPDATE (replace data / PUT)
  // =====================================================
  const updateAppObjectStore = useCallback(
    async <T,>(key: APP_OBJECT_STORE, payload: T) => {
      setAppObjectStoreData((p) => ({ ...p, [key]: payload }));
      await putIDB(storeName, { [`${storeName}_id`]: key, data: payload });
    },
    [],
  );

  // =====================================================
  // 🔹 CLEAR PER KEY
  // =====================================================
  const clearAppObjectStore = useCallback(
    async (key: APP_OBJECT_STORE) => {
      await deleteIDB(storeName, key);
      setAppObjectStoreData((prev) => ({ ...prev, [key]: null }));
    },
    [dbReady],
  );

  // =====================================================
  // 🔹 CLEAR ALL
  // =====================================================
  const clearAllAppObjectStore = useCallback(async () => {
    await Promise.all(
      Object.values(APP_OBJECT_STORE).map((key) => clearAppObjectStore(key)),
    );
  }, [clearAppObjectStore]);

  // =====================================================
  // RETURN
  // =====================================================
  return {
    appObjectStoreData,
    fetchAppObjectStore,
    patchAppObjectStore,
    updateAppObjectStore,
    clearAppObjectStore,
    clearAllAppObjectStore,
  };


}
