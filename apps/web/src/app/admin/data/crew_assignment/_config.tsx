/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "crew_assignment";
export const entityTitle = "Awak KRI";
export const entityEndpoint = "/admin/crew_assignment";
export const primaryKey = "assignment_id";

export const columns: ColumnField[] = [
  { key: "assignment_id", label: "Assignment Id" },
  { key: "ship_id", label: "Ship Id" },
  { key: "personnel_id", label: "Personnel Id" },
  { key: "crew_role", label: "Crew Role" },
  { key: "department", label: "Department" },
  { key: "watch_bill_duty", label: "Watch Bill Duty" },
  { key: "assigned_date", label: "Assigned Date" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Awak KRI", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
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
    name: "personnel_id",
    label: "Personnel Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "crew_role",
    label: "Crew Role",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "department",
    label: "Department",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "watch_bill_duty",
    label: "Watch Bill Duty",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "assigned_date",
    label: "Assigned Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "relieved_date",
    label: "Relieved Date",
    fieldType: "date",
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
  delete payload.assignment_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
