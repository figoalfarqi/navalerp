/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";
import { DetailItemsConfig } from "@/components/formCrud/DetailItemsTable";

export const entityName = "purchase_order";
export const entityTitle = "Purchase Order";
export const entityEndpoint = "/admin/purchase_order";
export const primaryKey = "po_id";

export const columns: ColumnField[] = [
  { key: "po_number", label: "Po Number" },
  { key: "contract_number", label: "No. Kontrak" },
  { key: "vendor_name", label: "Vendor" },
  { key: "unit_name", label: "Satuan Penerbit" },
  { key: "order_date", label: "Order Date", render: (item: any) => formatSmartDate(item.order_date) },
  { key: "delivery_deadline", label: "Delivery Deadline" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Purchase Order", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "po_number",
    col: "left",
    label: "Po Number",
    fieldType: "text",
    required: true,
    readOnly: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "PO",
  },
    {
    name: "contract_id",
    col: "right",
    label: "No. Kontrak",
    fieldType: "select",
    options: {
      url: "/admin/contract?limit=100",
      labelKey: "contract_number",
      valueKey: "contract_id",
    },
    required: false,
    disabled: mode === "view",
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
    name: "issuing_unit_id",
    col: "right",
    label: "Satuan Penerbit",
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
    name: "order_date",
    col: "left",
    label: "Order Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "delivery_deadline",
    col: "right",
    label: "Delivery Deadline",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
    {
    name: "destination_warehouse_id",
    col: "left",
    label: "Gudang Tujuan",
    fieldType: "select",
    options: {
      url: "/admin/warehouse?limit=100",
      labelKey: "warehouse_name",
      valueKey: "warehouse_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "total_amount",
    col: "right",
    label: "Total Amount",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "tax_amount",
    col: "left",
    label: "Tax Amount",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "right",
    label: "Status",
    fieldType: "select",
    options: [
      { label: "Issued", value: "ISSUED" },
      { label: "Partial Received", value: "PARTIAL_RECEIVED" },
      { label: "Completed", value: "COMPLETED" },
      { label: "Cancelled", value: "CANCELLED" },
    ],
    required: false,
    disabled: mode === "view",
  },
];

export const detailItemsConfig: DetailItemsConfig = {
  tableName: "items",
  title: "Rincian Item Material / Suku Cadang",
  itemName: "Item Material",
  qtyKey: "quantity",
  priceKey: "unit_price",
  totalKey: "total_price",
  syncHeaderTotalKey: "total_amount",
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
      key: "unit_price",
      label: "Harga Satuan",
      type: "currency",
      required: true,
      defaultValue: 0,
      width: "160px",
    },
    {
      key: "notes",
      label: "Catatan / Spesifikasi",
      type: "text",
      required: false,
      placeholder: "Catatan item...",
    },
  ],
};

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.po_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  delete payload.contract_number;
  delete payload.vendor_name;
  delete payload.unit_name;
  delete payload.destination_warehouse_name;
  delete payload.grand_total;

  if (payload.total_amount !== undefined) payload.total_amount = Number(payload.total_amount) || 0;
  if (payload.tax_amount !== undefined) payload.tax_amount = Number(payload.tax_amount) || 0;

  if (Array.isArray(payload.items)) {
    payload.items = payload.items.map((it: any) => ({
      material_id: it.material_id,
      quantity: Number(it.quantity) || 0,
      unit_price: Number(it.unit_price) || 0,
      total_price: Number(it.total_price) || (Number(it.quantity) || 0) * (Number(it.unit_price) || 0),
      notes: it.notes || null,
    }));
  }

  return payload;
};
