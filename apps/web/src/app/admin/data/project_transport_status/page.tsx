"use client";

import Table from "@/components/table/Table";
import {
  projectTransportStatusColumns,
  projectTransportStatusEndpoint,
  projectTransportStatusFilters,
  projectTransportStatusTableName,
} from "./_config";

export default function ProjectTransportStatusListPage() {
  return (
    <Table
      url={projectTransportStatusEndpoint}
      title="Status Project Transport"
      table_name={projectTransportStatusTableName}
      columns={projectTransportStatusColumns}
      filters={projectTransportStatusFilters}
    />
  );
}
