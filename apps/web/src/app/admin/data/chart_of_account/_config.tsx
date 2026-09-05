/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "chart_of_account";
export const entityTitle = "Bagan Akun (COA)";
export const entityEndpoint = "/admin/chart_of_account";
export const primaryKey = "account_id";

export const columns: ColumnField[] = [
  { key: "account_code", label: "Account Code" },
  { key: "account_name", label: "Account Name" },
  { key: "account_type", label: "Account Type" },
  { key: "parent_account_name", label: "Akun Induk" },
  { key: "is_active", label: "Is Active" },
  { key: "created_by", label: "Created By" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Bagan Akun (COA)", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "account_code",
    label: "Account Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "account_name",
    label: "Account Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "account_type",
    label: "Account Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "parent_account_id",
    label: "Akun Induk",
    fieldType: "select",
    options: {
      url: "/admin/chart_of_account?limit=100",
      labelKey: "account_name",
      valueKey: "account_id",
    },
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
  delete payload.account_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
