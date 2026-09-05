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
import { formatNumberID } from "@/utils/currencyFormater";
import { formatDateTime } from "@/utils/dateTime";
import {
  getTransportStatusLabel,
  transportStatusOptions,
} from "../project_transport/_labels";

export const projectTransportStatusEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_transport_status`;
export const projectTransportStatusTableName = "project_transport_status";
export const projectTransportStatusPrimaryKey =
  "project_transport_status_id";

export { transportStatusOptions };

export const projectTransportStatusColumns: ColumnField[] = [
  {
    key: "transport_number",
    label: "Transport",
    render: (item) => item.project_transport?.transport_number ?? "-",
    columnLength: 180,
  },
  {
    key: "project_name",
    label: "Project",
    render: (item) => item.project_transport?.project?.project_name ?? "-",
    columnLength: 220,
  },
  {
    key: "license_plate",
    label: "Truck",
    render: (item) => item.project_transport?.truck?.license_plate ?? "-",
    columnLength: 140,
  },
  {
    key: "project_transport_status_type_id",
    label: "Status",
    render: (item) =>
      getTransportStatusLabel(item.project_transport_status_type_id),
    sortable: "number",
    columnLength: 230,
  },
  {
    key: "status_time",
    label: "Waktu Status",
    render: (item) => formatDateTime(item.status_time),
    sortable: "date",
    columnLength: 180,
  },
  {
    key: "cargo_volume_cubic",
    label: "Volume",
    render: (item) =>
      item.cargo_volume_cubic == null
        ? "-"
        : `${formatNumberID(item.cargo_volume_cubic)} m³`,
    sortable: "number",
    columnLength: 140,
  },
  {
    key: "cargo_weight_ton",
    label: "Berat",
    render: (item) =>
      item.cargo_weight_ton == null
        ? "-"
        : `${formatNumberID(item.cargo_weight_ton)} ton`,
    sortable: "number",
    columnLength: 140,
  },
  {
    key: "is_fraud",
    label: "Fraud",
    render: (item) => (Number(item.is_fraud) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
  {
    key: "project_transport_status_note",
    label: "Catatan",
    sortable: "string",
    columnLength: 260,
  },
  {
    key: "is_active",
    label: "Aktif",
    render: (item) => (Number(item.is_active) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
];

export const projectTransportStatusFilters: FilterField[] = [
  {
    name: "project_transport_id",
    label: "Transport",
    fieldType: "select",
    options: {
      url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_transport?limit=999`,
      labelKey: "transport_number",
      valueKey: "project_transport_id",
    },
    col: "left",
  },
  {
    name: "project_transport_status_type_id",
    label: "Status",
    fieldType: "select",
    options: transportStatusOptions,
    col: "right",
  },
  {
    name: "status_time",
    label: "Waktu Status",
    fieldType: "dateAfterBefore",
    col: "left",
  },
  {
    name: "is_fraud",
    label: "Fraud",
    fieldType: "radio",
    options: [
      { label: "Ya", value: "1", variant: "red-outline" },
      { label: "Tidak", value: "0", variant: "green-outline" },
    ],
    col: "right",
  },
];

export function projectTransportStatusFields(
  mode: AdminCrudMode,
): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "project_transport_id",
      label: "Transport",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_transport?limit=999`,
        labelKey: "transport_number",
        valueKey: "project_transport_id",
      },
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "project_transport_status_type_id",
      label: "Status",
      fieldType: "select",
      options: transportStatusOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "status_time",
      label: "Waktu Status",
      fieldType: "datetime",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "cargo_box_length",
      label: "Panjang Box",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "left",
    },
    {
      name: "cargo_box_width",
      label: "Lebar Box",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "right",
    },
    {
      name: "cargo_box_height",
      label: "Tinggi Box",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "left",
    },
    {
      name: "cargo_volume_cubic",
      label: "Volume Cargo (m³)",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "right",
    },
    {
      name: "cargo_weight_ton",
      label: "Berat Cargo (ton)",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "left",
    },
    {
      name: "is_fraud",
      label: "Terindikasi Fraud",
      fieldType: "boolean",
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "project_transport_status_note",
      label: "Catatan Status / Fraud",
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
  "project_transport_id",
  "project_transport_status_type_id",
  "status_time",
  "cargo_box_length",
  "cargo_box_width",
  "cargo_box_height",
  "cargo_volume_cubic",
  "cargo_weight_ton",
  "project_transport_status_note",
  "is_fraud",
  "is_active",
] as const;

export const buildProjectTransportStatusPayload = (data: FormDataObject) =>
  pickAdminPayload(data, payloadFields);
