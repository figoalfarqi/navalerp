/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "journal_entry";
export const entityTitle = "Jurnal Akuntansi";
export const entityEndpoint = "/admin/journal_entry";
export const primaryKey = "journal_id";

export const columns: ColumnField[] = [
  { key: "entry_number", label: "Entry Number" },
  { key: "entry_date", label: "Entry Date" },
  { key: "description", label: "Description" },
  { key: "source_module", label: "Source Module" },
  { key: "source_reference_id", label: "Source Reference Id" },
  { key: "is_posted", label: "Is Posted" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Jurnal Akuntansi", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "entry_number",
    col: "left",
    label: "Entry Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "entry_date",
    col: "right",
    label: "Entry Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "description",
    col: "left",
    label: "Description",
    fieldType: "textarea",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "source_module",
    col: "right",
    label: "Source Module",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "source_reference_id",
    col: "left",
    label: "Source Reference Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_posted",
    col: "right",
    label: "Is Posted",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.journal_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
