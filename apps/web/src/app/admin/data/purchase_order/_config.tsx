/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "purchase_order";
export const entityTitle = "Purchase Order";
export const entityEndpoint = "/admin/purchase_order";
export const primaryKey = "po_id";

export const columns: ColumnField[] = [
  { key: "po_id", label: "Po Id" },
  { key: "po_number", label: "Po Number" },
  { key: "contract_id", label: "Contract Id" },
  { key: "vendor_id", label: "Vendor Id" },
  { key: "issuing_unit_id", label: "Issuing Unit Id" },
  { key: "order_date", label: "Order Date" },
  { key: "delivery_deadline", label: "Delivery Deadline" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Purchase Order", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "po_number",
    label: "Po Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "contract_id",
    label: "Contract Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "vendor_id",
    label: "Vendor Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "issuing_unit_id",
    label: "Issuing Unit Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "order_date",
    label: "Order Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "delivery_deadline",
    label: "Delivery Deadline",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "destination_warehouse_id",
    label: "Destination Warehouse Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "total_amount",
    label: "Total Amount",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "tax_amount",
    label: "Tax Amount",
    fieldType: "number",
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
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.po_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
