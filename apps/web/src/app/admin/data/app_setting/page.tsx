"use client";

import Table from "@/components/table/Table";
import {
  appSettingColumns,
  appSettingEndpoint,
  appSettingFilters,
  appSettingTableName,
} from "./_config";

export default function AppSettingListPage() {
  return (
    <Table
      url={appSettingEndpoint}
      title="Pengaturan Aplikasi"
      table_name={appSettingTableName}
      columns={appSettingColumns}
      filters={appSettingFilters}
    />
  );
}
