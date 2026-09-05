import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";

export const portEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/port`;
export const portTableName = "port";
export const portPrimaryKey = "port_id";

const cityOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city?limit=999`,
  labelKey: "city_name",
  valueKey: "city_id",
};

export const portColumns: ColumnField[] = [
  {
    key: "port_name",
    label: "Pelabuhan",
    sortable: "string",
    columnLength: 240,
  },
  {
    key: "city_id",
    label: "Kota",
    render: (item) => item.city?.city_name ?? "-",
    sortable: "table_key",
    columnLength: 180,
  },
  {
    key: "port_address",
    label: "Alamat",
    sortable: "string",
    columnLength: 300,
  },
  {
    key: "port_latitude",
    label: "Latitude",
    sortable: "number",
    columnLength: 130,
  },
  {
    key: "port_longitude",
    label: "Longitude",
    sortable: "number",
    columnLength: 130,
  },
  {
    key: "is_active",
    label: "Aktif",
    render: (item) => (Number(item.is_active) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
];

export const portFilters: FilterField[] = [
  {
    name: "port_name",
    label: "Pelabuhan",
    fieldType: "text",
    col: "left",
  },
  {
    name: "city_id",
    label: "Kota",
    fieldType: "select",
    options: cityOptions,
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

export function portFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "port_name",
      label: "Nama Pelabuhan",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "city_id",
      label: "Kota",
      fieldType: "select",
      options: cityOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "port_address",
      label: "Alamat",
      fieldType: "textarea",
      disabled,
      col: "left",
    },
    {
      name: "port_map_url",
      label: "Tautan Peta",
      fieldType: "text",
      disabled,
      col: "right",
    },
    {
      name: "port_latitude",
      label: "Latitude",
      fieldType: "number",
      validation: { min: -90, max: 90 },
      disabled,
      col: "left",
    },
    {
      name: "port_longitude",
      label: "Longitude",
      fieldType: "number",
      validation: { min: -180, max: 180 },
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

export const buildPortPayload = (data: FormDataObject) =>
  pickAdminPayload(data, [
    "port_name",
    "city_id",
    "port_address",
    "port_latitude",
    "port_longitude",
    "port_map_url",
    "is_active",
  ]);
