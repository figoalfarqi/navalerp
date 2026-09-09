/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "cui_monitoring_log";
export const entityTitle = "Log Sensor CUI";
export const entityEndpoint = "/admin/cui_monitoring_log";
export const primaryKey = "log_id";

export const columns: ColumnField[] = [
  { key: "sensor_code", label: "Kode Sensor" },
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
  { key: "sensor_type", label: "Tipe Sensor" },
  {
    key: "metric_value",
    label: "Nilai Ukur",
    render: (item: any) => `${item.metric_value} ${item.metric_unit || ""}`,
  },
  {
    key: "status",
    label: "Status",
    render: (item: any) => {
      const st = String(item.status).toUpperCase();
      let colorClass = "bg-slate-500/20 text-slate-300 border-slate-500/30";
      if (st === "NORMAL") colorClass = "bg-emerald-500/20 text-emerald-400 border-emerald-500/30";
      else if (st === "WARNING") colorClass = "bg-amber-500/20 text-amber-400 border-amber-500/30";
      else if (st === "CRITICAL_ANOMALY") colorClass = "bg-rose-500/20 text-rose-400 border-rose-500/30 font-bold animate-pulse";
      return (
        <span className={`px-2 py-0.5 rounded text-xs border ${colorClass}`}>
          {st}
        </span>
      );
    },
  },
  { key: "vessel_proximity_mmsi", label: "MMSI Kapal Terdekat" },
  { key: "log_time", label: "Waktu Sensor", render: (item: any) => formatSmartDate(item.log_time) },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Log Sensor CUI", fieldType: "text", col: "left", placeHolder: "Ketik kode sensor atau deskripsi..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "cui_asset_id",
    col: "left",
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
    name: "sensor_code",
    col: "right",
    label: "Kode Sensor",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "sensor_type",
    col: "left",
    label: "Tipe Sensor Pemantau",
    fieldType: "select",
    options: [
      { label: "ACOUSTIC_SONAR - Sonar Hidrofon Akustik Bawah Air", value: "ACOUSTIC_SONAR" },
      { label: "PRESSURE_TRANSDUCER - Sensor Transduser Tekanan Subsea", value: "PRESSURE_TRANSDUCER" },
      { label: "FIBER_STRAIN - Sensor Regangan Serat Optik (CUI Strain)", value: "FIBER_STRAIN" },
      { label: "SEISMOMETER - Seismometer Dasar Laut OBS", value: "SEISMOMETER" },
      { label: "MAGNETIC_ANOMALY - Sensor Anomali Medan Magnetik (MAD)", value: "MAGNETIC_ANOMALY" },
      { label: "TEMPERATURE_SENSOR - Sensor Termal Kedalaman Laut", value: "TEMPERATURE_SENSOR" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "log_time",
    col: "right",
    label: "Waktu Pencatatan Log",
    fieldType: "datetime",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "metric_value",
    col: "left",
    label: "Nilai Metrik Terukur",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "metric_unit",
    col: "right",
    label: "Satuan Metrik",
    fieldType: "select",
    options: [
      { label: "dB (Decibel)", value: "dB" },
      { label: "bar (Tekanan Bar)", value: "bar" },
      { label: "psi (Tekanan PSI)", value: "psi" },
      { label: "µε (Microstrain)", value: "µε" },
      { label: "m/s (Kecepatan Aliran)", value: "m/s" },
      { label: "nT (NanoTesla - Fluks Magnetik)", value: "nT" },
      { label: "°C (Derajat Celcius)", value: "°C" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "left",
    label: "Status Pembacaan Sensor",
    fieldType: "select",
    options: [
      { label: "NORMAL - Nilai Normal & Aman", value: "NORMAL" },
      { label: "WARNING - Melebihi Batas Ambang Peringatan", value: "WARNING" },
      { label: "CRITICAL_ANOMALY - Anomali Kritis / Gangguan Fisik", value: "CRITICAL_ANOMALY" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "vessel_proximity_mmsi",
    col: "right",
    label: "MMSI Kapal Terdekat (Jika Ada)",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "anomaly_score",
    col: "left",
    label: "Skor Anomali AI (0.00 - 1.00)",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "description",
    col: "right",
    label: "Deskripsi Telemetri / Catatan Sensor",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.log_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
