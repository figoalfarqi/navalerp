/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "ship";
export const entityTitle = "Kapal Perang KRI";
export const entityEndpoint = "/admin/ship";
export const primaryKey = "ship_id";

export const columns: ColumnField[] = [
  { key: "class_name", label: "Kelas Kapal" },
  { key: "unit_name", label: "Satuan / Pangkalan" },
  { key: "hull_number", label: "Hull Number" },
  { key: "ship_name", label: "Ship Name" },
  { key: "call_sign", label: "Call Sign" },
  { key: "commission_date", label: "Commission Date", render: (item: any) => formatSmartDate(item.commission_date) },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Kapal Perang KRI", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "class_id",
    col: "left",
    label: "Kelas Kapal",
    fieldType: "select",
    options: {
      url: "/admin/ship_class?limit=100",
      labelKey: "class_name",
      valueKey: "class_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "assigned_unit_id",
    col: "right",
    label: "Satuan / Pangkalan",
    fieldType: "select",
    options: {
      url: "/admin/org_unit?limit=100",
      labelKey: "unit_name",
      valueKey: "unit_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "hull_number",
    col: "left",
    label: "Hull Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "ship_name",
    col: "right",
    label: "Ship Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "call_sign",
    col: "left",
    label: "Call Sign",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "commission_date",
    col: "right",
    label: "Commission Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "home_port",
    col: "left",
    label: "Home Port",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "length_m",
    col: "right",
    label: "Length M",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "beam_m",
    col: "left",
    label: "Beam M",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "draft_m",
    col: "right",
    label: "Draft M",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "displacement_tons",
    col: "left",
    label: "Displacement Tons",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "max_speed_knots",
    col: "right",
    label: "Max Speed Knots",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "cruise_range_nm",
    col: "left",
    label: "Cruise Range Nm",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "crew_capacity",
    col: "right",
    label: "Crew Capacity",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "fuel_capacity_liters",
    col: "left",
    label: "Fuel Capacity Liters",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "fresh_water_capacity_liters",
    col: "right",
    label: "Fresh Water Capacity Liters",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "left",
    label: "Status",
    fieldType: "select",
    options: [
      { label: "Active", value: "ACTIVE" },
      { label: "Docked", value: "DOCKED" },
      { label: "Retired", value: "RETIRED" },
      { label: "Deployed", value: "DEPLOYED" },
      { label: "Standby", value: "STANDBY" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "current_readiness_status",
    col: "right",
    label: "Current Readiness Status",
    fieldType: "select",
    options: [
      { label: "Fully Mission Capable", value: "FULLY_MISSION_CAPABLE" },
      { label: "Partially Mission Capable", value: "PARTIALLY_MISSION_CAPABLE" },
      { label: "Non Mission Capable", value: "NON_MISSION_CAPABLE" },
    ],
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
