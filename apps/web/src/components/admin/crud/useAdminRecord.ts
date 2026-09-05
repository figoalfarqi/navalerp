"use client";

import { useEffect, useState } from "react";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { useAuth } from "@/context/AuthContext";
import type { AdminCrudMode } from "./adminCrud";

type UnknownRecord = Record<string, unknown>;

function unwrapRecord(value: unknown): UnknownRecord | null {
  if (!value || typeof value !== "object" || Array.isArray(value)) return null;
  const root = value as UnknownRecord;
  if (root.data && typeof root.data === "object" && !Array.isArray(root.data)) {
    return { ...(root.data as UnknownRecord) };
  }
  return { ...root };
}

interface UseAdminRecordOptions {
  endpoint: string;
  id: string;
  mode: Exclude<AdminCrudMode, "add">;
  primaryKey: string;
  normalize?: (record: UnknownRecord) => UnknownRecord;
}

export function useAdminRecord({
  endpoint,
  id,
  mode,
  primaryKey,
  normalize,
}: UseAdminRecordOptions) {
  const { getAPI } = useFetchAPI();
  const { adminToken, tokenLoaded } = useAuth();
  const [record, setRecord] = useState<UnknownRecord | null>(null);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!endpoint || !id) return;
    if (!tokenLoaded && adminToken === null) return;

    let active = true;
    void getAPI<unknown>(`${endpoint}/${id}`, { authToken: "admin" })
      .then((response) => {
        if (!active) return;
        if (response.code < 200 || response.code >= 300) {
          setError(response.message || "Data gagal dimuat.");
          return;
        }
        const unwrapped = unwrapRecord(response.data);
        if (!unwrapped) {
          setError("Data tidak ditemukan.");
          return;
        }
        const next = normalize ? normalize(unwrapped) : unwrapped;
        if (mode === "copy") delete next[primaryKey];
        setRecord(next);
      })
      .catch((err) => {
        if (!active) return;
        setError(err?.message || "Data gagal dimuat.");
      });

    return () => {
      active = false;
    };
  }, [endpoint, id, mode, normalize, primaryKey, getAPI, adminToken, tokenLoaded]);

  return { record, error };
}
