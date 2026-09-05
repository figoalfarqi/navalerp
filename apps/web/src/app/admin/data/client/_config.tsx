import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";

export const clientEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/client`;
export const clientTableName = "client";
export const clientPrimaryKey = "client_id";

const cityOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city?limit=999`,
  labelKey: "city_name",
  valueKey: "city_id",
};

export const clientColumns: ColumnField[] = [
  {
    key: "client_name",
    label: "Client",
    sortable: "string",
    columnLength: 220,
  },
  {
    key: "client_email",
    label: "Email",
    sortable: "string",
    columnLength: 220,
  },
  {
    key: "client_tin",
    label: "NPWP / TIN",
    sortable: "string",
    columnLength: 180,
  },
  {
    key: "number_of_day_until_due",
    label: "Jatuh Tempo (hari)",
    sortable: "number",
    columnLength: 160,
  },
  {
    key: "city_id",
    label: "Kota",
    render: (item) => item.city?.city_name ?? "-",
    sortable: "table_key",
    columnLength: 180,
  },
  {
    key: "client_address",
    label: "Alamat",
    sortable: "string",
    columnLength: 300,
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

export const clientFilters: FilterField[] = [
  {
    name: "client_name",
    label: "Nama Client",
    fieldType: "text",
    col: "left",
  },
  {
    name: "client_email",
    label: "Email",
    fieldType: "text",
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

export function clientFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "client_name",
      label: "Nama Client",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "client_email",
      label: "Email",
      fieldType: "email",
      disabled,
      col: "right",
    },
    {
      name: "client_tin",
      label: "NPWP / TIN",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "number_of_day_until_due",
      label: "Hari sampai Jatuh Tempo",
      fieldType: "number",
      validation: { min: 0 },
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
      name: "client_address",
      label: "Alamat",
      fieldType: "textarea",
      disabled,
      col: "right",
    },
    {
      name: "operating_hours",
      label: "Jam Operasional (JSON)",
      fieldType: "textarea",
      placeholder: '{"monday":{"open":"08:00","close":"17:00"}}',
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

export function normalizeClientRecord(record: Record<string, unknown>) {
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

export function buildClientPayload(data: FormDataObject) {
  const payload = pickAdminPayload(data, [
    "client_name",
    "client_email",
    "client_tin",
    "number_of_day_until_due",
    "city_id",
    "client_address",
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
