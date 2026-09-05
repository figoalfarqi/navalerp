"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildProjectPayload,
  projectEndpoint,
  projectFields,
  projectInitialData,
} from "../_config";

export default function CreateProjectPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Project"
      url={projectEndpoint}
      mode="add"
      initialData={projectInitialData}
      fields={projectFields("add")}
      buildPayload={(data) => buildProjectPayload(data)}
      onSuccess={() => router.push("/admin/data/project")}
    />
  );
}
