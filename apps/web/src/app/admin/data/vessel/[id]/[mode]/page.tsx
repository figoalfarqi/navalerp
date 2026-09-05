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
  buildVesselPayload,
  vesselEndpoint,
  vesselFields,
  vesselPrimaryKey,
} from "../../_config";

export default function VesselRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: vesselEndpoint,
    id,
    mode,
    primaryKey: vesselPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Kapal")}
      url={mode === "edit" ? `${vesselEndpoint}/${id}` : vesselEndpoint}
      mode={mode}
      initialData={record}
      fields={vesselFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildVesselPayload}
      onSuccess={() => router.push("/admin/data/vessel")}
    />
  );
}
