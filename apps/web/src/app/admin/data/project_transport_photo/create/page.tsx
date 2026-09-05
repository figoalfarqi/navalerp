"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildProjectTransportPhotoPayload,
  projectTransportPhotoEndpoint,
  projectTransportPhotoFields,
} from "../_config";

export default function CreateProjectTransportPhotoPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Foto Transport"
      url={projectTransportPhotoEndpoint}
      mode="add"
      initialData={{ photo_type_id: 2 }}
      fields={projectTransportPhotoFields("add")}
      buildPayload={(data) => buildProjectTransportPhotoPayload(data)}
      onSuccess={() => router.push("/admin/data/project_transport_photo")}
    />
  );
}
