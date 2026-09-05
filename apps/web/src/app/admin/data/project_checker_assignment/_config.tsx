import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import {
  pickAdminPayload,
} from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";
import { formatDateTime } from "@/utils/dateTime";

export const checkerAssignmentEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_checker_assignment`;
export const checkerAssignmentTableName = "project_checker_assignment";
export const checkerAssignmentPrimaryKey = "project_checker_assignment_id";

export const checkerAssignmentColumns: ColumnField[] = [
  {
    key: "project_id",
    label: "Project",
    render: (item) => item.project?.project_name ?? "-",
    sortable: "table_key",
    columnLength: 230,
  },
  {
    key: "checker_id",
    label: "Checker",
    render: (item) => item.checker?.app_user_name ?? "-",
    sortable: "table_key",
    columnLength: 200,
  },
  {
    key: "access_started_at",
    label: "Mulai Akses",
    render: (item) => formatDateTime(item.access_started_at),
    sortable: "date",
    columnLength: 180,
  },
  {
    key: "access_ended_at",
    label: "Akhir Akses",
    render: (item) =>
      item.access_ended_at ? formatDateTime(item.access_ended_at) : "-",
    sortable: "date",
    columnLength: 180,
  },
  {
    key: "is_default",
    label: "Default",
    render: (item) => (Number(item.is_default) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
  {
    key: "is_active",
    label: "Aktif",
    render: (item) => (Number(item.is_active) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
  {
    key: "assignment_note",
    label: "Catatan",
    sortable: "string",
    columnLength: 260,
  },
];

export const checkerAssignmentFilters: FilterField[] = [
  {
    name: "project_id",
    label: "Project",
    fieldType: "select",
    options: {
      url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project?limit=999`,
      labelKey: "project_name",
      valueKey: "project_id",
    },
    col: "left",
  },
  {
    name: "checker_id",
    label: "Checker",
    fieldType: "select",
    options: {
      url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/checker?limit=999`,
      labelKey: "app_user_name",
      valueKey: "app_user_id",
    },
    col: "right",
  },
  {
    name: "is_default",
    label: "Project Default",
    fieldType: "radio",
    options: [
      { label: "Ya", value: "1", variant: "green-outline" },
      { label: "Tidak", value: "0", variant: "gray-outline" },
    ],
    col: "left",
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

export function checkerAssignmentFields(
  mode: AdminCrudMode,
): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "project_id",
      label: "Project",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project?limit=999`,
        labelKey: "project_name",
        valueKey: "project_id",
      },
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "checker_id",
      label: "Checker",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/checker?limit=999`,
        labelKey: "app_user_name",
        valueKey: "app_user_id",
      },
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "access_started_at",
      label: "Mulai Akses",
      fieldType: "datetime",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "access_ended_at",
      label: "Akhir Akses",
      fieldType: "datetime",
      disabled,
      col: "right",
    },
    {
      name: "assignment_note",
      label: "Catatan Penugasan",
      fieldType: "textarea",
      disabled,
      col: "left",
    },
    {
      name: "is_default",
      label: "Project Default Checker",
      fieldType: "boolean",
      required: true,
      disabled,
      col: "right",
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

const payloadFields = [
  "project_id",
  "checker_id",
  "access_started_at",
  "access_ended_at",
  "assignment_note",
  "is_active",
  "is_default",
] as const;

export const buildCheckerAssignmentPayload = (data: FormDataObject) =>
  pickAdminPayload(data, payloadFields);
