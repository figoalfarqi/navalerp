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
  buildTruckAssignmentPayload,
  truckAssignmentEndpoint,
  truckAssignmentFields,
  truckAssignmentPrimaryKey,
} from "../../_config";

export default function ProjectTruckAssignmentRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: truckAssignmentEndpoint,
    id,
    mode,
    primaryKey: truckAssignmentPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Penugasan Truck")}
      url={
        mode === "edit"
          ? `${truckAssignmentEndpoint}/${id}`
          : truckAssignmentEndpoint
      }
      mode={mode}
      initialData={record}
      fields={truckAssignmentFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={(data) => buildTruckAssignmentPayload(data)}
      onSuccess={() => router.push("/admin/data/project_truck_assignment")}
    />
  );
}
