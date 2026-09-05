"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  appUserFields,
  appUserInitialData,
  buildAppUserPayload,
} from "../../app_user/_config";
import { checkerUserEndpoint, checkerUserListPath } from "../_config";

export default function CreateCheckerUserPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Checker"
      url={checkerUserEndpoint}
      mode="add"
      initialData={{ ...appUserInitialData, app_role_id: 2 }}
      fields={appUserFields("add", 2, 2)}
      buildPayload={(data) => buildAppUserPayload(data)}
      onSuccess={() => router.push(checkerUserListPath)}
    />
  );
}
