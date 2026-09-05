"use client";

import {
  CheckerLocation,
  CheckerPosition,
} from "@/types/checkerPosition.type";
import { ProjectRouteType } from "@/types/project.type";
import { useCallback, useEffect, useMemo, useState } from "react";

const POSITION_VERSION = "v1";
const TRUCK_COOLDOWN_VERSION = "v1";

type PositionDraft = {
  project_id: number;
  project_code: string;
  project_name: string;
  project_route_type: ProjectRouteType;
  location: CheckerLocation;
};

type SavePositionResult =
  | { ok: true; position: CheckerPosition }
  | { ok: false; remainingSeconds: number };

const safeParse = <T,>(value: string | null, fallback: T): T => {
  if (!value) return fallback;
  try {
    return JSON.parse(value) as T;
  } catch {
    return fallback;
  }
};

const getPositionStorageKey = (checkerId: number) =>
  `pml.checker.${checkerId}.position.${POSITION_VERSION}`;

const getTruckCooldownStorageKey = (checkerId: number) =>
  `pml.checker.${checkerId}.truck-cooldowns.${TRUCK_COOLDOWN_VERSION}`;

export function useCheckerLocalState({
  checkerId,
  locationCooldownMinutes,
  truckCooldownMinutes,
}: {
  checkerId: number;
  locationCooldownMinutes: number;
  truckCooldownMinutes: number;
}) {
  const [isLoaded, setIsLoaded] = useState(false);
  const [now, setNow] = useState(() => Date.now());
  const [position, setPosition] = useState<CheckerPosition | null>(null);
  const [truckCooldowns, setTruckCooldowns] = useState<Record<string, number>>(
    {},
  );

  const positionStorageKey = useMemo(
    () => getPositionStorageKey(checkerId),
    [checkerId],
  );
  const truckCooldownStorageKey = useMemo(
    () => getTruckCooldownStorageKey(checkerId),
    [checkerId],
  );

  useEffect(() => {
    if (!checkerId) return;

    const frame = window.requestAnimationFrame(() => {
      setPosition(
        safeParse<CheckerPosition | null>(
          localStorage.getItem(positionStorageKey),
          null,
        ),
      );
      setTruckCooldowns(
        safeParse<Record<string, number>>(
          localStorage.getItem(truckCooldownStorageKey),
          {},
        ),
      );
      setIsLoaded(true);
    });
    return () => window.cancelAnimationFrame(frame);
  }, [checkerId, positionStorageKey, truckCooldownStorageKey]);

  useEffect(() => {
    const timer = window.setInterval(() => setNow(Date.now()), 1000);
    return () => window.clearInterval(timer);
  }, []);

  useEffect(() => {
    const onStorage = (event: StorageEvent) => {
      if (event.key === positionStorageKey) {
        setPosition(safeParse<CheckerPosition | null>(event.newValue, null));
      }
      if (event.key === truckCooldownStorageKey) {
        setTruckCooldowns(
          safeParse<Record<string, number>>(event.newValue, {}),
        );
      }
    };

    window.addEventListener("storage", onStorage);
    return () => window.removeEventListener("storage", onStorage);
  }, [positionStorageKey, truckCooldownStorageKey]);

  const locationRemainingSeconds = Math.max(
    0,
    Math.ceil(((position?.location_locked_until ?? 0) - now) / 1000),
  );

  const savePosition = useCallback(
    (draft: PositionDraft): SavePositionResult => {
      const currentTime = Date.now();
      const isLocationChanged =
        position === null || position.location !== draft.location;
      const remainingSeconds = Math.max(
        0,
        Math.ceil(
          ((position?.location_locked_until ?? 0) - currentTime) / 1000,
        ),
      );

      if (position && isLocationChanged && remainingSeconds > 0) {
        return { ok: false, remainingSeconds };
      }

      const nextPosition: CheckerPosition = {
        ...draft,
        saved_at: currentTime,
        location_locked_until: isLocationChanged
          ? currentTime + locationCooldownMinutes * 60_000
          : (position?.location_locked_until ?? currentTime),
      };

      localStorage.setItem(positionStorageKey, JSON.stringify(nextPosition));
      setPosition(nextPosition);
      setNow(currentTime);
      return { ok: true, position: nextPosition };
    },
    [locationCooldownMinutes, position, positionStorageKey],
  );

  const getTruckRemainingSeconds = useCallback(
    (truckId?: number | null) => {
      if (!truckId) return 0;
      return Math.max(
        0,
        Math.ceil(((truckCooldowns[String(truckId)] ?? 0) - now) / 1000),
      );
    },
    [now, truckCooldowns],
  );

  const lockTruck = useCallback(
    (truckId: number) => {
      const deadline = Date.now() + truckCooldownMinutes * 60_000;
      setTruckCooldowns((current) => {
        const next = { ...current, [String(truckId)]: deadline };
        localStorage.setItem(truckCooldownStorageKey, JSON.stringify(next));
        return next;
      });
      setNow(Date.now());
    },
    [truckCooldownMinutes, truckCooldownStorageKey],
  );

  return {
    isLoaded,
    position,
    locationRemainingSeconds,
    savePosition,
    getTruckRemainingSeconds,
    lockTruck,
  };
}
