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
  appRoleEndpoint,
  appRoleFields,
  appRolePrimaryKey,
  buildAppRolePayload,
} from "../../_config";

export default function AppRoleRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: appRoleEndpoint,
    id,
    mode,
    primaryKey: appRolePrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Role")}
      url={mode === "edit" ? `${appRoleEndpoint}/${id}` : appRoleEndpoint}
      mode={mode}
      initialData={record}
      fields={appRoleFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildAppRolePayload}
      onSuccess={() => router.push("/admin/data/app_role")}
    />
  );
}
