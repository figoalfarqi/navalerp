/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "stock_balance";
export const entityTitle = "Saldo Stok";
export const entityEndpoint = "/admin/stock_balance";
export const primaryKey = "balance_id";

export const columns: ColumnField[] = [
  { key: "balance_id", label: "Balance Id" },
  { key: "warehouse_id", label: "Warehouse Id" },
  { key: "location_id", label: "Location Id" },
  { key: "material_id", label: "Material Id" },
  { key: "quantity_on_hand", label: "Quantity On Hand" },
  { key: "quantity_reserved", label: "Quantity Reserved" },
  { key: "quantity_in_transit", label: "Quantity In Transit" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Saldo Stok", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "warehouse_id",
    label: "Warehouse Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "location_id",
    label: "Location Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "material_id",
    label: "Material Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "quantity_on_hand",
    label: "Quantity On Hand",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "quantity_reserved",
    label: "Quantity Reserved",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "quantity_in_transit",
    label: "Quantity In Transit",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "last_count_date",
    label: "Last Count Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.balance_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
