"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildTruckPayload,
  truckEndpoint,
  truckFields,
} from "../_config";

export default function CreateTruckPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Truck"
      url={truckEndpoint}
      mode="add"
      initialData={{ ownership_status_id: 1, is_active: 1 }}
      fields={truckFields("add")}
      buildPayload={buildTruckPayload}
      onSuccess={() => router.push("/admin/data/truck")}
    />
  );
}
