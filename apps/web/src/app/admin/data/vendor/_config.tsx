/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "vendor";
export const entityTitle = "Rekanan Industri";
export const entityEndpoint = "/admin/vendor";
export const primaryKey = "vendor_id";

export const columns: ColumnField[] = [
  { key: "vendor_id", label: "Vendor Id" },
  { key: "vendor_code", label: "Vendor Code" },
  { key: "vendor_name", label: "Vendor Name" },
  { key: "tax_number", label: "Tax Number" },
  { key: "security_clearance_level", label: "Security Clearance Level" },
  { key: "defence_industry_license_no", label: "Defence Industry License No" },
  { key: "country", label: "Country" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Rekanan Industri", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "vendor_code",
    label: "Vendor Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "vendor_name",
    label: "Vendor Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "tax_number",
    label: "Tax Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "security_clearance_level",
    label: "Security Clearance Level",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "defence_industry_license_no",
    label: "Defence Industry License No",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "country",
    label: "Country",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "contact_person",
    label: "Contact Person",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "email",
    label: "Email",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "phone",
    label: "Phone",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "bank_account_name",
    label: "Bank Account Name",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "bank_account_no",
    label: "Bank Account No",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "bank_name",
    label: "Bank Name",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "performance_rating",
    label: "Performance Rating",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_approved",
    label: "Is Approved",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.vendor_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
