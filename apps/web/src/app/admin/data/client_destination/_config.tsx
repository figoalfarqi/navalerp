import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";

export const clientDestinationEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/client_destination`;
export const clientDestinationTableName = "client_destination";
export const clientDestinationPrimaryKey = "client_destination_id";

const clientOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/client?limit=999`,
  labelKey: "client_name",
  valueKey: "client_id",
};
const cityOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city?limit=999`,
  labelKey: "city_name",
  valueKey: "city_id",
};

export const clientDestinationColumns: ColumnField[] = [
  {
    key: "client_destination_name",
    label: "Lokasi Tujuan",
    sortable: "string",
    columnLength: 240,
  },
  {
    key: "client_id",
    label: "Client",
    render: (item) => item.client?.client_name ?? "-",
    sortable: "table_key",
    columnLength: 220,
  },
  {
    key: "city_id",
    label: "Kota",
    render: (item) => item.city?.city_name ?? "-",
    sortable: "table_key",
    columnLength: 180,
  },
  {
    key: "client_destination_address",
    label: "Alamat",
    sortable: "string",
    columnLength: 300,
  },
  {
    key: "client_destination_latitude",
    label: "Latitude",
    sortable: "number",
    columnLength: 130,
  },
  {
    key: "client_destination_longitude",
    label: "Longitude",
    sortable: "number",
    columnLength: 130,
  },
  {
    key: "operating_hours",
    label: "Jam Operasional",
    render: (item) =>
      item.operating_hours
        ? typeof item.operating_hours === "string"
          ? item.operating_hours
          : JSON.stringify(item.operating_hours)
        : "-",
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

export const clientDestinationFilters: FilterField[] = [
  {
    name: "client_destination_name",
    label: "Nama Tujuan",
    fieldType: "text",
    col: "left",
  },
  {
    name: "client_id",
    label: "Client",
    fieldType: "select",
    options: clientOptions,
    col: "right",
  },
  {
    name: "city_id",
    label: "Kota",
    fieldType: "select",
    options: cityOptions,
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

export function clientDestinationFields(
  mode: AdminCrudMode,
): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "client_id",
      label: "Client",
      fieldType: "select",
      options: clientOptions,
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "client_destination_name",
      label: "Nama Lokasi Tujuan",
      fieldType: "text",
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "city_id",
      label: "Kota",
      fieldType: "select",
      options: cityOptions,
      disabled,
      col: "left",
    },
    {
      name: "client_destination_address",
      label: "Alamat",
      fieldType: "textarea",
      disabled,
      col: "right",
    },
    {
      name: "client_destination_latitude",
      label: "Latitude",
      fieldType: "number",
      validation: { min: -90, max: 90 },
      disabled,
      col: "left",
    },
    {
      name: "client_destination_longitude",
      label: "Longitude",
      fieldType: "number",
      validation: { min: -180, max: 180 },
      disabled,
      col: "right",
    },
    {
      name: "client_destination_map_url",
      label: "Tautan Peta",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "operating_hours",
      label: "Jam Operasional (JSON)",
      fieldType: "textarea",
      placeholder: '{"monday":{"open":"08:00","close":"17:00"}}',
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

export function normalizeClientDestinationRecord(
  record: Record<string, unknown>,
) {
  return {
    ...record,
    operating_hours:
      typeof record.operating_hours === "string"
        ? record.operating_hours
        : record.operating_hours
          ? JSON.stringify(record.operating_hours, null, 2)
          : "",
  };
}

export function buildClientDestinationPayload(data: FormDataObject) {
  const payload = pickAdminPayload(data, [
    "client_id",
    "client_destination_name",
    "city_id",
    "client_destination_address",
    "client_destination_latitude",
    "client_destination_longitude",
    "client_destination_map_url",
    "operating_hours",
    "is_active",
  ]);
  if (typeof payload.operating_hours === "string") {
    try {
      payload.operating_hours = payload.operating_hours.trim()
        ? JSON.parse(payload.operating_hours)
        : null;
    } catch {
      // Keep the original value so the API can return a validation message.
    }
  }
  return payload;
}
