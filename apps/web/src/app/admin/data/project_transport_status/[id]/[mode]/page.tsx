"use client";

import { useParams, useRouter } from "next/navigation";
import AdminRecordState from "@/components/admin/crud/AdminRecordState";
import {
  crudTitle,
  type AdminCrudMode,
} from "@/components/admin/crud/adminCrud";
import { useAdminRecord } from "@/components/admin/crud/useAdminRecord";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildProjectTransportStatusPayload,
  projectTransportStatusEndpoint,
  projectTransportStatusFields,
  projectTransportStatusPrimaryKey,
} from "../../_config";

export default function ProjectTransportStatusRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: projectTransportStatusEndpoint,
    id,
    mode,
    primaryKey: projectTransportStatusPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Status Transport")}
      url={
        mode === "edit"
          ? `${projectTransportStatusEndpoint}/${id}`
          : projectTransportStatusEndpoint
      }
      mode={mode}
      initialData={record}
      fields={projectTransportStatusFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={(data) => buildProjectTransportStatusPayload(data)}
      onSuccess={() => router.push("/admin/data/project_transport_status")}
    />
  );
}
