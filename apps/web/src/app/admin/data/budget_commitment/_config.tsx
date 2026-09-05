/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "budget_commitment";
export const entityTitle = "Komitmen Anggaran";
export const entityEndpoint = "/admin/budget_commitment";
export const primaryKey = "commitment_id";

export const columns: ColumnField[] = [
  { key: "commitment_id", label: "Commitment Id" },
  { key: "commitment_number", label: "Commitment Number" },
  { key: "allocation_id", label: "Allocation Id" },
  { key: "contract_id", label: "Contract Id" },
  { key: "po_id", label: "Po Id" },
  { key: "work_order_id", label: "Work Order Id" },
  { key: "committed_amount", label: "Committed Amount" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Komitmen Anggaran", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "commitment_number",
    label: "Commitment Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "allocation_id",
    label: "Allocation Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "contract_id",
    label: "Contract Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "po_id",
    label: "Po Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "work_order_id",
    label: "Work Order Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "committed_amount",
    label: "Committed Amount",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "commitment_date",
    label: "Commitment Date",
    fieldType: "date",
    required: true,
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
  delete payload.commitment_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
