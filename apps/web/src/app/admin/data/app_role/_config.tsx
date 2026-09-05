import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";

export const appRoleEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/app_role`;
export const appRoleTableName = "app_role";
export const appRolePrimaryKey = "app_role_id";

export const appRoleTypeOptions = [
  { label: "Driver", value: 1 },
  { label: "Checker", value: 2 },
  { label: "Manajemen / Admin", value: 3 },
];

const roleTypeLabel = (value: unknown) =>
  appRoleTypeOptions.find((option) => option.value === Number(value))?.label ??
  "-";

export const appRoleColumns: ColumnField[] = [
  {
    key: "app_role_name",
    label: "Role",
    sortable: "string",
    columnLength: 220,
  },
  {
    key: "app_role_type_id",
    label: "Kelompok Role",
    render: (item) => roleTypeLabel(item.app_role_type_id),
    sortable: "number",
    columnLength: 190,
  },
  {
    key: "app_role_description",
    label: "Deskripsi",
    sortable: "string",
    columnLength: 320,
  },
  {
    key: "is_active",
    label: "Aktif",
    render: (item) => (Number(item.is_active) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
];

export const appRoleFilters: FilterField[] = [
  {
    name: "app_role_name",
    label: "Nama Role",
    fieldType: "text",
    col: "left",
  },
  {
    name: "app_role_type_id",
    label: "Kelompok Role",
    fieldType: "select",
    options: appRoleTypeOptions,
    col: "right",
  },
  {
    name: "is_active",
    label: "Aktif",
    fieldType: "radio",
    options: [
      { label: "Aktif", value: "1", variant: "green-outline" },
      { label: "Tidak aktif", value: "0", variant: "red-outline" },
      { label: "Semua", value: "2", variant: "purple-outline" },
    ],
    col: "right",
  },
];

export function appRoleFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "app_role_name",
      label: "Nama Role",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "app_role_type_id",
      label: "Kelompok Role",
      fieldType: "select",
      options: appRoleTypeOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "app_role_description",
      label: "Deskripsi",
      fieldType: "textarea",
      disabled,
      col: "left",
    },
    {
      name: "is_active",
      label: "Aktif",
      fieldType: "boolean",
      required: true,
      disabled,
      col: "right",
    },
  ];
}

export const buildAppRolePayload = (data: FormDataObject) =>
  pickAdminPayload(data, [
    "app_role_type_id",
    "app_role_name",
    "app_role_description",
    "is_active",
  ]);
