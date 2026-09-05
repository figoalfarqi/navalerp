import { DynamicOptions, FormDataObject, StaticOptions } from "@/components/formCrud/FormCrud";

export function IsPrimitive(value: unknown): value is string | number {
  return typeof value === "string" || typeof value === "number";
}

export function IsDynamicOptions(
  options: StaticOptions | DynamicOptions | undefined
): options is DynamicOptions {
  return !!options && !Array.isArray(options) && "url" in options;
};

export function IsStaticOptions(
  options: StaticOptions | DynamicOptions | undefined
): options is StaticOptions {
  return Array.isArray(options);
}


export function toTitleCase(str: string): string {
  return str
    .toLowerCase()
    .split(" ")
    .map(word =>
      word.charAt(0).toUpperCase() + word.slice(1)
    )
    .join(" ");
}





export function ResolveUrl(
  url: string,
  data: Record<string, any>
): string {
  return url.replace(/\{(\w+)\}/g, (_, key) =>
    data[key] != null ? encodeURIComponent(String(data[key])) : "0"
  );
};

function getValueByPath(data: any, path: string): string {
  try {
    const parts = path.split(".");
    let current: any = data;

    for (const part of parts) {
      const match = part.match(/^(\w+)\[(\-?\d+)\]$/);

      if (match) {
        const [, key, indexStr] = match;
        const arr = current?.[key];
        if (!Array.isArray(arr)) return "";

        let index = parseInt(indexStr, 10);
        if (index === -1) index = arr.length - 1;

        current = arr[index];
      } else {
        current = current?.[part];
      }

      if (current == null) return "";
    }

    return String(current);
  } catch {
    return "";
  }
}

export function resolveLabelValueKey(
  template: string,
  data: any,
  isFilter: boolean = false
): string {
  if (!template) return "";
  // console.log("resolveLabelValueKey:", template, data);
  // 🔹 JIKA BUKAN TEMPLATE (tidak ada { })
  if (!template.includes("{") || !template.includes("}")) {
    return isFilter ? template : getValueByPath(data, template);
  }

  // 🔹 JIKA TEMPLATE DENGAN { }
  return template.replace(/\{([^}]+)\}/g, (_, path) => {
    return getValueByPath(data, path);
  });
}


export function IsDependencySatisfied(
  dependsOn: string[] | undefined,
  parentFormData?: FormDataObject,
  formData?: FormDataObject
): boolean {
  if (!dependsOn?.length) return true;

  const source = parentFormData
    ? { ...parentFormData, ...formData }
    : formData ?? {};

  return dependsOn.every(
    (key) =>
      source[key] !== undefined &&
      source[key] !== null &&
      source[key] !== ""
  );
};

export function GetDependsOnLabel(dependsOn?: string[]) {
  if (!dependsOn?.length) return "";
  return dependsOn.map((key) => key.replaceAll("_id", "")).join(", ");
};


export function spaceToUnderscore(text: string): string {
  return text.trim().replace(/\s+/g, "_");
}

export function resolveFileUrl(path: string) {
  const baseUrl = process.env.NEXT_PUBLIC_API_FILESERVICE_URL;

  if (!path) return "";

  // kalau sudah full url → ganti origin-nya
  if (path.startsWith("http://") || path.startsWith("https://")) {
    const url = new URL(path);
    return `${baseUrl}${url.pathname}`;
  }

  // kalau cuma path
  return `${baseUrl}${path}`;
}


export const resolveDefaultFilterValues = (
  defaults: Record<string, any>,
  row: any
) =>
  Object.keys(defaults).reduce((acc, key) => {
    let value =
      typeof defaults[key] === "string"
        ? resolveLabelValueKey(defaults[key], row, true)
        : defaults[key];
    // convert ke number jika memungkinkan
    if (
      typeof value === "string" &&
      value.trim() !== "" &&
      !isNaN(Number(value))
    ) {

      value = Number(value);
    }

    acc[key] = value;

    return acc;
  }, {} as Record<string, any>);

export const setNestedValue = (obj: any, path: string, value: any) => {
  const keys = path.split(".");
  let current = obj;

  keys.forEach((key, idx) => {
    if (idx === keys.length - 1) {
      current[key] = value;
    } else {
      current[key] = current[key] ?? {};
      current = current[key];
    }
  });
};


// export const getNestedValue = (obj: any, path: string) => {
//   if (!obj || !path) return undefined;
//   // return path.split(".").reduce((acc, key) => acc?.[key], obj);
//   let value = fungsi nya disini
//   if (
//     typeof value === "string" &&
//     value.trim() !== "" &&
//     !isNaN(Number(value))
//   ) {
//     value = Number(value);
//   }
//   return value
// };


export const getNestedValue = (obj: any, path: string) => {
  if (!obj || !path) return undefined;

  // Mendukung format "truck.driver.app_user_id" dan
  // format template lama "{truck.driver.app_user_id}".
  const match = path.match(/^\{(.+)\}$/);
  const realPath = match?.[1] ?? path;

  const tokens = realPath
    .replace(/\[(\-?\d+)\]/g, ".$1")
    .split(".")
    .filter(Boolean);

  let value: any = obj;

  for (const token of tokens) {
    if (value == null) return undefined;

    // Array index
    if (Array.isArray(value)) {
      const index = Number(token);
      if (isNaN(index)) return undefined;

      value =
        index === -1 ? value[value.length - 1] : value[index];
    } else {
      value = value[token];
    }
  }

  // Auto convert string number → number
  if (
    typeof value === "string" &&
    value.trim() !== "" &&
    !isNaN(Number(value))
  ) {
    return Number(value);
  }

  return value;
};
