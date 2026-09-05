/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "document_category";
export const entityTitle = "Kategori Dokumen";
export const entityEndpoint = "/admin/document_category";
export const primaryKey = "category_id";

export const columns: ColumnField[] = [
  { key: "category_code", label: "Category Code" },
  { key: "category_name", label: "Category Name" },
  { key: "retention_years", label: "Retention Years" },
  { key: "confidentiality_level", label: "Confidentiality Level" },
  { key: "description", label: "Description" },
  { key: "created_at", label: "Created At" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Kategori Dokumen", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "category_code",
    label: "Category Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "category_name",
    label: "Category Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "retention_years",
    label: "Retention Years",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "confidentiality_level",
    label: "Confidentiality Level",
    fieldType: "text",
    required: false,
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
  delete payload.category_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
