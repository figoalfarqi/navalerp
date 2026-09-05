"use client";

import Table from "@/components/table/Table";
import {
  vesselColumns,
  vesselEndpoint,
  vesselFilters,
  vesselTableName,
} from "./_config";

export default function VesselListPage() {
  return (
    <Table
      url={vesselEndpoint}
      title="Kapal"
      table_name={vesselTableName}
      columns={vesselColumns}
      filters={vesselFilters}
    />
  );
}
