"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import { useAuth } from "@/context/AuthContext";
import {
  appUserFields,
  appUserInitialData,
  buildAppUserPayload,
} from "../../app_user/_config";
import { adminUserEndpoint, adminUserListPath } from "../_config";

export default function CreateAdminUserPage() {
  const router = useRouter();
  const { adminPayload } = useAuth();
  const isNormalAdmin = adminPayload.app_role_id === 6;
  const excludedRoleIds = adminPayload.app_role_id === 4 ? [5] : [];
  return (
    <FormCrud
      title="Tambah Pengguna Admin"
      url={adminUserEndpoint}
      mode="add"
      initialData={{
        ...appUserInitialData,
        ...(isNormalAdmin ? { app_role_id: 6 } : {}),
      }}
      fields={appUserFields(
        "add",
        3,
        isNormalAdmin ? 6 : undefined,
        excludedRoleIds,
      )}
      buildPayload={(data) => buildAppUserPayload(data)}
      onSuccess={() => router.push(adminUserListPath)}
    />
  );
}
