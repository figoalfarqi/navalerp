/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import FormCrud from "@/components/formCrud/FormCrud";
import { useFetchAPI } from "@/hooks/useFetchAPI";

export default function TruckTypeFormPage() {
  const { id, mode } = useParams() as {
    id: string;
    mode: "edit" | "copy" | "view";
  };
  const [initialData, setInitialData] = useState<any>(null);
  const router = useRouter();
  const { getAPI } = useFetchAPI();

  useEffect(() => {
    const fetchTruckType = async () => {
      try {
        const res = await getAPI<any>(
          `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck_type/${id}`,
          { authToken: "admin" },
        );
        const truckType = res.data.data || res.data;

        if (mode === "copy") {
          delete truckType.truck_type_id;
        }

        setInitialData(truckType);
      } catch (err) {
        console.error("Failed to fetch Truck Type", err);
      }
    };
    fetchTruckType();
  }, [getAPI, id, mode]);

  if (!initialData) return <p>Loading...</p>;

  const url =
    mode === "edit"
      ? `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck_type/${id}`
      : `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck_type`;

  return (
    <FormCrud
      title={
        mode === "edit"
          ? "Edit Truck Type"
          : mode === "copy"
            ? "Copy Truck Type"
            : "View Truck Type"
      }
      url={url}
      mode={mode}
      initialData={initialData}
      fields={[
        {
          name: "truck_type_name",
          label: "Truck Type Name",
          fieldType: "text",
          col: "left",
          required: true,
          disabled: mode === "view",
        },
        {
          name: "truck_type_description",
          label: "Truck Type Description",
          fieldType: "textarea",
          col: "left",
          disabled: mode === "view",
        },
        {
          name: "truck_box_length",
          label: "Truck Box Length (m)",
          fieldType: "number",
          col: "right",
          required: true,
          disabled: mode === "view",
        },
        {
          name: "truck_box_width",
          label: "Truck Box Width (m)",
          fieldType: "number",
          col: "left",
          required: true,
          disabled: mode === "view",
        },
        {
          name: "truck_box_height",
          label: "Truck Box Height (m)",
          fieldType: "number",
          col: "right",
          required: true,
          disabled: mode === "view",
        },
        {
          name: "truck_capacity",
          label: "Truck Capacity (ton)",
          fieldType: "number",
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
        router.push("/admin/data/truck_type");
      }}
    />
  );
}
