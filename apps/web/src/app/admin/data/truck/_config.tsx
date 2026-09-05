import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";
import { OwnershipStatusMap } from "@/consta/OwnershipStatusMap";

export const truckEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck`;
export const truckTableName = "truck";
export const truckPrimaryKey = "truck_id";

const truckTypeOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck_type?limit=999`,
  labelKey: "truck_type_name",
  valueKey: "truck_type_id",
};
const truckMerkOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck_merk?limit=999`,
  labelKey: "truck_merk_name",
  valueKey: "truck_merk_id",
};
const driverOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/driver?limit=999`,
  labelKey: "app_user_name",
  valueKey: "app_user_id",
};
const vendorOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/vendor?limit=999`,
  labelKey: "vendor_name",
  valueKey: "vendor_id",
};
const ownershipOptions = Object.entries(OwnershipStatusMap).map(
  ([value, label]) => ({ value: Number(value), label }),
);

export const truckColumns: ColumnField[] = [
  {
    key: "license_plate",
    label: "Plat Nomor",
    sortable: "string",
    columnLength: 150,
  },
  {
    key: "driver_id",
    label: "Driver",
    render: (item) => item.driver?.app_user_name ?? "-",
    sortable: "table_key",
    columnLength: 200,
  },
  {
    key: "truck_type_id",
    label: "Jenis Truck",
    render: (item) =>
      item.truck_type?.truck_type_name ?? "-",
    sortable: "table_key",
    columnLength: 180,
  },
  {
    key: "truck_merk_id",
    label: "Merk",
    render: (item) =>
      item.truck_merk?.truck_merk_name ?? "-",
    sortable: "table_key",
    columnLength: 160,
  },
  {
    key: "vendor_id",
    label: "Vendor",
    render: (item) => item.vendor?.vendor_name ?? "-",
    sortable: "table_key",
    columnLength: 200,
  },
  {
    key: "ownership_status_id",
    label: "Kepemilikan",
    render: (item) =>
      OwnershipStatusMap[Number(item.ownership_status_id)] ?? "-",
    sortable: "number",
    columnLength: 140,
  },
  {
    key: "production_year",
    label: "Tahun",
    sortable: "number",
    columnLength: 100,
  },
  {
    key: "number_of_tires",
    label: "Jumlah Ban",
    sortable: "number",
    columnLength: 110,
  },
  {
    key: "is_active",
    label: "Aktif",
    render: (item) => (Number(item.is_active) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
];

export const truckFilters: FilterField[] = [
  {
    name: "license_plate",
    label: "Plat Nomor",
    fieldType: "text",
    col: "left",
  },
  {
    name: "driver_id",
    label: "Driver",
    fieldType: "select",
    options: driverOptions,
    col: "right",
  },
  {
    name: "truck_type_id",
    label: "Jenis Truck",
    fieldType: "select",
    options: truckTypeOptions,
    col: "left",
  },
  {
    name: "truck_merk_id",
    label: "Merk",
    fieldType: "select",
    options: truckMerkOptions,
    col: "right",
  },
  {
    name: "vendor_id",
    label: "Vendor",
    fieldType: "select",
    options: vendorOptions,
    col: "left",
  },
  {
    name: "ownership_status_id",
    label: "Kepemilikan",
    fieldType: "select",
    options: ownershipOptions,
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

export function truckFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "license_plate",
      label: "Plat Nomor",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "driver_id",
      label: "Driver",
      fieldType: "select",
      options: driverOptions,
      disabled,
      col: "right",
    },
    {
      name: "truck_type_id",
      label: "Jenis Truck",
      fieldType: "select",
      options: truckTypeOptions,
      disabled,
      col: "left",
    },
    {
      name: "truck_merk_id",
      label: "Merk Truck",
      fieldType: "select",
      options: truckMerkOptions,
      disabled,
      col: "right",
    },
    {
      name: "vendor_id",
      label: "Vendor",
      fieldType: "select",
      options: vendorOptions,
      disabled,
      col: "left",
    },
    {
      name: "ownership_status_id",
      label: "Status Kepemilikan",
      fieldType: "select",
      options: ownershipOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "production_year",
      label: "Tahun Produksi",
      fieldType: "dateyear",
      disabled,
      col: "left",
    },
    {
      name: "number_of_tires",
      label: "Jumlah Ban",
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

export const buildTruckPayload = (data: FormDataObject) =>
  pickAdminPayload(data, [
    "truck_type_id",
    "truck_merk_id",
    "driver_id",
    "vendor_id",
    "license_plate",
    "ownership_status_id",
    "production_year",
    "number_of_tires",
    "is_active",
  ]);
