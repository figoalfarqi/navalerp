"use client";

import { useParams, useRouter } from "next/navigation";
import AdminRecordState from "@/components/admin/crud/AdminRecordState";
import {
  crudTitle,
  type AdminCrudMode,
} from "@/components/admin/crud/adminCrud";
import { useAdminRecord } from "@/components/admin/crud/useAdminRecord";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildFinancialTransactionPayload,
  financialTransactionEndpoint,
  financialTransactionFields,
  financialTransactionPrimaryKey,
} from "../../_config";

export default function FinancialTransactionRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: financialTransactionEndpoint,
    id,
    mode,
    primaryKey: financialTransactionPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Transaksi Finansial")}
      url={
        mode === "edit"
          ? `${financialTransactionEndpoint}/${id}`
          : financialTransactionEndpoint
      }
      mode={mode}
      initialData={record}
      fields={financialTransactionFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={(data) => buildFinancialTransactionPayload(data)}
      onSuccess={() =>
        router.push("/admin/data/project_financial_transaction")
      }
    />
  );
}
