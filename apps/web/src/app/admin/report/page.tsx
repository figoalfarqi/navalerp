"use client";

import { useEffect, useMemo, useRef, useState } from "react";
import ReportPdfDocument from "@/components/admin/report/ReportPdfDocument";
import {
  FaArrowRotateRight,
  FaFilePdf,
  FaPrint,
} from "@/components/icons";
import AdminSummaryCards from "@/components/admin/dashboard/AdminSummaryCards";
import { normalizeReportRows } from "@/components/admin/dashboard/dashboardAdapter";
import { extractAdminProjectOptions } from "@/components/admin/projectOptions";
import SelectField, {
  type SelectOption,
} from "@/components/form/SelectField";
import Button from "@/components/form/Button";
import DateField from "@/components/form/DateField";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import type {
  AdminDashboardSummary,
  AdminProjectOption,
  AdminReportRow,
} from "@/types/adminDashboard";
import {
  formatCurrencyIDR,
  formatNumberID,
} from "@/utils/currencyFormater";
import { formatDate } from "@/utils/dateTime";

const emptySummary: AdminDashboardSummary = {
  project_count: 0,
  transport_count: 0,
  completed_transport_count: 0,
  volume_cubic: 0,
  weight_ton: 0,
  total_income: 0,
  total_expense: 0,
  net_profit: 0,
};

const periodOptions: SelectOption[] = [
  { value: "monthly", label: "Bulanan" },
  { value: "weekly", label: "Mingguan" },
];

function today() {
  const current = new Date();
  return `${current.getFullYear()}-${String(current.getMonth() + 1).padStart(
    2,
    "0",
  )}-${String(current.getDate()).padStart(2, "0")}`;
}

