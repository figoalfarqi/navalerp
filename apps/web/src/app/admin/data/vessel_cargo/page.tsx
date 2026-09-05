"use client";

import Table from "@/components/table/Table";
import {
  vesselCargoColumns,
  vesselCargoEndpoint,
  vesselCargoFilters,
  vesselCargoTableName,
} from "./_config";

export default function VesselCargoListPage() {
  return (
    <Table
      url={vesselCargoEndpoint}
      title="Muatan Kapal"
      table_name={vesselCargoTableName}
      columns={vesselCargoColumns}
      filters={vesselCargoFilters}
    />
  );
}
