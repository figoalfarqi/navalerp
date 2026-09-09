/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";
import { DetailItemsConfig } from "@/components/formCrud/DetailItemsTable";

export const entityName = "journal_entry";
export const entityTitle = "Jurnal Akuntansi";
export const entityEndpoint = "/admin/journal_entry";
export const primaryKey = "journal_id";

export const columns: ColumnField[] = [
  { key: "entry_number", label: "Entry Number" },
  { key: "entry_date", label: "Entry Date", render: (item: any) => formatSmartDate(item.entry_date) },
  { key: "description", label: "Description" },
  { key: "source_module", label: "Source Module" },
  { key: "is_posted", label: "Is Posted" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Jurnal Akuntansi", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "entry_number",
    col: "left",
    label: "Entry Number",
    fieldType: "text",
    required: true,
    readOnly: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "JRN",
  },
  {
    name: "entry_date",
    col: "right",
    label: "Entry Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "description",
    col: "left",
    label: "Description",
    fieldType: "textarea",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "source_module",
    col: "right",
    label: "Source Module",
    fieldType: "select",
    options: [
      { label: "Procurement", value: "PROCUREMENT" },
      { label: "Inventory", value: "INVENTORY" },
      { label: "MRO", value: "MRO" },
      { label: "Asset Capitalization", value: "ASSET_CAPITALIZATION" },
      { label: "Payroll", value: "PAYROLL" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "source_reference_id",
    col: "left",
    label: "Source Reference Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_posted",
    col: "right",
    label: "Is Posted",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
];

export const detailItemsConfig: DetailItemsConfig = {
  tableName: "lines",
  title: "Rincian Baris Jurnal",
  itemName: "Baris Jurnal",
  columns: [
    {
      key: "account_id",
      label: "Akun (Chart of Account)",
      type: "select",
      required: true,
      placeholder: "Pilih Akun Rekening...",
      options: {
        url: "/admin/chart_of_account?limit=100",
        labelKey: "account_name",
        valueKey: "account_id",
        extraLabelKey: "account_code",
      },
    },
    {
      key: "debit",
      label: "Debit",
      type: "currency",
      required: false,
      defaultValue: 0,
      width: "160px",
    },
    {
      key: "credit",
      label: "Kredit",
      type: "currency",
      required: false,
      defaultValue: 0,
      width: "160px",
    },
    {
      key: "memo",
      label: "Memo / Keterangan",
      type: "text",
      required: false,
      placeholder: "Keterangan baris jurnal...",
    },
  ],
};

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.journal_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;

  if (Array.isArray(payload.lines)) {
    payload.lines = payload.lines.map((it: any) => ({
      account_id: it.account_id,
      debit: Number(it.debit) || 0,
      credit: Number(it.credit) || 0,
      memo: it.memo || null,
    }));
  }

  return payload;
};
