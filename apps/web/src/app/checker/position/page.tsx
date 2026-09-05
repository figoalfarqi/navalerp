"use client";

import Button from "@/components/form/Button";
import SelectField, { SelectOption } from "@/components/form/SelectField";
import MobilePageLoader from "@/components/mobile/MobilePageLoader";
import { useToast } from "@/components/ToastContext";
import { useAuth } from "@/context/AuthContext";
import { useCheckerLocalState } from "@/hooks/useCheckerLocalState";
import { useCheckerSettings } from "@/hooks/useCheckerSettings";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { CheckerLocation } from "@/types/checkerPosition.type";
import { Project } from "@/types/project.type";
import { ProjectCheckerAssignment } from "@/types/projectCheckerAssignment.type";
import { ProjectRoute } from "@/types/projectRoute.type";
import { useRouter } from "next/navigation";
import { useEffect, useMemo, useRef, useState } from "react";
import { FiClock, FiMapPin, FiUser } from "react-icons/fi";

type AssignedProject = {
  project: Project;
  isDefault: boolean;
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

const normalizeAssignedProjects = (rows: unknown[]): AssignedProject[] =>
  rows
    .map((row) => {
      const value = row as Partial<ProjectCheckerAssignment> &
        Partial<Project> & {
          project?: Project;
        };
      const project = value.project ?? (value as Project);
      if (!project?.project_id) return null;
      return {
        project,
        isDefault: Number(value.is_default ?? project.is_default ?? 0) === 1,
      };
    })
    .filter((row): row is AssignedProject => row !== null);

const locationLabel: Record<CheckerLocation, string> = {
  mine: "Mine",
  vessel: "Vessel",
  stockpile: "Stockpile",
  client: "Client",
};

const getValidLocations = (
  project: Project | undefined,
  routes: ProjectRoute[],
): CheckerLocation[] => {
  if (!project) return [];

  const source: CheckerLocation = project.route_type.startsWith("VESSEL")
    ? "vessel"
    : "mine";
  const locations = new Set<CheckerLocation>();

  routes.forEach((route) => {
    if (route.route_type === "SOURCE_TO_CLIENT") {
      locations.add(source);
      locations.add("client");
    }
    if (route.route_type === "SOURCE_TO_STOCKPILE") {
      locations.add(source);
      locations.add("stockpile");
    }
    if (route.route_type === "STOCKPILE_TO_CLIENT") {
      locations.add("stockpile");
      locations.add("client");
    }
  });

  if (locations.size === 0) {
    locations.add(source);
    if (project.route_type.includes("STOCKPILE")) locations.add("stockpile");
    locations.add("client");
  }

  return [...locations];
};

const formatRemaining = (seconds: number) => {
  const minutes = Math.floor(seconds / 60);
  const remainingSeconds = seconds % 60;
  return `${String(minutes).padStart(2, "0")}:${String(
    remainingSeconds,
  ).padStart(2, "0")}`;
};

export default function CheckerPositionPage() {
  const router = useRouter();
  const { showToast } = useToast();
  const { checkerPayload } = useAuth();
  const { getAPI } = useFetchAPI();
  const showToastRef = useRef(showToast);
  useEffect(() => {
    showToastRef.current = showToast;
  }, [showToast]);

  const { locationCooldownMinutes, truckCooldownMinutes } =
    useCheckerSettings();
  const {
    isLoaded: isLocalStateLoaded,
    position,
    locationRemainingSeconds,
    savePosition,
  } = useCheckerLocalState({
    checkerId: checkerPayload.app_user_id,
    locationCooldownMinutes,
    truckCooldownMinutes,
  });

  const [projects, setProjects] = useState<AssignedProject[]>([]);
  const [routes, setRoutes] = useState<ProjectRoute[]>([]);
  const [selectedProjectId, setSelectedProjectId] = useState<number | null>(
    null,
  );
  const [selectedLocation, setSelectedLocation] =
    useState<CheckerLocation | null>(null);
  const [isLoadingProjects, setIsLoadingProjects] = useState(true);
  const [isLoadingRoutes, setIsLoadingRoutes] = useState(false);
  const [isSaving, setIsSaving] = useState(false);

  useEffect(() => {
    let cancelled = false;

    const loadProjects = async () => {
      setIsLoadingProjects(true);
      const response = await getAPI<unknown>(
        `${process.env.NEXT_PUBLIC_API_BASE_URL}/checker/project?limit=399`,
        { authToken: "checker" },
      );

      if (cancelled) return;

      if (response.code === 200 && response.data) {
        setProjects(
          normalizeAssignedProjects(readItems<unknown>(response.data)),
        );
      } else {
        setProjects([]);
        showToastRef.current(
          3000,
          "error",
          response.message || "Gagal memuat project checker",
        );
      }
      setIsLoadingProjects(false);
    };

    void loadProjects();
    return () => {
      cancelled = true;
    };
  }, [getAPI]);

  useEffect(() => {
    if (!isLocalStateLoaded || projects.length === 0 || selectedProjectId) {
      return;
    }

    const storedProject = projects.find(
      ({ project }) => project.project_id === position?.project_id,
    );
    const initialProject =
      storedProject ?? projects.find((project) => project.isDefault) ?? projects[0];

    const frame = window.requestAnimationFrame(() =>
      setSelectedProjectId(initialProject.project.project_id),
    );
    return () => window.cancelAnimationFrame(frame);
  }, [isLocalStateLoaded, position?.project_id, projects, selectedProjectId]);

  useEffect(() => {
    if (!selectedProjectId) return;

    let cancelled = false;
    const loadRoutes = async () => {
      setIsLoadingRoutes(true);
      const response = await getAPI<unknown>(
        `${process.env.NEXT_PUBLIC_API_BASE_URL}/checker/project_route?project_id=${selectedProjectId}&limit=50`,
        { authToken: "checker" },
      );

      if (cancelled) return;
      setRoutes(
        response.code === 200 && response.data
          ? readItems<ProjectRoute>(response.data)
          : [],
      );
      setIsLoadingRoutes(false);
    };

    void loadRoutes();
    return () => {
      cancelled = true;
    };
  }, [getAPI, selectedProjectId]);

  const selectedProject = projects.find(
    ({ project }) => project.project_id === selectedProjectId,
  )?.project;
  const validLocations = useMemo(
    () => getValidLocations(selectedProject, routes),
    [routes, selectedProject],
  );

  useEffect(() => {
    if (!selectedProject || isLoadingRoutes) return;

    if (
      position?.project_id === selectedProject.project_id &&
      validLocations.includes(position.location)
    ) {
      const frame = window.requestAnimationFrame(() =>
        setSelectedLocation(position.location),
      );
      return () => window.cancelAnimationFrame(frame);
    }

    const frame = window.requestAnimationFrame(() =>
      setSelectedLocation((current) =>
        current && validLocations.includes(current) ? current : null,
      ),
    );
    return () => window.cancelAnimationFrame(frame);
  }, [isLoadingRoutes, position, selectedProject, validLocations]);

  const projectOptions: SelectOption[] = projects.map(
    ({ project, isDefault }) => ({
      value: project.project_id,
      label: `${project.project_code} — ${project.project_name}${
        isDefault ? " (Default)" : ""
      }`,
    }),
  );

  const locationOptions: SelectOption[] = validLocations.map((location) => ({
    value: location,
    label: locationLabel[location],
  }));

  const handleSave = () => {
    if (!selectedProject || !selectedLocation) {
      showToast(3000, "error", "Pilih project dan lokasi terlebih dahulu");
      return;
    }

    setIsSaving(true);
    const result = savePosition({
      project_id: selectedProject.project_id,
      project_code: selectedProject.project_code,
      project_name: selectedProject.project_name,
      project_route_type: selectedProject.route_type,
      location: selectedLocation,
    });
    setIsSaving(false);

    if (!result.ok) {
      showToast(
        4000,
        "error",
        `Lokasi baru dapat diganti dalam ${formatRemaining(
          result.remainingSeconds,
        )}`,
      );
      return;
    }

    showToast(2000, "success", "Project dan lokasi tersimpan di perangkat");
    router.replace("/checker/transport");
  };

  if (isLoadingProjects || !isLocalStateLoaded) {
    return <MobilePageLoader />;
  }

  return (
    <main className="mx-auto min-h-dvh max-w-md bg-slate-50 px-5 pb-8 pt-8">
      <div className="mb-7">
        <div className="mb-3 flex items-center justify-between">
          <div className="inline-flex h-12 w-12 items-center justify-center rounded-2xl bg-blue-600 text-white shadow-lg shadow-blue-200">
            <FiMapPin size={24} />
          </div>
          <button
            type="button"
            onClick={() => router.push("/checker/account")}
            className="inline-flex min-h-11 items-center gap-2 rounded-xl border border-slate-200 bg-white px-3 text-sm font-semibold text-slate-600 shadow-sm"
          >
            <FiUser size={18} />
            Akun
          </button>
        </div>
        <h1 className="text-2xl font-bold text-slate-900">Posisi checker</h1>
        <p className="mt-1 text-sm leading-6 text-slate-500">
          Pilih project terlebih dahulu, lalu tentukan titik kerja Anda.
        </p>
      </div>

      <section className="space-y-5 rounded-2xl border border-slate-200 bg-white p-5 shadow-sm">
        {projects.length === 0 ? (
          <div className="rounded-xl bg-amber-50 p-4 text-sm text-amber-800">
            Belum ada project aktif yang ditugaskan kepada akun ini.
          </div>
        ) : (
          <>
            <SelectField
              id="checker_project"
              label="Project"
              required
              value={selectedProjectId ?? ""}
              options={projectOptions}
              onChange={(value) => {
                setSelectedProjectId(value ? Number(value) : null);
                setSelectedLocation(null);
                if (!value) setRoutes([]);
              }}
              className="min-h-12 w-full rounded-lg"
              placeholder="Pilih project"
            />

            <SelectField
              id="checker_location"
              label="Lokasi checker"
              required
              value={selectedLocation ?? ""}
              options={locationOptions}
              onChange={(value) =>
                setSelectedLocation((value as CheckerLocation | null) ?? null)
              }
              disabled={isLoadingRoutes}
              className="min-h-12 w-full rounded-lg"
              placeholder={
                isLoadingRoutes ? "Memuat endpoint rute..." : "Pilih lokasi"
              }
            />

            {locationRemainingSeconds > 0 && (
              <div className="flex items-start gap-3 rounded-xl border border-amber-200 bg-amber-50 p-3 text-sm text-amber-800">
                <FiClock className="mt-0.5 shrink-0" size={18} />
                <p>
                  Perubahan lokasi dikunci sementara. Sisa waktu{" "}
                  <strong>{formatRemaining(locationRemainingSeconds)}</strong>.
                </p>
              </div>
            )}

            <Button
              id="save_checker_position"
              onClick={handleSave}
              disabled={isSaving || !selectedProjectId || !selectedLocation}
              wrapperClassName="w-full"
              className="min-h-12 w-full rounded-xl"
            >
              {isSaving ? "Menyimpan..." : "Simpan posisi"}
            </Button>
          </>
        )}
      </section>

      {position && (
        <p className="mt-4 text-center text-xs text-slate-400">
          Posisi tersimpan: {position.project_name} ·{" "}
          {locationLabel[position.location]}
        </p>
      )}
    </main>
  );
}
