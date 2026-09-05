import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";
import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import { formatCurrencyIDR, formatNumberID } from "@/utils/currencyFormater";
import { formatDate } from "@/utils/dateTime";

export const projectEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project`;
export const projectTableName = "project";
export const projectPrimaryKey = "project_id";

const routeTypeOptions = [
  { label: "Mine → Client", value: "MINE_CLIENT" },
  {
    label: "Mine → Stockpile → Client",
    value: "MINE_STOCKPILE_CLIENT",
  },
  { label: "Vessel → Client", value: "VESSEL_CLIENT" },
  {
    label: "Vessel → Stockpile → Client",
    value: "VESSEL_STOCKPILE_CLIENT",
  },
];

const projectStatusOptions = [
  { label: "Draft", value: "DRAFT" },
  { label: "Aktif", value: "ACTIVE" },
  { label: "Selesai", value: "COMPLETED" },
  { label: "Dibatalkan", value: "CANCELLED" },
];

const pricingUnitOptions = [
  { label: "Tidak digunakan", value: "NONE" },
  { label: "Per m³", value: "M3" },
  { label: "Per ton", value: "TON" },
];

export const projectColumns: ColumnField[] = [
  {
    key: "project_code",
    label: "Kode",
    sortable: "string",
    columnLength: 140,
  },
  {
    key: "project_name",
    label: "Nama Project",
    sortable: "string",
    columnLength: 240,
  },
  {
    key: "route_type",
    label: "Pola Rute",
    sortable: "string",
    columnLength: 220,
  },
  {
    key: "project_status",
    label: "Status",
    sortable: "string",
    columnLength: 130,
  },
  {
    key: "client_destination_id",
    label: "Tujuan",
    render: (item) =>
      item.client_destination?.client_destination_name ?? "-",
    sortable: "table_key",
    columnLength: 220,
  },
  {
    key: "planned_volume_cubic",
    label: "Rencana Volume",
    render: (item) =>
      item.planned_volume_cubic == null
        ? "-"
        : `${formatNumberID(item.planned_volume_cubic)} m³`,
    sortable: "number",
    columnLength: 160,
  },
  {
    key: "planned_weight_ton",
    label: "Rencana Berat",
    render: (item) =>
      item.planned_weight_ton == null
        ? "-"
        : `${formatNumberID(item.planned_weight_ton)} ton`,
    sortable: "number",
    columnLength: 160,
  },
  {
    key: "fixed_other_income",
    label: "Pendapatan Tetap",
    render: (item) => formatCurrencyIDR(item.fixed_other_income ?? 0),
    sortable: "number",
    columnLength: 180,
  },
  {
    key: "start_date",
    label: "Mulai",
    render: (item) => (item.start_date ? formatDate(item.start_date) : "-"),
    sortable: "date",
    columnLength: 130,
  },
  {
    key: "end_date",
    label: "Selesai",
    render: (item) => (item.end_date ? formatDate(item.end_date) : "-"),
    sortable: "date",
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

export const projectFilters: FilterField[] = [
  {
    name: "project_code",
    label: "Kode Project",
    fieldType: "text",
    col: "left",
  },
  {
    name: "project_name",
    label: "Nama Project",
    fieldType: "text",
    col: "right",
  },
  {
    name: "route_type",
    label: "Pola Rute",
    fieldType: "select",
    options: routeTypeOptions,
    col: "left",
  },
  {
    name: "project_status",
    label: "Status",
    fieldType: "select",
    options: projectStatusOptions,
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
    col: "left",
  },
];

export function projectFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "project_code",
      label: "Kode Project",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "project_name",
      label: "Nama Project",
      fieldType: "text",
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "route_type",
      label: "Pola Rute",
      fieldType: "select",
      options: routeTypeOptions,
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "project_status",
      label: "Status Project",
      fieldType: "select",
      options: projectStatusOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "mine_id",
      label: "Tambang (untuk pola Mine)",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/mine?limit=999`,
        labelKey: "mine_name",
        valueKey: "mine_id",
      },
      disabled,
      col: "left",
    },
    {
      name: "vessel_cargo_id",
      label: "Vessel Cargo (untuk pola Vessel)",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/vessel_cargo?limit=999`,
        labelKey: "voyage_number",
        valueKey: "vessel_cargo_id",
      },
      disabled,
      col: "right",
    },
    {
      name: "stockpile_cargo_id",
      label: "Stockpile Cargo (jika transit)",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/stockpile_cargo?limit=999&is_active=1`,
        labelKey:
          "{stockpile.stockpile_name} - {cargo_type.cargo_type_name}",
        valueKey: "stockpile_cargo_id",
      },
      disabled,
      col: "left",
    },
    {
      name: "client_destination_id",
      label: "Tujuan Client",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/client_destination?limit=999`,
        labelKey: "client_destination_name",
        valueKey: "client_destination_id",
      },
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "cargo_type_id",
      label: "Cargo Type",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/cargo_type?limit=999`,
        labelKey: "cargo_type_name",
        valueKey: "cargo_type_id",
      },
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "start_date",
      label: "Tanggal Mulai",
      fieldType: "date",
      disabled,
      col: "left",
    },
    {
      name: "end_date",
      label: "Tanggal Selesai",
      fieldType: "date",
      disabled,
      col: "right",
    },
    {
      name: "planned_volume_cubic",
      label: "Rencana Volume (m³)",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "left",
    },
    {
      name: "planned_weight_ton",
      label: "Rencana Berat (ton)",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "right",
    },
    {
      name: "volume_to_weight_conversion",
      label: "Konversi Volume ke Berat",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "left",
    },
    {
      name: "material_purchase_unit",
      label: "Unit Pembelian Material",
      fieldType: "select",
      options: pricingUnitOptions,
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "material_buy_price_per_cubic",
      label: "Harga Beli per m³",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "left",
    },
    {
      name: "material_buy_price_per_ton",
      label: "Harga Beli per ton",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "right",
    },
    {
      name: "material_sale_unit",
      label: "Unit Penjualan Material",
      fieldType: "select",
      options: pricingUnitOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "material_sell_price_per_cubic",
      label: "Harga Jual per m³",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "left",
    },
    {
      name: "material_sell_price_per_ton",
      label: "Harga Jual per ton",
      fieldType: "number",
      validation: { min: 0 },
      disabled,
      col: "right",
    },
    {
      name: "fixed_other_income",
      label: "Pendapatan Tetap Lain",
      fieldType: "number",
      validation: { min: 0 },
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "fixed_other_expense",
      label: "Pengeluaran Tetap Lain",
      fieldType: "number",
      validation: { min: 0 },
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "project_note",
      label: "Catatan Project",
      fieldType: "textarea",
      disabled,
      col: "left",
    },
    ...(mode === "edit" || mode === "view"
      ? [
          {
            name: "is_active",
            label: "Aktif",
            fieldType: "boolean" as const,
            required: true,
            disabled,
            col: "right" as const,
          },
        ]
      : []),
  ];
}

