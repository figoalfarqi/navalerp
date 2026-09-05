"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildVesselCargoPayload,
  vesselCargoEndpoint,
  vesselCargoFields,
} from "../_config";

export default function CreateVesselCargoPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Muatan Kapal"
      url={vesselCargoEndpoint}
      mode="add"
      initialData={{ is_active: 1 }}
      fields={vesselCargoFields("add")}
      buildPayload={buildVesselCargoPayload}
      onSuccess={() => router.push("/admin/data/vessel_cargo")}
    />
  );
}
