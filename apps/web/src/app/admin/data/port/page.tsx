"use client";

import Table from "@/components/table/Table";
import {
  portColumns,
  portEndpoint,
  portFilters,
  portTableName,
} from "./_config";

export default function PortListPage() {
  return (
    <Table
      url={portEndpoint}
      title="Pelabuhan"
      table_name={portTableName}
      columns={portColumns}
      filters={portFilters}
    />
  );
}
