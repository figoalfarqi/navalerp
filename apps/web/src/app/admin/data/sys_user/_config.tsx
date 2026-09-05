/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "sys_user";
export const entityTitle = "Pengguna Sistem";
export const entityEndpoint = "/admin/sys_user";
export const primaryKey = "user_id";

export const columns: ColumnField[] = [
  { key: "user_id", label: "User Id" },
  { key: "unit_id", label: "Unit Id" },
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
    label: "Unit Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "username",
    label: "Username",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "password_hash",
    label: "Password",
    fieldType: "password",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "full_name",
    label: "Full Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "email",
    label: "Email",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "phone",
    label: "Phone",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "military_id",
    label: "Military Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "rank_title",
    label: "Rank Title",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "department",
    label: "Department",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "role",
    label: "Role",
    fieldType: "text",
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
  {
    name: "last_login_at",
    label: "Last Login At",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "failed_login_attempts",
    label: "Failed Login Attempts",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "auth_version",
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
