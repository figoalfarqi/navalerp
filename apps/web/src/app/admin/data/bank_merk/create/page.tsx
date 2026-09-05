"use client";
import FormCrud from "@/components/formCrud/FormCrud";
import { useQueryParams } from "@/hooks/useQueryParams";
import { useRouter } from "next/navigation";

export default function BankMerkFormPage() {
  const router = useRouter();

  const mode = "add";
  const url = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/bank_merk`;
  const { searchParams } = useQueryParams();
  return (
    <FormCrud
      title={"Create Bank Merk"}
      url={url}
      mode={mode}
      initialData={{}}
      fields={[
        {
          name: "bank_merk_name",
          label: "Bank Merk",
          fieldType: "text",
          col: "left",
          required: true,
        },
        {
          name: "bank_merk_description",
          label: "Description",
          fieldType: "text",
          col: "left",
          required: false,
        },
      ]}
      onSuccess={() => {
        router.push(`/admin/data/bank_merk?${searchParams}`);
      }}
    />
  );
}
