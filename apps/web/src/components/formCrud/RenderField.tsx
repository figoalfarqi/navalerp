/* eslint-disable @typescript-eslint/no-explicit-any */
import { Dispatch, SetStateAction, useEffect } from "react";
import ChainSelectField from "../form/ChainSelectField";
import DateField from "../form/DateField";
import DateTimeField from "../form/DateTimeField";
import UploadFileField from "../form/FileField";
import SelectField, { SelectOption } from "../form/SelectField";
import SwitchField from "../form/SwitchField";
import TextAreaField from "../form/TextAreaField";
import TextField from "../form/TextField";
import TextFieldSkeleton from "../form/TextFieldSkeleton";
import { FormDataObject, FormField } from "./FormCrud";
import {
  getNestedValue,
  IsDynamicOptions,
  setNestedValue,
} from "@/utils/globalUtils";
import DurationField from "../form/DurationField";
import { authTokenType } from "@/hooks/useFetchAPI";
import BarcodeScanner from "../form/BarcodeScanner";
import YearField from "../form/YearField";

export interface RenderFieldProps {
  field: FormField;
  formData: FormDataObject;
  setFormData: Dispatch<SetStateAction<FormDataObject>>;
  authToken: authTokenType;
  errorForm: Record<string, string>;
  optionsMap: Record<string, SelectOption[]>;
  loading: boolean;
}
const RenderField = ({
  field,
  formData,
  setFormData,
  authToken,
  errorForm,
  optionsMap,
  loading,
}: RenderFieldProps) => {
  
  useEffect(() => {
    if (field.fieldType !== "select") return;
    if (!field.rowGetters) return;

    const currentValue = formData[field.name];
    if (!currentValue) return;

    const options = optionsMap[field.name] ?? [];
    const selectedRow = options.find(
      (opt: any) => opt.value === currentValue,
    )?.row;

    if (!selectedRow) return;

    const extraValues =
      field.rowGetters.reduce((acc: any, getter: any) => {
        if (!getter.key || !getter.source) return acc;

        const value = getNestedValue(selectedRow, getter.source);
        if (value === undefined) return acc;

        setNestedValue(acc, getter.key, value);
        return acc;
      }, {}) || {};

    setFormData((prev: any) => {
      const isDifferent = Object.keys(extraValues).some(
        (key) => prev[key] !== extraValues[key],
      );

      if (!isDifferent) return prev;

      return {
        ...prev,
        ...extraValues,
      };
    });
  }, [
    field.fieldType,
    field.name,
    field.rowGetters,
    formData,
    optionsMap,
    setFormData,
  ]);

  if (
    field.fieldType === "text" ||
    field.fieldType === "password" ||
    field.fieldType === "number" ||
    field.fieldType === "email"
  ) {
    return (
      <TextField
        type={field.fieldType}
        isPasswordShowable={field.fieldType === "password"}
        className="rounded px-2 py-1 w-full"
        required={field.required}
        disabled={field.disabled}
        readOnly={field.readOnly}
        value={(formData[field.name] as string) ?? ""}
        onChange={(val) => setFormData((prev) => ({ ...prev, [field.name]: val }))}
        id={""}
        label={field.label}
        error={errorForm[field.name]}
        placeholder={field.placeholder}
      />
    );
  }

  if (field.fieldType === "textarea") {
    return (
      <TextAreaField
        className="rounded px-2 py-1 w-full"
        required={field.required}
        disabled={field.disabled}
        value={(formData[field.name] as string) ?? ""}
        onChange={(val) => setFormData((prev) => ({ ...prev, [field.name]: val }))}
        id={""}
        label={field.label}
        error={errorForm[field.name]}
        placeholder={field.placeholder}
      />
    );
  }

  if (field.fieldType === "date") {
    return (
      <DateField
        className="rounded px-2 py-1 w-full"
        required={field.required}
        disabled={field.disabled}
        value={(formData[field.name] as string) ?? ""}
        onChange={(val) => setFormData((prev) => ({ ...prev, [field.name]: val }))}
        id={""}
        label={field.label}
        error={errorForm[field.name]}
        placeholder={field.placeholder}
      />
    );
  }

  if (field.fieldType === "dateyear") {
    return (
      <YearField
        className="rounded px-2 py-1 w-full"
        required={field.required}
        disabled={field.disabled}
        value={(formData[field.name] as string | number) ?? ""}
        onChange={(val) => setFormData((prev) => ({ ...prev, [field.name]: val }))}
        id={""}
        label={field.label}
        error={errorForm[field.name]}
        placeholder={field.placeholder}
        minYear={field.validation?.min}
        maxYear={field.validation?.max}
      />
    );
  }

  if (field.fieldType === "datetime") {
    return (
      <DateTimeField
        className="rounded px-2 py-1 w-full"
        required={field.required}
        disabled={field.disabled}
        value={(formData[field.name] as string) ?? ""}
        onChange={(val) => setFormData((prev) => ({ ...prev, [field.name]: val }))}
        id={""}
        label={field.label}
        error={errorForm[field.name]}
        placeholder={field.placeholder}
      />
    );
  }

  if (field.fieldType === "duration") {
    return (
      <DurationField
        className="rounded px-2 py-1 w-full"
        required={field.required}
        disabled={field.disabled}
        value={(formData[field.name] as number) ?? ""}
        onChange={(val) => setFormData((prev) => ({ ...prev, [field.name]: val }))}
        id={""}
        label={field.label}
        error={errorForm[field.name]}
        placeholder={field.placeholder}
      />
    );
  }

  if (field.fieldType === "select" && !loading) {
    return (
      <SelectField
        className="rounded px-2 py-1 w-full"
        required={field.required}
        disabled={field.disabled}
        value={(formData[field.name] as string) ?? ""}
        onChange={(val, _label, row) => {
          const extraValues =
            field.rowGetters?.reduce((acc: any, getter: any) => {
              if (!getter.key || !getter.source) return acc;
              const value = getNestedValue(row, getter.source);
              if (value === undefined) return acc;
              setNestedValue(acc, getter.key, value);
              return acc;
            }, {}) || {};
          const clearedValues = Object.fromEntries(
            (field.clearFieldsOnChange ?? []).map((fieldName) => [
              fieldName,
              null,
            ]),
          );

          setFormData((prev) => ({
            ...prev,
            ...clearedValues,
            [field.name]: val,
            ...extraValues,
          }));
        }}
        valueKey={field.name}
        labelKey={
          IsDynamicOptions(field.options) ? field.options.labelKey : undefined
        }
        tableProps={field.tableProps}
        id={""}
        label={field.label}
        options={(() => {
          const rawVal = formData[field.name];
          const baseOptions = optionsMap[field.name] ?? [];
          if (rawVal !== undefined && rawVal !== null && rawVal !== "") {
            const exists = baseOptions.some((opt) => String(opt.value) === String(rawVal));
            if (!exists) {
              const nameCandidates = [
                field.name.replace(/_id$/, "_name"),
                field.name.replace(/_user_id$/, "_name"),
                field.name.replace(/_id$/, "_title"),
                field.name.replace(/_id$/, "_code"),
              ];
              for (const cand of nameCandidates) {
                if (formData[cand] && typeof formData[cand] === "string") {
                  return [{ value: rawVal as string | number, label: formData[cand] as string }, ...baseOptions];
                }
              }
            }
          }
          return baseOptions;
        })()}
        error={errorForm[field.name]}
        placeholder={field.placeholder}
      ></SelectField>
    );
  }

  if (field.fieldType === "select" && loading) {
    return <TextFieldSkeleton isLabel />;
  }

  if (field.fieldType === "chainselect" && field.chains) {
    const errors = field.chains.reduce(
      (acc, chain) => {
        acc[chain.name] = errorForm[chain.name];
        return acc;
      },
      {} as Record<string, any>,
    );
    return (
      <ChainSelectField
        chains={field.chains}
        valueMap={formData}
        onChange={(newValues) =>
          setFormData((prev) => ({
            ...prev,
            ...newValues,
          }))
        }
        required={field.required}
        disabled={field.disabled}
        authTokenType={authToken}
        errors={errors}
      />
    );
  }
  if (field.fieldType === "boolean") {
    return (
      <SwitchField
        className=""
        required={field.required}
        disabled={field.disabled}
        value={
          formData[field.name] === true || formData[field.name] === 1
            ? 1
            : formData[field.name] === false || formData[field.name] === 0
              ? 0
              : 0
        }
        onChange={(val) => setFormData((prev) => ({ ...prev, [field.name]: val === 1 }))}
        id={field.name}
        label={field.label}
        error={""}
      />
    );
  }
  if (field.fieldType === "uploadimage") {
    return (
      <UploadFileField
        folder={field.fileFolder ?? "/admin"}
        authToken={authToken}
        allowedTypes={["image", "document"]}
        maxSizeMB={10}
        onUploaded={(url) => {
          setFormData((prev) => ({
            ...prev,
            [field.name]: url,
          }));
        }}
        value={(formData[field.name] as string) ?? ""}
        id={""}
        label={field.label}
        error={errorForm[field.name]}
      />
    );
  }
  if (field.fieldType === "barcodescanner") {
    return (
      <BarcodeScanner
        onScan={(value) => {
          const parsed: Record<string, string | number> = {};
          // truck_id:2;driver_id:8; value
          value
            .split(";")
            .filter(Boolean)
            .forEach((pair) => {
              const [key, val] = pair.split(":");
              if (key && val) {
                const cleanVal = val.trim();
                parsed[key.trim()] = isNaN(Number(cleanVal))
                  ? cleanVal
                  : Number(cleanVal);
              }
            });

          setFormData((prev: any) => ({
            ...prev,
            ...parsed,
          }));
        }}
      />
    );
  }
};

export default RenderField;
