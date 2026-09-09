/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "contract";
export const entityTitle = "Kontrak Militer";
export const entityEndpoint = "/admin/contract";
export const primaryKey = "contract_id";

export const columns: ColumnField[] = [
  { key: "tender_number", label: "No. Tender" },
  { key: "contract_number", label: "Contract Number" },
  { key: "vendor_name", label: "Vendor" },
  { key: "contract_title", label: "Contract Title" },
  { key: "contract_value", label: "Contract Value" },
  { key: "currency", label: "Currency" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Kontrak Militer", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "tender_id",
    col: "left",
    label: "No. Tender",
    fieldType: "select",
    options: {
      url: "/admin/tender?limit=100",
      labelKey: "tender_number",
      valueKey: "tender_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "contract_number",
    col: "right",
    label: "Contract Number",
    fieldType: "text",
    required: true,
    readOnly: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "KTR",
  },
    {
    name: "vendor_id",
    col: "left",
    label: "Vendor",
    fieldType: "select",
    options: {
      url: "/admin/vendor?limit=100",
      labelKey: "vendor_name",
      valueKey: "vendor_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "contract_title",
    col: "right",
    label: "Contract Title",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "contract_value",
    col: "left",
    label: "Contract Value",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "currency",
    col: "right",
    label: "Currency",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "start_date",
    col: "left",
    label: "Start Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "end_date",
    col: "right",
    label: "End Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "procurement_method",
    col: "left",
    label: "Procurement Method",
    fieldType: "select",
    options: [
      { label: "Direct Appointment", value: "DIRECT_APPOINTMENT" },
      { label: "Limited Tender", value: "LIMITED_TENDER" },
      { label: "Open Tender", value: "OPEN_TENDER" },
      { label: "G2G", value: "G2G" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "warranty_period_months",
    col: "right",
    label: "Warranty Period Months",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "tot_clause_summary",
    col: "left",
    label: "Tot Clause Summary",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "right",
    label: "Status",
    fieldType: "select",
    options: [
      { label: "Draft", value: "DRAFT" },
      { label: "Signed", value: "SIGNED" },
      { label: "Active", value: "ACTIVE" },
      { label: "Completed", value: "COMPLETED" },
      { label: "Terminated", value: "TERMINATED" },
    ],
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.contract_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
