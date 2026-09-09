/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "cui_asset";
export const entityTitle = "Aset Bawah Laut";
export const entityEndpoint = "/admin/cui_asset";
export const primaryKey = "cui_asset_id";

export const columns: ColumnField[] = [
  { key: "cui_asset_id", label: "Cui Asset Id" },
  { key: "asset_code", label: "Asset Code" },
  { key: "asset_name", label: "Asset Name" },
  { key: "asset_type", label: "Asset Type" },
  { key: "operator_name", label: "Operator Name" },
  { key: "theater_id", label: "Theater Id" },
  { key: "depth_meters", label: "Depth Meters" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Aset Bawah Laut", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "asset_code",
    label: "Asset Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "asset_name",
    label: "Asset Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "asset_type",
    label: "Asset Type",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "operator_name",
    label: "Operator Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "theater_id",
    label: "Theater Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "depth_meters",
    label: "Depth Meters",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "length_km",
    label: "Length Km",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "latitude",
    label: "Latitude",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "longitude",
    label: "Longitude",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "start_coordinates",
    label: "Start Coordinates",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "end_coordinates",
    label: "End Coordinates",
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
    name: "health_score",
    label: "Health Score",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "protection_priority",
    label: "Protection Priority",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "last_inspected_at",
    label: "Last Inspected At",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "next_inspection_due",
    label: "Next Inspection Due",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "notes",
    label: "Notes",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.cui_asset_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
