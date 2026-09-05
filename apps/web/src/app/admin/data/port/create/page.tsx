"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import { buildPortPayload, portEndpoint, portFields } from "../_config";

export default function CreatePortPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Pelabuhan"
      url={portEndpoint}
      mode="add"
      initialData={{ is_active: 1 }}
      fields={portFields("add")}
      buildPayload={buildPortPayload}
      onSuccess={() => router.push("/admin/data/port")}
    />
  );
}
