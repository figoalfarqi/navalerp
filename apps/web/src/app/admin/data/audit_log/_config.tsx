/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "audit_log";
export const entityTitle = "Audit Log";
export const entityEndpoint = "/admin/audit_log";
export const primaryKey = "log_id";

export const columns: ColumnField[] = [
  { key: "full_name", label: "Pengguna" },
  { key: "action", label: "Action" },
  { key: "entity_table", label: "Entity Table" },
  { key: "old_values", label: "Old Values" },
  { key: "new_values", label: "New Values" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Audit Log", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "user_id",
    col: "left",
    label: "Pengguna",
    fieldType: "select",
    options: {
      url: "/admin/sys_user?limit=100",
      labelKey: "full_name",
      valueKey: "user_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "action",
    col: "right",
    label: "Action",
    fieldType: "select",
    options: [
      { label: "Create", value: "CREATE" },
      { label: "Update", value: "UPDATE" },
      { label: "Delete", value: "DELETE" },
      { label: "Login", value: "LOGIN" },
      { label: "Logout", value: "LOGOUT" },
      { label: "Approval", value: "APPROVAL" },
      { label: "Export", value: "EXPORT" },
      { label: "Initialize", value: "INITIALIZE" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "entity_table",
    col: "left",
    label: "Entity Table",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "entity_id",
    col: "right",
    label: "Entity Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "old_values",
    col: "left",
    label: "Old Values",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "new_values",
    col: "right",
    label: "New Values",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "ip_address",
    col: "left",
    label: "Ip Address",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "user_agent",
    col: "right",
    label: "User Agent",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.log_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
