/* eslint-disable @next/next/no-img-element */
"use client";

import Camera from "@/components/camera";
import Button from "@/components/form/Button";
import SelectField, { SelectOption } from "@/components/form/SelectField";
import TextAreaField from "@/components/form/TextAreaField";
import TextField from "@/components/form/TextField";
import MobilePageLoader from "@/components/mobile/MobilePageLoader";
import { useToast } from "@/components/ToastContext";
import { useAuth } from "@/context/AuthContext";
import { useCheckerLocalState } from "@/hooks/useCheckerLocalState";
import { useCheckerSettings } from "@/hooks/useCheckerSettings";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { useUploadFile } from "@/hooks/useUploadFile";
import { CheckerLocation } from "@/types/checkerPosition.type";
import {
  CreateProjectTransportPayload,
  ProjectTransport,
} from "@/types/projectTransport.type";
import {
  CreateProjectTransportPhotoPayload,
  ProjectTransportPhoto,
} from "@/types/projectTransportPhoto.type";
import {
  CreateProjectTransportStatusPayload,
  ProjectTransportStatus,
  ProjectTransportStatusTypeId,
} from "@/types/projectTransportStatus.type";
import { ProjectRoute } from "@/types/projectRoute.type";
import { ProjectTruckAssignment } from "@/types/projectTruckAssignment.type";
import { resolveFileUrl } from "@/utils/globalUtils";
import { useRouter } from "next/navigation";
import { useCallback, useEffect, useMemo, useRef, useState } from "react";
import {
  FiAlertTriangle,
  FiCamera,
  FiCheckCircle,
  FiClock,
  FiMapPin,
  FiRefreshCw,
  FiTruck,
  FiX,
} from "react-icons/fi";

type EndpointRole = "origin" | "destination";

type SubmitResult = {
  isFraud: boolean;
  note?: string;
  transportNumber: string;
};

const EXPLICIT_STATUS_LABELS: Record<
  Extract<ProjectTransportStatusTypeId, 1 | 4 | 5 | 7>,
  string
> = {
  1: "Tiba di lokasi asal",
  4: "Berangkat menuju tujuan",
  5: "Tiba di lokasi tujuan",
  7: "Selesai bongkar",
};

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

const formatRemaining = (seconds: number) => {
  const minutes = Math.floor(seconds / 60);
  const remainingSeconds = seconds % 60;
  return `${String(minutes).padStart(2, "0")}:${String(
    remainingSeconds,
  ).padStart(2, "0")}`;
};

const getTruck = (assignment: ProjectTruckAssignment) => {
  return {
    truckId: assignment.truck_id,
    licensePlate:
      assignment.truck?.license_plate ??
      assignment.license_plate ??
      `Truck #${assignment.truck_id}`,
    driverId:
      assignment.truck?.driver_id ?? assignment.driver_id ?? undefined,
    vendorId:
      assignment.truck?.vendor_id ?? assignment.vendor_id ?? undefined,
  };
};

const getEndpointRole = (
  location: CheckerLocation,
  route: ProjectRoute,
): EndpointRole | null => {
  if (location === "mine" || location === "vessel") {
    return route.route_type.startsWith("SOURCE_") ? "origin" : null;
  }
  if (location === "client") {
    return route.route_type.endsWith("_CLIENT") ? "destination" : null;
  }
  if (location === "stockpile") {
    if (route.route_type === "SOURCE_TO_STOCKPILE") return "destination";
    if (route.route_type === "STOCKPILE_TO_CLIENT") return "origin";
  }
  return null;
};

const getNextExplicitStatus = (
  endpointRole: EndpointRole,
  statuses: ProjectTransportStatus[],
  hasTransport: boolean,
): 1 | 4 | 5 | 7 | null => {
  const existing = new Set(
    statuses.map((status) => status.project_transport_status_type_id),
  );

  if (endpointRole === "origin") {
    if (
      ([5, 6, 7, 8] as ProjectTransportStatusTypeId[]).some((statusTypeId) =>
        existing.has(statusTypeId),
      )
    ) {
      return null;
    }
    if (!hasTransport || !existing.has(1)) return 1;
    if (!existing.has(4)) return 4;
    return null;
  }

  if (!hasTransport || !existing.has(5)) return 5;
  if (!existing.has(7)) return 7;
  return null;
};

