/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "vendor";
export const entityTitle = "Rekanan Industri";
export const entityEndpoint = "/admin/vendor";
export const primaryKey = "vendor_id";

export const columns: ColumnField[] = [
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
    col: "left",
    label: "Vendor Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "vendor_name",
    col: "right",
    label: "Vendor Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "tax_number",
    col: "left",
    label: "Tax Number",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "security_clearance_level",
    col: "right",
    label: "Security Clearance Level",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "defence_industry_license_no",
    col: "left",
    label: "Defence Industry License No",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "country",
    col: "right",
    label: "Country",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "contact_person",
    col: "left",
    label: "Contact Person",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "email",
    col: "right",
    label: "Email",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "phone",
    col: "left",
    label: "Phone",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "bank_account_name",
    col: "right",
    label: "Bank Account Name",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "bank_account_no",
    col: "left",
    label: "Bank Account No",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "bank_name",
    col: "right",
    label: "Bank Name",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "performance_rating",
    col: "left",
    label: "Performance Rating",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "is_approved",
    col: "right",
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
