/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "payment";
export const entityTitle = "Pembayaran SP2D";
export const entityEndpoint = "/admin/payment";
export const primaryKey = "payment_id";

export const columns: ColumnField[] = [
  { key: "payment_id", label: "Payment Id" },
  { key: "payment_reference_no", label: "Payment Reference No" },
  { key: "spp_number", label: "Spp Number" },
  { key: "spm_number", label: "Spm Number" },
  { key: "invoice_id", label: "Invoice Id" },
  { key: "payment_date", label: "Payment Date" },
  { key: "amount_paid", label: "Amount Paid" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Pembayaran SP2D", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "payment_reference_no",
    label: "Payment Reference No",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "spp_number",
    label: "Spp Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "spm_number",
    label: "Spm Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "invoice_id",
    label: "Invoice Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "payment_date",
    label: "Payment Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "amount_paid",
    label: "Amount Paid",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "payment_method",
    label: "Payment Method",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "bank_source_account",
    label: "Bank Source Account",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "authorised_by_user_id",
    label: "Authorised By User Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.payment_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
