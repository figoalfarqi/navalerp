/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "military_rank";
export const entityTitle = "Pangkat Militer";
export const entityEndpoint = "/admin/military_rank";
export const primaryKey = "rank_id";

export const columns: ColumnField[] = [
  { key: "rank_code", label: "Rank Code" },
  { key: "rank_name", label: "Rank Name" },
  { key: "rank_category", label: "Rank Category" },
  { key: "nato_rank_code", label: "Nato Rank Code" },
  { key: "seniority_order", label: "Seniority Order" },
  { key: "is_active", label: "Is Active" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Pangkat Militer", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "rank_code",
    label: "Rank Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "rank_name",
    label: "Rank Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "rank_category",
    label: "Rank Category",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "nato_rank_code",
    label: "Nato Rank Code",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "seniority_order",
    label: "Seniority Order",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "is_active",
    label: "Is Active",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.rank_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
