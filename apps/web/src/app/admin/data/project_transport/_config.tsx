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
import { formatDateTime } from "@/utils/dateTime";
import { getTransportStatusLabel } from "./_labels";

export const projectTransportEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_transport`;
export const projectTransportTableName = "project_transport";
export const projectTransportPrimaryKey = "project_transport_id";

const materialUnitOptions = [
  { label: "Tidak digunakan", value: "NONE" },
  { label: "Per m³", value: "M3" },
  { label: "Per ton", value: "TON" },
];
const transportUnitOptions = [
  ...materialUnitOptions,
  { label: "Per transport", value: "TRANSPORT" },
];

export const projectTransportColumns: ColumnField[] = [
  {
    key: "transport_number",
    label: "No. Transport",
    sortable: "string",
    columnLength: 170,
  },
  {
    key: "project_name",
    label: "Project",
    render: (item) => item.project?.project_name ?? "-",
    sortable: "table_key",
    columnLength: 230,
  },
  {
    key: "route_name",
    label: "Rute",
    render: (item) => item.project_route?.route_name ?? "-",
    sortable: "table_key",
    columnLength: 200,
  },
  {
    key: "transported_at",
    label: "Waktu Transport",
    render: (item) => formatDateTime(item.transported_at),
    sortable: "date",
    columnLength: 180,
  },
  {
    key: "license_plate",
    label: "Truck",
    render: (item) => item.truck?.license_plate ?? "-",
    sortable: "table_key",
    columnLength: 140,
  },
  {
    key: "driver_name",
    label: "Driver",
    render: (item) => item.driver?.app_user_name ?? "-",
    sortable: "table_key",
    columnLength: 180,
  },
  {
    key: "transport_vendor_name",
    label: "Vendor Transport",
    render: (item) => item.transport_vendor?.vendor_name ?? "-",
    sortable: "string",
    columnLength: 190,
  },
  {
    key: "latest_status_type_id",
    label: "Status Terakhir",
    render: (item) =>
      item.latest_status?.project_transport_status_type_id == null
        ? "Belum ada status"
        : getTransportStatusLabel(
            item.latest_status.project_transport_status_type_id,
          ),
    columnLength: 230,
  },
  {
    key: "activity_summary",
    label: "Aktivitas",
    render: (item) =>
      `${Number(item.status_count ?? 0)} status · ${Number(
        item.photo_count ?? 0,
      )} foto`,
    columnLength: 140,
  },
  {
    key: "loaded_volume_cubic",
    label: "Volume Muat",
    render: (item) =>
      item.loaded_volume_cubic == null
        ? "-"
        : `${formatNumberID(item.loaded_volume_cubic)} m³`,
    sortable: "number",
    columnLength: 150,
  },
  {
    key: "loaded_weight_ton",
    label: "Berat Muat",
    render: (item) =>
      item.loaded_weight_ton == null
        ? "-"
        : `${formatNumberID(item.loaded_weight_ton)} ton`,
    sortable: "number",
    columnLength: 150,
  },
  {
    key: "transport_service_income_amount",
    label: "Pendapatan Jasa",
    render: (item) =>
      formatCurrencyIDR(item.transport_service_income_amount ?? 0),
    sortable: "number",
    columnLength: 180,
  },
  {
    key: "transport_expense_amount",
    label: "Biaya Angkut",
    render: (item) => formatCurrencyIDR(item.transport_expense_amount ?? 0),
    sortable: "number",
    columnLength: 180,
  },
  {
    key: "is_completed",
    label: "Selesai",
    render: (item) => (Number(item.is_completed) === 1 ? "Ya" : "Belum"),
    sortable: "boolean",
    columnLength: 100,
  },
];

export const projectTransportFilters: FilterField[] = [
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
    name: "transport_number",
    label: "Nomor Transport",
    fieldType: "text",
    col: "right",
  },
  {
    name: "truck_id",
    label: "Truck",
    fieldType: "select",
    options: {
      url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck?limit=999`,
      labelKey: "license_plate",
      valueKey: "truck_id",
    },
    col: "left",
  },
  {
    name: "transported_at",
    label: "Tanggal Transport",
    fieldType: "dateAfterBefore",
    col: "right",
  },
  {
    name: "is_completed",
    label: "Status",
    fieldType: "radio",
    options: [
      { label: "Selesai", value: "1", variant: "green-outline" },
      { label: "Belum", value: "0", variant: "gray-outline" },
    ],
    col: "left",
  },
];

