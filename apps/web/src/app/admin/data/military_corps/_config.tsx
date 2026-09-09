/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "military_corps";
export const entityTitle = "Korps Militer";
export const entityEndpoint = "/admin/military_corps";
export const primaryKey = "corps_id";

export const columns: ColumnField[] = [
  { key: "corps_code", label: "Corps Code" },
  { key: "corps_name", label: "Corps Name" },
  { key: "description", label: "Description" },
  { key: "created_at", label: "Created At", render: (item: any) => formatSmartDate(item.created_at) },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Korps Militer", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "corps_code",
    col: "left",
    label: "Corps Code",
    fieldType: "select",
    options: [
      { label: "P", value: "P" },
      { label: "T", value: "T" },
      { label: "E", value: "E" },
      { label: "S", value: "S" },
      { label: "M", value: "M" },
      { label: "K", value: "K" },
      { label: "KH", value: "KH" },
      { label: "PM", value: "PM" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "corps_name",
    col: "right",
    label: "Corps Name",
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
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.corps_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
