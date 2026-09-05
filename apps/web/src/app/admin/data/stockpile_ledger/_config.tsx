import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";
import { formatNumberID } from "@/utils/currencyFormater";
import { formatDateTime } from "@/utils/dateTime";
import { stockReferenceOptions } from "../stockpile_adjustment/_config";

export const stockpileLedgerEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/stockpile_ledger`;
export const stockpileLedgerTableName = "stockpile_ledger";

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

const referenceLabel = (value: unknown) =>
  stockReferenceOptions.find((option) => option.value === Number(value))
    ?.label ?? "-";

export const stockpileLedgerColumns: ColumnField[] = [
  {
    key: "created_at",
    label: "Waktu",
    render: (item) => formatDateTime(item.created_at),
    sortable: "date",
    columnLength: 180,
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
    key: "adjustment_id",
    label: "Penyesuaian",
    render: (item) =>
      item.adjustment?.adjustment_date
        ? `${item.adjustment.adjustment_date}${
            item.adjustment.reason ? ` · ${item.adjustment.reason}` : ""
          }`
        : "-",
    sortable: "table_key",
    columnLength: 130,
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
    label: "Perubahan m³",
    render: (item) => formatNumberID(Number(item.amount_volume_cubic ?? 0)),
    sortable: "number",
    columnLength: 150,
  },
  {
    key: "balance_volume_cubic",
    label: "Saldo m³",
    render: (item) => formatNumberID(Number(item.balance_volume_cubic ?? 0)),
    sortable: "number",
    columnLength: 150,
  },
  {
    key: "amount_weight_ton",
    label: "Perubahan ton",
    render: (item) => formatNumberID(Number(item.amount_weight_ton ?? 0)),
    sortable: "number",
    columnLength: 150,
  },
  {
    key: "balance_weight_ton",
    label: "Saldo ton",
    render: (item) => formatNumberID(Number(item.balance_weight_ton ?? 0)),
    sortable: "number",
    columnLength: 150,
  },
];

export const stockpileLedgerFilters: FilterField[] = [
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
    name: "created_at",
    label: "Waktu Pencatatan",
    fieldType: "dateAfterBefore",
    col: "right",
  },
];
