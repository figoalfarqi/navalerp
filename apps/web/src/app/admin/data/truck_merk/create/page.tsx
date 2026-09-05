"use client";
import FormCrud from "@/components/formCrud/FormCrud";
import { useQueryParams } from "@/hooks/useQueryParams";
import { useRouter } from "next/navigation";

export default function TruckMerkFormPage() {
  const router = useRouter();

  const mode = "add";
  const url = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck_merk`;
  const { searchParams } = useQueryParams();
  return (
    <FormCrud
      title={"Create Truck Merk"}
      url={url}
      mode={mode}
      initialData={{}}
      fields={[
        {
          name: "truck_merk_name",
          label: "Truck Merk Name",
          fieldType: "text",
          col: "left",
          required: true,
        },
      ]}
      onSuccess={() => {
        router.push(`/admin/data/truck_merk?${searchParams}`);
      }}
    />
  );
}
