/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import FormCrud from "@/components/formCrud/FormCrud";
import { useFetchAPI } from "@/hooks/useFetchAPI";

export default function CityFormPage() {
  const { id, mode } = useParams() as {
    id: string;
    mode: "edit" | "copy" | "view";
  };
  const [initialData, setInitialData] = useState<any>(null);
  const router = useRouter();

  const { getAPI } = useFetchAPI();

  useEffect(() => {
    const fetchCity = async () => {
      try {
        const res = await getAPI<any>(
          `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city/${id}`,
          { authToken: "admin" }
        );

        const city = res.data.data || res.data;

        if (mode === "copy") {
          delete city.city_id;
        }
        setInitialData(city);
      } catch (err) {
        console.error("Failed to fetch city", err);
      }
    };
    fetchCity();
  }, [getAPI, id, mode]);

  if (!initialData) return <p>Loading...</p>;

  const url =
    mode === "edit"
      ? `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city/${id}`
      : `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city`;

  return (
    <FormCrud
      title={
        mode === "edit"
          ? "Edit City"
          : mode === "copy"
          ? "Copy City"
          : "View City"
      }
      url={url}
      mode={mode}
      initialData={initialData}
      fields={[
        {
          name: "province_id",
          label: "Province",
          fieldType: "select",
          col: "left",
          required: true,
          disabled: mode === "view",
          options: {
            url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/province?limit=99`,
            labelKey: "province_name",
            valueKey: "province_id",
            autoSelectFirst: false,
          },
        },
        {
          name: "city_name",
          label: "City Name",
          fieldType: "text",
          col: "right",
          required: true,
          disabled: mode === "view",
        },
        {
          name: "is_active",
          label: "Is Active",
          fieldType: "boolean",
          col: "left",
          required: true,
          disabled: mode === "view",
        },
      ]}
      hideSubmit={mode === "view"}
      onSuccess={() => {
        router.push("/admin/data/city");
      }}
    />
  );
}
