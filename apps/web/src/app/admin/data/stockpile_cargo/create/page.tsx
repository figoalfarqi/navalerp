"use client";

import { useRouter } from "next/navigation";
import FormCrud from "@/components/formCrud/FormCrud";
import {
  buildStockpileCargoPayload,
  stockpileCargoEndpoint,
  stockpileCargoFields,
} from "../_config";

export default function CreateStockpileCargoPage() {
  const router = useRouter();
  return (
    <FormCrud
      title="Tambah Stok Muatan Stockpile"
      url={stockpileCargoEndpoint}
      mode="add"
      initialData={{
        current_volume_cubic: 0,
        current_weight_ton: 0,
        is_active: 1,
      }}
      fields={stockpileCargoFields("add")}
      buildPayload={buildStockpileCargoPayload}
      onSuccess={() => router.push("/admin/data/stockpile_cargo")}
    />
  );
}
