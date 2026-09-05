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
  buildTruckPayload,
  truckEndpoint,
  truckFields,
  truckPrimaryKey,
} from "../../_config";

export default function TruckRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: truckEndpoint,
    id,
    mode,
    primaryKey: truckPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Truck")}
      url={mode === "edit" ? `${truckEndpoint}/${id}` : truckEndpoint}
      mode={mode}
      initialData={record}
      fields={truckFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildTruckPayload}
      onSuccess={() => router.push("/admin/data/truck")}
    />
  );
}
