"use client";

import { useParams, useRouter } from "next/navigation";
import AdminRecordState from "@/components/admin/crud/AdminRecordState";
import {
  crudTitle,
  type AdminCrudMode,
} from "@/components/admin/crud/adminCrud";
import { useAdminRecord } from "@/components/admin/crud/useAdminRecord";
import FormCrud from "@/components/formCrud/FormCrud";
import TransportActivityDetail from "../../TransportActivityDetail";
import {
  buildProjectTransportPayload,
  projectTransportEndpoint,
  projectTransportFields,
  projectTransportPrimaryKey,
} from "../../_config";

export default function ProjectTransportRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: projectTransportEndpoint,
    id,
    mode,
    primaryKey: projectTransportPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <>
      <FormCrud
        title={crudTitle(mode, "Project Transport")}
        url={
          mode === "edit"
            ? `${projectTransportEndpoint}/${id}`
            : projectTransportEndpoint
        }
        mode={mode}
        initialData={record}
        fields={projectTransportFields(mode)}
        hideSubmit={mode === "view"}
        buildPayload={(data) => buildProjectTransportPayload(data)}
        onSuccess={() => router.push("/admin/data/project_transport")}
      />
      {mode === "view" && <TransportActivityDetail record={record} />}
    </>
  );
}
