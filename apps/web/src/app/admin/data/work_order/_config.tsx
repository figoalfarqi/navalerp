/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";
import { DetailItemsConfig } from "@/components/formCrud/DetailItemsTable";

export const entityName = "work_order";
export const entityTitle = "Perintah Kerja MRO";
export const entityEndpoint = "/admin/work_order";
export const primaryKey = "work_order_id";

export const columns: ColumnField[] = [
  { key: "failure_report_number", label: "Laporan Kerusakan" },
  { key: "pm_title", label: "Jadwal PM" },
  { key: "equipment_name", label: "Peralatan" },
  { key: "work_order_number", label: "Work Order Number" },
  { key: "work_order_type", label: "Work Order Type" },
  { key: "priority", label: "Priority" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Perintah Kerja MRO", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "failure_report_id",
    col: "left",
    label: "Laporan Kerusakan",
    fieldType: "select",
    options: {
      url: "/admin/failure_report?limit=100",
      labelKey: "report_number",
      valueKey: "report_id",
    },
    required: false,
    disabled: mode === "view",
  },
    {
    name: "pm_schedule_id",
    col: "right",
    label: "Jadwal PM",
    fieldType: "select",
    options: {
      url: "/admin/pm_schedule?limit=100",
      labelKey: "pm_title",
      valueKey: "pm_id",
    },
    required: false,
    disabled: mode === "view",
  },
    {
    name: "equipment_id",
    col: "left",
    label: "Peralatan",
    fieldType: "select",
    options: {
      url: "/admin/equipment?limit=100",
      labelKey: "equipment_name",
      valueKey: "equipment_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "work_order_number",
    col: "right",
    label: "Work Order Number",
    fieldType: "text",
    required: true,
    readOnly: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "WO",
  },
  {
    name: "work_order_type",
    col: "left",
    label: "Work Order Type",
    fieldType: "select",
    options: [
      { label: "Corrective", value: "CORRECTIVE" },
      { label: "Preventive", value: "PREVENTIVE" },
      { label: "Docking", value: "DOCKING" },
      { label: "Depot Level", value: "DEPOT_LEVEL" },
      { label: "Emergency", value: "EMERGENCY" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "priority",
    col: "right",
    label: "Priority",
    fieldType: "select",
    options: [
      { label: "Emergency", value: "EMERGENCY" },
      { label: "Urgent", value: "URGENT" },
      { label: "Routine", value: "ROUTINE" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "scheduled_start_date",
    col: "left",
    label: "Scheduled Start Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "scheduled_end_date",
    col: "right",
    label: "Scheduled End Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_start_date",
    col: "left",
    label: "Actual Start Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_end_date",
    col: "right",
    label: "Actual End Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
    {
    name: "lead_engineer_user_id",
    col: "left",
    label: "Insinyur Utama",
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
    name: "assigned_facility",
    col: "right",
    label: "Assigned Facility",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "left",
    label: "Status",
    fieldType: "select",
    options: [
      { label: "Draft", value: "DRAFT" },
      { label: "Approved", value: "APPROVED" },
      { label: "In Progress", value: "IN_PROGRESS" },
      { label: "Waiting Parts", value: "WAITING_PARTS" },
      { label: "Completed", value: "COMPLETED" },
      { label: "Inspected", value: "INSPECTED" },
      { label: "Cancelled", value: "CANCELLED" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "total_labor_hours",
    col: "right",
    label: "Total Labor Hours",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "estimated_cost",
    col: "left",
    label: "Estimated Cost",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_cost",
    col: "right",
    label: "Actual Cost",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "completion_notes",
    col: "left",
    label: "Completion Notes",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const detailItemsConfig: DetailItemsConfig = {
  tableName: "items",
  title: "Rincian Suku Cadang & Material Pemeliharaan",
  itemName: "Suku Cadang",
  qtyKey: "quantity_required",
  priceKey: "unit_cost",
  totalKey: "total_cost",
  columns: [
    {
      key: "material_id",
      label: "Material / Suku Cadang",
      type: "select",
      required: true,
      placeholder: "Pilih Material...",
      options: {
        url: "/admin/material?limit=100",
        labelKey: "material_name",
        valueKey: "material_id",
        extraLabelKey: "material_code",
        priceKey: "standard_cost",
      },
    },
    {
      key: "quantity_required",
      label: "Kuantitas Dibutuhkan",
      type: "number",
      required: true,
      defaultValue: 1,
      width: "150px",
    },
    {
      key: "quantity_issued",
      label: "Kuantitas Dikeluarkan",
      type: "number",
      required: false,
      defaultValue: 0,
      width: "150px",
    },
    {
      key: "unit_cost",
      label: "Biaya Satuan",
      type: "currency",
      required: false,
      defaultValue: 0,
      width: "160px",
    },
    {
      key: "is_critical_spare",
      label: "Suku Cadang Kritis?",
      type: "select",
      required: false,
      defaultValue: false,
      options: [
        { label: "Ya (Critical Spare)", value: true },
        { label: "Tidak (Standard Spare)", value: false },
      ],
      width: "170px",
    },
  ],
};

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.work_order_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  delete payload.failure_report_number;
  delete payload.pm_title;
  delete payload.equipment_name;
  delete payload.full_name;

  if (Array.isArray(payload.items)) {
    payload.items = payload.items.map((it: any) => ({
      material_id: it.material_id,
      quantity_required: Number(it.quantity_required) || 0,
      quantity_issued: Number(it.quantity_issued) || 0,
      unit_cost: Number(it.unit_cost) || 0,
      total_cost: (Number(it.quantity_required) || 0) * (Number(it.unit_cost) || 0),
      is_critical_spare: Boolean(it.is_critical_spare),
    }));
  }

  return payload;
};
