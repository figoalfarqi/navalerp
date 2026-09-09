/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "transport_unit";
export const entityTitle = "Armada Transportasi";
export const entityEndpoint = "/admin/transport_unit";
export const primaryKey = "transport_unit_id";

export const columns: ColumnField[] = [
  { key: "unit_code", label: "Unit Code" },
  { key: "unit_name", label: "Unit Name" },
  { key: "transport_type", label: "Transport Type" },
  { key: "cargo_capacity_tons", label: "Cargo Capacity Tons" },
  { key: "fuel_capacity_liters", label: "Fuel Capacity Liters" },
  { key: "unit_name", label: "Satuan Operasi" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Armada Transportasi", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "unit_code",
    col: "left",
    label: "Unit Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "unit_name",
    col: "right",
    label: "Unit Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "transport_type",
    col: "left",
    label: "Transport Type",
    fieldType: "select",
    options: [
      { label: "Naval Auxiliary Vessel", value: "NAVAL_AUXILIARY_VESSEL" },
      { label: "Land Truck", value: "LAND_TRUCK" },
      { label: "Air Cargo", value: "AIR_CARGO" },
      { label: "Cargo Barge", value: "CARGO_BARGE" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "cargo_capacity_tons",
    col: "right",
    label: "Cargo Capacity Tons",
    fieldType: "number",
    required: true,
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
    name: "operating_unit_id",
    col: "right",
    label: "Satuan Operasi",
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
    name: "status",
    col: "left",
    label: "Status",
    fieldType: "select",
    options: [
      { label: "Available", value: "AVAILABLE" },
      { label: "On Mission", value: "ON_MISSION" },
      { label: "Maintenance", value: "MAINTENANCE" },
      { label: "Decommissioned", value: "DECOMMISSIONED" },
    ],
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.transport_unit_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
