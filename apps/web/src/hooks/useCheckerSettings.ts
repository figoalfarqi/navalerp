"use client";

import { useFetchAPI } from "@/hooks/useFetchAPI";
import { useEffect, useState } from "react";

const DEFAULT_LOCATION_COOLDOWN_MINUTES = 5;
const DEFAULT_TRUCK_COOLDOWN_MINUTES = 10;

type SettingRow = {
  app_setting_key?: string;
  app_setting_value?: unknown;
};

const readItems = (value: unknown): SettingRow[] => {
  if (Array.isArray(value)) return value as SettingRow[];
  if (!value || typeof value !== "object") return [];

  const data = value as Record<string, unknown>;
  if ("app_setting_key" in data) return [data as SettingRow];
  if (Array.isArray(data.items)) return data.items as SettingRow[];
  if (data.data && typeof data.data === "object") {
    const nested = data.data as Record<string, unknown>;
    if ("app_setting_key" in nested) return [nested as SettingRow];
    if (Array.isArray(nested.items)) return nested.items as SettingRow[];
  }
  return [];
};

const toPositiveMinutes = (value: unknown, fallback: number) => {
  let candidate = value;

  if (typeof candidate === "string") {
    try {
      candidate = JSON.parse(candidate);
    } catch {
      candidate = Number(candidate);
    }
  }

  if (candidate && typeof candidate === "object") {
    const record = candidate as Record<string, unknown>;
    candidate =
      record.minutes ??
      record.value ??
      record.cooldown_minutes ??
      record.duration_minutes;
  }

  const parsed = Number(candidate);
  return Number.isFinite(parsed) && parsed > 0 ? parsed : fallback;
};

export function useCheckerSettings() {
  const { getAPI } = useFetchAPI();
  const [locationCooldownMinutes, setLocationCooldownMinutes] = useState(
    DEFAULT_LOCATION_COOLDOWN_MINUTES,
  );
  const [truckCooldownMinutes, setTruckCooldownMinutes] = useState(
    DEFAULT_TRUCK_COOLDOWN_MINUTES,
  );

  useEffect(() => {
    let cancelled = false;

    const loadSettings = async () => {
      const response = await getAPI<unknown>(
        `${process.env.NEXT_PUBLIC_API_BASE_URL}/checker/app_setting?limit=50`,
        { authToken: "checker" },
      );

      if (cancelled || response.code !== 200 || !response.data) return;

      const rows = readItems(response.data);
      const locationSetting = rows.find(
        (row) =>
          row.app_setting_key === "checker.location_cooldown_minutes",
      );
      const truckSetting = rows.find(
        (row) => row.app_setting_key === "checker.truck_cooldown_minutes",
      );

      setLocationCooldownMinutes(
        toPositiveMinutes(
          locationSetting?.app_setting_value,
          DEFAULT_LOCATION_COOLDOWN_MINUTES,
        ),
      );
      setTruckCooldownMinutes(
        toPositiveMinutes(
          truckSetting?.app_setting_value,
          DEFAULT_TRUCK_COOLDOWN_MINUTES,
        ),
      );
    };

    void loadSettings();
    return () => {
      cancelled = true;
    };
  }, [getAPI]);

  return {
    locationCooldownMinutes,
    truckCooldownMinutes,
  };
}
