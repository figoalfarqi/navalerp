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
  buildProjectTransportPhotoPayload,
  projectTransportPhotoEndpoint,
  projectTransportPhotoFields,
  projectTransportPhotoPrimaryKey,
} from "../../_config";

export default function ProjectTransportPhotoRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: projectTransportPhotoEndpoint,
    id,
    mode,
    primaryKey: projectTransportPhotoPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Foto Transport")}
      url={
        mode === "edit"
          ? `${projectTransportPhotoEndpoint}/${id}`
          : projectTransportPhotoEndpoint
      }
      mode={mode}
      initialData={record}
      fields={projectTransportPhotoFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={(data) => buildProjectTransportPhotoPayload(data)}
      onSuccess={() => router.push("/admin/data/project_transport_photo")}
    />
  );
}
