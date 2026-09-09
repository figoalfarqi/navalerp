/* eslint-disable @typescript-eslint/no-explicit-any */
import { useIsMobile } from "@/hooks/useIsMobile";
import { FormDataObject, FormField } from "./FormCrud";
import { Dispatch, SetStateAction, useMemo } from "react";
import RenderField from "./RenderField";
import {
  GetDependsOnLabel,
  IsDependencySatisfied,
  IsDynamicOptions,
  ResolveUrl,
} from "@/utils/globalUtils";
import { useDynamicSelectOptions } from "@/hooks/useDynamicSelectOptions";
import { SelectOption } from "../form/SelectField";
import { authTokenType } from "@/hooks/useFetchAPI";

export interface FieldLayoutProps {
  fields: FormField[];
  excludeOptions?: Record<string, Array<string | number>>;
  parentFormData?: FormDataObject;
  formData: FormDataObject;
  authToken: authTokenType;
  tablesDependsOn?: string[];
  tablesAddedData?: any;
  setFormData: Dispatch<SetStateAction<FormDataObject>>;
  errorForm: Record<string, string>;
}

const FieldLayout = ({
  fields,
  excludeOptions,
  parentFormData,
  formData,
  authToken,
  tablesDependsOn,
  tablesAddedData,
  setFormData,
  errorForm,
}: FieldLayoutProps) => {
  const { isMobile } = useIsMobile();

  const { dynamicOptionsMap, isLoadingDynamicOptions } =
    useDynamicSelectOptions({
      fields,
      formData,
      parentFormData,
      authToken
    });



  /* =========================
   * FILTER OPTIONS (NO FETCH)
   * ========================= */
  const filteredDynamicOptionsMap = useMemo(() => {
    if (!excludeOptions) return dynamicOptionsMap;

    const result: Record<string, SelectOption[]> = {};

    for (const [fieldName, opts] of Object.entries(dynamicOptionsMap)) {
      const excluded = excludeOptions[fieldName] ?? [];
      result[fieldName] = excluded.length
        ? opts.filter((opt) => !excluded.includes(opt.value))
        : opts;
    }

    return result;
  }, [dynamicOptionsMap, excludeOptions]);

  /* =========================
   * RENDER
   * ========================= */
  const renderField = (field: FormField) => {
    if (
      field.fieldType === "select" &&
      IsDynamicOptions(field.options) &&
      field.options.dependsOn?.length
    ) {
      const ready = IsDependencySatisfied(
        field.options.dependsOn,
        parentFormData,
        formData,
      );

      // console.log(
      //   "field.options.dependsOn",
      //   field.name,
      //   field.options.dependsOn,
      //   parentFormData,
      //   ready,
      // );

      if (!ready) {
        return (
          <RenderField
            field={{
              ...field,
              disabled: true,
              placeholder: `Pilih ${GetDependsOnLabel(
                field.options.dependsOn,
              )} terlebih dahulu`,
            }}
            formData={formData}
            authToken={authToken}
            setFormData={setFormData}
            errorForm={errorForm}
            optionsMap={filteredDynamicOptionsMap}
            loading={false}
          />
        );
      }
    }

    return (
      <RenderField
        field={field}
        formData={formData}
        authToken={authToken}
        setFormData={setFormData}
        errorForm={errorForm}
        optionsMap={filteredDynamicOptionsMap}
        loading={isLoadingDynamicOptions[field.name]}
      />
    );
  };

  const renderFieldList = (list: FormField[]) => (
    <div className={`${isMobile ? "col-span-2" : "col-span-1"} space-y-4`}>
      {list.map((field) => {
        // console.log(
        //   "fields in FieldLayout",
        //   field.tableProps?.default_filter_values?.customer_id &&
        //     ResolveUrl(
        //       field.tableProps?.default_filter_values?.customer_id as string,
        //       { ...formData, ...parentFormData }
        //     )
        // );
        const resolvedDefaultFilterValues = field.tableProps
          ?.default_filter_values
          ? Object.fromEntries(
              Object.entries(field.tableProps.default_filter_values).map(
                ([key, value]) => {
                  if (typeof value !== "string") return [key, value];

                  const resolved = ResolveUrl(value, {
                    ...formData,
                    ...parentFormData,
                  });

                  const asNumber = Number(resolved);
                  return [key, Number.isNaN(asNumber) ? resolved : asNumber];
                },
              ),
            )
          : undefined;
        let disabledData = {};
        if (
          tablesDependsOn &&
          tablesDependsOn.includes(field.name) &&
          tablesAddedData &&
          tablesAddedData.length > 0
        ) {
          disabledData = { disabled: true };
        }

        const resolvedField = {
          ...field,
          ...disabledData,
          tableProps: field.tableProps
            ? {
                ...field.tableProps,
                default_filter_values: resolvedDefaultFilterValues,
              }
            : undefined,
        };
        return <div key={field.name}>{renderField(resolvedField)}</div>;
      })}
    </div>
  );

  if (isMobile) {
    return (
      <div className="col-span-1 md:col-span-2">
        {renderFieldList(fields)}
      </div>
    );
  }

  const leftFields: FormField[] = [];
  const rightFields: FormField[] = [];

  const hasExplicitCol = fields.some((f) => f.col === "left" || f.col === "right");

  if (hasExplicitCol) {
    fields.forEach((field) => {
      if (field.col === "left") {
        leftFields.push(field);
      } else if (field.col === "right") {
        rightFields.push(field);
      } else {
        if (leftFields.length <= rightFields.length) {
          leftFields.push(field);
        } else {
          rightFields.push(field);
        }
      }
    });
  } else {
    fields.forEach((field, index) => {
      if (index % 2 === 0) {
        leftFields.push(field);
      } else {
        rightFields.push(field);
      }
    });
  }

  return (
    <>
      {renderFieldList(leftFields)}
      {renderFieldList(rightFields)}
    </>
  );
};

export default FieldLayout;
