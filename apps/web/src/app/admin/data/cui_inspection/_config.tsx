/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "cui_inspection";
export const entityTitle = "Inspeksi CUI";
export const entityEndpoint = "/admin/cui_inspection";
export const primaryKey = "inspection_id";

export const columns: ColumnField[] = [
  { key: "inspection_number", label: "Nomor Inspeksi" },
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
  {
    key: "ship_name",
    label: "Kapal Penginspeksi",
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
        <span className="text-slate-500 italic">-</span>
      ),
  },
  {
    key: "inspector_name",
    label: "Perwira Penginspeksi",
    render: (item: any) =>
      item.inspector_name ? (
        <div>
          <span className="text-slate-200 font-medium">{item.inspector_name}</span>
          {item.inspector_nrp && (
            <span className="block text-xs font-mono text-slate-400">{item.inspector_nrp}</span>
          )}
        </div>
      ) : (
        <span className="text-slate-500 italic">-</span>
      ),
  },
  { key: "inspection_date", label: "Tanggal Inspeksi", render: (item: any) => formatSmartDate(item.inspection_date) },
  { key: "method", label: "Metode" },
  { key: "condition_rating", label: "Kondisi Fisik" },
  {
    key: "remedial_action_required",
    label: "Tindakan Perbaikan",
    render: (item: any) =>
      item.remedial_action_required ? (
        <span className="px-2 py-0.5 rounded text-xs font-semibold bg-rose-500/20 text-rose-400 border border-rose-500/30">
          Perlu Perbaikan
        </span>
      ) : (
        <span className="px-2 py-0.5 rounded text-xs font-semibold bg-emerald-500/20 text-emerald-400 border border-emerald-500/30">
          Aman
        </span>
      ),
  },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Inspeksi CUI", fieldType: "text", col: "left", placeHolder: "Ketik nomor inspeksi atau temuan..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "inspection_number",
    col: "left",
    label: "Nomor Inspeksi",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "INSP-CUI",
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
    name: "ship_id",
    col: "left",
    label: "Kapal Penginspeksi (KRI)",
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
    name: "inspector_officer_id",
    col: "right",
    label: "Perwira Penginspeksi",
    fieldType: "select",
    options: {
      url: "/admin/personnel?limit=100",
      labelKey: "{nrp} - {full_name}",
      valueKey: "personnel_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "inspection_date",
    col: "left",
    label: "Tanggal Pelaksanaan Inspeksi",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "next_inspection_date",
    col: "right",
    label: "Tenggat Inspeksi Berikutnya",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "method",
    col: "left",
    label: "Metode Inspeksi Bawah Laut",
    fieldType: "select",
    options: [
      { label: "ROV Submersible (Wahana Bawah Air Nirawak)", value: "ROV_SUBMERSIBLE" },
      { label: "Diver Team (Penyelam Militer Dislambair)", value: "DIVER_TEAM" },
      { label: "Side Scan Sonar (Pemindaian Sonar Samping)", value: "SIDE_SCAN_SONAR" },
      { label: "Magnetometer (Deteksi Anomali Magnetik)", value: "MAGNETOMETER" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "condition_rating",
    col: "right",
    label: "Tingkat Kondisi Fisik",
    fieldType: "select",
    options: [
      { label: "EXCELLENT - Kondisi Sangat Baik & Utuh", value: "EXCELLENT" },
      { label: "GOOD - Kondisi Baik & Beroperasi Normal", value: "GOOD" },
      { label: "FAIR - Cukup / Terdapat Keausan Ringan", value: "FAIR" },
      { label: "DAMAGED_CRITICAL - Kerusakan Kritis / Bahaya", value: "DAMAGED_CRITICAL" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "remedial_action_required",
    col: "left",
    label: "Tindakan Perbaikan Diperlukan",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "findings",
    col: "right",
    label: "Temuan Hasil Inspeksi Detail",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.inspection_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  if (!payload.ship_id) payload.ship_id = null;
  if (!payload.inspector_officer_id) payload.inspector_officer_id = null;
  if (!payload.next_inspection_date) payload.next_inspection_date = null;
  return payload;
};
