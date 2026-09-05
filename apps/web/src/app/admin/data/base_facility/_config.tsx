/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "base_facility";
export const entityTitle = "Fasilitas Pangkalan";
export const entityEndpoint = "/admin/base_facility";
export const primaryKey = "facility_id";

export const columns: ColumnField[] = [
  { key: "unit_name", label: "Satuan Pangkalan" },
  { key: "facility_code", label: "Facility Code" },
  { key: "facility_name", label: "Facility Name" },
  { key: "facility_type", label: "Facility Type" },
  { key: "length_meters", label: "Length Meters" },
  { key: "draft_depth_meters", label: "Draft Depth Meters" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Fasilitas Pangkalan", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "base_unit_id",
    col: "left",
    label: "Satuan Pangkalan",
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
    name: "facility_code",
    col: "right",
    label: "Facility Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "facility_name",
    col: "left",
    label: "Facility Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "facility_type",
    col: "right",
    label: "Facility Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "length_meters",
    col: "left",
    label: "Length Meters",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "draft_depth_meters",
    col: "right",
    label: "Draft Depth Meters",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "max_displacement_tonnage",
    col: "left",
    label: "Max Displacement Tonnage",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "has_shore_power",
    col: "right",
    label: "Has Shore Power",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "has_fresh_water",
    col: "left",
    label: "Has Fresh Water",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "has_fuel_bunker_line",
    col: "right",
    label: "Has Fuel Bunker Line",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "left",
    label: "Status",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.facility_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
