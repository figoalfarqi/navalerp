import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";

export const vendorEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/vendor`;
export const vendorTableName = "vendor";
export const vendorPrimaryKey = "vendor_id";

const vendorTypeOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/vendor_type?limit=999`,
  labelKey: "vendor_type_name",
  valueKey: "vendor_type_id",
};
const bankOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/bank_merk?limit=999`,
  labelKey: "bank_merk_name",
  valueKey: "bank_merk_id",
};
const cityOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city?limit=999`,
  labelKey: "city_name",
  valueKey: "city_id",
};

export const vendorColumns: ColumnField[] = [
  {
    key: "vendor_name",
    label: "Vendor",
    sortable: "string",
    columnLength: 220,
  },
  {
    key: "vendor_type_id",
    label: "Jenis",
    render: (item) =>
      item.vendor_type?.vendor_type_name ?? "-",
    sortable: "table_key",
    columnLength: 180,
  },
  {
    key: "vendor_email",
    label: "Email",
    sortable: "string",
    columnLength: 220,
  },
  {
    key: "vendor_phone",
    label: "Telepon",
    sortable: "string",
    columnLength: 150,
  },
  {
    key: "city_id",
    label: "Kota",
    render: (item) => item.city?.city_name ?? "-",
    sortable: "table_key",
    columnLength: 180,
  },
  {
    key: "bank_account_number",
    label: "Rekening",
    sortable: "string",
    columnLength: 180,
  },
  {
    key: "is_active",
    label: "Aktif",
    render: (item) => (Number(item.is_active) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
];

export const vendorFilters: FilterField[] = [
  {
    name: "vendor_name",
    label: "Nama Vendor",
    fieldType: "text",
    col: "left",
  },
  {
    name: "vendor_type_id",
    label: "Jenis Vendor",
    fieldType: "select",
    options: vendorTypeOptions,
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

export function vendorFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "vendor_name",
      label: "Nama Vendor",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "vendor_type_id",
      label: "Jenis Vendor",
      fieldType: "select",
      options: vendorTypeOptions,
      disabled,
      col: "right",
    },
    {
      name: "vendor_email",
      label: "Email",
      fieldType: "email",
      disabled,
      col: "left",
    },
    {
      name: "vendor_phone",
      label: "Telepon",
      fieldType: "text",
      disabled,
      col: "right",
    },
    {
      name: "vendor_tin",
      label: "NPWP / TIN",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "city_id",
      label: "Kota",
      fieldType: "select",
      options: cityOptions,
      disabled,
      col: "right",
    },
    {
      name: "vendor_address",
      label: "Alamat",
      fieldType: "textarea",
      disabled,
      col: "left",
    },
    {
      name: "bank_merk_id",
      label: "Bank",
      fieldType: "select",
      options: bankOptions,
      disabled,
      col: "right",
    },
    {
      name: "bank_account_number",
      label: "Nomor Rekening",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "bank_account_name",
      label: "Nama Rekening",
      fieldType: "text",
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

export const buildVendorPayload = (data: FormDataObject) =>
  pickAdminPayload(data, [
    "vendor_type_id",
    "bank_merk_id",
    "vendor_name",
    "vendor_email",
    "vendor_phone",
    "vendor_tin",
    "city_id",
    "vendor_address",
    "bank_account_number",
    "bank_account_name",
    "is_active",
  ]);