const projectPayloadFields = [
  "project_code",
  "project_name",
  "route_type",
  "mine_id",
  "vessel_cargo_id",
  "stockpile_cargo_id",
  "client_destination_id",
  "cargo_type_id",
  "project_status",
  "start_date",
  "end_date",
  "planned_volume_cubic",
  "planned_weight_ton",
  "volume_to_weight_conversion",
  "material_purchase_unit",
  "material_buy_price_per_cubic",
  "material_buy_price_per_ton",
  "material_sale_unit",
  "material_sell_price_per_cubic",
  "material_sell_price_per_ton",
  "fixed_other_income",
  "fixed_other_expense",
  "project_note",
  "is_active",
] as const;

export function buildProjectPayload(data: FormDataObject) {
  const payload = pickAdminPayload(data, projectPayloadFields);
  const routeType = payload.route_type;
  if (routeType === "MINE_CLIENT") {
    payload.vessel_cargo_id = null;
    payload.stockpile_cargo_id = null;
  } else if (routeType === "MINE_STOCKPILE_CLIENT") {
    payload.vessel_cargo_id = null;
  } else if (routeType === "VESSEL_CLIENT") {
    payload.mine_id = null;
    payload.stockpile_cargo_id = null;
  } else if (routeType === "VESSEL_STOCKPILE_CLIENT") {
    payload.mine_id = null;
  }
  if (payload.material_purchase_unit !== "M3") {
    payload.material_buy_price_per_cubic = null;
  }
  if (payload.material_purchase_unit !== "TON") {
    payload.material_buy_price_per_ton = null;
  }
  if (payload.material_sale_unit !== "M3") {
    payload.material_sell_price_per_cubic = null;
  }
  if (payload.material_sale_unit !== "TON") {
    payload.material_sell_price_per_ton = null;
  }
  return payload;
}

export const projectInitialData = {
  project_status: "DRAFT",
  material_purchase_unit: "NONE",
  material_sale_unit: "NONE",
  fixed_other_income: 0,
  fixed_other_expense: 0,
  is_active: 1,
};
