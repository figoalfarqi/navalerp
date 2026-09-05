/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "tender";
export const entityTitle = "Tender & Lelang";
export const entityEndpoint = "/admin/tender";
export const primaryKey = "tender_id";

export const columns: ColumnField[] = [
  { key: "tender_number", label: "Tender Number" },
  { key: "title", label: "Title" },
  { key: "procurement_category", label: "Procurement Category" },
  { key: "estimated_budget", label: "Estimated Budget" },
  { key: "procurement_method", label: "Procurement Method" },
  { key: "start_date", label: "Start Date" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Tender & Lelang", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "tender_number",
    col: "left",
    label: "Tender Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
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
    name: "procurement_category",
    col: "left",
    label: "Procurement Category",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "estimated_budget",
    col: "right",
    label: "Estimated Budget",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "procurement_method",
    col: "left",
    label: "Procurement Method",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "start_date",
    col: "right",
    label: "Start Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "closing_date",
    col: "left",
    label: "Closing Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "right",
    label: "Status",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
    {
    name: "winner_vendor_id",
    col: "left",
    label: "Pemenang Tender",
    fieldType: "select",
    options: {
      url: "/admin/vendor?limit=100",
      labelKey: "vendor_name",
      valueKey: "vendor_id",
    },
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.tender_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
