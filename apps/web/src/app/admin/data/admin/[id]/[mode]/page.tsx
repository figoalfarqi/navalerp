"use client";

import { useParams, useRouter } from "next/navigation";
import AdminRecordState from "@/components/admin/crud/AdminRecordState";
import {
  crudTitle,
  type AdminCrudMode,
} from "@/components/admin/crud/adminCrud";
import { useAdminRecord } from "@/components/admin/crud/useAdminRecord";
import FormCrud from "@/components/formCrud/FormCrud";
import { useAuth } from "@/context/AuthContext";
import {
  appUserFields,
  appUserPrimaryKey,
  buildAppUserPayload,
} from "../../../app_user/_config";
import { adminUserEndpoint, adminUserListPath } from "../../_config";

export default function AdminUserRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { adminPayload } = useAuth();
  const isNormalAdmin = adminPayload.app_role_id === 6;
  const excludedRoleIds = adminPayload.app_role_id === 4 ? [5] : [];
  const { record, error } = useAdminRecord({
    endpoint: adminUserEndpoint,
    id,
    mode,
    primaryKey: appUserPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Pengguna Admin")}
      url={
        mode === "edit"
          ? `${adminUserEndpoint}/${id}`
          : adminUserEndpoint
      }
      mode={mode}
      initialData={record}
      fields={appUserFields(
        mode,
        3,
        isNormalAdmin ? 6 : undefined,
        excludedRoleIds,
      )}
      hideSubmit={mode === "view"}
      buildPayload={(data) => buildAppUserPayload(data)}
      onSuccess={() => router.push(adminUserListPath)}
    />
  );
}
