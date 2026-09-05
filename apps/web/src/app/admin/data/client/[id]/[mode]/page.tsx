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
  buildClientPayload,
  clientEndpoint,
  clientFields,
  clientPrimaryKey,
  normalizeClientRecord,
} from "../../_config";

export default function ClientRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: clientEndpoint,
    id,
    mode,
    primaryKey: clientPrimaryKey,
    normalize: normalizeClientRecord,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Client")}
      url={mode === "edit" ? `${clientEndpoint}/${id}` : clientEndpoint}
      mode={mode}
      initialData={record}
      fields={clientFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildClientPayload}
      onSuccess={() => router.push("/admin/data/client")}
    />
  );
}
