"use client";

import Table from "@/components/table/Table";
import {
  projectTransportColumns,
  projectTransportEndpoint,
  projectTransportFilters,
  projectTransportTableName,
} from "./_config";
import TransportExpandedDetail from "./TransportExpandedDetail";

export default function ProjectTransportListPage() {
  return (
    <Table
      url={projectTransportEndpoint}
      title="Daftar Project Transport"
      table_name={projectTransportTableName}
      columns={projectTransportColumns}
      filters={projectTransportFilters}
      renderExpandedRow={(row) => (
        <TransportExpandedDetail
          transportId={Number(row.project_transport_id)}
        />
      )}
    />
  );
}
