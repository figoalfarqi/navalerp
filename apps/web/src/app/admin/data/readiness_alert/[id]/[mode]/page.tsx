"use client";

import { useParams, useRouter } from "next/navigation";
import AdminRecordState from "@/components/admin/crud/AdminRecordState";
import { crudTitle, type AdminCrudMode } from "@/components/admin/crud/adminCrud";
import { useAdminRecord } from "@/components/admin/crud/useAdminRecord";
import FormCrud from "@/components/formCrud/FormCrud";
import { buildPayload, entityEndpoint, entityTitle, formFields, primaryKey } from "../../_config";

export default function ReadinessAlertDetailPage() {
  const { id, mode } = useParams<{ id: string; mode: Exclude<AdminCrudMode, "add"> }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: entityEndpoint,
    id,
    mode,
    primaryKey,
  });

  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, entityTitle)}
      url={mode === "edit" ? `${entityEndpoint}/${id}` : entityEndpoint}
      mode={mode}
      initialData={record}
      fields={formFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildPayload}
      onSuccess={() => router.push(`/admin/data/${entityEndpoint.replace('/admin/', '')}`)}
    />
  );
}
