/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "theater";
export const entityTitle = "Teater Operasi";
export const entityEndpoint = "/admin/theater";
export const primaryKey = "theater_id";

export const columns: ColumnField[] = [
  { key: "theater_code", label: "Theater Code" },
  { key: "theater_name", label: "Theater Name" },
  { key: "command_unit_name", label: "Komando Pengendali" },
  { key: "threat_level", label: "Threat Level" },
  { key: "description", label: "Description" },
  { key: "created_at", label: "Created At" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Teater Operasi", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "theater_code",
    label: "Theater Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "theater_name",
    label: "Theater Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "responsible_command_unit_id",
    label: "Komando Pengendali",
    fieldType: "select",
    options: {
      url: "/admin/org_unit?limit=100",
      labelKey: "unit_name",
      valueKey: "unit_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "threat_level",
    label: "Threat Level",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "description",
    label: "Description",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.theater_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
