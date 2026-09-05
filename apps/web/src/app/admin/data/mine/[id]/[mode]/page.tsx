/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import FormCrud, { FormField } from "@/components/formCrud/FormCrud";
import { useFetchAPI } from "@/hooks/useFetchAPI";

export default function MineFormPage() {
  const { id, mode } = useParams() as {
    id: string;
    mode: "edit" | "copy" | "view";
  };
  const [initialData, setInitialData] = useState<any>(null);
  const router = useRouter();

  const { getAPI } = useFetchAPI();

  useEffect(() => {
    const fetchMine = async () => {
      try {
        const res = await getAPI<any>(
          `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/mine/${id}`,
          { authToken: "admin" }
        );
        const mine = res.data.data || res.data;

        if (mode === "copy") {
          delete mine.mine_id;
        }
        setInitialData(mine);
      } catch (err) {
        console.error("Failed to fetch Mine", err);
      }
    };

    fetchMine();
  }, [getAPI, id, mode]);

  if (!initialData) return <p>Loading...</p>;

  const url =
    mode === "edit"
      ? `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/mine/${id}`
      : `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/mine`;

  const fields: FormField[] = [
    {
      name: "mine_name",
      label: "Mine Name",
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
      disabled: mode === "view",
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
      name: "mine_address",
      label: "Address",
      fieldType: "textarea",
      col: "left",
      required: true,
      disabled: mode === "view",
    },
    {
      name: "mine_latitude",
      label: "Latitude",
      fieldType: "number",
      col: "right",
      required: false,
      disabled: mode === "view",
      validation: {
        lt: 90,
        gt: -90,
      },
    },
    {
      name: "mine_longitude",
      label: "Longitude",
      fieldType: "number",
      col: "right",
      required: false,
      disabled: mode === "view",
      validation: {
        lt: 180,
        gt: -180,
      },
    },
    {
      name: "mine_map_url",
      label: "Google Maps URL",
      fieldType: "text",
      col: "right",
      required: false,
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
  ];

  return (
    <FormCrud
      title={
        mode === "edit"
          ? "Edit Mine"
          : mode === "copy"
          ? "Copy Mine"
          : mode === "view"
          ? "View Mine"
          : "Create Mine"
      }
      url={url}
      mode={mode}
      initialData={initialData}
      fields={fields}
      hideSubmit={mode === "view"}
      onSuccess={() => {
        router.push("/admin/data/mine");
      }}
    />
  );
}