function numberFields(
  disabled: boolean,
  definitions: [string, string][],
): FormField[] {
  return definitions.map(([name, label], index) => ({
    name,
    label,
    fieldType: "number",
    validation: { min: 0 },
    disabled,
    col: index % 2 === 0 ? "left" : "right",
  }));
}

export function projectTransportFields(mode: AdminCrudMode): FormField[] {
  const disabled = mode === "view";
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
      rowGetters: [
        {
          key: "volume_to_weight_conversion",
          source: "volume_to_weight_conversion",
        },
        { key: "material_purchase_unit", source: "material_purchase_unit" },
        {
          key: "material_buy_price_per_cubic",
          source: "material_buy_price_per_cubic",
        },
        {
          key: "material_buy_price_per_ton",
          source: "material_buy_price_per_ton",
        },
        { key: "material_sale_unit", source: "material_sale_unit" },
        {
          key: "material_sell_price_per_cubic",
          source: "material_sell_price_per_cubic",
        },
        {
          key: "material_sell_price_per_ton",
          source: "material_sell_price_per_ton",
        },
      ],
      clearFieldsOnChange: [
        "project_route_id",
        "truck_id",
        "project_truck_assignment_id",
        "driver_id",
        "transport_vendor_id",
        "volume_to_weight_conversion",
        "material_purchase_unit",
        "material_buy_price_per_cubic",
        "material_buy_price_per_ton",
        "material_sale_unit",
        "material_sell_price_per_cubic",
        "material_sell_price_per_ton",
        "transport_service_unit",
        "transport_service_price_per_cubic",
        "transport_service_price_per_ton",
        "transport_service_price_per_transport",
        "transport_cost_unit",
        "transport_cost_per_cubic",
        "transport_cost_per_ton",
        "transport_cost_per_transport",
        "road_money_amount",
        "loading_cost_amount",
        "unloading_cost_amount",
        "fuel_cost_amount",
        "toll_cost_amount",
        "other_income_amount",
        "other_expense_amount",
      ],
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "project_route_id",
      label: "Rute Project",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_route?limit=999&query=project_id%3D{project_id}`,
        labelKey: "route_name",
        valueKey: "project_route_id",
        dependsOn: ["project_id"],
      },
      rowGetters: [
        { key: "transport_service_unit", source: "transport_service_unit" },
        {
          key: "transport_service_price_per_cubic",
          source: "transport_service_price_per_cubic",
        },
        {
          key: "transport_service_price_per_ton",
          source: "transport_service_price_per_ton",
        },
        {
          key: "transport_service_price_per_transport",
          source: "transport_service_price_per_transport",
        },
        { key: "transport_cost_unit", source: "transport_cost_unit" },
        {
          key: "transport_cost_per_cubic",
          source: "transport_cost_per_cubic",
        },
        {
          key: "transport_cost_per_ton",
          source: "transport_cost_per_ton",
        },
        {
          key: "transport_cost_per_transport",
          source: "transport_cost_per_transport",
        },
        { key: "road_money_amount", source: "road_money_per_transport" },
        {
          key: "loading_cost_amount",
          source: "loading_cost_per_transport",
        },
        {
          key: "unloading_cost_amount",
          source: "unloading_cost_per_transport",
        },
        { key: "fuel_cost_amount", source: "fuel_cost_per_transport" },
        { key: "toll_cost_amount", source: "toll_cost_per_transport" },
        {
          key: "other_income_amount",
          source: "other_income_per_transport",
        },
        {
          key: "other_expense_amount",
          source: "other_expense_per_transport",
        },
      ],
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "transport_number",
      label: "Nomor Transport",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "delivery_note_number",
      label: "Nomor Delivery Note",
      fieldType: "text",
      disabled,
      col: "right",
    },
    {
      name: "transported_at",
      label: "Waktu Transport",
      fieldType: "datetime",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "truck_id",
      label: "Truck",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_truck_assignment?limit=999&is_active=1&query=project_id%3D{project_id}`,
        labelKey: "truck.license_plate",
        valueKey: "truck_id",
        dependsOn: ["project_id"],
      },
      rowGetters: [
        {
          key: "project_truck_assignment_id",
          source: "project_truck_assignment_id",
        },
        { key: "driver_id", source: "truck.driver.app_user_id" },
        { key: "transport_vendor_id", source: "truck.vendor.vendor_id" },
      ],
      clearFieldsOnChange: [
        "project_truck_assignment_id",
        "driver_id",
        "transport_vendor_id",
      ],
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "driver_id",
      label: "Driver",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/driver?limit=999`,
        labelKey: "app_user_name",
        valueKey: "app_user_id",
      },
      disabled: true,
      col: "right",
    },
    {
      name: "transport_vendor_id",
      label: "Vendor Transport",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/vendor?limit=999`,
        labelKey: "vendor_name",
        valueKey: "vendor_id",
      },
      disabled: true,
      col: "left",
    },
    ...numberFields(disabled, [
      ["loaded_volume_cubic", "Volume Muat (m³)"],
      ["loaded_weight_ton", "Berat Muat (ton)"],
      ["delivered_volume_cubic", "Volume Terkirim (m³)"],
      ["delivered_weight_ton", "Berat Terkirim (ton)"],
      ["purchase_volume_cubic", "Volume Pembelian (m³)"],
      ["purchase_weight_ton", "Berat Pembelian (ton)"],
      ["sale_volume_cubic", "Volume Penjualan (m³)"],
      ["sale_weight_ton", "Berat Penjualan (ton)"],
      ["transport_service_volume_cubic", "Volume Jasa (m³)"],
      ["transport_service_weight_ton", "Berat Jasa (ton)"],
      ["transport_cost_volume_cubic", "Volume Biaya Angkut (m³)"],
      ["transport_cost_weight_ton", "Berat Biaya Angkut (ton)"],
    ]),
    {
      name: "volume_to_weight_conversion",
      label: "Konversi Volume ke Berat",
      fieldType: "number",
      disabled: true,
      col: "left",
    },
    {
      name: "material_purchase_unit",
      label: "Unit Pembelian",
      fieldType: "select",
      options: materialUnitOptions,
      required: true,
      disabled: true,
      col: "left",
    },
    ...numberFields(true, [
      ["material_buy_price_per_cubic", "Harga Beli per m³"],
      ["material_buy_price_per_ton", "Harga Beli per ton"],
    ]),
    {
      name: "material_sale_unit",
      label: "Unit Penjualan",
      fieldType: "select",
      options: materialUnitOptions,
      required: true,
      disabled: true,
      col: "right",
    },
    ...numberFields(true, [
      ["material_sell_price_per_cubic", "Harga Jual per m³"],
      ["material_sell_price_per_ton", "Harga Jual per ton"],
    ]),
    {
      name: "transport_service_unit",
      label: "Unit Jasa Transport",
      fieldType: "select",
      options: transportUnitOptions,
      required: true,
      disabled: true,
      col: "left",
    },
    ...numberFields(true, [
      ["transport_service_price_per_cubic", "Harga Jasa per m³"],
      ["transport_service_price_per_ton", "Harga Jasa per ton"],
      ["transport_service_price_per_transport", "Harga Jasa per transport"],
    ]),
    {
      name: "transport_cost_unit",
      label: "Unit Biaya Transport",
      fieldType: "select",
      options: transportUnitOptions,
      required: true,
      disabled: true,
      col: "right",
    },
    ...numberFields(true, [
      ["transport_cost_per_cubic", "Biaya Transport per m³"],
      ["transport_cost_per_ton", "Biaya Transport per ton"],
      ["transport_cost_per_transport", "Biaya per transport"],
      ["road_money_amount", "Uang Jalan"],
      ["loading_cost_amount", "Biaya Loading"],
      ["unloading_cost_amount", "Biaya Unloading"],
      ["fuel_cost_amount", "Biaya BBM"],
      ["toll_cost_amount", "Biaya Tol"],
      ["other_income_amount", "Pendapatan Lain"],
      ["other_expense_amount", "Pengeluaran Lain"],
    ]),
    {
      name: "is_completed",
      label: "Transport Selesai",
      fieldType: "boolean",
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "transport_note",
      label: "Catatan Transport",
      fieldType: "textarea",
      disabled,
      col: "left",
    },
  ];
}

