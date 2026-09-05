/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "stock_adjustment";
export const entityTitle = "Penyesuaian Stok";
export const entityEndpoint = "/admin/stock_adjustment";
export const primaryKey = "adjustment_id";

export const columns: ColumnField[] = [
  { key: "warehouse_name", label: "Nama Gudang" },
  { key: "adjustment_number", label: "Adjustment Number" },
  { key: "adjustment_date", label: "Adjustment Date" },
  { key: "conductor_name", label: "Pelaksana Stock Take" },
  { key: "reason", label: "Reason" },
  { key: "status", label: "Status" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Penyesuaian Stok", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
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
    name: "adjustment_number",
    label: "Adjustment Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "adjustment_date",
    label: "Adjustment Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "conducted_by_user_id",
    label: "Pelaksana Stock Take",
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
    name: "reason",
    label: "Reason",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "status",
    label: "Status",
    fieldType: "text",
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
  delete payload.adjustment_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
