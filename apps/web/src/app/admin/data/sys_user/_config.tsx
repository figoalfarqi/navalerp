/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "sys_user";
export const entityTitle = "Pengguna Sistem";
export const entityEndpoint = "/admin/sys_user";
export const primaryKey = "user_id";

export const columns: ColumnField[] = [
  { key: "unit_name", label: "Satuan / Unit" },
  { key: "username", label: "Username" },
  { key: "full_name", label: "Full Name" },
  { key: "email", label: "Email" },
  { key: "phone", label: "Phone" },
  { key: "military_id", label: "Military Id" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Pengguna Sistem", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "unit_id",
    col: "left",
    label: "Satuan / Unit",
    fieldType: "select",
    options: {
      url: "/admin/org_unit?limit=100",
      labelKey: "unit_name",
      valueKey: "unit_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "username",
    col: "right",
    label: "Username",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "password_hash",
    col: "left",
    label: "Password",
    fieldType: "password",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "full_name",
    col: "right",
    label: "Full Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "email",
    col: "left",
    label: "Email",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "phone",
    col: "right",
    label: "Phone",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "military_id",
    col: "left",
    label: "Military Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "rank_title",
    col: "right",
    label: "Rank Title",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "department",
    col: "left",
    label: "Department",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "role",
    col: "right",
    label: "Role",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_active",
    col: "left",
    label: "Is Active",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "last_login_at",
    col: "right",
    label: "Last Login At",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "failed_login_attempts",
    col: "left",
    label: "Failed Login Attempts",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "auth_version",
    col: "right",
    label: "Auth Version",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.user_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
