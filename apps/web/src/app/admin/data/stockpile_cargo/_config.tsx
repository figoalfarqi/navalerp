import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";
import { formatNumberID } from "@/utils/currencyFormater";

export const stockpileCargoEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/stockpile_cargo`;
export const stockpileCargoTableName = "stockpile_cargo";
export const stockpileCargoPrimaryKey = "stockpile_cargo_id";

const stockpileOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/stockpile?limit=999`,
  labelKey: "stockpile_name",
  valueKey: "stockpile_id",
};
const cargoOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/cargo_type?limit=999`,
  labelKey: "cargo_type_name",
  valueKey: "cargo_type_id",
};

const quantity = (value: unknown, unit: string) =>
  value == null || value === ""
    ? "-"
    : `${formatNumberID(Number(value))} ${unit}`;

export const stockpileCargoColumns: ColumnField[] = [
  {
    key: "stockpile_id",
    label: "Stockpile",
    render: (item) => item.stockpile?.stockpile_name ?? "-",
    sortable: "table_key",
    columnLength: 220,
  },
  {
    key: "cargo_type_id",
    label: "Muatan",
    render: (item) => item.cargo_type?.cargo_type_name ?? "-",
    sortable: "table_key",
    columnLength: 190,
  },
  {
    key: "capacity_volume_cubic",
    label: "Kapasitas Volume",
    render: (item) => quantity(item.capacity_volume_cubic, "m³"),
    sortable: "number",
    columnLength: 170,
  },
  {
    key: "capacity_weight_ton",
    label: "Kapasitas Berat",
    render: (item) => quantity(item.capacity_weight_ton, "ton"),
    sortable: "number",
    columnLength: 170,
  },
  {
    key: "current_volume_cubic",
    label: "Volume Saat Ini",
    render: (item) => quantity(item.current_volume_cubic, "m³"),
    sortable: "number",
    columnLength: 160,
  },
  {
    key: "current_weight_ton",
    label: "Berat Saat Ini",
    render: (item) => quantity(item.current_weight_ton, "ton"),
    sortable: "number",
    columnLength: 160,
  },
  {
    key: "is_active",
    label: "Aktif",
    render: (item) => (Number(item.is_active) === 1 ? "Ya" : "Tidak"),
    sortable: "boolean",
    columnLength: 100,
  },
];

export const stockpileCargoFilters: FilterField[] = [
  {
    name: "stockpile_id",
    label: "Stockpile",
    fieldType: "select",
    options: stockpileOptions,
    col: "left",
  },
  {
    name: "cargo_type_id",
    label: "Jenis Muatan",
    fieldType: "select",
    options: cargoOptions,
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

export function stockpileCargoFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "stockpile_id",
      label: "Stockpile",
      fieldType: "select",
      options: stockpileOptions,
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "cargo_type_id",
      label: "Jenis Muatan",
      fieldType: "select",
      options: cargoOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "capacity_volume_cubic",
      label: "Kapasitas Volume (m³)",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "left",
    },
    {
      name: "capacity_weight_ton",
      label: "Kapasitas Berat (ton)",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "right",
    },
    {
      name: "current_volume_cubic",
      label: "Volume Saat Ini (m³)",
      fieldType: "number",
      validation: { min: 0 },
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "current_weight_ton",
      label: "Berat Saat Ini (ton)",
      fieldType: "number",
      validation: { min: 0 },
      required: true,
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

export const buildStockpileCargoPayload = (data: FormDataObject) =>
  pickAdminPayload(data, [
    "stockpile_id",
    "cargo_type_id",
    "capacity_volume_cubic",
    "capacity_weight_ton",
    "current_volume_cubic",
    "current_weight_ton",
    "is_active",
  ]);
