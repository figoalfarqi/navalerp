"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  appUserFields,
  appUserInitialData,
  buildAppUserPayload,
} from "../../app_user/_config";
import { driverUserEndpoint, driverUserListPath } from "../_config";

export default function CreateDriverUserPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Driver"
      url={driverUserEndpoint}
      mode="add"
      initialData={{ ...appUserInitialData, app_role_id: 1 }}
      fields={appUserFields("add", 1, 1)}
      buildPayload={(data) => buildAppUserPayload(data)}
      onSuccess={() => router.push(driverUserListPath)}
    />
  );
}
