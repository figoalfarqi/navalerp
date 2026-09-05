/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "route";
export const entityTitle = "Rute Pelayaran";
export const entityEndpoint = "/admin/route";
export const primaryKey = "route_id";

export const columns: ColumnField[] = [
  { key: "route_code", label: "Route Code" },
  { key: "route_name", label: "Route Name" },
  { key: "origin_facility_name", label: "Fasilitas Asal" },
  { key: "destination_facility_name", label: "Fasilitas Tujuan" },
  { key: "distance_nautical_miles", label: "Distance Nautical Miles" },
  { key: "estimated_transit_hours", label: "Estimated Transit Hours" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Rute Pelayaran", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "route_code",
    col: "left",
    label: "Route Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "route_name",
    col: "right",
    label: "Route Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "origin_facility_id",
    col: "left",
    label: "Fasilitas Asal",
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
    name: "destination_facility_id",
    col: "right",
    label: "Fasilitas Tujuan",
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
    name: "distance_nautical_miles",
    col: "left",
    label: "Distance Nautical Miles",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "estimated_transit_hours",
    col: "right",
    label: "Estimated Transit Hours",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "risk_level",
    col: "left",
    label: "Risk Level",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.route_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
