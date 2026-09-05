/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import FormCrud from "@/components/formCrud/FormCrud";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { useQueryParams } from "@/hooks/useQueryParams";

export default function BankMerkFormPage() {
  const { id, mode } = useParams() as {
    id: string;
    mode: "edit" | "copy" | "view";
  };
  const [initialData, setInitialData] = useState<any>(null);
  const { searchParams } = useQueryParams();

  const router = useRouter();
  const { getAPI } = useFetchAPI();

  useEffect(() => {
    const fetchBankMerk = async () => {
      try {
        const res = await getAPI<any>(
          `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/bank_merk/${id}`,
          { authToken: "admin" }
        );
        const bank_merk = res.data.data || res.data;

        if (mode === "copy") {
          delete bank_merk.bank_merk_id;
        }
        setInitialData(bank_merk);
      } catch (err) {
        console.error("Failed to fetch Bank Merk", err);
      }
    };
    fetchBankMerk();
  }, [getAPI, id, mode]);

  if (!initialData) return <p>Loading...</p>;

  const url =
    mode === "edit"
      ? `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/bank_merk/${id}`
      : `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/bank_merk`;

  return (
    <FormCrud
      title={
        mode === "edit"
          ? "Edit Bank Merk"
          : mode === "copy"
          ? "Copy Bank Merk"
          : "View Bank Merk"
      }
      url={url}
      mode={mode}
      initialData={initialData}
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
        router.push(`/admin/data/bank_merk?${searchParams}`);
      }}
    />
  );
}
