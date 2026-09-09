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
  { key: "military_id", label: "NRP / Identitas Militer" },
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
    fieldType: "select",
    options: [
      { label: "Laksamana TNI", value: "LAKSAMANA_TNI" },
      { label: "Laksamana Madya TNI", value: "LAKSAMANA_MADYA_TNI" },
      { label: "Laksamana Muda TNI", value: "LAKSAMANA_MUDA_TNI" },
      { label: "Laksamana Pertama TNI", value: "LAKSAMANA_PERTAMA_TNI" },
      { label: "Kolonel Laut", value: "KOLONEL_LAUT" },
      { label: "Letkol Laut", value: "LETKOL_LAUT" },
      { label: "Mayor Laut", value: "MAYOR_LAUT" },
      { label: "Kapten Laut", value: "KAPTEN_LAUT" },
      { label: "Lettu Laut", value: "LETTU_LAUT" },
      { label: "Letda Laut", value: "LETDA_LAUT" },
      { label: "PELTU", value: "PELTU" },
      { label: "PELDA", value: "PELDA" },
      { label: "SERMA", value: "SERMA" },
      { label: "SERKA", value: "SERKA" },
      { label: "SERTU", value: "SERTU" },
      { label: "SERDA", value: "SERDA" },
      { label: "KOPKA", value: "KOPKA" },
      { label: "KOPTU", value: "KOPTU" },
      { label: "KOPDA", value: "KOPDA" },
      { label: "KLK", value: "KLK" },
      { label: "KLS", value: "KLS" },
      { label: "KLD", value: "KLD" },
      { label: "PNS TNI", value: "PNS_TNI" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "department",
    col: "left",
    label: "Department",
    fieldType: "select",
    options: [
      { label: "SRENA", value: "SRENA" },
      { label: "SOPS", value: "SOPS" },
      { label: "SLOG", value: "SLOG" },
      { label: "SPERS", value: "SPERS" },
      { label: "SPOTMAR", value: "SPOTMAR" },
      { label: "Komando", value: "KOMANDO" },
      { label: "DEPOPS", value: "DEPOPS" },
      { label: "DEPSIN", value: "DEPSIN" },
      { label: "DEPLOG", value: "DEPLOG" },
      { label: "DEPSENAU", value: "DEPSENAU" },
      { label: "FASHARKAN", value: "FASHARKAN" },
      { label: "DISBEKAL", value: "DISBEKAL" },
      { label: "DISLAIKMATAL", value: "DISLAIKMATAL" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "role",
    col: "right",
    label: "Role",
    fieldType: "select",
    options: [
      { label: "Super Admin", value: "SUPER_ADMIN" },
      { label: "Command Officer", value: "COMMAND_OFFICER" },
      { label: "KRI Commander", value: "KRI_COMMANDER" },
      { label: "Logistics Officer", value: "LOGISTICS_OFFICER" },
      { label: "Maintenance Officer", value: "MAINTENANCE_OFFICER" },
      { label: "Personnel Officer", value: "PERSONNEL_OFFICER" },
      { label: "Finance Officer", value: "FINANCE_OFFICER" },
      { label: "Operator", value: "OPERATOR" },
    ],
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
