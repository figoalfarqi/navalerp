/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "cui_asset";
export const entityTitle = "Aset Bawah Laut (CUI)";
export const entityEndpoint = "/admin/cui_asset";
export const primaryKey = "cui_asset_id";

export const columns: ColumnField[] = [
  { key: "asset_code", label: "Kode Aset" },
  { key: "asset_name", label: "Nama Aset CUI" },
  { key: "asset_type", label: "Tipe Aset" },
  { key: "operator_name", label: "Operator / Pemilik" },
  {
    key: "theater_name",
    label: "Teater Pertahanan",
    render: (item: any) =>
      item.theater_name ? (
        <span className="px-2 py-0.5 rounded text-xs bg-indigo-500/20 text-indigo-300 border border-indigo-500/30 font-medium">
          {item.theater_name}
        </span>
      ) : (
        <span className="text-slate-500 italic">-</span>
      ),
  },
  {
    key: "status",
    label: "Status",
    render: (item: any) => {
      const st = String(item.status).toUpperCase();
      let colorClass = "bg-slate-500/20 text-slate-300 border-slate-500/30";
      if (st === "ACTIVE_MONITORED") colorClass = "bg-emerald-500/20 text-emerald-400 border-emerald-500/30";
      else if (st === "INSPECTION_REQUIRED") colorClass = "bg-amber-500/20 text-amber-400 border-amber-500/30";
      else if (st === "UNDER_MAINTENANCE") colorClass = "bg-blue-500/20 text-blue-400 border-blue-500/30";
      else if (st === "ALERT_ANOMALY") colorClass = "bg-rose-500/20 text-rose-400 border-rose-500/30 font-bold animate-pulse";
      return (
        <span className={`px-2 py-0.5 rounded text-xs border ${colorClass}`}>
          {st}
        </span>
      );
    },
  },
  {
    key: "protection_priority",
    label: "Prioritas Perlindungan",
    render: (item: any) => {
      const prio = String(item.protection_priority).toUpperCase();
      let colorClass = "bg-slate-500/20 text-slate-300 border-slate-500/30";
      if (prio === "CRITICAL_TIER_1") colorClass = "bg-rose-500/20 text-rose-400 border-rose-500/30 font-bold";
      else if (prio === "HIGH_TIER_2") colorClass = "bg-amber-500/20 text-amber-400 border-amber-500/30 font-semibold";
      else if (prio === "MEDIUM_TIER_3") colorClass = "bg-sky-500/20 text-sky-400 border-sky-500/30";
      return (
        <span className={`px-2 py-0.5 rounded text-xs border ${colorClass}`}>
          {prio}
        </span>
      );
    },
  },
  {
    key: "health_score",
    label: "Kesehatan",
    render: (item: any) => {
      const hs = Number(item.health_score || 0);
      let color = "text-emerald-400";
      if (hs < 60) color = "text-rose-400 font-bold";
      else if (hs < 80) color = "text-amber-400";
      return <span className={`font-mono ${color}`}>{hs}%</span>;
    },
  },
  { key: "depth_meters", label: "Kedalaman (m)", render: (item: any) => `${item.depth_meters ?? 0} m` },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Aset Bawah Laut", fieldType: "text", col: "left", placeHolder: "Ketik kode, nama, atau operator..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "asset_code",
    col: "left",
    label: "Kode Aset CUI",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "asset_name",
    col: "right",
    label: "Nama Aset Bawah Laut",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "asset_type",
    col: "left",
    label: "Tipe Aset Bawah Laut",
    fieldType: "select",
    options: [
      { label: "SUBMARINE_CABLE - Kabel Komunikasi / Listrik Bawah Laut", value: "SUBMARINE_CABLE" },
      { label: "SUBSEA_PIPELINE - Jalur Pipa Migas Subsea", value: "SUBSEA_PIPELINE" },
      { label: "LANDING_STATION - Stasiun Pendaratan Kabel Laut", value: "LANDING_STATION" },
      { label: "OFFSHORE_ENERGY - Anjungan & Turbin Energi Lepas Pantai", value: "OFFSHORE_ENERGY" },
      { label: "MONITORING_SYSTEM - Array Sensor Hidrofon & Sonar Tetap", value: "MONITORING_SYSTEM" },
      { label: "OTHER - Objek Infrastruktur Bawah Laut Lainnya", value: "OTHER" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "operator_name",
    col: "right",
    label: "Nama Operator / Pemilik",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "theater_id",
    col: "left",
    label: "Teater Wilayah Pertahanan",
    fieldType: "select",
    options: {
      url: "/admin/theater?limit=100",
      labelKey: "{theater_code} - {theater_name}",
      valueKey: "theater_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "right",
    label: "Status Operasional",
    fieldType: "select",
    options: [
      { label: "ACTIVE_MONITORED - Aktif dan Dipantau Sistem", value: "ACTIVE_MONITORED" },
      { label: "INSPECTION_REQUIRED - Membutuhkan Inspeksi Lapangan Segera", value: "INSPECTION_REQUIRED" },
      { label: "UNDER_MAINTENANCE - Dalam Proses Pemeliharaan / Perbaikan", value: "UNDER_MAINTENANCE" },
      { label: "ALERT_ANOMALY - Anomali Aktif Terdeteksi", value: "ALERT_ANOMALY" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "protection_priority",
    col: "left",
    label: "Tingkat Prioritas Pengamanan",
    fieldType: "select",
    options: [
      { label: "CRITICAL_TIER_1 - Objek Vital Nasional Strategis Utama", value: "CRITICAL_TIER_1" },
      { label: "HIGH_TIER_2 - Infrastruktur Prioritas Tinggi", value: "HIGH_TIER_2" },
      { label: "MEDIUM_TIER_3 - Infrastruktur Prioritas Menengah", value: "MEDIUM_TIER_3" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "health_score",
    col: "right",
    label: "Skor Kesehatan Fisik (0 - 100)",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "depth_meters",
    col: "left",
    label: "Kedalaman Rata-rata (Meter)",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "length_km",
    col: "right",
    label: "Panjang Jalur Terbentang (KM)",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "latitude",
    col: "left",
    label: "Koordinat Latitude (Lintang)",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "longitude",
    col: "right",
    label: "Koordinat Longitude (Bujur)",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "start_coordinates",
    col: "left",
    label: "Titik Koordinat Awal (Format Lat, Lon)",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "end_coordinates",
    col: "right",
    label: "Titik Koordinat Akhir (Format Lat, Lon)",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "last_inspected_at",
    col: "left",
    label: "Waktu Inspeksi Terakhir",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "next_inspection_due",
    col: "right",
    label: "Tenggat Inspeksi Berikutnya",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "notes",
    col: "left",
    label: "Catatan Tambahan & Spesifikasi Aset",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.cui_asset_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  if (!payload.theater_id) payload.theater_id = null;
  if (!payload.last_inspected_at) payload.last_inspected_at = null;
  if (!payload.next_inspection_due) payload.next_inspection_due = null;
  return payload;
};
