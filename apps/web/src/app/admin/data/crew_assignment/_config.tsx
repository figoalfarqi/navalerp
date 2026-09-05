/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "crew_assignment";
export const entityTitle = "Awak KRI";
export const entityEndpoint = "/admin/crew_assignment";
export const primaryKey = "assignment_id";

export const columns: ColumnField[] = [
  { key: "ship_name", label: "Kapal KRI" },
  { key: "personnel_name", label: "Nama Personel" },
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
    name: "personnel_id",
    col: "right",
    label: "Nama Personel",
    fieldType: "select",
    options: {
      url: "/admin/personnel?limit=100",
      labelKey: "full_name",
      valueKey: "personnel_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "crew_role",
    col: "left",
    label: "Crew Role",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "department",
    col: "right",
    label: "Department",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "watch_bill_duty",
    col: "left",
    label: "Watch Bill Duty",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "assigned_date",
    col: "right",
    label: "Assigned Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "relieved_date",
    col: "left",
    label: "Relieved Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_active",
    col: "right",
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
