// components/driver/FilterFormTable.tsx
/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useEffect, useMemo, useState } from "react";
import Button from "@/components/form/Button";
import TextField from "@/components/form/TextField";
import TextAreaField from "@/components/form/TextAreaField";
import SelectField, { SelectOption } from "@/components/form/SelectField";
import DateTimeField from "@/components/form/DateTimeField";
import SelectButtonField from "../form/SelectButtonField";
import TextFieldSkeleton from "../form/TextFieldSkeleton";
import DateField from "../form/DateField";
import { LuChevronDown } from "react-icons/lu";
import { useIsMobile } from "@/hooks/useIsMobile";
import { useQueryParams } from "@/hooks/useQueryParams";
import DateAfterBeforeField from "../form/DateAfterBeforeField";
import RadioField from "../form/RadioField";
import ChainSelectField, { ChainItem } from "../form/ChainSelectField";
import { DynamicOptions, StaticOptions } from "../formCrud/FormCrud";
import { useDynamicSelectOptions } from "@/hooks/useDynamicSelectOptions";
import { addDays } from "@/utils/dateTime";

export type FilterField = {
  name: string;
  label?: string;
  col?: "left" | "right";
  type?: string;
  fieldType?:
    | "text"
    | "number"
    | "textarea"
    | "select"
    | "selectDate"
    | "chainselect"
    | "selectButton"
    | "radio"
    | "date"
    | "dateAfterBefore"
    | "datetime"
    | string;
  placeHolder?: string;
  placeholder?: string;
  options?: StaticOptions | DynamicOptions;
  disabled?: boolean;
  chains?: ChainItem[];
  /** optional async fetcher */
  optionsFetcher?: () => Promise<{ label: string; value: any }[]>;
};

type FilterFormTableProps = {
  defaultFilterValues: Record<string, number | string | string[]>;
  filters: FilterField[];
  tableFor: "detail" | "normal" | "select";
  onApplyFilter: (data: Record<string, any>) => void;
  onResetFilter?: () => void;
};

