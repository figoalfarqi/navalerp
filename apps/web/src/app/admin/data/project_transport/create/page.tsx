"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildProjectTransportPayload,
  projectTransportEndpoint,
  projectTransportFields,
  projectTransportInitialData,
} from "../_config";

export default function CreateProjectTransportPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Project Transport"
      url={projectTransportEndpoint}
      mode="add"
      initialData={projectTransportInitialData}
      fields={projectTransportFields("add")}
      buildPayload={(data) => buildProjectTransportPayload(data)}
      onSuccess={() => router.push("/admin/data/project_transport")}
    />
  );
}
