"use client";

import Table from "@/components/table/Table";
import {
  stockpileCargoColumns,
  stockpileCargoEndpoint,
  stockpileCargoFilters,
  stockpileCargoTableName,
} from "./_config";

export default function StockpileCargoListPage() {
  return (
    <Table
      url={stockpileCargoEndpoint}
      title="Stok Muatan Stockpile"
      table_name={stockpileCargoTableName}
      columns={stockpileCargoColumns}
      filters={stockpileCargoFilters}
    />
  );
}
