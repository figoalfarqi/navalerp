/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "goods_receipt";
export const entityTitle = "Penerimaan BAPHP";
export const entityEndpoint = "/admin/goods_receipt";
export const primaryKey = "receipt_id";

export const columns: ColumnField[] = [
  { key: "receipt_id", label: "Receipt Id" },
  { key: "receipt_number", label: "Receipt Number" },
  { key: "po_id", label: "Po Id" },
  { key: "warehouse_id", label: "Warehouse Id" },
  { key: "received_date", label: "Received Date" },
  { key: "delivery_order_number", label: "Delivery Order Number" },
  { key: "inspected_by_user_id", label: "Inspected By User Id" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Penerimaan BAPHP", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "receipt_number",
    label: "Receipt Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "po_id",
    label: "Po Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "warehouse_id",
    label: "Warehouse Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "received_date",
    label: "Received Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "delivery_order_number",
    label: "Delivery Order Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "inspected_by_user_id",
    label: "Inspected By User Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "inspection_passed",
    label: "Inspection Passed",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "remarks",
    label: "Remarks",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.receipt_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
