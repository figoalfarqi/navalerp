"use client";

import Table from "@/components/table/Table";
import {
  appRoleColumns,
  appRoleEndpoint,
  appRoleFilters,
  appRoleTableName,
} from "./_config";

export default function AppRoleListPage() {
  return (
    <Table
      url={appRoleEndpoint}
      title="Role Aplikasi"
      table_name={appRoleTableName}
      columns={appRoleColumns}
      filters={appRoleFilters}
    />
  );
}
