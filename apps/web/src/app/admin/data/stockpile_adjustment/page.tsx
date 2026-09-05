"use client";

import Table from "@/components/table/Table";
import {
  stockpileAdjustmentColumns,
  stockpileAdjustmentEndpoint,
  stockpileAdjustmentFilters,
  stockpileAdjustmentTableName,
} from "./_config";

export default function StockpileAdjustmentListPage() {
  return (
    <Table
      url={stockpileAdjustmentEndpoint}
      title="Penyesuaian Stok"
      table_name={stockpileAdjustmentTableName}
      columns={stockpileAdjustmentColumns}
      filters={stockpileAdjustmentFilters}
    />
  );
}