const getStatusSequence = (
  explicitStatus: 1 | 4 | 5 | 7,
): ProjectTransportStatusTypeId[] => {
  if (explicitStatus === 1) return [1, 2];
  if (explicitStatus === 4) return [3, 4];
  if (explicitStatus === 5) return [5, 6];
  return [7, 8];
};

const createTransportNumber = (projectId: number, truckId: number) => {
  const timestamp = new Date()
    .toISOString()
    .replaceAll("-", "")
    .replaceAll(":", "")
    .replace("T", "")
    .replace(/\..+$/, "");
  return `TR-${projectId}-${timestamp}-${truckId}`;
};

export default function CheckerTransportPage() {
  const router = useRouter();
  const { checkerPayload } = useAuth();
  const { showToast } = useToast();
  const { getAPI, postAPI } = useFetchAPI();

  const { locationCooldownMinutes, truckCooldownMinutes } =
    useCheckerSettings();
  const {
    isLoaded: isLocalStateLoaded,
    position,
    getTruckRemainingSeconds,
    lockTruck,
  } = useCheckerLocalState({
    checkerId: checkerPayload.app_user_id,
    locationCooldownMinutes,
    truckCooldownMinutes,
  });
  const {
    captureImage,
    isUploading,
    progress: uploadProgress,
  } = useUploadFile();

  const [routes, setRoutes] = useState<ProjectRoute[]>([]);
  const [assignments, setAssignments] = useState<ProjectTruckAssignment[]>([]);
  const [activeTransports, setActiveTransports] = useState<ProjectTransport[]>(
    [],
  );
  const [selectedRouteId, setSelectedRouteId] = useState<number | null>(null);
  const [selectedTruckId, setSelectedTruckId] = useState<number | null>(null);
  const [selectedTransportId, setSelectedTransportId] = useState<number | null>(
    null,
  );
  const [statuses, setStatuses] = useState<ProjectTransportStatus[]>([]);
  const [cargoWeightTon, setCargoWeightTon] = useState<number | "">("");
  const [cargoVolumeCubic, setCargoVolumeCubic] = useState<number | "">("");
  const [note, setNote] = useState("");
  const [photoUrl, setPhotoUrl] = useState("");
  const [isLoading, setIsLoading] = useState(true);
  const [isLoadingStatuses, setIsLoadingStatuses] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [submitResult, setSubmitResult] = useState<SubmitResult | null>(null);
  const submissionLockRef = useRef(false);

  const loadOperationalData = useCallback(async () => {
    if (!position?.project_id) return;
    setIsLoading(true);

    const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
    const [routeResponse, assignmentResponse, transportResponse] =
      await Promise.all([
        getAPI<unknown>(
          `${baseUrl}/checker/project_route?project_id=${position.project_id}&limit=50`,
          { authToken: "checker" },
        ),
        getAPI<unknown>(
          `${baseUrl}/checker/project_truck_assignment?project_id=${position.project_id}&limit=399`,
          { authToken: "checker" },
        ),
        getAPI<unknown>(
          `${baseUrl}/checker/project_transport?project_id=${position.project_id}&is_completed=0&limit=399&order_by=transported_at&sort=desc`,
          { authToken: "checker" },
        ),
      ]);

    setRoutes(
      routeResponse.code === 200 && routeResponse.data
        ? readItems<ProjectRoute>(routeResponse.data).filter(
            (route) => Number(route.is_active ?? 1) === 1,
          )
        : [],
    );
    setAssignments(
      assignmentResponse.code === 200 && assignmentResponse.data
        ? readItems<ProjectTruckAssignment>(assignmentResponse.data).filter(
            (assignment) => Number(assignment.is_active ?? 1) === 1,
          )
        : [],
    );
    setActiveTransports(
      transportResponse.code === 200 && transportResponse.data
        ? readItems<ProjectTransport>(transportResponse.data).filter(
            (transport) => Number(transport.is_completed ?? 0) !== 1,
          )
        : [],
    );
    setIsLoading(false);
  }, [getAPI, position?.project_id]);

  useEffect(() => {
    if (!isLocalStateLoaded) return;
    if (!position) {
      router.replace("/checker/position");
      return;
    }
    void loadOperationalData();
  }, [isLocalStateLoaded, loadOperationalData, position, router]);

  const validRoutes = useMemo(
    () =>
      routes
        .filter(
          (route) =>
            position && getEndpointRole(position.location, route) !== null,
        )
        .sort((a, b) => a.route_sequence - b.route_sequence),
    [position, routes],
  );

  useEffect(() => {
    if (
      selectedRouteId &&
      validRoutes.some((route) => route.project_route_id === selectedRouteId)
    ) {
      return;
    }
    setSelectedRouteId(validRoutes[0]?.project_route_id ?? null);
  }, [selectedRouteId, validRoutes]);

  const selectedRoute = validRoutes.find(
    (route) => route.project_route_id === selectedRouteId,
  );
  const endpointRole =
    position && selectedRoute
      ? getEndpointRole(position.location, selectedRoute)
      : null;
  const selectedAssignment = assignments.find(
    (assignment) => assignment.truck_id === selectedTruckId,
  );

  const matchingTransports = useMemo(
    () =>
      activeTransports
        .filter(
          (transport) =>
            transport.truck_id === selectedTruckId &&
            transport.project_route_id === selectedRouteId,
        )
        .sort(
          (a, b) =>
            new Date(b.transported_at).getTime() -
            new Date(a.transported_at).getTime(),
        ),
    [activeTransports, selectedRouteId, selectedTruckId],
  );

  useEffect(() => {
    setSelectedTransportId(
      matchingTransports[0]?.project_transport_id ?? null,
    );
  }, [matchingTransports]);

  const selectedTransport = matchingTransports.find(
    (transport) => transport.project_transport_id === selectedTransportId,
  );

  useEffect(() => {
    if (!selectedTransportId) {
      setStatuses([]);
      return;
    }

    let cancelled = false;
    const loadStatuses = async () => {
      setIsLoadingStatuses(true);
      const response = await getAPI<unknown>(
        `${process.env.NEXT_PUBLIC_API_BASE_URL}/checker/project_transport_status?project_transport_id=${selectedTransportId}&limit=50&order_by=status_time&sort=asc`,
        { authToken: "checker" },
      );

      if (cancelled) return;
      const nestedStatuses =
        selectedTransport?.project_transport_statuses ??
        selectedTransport?.statuses ??
        [];
      setStatuses(
        response.code === 200 && response.data
          ? readItems<ProjectTransportStatus>(response.data)
          : nestedStatuses,
      );
      setIsLoadingStatuses(false);
    };

    void loadStatuses();
    return () => {
      cancelled = true;
    };
  }, [getAPI, selectedTransport, selectedTransportId]);

  const nextStatus =
    endpointRole && !isLoadingStatuses
      ? getNextExplicitStatus(endpointRole, statuses, Boolean(selectedTransport))
      : null;
  const hasOriginStatus = statuses.some(
    (status) => status.project_transport_status_type_id === 1,
  );
  const willBeFraud =
    endpointRole === "destination" && nextStatus === 5 && !hasOriginStatus;
  const truckRemainingSeconds = getTruckRemainingSeconds(selectedTruckId);

  const routeOptions: SelectOption[] = validRoutes.map((route) => ({
    value: route.project_route_id,
    label:
      route.route_name ||
      route.route_type
        .replaceAll("_", " ")
        .toLowerCase()
        .replace(/(^|\s)\S/g, (letter) => letter.toUpperCase()),
  }));

  const truckOptions: SelectOption[] = assignments.map((assignment) => {
    const truck = getTruck(assignment);
    const remaining = getTruckRemainingSeconds(truck.truckId);
    return {
      value: truck.truckId,
      label: `${truck.licensePlate}${
        remaining > 0 ? ` · tunggu ${formatRemaining(remaining)}` : ""
      }`,
      row: assignment as unknown as Record<string, unknown>,
    };
  });

  const transportOptions: SelectOption[] = matchingTransports.map(
    (transport) => ({
      value: transport.project_transport_id,
      label: `${transport.transport_number} · ${new Date(
        transport.transported_at,
      ).toLocaleTimeString("id-ID", {
        hour: "2-digit",
        minute: "2-digit",
      })}`,
    }),
  );

  const handleCapture = async (canvas: HTMLCanvasElement) => {
    if (!position || !selectedTruckId) return;
    const licensePlate = selectedAssignment
      ? getTruck(selectedAssignment).licensePlate
      : String(selectedTruckId);

    try {
      await captureImage({
        canvas,
        authToken: "checker",
        filename: `truck-${Date.now()}.jpg`,
        folder: `project_transport/${position.project_code}/${licensePlate}`,
        showOverlay: false,
        onSuccess: (url) => setPhotoUrl(url),
        onError: () =>
          showToast(3000, "error", "Foto gagal diunggah. Foto boleh dikosongkan."),
      });
    } catch {
      // Error sudah ditampilkan melalui callback.
    }
  };

  const createTransport = async () => {
    if (!position || !selectedRoute || !selectedAssignment) return null;
    const truck = getTruck(selectedAssignment);
    const payload: CreateProjectTransportPayload = {
      project_id: position.project_id,
      project_route_id: selectedRoute.project_route_id,
      project_truck_assignment_id:
        selectedAssignment.project_truck_assignment_id,
      truck_id: truck.truckId,
      transport_number: createTransportNumber(
        position.project_id,
        truck.truckId,
      ),
      transported_at: new Date().toISOString(),
      ...(truck.driverId ? { driver_id: truck.driverId } : {}),
      ...(truck.vendorId ? { transport_vendor_id: truck.vendorId } : {}),
    };

    const response = await postAPI<unknown>(
      `${process.env.NEXT_PUBLIC_API_BASE_URL}/checker/project_transport`,
      { authToken: "checker" },
      payload,
    );
    if (![200, 201, 202].includes(response.code) || !response.data) {
      throw new Error(response.message || "Gagal membuat transport");
    }

    return unwrapRecord<ProjectTransport>(response.data);
  };

  const createStatus = async (
    transportId: number,
    statusTypeId: ProjectTransportStatusTypeId,
    explicitStatus: 1 | 4 | 5 | 7,
    fraud: boolean,
    fraudNote: string,
  ) => {
    const isExplicit = statusTypeId === explicitStatus;
    const payload: CreateProjectTransportStatusPayload = {
      project_transport_id: transportId,
      project_transport_status_type_id: statusTypeId,
      status_time: new Date().toISOString(),
      is_fraud: isExplicit && fraud ? 1 : 0,
      ...(isExplicit && cargoVolumeCubic !== ""
        ? { cargo_volume_cubic: Number(cargoVolumeCubic) }
        : {}),
      ...(isExplicit && cargoWeightTon !== ""
        ? { cargo_weight_ton: Number(cargoWeightTon) }
        : {}),
      ...(isExplicit && fraudNote
        ? { project_transport_status_note: fraudNote }
        : isExplicit && note.trim()
          ? { project_transport_status_note: note.trim() }
          : {}),
    };

    const response = await postAPI<unknown>(
      `${process.env.NEXT_PUBLIC_API_BASE_URL}/checker/project_transport_status`,
      { authToken: "checker" },
      payload,
    );
    if (![200, 201, 202].includes(response.code) || !response.data) {
      throw new Error(response.message || `Gagal menyimpan status ${statusTypeId}`);
    }
    return unwrapRecord<ProjectTransportStatus>(response.data);
  };

  const savePhoto = async (statusId: number) => {
    if (!photoUrl) return true;
    const payload: CreateProjectTransportPhotoPayload = {
      project_transport_status_id: statusId,
      photo_url: photoUrl,
      photo_type_id: 2,
      photo_description: "Foto truck dari checker",
    };
    const response = await postAPI<ProjectTransportPhoto>(
      `${process.env.NEXT_PUBLIC_API_BASE_URL}/checker/project_transport_photo`,
      { authToken: "checker" },
      payload,
    );
    return [200, 201, 202].includes(response.code);
  };

  const handleSubmit = async () => {
    if (submissionLockRef.current) return;
    if (
      !position ||
      !selectedRoute ||
      !selectedAssignment ||
      !selectedTruckId ||
      !nextStatus
    ) {
      showToast(3000, "error", "Data truck, rute, atau status belum lengkap");
      return;
    }
    if (truckRemainingSeconds > 0) {
      showToast(
        3000,
        "error",
        `Truck ini dapat diinput lagi dalam ${formatRemaining(
          truckRemainingSeconds,
        )}`,
      );
      return;
    }
    const measurements = [cargoWeightTon, cargoVolumeCubic].filter(
      (value): value is number => value !== "",
    );
    if (
      measurements.some(
        (measurement) =>
          !Number.isFinite(measurement) || measurement < 0,
      )
    ) {
      showToast(3000, "error", "Berat dan volume tidak boleh bernilai negatif");
      return;
    }

    submissionLockRef.current = true;
    setIsSubmitting(true);
    setSubmitResult(null);
    try {
      const transport = selectedTransport ?? (await createTransport());
      if (!transport) throw new Error("Transport tidak berhasil dibuat");

      const existingTypes = new Set(
        statuses.map((status) => status.project_transport_status_type_id),
      );
      const fraudNote = willBeFraud
        ? [
            "Status tujuan dilaporkan tanpa status arrived at origin.",
            note.trim(),
          ]
            .filter(Boolean)
            .join(" ")
        : note.trim();
      let explicitStatusRecord: ProjectTransportStatus | null = null;

      for (const statusTypeId of getStatusSequence(nextStatus)) {
        if (existingTypes.has(statusTypeId)) continue;
        const createdStatus = await createStatus(
          transport.project_transport_id,
          statusTypeId,
          nextStatus,
          willBeFraud,
          fraudNote,
        );
        if (statusTypeId === nextStatus) explicitStatusRecord = createdStatus;
      }

      const fallbackExplicitStatus = statuses.find(
        (status) => status.project_transport_status_type_id === nextStatus,
      );
      const photoStatusId =
        explicitStatusRecord?.project_transport_status_id ??
        fallbackExplicitStatus?.project_transport_status_id;
      const photoSaved = photoStatusId ? await savePhoto(photoStatusId) : !photoUrl;

      lockTruck(selectedTruckId);
      setSubmitResult({
        isFraud: willBeFraud,
        note: willBeFraud ? fraudNote : undefined,
        transportNumber: transport.transport_number,
      });
      setCargoVolumeCubic("");
      setCargoWeightTon("");
      setNote("");
      setPhotoUrl("");
      showToast(
        3000,
        "success",
        photoSaved
          ? "Status transport berhasil disimpan"
          : "Status tersimpan, tetapi foto gagal dicatat",
      );
      await loadOperationalData();
      setSelectedTruckId(null);
      setSelectedTransportId(null);
      setStatuses([]);
    } catch (error) {
      await loadOperationalData();
      showToast(
        4000,
        "error",
        error instanceof Error
          ? error.message
          : "Gagal menyimpan status transport",
      );
    } finally {
      submissionLockRef.current = false;
      setIsSubmitting(false);
    }
  };

  if (!isLocalStateLoaded || isLoading) {
    return <MobilePageLoader />;
  }

  if (!position) return null;

  return (
    <main className="mx-auto min-h-dvh max-w-md bg-slate-50 pb-28">
      <header className="bg-gradient-to-br from-blue-700 to-blue-500 px-5 pb-6 pt-7 text-white">
        <div className="flex items-start justify-between gap-3">
          <div>
            <p className="text-xs font-medium uppercase tracking-[0.18em] text-blue-100">
              Input checker
            </p>
            <h1 className="mt-1 text-2xl font-bold">Laporan truck</h1>
          </div>
          <Button
            id="refresh_checker_transport"
            variant="gray-ghost"
            onClick={() => void loadOperationalData()}
            className="h-10 w-10 rounded-full bg-white/15 px-0! py-0! text-white hover:bg-white/25"
          >
            <FiRefreshCw className="h-5 w-4" />
          </Button>
        </div>
        <div className="mt-5 flex items-center justify-between gap-3 rounded-xl bg-white/12 px-3 py-2.5 text-sm backdrop-blur">
          <div className="min-w-0">
            <p className="truncate font-semibold">{position.project_name}</p>
            <p className="truncate text-xs text-blue-100">
              {position.project_code}
            </p>
          </div>
          <button
            type="button"
            onClick={() => router.push("/checker/position")}
            className="flex shrink-0 items-center gap-1 rounded-lg bg-white/15 px-2.5 py-1.5 text-xs font-semibold"
          >
            <FiMapPin />
            {position.location}
          </button>
        </div>
      </header>

      <div className="-mt-2 space-y-4 px-4">
        {submitResult && (
          <section
            className={`rounded-2xl border p-4 shadow-sm ${
              submitResult.isFraud
                ? "border-red-200 bg-red-50"
                : "border-emerald-200 bg-emerald-50"
            }`}
          >
            <div className="flex items-start gap-3">
              {submitResult.isFraud ? (
                <FiAlertTriangle
                  className="mt-0.5 shrink-0 text-red-600"
                  size={20}
                />
              ) : (
                <FiCheckCircle
                  className="mt-0.5 shrink-0 text-emerald-600"
                  size={20}
                />
              )}
              <div>
                <p className="font-semibold text-slate-900">
                  {submitResult.isFraud
                    ? "Laporan tersimpan sebagai fraud"
                    : "Laporan berhasil"}
                </p>
                <p className="mt-0.5 text-xs text-slate-600">
                  {submitResult.transportNumber}
                </p>
                {submitResult.note && (
                  <p className="mt-2 text-sm text-red-700">
                    {submitResult.note}
                  </p>
                )}
              </div>
            </div>
          </section>
        )}

        <section className="space-y-4 rounded-2xl border border-slate-200 bg-white p-4 shadow-sm">
          <div className="flex items-center gap-2">
            <span className="inline-flex h-9 w-9 items-center justify-center rounded-xl bg-blue-50 text-blue-600">
              <FiTruck size={19} />
            </span>
            <div>
              <h2 className="font-semibold text-slate-900">Truck dan rute</h2>
              <p className="text-xs text-slate-500">
                Hanya assignment aktif pada project ini
              </p>
            </div>
          </div>

          <SelectField
            id="checker_route"
            label="Rute"
            required
            value={selectedRouteId ?? ""}
            options={routeOptions}
            onChange={(value) => {
              setSelectedRouteId(value ? Number(value) : null);
              setStatuses([]);
            }}
            className="min-h-12 w-full rounded-lg"
            placeholder="Pilih rute"
          />

          <SelectField
            id="checker_truck"
            label="Plat nomor truck"
            required
            value={selectedTruckId ?? ""}
            options={truckOptions}
            onChange={(value) => {
              setSelectedTruckId(value ? Number(value) : null);
              setSubmitResult(null);
              setStatuses([]);
            }}
            className="min-h-12 w-full rounded-lg"
            placeholder="Pilih plat nomor"
          />

          {transportOptions.length > 1 && (
            <SelectField
              id="checker_active_transport"
              label="Transport aktif"
              value={selectedTransportId ?? ""}
              options={transportOptions}
              onChange={(value) =>
                setSelectedTransportId(value ? Number(value) : null)
              }
              className="min-h-12 w-full rounded-lg"
              placeholder="Pilih transport aktif"
            />
          )}

          {selectedTruckId && truckRemainingSeconds > 0 && (
            <div className="flex gap-2 rounded-xl bg-amber-50 p-3 text-sm text-amber-800">
              <FiClock className="mt-0.5 shrink-0" />
              <span>
                Truck dikunci selama {truckCooldownMinutes} menit. Sisa{" "}
                <strong>{formatRemaining(truckRemainingSeconds)}</strong>.
              </span>
            </div>
          )}

          {selectedTruckId && nextStatus && (
            <div className="rounded-xl border border-blue-100 bg-blue-50 p-3">
              <p className="text-xs font-medium uppercase tracking-wide text-blue-500">
                Status yang dilaporkan
              </p>
              <p className="mt-1 font-semibold text-blue-900">
                {nextStatus}. {EXPLICIT_STATUS_LABELS[nextStatus]}
              </p>
            </div>
          )}

          {selectedTruckId && !nextStatus && !isLoadingStatuses && (
            <div className="rounded-xl bg-slate-100 p-3 text-sm text-slate-600">
              Tidak ada transisi status yang valid untuk truck pada lokasi ini.
            </div>
          )}
        </section>

        <section className="space-y-4 rounded-2xl border border-slate-200 bg-white p-4 shadow-sm">
          <div>
            <h2 className="font-semibold text-slate-900">Muatan</h2>
            <p className="text-xs text-slate-500">
              Berat, volume, dan foto bersifat opsional.
            </p>
          </div>

          <div className="grid grid-cols-2 gap-3">
            <TextField
              id="cargo_weight_ton"
              type="number"
              label="Berat (ton)"
              value={cargoWeightTon}
              onChange={(value) =>
                setCargoWeightTon(value === "" ? "" : Number(value))
              }
              className="min-h-12 w-full rounded-lg"
              placeholder="0"
            />
            <TextField
              id="cargo_volume_cubic"
              type="number"
              label="Volume (m³)"
              value={cargoVolumeCubic}
              onChange={(value) =>
                setCargoVolumeCubic(value === "" ? "" : Number(value))
              }
              className="min-h-12 w-full rounded-lg"
              placeholder="0"
            />
          </div>

          <TextAreaField
            id="project_transport_status_note"
            label="Catatan"
            value={note}
            onChange={setNote}
            className="w-full rounded-lg"
            placeholder="Catatan tambahan (opsional)"
          />

          {willBeFraud && (
            <div className="flex items-start gap-2 rounded-xl border border-red-200 bg-red-50 p-3 text-sm text-red-700">
              <FiAlertTriangle className="mt-0.5 shrink-0" />
              <p>
                Truck tiba di tujuan tanpa laporan origin. Sistem akan menyimpan{" "}
                <strong>is_fraud = 1</strong> beserta catatan otomatis.
              </p>
            </div>
          )}

          <div className="space-y-2">
            <p className="text-sm font-medium text-slate-700">Foto truck</p>
            {photoUrl ? (
              <div className="relative overflow-hidden rounded-xl border border-slate-200">
                <img
                  src={resolveFileUrl(photoUrl)}
                  alt="Foto truck"
                  className="h-52 w-full object-cover"
                />
                <button
                  type="button"
                  onClick={() => setPhotoUrl("")}
                  className="absolute right-2 top-2 inline-flex h-9 w-9 items-center justify-center rounded-full bg-black/55 text-white"
                  aria-label="Hapus foto truck"
                >
                  <FiX />
                </button>
              </div>
            ) : (
              <Camera
                onCapture={(canvas) => void handleCapture(canvas)}
                buttonText="Ambil foto truck (opsional)"
                isBlurCheck
              >
                <div className="pointer-events-none absolute inset-4 rounded-xl border-2 border-dashed border-white/70" />
              </Camera>
            )}
            {isUploading && (
              <div className="flex items-center gap-2 text-xs text-blue-600">
                <FiCamera />
                Mengunggah foto {Math.round(uploadProgress)}%
              </div>
            )}
          </div>
        </section>

        <Button
          id="submit_checker_transport"
          onClick={() => void handleSubmit()}
          disabled={
            isSubmitting ||
            isUploading ||
            isLoadingStatuses ||
            !selectedTruckId ||
            !selectedRouteId ||
            !nextStatus ||
            truckRemainingSeconds > 0
          }
          wrapperClassName="w-full"
          className="min-h-14 w-full rounded-2xl text-base shadow-lg shadow-blue-200"
        >
          {isSubmitting
            ? "Menyimpan laporan..."
            : nextStatus
              ? `Laporkan: ${EXPLICIT_STATUS_LABELS[nextStatus]}`
              : "Pilih truck"}
        </Button>
      </div>
    </main>
  );
}
