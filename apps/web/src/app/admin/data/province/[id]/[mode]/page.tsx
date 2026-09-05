/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import FormCrud from "@/components/formCrud/FormCrud";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { useQueryParams } from "@/hooks/useQueryParams";

export default function ProvinceFormPage() {
  const { id, mode } = useParams() as {
    id: string;
    mode: "edit" | "copy" | "view";
  };
  const [initialData, setInitialData] = useState<any>(null);
  const { searchParams } = useQueryParams();

  const router = useRouter();
  const { getAPI } = useFetchAPI();

  useEffect(() => {
    const fetchProvince = async () => {
      try {
        const res = await getAPI<any>(
          `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/province/${id}`,
          { authToken: "admin" },
        );
        const province = res.data.data || res.data;

        if (mode === "copy") {
          delete province.province_id;
        }
        setInitialData(province);
      } catch (err) {
        console.error("Failed to fetch Province", err);
      }
    };
    fetchProvince();
  }, [getAPI, id, mode]);

  if (!initialData) return <p>Loading...</p>;

  const url =
    mode === "edit"
      ? `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/province/${id}`
      : `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/province`;

  return (
    <FormCrud
      title={
        mode === "edit"
          ? "Edit Province"
          : mode === "copy"
            ? "Copy Province"
            : "View Province"
      }
      url={url}
      mode={mode}
      initialData={initialData}
      fields={[
        {
          name: "province_name",
          label: "Province Name",
          fieldType: "text",
          col: "left",
          required: true,
          disabled: mode === "view",
        },
        {
          name: "province_real_name",
          label: "Province Real Name",
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
        router.push(`/admin/data/province?${searchParams}`);
      }}
    />
  );
}
