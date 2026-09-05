/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "ship_class";
export const entityTitle = "Kelas Kapal";
export const entityEndpoint = "/admin/ship_class";
export const primaryKey = "class_id";

export const columns: ColumnField[] = [
  { key: "class_id", label: "Class Id" },
  { key: "class_code", label: "Class Code" },
  { key: "class_name", label: "Class Name" },
  { key: "category", label: "Category" },
  { key: "specifications", label: "Specifications" },
  { key: "builder", label: "Builder" },
  { key: "total_built", label: "Total Built" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Kelas Kapal", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "class_code",
    label: "Class Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "class_name",
    label: "Class Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "category",
    label: "Category",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "specifications",
    label: "Specifications",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "builder",
    label: "Builder",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "total_built",
    label: "Total Built",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.class_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
