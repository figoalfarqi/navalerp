/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "readiness_alert";
export const entityTitle = "Peringatan Dini";
export const entityEndpoint = "/admin/readiness_alert";
export const primaryKey = "alert_id";

export const columns: ColumnField[] = [
  { key: "ship_name", label: "Kapal KRI" },
  { key: "equipment_name", label: "Peralatan" },
  { key: "severity", label: "Severity" },
  { key: "alert_type", label: "Alert Type" },
  { key: "alert_message", label: "Alert Message" },
  { key: "is_acknowledged", label: "Is Acknowledged" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Peringatan Dini", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
];

export const formFields = (mode: string): FormField[] => [
    {
    name: "ship_id",
    label: "Kapal KRI",
    fieldType: "select",
    options: {
      url: "/admin/ship?limit=100",
      labelKey: "ship_name",
      valueKey: "ship_id",
    },
    required: true,
    disabled: mode === "view",
  },
    {
    name: "equipment_id",
    label: "Peralatan",
    fieldType: "select",
    options: {
      url: "/admin/equipment?limit=100",
      labelKey: "equipment_name",
      valueKey: "equipment_id",
    },
    required: false,
    disabled: mode === "view",
  },
  {
    name: "severity",
    label: "Severity",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "alert_type",
    label: "Alert Type",
    fieldType: "text",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "alert_message",
    label: "Alert Message",
    fieldType: "textarea",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "is_acknowledged",
    label: "Is Acknowledged",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "acknowledged_by_user_id",
    label: "Acknowledged By User Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "acknowledged_at",
    label: "Acknowledged At",
    fieldType: "date",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.alert_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
