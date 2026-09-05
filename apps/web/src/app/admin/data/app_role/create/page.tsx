"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  appRoleEndpoint,
  appRoleFields,
  buildAppRolePayload,
} from "../_config";

export default function CreateAppRolePage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Role"
      url={appRoleEndpoint}
      mode="add"
      initialData={{ is_active: 1 }}
      fields={appRoleFields("add")}
      buildPayload={buildAppRolePayload}
      onSuccess={() => router.push("/admin/data/app_role")}
    />
  );
}
