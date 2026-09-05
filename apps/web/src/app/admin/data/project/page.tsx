"use client";

import Table from "@/components/table/Table";
import {
  projectColumns,
  projectEndpoint,
  projectFilters,
  projectTableName,
} from "./_config";

export default function ProjectListPage() {
  return (
    <Table
      url={projectEndpoint}
      title="Daftar Project"
      table_name={projectTableName}
      columns={projectColumns}
      filters={projectFilters}
    />
  );
}
