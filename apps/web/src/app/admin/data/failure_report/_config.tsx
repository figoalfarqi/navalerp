/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "failure_report";
export const entityTitle = "Laporan Kerusakan";
export const entityEndpoint = "/admin/failure_report";
export const primaryKey = "report_id";

export const columns: ColumnField[] = [
  { key: "equipment_name", label: "Peralatan" },
  { key: "reporter_name", label: "Pelapor" },
  { key: "report_number", label: "Report Number" },
  { key: "incident_date", label: "Incident Date" },
  { key: "severity", label: "Severity" },
  { key: "failure_mode", label: "Failure Mode" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Laporan Kerusakan", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "equipment_id",
    label: "Peralatan",
    fieldType: "select",
    options: {
      url: "/admin/equipment?limit=100",
      labelKey: "equipment_name",
      valueKey: "equipment_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "reported_by_user_id",
    label: "Pelapor",
    fieldType: "select",
    options: {
      url: "/admin/sys_user?limit=100",
      labelKey: "full_name",
      valueKey: "user_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "report_number",
    label: "Report Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "incident_date",
    label: "Incident Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "severity",
    label: "Severity",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "failure_mode",
    label: "Failure Mode",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "description",
    label: "Description",
    fieldType: "textarea",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "operational_impact",
    label: "Operational Impact",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "immediate_action_taken",
    label: "Immediate Action Taken",
    fieldType: "textarea",
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
  delete payload.report_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
