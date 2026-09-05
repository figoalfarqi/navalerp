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

export const truckAssignmentEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_truck_assignment`;
export const truckAssignmentTableName = "project_truck_assignment";
export const truckAssignmentPrimaryKey = "project_truck_assignment_id";

export const truckAssignmentColumns: ColumnField[] = [
  {
    key: "project_id",
    label: "Project",
    render: (item) => item.project?.project_name ?? "-",
    sortable: "table_key",
    columnLength: 230,
  },
  {
    key: "truck_id",
    label: "Truck",
    render: (item) => item.truck?.license_plate ?? "-",
    sortable: "table_key",
    columnLength: 160,
  },
  {
    key: "assignment_started_at",
    label: "Mulai Penugasan",
    render: (item) => formatDateTime(item.assignment_started_at),
    sortable: "date",
    columnLength: 180,
  },
  {
    key: "assignment_ended_at",
    label: "Akhir Penugasan",
    render: (item) =>
      item.assignment_ended_at
        ? formatDateTime(item.assignment_ended_at)
        : "-",
    sortable: "date",
    columnLength: 180,
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
    columnLength: 280,
  },
];

export const truckAssignmentFilters: FilterField[] = [
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
    name: "truck_id",
    label: "Truck",
    fieldType: "select",
    options: {
      url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck?limit=999`,
      labelKey: "license_plate",
      valueKey: "truck_id",
    },
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

export function truckAssignmentFields(mode: AdminCrudMode): FormField[] {
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
      name: "truck_id",
      label: "Truck",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck?limit=999`,
        labelKey: "license_plate",
        valueKey: "truck_id",
      },
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "assignment_started_at",
      label: "Mulai Penugasan",
      fieldType: "datetime",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "assignment_ended_at",
      label: "Akhir Penugasan",
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
  "truck_id",
  "assignment_started_at",
  "assignment_ended_at",
  "assignment_note",
  "is_active",
] as const;

export const buildTruckAssignmentPayload = (data: FormDataObject) =>
  pickAdminPayload(data, payloadFields);
