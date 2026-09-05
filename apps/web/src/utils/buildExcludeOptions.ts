import { FormDataObject } from "@/components/formCrud/FormCrud";
import { ExcludedFieldNames } from "@/consta/DetailFormData";
import { IsPrimitive } from "@/utils/globalUtils";

export type ExcludeOptions = Record<string, Array<string | number>>;

export const buildExcludeOptions = (
  list: FormDataObject[]
): ExcludeOptions => {
  const result: ExcludeOptions = Object.fromEntries(
    ExcludedFieldNames.map((f) => [f, []])
  );

  list.forEach((item) => {
    ExcludedFieldNames.forEach((field) => {
      const value = item[field];
      if (IsPrimitive(value) && !result[field].includes(value)) {
        result[field].push(value);
      }
    });
  });

  return result;
};
