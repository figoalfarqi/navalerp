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
  buildStockpileCargoPayload,
  stockpileCargoEndpoint,
  stockpileCargoFields,
  stockpileCargoPrimaryKey,
} from "../../_config";

export default function StockpileCargoRecordPage() {
  const { id, mode } = useParams<{
    id: string;
    mode: Exclude<AdminCrudMode, "add">;
  }>();
  const router = useRouter();
  const { record, error } = useAdminRecord({
    endpoint: stockpileCargoEndpoint,
    id,
    mode,
    primaryKey: stockpileCargoPrimaryKey,
  });
  if (!record) return <AdminRecordState error={error} />;

  return (
    <FormCrud
      title={crudTitle(mode, "Stok Muatan Stockpile")}
      url={
        mode === "edit"
          ? `${stockpileCargoEndpoint}/${id}`
          : stockpileCargoEndpoint
      }
      mode={mode}
      initialData={record}
      fields={stockpileCargoFields(mode)}
      hideSubmit={mode === "view"}
      buildPayload={buildStockpileCargoPayload}
      onSuccess={() => router.push("/admin/data/stockpile_cargo")}
    />
  );
}
