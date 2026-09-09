/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";
import { DetailItemsConfig } from "@/components/formCrud/DetailItemsTable";

export const entityName = "stock_transfer";
export const entityTitle = "Transfer Bebekal";
export const entityEndpoint = "/admin/stock_transfer";
export const primaryKey = "transfer_id";

export const columns: ColumnField[] = [
  { key: "transfer_number", label: "Transfer Number" },
  { key: "source_warehouse_name", label: "Gudang Asal" },
  { key: "dest_warehouse_name", label: "Gudang Tujuan" },
  { key: "movement_type", label: "Movement Type" },
  { key: "scheduled_departure", label: "Scheduled Departure" },
  { key: "actual_departure", label: "Actual Departure" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Transfer Bebekal", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "transfer_number",
    col: "left",
    label: "Transfer Number",
    fieldType: "text",
    required: true,
    readOnly: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "TRF",
  },
    {
    name: "from_warehouse_id",
    col: "right",
    label: "Gudang Asal",
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
    name: "to_warehouse_id",
    col: "left",
    label: "Gudang Tujuan",
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
    name: "movement_type",
    col: "right",
    label: "Movement Type",
    fieldType: "select",
    options: [
      { label: "Fleet Resupply", value: "FLEET_RESUPPLY" },
      { label: "Base Transfer", value: "BASE_TRANSFER" },
      { label: "Depot Dispatch", value: "DEPOT_DISPATCH" },
      { label: "Emergency Airdrop", value: "EMERGENCY_AIRDROP" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "scheduled_departure",
    col: "left",
    label: "Scheduled Departure",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_departure",
    col: "right",
    label: "Actual Departure",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "scheduled_arrival",
    col: "left",
    label: "Scheduled Arrival",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "actual_arrival",
    col: "right",
    label: "Actual Arrival",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "transporter_unit",
    col: "left",
    label: "Transporter Unit",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "right",
    label: "Status",
    fieldType: "select",
    options: [
      { label: "Planned", value: "PLANNED" },
      { label: "In Transit", value: "IN_TRANSIT" },
      { label: "Received", value: "RECEIVED" },
      { label: "Rejected", value: "REJECTED" },
    ],
    required: false,
    disabled: mode === "view",
  },
];

export const detailItemsConfig: DetailItemsConfig = {
  tableName: "items",
  title: "Rincian Item Mutasi / Transfer",
  itemName: "Item Bebekal",
  qtyKey: "quantity_shipped",
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
      },
    },
    {
      key: "quantity_shipped",
      label: "Jumlah Dikirim",
      type: "number",
      required: true,
      defaultValue: 1,
      width: "150px",
    },
    {
      key: "quantity_received",
      label: "Jumlah Diterima",
      type: "number",
      required: false,
      defaultValue: 0,
      width: "150px",
    },
    {
      key: "condition_on_receipt",
      label: "Kondisi",
      type: "select",
      required: false,
      defaultValue: "SERVICEABLE",
      options: [
        { label: "Siap Pakai (Serviceable)", value: "SERVICEABLE" },
        { label: "Rusak Ringan (Repairable)", value: "REPAIRABLE" },
        { label: "Tidak Layak Pakai (Unserviceable)", value: "UNSERVICEABLE" },
      ],
      width: "200px",
    },
  ],
};

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.transfer_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  delete payload.from_warehouse_name;
  delete payload.to_warehouse_name;

  if (Array.isArray(payload.items)) {
    payload.items = payload.items.map((it: any) => ({
      material_id: it.material_id,
      quantity_shipped: Number(it.quantity_shipped) || 0,
      quantity_received: Number(it.quantity_received) || 0,
      condition_on_receipt: it.condition_on_receipt || "SERVICEABLE",
    }));
  }

  return payload;
};
