"use client";

import Table from "@/components/table/Table";
import {
  vendorTypeColumns,
  vendorTypeEndpoint,
  vendorTypeFilters,
  vendorTypeTableName,
} from "./_config";

export default function VendorTypeListPage() {
  return (
    <Table
      url={vendorTypeEndpoint}
      title="Jenis Vendor"
      table_name={vendorTypeTableName}
      columns={vendorTypeColumns}
      filters={vendorTypeFilters}
    />
  );
}
