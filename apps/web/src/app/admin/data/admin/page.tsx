"use client";

import Table from "@/components/table/Table";
import { useAuth } from "@/context/AuthContext";
import {
  appUserColumns,
  appUserFilters,
  appUserTableName,
} from "../app_user/_config";
import { adminUserEndpoint } from "./_config";

export default function AdminUserListPage() {
  const { adminPayload } = useAuth();
  const isNormalAdmin = adminPayload.app_role_id === 6;

  return (
    <Table
      url={adminUserEndpoint}
      title="Pengguna Admin"
      table_name={appUserTableName}
      table_url="admin"
      columns={appUserColumns}
      filters={appUserFilters}
      default_filter_values={{
        app_role_type_id: 3,
        ...(isNormalAdmin ? { app_role_id: 6 } : {}),
      }}
    />
  );
}
