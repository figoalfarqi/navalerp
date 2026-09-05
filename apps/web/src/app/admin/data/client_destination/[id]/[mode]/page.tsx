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
  buildClientDestinationPayload,
  clientDestinationEndpoint,
  clientDestinationFields,
  clientDestinationPrimaryKey,
  normalizeClientDestinationRecord,
} from "../../_config";

export default function ClientDestinationRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: clientDestinationEndpoint,
    id,
    mode,
    primaryKey: clientDestinationPrimaryKey,
    normalize: normalizeClientDestinationRecord,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Tujuan Client")}
      url={
        mode === "edit"
          ? `${clientDestinationEndpoint}/${id}`
          : clientDestinationEndpoint
      }
      mode={mode}
      initialData={record}
      fields={clientDestinationFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildClientDestinationPayload}
      onSuccess={() => router.push("/admin/data/client_destination")}
    />
  );
}
