"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildCheckerAssignmentPayload,
  checkerAssignmentEndpoint,
  checkerAssignmentFields,
} from "../_config";

export default function CreateProjectCheckerAssignmentPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Penugasan Checker"
      url={checkerAssignmentEndpoint}
      mode="add"
      initialData={{
        access_started_at: new Date().toISOString(),
        is_default: 0,
        is_active: 1,
      }}
      fields={checkerAssignmentFields("add")}
      buildPayload={(data) => buildCheckerAssignmentPayload(data)}
      onSuccess={() =>
        router.push("/admin/data/project_checker_assignment")
      }
    />
  );
}
