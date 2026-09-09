/* eslint-disable @typescript-eslint/no-explicit-any */
import { formatSmartDate } from "@/utils/dateTime";
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";
import { DetailItemsConfig } from "@/components/formCrud/DetailItemsTable";

export const entityName = "shipment";
export const entityTitle = "Pengiriman Konvoi";
export const entityEndpoint = "/admin/shipment";
export const primaryKey = "shipment_id";

export const columns: ColumnField[] = [
  { key: "manifest_number", label: "Manifest Number" },
  { key: "route_name", label: "Rute Logistik" },
  { key: "transport_unit_name", label: "Unit Angkut" },
  { key: "origin_warehouse_name", label: "Gudang Asal" },
  { key: "destination_warehouse_name", label: "Gudang Tujuan" },
  { key: "departure_date", label: "Departure Date", render: (item: any) => formatSmartDate(item.departure_date) },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Pengiriman Konvoi", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "manifest_number",
    col: "left",
    label: "Manifest Number",
    fieldType: "text",
    required: true,
    readOnly: true,
    disabled: mode === "view",
    autoGenerate: true,
    autoPrefix: "MNF",
  },
    {
    name: "route_id",
    col: "right",
    label: "Rute Logistik",
    fieldType: "select",
    options: {
      url: "/admin/route?limit=100",
      labelKey: "route_name",
      valueKey: "route_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "transport_unit_id",
    col: "left",
    label: "Unit Angkut",
    fieldType: "select",
    options: {
      url: "/admin/transport_unit?limit=100",
      labelKey: "unit_code",
      valueKey: "transport_unit_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "origin_warehouse_id",
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
    name: "destination_warehouse_id",
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
    name: "departure_date",
    col: "right",
    label: "Departure Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "arrival_date",
    col: "left",
    label: "Arrival Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "escort_security_level",
    col: "right",
    label: "Escort Security Level",
    fieldType: "select",
    options: [
      { label: "Unescorted", value: "UNESCORTED" },
      { label: "Standard Convoy", value: "STANDARD_CONVOY" },
      { label: "Warship Escort", value: "WARSHIP_ESCORT" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "status",
    col: "left",
    label: "Status",
    fieldType: "select",
    options: [
      { label: "Planned", value: "PLANNED" },
      { label: "Loading", value: "LOADING" },
      { label: "In Transit", value: "IN_TRANSIT" },
      { label: "Delivered", value: "DELIVERED" },
      { label: "Cancelled", value: "CANCELLED" },
    ],
    required: false,
    disabled: mode === "view",
  },
    {
    name: "authorized_by_user_id",
    col: "right",
    label: "Diotorisasi Oleh",
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
  title: "Rincian Item Muatan Pengiriman",
  itemName: "Item Muatan",
  qtyKey: "quantity_dispatched",
  columns: [
    {
      key: "material_id",
      label: "Material / Muatan",
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
      key: "quantity_dispatched",
      label: "Jumlah Dikirim",
      type: "number",
      required: true,
      defaultValue: 1,
      width: "140px",
    },
    {
      key: "quantity_received",
      label: "Jumlah Diterima",
      type: "number",
      required: false,
      defaultValue: 0,
      width: "140px",
    },
    {
      key: "packaging_type",
      label: "Jenis Kemasan",
      type: "select",
      required: true,
      defaultValue: "CRATE",
      options: [
        { label: "Peti Kayu (Crate)", value: "CRATE" },
        { label: "Palet (Pallet)", value: "PALLET" },
        { label: "Drum", value: "DRUM" },
        { label: "Kotak Amunisi (Ammo Box)", value: "AMMO_BOX" },
        { label: "Kontainer (ISO Container)", value: "ISO_CONTAINER" },
      ],
      width: "180px",
    },
    {
      key: "weight_kg",
      label: "Berat (Kg)",
      type: "number",
      required: false,
      defaultValue: 0,
      width: "130px",
    },
    {
      key: "notes",
      label: "Catatan Muatan",
      type: "text",
      required: false,
      placeholder: "Keterangan kemasan/muatan...",
    },
  ],
};

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.shipment_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  delete payload.route_code;
  delete payload.transport_code;
  delete payload.origin_warehouse_name;
  delete payload.destination_warehouse_name;
  delete payload.full_name;

  if (Array.isArray(payload.items)) {
    payload.items = payload.items.map((it: any) => ({
      material_id: it.material_id,
      quantity_dispatched: Number(it.quantity_dispatched) || 0,
      quantity_received: Number(it.quantity_received) || 0,
      packaging_type: it.packaging_type || "CRATE",
      weight_kg: Number(it.weight_kg) || 0,
      notes: it.notes || null,
    }));
  }

  return payload;
};
