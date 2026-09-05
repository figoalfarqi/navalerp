/* eslint-disable @typescript-eslint/no-explicit-any */
import { useCallback, useEffect, useState, useRef } from "react";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { useDebounceFn } from "@/hooks/useDebounceFn";
import { FormField, FormDataObject } from "@/components/formCrud/FormCrud";
import { ResolveUrl } from "@/utils/globalUtils";

export function useCheckUniqueField({
  fields,
  formData,
  parentFormData,
}: {
  fields: FormField[];
  formData: FormDataObject;
  parentFormData?: FormDataObject;
}) {
  const { getAPI } = useFetchAPI();

  const [errorFormCheckUnique, setErrorFormCheckUnique] =
    useState<Record<string, string>>({});
  const [isLoadingCheckUniqu, setIsLoadingCheckUniqu] =
    useState<Record<string, boolean>>({});

  // Store previous values to compare
  const prevValuesRef = useRef<Record<string, any>>({});
  
  const mergedData = parentFormData
    ? { ...parentFormData, ...formData }
    : formData;

  // =====================================================
  // 🔥 CHECK UNIQUE FUNCTION (RETURNS BOOLEAN)
  // =====================================================
  const checkUnique = useCallback(async (field: FormField): Promise<boolean> => {
    const fieldName = field.name;
    const value = mergedData?.[fieldName];
    const minLength = 1;

    // Reset helper
    const resetError = () => {
      setErrorFormCheckUnique((prev) => {
        const { [fieldName]: _, ...rest } = prev;
        return rest;
      });
    };

    // Early returns
    if (!field.uniqueUrl) return true;

    if (!value || String(value).length < minLength) {
      resetError();
      setIsLoadingCheckUniqu((prev) => ({ ...prev, [fieldName]: false }));
      return true; // Valid because not enough chars to check
    }

    // Set loading state
    setIsLoadingCheckUniqu((prev) => ({ ...prev, [fieldName]: true }));
    resetError();

    try {
      const resolvedUrl = ResolveUrl(field.uniqueUrl, mergedData);
      const finalUrl = `${resolvedUrl}${resolvedUrl.includes("?") ? "&" : "?"}limit=1&sort=DESC&order_by=created_at`;

      const res = await getAPI<any>(finalUrl, { authToken: "admin" });
      const items = res.data?.data?.items ?? res.data?.items ?? [];

      const isUnique = items.length === 0;

      if (!isUnique) {
        setErrorFormCheckUnique((prev) => ({
          ...prev,
          [fieldName]: `${field.label} sudah digunakan`,
        }));
      }

      return isUnique;
    } catch (err) {
      console.error(`Check unique failed for ${fieldName}:`, err);
      return true; // Return true on error to not block submission
    } finally {
      setIsLoadingCheckUniqu((prev) => ({
        ...prev,
        [fieldName]: false,
      }));
    }
  }, [mergedData, getAPI]);

  // =====================================================
  // 🔥 DEBOUNCED CHECK (FOR REAL-TIME VALIDATION)
  // =====================================================
  const debouncedCheckUnique = useDebounceFn(
    async (field: FormField) => {
      await checkUnique(field);
    },
    500
  );

  // =====================================================
  // 🔥 EFFECT FOR AUTO-CHECK ON DATA CHANGE (WITH VALUE COMPARISON)
  // =====================================================
  useEffect(() => {
    const uniqueFields = fields.filter((field) => field.uniqueUrl);
    let isMounted = true;

    const runChecks = async () => {
      for (const field of uniqueFields) {
        if (!isMounted) break;
        
        const fieldName = field.name;
        const currentValue = mergedData?.[fieldName];
        const previousValue = prevValuesRef.current[fieldName];
        
        // Only check if the value has actually changed
        if (currentValue !== previousValue) {
          // Update previous value immediately to prevent multiple checks
          prevValuesRef.current[fieldName] = currentValue;
          
          // Only trigger check if value meets minimum length
          if (currentValue && String(currentValue).length >= 1) {
            await debouncedCheckUnique(field);
          } else {
            // Clear error if value is empty
            setErrorFormCheckUnique((prev) => {
              const { [fieldName]: _, ...rest } = prev;
              return rest;
            });
          }
        }
      }
    };

    runChecks();

    return () => {
      isMounted = false;
    };
  }, [mergedData, fields, debouncedCheckUnique, setErrorFormCheckUnique]);

  // Reset prevValues when fields change
  useEffect(() => {
    prevValuesRef.current = {};
  }, [fields]);

  // =====================================================
  // 🔥 MANUAL CHECK (FOR SUBMIT)
  // =====================================================
  const checkUniqueNow = useCallback(
    async (fieldName?: string): Promise<boolean> => {
      // Filter fields that need checking
      const targetFields = fieldName
        ? fields.filter((f) => f.name === fieldName && f.uniqueUrl)
        : fields.filter((f) => f.uniqueUrl);

      if (!targetFields.length) return true;

      // Run all checks in parallel
      const results = await Promise.all(
        targetFields.map((field) => checkUnique(field))
      );

      // Return false if any field is not unique
      return results.every(Boolean);
    },
    [fields, checkUnique]
  );

  // =====================================================
  // 🔥 UTILITY TO CHECK IF ANY FIELD IS LOADING
  // =====================================================
  const isAnyLoading = useCallback((): boolean => {
    return Object.values(isLoadingCheckUniqu).some(Boolean);
  }, [isLoadingCheckUniqu]);

  // =====================================================
  // 🔥 UTILITY TO GET ERROR FOR SPECIFIC FIELD
  // =====================================================
  const getFieldError = useCallback((fieldName: string): string | undefined => {
    return errorFormCheckUnique[fieldName];
  }, [errorFormCheckUnique]);

  // =====================================================
  // 🔥 UTILITY TO MANUALLY TRIGGER CHECK FOR SPECIFIC FIELD
  // =====================================================
  const triggerCheck = useCallback((fieldName: string) => {
    const field = fields.find(f => f.name === fieldName);
    if (field?.uniqueUrl) {
      debouncedCheckUnique(field);
    }
  }, [fields, debouncedCheckUnique]);

  return {
    errorFormCheckUnique,
    isLoadingCheckUniqu,
    checkUniqueNow,
    isAnyLoading,
    getFieldError,
    triggerCheck,
  };
}