export default function FilterFormTable({
  defaultFilterValues,
  filters,
  tableFor,
  onApplyFilter,
  onResetFilter,
}: FilterFormTableProps) {
  const paramNotFilters = ["order_by", "sort", "global_limit"];
  const [openTab, setOpenTab] = useState("");
  const [filterFormTableData, setFilterFormTableData] =
    useState<Record<string, any>>(defaultFilterValues);

  const { dynamicOptionsMap, isLoadingDynamicOptions } =
    useDynamicSelectOptions({
      fields: filters,
      // formData,
      // parentFormData,
      authToken: "admin",
    });

  const [isLoadingOptionsFetcher, setIsLoadingOptionsFetcher] = useState<
    Record<string, boolean>
  >({});
  const [optionsFetcherMap, setOptionsFetcherMap] = useState<
    Record<string, SelectOption[]>
  >({});

  const { isMobile } = useIsMobile();
  const { getAllParams, setParams } = useQueryParams();

  // 🧠 Ambil async options kalau ada
  useEffect(() => {
    filters.forEach(async (filter) => {
      if (filter.optionsFetcher) {
        setIsLoadingOptionsFetcher((prev) => ({
          ...prev,
          [filter.name]: true,
        }));

        const opts = await filter.optionsFetcher();
        setOptionsFetcherMap((prev) => ({ ...prev, [filter.name]: opts }));
        setIsLoadingOptionsFetcher((prev) => ({
          ...prev,
          [filter.name]: false,
        }));
      }
    });
  }, []);

  const toggleOpen = (tab: string) =>
    setOpenTab((prev) => (prev == tab ? "" : tab));

  // const handleChangeFilter = (name: string, value: any) => {
  //   setFilterFormTableData((prev) => ({ ...prev, [name]: value }));
  // };

  const handleChangeFilter = (name: string, value: any) => {
    const nextState = {
      ...filterFormTableData,
      [name]: value,
    };

    setFilterFormTableData(nextState);
    return nextState;
  };

  const handleApplyFilter = (data = filterFormTableData) => {
    const activeFilters = Object.fromEntries(
      Object.entries(data).filter(
        ([_, v]) => v !== "" && v !== undefined && v !== null,
      ),
    );

    onApplyFilter(activeFilters);

    if (tableFor !== "select" && tableFor !== "detail") {
      setParams(activeFilters);
    }
  };

  // const handleApplyFilter = () => {
  //   const activeFilters = Object.fromEntries(
  //     Object.entries(filterFormTableData).filter(
  //       ([_, v]) => v !== "" && v !== undefined && v !== null,
  //     ),
  //   );
  //   onApplyFilter(activeFilters);
  //   if (tableFor != "select" && tableFor != "detail") {
  //     setParams(activeFilters);
  //   }
  // };

  const handleResetFilter = () => {
    setFilterFormTableData(defaultFilterValues);
    onResetFilter?.();
    const activeFilters = getAllParams();
    const forDeleteFilters = Object.entries(activeFilters).reduce(
      (acc, [key]) => {
        if (paramNotFilters.includes(key)) {
          return acc;
        } else {
          acc[key] = "";
        }
        return acc;
      },
      {} as Record<string, any>,
    );
    setParams(forDeleteFilters);
  };

  const activeFilterCount = Object.entries(filterFormTableData).filter(
    ([key, value]) => {
      if (value === "" || value === null || value === undefined) {
        return false;
      }
      if (
        ["project_id"].includes(key) &&
        defaultFilterValues &&
        defaultFilterValues[key] === value
      ) {
        return false;
      }
      if (
        (key === "is_active" || key === "app_user_status_id") &&
        value === "1"
      ) {
        return false;
      }
      if (
        [
          "completed_at_before",
          "completed_at_after",
          "stockpile_transport_for_page",
          "client_transport_for_page",
        ].includes(key)
      ) {
        return false;
      }
      return true;
    },
  ).length;

  const buttonHeight = 60;
  const filterFieldHeight = 86;
  const maxFilterHeight = useMemo(() => {
    const totalFields = filters.reduce((count, filter) => {
      return count + (filter.fieldType === "dateAfterBefore" ? 2 : 1);
    }, 4);
    return isMobile
      ? totalFields * filterFieldHeight + buttonHeight
      : Math.ceil(totalFields / 2) * filterFieldHeight + buttonHeight;
  }, [filters.length, isMobile]);

  useEffect(() => {
    const activeFilters = getAllParams();
    const parsedFilters = Object.entries(activeFilters).reduce(
      (acc, [key, value]) => {
        if (paramNotFilters.includes(key)) {
          return acc;
        } else if (
          key.endsWith("_id") &&
          !["app_user_status_id"].includes(key)
        ) {
          const num = parseInt(value as string, 10);
          acc[key] = isNaN(num) ? value : num;
        } else {
          acc[key] = value;
        }
        return acc;
      },
      {} as Record<string, any>,
    );

    const finalFilters = {
      ...defaultFilterValues,
      ...parsedFilters,
    };
    setFilterFormTableData(finalFilters);
    onApplyFilter(finalFilters);
  }, []);

  const renderField = (
    filter: FilterField,
    opts: { label: string; value: any }[],
  ) => {
    if (filter.fieldType === "text") {
      return (
        <TextField
          key={filter.name}
          id={filter.name}
          label={filter.label || filter.name}
          value={filterFormTableData[filter.name] || ""}
          className="w-full"
          placeholder={filter.placeHolder}
          onChange={(val) => handleChangeFilter(filter.name, val)}
          disabled={filter.disabled}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              handleApplyFilter();
            }
          }}
        />
      );
    }

    if (filter.fieldType === "number") {
      return (
        <TextField
          key={filter.name}
          id={filter.name}
          type="number"
          label={filter.label || filter.name}
          value={filterFormTableData[filter.name] || ""}
          className="w-full"
          onChange={(val) => handleChangeFilter(filter.name, val)}
          disabled={filter.disabled}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              handleApplyFilter();
            }
          }}
        />
      );
    }

    if (filter.fieldType === "textarea") {
      return (
        <TextAreaField
          key={filter.name}
          id={filter.name}
          label={filter.label || filter.name}
          value={filterFormTableData[filter.name] || ""}
          className="w-full"
          onChange={(val) => handleChangeFilter(filter.name, val)}
          disabled={filter.disabled}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              handleApplyFilter();
            }
          }}
        />
      );
    }

    if (filter.fieldType === "select") {
      return (
        <SelectField
          key={filter.name}
          id={filter.name}
          label={filter.label || filter.name}
          options={opts}
          value={filterFormTableData[filter.name] || ""}
          className="w-full"
          onChange={(val) => {
            const nextState = handleChangeFilter(filter.name, val);
            handleApplyFilter(nextState);
          }}
          disabled={filter.disabled}
        />
      );
    }

    if (filter.fieldType === "selectDate") {
      return (
        <SelectField
          key={filter.name}
          id={filter.name}
          label={filter.label || filter.name}
          options={[
            { label: "2 atau lebih hari lalu", value: "-2" },
            { label: "Kemarin", value: "-1" },
            { label: "Hari ini", value: "0" },
            { label: "Besok", value: "1" },
            { label: "Lusa", value: "2" },
            { label: "3 atau lebih hari lagi", value: "3" },
          ]}
          value={filterFormTableData[filter.name] || ""}
          className="w-full"
          onChange={(val) => {
            const today = new Date();
            const dayOffset = Number(val);
            const target = addDays(today, dayOffset);

            const nextState = {
              ...filterFormTableData,
              [filter.name]: val,
              [`${filter.name}_after`]: new Date(target).toISOString(),
              [`${filter.name}_before`]: new Date(target).toISOString(),
            };

            if (val === "-2") {
              nextState[`${filter.name}_after`] = "";
            } else if (val === "3") {
              nextState[`${filter.name}_before`] = "";
            } else if (val == null) {
              nextState[`${filter.name}_after`] = "";
              nextState[`${filter.name}_before`] = "";
            }

            setFilterFormTableData(nextState);
            handleApplyFilter(nextState);
          }}
          disabled={filter.disabled}
        />
      );
    }

    if (filter.fieldType === "chainselect" && filter.chains) {
      return (
        <ChainSelectField
          key={filter.name}
          chains={filter.chains}
          valueMap={filterFormTableData}
          onChange={(newValues) => {
            setFilterFormTableData((prev) => ({
              ...prev,
              ...newValues,
            }));
            const nextState = {
              ...filterFormTableData,
              ...newValues,
            };
            handleApplyFilter(nextState);
          }}
          authTokenType="admin"
          disabled={filter.disabled}
        />
      );
    }

    if (filter.fieldType === "selectButton") {
      return (
        <SelectButtonField
          key={filter.name}
          id={filter.name}
          label={filter.label || filter.name}
          options={opts}
          value={filterFormTableData[filter.name] || ""}
          className="w-full"
          onChange={(val) => {
            const nextState = handleChangeFilter(filter.name, val);
            handleApplyFilter(nextState);
          }}
          disabled={filter.disabled}
        />
      );
    }

    if (filter.fieldType === "radio") {
      return (
        <RadioField
          key={filter.name}
          id={filter.name}
          label={filter.label || filter.name}
          options={opts}
          value={filterFormTableData[filter.name] || ""}
          className="w-full"
          onChange={(val) => {
            const nextState = handleChangeFilter(filter.name, val);
            handleApplyFilter(nextState);
          }}
        />
      );
    }

    if (filter.fieldType === "date") {
      return (
        <DateField
          key={filter.name}
          id={filter.name}
          label={filter.label || filter.name}
          value={`${filterFormTableData[filter.name]}` || ""}
          className="w-full"
          onChange={(val) => {
            const nextState = handleChangeFilter(filter.name, val);
            handleApplyFilter(nextState);
          }}
          disabled={filter.disabled}
        />
      );
    }

    if (filter.fieldType === "dateAfterBefore") {
      return (
        <DateAfterBeforeField
          key={filter.name}
          afterName={`${filter.name}_after`}
          beforeName={`${filter.name}_before`}
          afterLabel={`${filter.label || filter.name} After`}
          beforeLabel={`${filter.label || filter.name} Before`}
          afterValue={filterFormTableData[`${filter.name}_after`]}
          beforeValue={filterFormTableData[`${filter.name}_before`]}
          onChange={(name, value) => {
            const nextState = {
              ...filterFormTableData,
              [name]: value,
            };

            setFilterFormTableData(nextState);
            handleApplyFilter(nextState);
          }}
        />
      );
    }

    if (filter.fieldType === "datetime") {
      return (
        <DateTimeField
          key={filter.name}
          id={filter.name}
          label={filter.label || filter.name}
          value={filterFormTableData[filter.name] || ""}
          className="w-full"
          onChange={async (val) => {
            const nextState = handleChangeFilter(filter.name, val);
            handleApplyFilter(nextState);
          }}
          disabled={filter.disabled}
        />
      );
    }
  };

  const renderFilterList = (list: FilterField[]) => (
    <div className="space-y-4">
      {list.map((filter) => {
        const opts =
          dynamicOptionsMap[filter.name] ??
          optionsFetcherMap[filter.name] ??
          filter.options ??
          [];
        const isLoadingSelect =
          ["select", "selectButton"].includes((filter.fieldType as string) ?? "") &&
          !Array.isArray(opts);

        if (
          isLoadingDynamicOptions[filter.name] ||
          isLoadingOptionsFetcher[filter.name] ||
          isLoadingSelect
        ) {
          return <TextFieldSkeleton isLabel key={`loading-${filter.name}`} />;
        }

        return renderField(filter, opts);
      })}
    </div>
  );

  return (
    <div className="mb-1 w-full">
      <div
        className={`flex z-20 cursor-pointer select-none`}
        onClick={() => toggleOpen("filter")}
      >
        <div className={`flex flex-col items-center py-1 gap-1 relative`}>
          <div>Filter</div>
          {activeFilterCount > 0 && (
            <div className="absolute translate-x-[240%] -translate-y-[30%] w-5 h-5 rounded-full text-center bg-blue-100 text-blue-600 text-sm leading-snug">
              {activeFilterCount}
            </div>
          )}
        </div>
        <LuChevronDown
          className={`text-3xl py-1 text-gray-600 transform transition-transform duration-300 ${
            openTab ? "rotate-180" : "rotate-0"
          }`}
        />
      </div>
      <div
        className={`transition-all duration-300 ease-in-out overflow-hidden ${
          openTab == "filter" ? "mt-2" : ""
        }`}
        style={{
          maxHeight: openTab == "filter" ? `${maxFilterHeight}px` : "0px",
        }}
      >
        <div
          className={`grid ${isMobile ? "grid-cols-1" : "grid-cols-2"} gap-4`}
        >
          {isMobile ? (
            <div className="col-span-2">{renderFilterList(filters)}</div>
          ) : (
            <>
              {renderFilterList(filters.filter((f) => f.col === "left"))}
              {renderFilterList(filters.filter((f) => f.col === "right"))}
            </>
          )}
          {/* Buttons */}
          <div className="col-span-2 flex justify-end gap-3 pb-2">
            <Button
              id="reset-filters"
              variant="gray-outline"
              onClick={handleResetFilter}
              className="text-sm"
            >
              Reset
            </Button>
            <Button
              id="apply-filters"
              variant="blue-solid"
              onClick={handleApplyFilter}
              className="text-sm"
            >
              Apply
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}
