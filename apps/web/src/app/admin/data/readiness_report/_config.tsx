/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "readiness_report";
export const entityTitle = "Kesiapan KRI";
export const entityEndpoint = "/admin/readiness_report";
export const primaryKey = "snapshot_id";

export const columns: ColumnField[] = [
  { key: "snapshot_id", label: "Snapshot Id" },
  { key: "ship_id", label: "Ship Id" },
  { key: "snapshot_timestamp", label: "Snapshot Timestamp" },
  { key: "readiness_category", label: "Readiness Category" },
  { key: "mro_readiness_score", label: "Mro Readiness Score" },
  { key: "personnel_manning_score", label: "Personnel Manning Score" },
  { key: "logistics_supply_score", label: "Logistics Supply Score" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Kesiapan KRI", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "ship_id",
    label: "Ship Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "snapshot_timestamp",
    label: "Snapshot Timestamp",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "readiness_category",
    label: "Readiness Category",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "mro_readiness_score",
    label: "Mro Readiness Score",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "personnel_manning_score",
    label: "Personnel Manning Score",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "logistics_supply_score",
    label: "Logistics Supply Score",
    fieldType: "number",
    required: true,
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
  delete payload.snapshot_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
