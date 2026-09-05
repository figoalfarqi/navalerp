"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildFinancialTransactionPayload,
  financialTransactionEndpoint,
  financialTransactionFields,
  financialTransactionInitialData,
} from "../_config";

export default function CreateFinancialTransactionPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Transaksi Finansial"
      url={financialTransactionEndpoint}
      mode="add"
      initialData={financialTransactionInitialData}
      fields={financialTransactionFields("add")}
      buildPayload={(data) => buildFinancialTransactionPayload(data)}
      onSuccess={() =>
        router.push("/admin/data/project_financial_transaction")
      }
    />
  );
}
