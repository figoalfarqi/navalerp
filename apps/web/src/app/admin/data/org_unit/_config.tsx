/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "org_unit";
export const entityTitle = "Satuan Kerja";
export const entityEndpoint = "/admin/org_unit";
export const primaryKey = "unit_id";

export const columns: ColumnField[] = [
  { key: "unit_id", label: "Unit Id" },
  { key: "parent_unit_id", label: "Parent Unit Id" },
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
    label: "Parent Unit Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "unit_code",
    label: "Unit Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "unit_name",
    label: "Unit Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "unit_type",
    label: "Unit Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "description",
    label: "Description",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "command_level",
    label: "Command Level",
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
    name: "address",
    label: "Address",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "phone",
    label: "Phone",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_active",
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
