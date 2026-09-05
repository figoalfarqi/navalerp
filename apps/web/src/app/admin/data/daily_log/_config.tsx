/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "daily_log";
export const entityTitle = "Log Harian KRI";
export const entityEndpoint = "/admin/daily_log";
export const primaryKey = "log_id";

export const columns: ColumnField[] = [
  { key: "log_id", label: "Log Id" },
  { key: "ship_id", label: "Ship Id" },
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
    label: "Ship Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "log_date",
    label: "Log Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "latitude",
    label: "Latitude",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "longitude",
    label: "Longitude",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "heading_degrees",
    label: "Heading Degrees",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "speed_knots",
    label: "Speed Knots",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "sea_state",
    label: "Sea State",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "weather_condition",
    label: "Weather Condition",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "fuel_remaining_liters",
    label: "Fuel Remaining Liters",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "fresh_water_remaining_tons",
    label: "Fresh Water Remaining Tons",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "tactical_summary",
    label: "Tactical Summary",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "logged_by_user_id",
    label: "Logged By User Id",
    fieldType: "text",
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
