import type { FormDataObject, FormDataValue } from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";

export type AdminFilterField = FilterField;
export type AdminCrudMode = "add" | "edit" | "copy" | "view";

export function pickAdminPayload(
  data: FormDataObject,
  fields: readonly string[],
): Record<string, FormDataValue> {
  return fields.reduce<Record<string, FormDataValue>>((payload, field) => {
    const value = data[field];
    payload[field] = value === "" || value === undefined ? null : value;
    return payload;
  }, {});
}

export function crudTitle(
  mode: AdminCrudMode,
  entityLabel: string,
): string {
  const prefix: Record<AdminCrudMode, string> = {
    add: "Tambah",
    edit: "Edit",
    copy: "Salin",
    view: "Detail",
  };
  return `${prefix[mode]} ${entityLabel}`;
}
