/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "work_order";
export const entityTitle = "Perintah Kerja MRO";
export const entityEndpoint = "/admin/work_order";
export const primaryKey = "work_order_id";

export const columns: ColumnField[] = [
  { key: "work_order_id", label: "Work Order Id" },
  { key: "failure_report_id", label: "Failure Report Id" },
  { key: "pm_schedule_id", label: "Pm Schedule Id" },
  { key: "equipment_id", label: "Equipment Id" },
  { key: "work_order_number", label: "Work Order Number" },
  { key: "work_order_type", label: "Work Order Type" },
  { key: "priority", label: "Priority" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Perintah Kerja MRO", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "failure_report_id",
    label: "Failure Report Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "pm_schedule_id",
    label: "Pm Schedule Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "equipment_id",
    label: "Equipment Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "work_order_number",
    label: "Work Order Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "work_order_type",
    label: "Work Order Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "priority",
    label: "Priority",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "scheduled_start_date",
    label: "Scheduled Start Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "scheduled_end_date",
    label: "Scheduled End Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_start_date",
    label: "Actual Start Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_end_date",
    label: "Actual End Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "lead_engineer_user_id",
    label: "Lead Engineer User Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "assigned_facility",
    label: "Assigned Facility",
    fieldType: "text",
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
  {
    name: "total_labor_hours",
    label: "Total Labor Hours",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "estimated_cost",
    label: "Estimated Cost",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_cost",
    label: "Actual Cost",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "completion_notes",
    label: "Completion Notes",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.work_order_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
