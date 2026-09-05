/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "berth_booking";
export const entityTitle = "Penjadwalan Sandar";
export const entityEndpoint = "/admin/berth_booking";
export const primaryKey = "booking_id";

export const columns: ColumnField[] = [
  { key: "facility_name", label: "Fasilitas Pangkalan" },
  { key: "ship_name", label: "Kapal KRI" },
  { key: "booking_purpose", label: "Booking Purpose" },
  { key: "eta", label: "Eta" },
  { key: "etd", label: "Etd" },
  { key: "actual_berth_time", label: "Actual Berth Time" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Penjadwalan Sandar", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "facility_id",
    col: "left",
    label: "Fasilitas Pangkalan",
    fieldType: "select",
    options: {
      url: "/admin/base_facility?limit=100",
      labelKey: "facility_name",
      valueKey: "facility_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "ship_id",
    col: "right",
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
    name: "booking_purpose",
    col: "left",
    label: "Booking Purpose",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "eta",
    col: "right",
    label: "Eta",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "etd",
    col: "left",
    label: "Etd",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "actual_berth_time",
    col: "right",
    label: "Actual Berth Time",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_unberth_time",
    col: "left",
    label: "Actual Unberth Time",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "shore_power_kwh_used",
    col: "right",
    label: "Shore Power Kwh Used",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "fresh_water_ton_used",
    col: "left",
    label: "Fresh Water Ton Used",
    fieldType: "number",
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
    {
    name: "approved_by_user_id",
    col: "left",
    label: "Disetujui Oleh",
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
    name: "remarks",
    col: "right",
    label: "Remarks",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.booking_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
