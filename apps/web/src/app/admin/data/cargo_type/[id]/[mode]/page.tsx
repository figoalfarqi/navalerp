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
  buildCargoTypePayload,
  cargoTypeEndpoint,
  cargoTypeFields,
  cargoTypePrimaryKey,
} from "../../_config";

export default function CargoTypeRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: cargoTypeEndpoint,
    id,
    mode,
    primaryKey: cargoTypePrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Jenis Muatan")}
      url={mode === "edit" ? `${cargoTypeEndpoint}/${id}` : cargoTypeEndpoint}
      mode={mode}
      initialData={record}
      fields={cargoTypeFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildCargoTypePayload}
      onSuccess={() => router.push("/admin/data/cargo_type")}
    />
  );
}
