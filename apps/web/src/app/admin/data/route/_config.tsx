/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "route";
export const entityTitle = "Rute Pelayaran";
export const entityEndpoint = "/admin/route";
export const primaryKey = "route_id";

export const columns: ColumnField[] = [
  { key: "route_id", label: "Route Id" },
  { key: "route_code", label: "Route Code" },
  { key: "route_name", label: "Route Name" },
  { key: "origin_facility_id", label: "Origin Facility Id" },
  { key: "destination_facility_id", label: "Destination Facility Id" },
  { key: "distance_nautical_miles", label: "Distance Nautical Miles" },
  { key: "estimated_transit_hours", label: "Estimated Transit Hours" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Rute Pelayaran", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "route_code",
    label: "Route Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "route_name",
    label: "Route Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "origin_facility_id",
    label: "Origin Facility Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "destination_facility_id",
    label: "Destination Facility Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "distance_nautical_miles",
    label: "Distance Nautical Miles",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "estimated_transit_hours",
    label: "Estimated Transit Hours",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "risk_level",
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
