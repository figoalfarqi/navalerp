import {
  pickAdminPayload,
  type AdminCrudMode,
} from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";
import { formatDateTime } from "@/utils/dateTime";

export const appUserEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/app_user`;
export const appUserTableName = "app_user";
export const appUserPrimaryKey = "app_user_id";

const userStatusOptions = [
  { label: "Pending", value: 0 },
  { label: "Aktif", value: 1 },
  { label: "Tidak aktif", value: 2 },
];

const fixedRoleLabels: Record<number, string> = {
  1: "Driver",
  2: "Checker",
  3: "Owner",
  4: "IT Developer",
  5: "Super Admin",
  6: "Admin",
};

function roleOptions(
  roleTypeId?: 1 | 2 | 3,
  excludedRoleIds: number[] = [],
) {
  const suffix = roleTypeId
    ? `&query=app_role_type_id%3D${roleTypeId}`
    : "";
  return {
    url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/app_role?limit=999${suffix}`,
    labelKey: "app_role_name",
    valueKey: "app_role_id",
    excludedValues: excludedRoleIds,
  };
}

export const appUserColumns: ColumnField[] = [
  {
    key: "app_user_name",
    label: "Nama",
    sortable: "string",
    columnLength: 220,
  },
  {
    key: "app_role_id",
    label: "Role",
    render: (item) => item.app_role?.app_role_name ?? "-",
    sortable: "table_key",
    columnLength: 160,
  },
  {
    key: "username",
    label: "Username",
    sortable: "string",
    columnLength: 180,
  },
  {
    key: "app_user_phone",
    label: "Telepon",
    sortable: "string",
    columnLength: 160,
  },
  {
    key: "city_id",
    label: "Kota",
    render: (item) => item.city?.city_name ?? "-",
    sortable: "table_key",
    columnLength: 170,
  },
  {
    key: "app_user_status_id",
    label: "Status",
    render: (item) =>
      userStatusOptions.find(
        (option) => option.value === Number(item.app_user_status_id),
      )?.label ?? "-",
    sortable: "number",
    columnLength: 120,
  },
  {
    key: "updated_at",
    label: "Diperbarui",
    render: (item) => formatDateTime(item.updated_at),
    sortable: "date",
    columnLength: 180,
  },
];

export const appUserFilters: FilterField[] = [
  {
    name: "app_user_name",
    label: "Nama",
    fieldType: "text",
    col: "left",
  },
  {
    name: "username",
    label: "Username",
    fieldType: "text",
    col: "right",
  },
  {
    name: "app_user_phone",
    label: "Telepon",
    fieldType: "text",
    col: "left",
  },
  {
    name: "app_role_id",
    label: "Role",
    fieldType: "select",
    options: roleOptions(),
    col: "right",
  },
  {
    name: "app_user_status_id",
    label: "Status",
    fieldType: "radio",
    options: [
      { label: "Aktif", value: "1", variant: "green-outline" },
      { label: "Pending", value: "0", variant: "gray-outline" },
      { label: "Tidak aktif", value: "2", variant: "red-outline" },
      { label: "Semua", value: "99", variant: "purple-outline" },
    ],
    col: "left",
  },
];

