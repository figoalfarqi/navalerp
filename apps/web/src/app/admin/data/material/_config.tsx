/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "material";
export const entityTitle = "Material & Suku Cadang";
export const entityEndpoint = "/admin/material";
export const primaryKey = "material_id";

export const columns: ColumnField[] = [
  { key: "material_id", label: "Material Id" },
  { key: "material_code", label: "Material Code" },
  { key: "nsn", label: "Nsn" },
  { key: "part_number", label: "Part Number" },
  { key: "oem_name", label: "Oem Name" },
  { key: "material_name", label: "Material Name" },
  { key: "category", label: "Category" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Material & Suku Cadang", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "material_code",
    label: "Material Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "nsn",
    label: "Nsn",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "part_number",
    label: "Part Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "oem_name",
    label: "Oem Name",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "material_name",
    label: "Material Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "category",
    label: "Category",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "uom",
    label: "Uom",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "weight_kg",
    label: "Weight Kg",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "min_stock_level",
    label: "Min Stock Level",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "max_stock_level",
    label: "Max Stock Level",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "reorder_point",
    label: "Reorder Point",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "safety_stock",
    label: "Safety Stock",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "shelf_life_days",
    label: "Shelf Life Days",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_controlled_item",
    label: "Is Controlled Item",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "unit_price_idr",
    label: "Unit Price Idr",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.material_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
