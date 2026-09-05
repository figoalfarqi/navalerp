/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "pm_schedule";
export const entityTitle = "Jadwal PMS";
export const entityEndpoint = "/admin/pm_schedule";
export const primaryKey = "pm_id";

export const columns: ColumnField[] = [
  { key: "equipment_name", label: "Peralatan" },
  { key: "pm_code", label: "Pm Code" },
  { key: "pm_title", label: "Pm Title" },
  { key: "interval_hours", label: "Interval Hours" },
  { key: "interval_days", label: "Interval Days" },
  { key: "last_performed_at", label: "Last Performed At" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Jadwal PMS", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
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
    name: "pm_code",
    label: "Pm Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "pm_title",
    label: "Pm Title",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "interval_hours",
    label: "Interval Hours",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "interval_days",
    label: "Interval Days",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "last_performed_at",
    label: "Last Performed At",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "next_due_at",
    label: "Next Due At",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "task_instructions",
    label: "Task Instructions",
    fieldType: "textarea",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "estimated_duration_hours",
    label: "Estimated Duration Hours",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_active",
    label: "Is Active",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.pm_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
