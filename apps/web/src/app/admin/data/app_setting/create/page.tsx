"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  appSettingEndpoint,
  appSettingFields,
  buildAppSettingPayload,
} from "../_config";

export default function CreateAppSettingPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Pengaturan Aplikasi"
      url={appSettingEndpoint}
      mode="add"
      initialData={{ app_setting_value: "{}", is_active: 1 }}
      fields={appSettingFields("add")}
      buildPayload={(data) => buildAppSettingPayload(data)}
      onSuccess={() => router.push("/admin/data/app_setting")}
    />
  );
}
