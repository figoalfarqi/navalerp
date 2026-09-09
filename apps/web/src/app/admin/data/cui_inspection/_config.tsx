/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "cui_inspection";
export const entityTitle = "Inspeksi CUI";
export const entityEndpoint = "/admin/cui_inspection";
export const primaryKey = "inspection_id";

export const columns: ColumnField[] = [
  { key: "inspection_id", label: "Inspection Id" },
  { key: "inspection_number", label: "Inspection Number" },
  { key: "cui_asset_id", label: "Cui Asset Id" },
  { key: "ship_id", label: "Ship Id" },
  { key: "inspection_date", label: "Inspection Date" },
  { key: "inspector_officer_id", label: "Inspector Officer Id" },
  { key: "method", label: "Method" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Inspeksi CUI", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "inspection_number",
    label: "Inspection Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "cui_asset_id",
    label: "Cui Asset Id",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "ship_id",
    label: "Ship Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "inspection_date",
    label: "Inspection Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "inspector_officer_id",
    label: "Inspector Officer Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "method",
    label: "Method",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "condition_rating",
    label: "Condition Rating",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "findings",
    label: "Findings",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "remedial_action_required",
    label: "Remedial Action Required",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "next_inspection_date",
    label: "Next Inspection Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.inspection_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
