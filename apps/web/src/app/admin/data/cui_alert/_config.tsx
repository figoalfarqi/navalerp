/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "cui_alert";
export const entityTitle = "Peringatan CUI";
export const entityEndpoint = "/admin/cui_alert";
export const primaryKey = "alert_id";

export const columns: ColumnField[] = [
  { key: "alert_code", label: "Kode Peringatan" },
  {
    key: "asset_name",
    label: "Aset Bawah Laut (CUI)",
    render: (item: any) =>
      item.asset_name ? (
        <div>
          <span className="font-semibold text-slate-200">{item.asset_name}</span>
          {item.asset_code && (
            <span className="block text-xs font-mono text-cyan-400/80">{item.asset_code}</span>
          )}
        </div>
      ) : (
        <span className="text-slate-500 italic">-</span>
      ),
  },
  { key: "alert_type", label: "Jenis Ancaman" },
  {
    key: "severity",
    label: "Tingkat Keparahan",
    render: (item: any) => {
      const sev = String(item.severity).toUpperCase();
      let colorClass = "bg-blue-500/20 text-blue-400 border-blue-500/30";
      if (sev === "CRITICAL") colorClass = "bg-rose-500/20 text-rose-400 border-rose-500/30 font-bold";
      else if (sev === "HIGH") colorClass = "bg-amber-500/20 text-amber-400 border-amber-500/30 font-semibold";
      else if (sev === "MEDIUM") colorClass = "bg-yellow-500/20 text-yellow-400 border-yellow-500/30";
      return (
        <span className={`px-2 py-0.5 rounded text-xs border ${colorClass}`}>
          {sev}
        </span>
      );
    },
  },
  {
    key: "status",
    label: "Status",
    render: (item: any) => {
      const st = String(item.status).toUpperCase();
      let colorClass = "bg-slate-500/20 text-slate-300 border-slate-500/30";
      if (st === "ACTIVE") colorClass = "bg-rose-500/20 text-rose-400 border-rose-500/30 animate-pulse";
      else if (st === "INVESTIGATING") colorClass = "bg-amber-500/20 text-amber-400 border-amber-500/30";
      else if (st === "DISPATCHED") colorClass = "bg-cyan-500/20 text-cyan-400 border-cyan-500/30";
      else if (st === "RESOLVED") colorClass = "bg-emerald-500/20 text-emerald-400 border-emerald-500/30";
      return (
        <span className={`px-2 py-0.5 rounded text-xs border ${colorClass}`}>
          {st}
        </span>
      );
    },
  },
  {
    key: "ship_name",
    label: "Kapal Patroli Ditugaskan (KRI)",
    render: (item: any) =>
      item.ship_name ? (
        <div className="flex items-center gap-1.5">
          {item.hull_number && (
            <span className="px-1.5 py-0.5 rounded text-[11px] font-mono bg-blue-500/20 text-blue-300 border border-blue-500/30">
              {item.hull_number}
            </span>
          )}
          <span className="text-slate-200">{item.ship_name}</span>
        </div>
      ) : (
        <span className="text-slate-500 italic">Belum Ditugaskan</span>
      ),
  },
  { key: "detected_at", label: "Waktu Terdeteksi", render: (item: any) => formatSmartDate(item.detected_at) },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Peringatan CUI", fieldType: "text", col: "left", placeHolder: "Ketik kode peringatan..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "alert_code",
    col: "left",
    label: "Kode Peringatan",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "ALT-CUI",
  },
  {
    name: "cui_asset_id",
    col: "right",
    label: "Aset Bawah Laut (CUI)",
    fieldType: "select",
    options: {
      url: "/admin/cui_asset?limit=100",
      labelKey: "{asset_code} - {asset_name}",
      valueKey: "cui_asset_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "alert_type",
    col: "left",
    label: "Jenis Ancaman / Anomali",
    fieldType: "select",
    options: [
      { label: "VESSEL_ANCHOR_DRAG_RISK - Risiko Hantaman Jangkar Kapal", value: "VESSEL_ANCHOR_DRAG_RISK" },
      { label: "SEISMIC_DISTURBANCE - Gangguan Seismik Bawah Laut", value: "SEISMIC_DISTURBANCE" },
      { label: "PRESSURE_DROP - Penurunan Tekanan Aliran Pipa", value: "PRESSURE_DROP" },
      { label: "ACOUSTIC_ANOMALY - Anomali Sonar / Suara Bawah Air", value: "ACOUSTIC_ANOMALY" },
      { label: "UNAUTHORIZED_SUBMERSIBLE - Deteksi Wahana Bawah Air Asing", value: "UNAUTHORIZED_SUBMERSIBLE" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "severity",
    col: "right",
    label: "Tingkat Keparahan Ancaman",
    fieldType: "select",
    options: [
      { label: "CRITICAL - Bahaya Segera Terhadap Objek Vital", value: "CRITICAL" },
      { label: "HIGH - Risiko Tinggi Butuh Intervensi Cepat", value: "HIGH" },
      { label: "MEDIUM - Pantauan Intensif Diperlukan", value: "MEDIUM" },
      { label: "LOW - Tingkat Rendah / Informatif", value: "LOW" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "left",
    label: "Status Penanganan",
    fieldType: "select",
    options: [
      { label: "ACTIVE - Masih Aktif / Belum Ditangani", value: "ACTIVE" },
      { label: "INVESTIGATING - Sedang Tahap Verifikasi Intelijen", value: "INVESTIGATING" },
      { label: "DISPATCHED - KRI / Unsur Patroli Diluncurkan ke Lokasi", value: "DISPATCHED" },
      { label: "RESOLVED - Situasi Telah Dinormalisasi", value: "RESOLVED" },
      { label: "FALSE_ALARM - Terkonfirmasi Alarm Palsu", value: "FALSE_ALARM" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "assigned_ship_id",
    col: "right",
    label: "Kapal Patroli Ditugaskan (KRI)",
    fieldType: "select",
    options: {
      url: "/admin/ship?limit=100",
      labelKey: "{hull_number} - {ship_name}",
      valueKey: "ship_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "detected_at",
    col: "left",
    label: "Waktu Terdeteksi Oleh Sistem",
    fieldType: "datetime",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "ai_confidence",
    col: "right",
    label: "Tingkat Akurasi AI (0.00 - 1.00)",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "recommended_action",
    col: "left",
    label: "Rekomendasi Tindakan Taktis",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "resolution_notes",
    col: "right",
    label: "Catatan Tindak Lanjut / Resolusi",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "resolved_at",
    col: "left",
    label: "Waktu Situasi Terselesaikan",
    fieldType: "datetime",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.alert_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  if (!payload.assigned_ship_id) payload.assigned_ship_id = null;
  if (!payload.resolved_at) payload.resolved_at = null;
  return payload;
};
