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
  buildVesselCargoPayload,
  vesselCargoEndpoint,
  vesselCargoFields,
  vesselCargoPrimaryKey,
} from "../../_config";

export default function VesselCargoRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: vesselCargoEndpoint,
    id,
    mode,
    primaryKey: vesselCargoPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Muatan Kapal")}
      url={
        mode === "edit"
          ? `${vesselCargoEndpoint}/${id}`
          : vesselCargoEndpoint
      }
      mode={mode}
      initialData={record}
      fields={vesselCargoFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildVesselCargoPayload}
      onSuccess={() => router.push("/admin/data/vessel_cargo")}
    />
  );
}
