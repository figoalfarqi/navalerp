import { useIsMobile } from "@/hooks/useIsMobile";
import { FormDataObject, FormDataValue, FormField } from "./FormCrud";
import FieldLayout from "./FieldLayout";
import { Mode } from "@/hooks/useDetails";
import { LuEllipsis, LuPlus } from "@/components/icons";
import Button from "../form/Button";
import { Fragment, useMemo, useState } from "react";
import { IsPrimitive, IsStaticOptions } from "@/utils/globalUtils";
import { AnimatePresence, Reorder, useDragControls } from "framer-motion";
import { useNanoId } from "@/hooks/useNanoId";
import { useToast } from "../ToastContext";
import { DividerFieldNames, SequenceFieldNames } from "@/consta/DetailFormData";
import DraggableItem from "./DraggableItem";
import { buildExcludeOptions } from "@/utils/buildExcludeOptions";
import { SelectOption } from "../form/SelectField";
import { addMinutes } from "@/utils/dateTime";
import { handleCascadeArrivalTime } from "@/utils/detailForm";

interface DetailFormProps {
  title: string;
  table_key: string;
  table_name: string;
  fields: FormField[];
  mode: Mode;
  maxItems: number;
  formData: FormDataObject;
  parentFormData?: FormDataObject;
  setFormData: React.Dispatch<React.SetStateAction<FormDataObject>>;
  errorForm: Record<string, string>;
}

