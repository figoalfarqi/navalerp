import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";

export const cargoTypeEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/cargo_type`;
export const cargoTypeTableName = "cargo_type";
export const cargoTypePrimaryKey = "cargo_type_id";

export const cargoTypeColumns: ColumnField[] = [
  {
    key: "cargo_type_name",
    label: "Jenis Muatan",
    sortable: "string",
    columnLength: 220,
  },
  {
    key: "cargo_type_grade",
    label: "Grade",
    sortable: "string",
    columnLength: 160,
  },
  {
    key: "cargo_type_description",
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
];

export const cargoTypeFilters: FilterField[] = [
  {
    name: "cargo_type_name",
    label: "Jenis Muatan",
    fieldType: "text",
    col: "left",
  },
  {
    name: "cargo_type_grade",
    label: "Grade",
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

export function cargoTypeFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "cargo_type_name",
      label: "Jenis Muatan",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "cargo_type_grade",
      label: "Grade",
      fieldType: "text",
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "cargo_type_description",
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

export const buildCargoTypePayload = (data: FormDataObject) =>
  pickAdminPayload(data, [
    "cargo_type_name",
    "cargo_type_description",
    "cargo_type_grade",
    "is_active",
  ]);
