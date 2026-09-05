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
  buildStockpileAdjustmentPayload,
  stockpileAdjustmentEndpoint,
  stockpileAdjustmentFields,
  stockpileAdjustmentPrimaryKey,
} from "../../_config";

export default function StockpileAdjustmentRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: stockpileAdjustmentEndpoint,
    id,
    mode,
    primaryKey: stockpileAdjustmentPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Penyesuaian Stok")}
      url={
        mode === "edit"
          ? `${stockpileAdjustmentEndpoint}/${id}`
          : stockpileAdjustmentEndpoint
      }
      mode={mode}
      initialData={record}
      fields={stockpileAdjustmentFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildStockpileAdjustmentPayload}
      onSuccess={() => router.push("/admin/data/stockpile_adjustment")}
    />
  );
}
