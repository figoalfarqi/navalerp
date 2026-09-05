/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "stock_transfer";
export const entityTitle = "Transfer Bebekal";
export const entityEndpoint = "/admin/stock_transfer";
export const primaryKey = "transfer_id";

export const columns: ColumnField[] = [
  { key: "transfer_id", label: "Transfer Id" },
  { key: "transfer_number", label: "Transfer Number" },
  { key: "from_warehouse_id", label: "From Warehouse Id" },
  { key: "to_warehouse_id", label: "To Warehouse Id" },
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
    label: "Transfer Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "from_warehouse_id",
    label: "From Warehouse Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "to_warehouse_id",
    label: "To Warehouse Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "movement_type",
    label: "Movement Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "scheduled_departure",
    label: "Scheduled Departure",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_departure",
    label: "Actual Departure",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "scheduled_arrival",
    label: "Scheduled Arrival",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_arrival",
    label: "Actual Arrival",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "transporter_unit",
    label: "Transporter Unit",
    fieldType: "text",
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
  delete payload.transfer_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
