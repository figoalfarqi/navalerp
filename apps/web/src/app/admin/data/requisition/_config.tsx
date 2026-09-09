/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";
import { DetailItemsConfig } from "@/components/formCrud/DetailItemsTable";

export const entityName = "requisition";
export const entityTitle = "Permintaan Pengadaan";
export const entityEndpoint = "/admin/requisition";
export const primaryKey = "requisition_id";

export const columns: ColumnField[] = [
  { key: "requisition_number", label: "Requisition Number" },
  { key: "unit_name", label: "Satuan Pengaju" },
  { key: "work_order_number", label: "Perintah Kerja (WO)" },
  { key: "priority", label: "Priority" },
  { key: "requested_date", label: "Requested Date", render: (item: any) => formatSmartDate(item.requested_date) },
  { key: "required_by_date", label: "Required By Date", render: (item: any) => formatSmartDate(item.required_by_date) },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Permintaan Pengadaan", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "requisition_number",
    col: "left",
    label: "Requisition Number",
    fieldType: "text",
    required: true,
    readOnly: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "PR",
  },
    {
    name: "origin_unit_id",
    col: "right",
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
    col: "left",
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
    col: "right",
    label: "Priority",
    fieldType: "select",
    options: [
      { label: "Emergency", value: "EMERGENCY" },
      { label: "High", value: "HIGH" },
      { label: "Regular", value: "REGULAR" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "requested_date",
    col: "left",
    label: "Requested Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "required_by_date",
    col: "right",
    label: "Required By Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "approval_status",
    col: "left",
    label: "Approval Status",
    fieldType: "select",
    options: [
      { label: "Pending", value: "PENDING" },
      { label: "Approved", value: "APPROVED" },
      { label: "Rejected", value: "REJECTED" },
    ],
    required: false,
    disabled: mode === "view",
  },
    {
    name: "approved_by_user_id",
    col: "right",
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
    col: "left",
    label: "Approved At",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "total_estimated_cost",
    col: "right",
    label: "Total Estimated Cost",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "justification",
    col: "left",
    label: "Justification",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const detailItemsConfig: DetailItemsConfig = {
  tableName: "items",
  title: "Rincian Item Permintaan Material",
  itemName: "Item Material",
  qtyKey: "quantity",
  priceKey: "estimated_unit_price",
  totalKey: "estimated_total_price",
  syncHeaderTotalKey: "total_estimated_cost",
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
      key: "quantity",
      label: "Jumlah (Qty)",
      type: "number",
      required: true,
      defaultValue: 1,
      width: "130px",
    },
    {
      key: "estimated_unit_price",
      label: "Estimasi Harga Satuan",
      type: "currency",
      required: true,
      defaultValue: 0,
      width: "170px",
    },
    {
      key: "notes",
      label: "Catatan Kebutuhan",
      type: "text",
      required: false,
      placeholder: "Catatan spesifikasi...",
    },
  ],
};

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.requisition_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  delete payload.unit_name;
  delete payload.work_order_number;
  delete payload.full_name;

  if (payload.total_estimated_cost !== undefined) payload.total_estimated_cost = Number(payload.total_estimated_cost) || 0;

  if (Array.isArray(payload.items)) {
    payload.items = payload.items.map((it: any) => ({
      material_id: it.material_id,
      quantity: Number(it.quantity) || 0,
      estimated_unit_price: Number(it.estimated_unit_price) || 0,
      estimated_total_price: Number(it.estimated_total_price) || (Number(it.quantity) || 0) * (Number(it.estimated_unit_price) || 0),
      notes: it.notes || null,
    }));
  }

  return payload;
};
