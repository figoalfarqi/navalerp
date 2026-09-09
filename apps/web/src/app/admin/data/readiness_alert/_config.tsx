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
    col: "left",
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
    col: "right",
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
    col: "left",
    label: "Severity",
    fieldType: "select",
    options: [
      { label: "Critical", value: "CRITICAL" },
      { label: "High", value: "HIGH" },
      { label: "Medium", value: "MEDIUM" },
      { label: "Low", value: "LOW" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "alert_type",
    col: "right",
    label: "Alert Type",
    fieldType: "select",
    options: [
      { label: "CASREP Defect", value: "CASREP_DEFECT" },
      { label: "Critical Spare Deficit", value: "CRITICAL_SPARE_DEFICIT" },
      { label: "Crew Shortage", value: "CREW_SHORTAGE" },
      { label: "Expired Kelaikan", value: "EXPIRED_KELAIKAN" },
    ],
    required: true,
    disabled: mode === "view",
  },
  {
    name: "alert_message",
    col: "left",
    label: "Alert Message",
    fieldType: "textarea",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "is_acknowledged",
    col: "right",
    label: "Is Acknowledged",
    fieldType: "boolean",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "acknowledged_by_user_id",
    col: "left",
    label: "Acknowledged By User Id",
    fieldType: "text",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "acknowledged_at",
    col: "right",
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
