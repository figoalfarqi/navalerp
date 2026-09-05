/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useCallback, useRef, useState } from "react";
import { useIndexedDB } from "@/hooks/useIndexedDB";
import { StoreName, STORES } from "@/types/tableIDB";

type FormKey = string;
type FormDataValue = any;

export function useFormObjectStore() {
  const { dbReady, getIDB, putIDB, deleteIDB } = useIndexedDB();

  const storeName: StoreName = STORES.FORM_OBJECT;

  // 🔹 semua form disimpan di sini
  const [formStoreData, setFormStoreData] = useState<Record<FormKey, FormDataValue>>({});

  // 🔸 debounce per form key
  const debounceTimers = useRef<Record<string, NodeJS.Timeout>>({});

  function waitFor(condition: () => boolean) {
    return new Promise<void>((resolve) => {
      if (condition()) return resolve();

      const interval = setInterval(() => {
        if (condition()) {
          clearInterval(interval);
          resolve();
        }
      }, 50);
    });
  }


  // =====================================================
  // 🔹 AUTO SAVE (debounced)
  // =====================================================
  const autoSave = useCallback(
    (key: FormKey, data: FormDataValue) => {
      if (!dbReady) return;

      if (debounceTimers.current[key]) {
        clearTimeout(debounceTimers.current[key]);
      }

      debounceTimers.current[key] = setTimeout(() => {
        const wrapped = { form_objects_id: key, data: data };
        putIDB(storeName, wrapped).catch((err) =>
          console.warn(`❌ Failed to autosave form "${key}":`, err)
        );
      }, 300);
    },
    [dbReady, putIDB, storeName]
  );


  // =====================================================
  // 🔹 SET FORM (local + IDB)
  // =====================================================
  const setFormStore = useCallback(
    (key: FormKey, data: FormDataValue) => {
      setFormStoreData((prev) => ({
        ...prev,
        [key]: data,
      }));

      autoSave(key, data);
    },
    [autoSave]
  );


  // =====================================================
  // 🔹 GET FORM (lazy load from IDB)
  // =====================================================
  const getFormStore = useCallback(
    async (key: FormKey) => {
      // 1️⃣ dari memory
      if (formStoreData[key]) return formStoreData[key];

      // 2️⃣ DB belum siap
      if (!dbReady) return null;

      try {
        const result = await getIDB(storeName, key);
        if (result?.data !== undefined) {
          setFormStoreData((prev) => ({
            ...prev,
            [key]: result.data,
          }));
          return result.data;
        }

        return null;
      } catch (err) {
        console.error(`❌ Failed to get form "${key}":`, err);
        return null;
      }
    },
    [dbReady, formStoreData, getIDB, storeName]
  );


  // =====================================================
  // 🔹 CLEAR FORM
  // =====================================================
  const clearFormStore = useCallback(
    async (key: FormKey) => {
      setFormStoreData((prev) => {
        const next = { ...prev };
        delete next[key];
        return next;
      });

      if (!dbReady) return;

      try {
        await deleteIDB(storeName, key);
      } catch (err) {
        console.warn(`⚠️ Failed to delete form "${key}" from IDB`, err);
      }
    },
    [dbReady, deleteIDB]
  );


  // =====================================================
  // RETURN
  // =====================================================
  return {
    formStoreData,     // optional (debug / inspect)
    setFormStore,
    getFormStore,
    clearFormStore,
  };
}




