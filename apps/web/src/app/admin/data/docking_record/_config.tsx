/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "docking_record";
export const entityTitle = "Riwayat Docking";
export const entityEndpoint = "/admin/docking_record";
export const primaryKey = "docking_id";

export const columns: ColumnField[] = [
  { key: "ship_name", label: "Kapal KRI" },
  { key: "shipyard_name", label: "Shipyard Name" },
  { key: "docking_type", label: "Docking Type" },
  { key: "entry_date", label: "Entry Date" },
  { key: "scheduled_exit_date", label: "Scheduled Exit Date" },
  { key: "actual_exit_date", label: "Actual Exit Date" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Riwayat Docking", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
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
    name: "shipyard_name",
    col: "right",
    label: "Shipyard Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "docking_type",
    col: "left",
    label: "Docking Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "entry_date",
    col: "right",
    label: "Entry Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "scheduled_exit_date",
    col: "left",
    label: "Scheduled Exit Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "actual_exit_date",
    col: "right",
    label: "Actual Exit Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "sea_trial_passed",
    col: "left",
    label: "Sea Trial Passed",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "classification_surveyor",
    col: "right",
    label: "Classification Surveyor",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "certificate_number",
    col: "left",
    label: "Certificate Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "total_docking_cost",
    col: "right",
    label: "Total Docking Cost",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "docking_summary",
    col: "left",
    label: "Docking Summary",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.docking_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
