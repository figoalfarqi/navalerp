"use client";

import Table from "@/components/table/Table";
import {
  cargoTypeColumns,
  cargoTypeEndpoint,
  cargoTypeFilters,
  cargoTypeTableName,
} from "./_config";

export default function CargoTypeListPage() {
  return (
    <Table
      url={cargoTypeEndpoint}
      title="Jenis Muatan"
      table_name={cargoTypeTableName}
      columns={cargoTypeColumns}
      filters={cargoTypeFilters}
    />
  );
}
