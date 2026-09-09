/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
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
  { key: "incident_date", label: "Incident Date", render: (item: any) => formatSmartDate(item.incident_date) },
  { key: "severity", label: "Severity" },
  { key: "failure_mode", label: "Failure Mode" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Laporan Kerusakan", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "equipment_id",
    col: "left",
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
    col: "right",
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
    col: "left",
    label: "Report Number",
    fieldType: "text",
    required: true,
    readOnly: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "REP",
  },
  {
    name: "incident_date",
    col: "right",
    label: "Incident Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "severity",
    col: "left",
    label: "Severity",
    fieldType: "select",
    options: [
      { label: "Cat1", value: "CAT1" },
      { label: "Cat2", value: "CAT2" },
      { label: "Cat3", value: "CAT3" },
      { label: "Cat4", value: "CAT4" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "failure_mode",
    col: "right",
    label: "Failure Mode",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "description",
    col: "left",
    label: "Description",
    fieldType: "textarea",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "operational_impact",
    col: "right",
    label: "Operational Impact",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "immediate_action_taken",
    col: "left",
    label: "Immediate Action Taken",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "right",
    label: "Status",
    fieldType: "select",
    options: [
      { label: "Pending", value: "PENDING" },
      { label: "Assessed", value: "ASSESSED" },
      { label: "Work Order Created", value: "WORK_ORDER_CREATED" },
      { label: "Closed", value: "CLOSED" },
    ],
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
