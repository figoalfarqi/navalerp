/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { ComponentType, useEffect, useMemo, useRef, useState } from "react";
import Button from "../form/Button";
import { usePathname, useRouter } from "next/navigation";
import LoaderDots from "../form/LoaderDots";
import { useToast } from "../ToastContext";
import { authTokenType, useFetchAPI } from "@/hooks/useFetchAPI";
import { ChainItem } from "../form/ChainSelectField";
import { useIsMobile } from "@/hooks/useIsMobile";
import { useQueryParams } from "@/hooks/useQueryParams";
import { useUploadFile } from "@/hooks/useUploadFile";
import { TableProps } from "../table/Table";
import { Mode } from "@/hooks/useDetails";
import FieldLayout from "./FieldLayout";
import TablesLayout from "./TablesLayout";
import DetailLayout from "./DetailLayout";
import { DividerFieldNames } from "@/consta/DetailFormData";
import { IsStaticOptions } from "@/utils/globalUtils";
import { SelectOption } from "../form/SelectField";
import { RadioOption } from "../form/RadioField";
import { useCheckUniqueField } from "@/hooks/useCheckUniqueField";
import DetailItemsTable, { DetailItemsConfig } from "./DetailItemsTable";
import { AUTO_NUMBER_CONFIGS, fetchAutoNumber } from "@/utils/numberGenerator";

export type StaticOptions = SelectOption[] | RadioOption[];
export type DynamicOptions = {
  url: string;
  labelKey: string;
  valueKey: string;
  autoSelectFirst?: boolean;
  dependsOn?: string[];
  excludedValues?: Array<string | number>;
};

export type rowGetter = {
  key: string;
  source: string;
};
export interface FormField {
  name: string;
  label: string;
  validation?: {
    min?: number;
    max?: number;
    gt?: number;
    lt?: number;
  };
  fieldType:
    | "text"
    | "password"
    | "number"
    | "email"
    | "textarea"
    | "date"
    | "dateyear"
    | "datetime"
    | "duration"
    | "select"
    | "chainselect"
    | "boolean"
    | "uploadimage"
    | "barcodescanner";
  col?: "left" | "right";
  required?: boolean;
  disabled?: boolean;
  readOnly?: boolean;
  autoGenerate?: boolean;
  autoPrefix?: string;
  autoEntity?: string;
  placeholder?: string;
  chains?: ChainItem[];
  rowGetters?: rowGetter[];
  clearFieldsOnChange?: string[];
  tableProps?: TableProps;
  fileFolder?: string;
  options?: StaticOptions | DynamicOptions;
  defaultValueUrl?: string; // URL untuk ambil default value
  uniqueUrl?: string; // untuk unique field
}

// export interface ValidationOnlyField {
//   name: string;
//   type: "array";
//   fields: FormField[];
// }

export interface DetailActionFromCrud {
  setMode: (mode: Mode) => void;
  add: (data: any) => void;
  update: (data: any) => void;
  resetItem: (item: any) => void;
}

export interface tablesProps {
  title: string;
  formCrudProps: FormCrudProps;
  tableProps: TableProps;
}

export interface DetailFormProps {
  title: string;
  table_key: string;
  table_name: string;
  fields: FormField[];
  mode: Mode;
  maxItems: number;
  detailFormGetUrl?: string;
  detailFormDependsOn?: string[];
}
export type FormDataValue =
  | string
  | number
  | boolean
  | null
  | undefined
  | FormDataObject
  | FormDataValue[];

export interface FormDataObject {
  [key: string]: FormDataValue;
}

export interface FormCrudProps {
  title: string;
  url: string;
  mode: Mode;
  initialData?: any;
  fields: FormField[];
  hideSubmit?: boolean;
  onSuccess?: (res: any) => void;
  AddedComponent?: ComponentType;
  tablesDependsOn?: string[];
  resetKey?: string;
  authToken?: authTokenType;
  tablesAddedData?: any;
  tables?: tablesProps[];
  detailForms?: DetailFormProps[];
  detailItemsConfig?: DetailItemsConfig;
  formCrudFor?: "detail" | "normal";
  detailAction?: DetailActionFromCrud;
  buildPayload?: (formData: FormDataObject, tables: tablesProps[]) => void;
  parentFormData?: FormDataObject;
}

