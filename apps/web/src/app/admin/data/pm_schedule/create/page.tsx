"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import { buildPayload, entityEndpoint, entityTitle, formFields } from "../_config";

export default function PmScheduleCreatePage() {
  const router = useRouter();
  return (
    <FormCrud
      title={`Tambah ${entityTitle}`}
      url={entityEndpoint}
      mode="add"
      fields={formFields("add")}
      buildPayload={buildPayload}
      onSuccess={() => router.push(`/admin/data/${entityEndpoint.replace('/admin/', '')}`)}
    />
  );
}
