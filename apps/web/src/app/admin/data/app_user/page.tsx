"use client";

import Table from "@/components/table/Table";
import {
  appUserColumns,
  appUserEndpoint,
  appUserFilters,
  appUserTableName,
} from "./_config";

export default function AppUserListPage() {
  return (
    <Table
      url={appUserEndpoint}
      title="Semua Pengguna"
      table_name={appUserTableName}
      columns={appUserColumns}
      filters={appUserFilters}
    />
  );
}
