"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildStockpileAdjustmentPayload,
  stockpileAdjustmentEndpoint,
  stockpileAdjustmentFields,
} from "../_config";

const today = () => new Date().toISOString().slice(0, 10);

export default function CreateStockpileAdjustmentPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Penyesuaian Stok"
      url={stockpileAdjustmentEndpoint}
      mode="add"
      initialData={{
        adjustment_date: today(),
        reference_type_id: 3,
        amount_volume_cubic: 0,
        amount_weight_ton: 0,
      }}
      fields={stockpileAdjustmentFields("add")}
      buildPayload={buildStockpileAdjustmentPayload}
      onSuccess={() => router.push("/admin/data/stockpile_adjustment")}
    />
  );
}
