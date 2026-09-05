"use client";

import Table from "@/components/table/Table";
import {
  clientColumns,
  clientEndpoint,
  clientFilters,
  clientTableName,
} from "./_config";

export default function ClientListPage() {
  return (
    <Table
      url={clientEndpoint}
      title="Client"
      table_name={clientTableName}
      columns={clientColumns}
      filters={clientFilters}
    />
  );
}
