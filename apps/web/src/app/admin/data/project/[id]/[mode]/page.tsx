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
  buildProjectPayload,
  projectEndpoint,
  projectFields,
  projectPrimaryKey,
} from "../../_config";

export default function ProjectRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: projectEndpoint,
    id,
    mode,
    primaryKey: projectPrimaryKey,
  });

  if (!record) return <AdminRecordState error={error} />;
  const submitUrl =
    mode === "edit" ? `${projectEndpoint}/${id}` : projectEndpoint;

  return (
    <FormCrud
      title={crudTitle(mode, "Project")}
      url={submitUrl}
      mode={mode}
      initialData={record}
      fields={projectFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={(data) => buildProjectPayload(data)}
      onSuccess={() => router.push("/admin/data/project")}
    />
  );
}
