import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";

export const vendorTypeEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/vendor_type`;
export const vendorTypeTableName = "vendor_type";
export const vendorTypePrimaryKey = "vendor_type_id";

export const vendorTypeColumns: ColumnField[] = [
  {
    key: "vendor_type_name",
    label: "Jenis Vendor",
    sortable: "string",
    columnLength: 220,
  },
  {
    key: "vendor_type_description",
    label: "Deskripsi",
    sortable: "string",
    columnLength: 320,
  },
  {
    key: "is_active",
    label: "Aktif",
    render: (item) => (Number(item.is_active) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
];

export const vendorTypeFilters: FilterField[] = [
  {
    name: "vendor_type_name",
    label: "Jenis Vendor",
    fieldType: "text",
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

export function vendorTypeFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "vendor_type_name",
      label: "Jenis Vendor",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "vendor_type_description",
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

export const buildVendorTypePayload = (data: FormDataObject) =>
  pickAdminPayload(data, [
    "vendor_type_name",
    "vendor_type_description",
    "is_active",
  ]);
