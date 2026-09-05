/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import FormCrud from "@/components/formCrud/FormCrud";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { useQueryParams } from "@/hooks/useQueryParams";

export default function TruckMerkFormPage() {
  const { id, mode } = useParams() as {
    id: string;
    mode: "edit" | "copy" | "view";
  };
  const [initialData, setInitialData] = useState<any>(null);
  const { searchParams } = useQueryParams();

  const router = useRouter();
  const { getAPI } = useFetchAPI();

  useEffect(() => {
    const fetchTruckMerk = async () => {
      try {
        const res = await getAPI<any>(
          `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck_merk/${id}`,
          { authToken: "admin" }
        );
        const truck_merk = res.data.data || res.data;

        if (mode === "copy") {
          delete truck_merk.truck_merk_id;
        }
        setInitialData(truck_merk);
      } catch (err) {
        console.error("Failed to fetch Truck Merk", err);
      }
    };
    fetchTruckMerk();
  }, [getAPI, id, mode]);

  if (!initialData) return <p>Loading...</p>;

   
  const url =
    mode === "edit"
      ? `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck_merk/${id}`
      : `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck_merk`;

  return (
    <FormCrud
      title={
        mode === "edit"
          ? "Edit Truck Merk"
          : mode === "copy"
          ? "Copy Truck Merk"
          : "View Truck Merk"
      }
      url={url}
      mode={mode}
      initialData={initialData}
      fields={[
        {
          name: "truck_merk_name",
          label: "Truck Merk Name",
          fieldType: "text",
          col: "left",
          required: true,
          disabled: mode === "view",
        },
        {
          name: "is_active",
          label: "Is Active",
          fieldType: "boolean",
          col: "right",
          required: true,
          disabled: mode === "view",
        },
      ]}
      hideSubmit={mode === "view"}
      onSuccess={() => {
        router.push(`/admin/data/truck_merk?${searchParams}`);
      }}
    />
  );
}
