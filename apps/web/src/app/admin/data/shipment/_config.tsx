/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "shipment";
export const entityTitle = "Pengiriman Konvoi";
export const entityEndpoint = "/admin/shipment";
export const primaryKey = "shipment_id";

export const columns: ColumnField[] = [
  { key: "manifest_number", label: "Manifest Number" },
  { key: "route_name", label: "Rute Logistik" },
  { key: "transport_unit_name", label: "Unit Angkut" },
  { key: "origin_warehouse_name", label: "Gudang Asal" },
  { key: "destination_warehouse_name", label: "Gudang Tujuan" },
  { key: "departure_date", label: "Departure Date" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Pengiriman Konvoi", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "manifest_number",
    col: "left",
    label: "Manifest Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "route_id",
    col: "right",
    label: "Rute Logistik",
    fieldType: "select",
    options: {
      url: "/admin/route?limit=100",
      labelKey: "route_name",
      valueKey: "route_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "transport_unit_id",
    col: "left",
    label: "Unit Angkut",
    fieldType: "select",
    options: {
      url: "/admin/transport_unit?limit=100",
      labelKey: "unit_code",
      valueKey: "transport_unit_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "origin_warehouse_id",
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
    name: "destination_warehouse_id",
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
    name: "departure_date",
    col: "right",
    label: "Departure Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "arrival_date",
    col: "left",
    label: "Arrival Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "escort_security_level",
    col: "right",
    label: "Escort Security Level",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "left",
    label: "Status",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
    {
    name: "authorized_by_user_id",
    col: "right",
    label: "Diotorisasi Oleh",
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
    col: "left",
    label: "Remarks",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.shipment_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
