"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildVesselPayload,
  vesselEndpoint,
  vesselFields,
} from "../_config";

export default function CreateVesselPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Kapal"
      url={vesselEndpoint}
      mode="add"
      initialData={{ is_active: 1 }}
      fields={vesselFields("add")}
      buildPayload={buildVesselPayload}
      onSuccess={() => router.push("/admin/data/vessel")}
    />
  );
}
