import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";

export const vesselEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/vessel`;
export const vesselTableName = "vessel";
export const vesselPrimaryKey = "vessel_id";

const vendorOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/vendor?limit=999`,
  labelKey: "vendor_name",
  valueKey: "vendor_id",
};

export const vesselColumns: ColumnField[] = [
  {
    key: "vessel_name",
    label: "Kapal",
    sortable: "string",
    columnLength: 240,
  },
  {
    key: "vendor_id",
    label: "Vendor",
    render: (item) => item.vendor?.vendor_name ?? "-",
    sortable: "table_key",
    columnLength: 220,
  },
  {
    key: "imo_number",
    label: "Nomor IMO",
    sortable: "string",
    columnLength: 160,
  },
  {
    key: "registration_number",
    label: "Nomor Registrasi",
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

export const vesselFilters: FilterField[] = [
  {
    name: "vessel_name",
    label: "Nama Kapal",
    fieldType: "text",
    col: "left",
  },
  {
    name: "vendor_id",
    label: "Vendor",
    fieldType: "select",
    options: vendorOptions,
    col: "right",
  },
  {
    name: "imo_number",
    label: "Nomor IMO",
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

export function vesselFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "vessel_name",
      label: "Nama Kapal",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "vendor_id",
      label: "Vendor",
      fieldType: "select",
      options: vendorOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "imo_number",
      label: "Nomor IMO",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "registration_number",
      label: "Nomor Registrasi",
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

export const buildVesselPayload = (data: FormDataObject) =>
  pickAdminPayload(data, [
    "vendor_id",
    "vessel_name",
    "imo_number",
    "registration_number",
    "is_active",
  ]);
