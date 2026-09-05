/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "daily_log";
export const entityTitle = "Log Harian KRI";
export const entityEndpoint = "/admin/daily_log";
export const primaryKey = "log_id";

export const columns: ColumnField[] = [
  { key: "ship_name", label: "Kapal KRI" },
  { key: "log_date", label: "Log Date" },
  { key: "latitude", label: "Latitude" },
  { key: "longitude", label: "Longitude" },
  { key: "heading_degrees", label: "Heading Degrees" },
  { key: "speed_knots", label: "Speed Knots" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Log Harian KRI", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "ship_id",
    col: "left",
    label: "Kapal KRI",
    fieldType: "select",
    options: {
      url: "/admin/ship?limit=100",
      labelKey: "ship_name",
      valueKey: "ship_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "log_date",
    col: "right",
    label: "Log Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "latitude",
    col: "left",
    label: "Latitude",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "longitude",
    col: "right",
    label: "Longitude",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "heading_degrees",
    col: "left",
    label: "Heading Degrees",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "speed_knots",
    col: "right",
    label: "Speed Knots",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "sea_state",
    col: "left",
    label: "Sea State",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "weather_condition",
    col: "right",
    label: "Weather Condition",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "fuel_remaining_liters",
    col: "left",
    label: "Fuel Remaining Liters",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "fresh_water_remaining_tons",
    col: "right",
    label: "Fresh Water Remaining Tons",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "tactical_summary",
    col: "left",
    label: "Tactical Summary",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
    {
    name: "logged_by_user_id",
    col: "right",
    label: "Petugas Jurnal",
    fieldType: "select",
    options: {
      url: "/admin/sys_user?limit=100",
      labelKey: "full_name",
      valueKey: "user_id",
    },
    required: true,
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
