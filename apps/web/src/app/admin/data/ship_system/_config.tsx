/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "ship_system";
export const entityTitle = "Sistem Kapal";
export const entityEndpoint = "/admin/ship_system";
export const primaryKey = "system_id";

export const columns: ColumnField[] = [
  { key: "ship_name", label: "Kapal KRI" },
  { key: "parent_system_name", label: "Sistem Induk" },
  { key: "system_code", label: "System Code" },
  { key: "system_name", label: "System Name" },
  { key: "system_category", label: "System Category" },
  { key: "system_level", label: "System Level" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Sistem Kapal", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "ship_id",
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
    name: "parent_system_id",
    label: "Sistem Induk",
    fieldType: "select",
    options: {
      url: "/admin/ship_system?limit=100",
      labelKey: "system_name",
      valueKey: "system_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "system_code",
    label: "System Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "system_name",
    label: "System Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "system_category",
    label: "System Category",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "system_level",
    label: "System Level",
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
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.system_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
