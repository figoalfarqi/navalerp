import type { AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { pickAdminPayload } from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";
import { formatNumberID } from "@/utils/currencyFormater";

export const stockpileAdjustmentEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/stockpile_adjustment`;
export const stockpileAdjustmentTableName = "stockpile_adjustment";
export const stockpileAdjustmentPrimaryKey = "stockpile_adjustment_id";

const projectOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project?limit=999`,
  labelKey: "project_name",
  valueKey: "project_id",
};
const stockpileCargoOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/stockpile_cargo?limit=999&is_active=1`,
  labelKey: "{stockpile.stockpile_name} - {cargo_type.cargo_type_name}",
  valueKey: "stockpile_cargo_id",
};
const transportOptions = {
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_transport?limit=999`,
  labelKey: "transport_number",
  valueKey: "project_transport_id",
  dependsOn: ["project_id"],
};
export const stockReferenceOptions = [
  { label: "Transport masuk", value: 1 },
  { label: "Transport keluar", value: 2 },
  { label: "Penyesuaian stockpile", value: 3 },
  { label: "Manual", value: 4 },
  { label: "Koreksi", value: 5 },
];

const referenceLabel = (value: unknown) =>
  stockReferenceOptions.find((option) => option.value === Number(value))
    ?.label ?? "-";

export const stockpileAdjustmentColumns: ColumnField[] = [
  {
    key: "adjustment_date",
    label: "Tanggal",
    sortable: "date",
    columnLength: 130,
  },
  {
    key: "project_id",
    label: "Project",
    render: (item) => item.project?.project_name ?? "-",
    sortable: "table_key",
    columnLength: 220,
  },
  {
    key: "stockpile_cargo_id",
    label: "Stok Muatan",
    render: (item) => {
      const stockpile =
        item.stockpile_cargo?.stockpile?.stockpile_name ?? "";
      const cargo =
        item.stockpile_cargo?.cargo_type?.cargo_type_name ?? "";
      return [stockpile, cargo].filter(Boolean).join(" / ") || "-";
    },
    sortable: "table_key",
    columnLength: 240,
  },
  {
    key: "project_transport_id",
    label: "Transport",
    render: (item) =>
      item.project_transport?.transport_number ?? "-",
    sortable: "table_key",
    columnLength: 170,
  },
  {
    key: "reference_type_id",
    label: "Referensi",
    render: (item) => referenceLabel(item.reference_type_id),
    sortable: "number",
    columnLength: 170,
  },
  {
    key: "amount_volume_cubic",
    label: "Volume (m³)",
    render: (item) => formatNumberID(Number(item.amount_volume_cubic ?? 0)),
    sortable: "number",
    columnLength: 140,
  },
  {
    key: "amount_weight_ton",
    label: "Berat (ton)",
    render: (item) => formatNumberID(Number(item.amount_weight_ton ?? 0)),
    sortable: "number",
    columnLength: 140,
  },
  {
    key: "reason",
    label: "Alasan",
    sortable: "string",
    columnLength: 280,
  },
];

export const stockpileAdjustmentFilters: FilterField[] = [
  {
    name: "project_id",
    label: "Project",
    fieldType: "select",
    options: projectOptions,
    col: "left",
  },
  {
    name: "stockpile_cargo_id",
    label: "Stok Muatan",
    fieldType: "select",
    options: stockpileCargoOptions,
    col: "right",
  },
  {
    name: "reference_type_id",
    label: "Jenis Referensi",
    fieldType: "select",
    options: stockReferenceOptions,
    col: "left",
  },
  {
    name: "adjustment_date",
    label: "Tanggal",
    fieldType: "dateAfterBefore",
    col: "right",
  },
];

export function stockpileAdjustmentFields(
  mode: AdminCrudMode,
): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "project_id",
      label: "Project",
      fieldType: "select",
      options: projectOptions,
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "stockpile_cargo_id",
      label: "Stok Muatan Stockpile",
      fieldType: "select",
      options: stockpileCargoOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "adjustment_date",
      label: "Tanggal Penyesuaian",
      fieldType: "date",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "project_transport_id",
      label: "Transport Terkait",
      fieldType: "select",
      options: transportOptions,
      disabled,
      col: "right",
    },
    {
      name: "reference_type_id",
      label: "Jenis Referensi",
      fieldType: "select",
      options: stockReferenceOptions,
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "amount_volume_cubic",
      label: "Perubahan Volume (m³)",
      fieldType: "number",
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "amount_weight_ton",
      label: "Perubahan Berat (ton)",
      fieldType: "number",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "reason",
      label: "Alasan",
      fieldType: "textarea",
      disabled,
      col: "right",
    },
  ];
}

export const buildStockpileAdjustmentPayload = (data: FormDataObject) =>
  pickAdminPayload(data, [
    "project_id",
    "stockpile_cargo_id",
    "adjustment_date",
    "project_transport_id",
    "reference_type_id",
    "amount_volume_cubic",
    "amount_weight_ton",
    "reason",
  ]);
