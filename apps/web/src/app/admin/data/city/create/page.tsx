"use client";
import FormCrud from "@/components/formCrud/FormCrud";
import { useRouter } from "next/navigation";

export default function CityFormPage() {
  const router = useRouter();

  const mode = "add";
  const url = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city`;

  return (
    <FormCrud
      title={"Create City"}
      url={url}
      mode={mode}
      initialData={{}}
      fields={[
        {
          name: "province_id",
          label: "Province",
          fieldType: "select",
          col: "left",
          required: true,
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
        },
      ]}
      onSuccess={() => {
        router.push("/admin/data/city");
      }}
    />
  );
}
