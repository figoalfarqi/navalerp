/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useMemo, useState } from "react";
import { authTokenType, useFetchAPI } from "@/hooks/useFetchAPI";
import {
    IsDependencySatisfied,
    ResolveUrl,
    IsDynamicOptions,
    resolveLabelValueKey,
} from "@/utils/globalUtils";
import { FormDataObject, FormField } from "@/components/formCrud/FormCrud";
import { FilterField } from "@/components/table/FilterFormTable";
import { SelectOption } from "@/components/form/SelectField";

export const useDynamicSelectOptions = ({
    fields,
    formData,
    parentFormData,
    authToken
}: {
    fields: FormField[] | FilterField[];
    formData?: FormDataObject;
    parentFormData?: FormDataObject;
    authToken: authTokenType;
}) => {
    const { getAPI } = useFetchAPI();

    const [dynamicOptionsMap, setDynamicOptionsMap] = useState<Record<string, SelectOption[]>>(
        {}
    );
    const [isLoadingDynamicOptions, setIsLoadingDynamicOptions] = useState<Record<string, boolean>>({});

    const getDependencyValues = (
        dependsOn: string[] | undefined,
        parentFormData?: FormDataObject,
        formData?: FormDataObject
    ) => {
        if (!dependsOn?.length) return {};

        const source = parentFormData
            ? { ...parentFormData, ...formData }
            : formData ?? {};

        return dependsOn.reduce((acc, key) => {
            acc[key] = source[key];
            return acc;
        }, {} as Record<string, any>);
    };

    const dependencySignature = useMemo(
        () =>
            fields
                .filter(
                    (f) =>
                        ["select", "selectButton"].includes((f.fieldType as string) ?? "") &&
                        IsDynamicOptions(f.options)
                )
                .map((f) => {
                    if (!IsDynamicOptions(f.options)) return "";
                    return JSON.stringify({
                        name: f.name,
                        url: f.options.url,
                        excludedValues: f.options.excludedValues ?? [],
                        dependencies: getDependencyValues(
                            f.options.dependsOn,
                            parentFormData,
                            formData
                        ),
                    });
                })
                .join("|"),
        [fields, formData, parentFormData]
    );

    useEffect(() => {
        let cancelled = false;

        const loadOptions = async () => {
            const nextOptions: Record<string, SelectOption[]> = {};

            for (const field of fields) {
                if (!["select", "selectButton"].includes((field.fieldType as string) ?? "") || !field.options) continue;

                if (Array.isArray(field.options) && !IsDynamicOptions(field.options)) {
                    nextOptions[field.name] = field.options;
                    continue;
                }
                setIsLoadingDynamicOptions((prev) => ({ ...prev, [field.name]: true }));
                const dynamicOpts = field.options;
                if (
                    !IsDependencySatisfied(
                        dynamicOpts.dependsOn,
                        parentFormData,
                        formData
                    )
                ) {
                    nextOptions[field.name] = [];
                    continue;
                }

                const resolvedUrl = formData ? ResolveUrl(
                    dynamicOpts.url,
                    parentFormData ? { ...parentFormData, ...formData } : formData
                ) : dynamicOpts.url;

                try {
                    const res = await getAPI<any>(resolvedUrl, {
                        authToken: authToken,
                    });

                    const items = res.data?.data?.items ?? res.data?.items ?? [];
                    const {
                        valueKey,
                        labelKey,
                        excludedValues = [],
                    } = dynamicOpts;

                    nextOptions[field.name] = items
                        .filter(
                            (item: any) =>
                                !excludedValues.includes(item?.[valueKey]),
                        )
                        .map((item: any) => ({
                            value: item?.[valueKey],
                            label: resolveLabelValueKey(labelKey, item),
                            row: item,
                        }));
                    setIsLoadingDynamicOptions((prev) => ({ ...prev, [field.name]: false }));

                } catch (err) {
                    console.error(`Failed to fetch options for ${field.name}`, err);
                    nextOptions[field.name] = [];
                    setIsLoadingDynamicOptions((prev) => ({ ...prev, [field.name]: false }));
                }
            }

            if (!cancelled) {
                setDynamicOptionsMap(nextOptions);
            }
        };

        loadOptions();

        return () => {
            cancelled = true;
        };
    }, [dependencySignature]);

    return { dynamicOptionsMap, isLoadingDynamicOptions };
};
