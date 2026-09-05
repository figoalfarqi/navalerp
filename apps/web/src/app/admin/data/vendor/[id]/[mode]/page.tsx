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
  buildVendorPayload,
  vendorEndpoint,
  vendorFields,
  vendorPrimaryKey,
} from "../../_config";

export default function VendorRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: vendorEndpoint,
    id,
    mode,
    primaryKey: vendorPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Vendor")}
      url={mode === "edit" ? `${vendorEndpoint}/${id}` : vendorEndpoint}
      mode={mode}
      initialData={record}
      fields={vendorFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildVendorPayload}
      onSuccess={() => router.push("/admin/data/vendor")}
    />
  );
}
