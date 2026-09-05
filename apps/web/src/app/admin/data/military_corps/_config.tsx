/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "military_corps";
export const entityTitle = "Korps Militer";
export const entityEndpoint = "/admin/military_corps";
export const primaryKey = "corps_id";

export const columns: ColumnField[] = [
  { key: "corps_id", label: "Corps Id" },
  { key: "corps_code", label: "Corps Code" },
  { key: "corps_name", label: "Corps Name" },
  { key: "description", label: "Description" },
  { key: "created_at", label: "Created At" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Korps Militer", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "corps_code",
    label: "Corps Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "corps_name",
    label: "Corps Name",
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
  delete payload.corps_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
