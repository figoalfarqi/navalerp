"use client";
import { FilterField } from "@/components/table/FilterFormTable";
/* eslint-disable @typescript-eslint/no-explicit-any */
import Table, { ColumnField, TableProps } from "@/components/table/Table";
import { formatDateTime } from "@/utils/dateTime";

const columns: ColumnField[] = [
  {
    key: "stockpile_name",
    label: "Stockpile Name",
    columnLength: 200,
    sortable: "string",
  },
  {
    key: "city_id",
    render: (item: any) => item.city?.city_name || "-",
    label: "City",
    columnLength: 200,
    sortable: "table_key",
  },
  {
    key: "stockpile_address",
    label: "Address",
    columnLength: 300,
    sortable: "string",
  },
  {
    key: "stockpile_latitude",
    label: "Latitude",
    columnLength: 300,
    sortable: "string",
  },
  {
    key: "stockpile_longitude",
    label: "Longitude",
    columnLength: 300,
    sortable: "string",
  },
  {
    key: "stockpile_map_url",
    label: "Map URL",
    columnLength: 300,
    sortable: "string",
  },
  {
    key: "is_active",
    label: "Is Active",
    render: (val) =>
      val.is_active == "1" ? (
        <div className="text-green-500">Aktif</div>
      ) : (
        <div className="text-red-500">Tidak Aktif</div>
      ),
    columnLength: 200,
    sortable: "boolean",
  },
  {
    key: "created_at",
    label: "Created At",
    render: (val) => formatDateTime(val.created_at),
    columnLength: 200,
    sortable: "date",
  },
  {
    key: "updated_at",
    label: "Updated At",
    render: (val) => formatDateTime(val.updated_at),
    columnLength: 200,
    sortable: "date",
  },
];

const filters: FilterField[] = [
  {
    name: "stockpile_name",
    label: "Stockpile Name",
    fieldType: "text",
    col: "right",
    
  },
];

export const createTablePropsStockpile = ({
  default_filter_values = {},
}: {
  default_filter_values?: Record<string, string | number | string[]>;
}): TableProps => ({
  url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/stockpile`,
  title: "Client Destination List",
  table_name: "stockpile",
  columns: columns,
  filters: filters,
  default_filter_values: default_filter_values,
  opendata: "stockpile_map_url",
});

export default function StockpilePage() {
  return <Table {...createTablePropsStockpile({})} />;
}