const DetailForm = ({
  title,
  table_key,
  table_name,
  fields,
  mode,
  maxItems,
  formData,
  parentFormData,
  setFormData,
  errorForm,
}: DetailFormProps) => {
  const { isMobile } = useIsMobile();
  const { generateId } = useNanoId();
  const { showToast } = useToast();

  type ExcludeOptions = Record<string, Array<string | number>>;
  // const [excludeOptions, setExcludeOptions] = useState<ExcludeOptions>(() =>
  //   Object.fromEntries(ExcludedFieldNames.map((field) => [field, []]))
  // );

  const excludeOptions = useMemo(() => {
    const list = (formData[table_name] as FormDataObject[]) ?? [];
    return buildExcludeOptions(list);
  }, [formData, table_name]);

  const dividerField = fields.find((field) =>
    DividerFieldNames.includes(field.name),
  );

  const sequenceField = fields.find((field) =>
    SequenceFieldNames.includes(field.name),
  );

  // console.log("formData[table_name]", formData[table_name]);

  const normalizeSequence = (
    list: FormDataObject[],
    sequenceFieldName: string,
  ) => {
    return list.map((item, index) => {
      const updated = { ...item };
      updated[sequenceFieldName] = index + 1; // mulai dari 1
      return updated;
    });
  };

  // ➕ tambah item
  const handleAdd = (initialData: Record<string, string | number>) => {
    if (
      (formData[table_name]
        ? (formData[table_name] as FormDataObject[]).length
        : 0) < maxItems
    ) {
      setFormData((prev) => ({
        ...prev,
        [table_name]: [
          ...((prev[table_name] as FormDataObject[]) ?? []),
          { ...initialData, flag: "added", __key: generateId() },
        ],
      }));
    } else {
      showToast(3000, "error", `Maximum ${maxItems} items allowed.`);
    }
  };

  // ❌ hapus item
  const handleDelete = ({
    index,
    dividerFieldName,
    option,
    sequenceFieldName,
  }: {
    index: number;
    dividerFieldName?: string;
    option?: SelectOption;
    sequenceFieldName?: string;
  }) => {
    const list = formData[table_name] as FormDataObject[];
    let item;
    let dividedList;
    if (dividerFieldName && option) {
      dividedList = list.filter(
        (li) => (li as FormDataObject)?.[dividerFieldName] === option.value,
      );
      item = dividedList?.[index];
    } else {
      item = list?.[index];
    }
    if (!item) return;

    // 🗑️ TANDAI SEBAGAI DELETED (jika bukan item baru)
    if (item.flag !== "added") {
      setFormData((prev) => ({
        ...prev,
        [`deleted_${table_name.slice(0, -1)}_ids`]: [
          ...((prev[
            `deleted_${table_name.slice(0, -1)}_ids`
          ] as FormDataValue[]) ?? []),
          item[table_key],
        ],
      }));
    }

    // ❌ HAPUS ITEM DARI LIST
    if (!dividerFieldName || !option) {
      // ===== MODE TANPA DIVIDER =====

      setFormData((prev) => {
        const newList = (prev[table_name] as FormDataObject[]).filter(
          (_: any, i: number) => i !== index,
        );
        return {
          ...prev,
          [table_name]: sequenceFieldName
            ? normalizeSequence(newList, sequenceFieldName)
            : newList,
        };
      });
    } else {
      // ===== MODE DENGAN DIVIDER =====

      setFormData((prev) => {
        const list = prev[table_name] as FormDataObject[];
        const grouped = list.filter(
          (li) => li?.[dividerFieldName] === option.value,
        );
        const others = list.filter(
          (li) => li?.[dividerFieldName] !== option.value,
        );

        const groupWithDeleted = grouped.filter(
          (_: any, i: number) => i !== index,
        );

        return {
          ...prev,
          [table_name]: [
            ...others,
            ...(sequenceFieldName
              ? normalizeSequence(groupWithDeleted, sequenceFieldName)
              : groupWithDeleted),
          ],
        };
      });
    }
  };

  const createItemSetter =
    ({
      index,
      currentItem,
      dividerFieldName,
      option,
    }: {
      index: number;
      currentItem: FormDataObject;
      dividerFieldName?: string;
      option?: SelectOption;
    }) =>
    (value: React.SetStateAction<FormDataObject>) => {
      const resolvedValue =
        typeof value === "function" ? value(currentItem) : value;
      const updatedValue =
        resolvedValue.flag !== "added"
          ? { ...resolvedValue, flag: "updated" }
          : resolvedValue;
      handleUpdateItem({
        index,
        newItem: updatedValue,
        dividerFieldName,
        option,
      });
    };

  type CascadeParams = {
    prevItem: FormDataObject;
    newItem: FormDataObject;
    grouped: FormDataObject[];
    others: FormDataObject[];
    intervalMinutes?: number;
  };

  // 📝 update item (dipakai FieldLayout)
  const handleUpdateItem = ({
    index,
    newItem,
    dividerFieldName,
    option,
  }: {
    index: number;
    newItem: FormDataObject;
    dividerFieldName?: string;
    option?: SelectOption;
  }) => {
    setFormData((prev) => {
      const list = prev[table_name] as FormDataObject[];

      // ===== MODE TANPA DIVIDER =====
      if (!dividerFieldName || !option) {
        const newList = list.map((item, i) => (i === index ? newItem : item));
        return {
          ...prev,
          [table_name]: newList,
        };
      }

      // ===== MODE DENGAN DIVIDER =====
      const grouped = list.filter(
        (li) => li?.[dividerFieldName] === option.value,
      );
      let others = list.filter((li) => li?.[dividerFieldName] !== option.value);

      const prevItem = grouped[index] ?? {};
      let updatedGroup = grouped.map((item, i) =>
        i === index ? newItem : item,
      );

      // =============================================
      // START: hanya untuk create delivery order
      // =============================================
      let completed_at = prev.completed_at;
      const DEFAULT_INTERVAL_MINUTES = 15;

      const result = handleCascadeArrivalTime({
        prevItem,
        newItem,
        updatedGroup,
        others,
        completed_at,
        DEFAULT_INTERVAL_MINUTES,
      });

      updatedGroup = result.updatedGroup;
      others = result.others;
      completed_at = result.completed_at;

      // =============================================
      // END: hanya untuk create delivery order
      // =============================================

      return {
        ...prev,
        completed_at: completed_at,
        [table_name]: [...others, ...updatedGroup],
      };
    });
  };

  const getRowErrors = (
    errorForm: Record<string, string>,
    tableName: string,
    index: number,
    dividerFieldName?: string,
    option?: SelectOption,
  ): Record<string, string> => {
    const prefix =
      dividerFieldName && option
        ? `${tableName}.${dividerFieldName}:${option?.value}.${index}.`
        : `${tableName}.${index}.`;
    const rowErrors: Record<string, string> = {};

    Object.entries(errorForm).forEach(([key, message]) => {
      if (key.startsWith(prefix)) {
        const fieldName = key.replace(prefix, "");
        rowErrors[fieldName] = message;
      }
    });

    return rowErrors;
  };

  const getItemId = (item: FormDataObject): string =>
    (item[table_key] as string) ?? item.__key;

  const renderDataList = ({
    dataList,
    option,
    dividerFieldName,
    sequenceFieldName,
  }: {
    dataList: FormDataObject[];
    option?: SelectOption;
    dividerFieldName?: string;
    sequenceFieldName?: string;
  }) => {
    return (
      <>
        <Reorder.Group
          axis="y"
          values={dataList.map(getItemId)}
          onReorder={(newOrder) => {
            if (!sequenceFieldName) return;
            const reordered = newOrder
              .map((id) => dataList.find((item) => getItemId(item) === id))
              .filter(Boolean) as FormDataObject[];
            const normalized = normalizeSequence(reordered, sequenceFieldName);
            if (dividerFieldName && option) {
              setFormData((prev) => ({
                ...prev,
                [table_name]: [
                  ...(prev[table_name] as FormDataObject[]).filter(
                    (data) => data[dividerFieldName] != option.value,
                  ),
                  ...normalized,
                ],
              }));
            } else {
              setFormData((prev) => ({
                ...prev,
                [table_name]: normalized,
              }));
            }
          }}
          className="space-y-4"
        >
          <AnimatePresence initial={false}>
            {dataList.map((item, index) => {
              const localFields = fields
                .filter((field) => !DividerFieldNames.includes(field.name))
                .filter((field2) => !SequenceFieldNames.includes(field2.name));
              const id = getItemId(item);
              return (
                <DraggableItem
                  key={id}
                  id={id}
                  dragEnabled={!!sequenceFieldName}
                  onDelete={() =>
                    handleDelete({
                      index,
                      dividerFieldName,
                      option,
                      sequenceFieldName,
                    })
                  }
                >
                  <div
                    className={`grid ${
                      isMobile ? "grid-cols-1" : "grid-cols-2"
                    } gap-4`}
                  >
                    <FieldLayout
                      fields={localFields}
                      excludeOptions={Object.fromEntries(
                        Object.entries(excludeOptions).map(
                          ([fieldName, values]) => [
                            fieldName,
                            values.filter((v) => v !== item[fieldName]), // ⛔ jangan exclude milik sendiri
                          ],
                        ),
                      )}
                      formData={item}
                      parentFormData={{ ...parentFormData, ...formData }}
                      authToken="admin"
                      setFormData={createItemSetter({
                        index,
                        currentItem: item,
                        dividerFieldName,
                        option,
                      })}
                      errorForm={getRowErrors(
                        errorForm,
                        table_name,
                        index,
                        dividerFieldName,
                        option,
                      )}
                    />
                  </div>
                </DraggableItem>
              );
            })}
          </AnimatePresence>
        </Reorder.Group>
        <div
          onClick={() => {
            const payload: Record<string, any> = {};
            if (sequenceFieldName) {
              const lastSeq =
                (dataList[dataList.length - 1]?.[
                  sequenceFieldName
                ] as number) ?? 0;
              payload[sequenceFieldName] = lastSeq + 1;
            }
            if (dividerFieldName && option) {
              payload[dividerFieldName] = option.value;
            }
            handleAdd(payload);
          }}
          className={`flex ${
            mode === "view" ? "hidden" : ""
          } text-blue-400 hover:text-blue-500 text-base font-semibold mt-4 cursor-pointer`}
        >
          <LuPlus size={24} /> {title} {option && option.label}
        </div>
        {errorForm[table_name] && (
          <p className="text-red-500 text-sm mt-1">{errorForm[table_name]}</p>
        )}
        {errorForm[`${table_name}.${dividerFieldName}:${option?.value}`] && (
          <p className="text-red-500 text-sm mt-1">
            {errorForm[`${table_name}.${dividerFieldName}:${option?.value}`]}
          </p>
        )}
      </>
    );
  };

  return (
    <div className="space-y-4">
      <h2 className="text-lg font-bold mb-4">{title}</h2>
      {IsStaticOptions(dividerField?.options)
        ? dividerField?.options.map((option) => {
            return (
              <Fragment key={option.value}>
                <div className="text-base font-semibold">{option.label}</div>
                {renderDataList({
                  dataList:
                    ((formData[table_name] as FormDataObject[]) ?? []).filter(
                      (data) => data[dividerField.name] === option.value,
                    ) ?? [],
                  option: option,
                  dividerFieldName: dividerField.name,
                  sequenceFieldName: sequenceField?.name,
                })}
              </Fragment>
            );
          })
        : renderDataList({
            dataList: (formData[table_name] as FormDataObject[]) ?? [],
          })}
    </div>
  );
};
export default DetailForm;