const payloadFields = [
  "project_id",
  "project_route_id",
  "transport_number",
  "delivery_note_number",
  "transported_at",
  "truck_id",
  "project_truck_assignment_id",
  "driver_id",
  "transport_vendor_id",
  "loaded_volume_cubic",
  "loaded_weight_ton",
  "delivered_volume_cubic",
  "delivered_weight_ton",
  "purchase_volume_cubic",
  "purchase_weight_ton",
  "sale_volume_cubic",
  "sale_weight_ton",
  "transport_service_volume_cubic",
  "transport_service_weight_ton",
  "transport_cost_volume_cubic",
  "transport_cost_weight_ton",
  "volume_to_weight_conversion",
  "material_purchase_unit",
  "material_buy_price_per_cubic",
  "material_buy_price_per_ton",
  "material_sale_unit",
  "material_sell_price_per_cubic",
  "material_sell_price_per_ton",
  "transport_service_unit",
  "transport_service_price_per_cubic",
  "transport_service_price_per_ton",
  "transport_service_price_per_transport",
  "transport_cost_unit",
  "transport_cost_per_cubic",
  "transport_cost_per_ton",
  "transport_cost_per_transport",
  "road_money_amount",
  "loading_cost_amount",
  "unloading_cost_amount",
  "fuel_cost_amount",
  "toll_cost_amount",
  "other_income_amount",
  "other_expense_amount",
  "is_completed",
  "transport_note",
] as const;

