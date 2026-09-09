/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "material";
export const entityTitle = "Material & Suku Cadang";
export const entityEndpoint = "/admin/material";
export const primaryKey = "material_id";

export const columns: ColumnField[] = [
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
    col: "left",
    label: "Material Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "nsn",
    col: "right",
    label: "Nsn",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "part_number",
    col: "left",
    label: "Part Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "oem_name",
    col: "right",
    label: "Oem Name",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "material_name",
    col: "left",
    label: "Material Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "category",
    col: "right",
    label: "Category",
    fieldType: "select",
    options: [
      { label: "Technical Spares", value: "TECHNICAL_SPARES" },
      { label: "Consumables", value: "CONSUMABLES" },
      { label: "Fuel Energy", value: "FUEL_ENERGY" },
      { label: "Ammunition", value: "AMMUNITION" },
      { label: "General Supplies", value: "GENERAL_SUPPLIES" },
      { label: "Strategic Reserve", value: "STRATEGIC_RESERVE" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "uom",
    col: "left",
    label: "Uom",
    fieldType: "select",
    options: [
      { label: "Unit", value: "UNIT" },
      { label: "Set", value: "SET" },
      { label: "Liter", value: "LITER" },
      { label: "Kg", value: "KG" },
      { label: "Ton", value: "TON" },
      { label: "Meter", value: "METER" },
      { label: "Box", value: "BOX" },
      { label: "Drum", value: "DRUM" },
      { label: "Crate", value: "CRATE" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "weight_kg",
    col: "right",
    label: "Weight Kg",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "min_stock_level",
    col: "left",
    label: "Min Stock Level",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "max_stock_level",
    col: "right",
    label: "Max Stock Level",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "reorder_point",
    col: "left",
    label: "Reorder Point",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "safety_stock",
    col: "right",
    label: "Safety Stock",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "shelf_life_days",
    col: "left",
    label: "Shelf Life Days",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_controlled_item",
    col: "right",
    label: "Is Controlled Item",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "unit_price_idr",
    col: "left",
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
