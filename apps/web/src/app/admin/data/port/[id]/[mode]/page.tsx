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
  buildPortPayload,
  portEndpoint,
  portFields,
  portPrimaryKey,
} from "../../_config";

export default function PortRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: portEndpoint,
    id,
    mode,
    primaryKey: portPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Pelabuhan")}
      url={mode === "edit" ? `${portEndpoint}/${id}` : portEndpoint}
      mode={mode}
      initialData={record}
      fields={portFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildPortPayload}
      onSuccess={() => router.push("/admin/data/port")}
    />
  );
}
