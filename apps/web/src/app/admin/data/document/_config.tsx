/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "document";
export const entityTitle = "Dokumen Militer";
export const entityEndpoint = "/admin/document";
export const primaryKey = "document_id";

export const columns: ColumnField[] = [
  { key: "document_number", label: "Document Number" },
  { key: "title", label: "Title" },
  { key: "category_name", label: "Kategori Dokumen" },
  { key: "unit_name", label: "Satuan Asal" },
  { key: "classification_level", label: "Classification Level" },
  { key: "effective_date", label: "Effective Date", render: (item: any) => formatSmartDate(item.effective_date) },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Dokumen Militer", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "document_number",
    col: "left",
    label: "Document Number",
    fieldType: "text",
    required: true,
    readOnly: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "DOC",
  },
  {
    name: "title",
    col: "right",
    label: "Title",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "category_id",
    col: "left",
    label: "Kategori Dokumen",
    fieldType: "select",
    options: {
      url: "/admin/document_category?limit=100",
      labelKey: "category_name",
      valueKey: "category_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "originating_unit_id",
    col: "right",
    label: "Satuan Asal",
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
    name: "classification_level",
    col: "left",
    label: "Classification Level",
    fieldType: "select",
    options: [
      { label: "Sangat Rahasia", value: "SANGAT_RAHASIA" },
      { label: "Rahasia Negara", value: "RAHASIA_NEGARA" },
      { label: "Rahasia", value: "RAHASIA" },
      { label: "Terbatas", value: "TERBATAS" },
      { label: "Biasa", value: "BIASA" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "effective_date",
    col: "right",
    label: "Effective Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "expiry_date",
    col: "left",
    label: "Expiry Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "right",
    label: "Status",
    fieldType: "select",
    options: [
      { label: "Draft", value: "DRAFT" },
      { label: "Review", value: "REVIEW" },
      { label: "Approved", value: "APPROVED" },
      { label: "Active", value: "ACTIVE" },
      { label: "Archived", value: "ARCHIVED" },
      { label: "Superseded", value: "SUPERSEDED" },
    ],
    required: false,
    disabled: mode === "view",
  },
    {
    name: "approved_by_user_id",
    col: "left",
    label: "Disetujui Oleh",
    fieldType: "select",
    options: {
      url: "/admin/sys_user?limit=100",
      labelKey: "full_name",
      valueKey: "user_id",
    },
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
