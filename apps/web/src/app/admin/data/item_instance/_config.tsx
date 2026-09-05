/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "item_instance";
export const entityTitle = "Nomor Seri / Batch";
export const entityEndpoint = "/admin/item_instance";
export const primaryKey = "instance_id";

export const columns: ColumnField[] = [
  { key: "warehouse_name", label: "Nama Gudang" },
  { key: "location_id", label: "Location Id" },
  { key: "material_name", label: "Material" },
  { key: "batch_number", label: "Batch Number" },
  { key: "serial_number", label: "Serial Number" },
  { key: "lot_number", label: "Lot Number" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Nomor Seri / Batch", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "warehouse_id",
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
    name: "location_id",
    label: "Location Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
    {
    name: "material_id",
    label: "Material",
    fieldType: "select",
    options: {
      url: "/admin/material?limit=100",
      labelKey: "material_name",
      valueKey: "material_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "batch_number",
    label: "Batch Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "serial_number",
    label: "Serial Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "lot_number",
    label: "Lot Number",
    fieldType: "text",
    required: false,
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
    name: "manufactured_date",
    label: "Manufactured Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "condition",
    label: "Condition",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "inspection_due_date",
    label: "Inspection Due Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "quantity",
    label: "Quantity",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.instance_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
