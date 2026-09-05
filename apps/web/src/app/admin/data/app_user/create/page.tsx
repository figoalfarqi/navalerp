"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  appUserEndpoint,
  appUserFields,
  appUserInitialData,
  buildAppUserPayload,
} from "../_config";

export default function CreateAppUserPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Pengguna"
      url={appUserEndpoint}
      mode="add"
      initialData={appUserInitialData}
      fields={appUserFields("add")}
      buildPayload={(data) => buildAppUserPayload(data)}
      onSuccess={() => router.push("/admin/data/app_user")}
    />
  );
}
