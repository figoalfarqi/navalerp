import {
  pickAdminPayload,
  type AdminCrudMode,
} from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";
import { formatCurrencyIDR, formatNumberID } from "@/utils/currencyFormater";

export const projectRouteEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_route`;
export const projectRouteTableName = "project_route";
export const projectRoutePrimaryKey = "project_route_id";

const patternOptions = [
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

const routeTypeOptions = [
  { label: "Source → Client", value: "SOURCE_TO_CLIENT" },
  { label: "Source → Stockpile", value: "SOURCE_TO_STOCKPILE" },
  { label: "Stockpile → Client", value: "STOCKPILE_TO_CLIENT" },
];

const priceUnitOptions = [
  { label: "Tidak digunakan", value: "NONE" },
  { label: "Per m³", value: "M3" },
  { label: "Per ton", value: "TON" },
  { label: "Per transport", value: "TRANSPORT" },
];

export const projectRouteColumns: ColumnField[] = [
  {
    key: "project_id",
    label: "Project",
    render: (item) => item.project?.project_name ?? "-",
    sortable: "table_key",
    columnLength: 230,
  },
  {
    key: "route_sequence",
    label: "Urutan",
    sortable: "number",
    columnLength: 100,
  },
  {
    key: "route_name",
    label: "Nama Rute",
    sortable: "string",
    columnLength: 220,
  },
  {
    key: "route_type",
    label: "Tipe Rute",
    sortable: "string",
    columnLength: 210,
  },
  {
    key: "distance_km",
    label: "Jarak",
    render: (item) =>
      item.distance_km == null
        ? "-"
        : `${formatNumberID(item.distance_km)} km`,
    sortable: "number",
    columnLength: 130,
  },
  {
    key: "transport_service_price_per_transport",
    label: "Jasa / Transport",
    render: (item) =>
      item.transport_service_price_per_transport == null
        ? "-"
        : formatCurrencyIDR(item.transport_service_price_per_transport),
    sortable: "number",
    columnLength: 180,
  },
  {
    key: "transport_cost_per_transport",
    label: "Biaya / Transport",
    render: (item) =>
      item.transport_cost_per_transport == null
        ? "-"
        : formatCurrencyIDR(item.transport_cost_per_transport),
    sortable: "number",
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

export const projectRouteFilters: FilterField[] = [
  {
    name: "project_id",
    label: "Project",
    fieldType: "select",
    options: {
      url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project?limit=999`,
      labelKey: "project_name",
      valueKey: "project_id",
    },
    col: "left",
  },
  {
    name: "route_type",
    label: "Tipe Rute",
    fieldType: "select",
    options: routeTypeOptions,
    col: "right",
  },
  {
    name: "route_name",
    label: "Nama Rute",
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

export function projectRouteFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
  const numericFields: FormField[] = [
    ["distance_km", "Jarak (km)"],
    ["transport_service_price_per_cubic", "Harga Jasa per m³"],
    ["transport_service_price_per_ton", "Harga Jasa per ton"],
    ["transport_service_price_per_transport", "Harga Jasa per transport"],
    ["transport_cost_per_cubic", "Biaya Angkut per m³"],
    ["transport_cost_per_ton", "Biaya Angkut per ton"],
    ["transport_cost_per_transport", "Biaya Angkut per transport"],
    ["road_money_per_transport", "Uang Jalan / transport"],
    ["loading_cost_per_transport", "Biaya Loading / transport"],
    ["unloading_cost_per_transport", "Biaya Unloading / transport"],
    ["fuel_cost_per_transport", "Biaya BBM / transport"],
    ["toll_cost_per_transport", "Biaya Tol / transport"],
    ["other_income_per_transport", "Pendapatan Lain / transport"],
    ["other_expense_per_transport", "Pengeluaran Lain / transport"],
  ].map(([name, label], index) => ({
    name,
    label,
    fieldType: "number",
    validation: { min: 0 },
    disabled,
    col: index % 2 === 0 ? "left" : "right",
  }));

  return [
    {
      name: "project_id",
      label: "Project",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project?limit=999`,
        labelKey: "project_name",
        valueKey: "project_id",
      },
      rowGetters: [{ key: "project_pattern", source: "route_type" }],
      clearFieldsOnChange: ["project_pattern"],
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "project_pattern",
      label: "Pola Project",
      fieldType: "select",
      options: patternOptions,
      required: true,
      disabled: true,
      col: "right",
    },
    {
      name: "route_sequence",
      label: "Urutan Rute",
      fieldType: "number",
      validation: { min: 1, max: 2 },
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "route_type",
      label: "Tipe Rute",
      fieldType: "select",
      options: routeTypeOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "route_name",
      label: "Nama Rute",
      fieldType: "text",
      disabled,
      col: "left",
    },
    {
      name: "transport_service_unit",
      label: "Unit Harga Jasa",
      fieldType: "select",
      options: priceUnitOptions,
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "transport_cost_unit",
      label: "Unit Biaya Angkut",
      fieldType: "select",
      options: priceUnitOptions,
      required: true,
      disabled,
      col: "right",
    },
    ...numericFields,
    {
      name: "route_note",
      label: "Catatan Rute",
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

const payloadFields = [
  "project_id",
  "project_pattern",
  "route_sequence",
  "route_type",
  "route_name",
  "distance_km",
  "transport_service_unit",
  "transport_service_price_per_cubic",
  "transport_service_price_per_ton",
  "transport_service_price_per_transport",
  "transport_cost_unit",
  "transport_cost_per_cubic",
  "transport_cost_per_ton",
  "transport_cost_per_transport",
  "road_money_per_transport",
  "loading_cost_per_transport",
  "unloading_cost_per_transport",
  "fuel_cost_per_transport",
  "toll_cost_per_transport",
  "other_income_per_transport",
  "other_expense_per_transport",
  "route_note",
  "is_active",
] as const;

function clearUnitPrices(
  payload: Record<string, unknown>,
  unitField: string,
  cubicField: string,
  tonField: string,
  transportField: string,
) {
  const unit = payload[unitField];
  if (unit !== "M3") payload[cubicField] = null;
  if (unit !== "TON") payload[tonField] = null;
  if (unit !== "TRANSPORT") payload[transportField] = null;
}

export function buildProjectRoutePayload(data: FormDataObject) {
  const payload = pickAdminPayload(data, payloadFields);
  clearUnitPrices(
    payload,
    "transport_service_unit",
    "transport_service_price_per_cubic",
    "transport_service_price_per_ton",
    "transport_service_price_per_transport",
  );
  clearUnitPrices(
    payload,
    "transport_cost_unit",
    "transport_cost_per_cubic",
    "transport_cost_per_ton",
    "transport_cost_per_transport",
  );
  return payload;
}

export const projectRouteInitialData = {
  route_sequence: 1,
  transport_service_unit: "NONE",
  transport_cost_unit: "NONE",
  road_money_per_transport: 0,
  loading_cost_per_transport: 0,
  unloading_cost_per_transport: 0,
  fuel_cost_per_transport: 0,
  toll_cost_per_transport: 0,
  other_income_per_transport: 0,
  other_expense_per_transport: 0,
  is_active: 1,
};
