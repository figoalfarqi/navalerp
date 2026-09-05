"use client";
import FormCrud from "@/components/formCrud/FormCrud";
import { useRouter } from "next/navigation";

export default function TruckTypeFormPage() {
  const router = useRouter();

  const mode = "add";
  const url = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck_type`;

  return (
    <FormCrud
      title={"Create Truck Type"}
      url={url}
      mode={mode}
      initialData={{}}
      fields={[
        {
          name: "truck_type_name",
          label: "Truck Type Name",
          fieldType: "text",
          col: "left",
          required: true,
        },
        {
          name: "truck_type_description",
          label: "Truck Type Description",
          fieldType: "textarea",
          col: "left",
        },
        {
          name: "truck_box_length",
          label: "Truck Box Length (m)",
          fieldType: "number",
          col: "right",
          required: true,
        },
        {
          name: "truck_box_width",
          label: "Truck Box Width (m)",
          fieldType: "number",
          col: "left",
          required: true,
        },
        {
          name: "truck_box_height",
          label: "Truck Box Height (m)",
          fieldType: "number",
          col: "right",
          required: true,
        },
        {
          name: "truck_capacity",
          label: "Truck Capacity (ton)",
          fieldType: "number",
          col: "left",
          required: true,
        },
      ]}
      onSuccess={() => {
        router.push("/admin/data/truck_type");
      }}
    />
  );
}
