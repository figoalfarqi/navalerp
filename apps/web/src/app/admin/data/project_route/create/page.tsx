"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildProjectRoutePayload,
  projectRouteEndpoint,
  projectRouteFields,
  projectRouteInitialData,
} from "../_config";

export default function CreateProjectRoutePage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Rute Project"
      url={projectRouteEndpoint}
      mode="add"
      initialData={projectRouteInitialData}
      fields={projectRouteFields("add")}
      buildPayload={(data) => buildProjectRoutePayload(data)}
      onSuccess={() => router.push("/admin/data/project_route")}
    />
  );
}
