"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildClientDestinationPayload,
  clientDestinationEndpoint,
  clientDestinationFields,
} from "../_config";

export default function CreateClientDestinationPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Tujuan Client"
      url={clientDestinationEndpoint}
      mode="add"
      initialData={{ is_active: 1 }}
      fields={clientDestinationFields("add")}
      buildPayload={buildClientDestinationPayload}
      onSuccess={() => router.push("/admin/data/client_destination")}
    />
  );
}
