"use client";

import Table from "@/components/table/Table";
import {
  projectTransportPhotoColumns,
  projectTransportPhotoEndpoint,
  projectTransportPhotoFilters,
  projectTransportPhotoTableName,
} from "./_config";

export default function ProjectTransportPhotoListPage() {
  return (
    <Table
      url={projectTransportPhotoEndpoint}
      title="Foto Project Transport"
      table_name={projectTransportPhotoTableName}
      columns={projectTransportPhotoColumns}
      filters={projectTransportPhotoFilters}
      opendata="photo_url"
    />
  );
}
