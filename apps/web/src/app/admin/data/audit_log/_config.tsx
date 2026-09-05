/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "audit_log";
export const entityTitle = "Audit Log";
export const entityEndpoint = "/admin/audit_log";
export const primaryKey = "log_id";

export const columns: ColumnField[] = [
  { key: "log_id", label: "Log Id" },
  { key: "user_id", label: "User Id" },
  { key: "action", label: "Action" },
  { key: "entity_table", label: "Entity Table" },
  { key: "entity_id", label: "Entity Id" },
  { key: "old_values", label: "Old Values" },
  { key: "new_values", label: "New Values" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Audit Log", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "user_id",
    label: "User Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "action",
    label: "Action",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "entity_table",
    label: "Entity Table",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "entity_id",
    label: "Entity Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "old_values",
    label: "Old Values",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "new_values",
    label: "New Values",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "ip_address",
    label: "Ip Address",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "user_agent",
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
