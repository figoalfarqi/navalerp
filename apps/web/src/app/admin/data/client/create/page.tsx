"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildClientPayload,
  clientEndpoint,
  clientFields,
} from "../_config";

export default function CreateClientPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Client"
      url={clientEndpoint}
      mode="add"
      initialData={{ is_active: 1 }}
      fields={clientFields("add")}
      buildPayload={buildClientPayload}
      onSuccess={() => router.push("/admin/data/client")}
    />
  );
}
