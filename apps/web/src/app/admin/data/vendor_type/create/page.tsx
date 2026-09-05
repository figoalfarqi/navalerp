"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildVendorTypePayload,
  vendorTypeEndpoint,
  vendorTypeFields,
} from "../_config";

export default function CreateVendorTypePage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Jenis Vendor"
      url={vendorTypeEndpoint}
      mode="add"
      initialData={{ is_active: 1 }}
      fields={vendorTypeFields("add")}
      buildPayload={buildVendorTypePayload}
      onSuccess={() => router.push("/admin/data/vendor_type")}
    />
  );
}
