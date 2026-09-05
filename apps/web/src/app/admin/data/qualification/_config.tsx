/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "qualification";
export const entityTitle = "Kualifikasi & Brevet";
export const entityEndpoint = "/admin/qualification";
export const primaryKey = "qualification_id";

export const columns: ColumnField[] = [
  { key: "qualification_code", label: "Qualification Code" },
  { key: "qualification_name", label: "Qualification Name" },
  { key: "qualification_category", label: "Qualification Category" },
  { key: "issuing_institution", label: "Issuing Institution" },
  { key: "validity_years", label: "Validity Years" },
  { key: "description", label: "Description" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Kualifikasi & Brevet", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "qualification_code",
    col: "left",
    label: "Qualification Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "qualification_name",
    col: "right",
    label: "Qualification Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "qualification_category",
    col: "left",
    label: "Qualification Category",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "issuing_institution",
    col: "right",
    label: "Issuing Institution",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "validity_years",
    col: "left",
    label: "Validity Years",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "description",
    col: "right",
    label: "Description",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.qualification_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
