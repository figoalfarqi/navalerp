"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildTruckAssignmentPayload,
  truckAssignmentEndpoint,
  truckAssignmentFields,
} from "../_config";

export default function CreateProjectTruckAssignmentPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Penugasan Truck"
      url={truckAssignmentEndpoint}
      mode="add"
      initialData={{
        assignment_started_at: new Date().toISOString(),
        is_active: 1,
      }}
      fields={truckAssignmentFields("add")}
      buildPayload={(data) => buildTruckAssignmentPayload(data)}
      onSuccess={() => router.push("/admin/data/project_truck_assignment")}
    />
  );
}
