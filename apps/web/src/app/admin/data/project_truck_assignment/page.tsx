"use client";

import Table from "@/components/table/Table";
import {
  truckAssignmentColumns,
  truckAssignmentEndpoint,
  truckAssignmentFilters,
  truckAssignmentTableName,
} from "./_config";

export default function ProjectTruckAssignmentListPage() {
  return (
    <Table
      url={truckAssignmentEndpoint}
      title="Penugasan Truck per Project"
      table_name={truckAssignmentTableName}
      columns={truckAssignmentColumns}
      filters={truckAssignmentFilters}
    />
  );
}
