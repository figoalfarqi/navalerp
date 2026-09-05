/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "document";
export const entityTitle = "Dokumen Militer";
export const entityEndpoint = "/admin/document";
export const primaryKey = "document_id";

export const columns: ColumnField[] = [
  { key: "document_id", label: "Document Id" },
  { key: "document_number", label: "Document Number" },
  { key: "title", label: "Title" },
  { key: "category_id", label: "Category Id" },
  { key: "originating_unit_id", label: "Originating Unit Id" },
  { key: "classification_level", label: "Classification Level" },
  { key: "effective_date", label: "Effective Date" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Dokumen Militer", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "document_number",
    label: "Document Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "title",
    label: "Title",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "category_id",
    label: "Category Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "originating_unit_id",
    label: "Originating Unit Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "classification_level",
    label: "Classification Level",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "effective_date",
    label: "Effective Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "expiry_date",
    label: "Expiry Date",
    fieldType: "date",
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
    name: "approved_by_user_id",
    label: "Approved By User Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.document_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
