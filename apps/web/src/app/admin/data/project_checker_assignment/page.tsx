"use client";

import Table from "@/components/table/Table";
import {
  checkerAssignmentColumns,
  checkerAssignmentEndpoint,
  checkerAssignmentFilters,
  checkerAssignmentTableName,
} from "./_config";

export default function ProjectCheckerAssignmentListPage() {
  return (
    <Table
      url={checkerAssignmentEndpoint}
      title="Penugasan Checker per Project"
      table_name={checkerAssignmentTableName}
      columns={checkerAssignmentColumns}
      filters={checkerAssignmentFilters}
    />
  );
}
