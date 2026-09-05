/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "stock_transfer";
export const entityTitle = "Transfer Bebekal";
export const entityEndpoint = "/admin/stock_transfer";
export const primaryKey = "transfer_id";

export const columns: ColumnField[] = [
  { key: "transfer_number", label: "Transfer Number" },
  { key: "source_warehouse_name", label: "Gudang Asal" },
  { key: "dest_warehouse_name", label: "Gudang Tujuan" },
  { key: "movement_type", label: "Movement Type" },
  { key: "scheduled_departure", label: "Scheduled Departure" },
  { key: "actual_departure", label: "Actual Departure" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Transfer Bebekal", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "transfer_number",
    col: "left",
    label: "Transfer Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "from_warehouse_id",
    col: "right",
    label: "Gudang Asal",
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
    name: "to_warehouse_id",
    col: "left",
    label: "Gudang Tujuan",
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
    name: "movement_type",
    col: "right",
    label: "Movement Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "scheduled_departure",
    col: "left",
    label: "Scheduled Departure",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_departure",
    col: "right",
    label: "Actual Departure",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "scheduled_arrival",
    col: "left",
    label: "Scheduled Arrival",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_arrival",
    col: "right",
    label: "Actual Arrival",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "transporter_unit",
    col: "left",
    label: "Transporter Unit",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "right",
    label: "Status",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.transfer_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
