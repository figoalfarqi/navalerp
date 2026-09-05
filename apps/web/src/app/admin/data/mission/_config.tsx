/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "mission";
export const entityTitle = "Misi Tempur";
export const entityEndpoint = "/admin/mission";
export const primaryKey = "mission_id";

export const columns: ColumnField[] = [
  { key: "theater_name", label: "Teater Operasi" },
  { key: "mission_code", label: "Mission Code" },
  { key: "mission_name", label: "Mission Name" },
  { key: "mission_type", label: "Mission Type" },
  { key: "start_date", label: "Start Date" },
  { key: "end_date", label: "End Date" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Misi Tempur", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "theater_id",
    col: "left",
    label: "Teater Operasi",
    fieldType: "select",
    options: {
      url: "/admin/theater?limit=100",
      labelKey: "theater_name",
      valueKey: "theater_id",
    },
    required: true,
    disabled: mode === "view",
  },
  {
    name: "mission_code",
    col: "right",
    label: "Mission Code",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "mission_name",
    col: "left",
    label: "Mission Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "mission_type",
    col: "right",
    label: "Mission Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "start_date",
    col: "left",
    label: "Start Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "end_date",
    col: "right",
    label: "End Date",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
    {
    name: "commanding_officer_user_id",
    col: "left",
    label: "Komandan Operasi",
    fieldType: "select",
    options: {
      url: "/admin/sys_user?limit=100",
      labelKey: "full_name",
      valueKey: "user_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "mission_status",
    col: "right",
    label: "Mission Status",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.mission_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
