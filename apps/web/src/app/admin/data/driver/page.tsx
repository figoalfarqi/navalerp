"use client";

import Table from "@/components/table/Table";
import {
  appUserColumns,
  appUserFilters,
  appUserTableName,
} from "../app_user/_config";
import { driverUserEndpoint } from "./_config";

export default function DriverUserListPage() {
  return (
    <Table
      url={driverUserEndpoint}
      title="Pengguna Driver"
      table_name={appUserTableName}
      table_url="driver"
      columns={appUserColumns}
      filters={appUserFilters}
      default_filter_values={{ app_role_type_id: 1 }}
    />
  );
}
