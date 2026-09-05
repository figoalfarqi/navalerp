"use client";

import Table from "@/components/table/Table";
import {
  vendorColumns,
  vendorEndpoint,
  vendorFilters,
  vendorTableName,
} from "./_config";

export default function VendorListPage() {
  return (
    <Table
      url={vendorEndpoint}
      title="Vendor"
      table_name={vendorTableName}
      columns={vendorColumns}
      filters={vendorFilters}
    />
  );
}
