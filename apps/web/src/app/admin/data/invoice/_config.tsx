/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "invoice";
export const entityTitle = "Tagihan Rekanan";
export const entityEndpoint = "/admin/invoice";
export const primaryKey = "invoice_id";

export const columns: ColumnField[] = [
  { key: "invoice_number", label: "Invoice Number" },
  { key: "vendor_name", label: "Vendor" },
  { key: "contract_number", label: "No. Kontrak" },
  { key: "po_number", label: "No. PO" },
  { key: "invoice_date", label: "Invoice Date" },
  { key: "due_date", label: "Due Date" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Tagihan Rekanan", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "invoice_number",
    col: "left",
    label: "Invoice Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "vendor_id",
    col: "right",
    label: "Vendor",
    fieldType: "select",
    options: {
      url: "/admin/vendor?limit=100",
      labelKey: "vendor_name",
      valueKey: "vendor_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "contract_id",
    col: "left",
    label: "No. Kontrak",
    fieldType: "select",
    options: {
      url: "/admin/contract?limit=100",
      labelKey: "contract_number",
      valueKey: "contract_id",
    },
    required: false,
    disabled: mode === "view",
  },
    {
    name: "po_id",
    col: "right",
    label: "No. PO",
    fieldType: "select",
    options: {
      url: "/admin/purchase_order?limit=100",
      labelKey: "po_number",
      valueKey: "po_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "invoice_date",
    col: "left",
    label: "Invoice Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "due_date",
    col: "right",
    label: "Due Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "tax_invoice_number",
    col: "left",
    label: "Tax Invoice Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "subtotal",
    col: "right",
    label: "Subtotal",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "tax_amount",
    col: "left",
    label: "Tax Amount",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "verification_status",
    col: "right",
    label: "Verification Status",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
    {
    name: "verified_by_user_id",
    col: "left",
    label: "Diverifikasi Oleh",
    fieldType: "select",
    options: {
      url: "/admin/sys_user?limit=100",
      labelKey: "full_name",
      valueKey: "user_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "payment_status",
    col: "right",
    label: "Payment Status",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.invoice_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