const EMPTY_INITIAL_DATA: FormDataObject = {};

export default function FormCrud({
  title,
  url,
  mode,
  initialData = EMPTY_INITIAL_DATA,
  fields,
  hideSubmit,
  onSuccess,
  AddedComponent,
  tablesDependsOn,
  resetKey,
  authToken = "admin",
  tablesAddedData,
  tables,
  detailForms,
  detailItemsConfig,
  formCrudFor = "normal",
  detailAction,
  buildPayload,
  parentFormData,
}: FormCrudProps) {
  const { postAPI, putAPI, getAPI } = useFetchAPI();
  const [formData, setFormData] = useState<FormDataObject>(initialData || {});
  const [errorForm, setErrorForm] = useState<{ [key: string]: string }>({});

  const [loadingSumbit, setLoadingSumbit] = useState(false);

  // Auto-generate document numbering for add and copy modes
  useEffect(() => {
    if (!["add", "copy"].includes(mode)) return;

    fields.forEach(async (field) => {
      const isAuto = field.autoGenerate || AUTO_NUMBER_CONFIGS[field.name];
      if (!isAuto) return;

      if (!formData[field.name]) {
        const nextNum = await fetchAutoNumber(
          field.name,
          field.autoPrefix,
          field.autoEntity,
          getAPI
        );
        setFormData((prev) => {
          if (prev[field.name]) return prev;
          return {
            ...prev,
            [field.name]: nextNum,
          };
        });
      }
    });
  }, [mode, fields, getAPI]);

  const isForDetail = formCrudFor === "detail";

  const router = useRouter();
  const pathname = usePathname(); // misal: /admin/data/driver/2/edit
  const { showToast } = useToast();

  const { isMobile } = useIsMobile();
  const { searchParams } = useQueryParams();
  const { uploadFileData } = useUploadFile();

  const { errorFormCheckUnique, isLoadingCheckUniqu, checkUniqueNow } =
    useCheckUniqueField({
      fields,
      formData,
    });

  const errorFormCombined = useMemo(
    () => ({
      ...errorForm,
      ...errorFormCheckUnique,
    }),
    [errorForm, errorFormCheckUnique],
  );

  // Ambil base route: /admin/data/driver
  const querySuffix =
    searchParams && searchParams.toString() ? `?${searchParams.toString()}` : "";
  const baseRoute =
    pathname.split("/").slice(0, 4).join("/") + querySuffix;

  const handleValidate = async (): Promise<boolean> => {
    const newErrors: Record<string, string> = {};

    const validateValue = (field: FormField, value: any, errorKey: string) => {
      const v = field.validation || {};

      if (
        field.required &&
        field.fieldType !== "chainselect" &&
        (value === "" || value === null || value === undefined)
      ) {
        newErrors[errorKey] = "This field is required";
        return;
      }

      if (value === "" || value === null || value === undefined) return;

      if (field.fieldType === "number") {
        const numVal = Number(value);
        if (isNaN(numVal)) {
          newErrors[errorKey] = "Must be a number";
          return;
        }
        if (v.min !== undefined && numVal < v.min) {
          newErrors[errorKey] = `Must be at least ${v.min}`;
        }
        if (v.max !== undefined && numVal > v.max) {
          newErrors[errorKey] = `Must be at most ${v.max}`;
        }
        if (v.gt !== undefined && numVal <= v.gt) {
          newErrors[errorKey] = `Must be greater than ${v.gt}`;
        }
        if (v.lt !== undefined && numVal >= v.lt) {
          newErrors[errorKey] = `Must be less than ${v.lt}`;
        }
      }

      if (field.fieldType !== "number" && typeof value === "string") {
        if (v.min !== undefined && value.length < v.min) {
          newErrors[errorKey] = `Minimum ${v.min} characters`;
        }
        if (v.max !== undefined && value.length > v.max) {
          newErrors[errorKey] = `Maximum ${v.max} characters`;
        }
      }
    };

    for (const field of fields) {
      const value = formData[field.name];

      // 🔹 CASE 1: FIELD BIASA
      if (!Array.isArray(value)) {
        validateValue(field, value, field.name);
        continue;
      }
    }
    if (detailForms && detailForms.length > 0) {
      for (const group of detailForms) {
        const list = formData[group.table_name];
        if (!list) newErrors[group.table_name] = `Minimal 1 data`;
        if (!Array.isArray(list)) continue;
        if (list.length < 1) newErrors[group.table_name] = `Minimal 1 data`;
        const dividerField = group.fields.find((field) =>
          DividerFieldNames.includes(field.name),
        );
        if (IsStaticOptions(dividerField?.options)) {
          dividerField.options.map((option) => {
            const dividedList = list.filter(
              (li) =>
                (li as FormDataObject)?.[dividerField.name] === option.value,
            );
            if (dividedList.length < 1) {
              newErrors[
                `${group.table_name}.${dividerField.name}:${option.value}`
              ] = `Minimal 1 data`;
            }
            dividedList.forEach((item, index) => {
              group.fields.forEach((field) => {
                validateValue(
                  field,
                  (item as FormDataObject)[field.name],
                  `${group.table_name}.${dividerField.name}:${option.value}.${index}.${field.name}`,
                );
              });
            });
          });
        } else {
          list.forEach((item, index) => {
            group.fields.forEach((field) => {
              validateValue(
                field,
                (item as FormDataObject)[field.name],
                `${group.table_name}.${index}.${field.name}`,
              );
            });
          });
        }
      }
    }
    setErrorForm(newErrors);
    
    const isUnique = await checkUniqueNow();
    return Object.keys(newErrors).length === 0 && isUnique;
  };

  const tablesDependsOnSignatureKey = useMemo(() => {
    if (!tablesDependsOn) return "";

    return tablesDependsOn.map((key) => String(formData[key] ?? "")).join("|");
  }, [tablesDependsOn, formData]);

  const didInitRef = useRef(false);
  const prevInitialDataRef = useRef<any>(initialData);

  useEffect(() => {
    if (formCrudFor === "detail") {
      if (!didInitRef.current) {
        didInitRef.current = true;
        setFormData(initialData);
        return;
      }

      if (JSON.stringify(formData) !== JSON.stringify(initialData)) {
        setFormData(initialData);
      }
      return;
    }

    // Normal forms: only update formData if initialData actually changed from outside (e.g. async fetch)
    const prev = prevInitialDataRef.current;
    if (JSON.stringify(prev) !== JSON.stringify(initialData)) {
      prevInitialDataRef.current = initialData;
      if (initialData && typeof initialData === "object" && Object.keys(initialData).length > 0) {
        setFormData(initialData);
      }
    }
  }, [initialData, resetKey, formCrudFor]);

  useEffect(() => {
    if (Object.keys(errorForm).length > 0) {
      handleValidate();
    }
  }, [formData]);

  // ==============================================
  //  START : ini hanya untuk create delivery order
  // ==============================================
  // const prevRef = useRef<{
  //   truck_type_id?: number;
  //   project_detail_id?: number;
  // }>({});

  // useEffect(() => {
  //   const prev = prevRef.current;
    
  //   const truckTypeChanged = prev.truck_type_id !== formData.truck_type_id;
    
  //   const projectDetailChanged =
  //   prev.project_detail_id !== formData.project_detail_id;
    
  //   if (!formData.truck_type_id || !formData.project_detail_id) {
  //     prevRef.current = {
  //       truck_type_id: formData.truck_type_id as number,
  //       project_detail_id: formData.project_detail_id as number,
  //     };
  //     return;
  //   }
  //   if (!Array.isArray(formData.travel_allowances)) {

  //     prevRef.current = {
  //       truck_type_id: formData.truck_type_id as number,
  //       project_detail_id: formData.project_detail_id as number,
  //     };
  //     return;
  //   }
    
  //   // 🔁 Jalankan logic hanya jika salah satu berubah
  //   if (truckTypeChanged || projectDetailChanged) {
  //     const taData: any = formData.travel_allowances.find(
  //       (ta: any) => ta.truck_type_id === formData.truck_type_id,
  //     );

  //     if (taData?.travel_allowance_amount_1) {
  //       setFormData((prev) => ({
  //         ...prev,
  //         travel_allowance_amount_1: taData.travel_allowance_amount_1,
  //         travel_allowance_amount_2: taData.travel_allowance_amount_2 ?? 0,
  //       }));
  //     } else {
  //       // 🔍 INFO: field mana yang berubah
  //       if (truckTypeChanged) {
  //         setFormData((prev) => ({
  //           ...prev,
  //           truck_type_id: null,
  //         }));
  //       }

  //       if (projectDetailChanged) {
  //         setFormData((prev) => ({
  //           ...prev,
  //           project_detail_id: null,
  //         }));
  //       }
  //       showToast(
  //         3000,
  //         "error",
  //         "Travel Allowance tidak ada untuk truck type rute ini",
  //       );
  //     }
  //   }

  //   // 🧠 Simpan nilai sekarang untuk render berikutnya
  //   prevRef.current = {
  //     truck_type_id: formData.truck_type_id as number,
  //     project_detail_id: formData.project_detail_id as number,
  //   };
  // }, [formData?.truck_type_id, formData?.project_detail_id]);

  const handleDetailItemsChange = (newItems: any[]) => {
    if (!detailItemsConfig) return;
    setFormData((prev) => ({
      ...prev,
      [detailItemsConfig.tableName]: newItems,
    }));
  };

  const handleSyncHeaderTotal = (total: number) => {
    if (!detailItemsConfig?.syncHeaderTotalKey) return;
    const targetKey = detailItemsConfig.syncHeaderTotalKey;
    setFormData((prev) => {
      if (Number(prev[targetKey]) === total) return prev;
      return {
        ...prev,
        [targetKey]: total,
      };
    });
  };

  // Handle form submit
  const handleSubmit = async (e: React.FormEvent) => {
    setLoadingSumbit(true);
    e.preventDefault();

    const isValid = await handleValidate();
    if (!isValid) {
      showToast(3000, "error", "Please fix validation errors!");
      setLoadingSumbit(false);
      return;
    }

    try {
      let res;
      const payload = buildPayload
        ? buildPayload(formData, tables ?? [])
        : formData;

      if (["add", "copy"].includes(mode)) {
        res = await postAPI(url, { authToken: authToken}, payload);
      } else {
        res = await putAPI(url, { authToken: authToken }, payload);
      }
      if (res && res.code) {
        if ([200, 201].includes(res.code)) {
          setErrorForm({});
          if (onSuccess) onSuccess(res.data);
          showToast(
            3000,
            "success",
            `${
              mode === "add"
                ? "Data berhasil disimpan!"
                : mode === "copy"
                  ? "Data berhasil disalin!"
                  : "Data berhasil diupdate!"
            }`,
          );
        } else {
          setErrorForm(res.errors || {});
          showToast(3000, "error", `Gagal menyimpan data! ${res.message}`);
          setLoadingSumbit(false);
        }
      } else {
        showToast(3000, "error", "Gagal menyimpan data!");
        setLoadingSumbit(false);
      }
    } catch (err) {
      console.error("Submit failed", err);
      showToast(3000, "error", "Gagal menyimpan data!");
      setLoadingSumbit(false);
    }
  };

  const Wrapper = isForDetail ? "div" : "form";
  return (
    <div className={`${isForDetail ? "" : "pt-3 px-3"} w-full`}>
      <div
        className={`${
          isForDetail
            ? ""
            : "rounded-md shadow p-4 border-t-4 border-[#0a2540] overflow-visible md:overflow-auto h-auto md:h-[calc(100vh-80px)] md:app-scrollbar"
        } bg-white`}
      >
        <h2 className={`${isForDetail ? "text-lg" : "text-xl"} font-bold mb-2`}>
          {title}
        </h2>
        <Wrapper
          {...(isForDetail ? {} : { onSubmit: handleSubmit })}
          className="grid grid-cols-1 md:grid-cols-2 gap-4"
        >
          {formData && (
            <FieldLayout
              fields={fields}
              formData={formData}
              setFormData={setFormData}
              tablesDependsOn={tablesDependsOn}
              tablesAddedData={tablesAddedData}
              authToken={authToken}
              errorForm={errorFormCombined}
            />
          )}
          {tables && tables.length > 0 && (
            <>
              {(tablesDependsOn
                ? // 🔹 Jika ADA tablesDependsOn → pakai aturan lama
                  tablesDependsOn.every(
                    (key) =>
                      formData[key] !== undefined && formData[key] !== null,
                  )
                : // 🔹 Jika TIDAK ADA tablesDependsOn → tunggu formData created
                  // formData && Object.keys(formData).length > 0
                  true) && (
                <TablesLayout
                  tables={tables}
                  parentFormData={formData}
                  {...(tablesDependsOn
                    ? { resetKey: tablesDependsOnSignatureKey }
                    : {})}
                />
              )}
            </>
          )}

          {formData && detailForms && detailForms?.length > 0 && (
            <DetailLayout
              detailForms={detailForms}
              parentFormData={parentFormData}
              formData={formData}
              setFormData={setFormData}
              errorForm={errorFormCombined}
            />
          )}

          {detailItemsConfig && (
            <DetailItemsTable
              config={detailItemsConfig}
              items={
                Array.isArray(formData[detailItemsConfig.tableName])
                  ? (formData[detailItemsConfig.tableName] as any[])
                  : []
              }
              onChange={handleDetailItemsChange}
              mode={mode}
              onSyncHeaderTotal={handleSyncHeaderTotal}
            />
          )}
          <div className="col-span-1 md:col-span-2 flex justify-end mt-2">
            {!isForDetail && (
              <Button
                type="reset"
                size="md"
                variant="yellow-solid"
                className="mr-2"
                onClick={() => router.push(baseRoute)}
                id={"button-reset"}
              >
                Back
              </Button>
            )}
            {!hideSubmit && (
              <>
                <Button
                  type="reset"
                  size="md"
                  variant="gray-solid"
                  className="mr-2"
                  onClick={() => {
                    if (isForDetail) {
                      detailAction?.resetItem(formData);
                    } else {
                      const resetData: any = { ...initialData };
                      fields.forEach((f) => {
                        if (
                          (f.autoGenerate || AUTO_NUMBER_CONFIGS[f.name]) &&
                          formData[f.name]
                        ) {
                          resetData[f.name] = formData[f.name];
                        }
                      });
                      setFormData(resetData);
                    }
                  }}
                  id={"button-reset"}
                >
                  Reset
                </Button>
                <Button
                  type={isForDetail ? "button" : "submit"}
                  onClick={
                    isForDetail
                      ? async () => {
                          const isValid = await handleValidate();
                          if (isValid) {
                            if (["add", "copy"].includes(mode)) {
                              detailAction?.add({
                                ...formData,
                                created_at:
                                  formData.created_at ??
                                  new Date().toISOString(),
                                updated_at:
                                  formData.updated_at ??
                                  new Date().toISOString(),
                              });
                              setFormData(initialData);
                            } else {
                              if (
                                JSON.stringify(formData) !==
                                JSON.stringify(initialData)
                              )
                                detailAction?.update(formData);
                            }
                          } else {
                            showToast(
                              3000,
                              "error",
                              "Please fix validation errors!",
                            );
                          }
                        }
                      : undefined
                  }
                  size="md"
                  variant="blue-solid"
                  className=""
                  id={"button-save"}
                  disabled={
                    loadingSumbit ||
                    (JSON.stringify(formData) === JSON.stringify(initialData) &&
                      isForDetail)
                  }
                >
                  {loadingSumbit || uploadFileData.isUploading ? (
                    <LoaderDots />
                  ) : mode === "add" ? (
                    "Simpan"
                  ) : mode === "copy" ? (
                    "Salin"
                  ) : (
                    "Update"
                  )}
                </Button>
              </>
            )}
          </div>
        </Wrapper>
      </div>
      {AddedComponent && <AddedComponent />}
    </div>
  );
}
