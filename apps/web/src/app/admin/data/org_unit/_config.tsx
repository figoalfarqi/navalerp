/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "org_unit";
export const entityTitle = "Satuan Kerja";
export const entityEndpoint = "/admin/org_unit";
export const primaryKey = "unit_id";

export const columns: ColumnField[] = [
  { key: "parent_unit_name", label: "Satuan Induk" },
  { key: "unit_code", label: "Unit Code" },
  { key: "unit_name", label: "Unit Name" },
  { key: "unit_type", label: "Unit Type" },
  { key: "description", label: "Description" },
  { key: "command_level", label: "Command Level" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Satuan Kerja", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "parent_unit_id",
    col: "left",
    label: "Satuan Induk",
    fieldType: "select",
    options: {
      url: "/admin/org_unit?limit=100",
      labelKey: "unit_name",
      valueKey: "unit_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "unit_code",
    col: "right",
    label: "Unit Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "unit_name",
    col: "left",
    label: "Unit Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "unit_type",
    col: "right",
    label: "Unit Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "description",
    col: "left",
    label: "Description",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "command_level",
    col: "right",
    label: "Command Level",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "latitude",
    col: "left",
    label: "Latitude",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "longitude",
    col: "right",
    label: "Longitude",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "address",
    col: "left",
    label: "Address",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "phone",
    col: "right",
    label: "Phone",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_active",
    col: "left",
    label: "Is Active",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.unit_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
