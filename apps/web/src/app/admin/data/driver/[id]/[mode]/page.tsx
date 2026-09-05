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
  appUserFields,
  appUserPrimaryKey,
  buildAppUserPayload,
} from "../../../app_user/_config";
import { driverUserEndpoint, driverUserListPath } from "../../_config";

export default function DriverUserRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: driverUserEndpoint,
    id,
    mode,
    primaryKey: appUserPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Driver")}
      url={
        mode === "edit"
          ? `${driverUserEndpoint}/${id}`
          : driverUserEndpoint
      }
      mode={mode}
      initialData={record}
      fields={appUserFields(mode, 1, 1)}
      hideSubmit={mode === "view"}
      buildPayload={(data) => buildAppUserPayload(data)}
      onSuccess={() => router.push(driverUserListPath)}
    />
  );
}
