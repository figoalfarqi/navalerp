/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "contract";
export const entityTitle = "Kontrak Militer";
export const entityEndpoint = "/admin/contract";
export const primaryKey = "contract_id";

export const columns: ColumnField[] = [
  { key: "contract_id", label: "Contract Id" },
  { key: "tender_id", label: "Tender Id" },
  { key: "contract_number", label: "Contract Number" },
  { key: "vendor_id", label: "Vendor Id" },
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
    label: "Tender Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "contract_number",
    label: "Contract Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "vendor_id",
    label: "Vendor Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "contract_title",
    label: "Contract Title",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "contract_value",
    label: "Contract Value",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "currency",
    label: "Currency",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "start_date",
    label: "Start Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "end_date",
    label: "End Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "procurement_method",
    label: "Procurement Method",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "warranty_period_months",
    label: "Warranty Period Months",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "tot_clause_summary",
    label: "Tot Clause Summary",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    label: "Status",
    fieldType: "text",
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
