/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useEffect, useState, useRef } from "react";
import SelectField from "./SelectField";
import { authTokenType, useFetchAPI } from "@/hooks/useFetchAPI";
import TextFieldSkeleton from "./TextFieldSkeleton";
import { rowGetter } from "../formCrud/FormCrud";
import {
  getNestedValue,
  resolveLabelValueKey,
  setNestedValue,
} from "@/utils/globalUtils";

export interface ChainItem {
  name: string;
  label: string;
  url: string; // bisa berisi {city_id}, {district_id}, dll
  labelKey: string;
  valueKey: string;
  rowGetters?: rowGetter[];
}

interface ChainSelectFieldProps {
  chains: ChainItem[];
  valueMap: { [key: string]: any };
  onChange: (newValues: { [key: string]: any }) => void;
  required?: boolean;
  disabled?: boolean;
  authTokenType: authTokenType;
  errors?: Record<string, string | undefined>;
}

export default function ChainSelectField({
  chains,
  valueMap,
  onChange,
  required,
  disabled,
  authTokenType,
  errors = {},
}: ChainSelectFieldProps) {
  const { getAPI } = useFetchAPI();
  const [options, setOptions] = useState<{ [key: string]: any[] }>({});
  const [loading, setLoading] = useState<{ [key: string]: boolean }>({});
  const prevValueMap = useRef<{ [key: string]: any }>({});
  // console.log(valueMap)
  // Load first select
  useEffect(() => {
    const first = chains[0];
    if (!first) return;

    const loadFirst = async () => {
      setLoading((p) => ({ ...p, [first.name]: true }));
      try {
        const res = await getAPI<any>(first.url, { authToken: authTokenType });
        const items = res.data?.data?.items || res.data?.items || [];
        setOptions((p) => ({
          ...p,
          [first.name]: items.map((item: any) => ({
            // label: item[first.labelKey],
            label: resolveLabelValueKey(first.labelKey, item),

            value: item[first.valueKey],
            row: item,
          })),
        }));
      } catch (err) {
        console.error(`Failed to fetch ${first.name}`, err);
      } finally {
        setLoading((p) => ({ ...p, [first.name]: false }));
      }
    };

    loadFirst();
  }, [chains[0]]);

  // Observe parent changes → fetch next level
  useEffect(() => {
    const checkAndLoadNext = async (idx: number) => {
      const current = chains[idx];
      const next = chains[idx + 1];
      if (!next) return;

      const prevVal = prevValueMap.current[current.name];
      const currVal = valueMap[current.name];

      // parent tidak berubah → skip
      if (prevVal === currVal) return;

      // parent dikosongkan → reset child & options
      if (!currVal) {
        setOptions((p) => ({ ...p, [next.name]: [] }));
        if (valueMap[next.name]) {
          onChange({ [next.name]: null });
        }
        return;
      }

      // build url
      let url = next.url;
      chains.forEach((c) => {
        url = url.replace(`{${c.name}}`, valueMap[c.name] ?? "");
      });

      setLoading((p) => ({ ...p, [next.name]: true }));

      try {
        const res = await getAPI<any>(url, { authToken: authTokenType });
        const items = res.data?.data?.items || res.data?.items || [];

        const newOptions = items.map((item: any) => ({
          // label: item[next.labelKey],
          label: resolveLabelValueKey(next.labelKey, item),
          value: item[next.valueKey],
          row: item,
        }));

        setOptions((p) => ({
          ...p,
          [next.name]: newOptions,
        }));

        // 🔥 CEK VALUE CHILD
        const currentChildValue = valueMap[next.name];
        const stillValid = newOptions.some(
          (opt: any) => opt.value === currentChildValue
        );

        if (currentChildValue && !stillValid) {
          onChange({ [next.name]: null });
        }
      } catch (err) {
        console.error(`Failed to fetch ${next.name}`, err);
      } finally {
        setLoading((p) => ({ ...p, [next.name]: false }));
      }
    };

    // Check all parent-child pairs
    chains.forEach((_, idx) => {
      if (idx < chains.length - 1) checkAndLoadNext(idx);
    });

    // Simpan state lama untuk perbandingan
    prevValueMap.current = { ...valueMap };
  }, [valueMap, chains]);

  return (
    <div className="space-y-4">
      {chains.map((chain, idx) => {
        if (loading[chain.name]) {
          return <TextFieldSkeleton key={chain.name} isLabel />;
        }
        const isDisabled =
          disabled || (idx > 0 && !valueMap[chains[idx - 1].name]);

        const placeholder =
          idx > 0 && !valueMap[chains[idx - 1].name]
            ? `Pilih ${chains[idx - 1].label} terlebih dahulu`
            : `Pilih ${chain.label}`;

        return (
          <SelectField
            key={chain.name}
            label={chain.label}
            required={required}
            disabled={isDisabled}
            value={valueMap[chain.name] || null}
            // onChange={(val, label, row) => onChange({ [chain.name]: val })}
            onChange={(val, label, row) => {
              const extraValues =
                chain.rowGetters?.reduce((acc: any, getter: any) => {
                  if (!getter.key || !getter.source) return acc;
                  const value = getNestedValue(row, getter.source);
                  if (value === undefined) return acc;
                  setNestedValue(acc, getter.key, value);
                  return acc;
                }, {}) || {};
              // console.log("uhtlk", extraValues, chain.rowGetters, row)

              onChange({
                [chain.name]: val,
                // [`${chain.name}_name`]: label,
                ...extraValues,
              });
            }}
            options={options[chain.name] || []}
            placeholder={placeholder}
            id={`select-${chain.name}`}
            className="rounded px-2 py-1 w-full"
            error={errors[chain.name]}
          />
        );
      })}
    </div>
  );
}
