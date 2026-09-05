/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import FormCrud, { FormField } from "@/components/formCrud/FormCrud";
import { useFetchAPI } from "@/hooks/useFetchAPI";

export default function StockpileFormPage() {
  const { id, mode } = useParams() as {
    id: string;
    mode: "edit" | "copy" | "view";
  };
  const [initialData, setInitialData] = useState<any>(null);
  const router = useRouter();

  const { getAPI } = useFetchAPI();

  useEffect(() => {
    const fetchStockpile = async () => {
      try {
        const res = await getAPI<any>(
          `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/stockpile/${id}`,
          { authToken: "admin" }
        );
        const stockpile = res.data.data || res.data;

        if (mode === "copy") {
          delete stockpile.stockpile_id;
        }
        setInitialData(stockpile);
      } catch (err) {
        console.error("Failed to fetch Stockpile", err);
      }
    };

    fetchStockpile();
  }, [getAPI, id, mode]);

  if (!initialData) return <p>Loading...</p>;

  const url =
    mode === "edit"
      ? `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/stockpile/${id}`
      : `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/stockpile`;

  const fields: FormField[] = [
    {
      name: "stockpile_name",
      label: "Stockpile Name",
      fieldType: "text",
      col: "left",
      required: true,
      disabled: mode === "view",
    },
    {
      fieldType: "chainselect",
      label: "Lokasi",
      name: "Lokasi",
      col: "left",
      required: false,
      chains: [
        {
          name: "city_id",
          label: "City",
          url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city?limit=399`,
          labelKey: "city_name",
          valueKey: "city_id",
        },
      ],
    },
    {
      name: "stockpile_address",
      label: "Address",
      fieldType: "textarea",
      col: "left",
      required: true,
      disabled: mode === "view",
    },
    {
      name: "stockpile_latitude",
      label: "Latitude",
      fieldType: "number",
      col: "right",
      required: false,
      validation: {
        lt: 90,
        gt: -90,
      },
    },
    {
      name: "stockpile_longitude",
      label: "Longitude",
      fieldType: "number",
      col: "right",
      required: false,
      validation: {
        lt: 180,
        gt: -180,
      },
    },
    {
      name: "stockpile_map_url",
      label: "Google Maps URL",
      fieldType: "text",
      col: "right",
      required: false,
    },
    {
      name: "is_active",
      label: "Is Active",
      fieldType: "boolean",
      col: "left",
      required: true,
    },
  ];

  return (
    <FormCrud
      title={
        mode === "edit"
          ? "Edit Stockpile"
          : mode === "copy"
          ? "Copy Stockpile"
          : mode === "view"
          ? "View Stockpile"
          : "Create Stockpile"
      }
      url={url}
      mode={mode}
      initialData={initialData}
      fields={fields}
      hideSubmit={mode === "view"}
      onSuccess={() => {
        router.push("/admin/data/stockpile");
      }}
    />
  );
}
