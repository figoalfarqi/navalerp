/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "budget_program";
export const entityTitle = "Program DIPA";
export const entityEndpoint = "/admin/budget_program";
export const primaryKey = "program_id";

export const columns: ColumnField[] = [
  { key: "fiscal_year", label: "Fiscal Year" },
  { key: "dipa_number", label: "Dipa Number" },
  { key: "program_code", label: "Program Code" },
  { key: "program_name", label: "Program Name" },
  { key: "total_budget", label: "Total Budget" },
  { key: "responsible_unit_name", label: "Satuan Penanggung Jawab" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Program DIPA", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "fiscal_year",
    col: "left",
    label: "Fiscal Year",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "dipa_number",
    col: "right",
    label: "Dipa Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "program_code",
    col: "left",
    label: "Program Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "program_name",
    col: "right",
    label: "Program Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "total_budget",
    col: "left",
    label: "Total Budget",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "responsible_unit_id",
    col: "right",
    label: "Satuan Penanggung Jawab",
    fieldType: "select",
    options: {
      url: "/admin/org_unit?limit=100",
      labelKey: "unit_name",
      valueKey: "unit_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "left",
    label: "Status",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.program_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
