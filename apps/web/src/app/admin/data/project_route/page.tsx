"use client";

import Table from "@/components/table/Table";
import {
  projectRouteColumns,
  projectRouteEndpoint,
  projectRouteFilters,
  projectRouteTableName,
} from "./_config";

export default function ProjectRouteListPage() {
  return (
    <Table
      url={projectRouteEndpoint}
      title="Daftar Rute Project"
      table_name={projectRouteTableName}
      columns={projectRouteColumns}
      filters={projectRouteFilters}
    />
  );
}
