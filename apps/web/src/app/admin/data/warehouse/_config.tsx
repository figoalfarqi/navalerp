/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "warehouse";
export const entityTitle = "Gudang Militer";
export const entityEndpoint = "/admin/warehouse";
export const primaryKey = "warehouse_id";

export const columns: ColumnField[] = [
  { key: "warehouse_id", label: "Warehouse Id" },
  { key: "unit_id", label: "Unit Id" },
  { key: "warehouse_code", label: "Warehouse Code" },
  { key: "warehouse_name", label: "Warehouse Name" },
  { key: "warehouse_type", label: "Warehouse Type" },
  { key: "capacity_m3", label: "Capacity M3" },
  { key: "manager_user_id", label: "Manager User Id" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Gudang Militer", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "unit_id",
    label: "Unit Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "warehouse_code",
    label: "Warehouse Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "warehouse_name",
    label: "Warehouse Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "warehouse_type",
    label: "Warehouse Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "capacity_m3",
    label: "Capacity M3",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "manager_user_id",
    label: "Manager User Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "location_address",
    label: "Location Address",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_active",
    label: "Is Active",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.warehouse_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
