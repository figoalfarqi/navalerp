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

export const appSettingEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/app_setting`;
export const appSettingTableName = "app_setting";
export const appSettingPrimaryKey = "app_setting_id";

export const appSettingColumns: ColumnField[] = [
  {
    key: "app_setting_key",
    label: "Key",
    sortable: "string",
    columnLength: 260,
  },
  {
    key: "app_setting_value",
    label: "Value",
    render: (item) =>
      typeof item.app_setting_value === "string"
        ? item.app_setting_value
        : JSON.stringify(item.app_setting_value),
    columnLength: 320,
  },
  {
    key: "app_setting_description",
    label: "Deskripsi",
    sortable: "string",
    columnLength: 300,
  },
  {
    key: "is_active",
    label: "Aktif",
    render: (item) => (Number(item.is_active) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
  {
    key: "updated_at",
    label: "Diperbarui",
    render: (item) => formatDateTime(item.updated_at),
    sortable: "date",
    columnLength: 180,
  },
];

export const appSettingFilters: FilterField[] = [
  {
    name: "app_setting_key",
    label: "Key",
    fieldType: "text",
    col: "left",
  },
  {
    name: "app_setting_description",
    label: "Deskripsi",
    fieldType: "text",
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
    col: "left",
  },
];

export function appSettingFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "app_setting_key",
      label: "Setting Key",
      fieldType: "text",
      required: true,
      disabled,
      placeholder: "contoh: checker.location_cooldown_minutes",
      col: "left",
    },
    {
      name: "app_setting_value",
      label: "Setting Value (JSON)",
      fieldType: "textarea",
      required: true,
      disabled,
      placeholder: '{"minutes":5}',
      col: "right",
    },
    {
      name: "app_setting_description",
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

export function normalizeAppSettingRecord(record: Record<string, unknown>) {
  return {
    ...record,
    app_setting_value:
      typeof record.app_setting_value === "string"
        ? record.app_setting_value
        : JSON.stringify(record.app_setting_value ?? {}, null, 2),
  };
}

export function buildAppSettingPayload(data: FormDataObject) {
  const payload = pickAdminPayload(data, [
    "app_setting_key",
    "app_setting_value",
    "app_setting_description",
    "is_active",
  ]);
  if (typeof payload.app_setting_value === "string") {
    try {
      payload.app_setting_value = JSON.parse(payload.app_setting_value);
    } catch {
      // Backend returns the field-level validation message for invalid JSON.
    }
  }
  return payload;
}
