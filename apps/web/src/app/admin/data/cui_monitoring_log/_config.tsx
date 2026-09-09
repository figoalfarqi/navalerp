/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "cui_monitoring_log";
export const entityTitle = "Log Sensor CUI";
export const entityEndpoint = "/admin/cui_monitoring_log";
export const primaryKey = "log_id";

export const columns: ColumnField[] = [
  { key: "log_id", label: "Log Id" },
  { key: "cui_asset_id", label: "Cui Asset Id" },
  { key: "sensor_code", label: "Sensor Code" },
  { key: "sensor_type", label: "Sensor Type" },
  { key: "log_time", label: "Log Time" },
  { key: "metric_value", label: "Metric Value" },
  { key: "metric_unit", label: "Metric Unit" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Log Sensor CUI", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "cui_asset_id",
    label: "Cui Asset Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "sensor_code",
    label: "Sensor Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "sensor_type",
    label: "Sensor Type",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "log_time",
    label: "Log Time",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "metric_value",
    label: "Metric Value",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "metric_unit",
    label: "Metric Unit",
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
    name: "vessel_proximity_mmsi",
    label: "Vessel Proximity Mmsi",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "anomaly_score",
    label: "Anomaly Score",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "description",
    label: "Description",
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