function clearPriceFields(
  payload: Record<string, unknown>,
  unitField: string,
  cubicField: string,
  tonField: string,
  transportField?: string,
) {
  const unit = payload[unitField];
  if (unit !== "M3") payload[cubicField] = null;
  if (unit !== "TON") payload[tonField] = null;
  if (transportField && unit !== "TRANSPORT") payload[transportField] = null;
}

export function buildProjectTransportPayload(data: FormDataObject) {
  const payload = pickAdminPayload(data, payloadFields);
  if (!payload.truck_id) payload.project_truck_assignment_id = null;
  clearPriceFields(
    payload,
    "material_purchase_unit",
    "material_buy_price_per_cubic",
    "material_buy_price_per_ton",
  );
  clearPriceFields(
    payload,
    "material_sale_unit",
    "material_sell_price_per_cubic",
    "material_sell_price_per_ton",
  );
  clearPriceFields(
    payload,
    "transport_service_unit",
    "transport_service_price_per_cubic",
    "transport_service_price_per_ton",
    "transport_service_price_per_transport",
  );
  clearPriceFields(
    payload,
    "transport_cost_unit",
    "transport_cost_per_cubic",
    "transport_cost_per_ton",
    "transport_cost_per_transport",
  );
  return payload;
}

export const projectTransportInitialData = {
  transported_at: new Date().toISOString(),
  material_purchase_unit: "NONE",
  material_sale_unit: "NONE",
  transport_service_unit: "NONE",
  transport_cost_unit: "NONE",
  road_money_amount: 0,
  loading_cost_amount: 0,
  unloading_cost_amount: 0,
  fuel_cost_amount: 0,
  toll_cost_amount: 0,
  other_income_amount: 0,
  other_expense_amount: 0,
  is_completed: 0,
};
