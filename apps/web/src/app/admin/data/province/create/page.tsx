"use client";
import FormCrud from "@/components/formCrud/FormCrud";
import { useQueryParams } from "@/hooks/useQueryParams";
import { useRouter } from "next/navigation";

export default function ProvinceFormPage() {
  const router = useRouter();

  const mode = "add";
  const url = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/province`;
  const { searchParams } = useQueryParams();
  return (
    <FormCrud
      title={"Create Province"}
      url={url}
      mode={mode}
      initialData={{}}
      fields={[
        {
          name: "province_name",
          label: "Province Name",
          fieldType: "text",
          col: "left",
          required: true,
        },
        {
          name: "province_real_name",
          label: "Province Real Name",
          fieldType: "text",
          col: "left",
          required: true,
        },
      ]}
      onSuccess={() => {
        router.push(`/admin/data/province?${searchParams}`);
      }}
    />
  );
}
