/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "requisition";
export const entityTitle = "Permintaan Pengadaan";
export const entityEndpoint = "/admin/requisition";
export const primaryKey = "requisition_id";

export const columns: ColumnField[] = [
  { key: "requisition_number", label: "Requisition Number" },
  { key: "unit_name", label: "Satuan Pengaju" },
  { key: "work_order_number", label: "Perintah Kerja (WO)" },
  { key: "priority", label: "Priority" },
  { key: "requested_date", label: "Requested Date" },
  { key: "required_by_date", label: "Required By Date" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Permintaan Pengadaan", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "requisition_number",
    label: "Requisition Number",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "origin_unit_id",
    label: "Satuan Pengaju",
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
    name: "work_order_id",
    label: "Perintah Kerja (WO)",
    fieldType: "select",
    options: {
      url: "/admin/work_order?limit=100",
      labelKey: "work_order_number",
      valueKey: "work_order_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "priority",
    label: "Priority",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "requested_date",
    label: "Requested Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "required_by_date",
    label: "Required By Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "approval_status",
    label: "Approval Status",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
    {
    name: "approved_by_user_id",
    label: "Disetujui Oleh",
    fieldType: "select",
    options: {
      url: "/admin/sys_user?limit=100",
      labelKey: "full_name",
      valueKey: "user_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "approved_at",
    label: "Approved At",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "total_estimated_cost",
    label: "Total Estimated Cost",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "justification",
    label: "Justification",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.requisition_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
