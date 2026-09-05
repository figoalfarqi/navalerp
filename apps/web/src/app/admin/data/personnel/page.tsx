"use client";

import Table from "@/components/table/Table";
import { columns, entityEndpoint, entityName, entityTitle, filterFields, primaryKey } from "./_config";

export default function PersonnelListPage() {
  return (
    <Table
      title={entityTitle}
      url={entityEndpoint}
      table_name={entityName}
      table_url={`/admin/data/${entityName}`}
      columns={columns}
      primaryKey={primaryKey}
      filters={filterFields}
    />
  );
}
