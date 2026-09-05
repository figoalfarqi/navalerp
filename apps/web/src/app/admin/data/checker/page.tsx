"use client";

import Table from "@/components/table/Table";
import {
  appUserColumns,
  appUserFilters,
  appUserTableName,
} from "../app_user/_config";
import { checkerUserEndpoint } from "./_config";

export default function CheckerUserListPage() {
  return (
    <Table
      url={checkerUserEndpoint}
      title="Pengguna Checker"
      table_name={appUserTableName}
      table_url="checker"
      columns={appUserColumns}
      filters={appUserFilters}
      default_filter_values={{ app_role_type_id: 2 }}
    />
  );
}
