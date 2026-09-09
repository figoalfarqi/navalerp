/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";
import { DetailItemsConfig } from "@/components/formCrud/DetailItemsTable";

export const entityName = "stock_adjustment";
export const entityTitle = "Penyesuaian Stok";
export const entityEndpoint = "/admin/stock_adjustment";
export const primaryKey = "adjustment_id";

export const columns: ColumnField[] = [
  { key: "warehouse_name", label: "Nama Gudang" },
  { key: "adjustment_number", label: "Adjustment Number" },
  { key: "adjustment_date", label: "Adjustment Date", render: (item: any) => formatSmartDate(item.adjustment_date) },
  { key: "conductor_name", label: "Pelaksana Stock Take" },
  { key: "reason", label: "Reason" },
  { key: "status", label: "Status" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Penyesuaian Stok", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "warehouse_id",
    col: "left",
    label: "Nama Gudang",
    fieldType: "select",
    options: {
      url: "/admin/warehouse?limit=100",
      labelKey: "warehouse_name",
      valueKey: "warehouse_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "adjustment_number",
    col: "right",
    label: "Adjustment Number",
    fieldType: "text",
    required: true,
    readOnly: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "ADJ",
  },
  {
    name: "adjustment_date",
    col: "left",
    label: "Adjustment Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "conducted_by_user_id",
    col: "right",
    label: "Pelaksana Stock Take",
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
    name: "reason",
    col: "left",
    label: "Reason",
    fieldType: "select",
    options: [
      { label: "Stock Opname Variance", value: "STOCK_OPNAME_VARIANCE" },
      { label: "Damage", value: "DAMAGE" },
      { label: "Expiry", value: "EXPIRY" },
      { label: "Salvage", value: "SALVAGE" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "right",
    label: "Status",
    fieldType: "select",
    options: [
      { label: "Draft", value: "DRAFT" },
      { label: "Approved", value: "APPROVED" },
      { label: "Cancelled", value: "CANCELLED" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "remarks",
    col: "left",
    label: "Remarks",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const detailItemsConfig: DetailItemsConfig = {
  tableName: "items",
  title: "Rincian Item Penyesuaian Stok",
  itemName: "Item Penyesuaian",
  qtyKey: "physical_quantity",
  priceKey: "unit_cost",
  totalKey: "total_adjustment_value",
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
      key: "book_quantity",
      label: "Kuantitas Sistem (Buku)",
      type: "number",
      required: true,
      defaultValue: 0,
      width: "150px",
    },
    {
      key: "physical_quantity",
      label: "Kuantitas Fisik Aktual",
      type: "number",
      required: true,
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
  ],
};

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.adjustment_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  delete payload.warehouse_name;
  delete payload.full_name;

  if (Array.isArray(payload.items)) {
    payload.items = payload.items.map((it: any) => {
      const bookQty = Number(it.book_quantity) || 0;
      const physQty = Number(it.physical_quantity) || 0;
      const unitCost = Number(it.unit_cost) || 0;
      const diffQty = physQty - bookQty;
      return {
        material_id: it.material_id,
        book_quantity: bookQty,
        physical_quantity: physQty,
        difference_quantity: diffQty,
        unit_cost: unitCost,
        total_adjustment_value: diffQty * unitCost,
      };
    });
  }

  return payload;
};
