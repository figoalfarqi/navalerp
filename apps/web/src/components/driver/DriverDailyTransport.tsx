"use client";

import ItemCardSkeleton from "@/components/driver/ItemCardSkeleton";
import Button from "@/components/form/Button";
import Modal from "@/components/Modal";
import ProjectTransportCard, {
  getProjectTransportStatusLabel,
  getTransportStatuses,
} from "@/components/mobile/ProjectTransportCard";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { ProjectTransport } from "@/types/projectTransport.type";
import { formatDateTime, toLocalDateString } from "@/utils/dateTime";
import { useCallback, useEffect, useState } from "react";
import {
  FiCalendar,
  FiMap,
  FiRefreshCw,
  FiTrendingUp,
  FiTruck,
} from "react-icons/fi";

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

const unwrapRecord = <T,>(value: unknown): T | null => {
  if (!value || typeof value !== "object") return null;
  const record = value as Record<string, unknown>;
  if (record.data && typeof record.data === "object") return record.data as T;
  return value as T;
};

const getRouteEndpoints = (transport: ProjectTransport) => {
  const project = transport.project;
  const route = transport.project_route;
  const routeType = route?.route_type ?? transport.route_type;

  const originName =
    transport.origin_name ?? transport.origin_location_name;
  const destinationName =
    transport.destination_name ?? transport.destination_location_name;

  if (originName || destinationName) {
    return [
      originName || "Lokasi asal",
      destinationName || "Lokasi tujuan",
    ];
  }

  if (!project || !route) {
    if (routeType === "SOURCE_TO_STOCKPILE") return ["Sumber", "Stockpile"];
    if (routeType === "STOCKPILE_TO_CLIENT") return ["Stockpile", "Client"];
    if (routeType === "SOURCE_TO_CLIENT") return ["Sumber", "Client"];
    return ["Lokasi asal", "Lokasi tujuan"];
  }

  const source =
    project.mine?.mine_name ??
    project.vessel_cargo?.vessel?.vessel_name ??
    project.vessel_cargo?.port?.port_name ??
    "Sumber";
  const stockpile =
    project.stockpile_cargo?.stockpile?.stockpile_name ?? "Stockpile";
  const client =
    project.client_destination?.client_destination_name ?? "Client";

  if (routeType === "SOURCE_TO_STOCKPILE") return [source, stockpile];
  if (routeType === "STOCKPILE_TO_CLIENT") return [stockpile, client];
  return [source, client];
};

const sumMeasurement = (
  transports: ProjectTransport[],
  deliveredKey: "delivered_weight_ton" | "delivered_volume_cubic",
  loadedKey: "loaded_weight_ton" | "loaded_volume_cubic",
) =>
  transports.reduce(
    (total, transport) =>
      total + Number(transport[deliveredKey] ?? transport[loadedKey] ?? 0),
    0,
  );

