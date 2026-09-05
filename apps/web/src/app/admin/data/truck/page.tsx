"use client";

import Table from "@/components/table/Table";
import {
  truckColumns,
  truckEndpoint,
  truckFilters,
  truckTableName,
} from "./_config";

export default function TruckListPage() {
  return (
    <Table
      url={truckEndpoint}
      title="Truck"
      table_name={truckTableName}
      columns={truckColumns}
      filters={truckFilters}
    />
  );
}
