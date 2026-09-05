"use client";

import {
  Bar,
  BarChart,
  CartesianGrid,
  Legend,
  Line,
  LineChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";
import type {
  AdminDailyMetric,
  AdminProjectComparison,
} from "@/types/adminDashboard";
import { formatNumberID } from "@/utils/currencyFormater";

interface AdminDashboardChartProps {
  projectId: number | null;
  projects: AdminProjectComparison[];
  daily: AdminDailyMetric[];
}

const tooltipStyle = {
  borderRadius: 12,
  borderColor: "#e2e8f0",
  boxShadow: "0 10px 30px rgba(15,23,42,.08)",
};

export default function AdminDashboardChart({
  projectId,
  projects,
  daily,
}: AdminDashboardChartProps) {
  const isAllProjects = projectId === null;
  const dataEmpty = isAllProjects ? projects.length === 0 : daily.length === 0;

  return (
    <section className="rounded-2xl border border-slate-200 bg-white p-4 shadow-sm sm:p-6">
      <div className="mb-5">
        <p className="text-xs font-semibold uppercase tracking-[0.18em] text-blue-600">
          {isAllProjects ? "Perbandingan project" : "Pergerakan harian"}
        </p>
        <h2 className="mt-1 text-lg font-semibold text-slate-900">
          {isAllProjects
            ? "Transport selesai per project"
            : "Volume dan berat harian"}
        </h2>
        <p className="mt-1 text-sm text-slate-500">
          {isAllProjects
            ? "Tampilan semua project membandingkan jumlah perjalanan bulan terpilih."
            : "Tampilan satu project menunjukkan realisasi setiap hari dalam bulan terpilih."}
        </p>
      </div>

      {dataEmpty ? (
        <div className="grid h-80 place-items-center rounded-xl border border-dashed border-slate-200 bg-slate-50 text-sm text-slate-500">
          Belum ada data untuk filter ini.
        </div>
      ) : (
        <div className="h-80 w-full">
          <ResponsiveContainer width="100%" height="100%">
            {isAllProjects ? (
              <BarChart
                data={projects}
                margin={{ top: 8, right: 8, left: 0, bottom: 20 }}
              >
                <CartesianGrid strokeDasharray="4 4" stroke="#e2e8f0" />
                <XAxis
                  dataKey="project_name"
                  tick={{ fontSize: 11, fill: "#64748b" }}
                  angle={-15}
                  textAnchor="end"
                  height={58}
                />
                <YAxis
                  allowDecimals={false}
                  tick={{ fontSize: 11, fill: "#64748b" }}
                />
                <Tooltip contentStyle={tooltipStyle} />
                <Legend />
                <Bar
                  dataKey="transport_count"
                  name="Total transport"
                  fill="#38bdf8"
                  radius={[6, 6, 0, 0]}
                />
                <Bar
                  dataKey="completed_transport_count"
                  name="Selesai"
                  fill="#155eaa"
                  radius={[6, 6, 0, 0]}
                />
              </BarChart>
            ) : (
              <LineChart
                data={daily}
                margin={{ top: 8, right: 14, left: 2, bottom: 8 }}
              >
                <CartesianGrid strokeDasharray="4 4" stroke="#e2e8f0" />
                <XAxis
                  dataKey="label"
                  tick={{ fontSize: 11, fill: "#64748b" }}
                />
                <YAxis
                  tickFormatter={(value) => formatNumberID(Number(value))}
                  tick={{ fontSize: 11, fill: "#64748b" }}
                />
                <Tooltip
                  contentStyle={tooltipStyle}
                  formatter={(value) => formatNumberID(Number(value ?? 0))}
                />
                <Legend />
                <Line
                  type="monotone"
                  dataKey="volume_cubic"
                  name="Volume (m³)"
                  stroke="#06b6d4"
                  strokeWidth={3}
                  dot={{ r: 3 }}
                />
                <Line
                  type="monotone"
                  dataKey="weight_ton"
                  name="Berat (ton)"
                  stroke="#155eaa"
                  strokeWidth={3}
                  dot={{ r: 3 }}
                />
              </LineChart>
            )}
          </ResponsiveContainer>
        </div>
      )}
    </section>
  );
}
