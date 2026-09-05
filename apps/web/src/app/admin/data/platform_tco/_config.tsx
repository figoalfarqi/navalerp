/* eslint-disable @typescript-eslint/no-explicit-any */
import { ColumnField } from "@/components/table/Table";
import { FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";

export const entityName = "platform_tco";
export const entityTitle = "Total Cost of Ownership";
export const entityEndpoint = "/admin/platform_tco";
export const primaryKey = "tco_id";

export const columns: ColumnField[] = [
  { key: "ship_name", label: "Kapal KRI" },
  { key: "fiscal_year", label: "Fiscal Year" },
  { key: "acquisition_amortization", label: "Acquisition Amortization" },
  { key: "fuel_lube_cost", label: "Fuel Lube Cost" },
  { key: "mro_spareparts_cost", label: "Mro Spareparts Cost" },
  { key: "docking_services_cost", label: "Docking Services Cost" },
];

export const filterFields: FilterField[] = [
  { name: "search", label: "Cari Total Cost of Ownership", fieldType: "text", col: "left", placeHolder: "Ketik kata kunci pencarian..." },
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
    name: "fiscal_year",
    label: "Fiscal Year",
    fieldType: "number",
    required: true,
    disabled: mode === "view",
  },
  {
    name: "acquisition_amortization",
    label: "Acquisition Amortization",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "fuel_lube_cost",
    label: "Fuel Lube Cost",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "mro_spareparts_cost",
    label: "Mro Spareparts Cost",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "docking_services_cost",
    label: "Docking Services Cost",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "crew_payroll_allowances",
    label: "Crew Payroll Allowances",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "modernization_upgrades_cost",
    label: "Modernization Upgrades Cost",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "operating_hours_sea",
    label: "Operating Hours Sea",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "cost_per_operating_hour",
    label: "Cost Per Operating Hour",
    fieldType: "number",
    required: false,
    disabled: mode === "view",
  },
  {
    name: "remarks",
    label: "Remarks",
    fieldType: "textarea",
    required: false,
    disabled: mode === "view",
  },
];

export const buildPayload = (data: any) => {
  const payload: any = { ...data };
  delete payload.tco_id;
  delete payload.created_at;
  delete payload.updated_at;
  delete payload.deleted_at;
  delete payload.created_by;
  delete payload.updated_by;
  delete payload.deleted_by;
  return payload;
};
