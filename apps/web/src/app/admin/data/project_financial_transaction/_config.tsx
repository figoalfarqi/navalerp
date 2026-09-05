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
import { formatCurrencyIDR } from "@/utils/currencyFormater";
import { formatDate, formatDateTime } from "@/utils/dateTime";

export const financialTransactionEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_financial_transaction`;
export const financialTransactionTableName =
  "project_financial_transaction";
export const financialTransactionPrimaryKey =
  "project_financial_transaction_id";

const transactionKindOptions = [
  { label: "Piutang", value: "RECEIVABLE" },
  { label: "Penerimaan", value: "RECEIPT" },
  { label: "Hutang", value: "PAYABLE" },
  { label: "Pembayaran", value: "PAYMENT" },
  { label: "Penyesuaian", value: "ADJUSTMENT" },
];

const incomeFields = [
  "material_sale_income",
  "transport_service_income",
  "other_income",
] as const;
const expenseFields = [
  "material_purchase_expense",
  "transport_expense",
  "road_money_expense",
  "loading_expense",
  "unloading_expense",
  "fuel_expense",
  "toll_expense",
  "other_expense",
] as const;

function totalFor(
  item: Record<string, unknown>,
  fields: readonly string[],
): number {
  return fields.reduce((total, field) => total + Number(item[field] ?? 0), 0);
}

export const financialTransactionColumns: ColumnField[] = [
  {
    key: "transaction_number",
    label: "Nomor Transaksi",
    sortable: "string",
    columnLength: 190,
  },
  {
    key: "project_id",
    label: "Project",
    render: (item) => item.project?.project_name ?? "-",
    sortable: "table_key",
    columnLength: 220,
  },
  {
    key: "transaction_kind",
    label: "Jenis",
    render: (item) =>
      transactionKindOptions.find(
        (option) => option.value === item.transaction_kind,
      )?.label ?? item.transaction_kind,
    sortable: "string",
    columnLength: 150,
  },
  {
    key: "transaction_date",
    label: "Tanggal",
    render: (item) => formatDate(item.transaction_date),
    sortable: "date",
    columnLength: 130,
  },
  {
    key: "due_date",
    label: "Jatuh Tempo",
    render: (item) => (item.due_date ? formatDate(item.due_date) : "-"),
    sortable: "date",
    columnLength: 140,
  },
  {
    key: "paid_at",
    label: "Dibayar",
    render: (item) => (item.paid_at ? formatDateTime(item.paid_at) : "-"),
    sortable: "date",
    columnLength: 180,
  },
  {
    key: "total_income",
    label: "Total Pendapatan",
    render: (item) => formatCurrencyIDR(totalFor(item, incomeFields)),
    sortable: "number",
    columnLength: 180,
  },
  {
    key: "total_expense",
    label: "Total Pengeluaran",
    render: (item) => formatCurrencyIDR(totalFor(item, expenseFields)),
    sortable: "number",
    columnLength: 180,
  },
  {
    key: "is_posted",
    label: "Posted",
    render: (item) => (Number(item.is_posted) === 1 ? "Ya" : "Belum"),
    sortable: "boolean",
    columnLength: 100,
  },
];

export const financialTransactionFilters: FilterField[] = [
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
    name: "transaction_kind",
    label: "Jenis Transaksi",
    fieldType: "select",
    options: transactionKindOptions,
    col: "right",
  },
  {
    name: "transaction_number",
    label: "Nomor Transaksi",
    fieldType: "text",
    col: "left",
  },
  {
    name: "transaction_date",
    label: "Tanggal Transaksi",
    fieldType: "dateAfterBefore",
    col: "right",
  },
  {
    name: "is_posted",
    label: "Posted",
    fieldType: "radio",
    options: [
      { label: "Ya", value: "1", variant: "green-outline" },
      { label: "Belum", value: "0", variant: "gray-outline" },
    ],
    col: "left",
  },
];

export function financialTransactionFields(
  mode: AdminCrudMode,
): FormField[] {
  const disabled = mode === "view";
  const amountDefinitions: [string, string][] = [
    ["material_sale_income", "Pendapatan Penjualan Material"],
    ["transport_service_income", "Pendapatan Jasa Transport"],
    ["other_income", "Pendapatan Lain"],
    ["material_purchase_expense", "Pembelian Material"],
    ["transport_expense", "Biaya Transport"],
    ["road_money_expense", "Uang Jalan"],
    ["loading_expense", "Biaya Loading"],
    ["unloading_expense", "Biaya Unloading"],
    ["fuel_expense", "Biaya BBM"],
    ["toll_expense", "Biaya Tol"],
    ["other_expense", "Pengeluaran Lain"],
  ];

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
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "project_transport_id",
      label: "Transport (opsional)",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_transport?limit=999&query=project_id%3D{project_id}`,
        labelKey: "transport_number",
        valueKey: "project_transport_id",
        dependsOn: ["project_id"],
      },
      disabled,
      col: "right",
    },
    {
      name: "transaction_number",
      label: "Nomor Transaksi",
      fieldType: "text",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "transaction_kind",
      label: "Jenis Transaksi",
      fieldType: "select",
      options: transactionKindOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "transaction_date",
      label: "Tanggal Transaksi",
      fieldType: "date",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "due_date",
      label: "Jatuh Tempo",
      fieldType: "date",
      disabled,
      col: "right",
    },
    {
      name: "paid_at",
      label: "Waktu Pembayaran",
      fieldType: "datetime",
      disabled,
      col: "left",
    },
    {
      name: "client_id",
      label: "Client",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/client?limit=999`,
        labelKey: "client_name",
        valueKey: "client_id",
      },
      disabled,
      col: "right",
    },
    {
      name: "vendor_id",
      label: "Vendor",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/vendor?limit=999`,
        labelKey: "vendor_name",
        valueKey: "vendor_id",
      },
      disabled,
      col: "left",
    },
    {
      name: "reference_number",
      label: "Nomor Referensi",
      fieldType: "text",
      disabled,
      col: "right",
    },
    ...amountDefinitions.map<FormField>(([name, label], index) => ({
      name,
      label,
      fieldType: "number",
      validation: { min: 0 },
      required: true,
      disabled,
      col: index % 2 === 0 ? "left" : "right",
    })),
    {
      name: "transaction_note",
      label: "Catatan Transaksi",
      fieldType: "textarea",
      disabled,
      col: "left",
    },
    {
      name: "is_posted",
      label: "Posted",
      fieldType: "boolean",
      required: true,
      disabled,
      col: "right",
    },
  ];
}

