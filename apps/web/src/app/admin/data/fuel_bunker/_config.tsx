/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "fuel_bunker";
export const entityTitle = "Bunker BBM";
export const entityEndpoint = "/admin/fuel_bunker";
export const primaryKey = "bunker_id";

export const columns: ColumnField[] = [
  { key: "ship_name", label: "Kapal KRI" },
  { key: "facility_name", label: "Fasilitas Pangkalan" },
  { key: "fuel_type", label: "Fuel Type" },
  { key: "quantity_liters", label: "Quantity Liters" },
  { key: "density_15c", label: "Density 15c" },
  { key: "flow_rate_lph", label: "Flow Rate Lph" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Bunker BBM", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "ship_id",
    col: "left",
    label: "Kapal KRI",
    fieldType: "select",
    options: {
      url: "/admin/ship?limit=100",
      labelKey: "ship_name",
      valueKey: "ship_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "facility_id",
    col: "right",
    label: "Fasilitas Pangkalan",
    fieldType: "select",
    options: {
      url: "/admin/base_facility?limit=100",
      labelKey: "facility_name",
      valueKey: "facility_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "fuel_type",
    col: "left",
    label: "Fuel Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "quantity_liters",
    col: "right",
    label: "Quantity Liters",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "density_15c",
    col: "left",
    label: "Density 15c",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "flow_rate_lph",
    col: "right",
    label: "Flow Rate Lph",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "bunkering_start_time",
    col: "left",
    label: "Bunkering Start Time",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "bunkering_end_time",
    col: "right",
    label: "Bunkering End Time",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "receipt_voucher_no",
    col: "left",
    label: "Receipt Voucher No",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "authorised_by_user_id",
    col: "right",
    label: "Diotorisasi Oleh",
    fieldType: "select",
    options: {
      url: "/admin/sys_user?limit=100",
      labelKey: "full_name",
      valueKey: "user_id",
    },
    required: true,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.bunker_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
