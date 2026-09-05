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
  buildProjectRoutePayload,
  projectRouteEndpoint,
  projectRouteFields,
  projectRoutePrimaryKey,
} from "../../_config";

export default function ProjectRouteRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: projectRouteEndpoint,
    id,
    mode,
    primaryKey: projectRoutePrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Rute Project")}
      url={
        mode === "edit"
          ? `${projectRouteEndpoint}/${id}`
          : projectRouteEndpoint
      }
      mode={mode}
      initialData={record}
      fields={projectRouteFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={(data) => buildProjectRoutePayload(data)}
      onSuccess={() => router.push("/admin/data/project_route")}
    />
  );
}
