"use client";

import { useEffect, useMemo, useRef, useState } from "react";
import { FaArrowRotateRight } from "react-icons/fa6";
import AdminDashboardChart from "@/components/admin/dashboard/AdminDashboardChart";
import AdminSummaryCards from "@/components/admin/dashboard/AdminSummaryCards";
import { normalizeDashboardResponse } from "@/components/admin/dashboard/dashboardAdapter";
import { extractAdminProjectOptions } from "@/components/admin/projectOptions";
import SelectField, {
  type SelectOption,
} from "@/components/form/SelectField";
import MonthField from "@/components/form/MonthField";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import type {
  AdminDashboardData,
  AdminProjectOption,
} from "@/types/adminDashboard";

const emptyDashboard: AdminDashboardData = {
  summary: {
    project_count: 0,
    transport_count: 0,
    completed_transport_count: 0,
    volume_cubic: 0,
    weight_ton: 0,
    total_income: 0,
    total_expense: 0,
    net_profit: 0,
  },
  projects: [],
  daily: [],
};

function currentMonth() {
  const now = new Date();
  return `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, "0")}`;
}

export default function AdminDashboardPage() {
  const { getAPI } = useFetchAPI();
  const getAPIRef = useRef(getAPI);
  const [projectId, setProjectId] = useState<number | null>(null);
  const [month, setMonth] = useState(currentMonth);
  const [projects, setProjects] = useState<AdminProjectOption[]>([]);
  const [dashboard, setDashboard] =
    useState<AdminDashboardData>(emptyDashboard);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [reloadKey, setReloadKey] = useState(0);

  useEffect(() => {
    let active = true;
    void getAPIRef
      .current<unknown>(
        `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project?limit=999`,
        { authToken: "admin" },
      )
      .then((response) => {
        if (active && response.code >= 200 && response.code < 300) {
          setProjects(extractAdminProjectOptions(response.data));
        }
      });
    return () => {
      active = false;
    };
  }, []);

  useEffect(() => {
    let active = true;
    const params = new URLSearchParams({
      project_id: projectId === null ? "" : String(projectId),
      month,
    });
    queueMicrotask(() => {
      if (!active) return;
      setLoading(true);
      setError("");
    });

    void getAPIRef
      .current<unknown>(
        `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/dashboard?${params.toString()}`,
        { authToken: "admin", timeout: 20000 },
      )
      .then((response) => {
        if (!active) return;
        if (response.code < 200 || response.code >= 300) {
          setError(response.message || "Dashboard gagal dimuat.");
          setDashboard(emptyDashboard);
          return;
        }
        setDashboard(normalizeDashboardResponse(response.data));
      })
      .finally(() => {
        if (active) setLoading(false);
      });

    return () => {
      active = false;
    };
  }, [month, projectId, reloadKey]);

  const selectedProject = useMemo(
    () => projects.find((project) => project.project_id === projectId),
    [projectId, projects],
  );
  const projectOptions = useMemo<SelectOption[]>(
    () => [
      { value: "all", label: "Semua project" },
      ...projects.map((project) => ({
        value: project.project_id,
        label: project.project_code
          ? `${project.project_code} — ${project.project_name}`
          : project.project_name,
      })),
    ],
    [projects],
  );

  return (
    <div className="space-y-5 p-1 md:p-2">
      <section className="flex flex-col justify-between gap-4 rounded-2xl bg-gradient-to-r from-[#155eaa] to-[#0b9cad] p-5 text-white shadow-lg shadow-blue-900/10 sm:flex-row sm:items-end sm:p-7">
        <div>
          <p className="text-xs font-semibold uppercase tracking-[0.2em] text-cyan-100">
            Operational intelligence
          </p>
          <h1 className="mt-2 text-2xl font-semibold sm:text-3xl">
            {selectedProject?.project_name ?? "Semua Project"}
          </h1>
          <p className="mt-2 max-w-2xl text-sm text-blue-50/85">
            Ringkasan transport, kuantitas, dan performa keuangan berdasarkan
            project serta bulan operasional.
          </p>
        </div>
        <div className="rounded-xl bg-white/10 px-4 py-3 text-sm backdrop-blur">
          Periode: <strong>{month}</strong>
        </div>
      </section>

      <section className="flex flex-col gap-3 rounded-2xl border border-slate-200 bg-white p-4 shadow-sm sm:flex-row sm:items-end">
        <SelectField
          id="dashboard-project"
          label="Project"
          value={projectId ?? "all"}
          options={projectOptions}
          onChange={(value) =>
            setProjectId(value === "all" || value === null ? null : Number(value))
          }
          uncloseable
          wrapperClassName="min-w-0 flex-1"
          labelClassName="!mb-1.5 text-xs font-semibold uppercase tracking-wide !text-slate-500"
          className="h-11 w-full rounded-xl !border-slate-200 bg-white text-sm"
        />
        <MonthField
          id="dashboard-month"
          label="Bulan"
          value={month}
          onChange={(value) => {
            if (value) setMonth(value);
          }}
          required
          wrapperClassName="sm:w-56"
          labelClassName="!mb-1.5 text-xs font-semibold uppercase tracking-wide !text-slate-500"
        />
        <button
          type="button"
          onClick={() => setReloadKey((key) => key + 1)}
          className="inline-flex h-11 items-center justify-center gap-2 rounded-xl border border-slate-200 px-4 text-sm font-medium text-slate-600 hover:border-blue-200 hover:bg-blue-50 hover:text-blue-700"
        >
          <FaArrowRotateRight className={loading ? "animate-spin" : ""} />
          Muat ulang
        </button>
      </section>

      {error && (
        <div className="rounded-xl border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-700">
          {error}
        </div>
      )}

      <div className={loading ? "animate-pulse opacity-60" : ""}>
        <AdminSummaryCards summary={dashboard.summary} />
      </div>
      <div className={loading ? "animate-pulse opacity-60" : ""}>
        <AdminDashboardChart
          projectId={projectId}
          projects={dashboard.projects}
          daily={dashboard.daily}
        />
      </div>
    </div>
  );
}
