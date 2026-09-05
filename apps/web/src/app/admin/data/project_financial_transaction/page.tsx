"use client";

import Table from "@/components/table/Table";
import {
  financialTransactionColumns,
  financialTransactionEndpoint,
  financialTransactionFilters,
  financialTransactionTableName,
} from "./_config";

export default function FinancialTransactionListPage() {
  return (
    <Table
      url={financialTransactionEndpoint}
      title="Transaksi Finansial Project"
      table_name={financialTransactionTableName}
      columns={financialTransactionColumns}
      filters={financialTransactionFilters}
    />
  );
}
