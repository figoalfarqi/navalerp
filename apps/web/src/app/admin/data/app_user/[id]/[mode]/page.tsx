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
  appUserEndpoint,
  appUserFields,
  appUserPrimaryKey,
  buildAppUserPayload,
} from "../../_config";

export default function AppUserRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: appUserEndpoint,
    id,
    mode,
    primaryKey: appUserPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Pengguna")}
      url={mode === "edit" ? `${appUserEndpoint}/${id}` : appUserEndpoint}
      mode={mode}
      initialData={record}
      fields={appUserFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={(data) => buildAppUserPayload(data)}
      onSuccess={() => router.push("/admin/data/app_user")}
    />
  );
}