export default function AdminReportPage() {
  const { getAPI } = useFetchAPI();
  const getAPIRef = useRef(getAPI);
  const [projects, setProjects] = useState<AdminProjectOption[]>([]);
  const [projectId, setProjectId] = useState<number | null>(null);
  const [period, setPeriod] = useState<"weekly" | "monthly">("monthly");
  const [date, setDate] = useState(today);
  const [summary, setSummary] = useState(emptySummary);
  const [rows, setRows] = useState<AdminReportRow[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [reloadKey, setReloadKey] = useState(0);
  const [exportingPdf, setExportingPdf] = useState(false);

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
      period,
      date,
    });
    queueMicrotask(() => {
      if (!active) return;
      setLoading(true);
      setError("");
    });

    void getAPIRef
      .current<unknown>(
        `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/report?${params.toString()}`,
        { authToken: "admin", timeout: 30000 },
      )
      .then((response) => {
        if (!active) return;
        if (response.code < 200 || response.code >= 300) {
          setError(response.message || "Laporan gagal dimuat.");
          setSummary(emptySummary);
          setRows([]);
          return;
        }
        const report = normalizeReportRows(response.data);
        setSummary(report.summary);
        setRows(report.rows);
      })
      .finally(() => {
        if (active) setLoading(false);
      });

    return () => {
      active = false;
    };
  }, [date, period, projectId, reloadKey]);

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

  const exportPdf = async () => {
    if (rows.length === 0 || exportingPdf) return;
    try {
      setExportingPdf(true);
      const projectLabel = selectedProject
        ? `${selectedProject.project_code ? `${selectedProject.project_code} - ` : ""}${selectedProject.project_name}`
        : "Semua Project";
      const periodLabel = period === "weekly" ? "Mingguan" : "Bulanan";

      const { pdf } = await import("@react-pdf/renderer");
      const blob = await pdf(
        <ReportPdfDocument
          projectLabel={projectLabel}
          periodLabel={periodLabel}
          date={date}
          summary={summary}
          rows={rows}
        />,
      ).toBlob();

      const url = URL.createObjectURL(blob);
      const link = document.createElement("a");
      link.href = url;
      link.download = `laporan-${period}-${date}.pdf`;
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
      URL.revokeObjectURL(url);
    } catch (err) {
      console.error("Gagal mengekspor PDF:", err);
    } finally {
      setExportingPdf(false);
    }
  };

  return (
    <div className="space-y-5 p-1 md:p-2">
      <section className="print:border-0 print:shadow-none rounded-2xl border border-slate-200 bg-white p-5 shadow-sm sm:p-7">
        <div className="flex flex-col justify-between gap-4 sm:flex-row sm:items-start">
          <div>
            <p className="text-xs font-semibold uppercase tracking-[0.18em] text-blue-600">
              Reporting center
            </p>
            <h1 className="mt-1 text-2xl font-semibold text-slate-900">
              Laporan Operasional & Keuangan
            </h1>
            <p className="mt-2 max-w-2xl text-sm text-slate-500">
              Gunakan filter project dan periode untuk laporan mingguan atau
              bulanan, lalu cetak langsung atau simpan sebagai PDF.
            </p>
          </div>
          <div className="flex flex-wrap items-center gap-2 print:hidden">
            <Button
              id="report-print-btn"
              type="button"
              variant="gray-outline"
              size="md"
              onClick={() => window.print()}
            >
              <FaPrint /> Cetak
            </Button>
            <Button
              id="report-export-pdf-btn"
              type="button"
              variant="red-solid"
              size="md"
              onClick={exportPdf}
              disabled={rows.length === 0 || exportingPdf}
            >
              <FaFilePdf /> {exportingPdf ? "Mengekspor..." : "Export PDF"}
            </Button>
          </div>
        </div>
      </section>

      <section className="grid gap-3 rounded-2xl border border-slate-200 bg-white p-4 shadow-sm print:hidden sm:grid-cols-2 xl:grid-cols-[minmax(260px,1fr)_180px_190px_auto] xl:items-end">
        <SelectField
          id="report-project"
          label="Project"
          value={projectId ?? "all"}
          options={projectOptions}
          onChange={(value) =>
            setProjectId(value === "all" || value === null ? null : Number(value))
          }
          uncloseable
        />
        <SelectField
          id="report-period"
          label="Periode"
          value={period}
          options={periodOptions}
          onChange={(value) =>
            setPeriod(value === "weekly" ? "weekly" : "monthly")
          }
          uncloseable
        />
        <DateField
          id="report-date"
          label="Tanggal acuan"
          value={date}
          uncloseable
          onChange={(val) => {
            if (!val) return;
            const d = new Date(val);
            if (!isNaN(d.getTime())) {
              const yyyy = d.getFullYear();
              const mm = String(d.getMonth() + 1).padStart(2, "0");
              const dd = String(d.getDate()).padStart(2, "0");
              setDate(`${yyyy}-${mm}-${dd}`);
            }
          }}
        />
        <div className="flex items-end pb-0.5">
          <Button
            id="report-apply-btn"
            type="button"
            variant="blue-solid"
            size="md"
            onClick={() => setReloadKey((key) => key + 1)}
          >
            <FaArrowRotateRight className={loading ? "animate-spin" : ""} />
            Terapkan
          </Button>
        </div>
      </section>

      {error && (
        <div className="rounded-xl border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-700">
          {error}
        </div>
      )}

      <div className={loading ? "animate-pulse opacity-60" : ""}>
        <AdminSummaryCards summary={summary} />
      </div>

      <section className="overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm print:border-slate-300 print:shadow-none">
        <div className="border-b border-slate-200 px-5 py-4">
          <h2 className="font-semibold text-slate-900">
            {selectedProject?.project_name ?? "Semua Project"}
          </h2>
          <p className="text-xs text-slate-500">
            {period === "weekly" ? "Laporan mingguan" : "Laporan bulanan"} —
            tanggal acuan {formatDate(date)}
          </p>
        </div>
        <div className="overflow-x-auto">
          <table className="w-full min-w-[1050px] border-collapse text-sm">
            <thead className="bg-slate-50 text-left text-xs uppercase tracking-wide text-slate-500">
              <tr>
                {[
                  "Project",
                  "Periode",
                  "Transport",
                  "Selesai",
                  "Volume",
                  "Berat",
                  "Pendapatan",
                  "Pengeluaran",
                  "Laba Bersih",
                ].map((heading) => (
                  <th key={heading} className="px-4 py-3 font-semibold">
                    {heading}
                  </th>
                ))}
              </tr>
            </thead>
            <tbody className="divide-y divide-slate-100">
              {rows.map((row, index) => (
                <tr
                  key={`${row.project_id ?? row.project_name}-${row.period_label}-${index}`}
                  className="hover:bg-slate-50"
                >
                  <td className="px-4 py-3 font-medium text-slate-800">
                    {row.project_code && (
                      <span className="mr-1 text-xs text-slate-400">
                        {row.project_code}
                      </span>
                    )}
                    {row.project_name}
                  </td>
                  <td className="px-4 py-3 text-slate-600">
                    {row.period_label}
                  </td>
                  <td className="px-4 py-3">
                    {formatNumberID(row.transport_count)}
                  </td>
                  <td className="px-4 py-3">
                    {formatNumberID(row.completed_transport_count)}
                  </td>
                  <td className="px-4 py-3">
                    {formatNumberID(row.volume_cubic)} m³
                  </td>
                  <td className="px-4 py-3">
                    {formatNumberID(row.weight_ton)} ton
                  </td>
                  <td className="px-4 py-3 text-emerald-700">
                    {formatCurrencyIDR(row.total_income)}
                  </td>
                  <td className="px-4 py-3 text-red-700">
                    {formatCurrencyIDR(row.total_expense)}
                  </td>
                  <td
                    className={`px-4 py-3 font-semibold ${
                      row.net_profit < 0 ? "text-red-700" : "text-emerald-700"
                    }`}
                  >
                    {formatCurrencyIDR(row.net_profit)}
                  </td>
                </tr>
              ))}
              {!loading && rows.length === 0 && (
                <tr>
                  <td
                    colSpan={9}
                    className="px-4 py-16 text-center text-slate-500"
                  >
                    Belum ada data laporan untuk filter ini.
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </section>
    </div>
  );
}
