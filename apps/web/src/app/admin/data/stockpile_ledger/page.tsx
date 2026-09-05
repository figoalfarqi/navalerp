"use client";

import Table from "@/components/table/Table";
import {
  stockpileLedgerColumns,
  stockpileLedgerEndpoint,
  stockpileLedgerFilters,
  stockpileLedgerTableName,
} from "./_config";

export default function StockpileLedgerListPage() {
  return (
    <Table
      url={stockpileLedgerEndpoint}
      title="Ledger Stockpile"
      table_name={stockpileLedgerTableName}
      columns={stockpileLedgerColumns}
      filters={stockpileLedgerFilters}
      isCreateable={false}
      readOnly
    />
  );
}
