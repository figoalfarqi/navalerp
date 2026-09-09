/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "ship_class";
export const entityTitle = "Kelas Kapal";
export const entityEndpoint = "/admin/ship_class";
export const primaryKey = "class_id";

export const columns: ColumnField[] = [
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
    col: "left",
    label: "Class Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "class_name",
    col: "right",
    label: "Class Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "category",
    col: "left",
    label: "Category",
    fieldType: "select",
    options: [
      { label: "Frigate", value: "FRIGATE" },
      { label: "Corvette", value: "CORVETTE" },
      { label: "Submarine", value: "SUBMARINE" },
      { label: "LPD", value: "LPD" },
      { label: "Patrol", value: "PATROL" },
      { label: "Aircraft", value: "AIRCRAFT" },
      { label: "Fast Attack", value: "FAST_ATTACK" },
      { label: "Auxiliary", value: "AUXILIARY" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "specifications",
    col: "right",
    label: "Specifications",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "builder",
    col: "left",
    label: "Builder",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "total_built",
    col: "right",
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
