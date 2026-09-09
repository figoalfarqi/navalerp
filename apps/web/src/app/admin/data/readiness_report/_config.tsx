/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "readiness_report";
export const entityTitle = "Kesiapan KRI";
export const entityEndpoint = "/admin/readiness_report";
export const primaryKey = "snapshot_id";

export const columns: ColumnField[] = [
  { key: "ship_name", label: "Kapal KRI" },
  { key: "snapshot_timestamp", label: "Snapshot Timestamp", render: (item: any) => formatSmartDate(item.snapshot_timestamp) },
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
    name: "snapshot_timestamp",
    col: "right",
    label: "Snapshot Timestamp",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "readiness_category",
    col: "left",
    label: "Readiness Category",
    fieldType: "select",
    options: [
      { label: "C-1", value: "C-1" },
      { label: "C-2", value: "C-2" },
      { label: "C-3", value: "C-3" },
      { label: "C-4", value: "C-4" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "mro_readiness_score",
    col: "right",
    label: "Mro Readiness Score",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "personnel_manning_score",
    col: "left",
    label: "Personnel Manning Score",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "logistics_supply_score",
    col: "right",
    label: "Logistics Supply Score",
    fieldType: "number",
    required: true,
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
  delete payload.snapshot_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
