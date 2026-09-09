"use client";

import Button from "@/components/form/Button";
import ItemCardSkeleton from "@/components/driver/ItemCardSkeleton";
import ProjectTransportCard from "@/components/mobile/ProjectTransportCard";
import { useAuth } from "@/context/AuthContext";
import { useCheckerLocalState } from "@/hooks/useCheckerLocalState";
import { useCheckerSettings } from "@/hooks/useCheckerSettings";
import { authTokenType, useFetchAPI } from "@/hooks/useFetchAPI";
import { ProjectTransport } from "@/types/projectTransport.type";
import { ProjectTransportStatus } from "@/types/projectTransportStatus.type";
import { useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
import { FiMapPin, FiRefreshCw } from "@/components/icons";

const readItems = <T,>(value: unknown): T[] => {
  if (Array.isArray(value)) return value as T[];
  if (!value || typeof value !== "object") return [];
  const record = value as Record<string, unknown>;
  if (Array.isArray(record.items)) return record.items as T[];
  if (record.data && typeof record.data === "object") {
    const nested = record.data as Record<string, unknown>;
    if (Array.isArray(nested.items)) return nested.items as T[];
  }
  return [];
};

export default function TransportData({
  transportType,
  authToken,
}: {
  transportType: "active" | "history";
  authToken: authTokenType;
}) {
  const router = useRouter();
  const { checkerPayload } = useAuth();
  const { getAPI } = useFetchAPI();
  const { locationCooldownMinutes, truckCooldownMinutes } =
    useCheckerSettings();
  const { isLoaded: isPositionLoaded, position } = useCheckerLocalState({
    checkerId: checkerPayload.app_user_id,
    locationCooldownMinutes,
    truckCooldownMinutes,
  });

  const [transports, setTransports] = useState<ProjectTransport[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState("");

  const loadTransports = useCallback(async () => {
    if (!position) return;
    setIsLoading(true);
    setError("");

    const params = new URLSearchParams({
      project_id: String(position.project_id),
      is_completed: transportType === "active" ? "0" : "1",
      limit: "50",
      order_by: "transported_at",
      sort: "desc",
    });
    const response = await getAPI<unknown>(
      `${process.env.NEXT_PUBLIC_API_BASE_URL}/${authToken}/project_transport?${params.toString()}`,
      { authToken },
    );

    if (response.code !== 200 || !response.data) {
      setTransports([]);
      setError(response.message || "Gagal memuat transport");
      setIsLoading(false);
      return;
    }

    const expectedCompleted = transportType === "history" ? 1 : 0;
    const items = readItems<ProjectTransport>(response.data).filter(
      (transport) =>
        Number(transport.is_completed ?? 0) === expectedCompleted,
    );
    const enriched = await Promise.all(
      items.map(async (transport) => {
        if (
          transport.project_transport_statuses?.length ||
          transport.statuses?.length
        ) {
          return transport;
        }

        const statusResponse = await getAPI<unknown>(
          `${process.env.NEXT_PUBLIC_API_BASE_URL}/${authToken}/project_transport_status?project_transport_id=${transport.project_transport_id}&limit=20&order_by=status_time&sort=desc`,
          { authToken },
        );
        const statuses =
          statusResponse.code === 200 && statusResponse.data
            ? readItems<ProjectTransportStatus>(statusResponse.data)
            : [];
        return { ...transport, project_transport_statuses: statuses };
      }),
    );

    setTransports(enriched);
    setIsLoading(false);
  }, [authToken, getAPI, position, transportType]);

  useEffect(() => {
    if (!isPositionLoaded) return;
    if (!position) {
      router.replace("/checker/position");
      return;
    }
    const frame = window.requestAnimationFrame(() => void loadTransports());
    return () => window.cancelAnimationFrame(frame);
  }, [isPositionLoaded, loadTransports, position, router]);

  if (!isPositionLoaded) return <ItemCardSkeleton />;

  return (
    <main className="mx-auto min-h-dvh max-w-md bg-slate-50 pb-24">
      <header className="border-b border-slate-200 bg-white px-5 pb-5 pt-7">
        <div className="flex items-start justify-between gap-3">
          <div>
            <p className="text-xs font-semibold uppercase tracking-[0.16em] text-blue-600">
              {transportType === "active" ? "Operasional" : "Arsip"}
            </p>
            <h1 className="mt-1 text-2xl font-bold text-slate-900">
              {transportType === "active"
                ? "Transport aktif"
                : "Riwayat transport"}
            </h1>
          </div>
          <Button
            id={`reload_${transportType}_transport`}
            variant="gray-outline"
            onClick={() => void loadTransports()}
            disabled={isLoading}
            className="h-10 w-10 rounded-full p-0"
          >
            <FiRefreshCw className={isLoading ? "animate-spin" : ""} />
          </Button>
        </div>

        {position && (
          <button
            type="button"
            onClick={() => router.push("/checker/position")}
            className="mt-4 flex w-full items-center gap-3 rounded-xl bg-blue-50 px-3 py-2.5 text-left"
          >
            <span className="inline-flex h-9 w-9 shrink-0 items-center justify-center rounded-lg bg-blue-600 text-white">
              <FiMapPin />
            </span>
            <span className="min-w-0">
              <span className="block truncate text-sm font-semibold text-slate-800">
                {position.project_name}
              </span>
              <span className="block text-xs capitalize text-slate-500">
                {position.location}
              </span>
            </span>
          </button>
        )}
      </header>

      <section className="space-y-3 px-4 py-5">
        {isLoading &&
          [...Array(3)].map((_, index) => (
            <ItemCardSkeleton key={index} />
          ))}

        {!isLoading &&
          transports.map((transport) => (
            <ProjectTransportCard
              key={transport.project_transport_id}
              transport={transport}
              showProject={false}
            />
          ))}

        {!isLoading && transports.length === 0 && (
          <div className="rounded-2xl border border-dashed border-slate-300 bg-white px-5 py-12 text-center">
            <p className="font-semibold text-slate-700">
              {error || "Belum ada data transport"}
            </p>
            <p className="mt-1 text-sm text-slate-400">
              Tarik data terbaru dengan tombol refresh.
            </p>
          </div>
        )}
      </section>
    </main>
  );
}
