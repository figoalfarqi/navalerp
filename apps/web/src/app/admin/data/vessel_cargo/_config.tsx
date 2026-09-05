import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";
import { formatDateTime } from "@/utils/dateTime";

export const vesselCargoEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/vessel_cargo`;
export const vesselCargoTableName = "vessel_cargo";
export const vesselCargoPrimaryKey = "vessel_cargo_id";

const vesselOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/vessel?limit=999`,
  labelKey: "vessel_name",
  valueKey: "vessel_id",
};
const portOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/port?limit=999`,
  labelKey: "port_name",
  valueKey: "port_id",
};
const cargoTypeOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/cargo_type?limit=999`,
  labelKey: "cargo_type_name",
  valueKey: "cargo_type_id",
};

export const vesselCargoColumns: ColumnField[] = [
  {
    key: "vessel_id",
    label: "Kapal",
    render: (item) => item.vessel?.vessel_name ?? "-",
    sortable: "table_key",
    columnLength: 220,
  },
  {
    key: "port_id",
    label: "Pelabuhan",
    render: (item) => item.port?.port_name ?? "-",
    sortable: "table_key",
    columnLength: 220,
  },
  {
    key: "cargo_type_id",
    label: "Muatan",
    render: (item) => item.cargo_type?.cargo_type_name ?? "-",
    sortable: "table_key",
    columnLength: 180,
  },
  {
    key: "voyage_number",
    label: "Voyage",
    sortable: "string",
    columnLength: 140,
  },
  {
    key: "bill_of_lading_number",
    label: "Bill of Lading",
    sortable: "string",
    columnLength: 180,
  },
  {
    key: "arrival_at",
    label: "Tiba",
    render: (item) => (item.arrival_at ? formatDateTime(item.arrival_at) : "-"),
    sortable: "date",
    columnLength: 180,
  },
  {
    key: "manifest_volume_cubic",
    label: "Manifest (m³)",
    sortable: "number",
    columnLength: 150,
  },
  {
    key: "manifest_weight_ton",
    label: "Manifest (ton)",
    sortable: "number",
    columnLength: 150,
  },
  {
    key: "is_active",
    label: "Aktif",
    render: (item) => (Number(item.is_active) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
];

export const vesselCargoFilters: FilterField[] = [
  {
    name: "vessel_id",
    label: "Kapal",
    fieldType: "select",
    options: vesselOptions,
    col: "left",
  },
  {
    name: "port_id",
    label: "Pelabuhan",
    fieldType: "select",
    options: portOptions,
    col: "right",
  },
  {
    name: "cargo_type_id",
    label: "Muatan",
    fieldType: "select",
    options: cargoTypeOptions,
    col: "left",
  },
  {
    name: "voyage_number",
    label: "Voyage",
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
    col: "right",
  },
];

export function vesselCargoFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "vessel_id",
      label: "Kapal",
      fieldType: "select",
      options: vesselOptions,
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "port_id",
      label: "Pelabuhan Bongkar",
      fieldType: "select",
      options: portOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "cargo_type_id",
      label: "Jenis Muatan",
      fieldType: "select",
      options: cargoTypeOptions,
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "voyage_number",
      label: "Nomor Voyage",
      fieldType: "text",
      disabled,
      col: "right",
    },
    {
      name: "bill_of_lading_number",
      label: "Nomor Bill of Lading",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "arrival_at",
      label: "Waktu Tiba",
      fieldType: "datetime",
      disabled,
      col: "right",
    },
    {
      name: "unloading_started_at",
      label: "Mulai Bongkar",
      fieldType: "datetime",
      disabled,
      col: "left",
    },
    {
      name: "unloading_completed_at",
      label: "Selesai Bongkar",
      fieldType: "datetime",
      disabled,
      col: "right",
    },
    {
      name: "manifest_volume_cubic",
      label: "Volume Manifest (m³)",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "left",
    },
    {
      name: "manifest_weight_ton",
      label: "Berat Manifest (ton)",
      fieldType: "number",
      validation: { min: 0 },
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

export const buildVesselCargoPayload = (data: FormDataObject) =>
  pickAdminPayload(data, [
    "vessel_id",
    "port_id",
    "cargo_type_id",
    "voyage_number",
    "bill_of_lading_number",
    "arrival_at",
    "unloading_started_at",
    "unloading_completed_at",
    "manifest_volume_cubic",
    "manifest_weight_ton",
    "is_active",
  ]);
