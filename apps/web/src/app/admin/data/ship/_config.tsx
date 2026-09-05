/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "ship";
export const entityTitle = "Kapal Perang KRI";
export const entityEndpoint = "/admin/ship";
export const primaryKey = "ship_id";

export const columns: ColumnField[] = [
  { key: "ship_id", label: "Ship Id" },
  { key: "class_id", label: "Class Id" },
  { key: "assigned_unit_id", label: "Assigned Unit Id" },
  { key: "hull_number", label: "Hull Number" },
  { key: "ship_name", label: "Ship Name" },
  { key: "call_sign", label: "Call Sign" },
  { key: "commission_date", label: "Commission Date" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Kapal Perang KRI", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "class_id",
    label: "Class Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "assigned_unit_id",
    label: "Assigned Unit Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "hull_number",
    label: "Hull Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "ship_name",
    label: "Ship Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "call_sign",
    label: "Call Sign",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "commission_date",
    label: "Commission Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "home_port",
    label: "Home Port",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "length_m",
    label: "Length M",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "beam_m",
    label: "Beam M",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "draft_m",
    label: "Draft M",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "displacement_tons",
    label: "Displacement Tons",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "max_speed_knots",
    label: "Max Speed Knots",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "cruise_range_nm",
    label: "Cruise Range Nm",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "crew_capacity",
    label: "Crew Capacity",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "fuel_capacity_liters",
    label: "Fuel Capacity Liters",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "fresh_water_capacity_liters",
    label: "Fresh Water Capacity Liters",
    fieldType: "number",
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
    name: "current_readiness_status",
    label: "Current Readiness Status",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.ship_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
