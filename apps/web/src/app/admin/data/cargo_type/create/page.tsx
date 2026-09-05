"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildCargoTypePayload,
  cargoTypeEndpoint,
  cargoTypeFields,
} from "../_config";

export default function CreateCargoTypePage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Jenis Muatan"
      url={cargoTypeEndpoint}
      mode="add"
      initialData={{ is_active: 1 }}
      fields={cargoTypeFields("add")}
      buildPayload={buildCargoTypePayload}
      onSuccess={() => router.push("/admin/data/cargo_type")}
    />
  );
}
