"use client";

import Table from "@/components/table/Table";
import {
  clientDestinationColumns,
  clientDestinationEndpoint,
  clientDestinationFilters,
  clientDestinationTableName,
} from "./_config";

export default function ClientDestinationListPage() {
  return (
    <Table
      url={clientDestinationEndpoint}
      title="Tujuan Client"
      table_name={clientDestinationTableName}
      columns={clientDestinationColumns}
      filters={clientDestinationFilters}
      opendata="client_destination_map_url"
    />
  );
}