export default function DriverDailyTransport() {
  const { getAPI } = useFetchAPI();

  const [selectedDate, setSelectedDate] = useState(() =>
    toLocalDateString(new Date()),
  );
  const [transports, setTransports] = useState<ProjectTransport[]>([]);
  const [selectedTransport, setSelectedTransport] =
    useState<ProjectTransport | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [isLoadingDetail, setIsLoadingDetail] = useState(false);
  const [error, setError] = useState("");

  const loadTransports = useCallback(async () => {
    setIsLoading(true);
    setError("");
    const response = await getAPI<unknown>(
      `${process.env.NEXT_PUBLIC_API_BASE_URL}/driver/project_transport?date=${selectedDate}&limit=100&order_by=transported_at&sort=desc`,
      { authToken: "driver" },
    );

    if (response.code === 200 && response.data) {
      setTransports(readItems<ProjectTransport>(response.data));
    } else {
      setTransports([]);
      setError(response.message || "Gagal memuat perjalanan");
    }
    setIsLoading(false);
  }, [getAPI, selectedDate]);

  useEffect(() => {
    const frame = window.requestAnimationFrame(() => void loadTransports());
    return () => window.cancelAnimationFrame(frame);
  }, [loadTransports]);

  const openDetail = async (transport: ProjectTransport) => {
    setSelectedTransport(transport);
    setIsLoadingDetail(true);
    const response = await getAPI<unknown>(
      `${process.env.NEXT_PUBLIC_API_BASE_URL}/driver/project_transport/${transport.project_transport_id}`,
      { authToken: "driver" },
    );
    if (response.code === 200 && response.data) {
      setSelectedTransport(
        unwrapRecord<ProjectTransport>(response.data) ?? transport,
      );
    }
    setIsLoadingDetail(false);
  };

  const totalWeight = sumMeasurement(
    transports,
    "delivered_weight_ton",
    "loaded_weight_ton",
  );
  const totalVolume = sumMeasurement(
    transports,
    "delivered_volume_cubic",
    "loaded_volume_cubic",
  );
  const routeEndpoints = selectedTransport
    ? getRouteEndpoints(selectedTransport)
    : ["", ""];
  const selectedStatuses = selectedTransport
    ? [...getTransportStatuses(selectedTransport)].sort(
        (a, b) =>
          new Date(a.status_time).getTime() -
          new Date(b.status_time).getTime(),
      )
    : [];

  return (
    <main className="mx-auto min-h-dvh max-w-md bg-slate-50 pb-28">
      <header className="bg-gradient-to-br from-slate-900 via-slate-800 to-blue-900 px-5 pb-7 pt-7 text-white">
        <div className="flex items-start justify-between">
          <div>
            <p className="text-xs font-semibold uppercase tracking-[0.17em] text-blue-200">
              Driver
            </p>
            <h1 className="mt-1 text-2xl font-bold">Perjalanan harian</h1>
            <p className="mt-1 text-sm text-slate-300">
              Rute dan muatan truck Anda
            </p>
          </div>
          <Button
            id="reload_driver_transport"
            variant="gray-ghost"
            onClick={() => void loadTransports()}
            disabled={isLoading}
            className="h-10 w-10 rounded-full bg-white/10 p-0 text-white hover:bg-white/20"
          >
            <FiRefreshCw className={isLoading ? "animate-spin" : ""} />
          </Button>
        </div>

        <label className="mt-5 block">
          <span className="mb-1.5 flex items-center gap-2 text-xs font-medium text-slate-300">
            <FiCalendar />
            Tanggal perjalanan
          </span>
          <input
            type="date"
            value={selectedDate}
            onChange={(event) => setSelectedDate(event.target.value)}
            className="min-h-12 w-full rounded-xl border border-white/15 bg-white/10 px-3 text-base text-white outline-none [color-scheme:dark] focus:border-blue-300"
          />
        </label>
      </header>

      <section className="-mt-3 grid grid-cols-3 gap-2 px-4">
        <div className="rounded-2xl border border-slate-200 bg-white p-3 shadow-sm">
          <FiTruck className="text-blue-600" />
          <p className="mt-2 text-lg font-bold text-slate-900">
            {transports.length}
          </p>
          <p className="text-[10px] text-slate-500">Perjalanan</p>
        </div>
        <div className="rounded-2xl border border-slate-200 bg-white p-3 shadow-sm">
          <FiTrendingUp className="text-emerald-600" />
          <p className="mt-2 truncate text-lg font-bold text-slate-900">
            {totalWeight.toLocaleString("id-ID", {
              maximumFractionDigits: 3,
            })}
          </p>
          <p className="text-[10px] text-slate-500">Total ton</p>
        </div>
        <div className="rounded-2xl border border-slate-200 bg-white p-3 shadow-sm">
          <FiMap className="text-violet-600" />
          <p className="mt-2 truncate text-lg font-bold text-slate-900">
            {totalVolume.toLocaleString("id-ID", {
              maximumFractionDigits: 3,
            })}
          </p>
          <p className="text-[10px] text-slate-500">Total m³</p>
        </div>
      </section>

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
              onClick={() => void openDetail(transport)}
            />
          ))}

        {!isLoading && transports.length === 0 && (
          <div className="rounded-2xl border border-dashed border-slate-300 bg-white px-6 py-12 text-center">
            <FiTruck className="mx-auto text-slate-300" size={34} />
            <p className="mt-3 font-semibold text-slate-700">
              {error || "Tidak ada perjalanan pada tanggal ini"}
            </p>
            <p className="mt-1 text-sm text-slate-400">
              Pilih tanggal lain untuk melihat data.
            </p>
          </div>
        )}
      </section>

      <Modal
        isOpen={selectedTransport !== null}
        title="Detail perjalanan"
        confirmText=""
        cancelText="Tutup"
        onCancel={() => setSelectedTransport(null)}
      >
        {selectedTransport && (
          <div className="space-y-4">
            {isLoadingDetail && (
              <div className="h-1 overflow-hidden rounded-full bg-blue-100">
                <div className="h-full w-1/2 animate-pulse rounded-full bg-blue-600" />
              </div>
            )}

            <div>
              <p className="text-lg font-bold text-slate-900">
                {selectedTransport.truck?.license_plate ||
                  selectedTransport.license_plate ||
                  selectedTransport.transport_number}
              </p>
              <p className="text-sm text-slate-500">
                {selectedTransport.project?.project_name ||
                  selectedTransport.project_name ||
                  selectedTransport.project_code ||
                  `Project #${selectedTransport.project_id}`}
              </p>
            </div>

            <div className="rounded-xl bg-slate-50 p-4">
              <p className="text-xs font-semibold uppercase tracking-wide text-slate-400">
                Rute
              </p>
              <p className="mt-1 text-sm font-semibold text-slate-800">
                {selectedTransport.project_route?.route_name ||
                  selectedTransport.route_name ||
                  selectedTransport.project_route?.route_type?.replaceAll(
                    "_",
                    " ",
                  ) ||
                  selectedTransport.route_type?.replaceAll("_", " ") ||
                  `Rute #${selectedTransport.project_route_id}`}
              </p>
              <div className="mt-3 flex items-center gap-3">
                <span className="h-3 w-3 shrink-0 rounded-full bg-blue-600" />
                <span className="font-medium text-slate-800">
                  {routeEndpoints[0]}
                </span>
              </div>
              <div className="ml-[5px] h-7 border-l-2 border-dashed border-slate-300" />
              <div className="flex items-center gap-3">
                <span className="h-3 w-3 shrink-0 rounded-full bg-emerald-500" />
                <span className="font-medium text-slate-800">
                  {routeEndpoints[1]}
                </span>
              </div>
              {(selectedTransport.project_route?.distance_km ??
                selectedTransport.distance_km) != null && (
                <p className="mt-3 text-xs text-slate-500">
                  Jarak:{" "}
                  {selectedTransport.project_route?.distance_km ??
                    selectedTransport.distance_km}{" "}
                  km
                </p>
              )}
            </div>

            <div className="grid grid-cols-2 gap-3">
              <div className="rounded-xl border border-slate-200 p-3">
                <p className="text-xs text-slate-400">Berat muat / terkirim</p>
                <p className="mt-1 font-semibold text-slate-800">
                  {selectedTransport.loaded_weight_ton ?? "-"} /{" "}
                  {selectedTransport.delivered_weight_ton ?? "-"} ton
                </p>
              </div>
              <div className="rounded-xl border border-slate-200 p-3">
                <p className="text-xs text-slate-400">Volume muat / terkirim</p>
                <p className="mt-1 font-semibold text-slate-800">
                  {selectedTransport.loaded_volume_cubic ?? "-"} /{" "}
                  {selectedTransport.delivered_volume_cubic ?? "-"} m³
                </p>
              </div>
            </div>

            {selectedStatuses.length > 0 && (
              <div>
                <p className="mb-2 text-sm font-semibold text-slate-800">
                  Timeline status
                </p>
                <div className="space-y-2">
                  {selectedStatuses.map((status) => (
                    <div
                      key={status.project_transport_status_id}
                      className="flex items-start justify-between gap-3 rounded-lg bg-slate-50 px-3 py-2 text-sm"
                    >
                      <span className="font-medium text-slate-700">
                        {getProjectTransportStatusLabel(
                          status.project_transport_status_type_id,
                        )}
                      </span>
                      <span className="text-right">
                        <span className="block text-xs text-slate-400">
                          {formatDateTime(status.status_time)}
                        </span>
                        {status.is_fraud === 1 && (
                          <span className="mt-1 block text-xs font-semibold text-red-600">
                            Fraud
                          </span>
                        )}
                      </span>
                    </div>
                  ))}
                </div>
                {selectedStatuses.some((status) => status.is_fraud === 1) && (
                  <div className="mt-3 rounded-lg border border-red-200 bg-red-50 p-3 text-sm text-red-700">
                    {selectedStatuses.find((status) => status.is_fraud === 1)
                      ?.project_transport_status_note ||
                      "Status tujuan tercatat tanpa status origin."}
                  </div>
                )}
              </div>
            )}
          </div>
        )}
      </Modal>
    </main>
  );
}
