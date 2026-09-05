/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "goods_receipt";
export const entityTitle = "Penerimaan BAPHP";
export const entityEndpoint = "/admin/goods_receipt";
export const primaryKey = "receipt_id";

export const columns: ColumnField[] = [
  { key: "receipt_number", label: "Receipt Number" },
  { key: "po_number", label: "No. PO" },
  { key: "warehouse_name", label: "Nama Gudang" },
  { key: "received_date", label: "Received Date" },
  { key: "delivery_order_number", label: "Delivery Order Number" },
  { key: "inspector_name", label: "Inspektur" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Penerimaan BAPHP", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "receipt_number",
    col: "left",
    label: "Receipt Number",
    fieldType: "text",
    required: true,
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
    required: true,
    disabled: mode === "view",
  },
    {
    name: "warehouse_id",
    col: "left",
    label: "Nama Gudang",
    fieldType: "select",
    options: {
      url: "/admin/warehouse?limit=100",
      labelKey: "warehouse_name",
      valueKey: "warehouse_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "received_date",
    col: "right",
    label: "Received Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "delivery_order_number",
    col: "left",
    label: "Delivery Order Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
    {
    name: "inspected_by_user_id",
    col: "right",
    label: "Inspektur",
    fieldType: "select",
    options: {
      url: "/admin/sys_user?limit=100",
      labelKey: "full_name",
      valueKey: "user_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "inspection_passed",
    col: "left",
    label: "Inspection Passed",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "remarks",
    col: "right",
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
