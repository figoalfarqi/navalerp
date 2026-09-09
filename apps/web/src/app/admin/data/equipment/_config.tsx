/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "equipment";
export const entityTitle = "Peralatan Mesin";
export const entityEndpoint = "/admin/equipment";
export const primaryKey = "equipment_id";

export const columns: ColumnField[] = [
  { key: "system_name", label: "Sistem Kapal" },
  { key: "serial_number", label: "Serial Number" },
  { key: "equipment_tag", label: "Equipment Tag" },
  { key: "equipment_name", label: "Equipment Name" },
  { key: "manufacturer", label: "Manufacturer" },
  { key: "model_number", label: "Model Number" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Peralatan Mesin", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "system_id",
    col: "left",
    label: "Sistem Kapal",
    fieldType: "select",
    options: {
      url: "/admin/ship_system?limit=100",
      labelKey: "system_name",
      valueKey: "system_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "serial_number",
    col: "right",
    label: "Serial Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "equipment_tag",
    col: "left",
    label: "Equipment Tag",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "equipment_name",
    col: "right",
    label: "Equipment Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "manufacturer",
    col: "left",
    label: "Manufacturer",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "model_number",
    col: "right",
    label: "Model Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "country_of_origin",
    col: "left",
    label: "Country Of Origin",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "installation_date",
    col: "right",
    label: "Installation Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "total_operating_hours",
    col: "left",
    label: "Total Operating Hours",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "design_life_hours",
    col: "right",
    label: "Design Life Hours",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "criticality_level",
    col: "left",
    label: "Criticality Level",
    fieldType: "select",
    options: [
      { label: "Critical Safety", value: "CRITICAL_SAFETY" },
      { label: "Mission Essential", value: "MISSION_ESSENTIAL" },
      { label: "Routine", value: "ROUTINE" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "health_status",
    col: "right",
    label: "Health Status",
    fieldType: "select",
    options: [
      { label: "Operational", value: "OPERATIONAL" },
      { label: "Degraded", value: "DEGRADED" },
      { label: "Critical", value: "CRITICAL" },
      { label: "Non Operational", value: "NON_OPERATIONAL" },
    ],
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.equipment_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