export function appUserFields(
  mode: AdminCrudMode,
  roleTypeId?: 1 | 2 | 3,
  fixedRoleId?: number,
  excludedRoleIds: number[] = [],
): FormField[] {
  const disabled = mode === "view";
  const passwordRequired = mode === "add" || mode === "copy";
  return [
    {
      name: "app_role_id",
      label: "Role",
      fieldType: "select",
      options:
        fixedRoleId == null
          ? roleOptions(roleTypeId, excludedRoleIds)
          : [
              {
                label: fixedRoleLabels[fixedRoleId] ?? `Role ${fixedRoleId}`,
                value: fixedRoleId,
              },
            ],
      required: true,
      disabled: disabled || fixedRoleId != null,
      col: "left",
    },
    {
      name: "app_user_status_id",
      label: "Status User",
      fieldType: "select",
      options: userStatusOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "username",
      label: "Username",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "password",
      label:
        mode === "edit"
          ? "Password Baru (kosongkan jika tidak berubah)"
          : "Password",
      fieldType: "password",
      required: passwordRequired,
      disabled,
      col: "right",
    },
    {
      name: "app_user_name",
      label: "Nama Lengkap",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "app_user_preferred_name",
      label: "Nama Panggilan",
      fieldType: "text",
      disabled,
      col: "right",
    },
    {
      name: "app_user_phone",
      label: "Nomor Telepon",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "city_id",
      label: "Kota",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city?limit=999`,
        labelKey: "city_name",
        valueKey: "city_id",
      },
      disabled,
      col: "right",
    },
    {
      name: "client_id",
      label: "Client Terkait",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/client?limit=999`,
        labelKey: "client_name",
        valueKey: "client_id",
      },
      disabled,
      col: "left",
    },
    {
      name: "app_user_address",
      label: "Alamat",
      fieldType: "textarea",
      disabled,
      col: "right",
    },
    {
      name: "app_user_photo_url",
      label: "Foto User",
      fieldType: "uploadimage",
      fileFolder: "/app-user",
      disabled,
      col: "left",
    },
    {
      name: "id_card_number",
      label: "Nomor KTP",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "id_card_photo_url",
      label: "Foto KTP",
      fieldType: "uploadimage",
      fileFolder: "/app-user/document",
      disabled,
      col: "right",
    },
    {
      name: "family_card_number",
      label: "Nomor Kartu Keluarga",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "family_card_photo_url",
      label: "Foto Kartu Keluarga",
      fieldType: "uploadimage",
      fileFolder: "/app-user/document",
      disabled,
      col: "right",
    },
    {
      name: "driver_license_b_number",
      label: "Nomor SIM B",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "driver_license_b_expiry",
      label: "Masa Berlaku SIM B",
      fieldType: "date",
      disabled,
      col: "right",
    },
    {
      name: "driver_license_b_photo_url",
      label: "Foto SIM B",
      fieldType: "uploadimage",
      fileFolder: "/app-user/document",
      disabled,
      col: "left",
    },
    {
      name: "tax_id_number",
      label: "Nomor NPWP",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "tax_id_photo_url",
      label: "Foto NPWP",
      fieldType: "uploadimage",
      fileFolder: "/app-user/document",
      disabled,
      col: "right",
    },
    {
      name: "bpjs_number",
      label: "Nomor BPJS",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "bpjs_photo_url",
      label: "Foto BPJS",
      fieldType: "uploadimage",
      fileFolder: "/app-user/document",
      disabled,
      col: "right",
    },
    {
      name: "bank_merk_id",
      label: "Bank",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/bank_merk?limit=999`,
        labelKey: "bank_merk_name",
        valueKey: "bank_merk_id",
      },
      disabled,
      col: "left",
    },
    {
      name: "bank_account_number",
      label: "Nomor Rekening",
      fieldType: "text",
      disabled,
      col: "right",
    },
    {
      name: "bank_account_name",
      label: "Nama Pemilik Rekening",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "salary_percentage",
      label: "Persentase Gaji",
      fieldType: "number",
      validation: { min: 0, max: 100 },
      disabled,
      col: "right",
    },
  ];
}

const payloadFields = [
  "app_role_id",
  "bank_merk_id",
  "client_id",
  "username",
  "password",
  "app_user_status_id",
  "app_user_name",
  "app_user_preferred_name",
  "app_user_phone",
  "city_id",
  "app_user_address",
  "app_user_photo_url",
  "id_card_photo_url",
  "id_card_number",
  "family_card_photo_url",
  "family_card_number",
  "driver_license_b_photo_url",
  "driver_license_b_number",
  "driver_license_b_expiry",
  "tax_id_photo_url",
  "tax_id_number",
  "bpjs_photo_url",
  "bpjs_number",
  "bank_account_number",
  "bank_account_name",
  "salary_percentage",
] as const;

export function buildAppUserPayload(data: FormDataObject) {
  const payload = pickAdminPayload(data, payloadFields);
  if (!data.password) delete payload.password;
  return payload;
}

export const appUserInitialData = { app_user_status_id: 1 };
