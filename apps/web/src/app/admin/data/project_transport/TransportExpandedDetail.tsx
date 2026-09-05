"use client";

import { useAdminRecord } from "@/components/admin/crud/useAdminRecord";
import TransportActivityDetail from "./TransportActivityDetail";
import {
  projectTransportEndpoint,
  projectTransportPrimaryKey,
} from "./_config";

export default function TransportExpandedDetail({
  transportId,
}: {
  transportId: number;
}) {
  const { record, error } = useAdminRecord({
    endpoint: projectTransportEndpoint,
    id: String(transportId),
    mode: "view",
    primaryKey: projectTransportPrimaryKey,
  });

  if (error) {
    return (
      <div className="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-700">
        {error}
      </div>
    );
  }

  if (!record) {
    return (
      <div className="rounded-lg border border-slate-200 bg-white px-4 py-5 text-sm text-slate-500">
        Memuat status dan foto transport...
      </div>
    );
  }

  return <TransportActivityDetail record={record} embedded />;
}
