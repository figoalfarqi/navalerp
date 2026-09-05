"use client";

import { useEffect, useMemo, useRef, useState } from "react";
import jsPDF from "jspdf";
import autoTable from "jspdf-autotable";
import {
  FaArrowRotateRight,
  FaFilePdf,
  FaPrint,
} from "react-icons/fa6";
import AdminSummaryCards from "@/components/admin/dashboard/AdminSummaryCards";
import { normalizeReportRows } from "@/components/admin/dashboard/dashboardAdapter";
import { extractAdminProjectOptions } from "@/components/admin/projectOptions";
import SelectField, {
  type SelectOption,
} from "@/components/form/SelectField";
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

  const exportPdf = () => {
    const document = new jsPDF({
      orientation: "landscape",
      unit: "mm",
      format: "a4",
    });
    const projectLabel = selectedProject
      ? `${selectedProject.project_code ? `${selectedProject.project_code} - ` : ""}${selectedProject.project_name}`
      : "Semua Project";
    document.setFontSize(16);
    document.text("Laporan Operasional & Keuangan NavalERP", 14, 15);
    document.setFontSize(9);
    document.setTextColor(90);
    document.text(
      `${projectLabel} | ${period === "weekly" ? "Mingguan" : "Bulanan"} | Acuan ${date}`,
      14,
      21,
    );
    document.text(
      `Transport ${formatNumberID(summary.transport_count)} | Volume ${formatNumberID(summary.volume_cubic)} m3 | Berat ${formatNumberID(summary.weight_ton)} ton | Laba ${formatCurrencyIDR(summary.net_profit)}`,
      14,
      27,
    );

    autoTable(document, {
      startY: 33,
      head: [
        [
          "Project",
          "Periode",
          "Transport",
          "Selesai",
          "Volume (m3)",
          "Berat (ton)",
          "Pendapatan",
          "Pengeluaran",
          "Laba Bersih",
        ],
      ],
      body: rows.map((row) => [
        `${row.project_code ? `${row.project_code} - ` : ""}${row.project_name}`,
        row.period_label,
        formatNumberID(row.transport_count),
        formatNumberID(row.completed_transport_count),
        formatNumberID(row.volume_cubic),
        formatNumberID(row.weight_ton),
        formatCurrencyIDR(row.total_income),
        formatCurrencyIDR(row.total_expense),
        formatCurrencyIDR(row.net_profit),
      ]),
      styles: { fontSize: 7, cellPadding: 2 },
      headStyles: { fillColor: [21, 94, 170] },
      alternateRowStyles: { fillColor: [245, 248, 252] },
    });

    document.save(`laporan-${period}-${date}.pdf`);
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
          <div className="flex flex-wrap gap-2 print:hidden">
            <button
              type="button"
              onClick={() => window.print()}
              className="inline-flex h-10 items-center gap-2 rounded-xl border border-slate-200 px-4 text-sm font-medium text-slate-700 hover:bg-slate-50"
            >
              <FaPrint /> Cetak
            </button>
            <button
              type="button"
              onClick={exportPdf}
              disabled={rows.length === 0}
              className="inline-flex h-10 items-center gap-2 rounded-xl bg-red-600 px-4 text-sm font-medium text-white hover:bg-red-700 disabled:cursor-not-allowed disabled:opacity-50"
            >
              <FaFilePdf /> Export PDF
            </button>
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
          labelClassName="!mb-1.5 text-xs font-semibold uppercase tracking-wide !text-slate-500"
          className="h-11 w-full rounded-xl !border-slate-200 bg-white text-sm"
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
          labelClassName="!mb-1.5 text-xs font-semibold uppercase tracking-wide !text-slate-500"
          className="h-11 w-full rounded-xl !border-slate-200 bg-white text-sm"
        />
        <label className="flex flex-col gap-1.5">
          <span className="text-xs font-semibold uppercase tracking-wide text-slate-500">
            Tanggal acuan
          </span>
          <input
            type="date"
            value={date}
            onChange={(event) => setDate(event.target.value)}
            className="h-11 rounded-xl border border-slate-200 px-3 text-sm outline-none focus:border-blue-500 focus:ring-4 focus:ring-blue-100"
          />
        </label>
        <button
          type="button"
          onClick={() => setReloadKey((key) => key + 1)}
          className="inline-flex h-11 items-center justify-center gap-2 rounded-xl bg-blue-600 px-4 text-sm font-medium text-white hover:bg-blue-700"
        >
          <FaArrowRotateRight className={loading ? "animate-spin" : ""} />
          Terapkan
        </button>
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
            tanggal acuan {date}
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
