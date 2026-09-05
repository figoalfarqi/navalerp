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
  buildVendorTypePayload,
  vendorTypeEndpoint,
  vendorTypeFields,
  vendorTypePrimaryKey,
} from "../../_config";

export default function VendorTypeRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: vendorTypeEndpoint,
    id,
    mode,
    primaryKey: vendorTypePrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Jenis Vendor")}
      url={mode === "edit" ? `${vendorTypeEndpoint}/${id}` : vendorTypeEndpoint}
      mode={mode}
      initialData={record}
      fields={vendorTypeFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildVendorTypePayload}
      onSuccess={() => router.push("/admin/data/vendor_type")}
    />
  );
}