const payloadFields = [
  "project_id",
  "project_transport_id",
  "transaction_number",
  "transaction_kind",
  "transaction_date",
  "due_date",
  "paid_at",
  "client_id",
  "vendor_id",
  "reference_number",
  ...incomeFields,
  ...expenseFields,
  "transaction_note",
  "is_posted",
] as const;

export function buildFinancialTransactionPayload(data: FormDataObject) {
  const payload = pickAdminPayload(data, payloadFields);
  const kind = payload.transaction_kind;
  if (kind === "RECEIVABLE" || kind === "RECEIPT") {
    expenseFields.forEach((field) => {
      payload[field] = 0;
    });
  } else if (kind === "PAYABLE" || kind === "PAYMENT") {
    incomeFields.forEach((field) => {
      payload[field] = 0;
    });
  }
  return payload;
}

export const financialTransactionInitialData = {
  transaction_kind: "RECEIVABLE",
  transaction_date: new Date().toISOString().slice(0, 10),
  material_sale_income: 0,
  transport_service_income: 0,
  other_income: 0,
  material_purchase_expense: 0,
  transport_expense: 0,
  road_money_expense: 0,
  loading_expense: 0,
  unloading_expense: 0,
  fuel_expense: 0,
  toll_expense: 0,
  other_expense: 0,
  is_posted: 0,
};
