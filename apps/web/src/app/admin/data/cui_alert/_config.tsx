/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "cui_alert";
export const entityTitle = "Peringatan CUI";
export const entityEndpoint = "/admin/cui_alert";
export const primaryKey = "alert_id";

export const columns: ColumnField[] = [
  { key: "alert_id", label: "Alert Id" },
  { key: "alert_code", label: "Alert Code" },
  { key: "cui_asset_id", label: "Cui Asset Id" },
  { key: "alert_type", label: "Alert Type" },
  { key: "severity", label: "Severity" },
  { key: "detected_at", label: "Detected At" },
  { key: "assigned_ship_id", label: "Assigned Ship Id" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Peringatan CUI", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "alert_code",
    label: "Alert Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "cui_asset_id",
    label: "Cui Asset Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "alert_type",
    label: "Alert Type",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "severity",
    label: "Severity",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "detected_at",
    label: "Detected At",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "assigned_ship_id",
    label: "Assigned Ship Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    label: "Status",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "ai_confidence",
    label: "Ai Confidence",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "recommended_action",
    label: "Recommended Action",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "resolution_notes",
    label: "Resolution Notes",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "resolved_at",
    label: "Resolved At",
    fieldType: "date",
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
  return payload;
};
