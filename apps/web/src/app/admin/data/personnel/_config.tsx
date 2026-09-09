/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "personnel";
export const entityTitle = "Prajurit TNI AL";
export const entityEndpoint = "/admin/personnel";
export const primaryKey = "personnel_id";

export const columns: ColumnField[] = [
  { key: "nrp", label: "Nrp" },
  { key: "full_name", label: "Full Name" },
  { key: "rank_name", label: "Pangkat" },
  { key: "corps_name", label: "Korps" },
  { key: "unit_name", label: "Satuan Sekarang" },
  { key: "current_position", label: "Current Position" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Prajurit TNI AL", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
  {
    name: "nrp",
    col: "left",
    label: "Nrp",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "full_name",
    col: "right",
    label: "Full Name",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "rank_id",
    col: "left",
    label: "Pangkat",
    fieldType: "select",
    options: {
      url: "/admin/military_rank?limit=100",
      labelKey: "rank_name",
      valueKey: "rank_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "corps_id",
    col: "right",
    label: "Korps",
    fieldType: "select",
    options: {
      url: "/admin/military_corps?limit=100",
      labelKey: "corps_name",
      valueKey: "corps_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "current_unit_id",
    col: "left",
    label: "Satuan Sekarang",
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
    name: "current_position",
    col: "right",
    label: "Current Position",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "birth_place",
    col: "left",
    label: "Birth Place",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "birth_date",
    col: "right",
    label: "Birth Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "gender",
    col: "left",
    label: "Gender",
    fieldType: "select",
    options: [
      { label: "Male", value: "MALE" },
      { label: "Female", value: "FEMALE" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "blood_type",
    col: "right",
    label: "Blood Type",
    fieldType: "select",
    options: [
      { label: "A", value: "A" },
      { label: "B", value: "B" },
      { label: "AB", value: "AB" },
      { label: "O", value: "O" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "religion",
    col: "left",
    label: "Religion",
    fieldType: "select",
    options: [
      { label: "Islam", value: "ISLAM" },
      { label: "Protestan", value: "PROTESTAN" },
      { label: "Katolik", value: "KATOLIK" },
      { label: "Hindu", value: "HINDU" },
      { label: "Buddha", value: "BUDDHA" },
      { label: "Konghucu", value: "KONGHUCU" },
    ],
    required: false,
    disabled: mode === "view",
  },
  {
    name: "education_level",
    col: "right",
    label: "Education Level",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "service_entry_date",
    col: "left",
    label: "Service Entry Date",
    fieldType: "date",
    required: true,
    disabled: mode === "view",
  },
    {
    name: "user_id",
    col: "right",
    label: "Pengguna",
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
    name: "status",
    col: "left",
    label: "Status",
    fieldType: "select",
    options: [
      { label: "Active", value: "ACTIVE" },
      { label: "On Leave", value: "ON_LEAVE" },
      { label: "Retired", value: "RETIRED" },
      { label: "Deceased", value: "DECEASED" },
      { label: "WIA", value: "WIA" },
      { label: "KIA", value: "KIA" },
    ],
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.personnel_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
