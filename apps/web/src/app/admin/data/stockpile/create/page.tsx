"use client";
import FormCrud, { FormField } from "@/components/formCrud/FormCrud";
import { useRouter } from "next/navigation";

export default function StockpileFormPage() {
  const router = useRouter();

  const mode = "add";
  const url = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/stockpile`;

  const fields: FormField[] = [
    {
      name: "stockpile_name",
      label: "Stockpile Name",
      fieldType: "text",
      col: "left",
      required: true,
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
      title={"Create Stockpile"}
      url={url}
      mode={mode}
      fields={fields}
      onSuccess={() => {
        router.push("/admin/data/stockpile");
      }}
      initialData={{is_active:1}}
    />
  );
}
