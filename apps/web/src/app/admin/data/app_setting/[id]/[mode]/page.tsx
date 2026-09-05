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
  appSettingEndpoint,
  appSettingFields,
  appSettingPrimaryKey,
  buildAppSettingPayload,
  normalizeAppSettingRecord,
} from "../../_config";

export default function AppSettingRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: appSettingEndpoint,
    id,
    mode,
    primaryKey: appSettingPrimaryKey,
    normalize: normalizeAppSettingRecord,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Pengaturan Aplikasi")}
      url={
        mode === "edit"
          ? `${appSettingEndpoint}/${id}`
          : appSettingEndpoint
      }
      mode={mode}
      initialData={record}
      fields={appSettingFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={(data) => buildAppSettingPayload(data)}
      onSuccess={() => router.push("/admin/data/app_setting")}
    />
  );
}
