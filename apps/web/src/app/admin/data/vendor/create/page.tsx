"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildVendorPayload,
  vendorEndpoint,
  vendorFields,
} from "../_config";

export default function CreateVendorPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Vendor"
      url={vendorEndpoint}
      mode="add"
      initialData={{ is_active: 1 }}
      fields={vendorFields("add")}
      buildPayload={buildVendorPayload}
      onSuccess={() => router.push("/admin/data/vendor")}
    />
  );
}
