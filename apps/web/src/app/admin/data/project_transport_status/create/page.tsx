"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildProjectTransportStatusPayload,
  projectTransportStatusEndpoint,
  projectTransportStatusFields,
} from "../_config";

export default function CreateProjectTransportStatusPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Status Transport"
      url={projectTransportStatusEndpoint}
      mode="add"
      initialData={{
        status_time: new Date().toISOString(),
        is_fraud: 0,
        is_active: 1,
      }}
      fields={projectTransportStatusFields("add")}
      buildPayload={(data) => buildProjectTransportStatusPayload(data)}
      onSuccess={() => router.push("/admin/data/project_transport_status")}
    />
  );
}
