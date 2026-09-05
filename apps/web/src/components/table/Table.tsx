/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import {
  Fragment,
  useState,
  useEffect,
  useRef,
  useCallback,
  useMemo,
} from "react";
import { useRouter } from "next/navigation";
import TableToolbar from "./TableToolbar";
import Modal from "../Modal";
import Button from "../form/Button";
import { useToast } from "../ToastContext";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import FilterFormTable, { FilterField } from "./FilterFormTable";
import { buildQueryString } from "@/utils/queryString";
import { FaListCheck } from "react-icons/fa6";
import { IoReload } from "react-icons/io5";
import LoadingRow from "./LoadingRow";
import { useQueryParams } from "@/hooks/useQueryParams";
import { HiArrowLongUp } from "react-icons/hi2";
import { useThrottleLeadingTrailing } from "@/hooks/useThrottleLeadingTrailing";
import { useDebounceFn } from "@/hooks/useDebounceFn";
import SelectField from "../form/SelectField";
import { Mode, ModifiedDataTables } from "@/hooks/useDetails";
import {
  resolveDefaultFilterValues,
  resolveFileUrl,
} from "@/utils/globalUtils";
import { getLastDeliveryOrderStatusTypeIdStatusTime } from "@/utils/deliveryOrder";
import { DeliveryOrderStatusTypeMap } from "@/consta/DeliveryOrderStatusTypeMap";
import { formatDateTime } from "@/utils/dateTime";
import { InfoRow } from "../InfoRow";
import { TextSkeleton } from "../TextSkeleton";
import { formatCurrencyIDR } from "@/utils/currencyFormater";
import ImageViewer from "../ImageViewer";
import { BankMerkMap } from "@/consta/BankMerkMap";
import { useAuth } from "@/context/AuthContext";
import { isApproveableData } from "@/utils/approvement";
import { ApprovalBaseTableNameType } from "../approval/ApprovalTable";

export type ColumnField = {
  key: string;
  label: string;
  render?: (
    item: any,
    isSelected?: boolean,
    setPhotoViewData?: (
      data: {
        table_image_id: number;
        image_url: string;
        approval_data: Record<string, any>[];
      }[],
    ) => void,
    setPhotoViewIndex?: (index: number | null) => void,
    setapprovalBaseTableName?: (table_name: ApprovalBaseTableNameType) => void,
    setItemInImageViewer?: (item: any) => void,
    setDestinationIndexInImageViewer?: (index: number | null) => void,
  ) => React.ReactNode;
  columnLength?: number;
  sortable?: "string" | "number" | "date" | "boolean" | "table_key";
  // filterKey?: string;
  // filterOptions?:
  //   | { value: string; label: string }[]
  //   | { url: string; labelKey: string; valueKey: string };
  // autoSelectFirst?: boolean;
};

export interface DetailActionTable {
  initialAdder?: Record<string, number | string>;
  modifiedDataTables: ModifiedDataTables;
  setMode: (mode: Mode) => void;
  setFormData: (data: any) => void;
  remove: (id: number) => void;
  restore: (id: number) => void;
  resetTable: () => void;
}

export interface BaseTableProps {
  url: string;
  table_data?: Record<string, any>[];
  title: string;
  table_name: string;
  table_url?: string;
  default_filter_values?: Record<string, number | string | string[]>;
  filters: FilterField[];
  columns: ColumnField[];
  opendata?: string;
  tableFor?: "detail" | "normal" | "select";
  detailAction?: DetailActionTable;
  onCellClick?: (row: any) => void;
  renderExpandedRow?: (row: any) => React.ReactNode;
  addedToolbarButtons?: ("travel_allowance" | "open_data")[];
  isCreateable?: boolean;
  readOnly?: boolean;
}

export interface TableProps extends BaseTableProps {
  assignTable?: BaseTableProps & {
    assignUrl: string;
    assignIcon: React.ReactNode;
  };
}

export default function Table({
  url,
  table_data,
  title,
  table_name,
  table_url,
  default_filter_values,
  filters,
  columns,
  opendata,
  tableFor = "normal",
  detailAction,
  assignTable,
  onCellClick,
  renderExpandedRow,
  addedToolbarButtons,
  isCreateable = true,
  readOnly = false,
}: TableProps) {
  const key_table = `${table_name}_id`;
  const table_web_url = table_url ? table_url : table_name;
  const isForDetail = tableFor === "detail";
  const isForSelect = tableFor === "select";
  // Hitung di luar — jadi bisa digunakan di mana saja
  const hasIsActiveFilter = useMemo(
    () => filters.some((f) => f.name === "is_active"),
    [filters],
  );

  const hasUserStatusIDFilter = useMemo(
    () => filters.some((f) => f.name === "app_user_status_id"),
    [filters],
  );
  const defaultFilterValues = useMemo((): Record<
    string,
    number | string | string[]
  > => {
    if (hasIsActiveFilter) {
      return { is_active: "1", ...default_filter_values };
    } else if (hasUserStatusIDFilter) {
      return { app_user_status_id: "1", ...default_filter_values };
    }
    return default_filter_values ?? {};
  }, [hasIsActiveFilter, hasUserStatusIDFilter]);

  const [filterValues, setFilterValues] =
    useState<Record<string, number | string | string[]>>(defaultFilterValues);
  const { getParam, searchParams, setParams } = useQueryParams();
  const orderByParam = getParam("order_by");
  const sortParam = getParam("sort");
  const globalLimitParam = getParam("global_limit");
  const isValidSort = (v: string): v is "asc" | "desc" =>
    v === "asc" || v === "desc";
  const defaultSort = useMemo(() => {
    if (orderByParam && sortParam && isValidSort(sortParam)) {
      return { order_by: orderByParam, sort: sortParam };
    }
    return {};
  }, []);
  const [sortValues, setSortValues] = useState<{
    order_by?: string;
    sort?: "asc" | "desc" | "";
  }>(defaultSort);

  const defaultLimit = useMemo((): number => {
    const limitMap: Record<string, number> = {
      "10": 10,
      "25": 25,
      "50": 50,
      "100": 100,
    };
    return (
      limitMap[globalLimitParam ?? ""] ??
      (isForDetail || isForSelect ? 999 : 10)
    );
  }, [globalLimitParam]);
  const [selectedRows, setSelectedRows] = useState<any[]>([]);
  const [expandedRowKey, setExpandedRowKey] = useState<
    string | number | null
  >(null);
  const [multipleMode, setMultipleMode] = useState(false);

  const [dataItems, setDataItems] = useState<any[]>([]);
  const [dataCombineds, setDataCombineds] = useState<any[]>([]);

  const [photoViewIndex, setPhotoViewIndex] = useState<number | null>(null);

  const [itemInImageViewer, setItemInImageViewer] = useState<any | null>();
  const [destinationIndexInImageViewer, setDestinationIndexInImageViewer] =
    useState<any | null>();

  const [photoViewData, setPhotoViewData] = useState<
    {
      table_image_id: number;
      image_url: string;
      approval_data: Record<string, any>[];
    }[]
  >([]);
  const [approvalBaseTableName, setapprovalBaseTableName] =
    useState<ApprovalBaseTableNameType>();

  const [modalDeleteOpen, setModalDeleteOpen] = useState(false);
  const [modalApproveTAOpen, setModalApproveTAOpen] = useState<
    "approveTA" | "rejectTA" | null
  >(null);
  const [modalActivateOpen, setModalActivateOpen] = useState(false);
  const [modalAssignOpen, setModalAssignOpen] = useState(false);

  const [isLoadingAction, setIsLoadingAction] = useState(false);
  const [assignData, setAssignData] = useState<Record<string, string>>();
  const [activateData, setActivateData] = useState(false);
  const [loadingAction, setLoadingAction] = useState(false);

  // const [errorTA, setErrorTA] = useState<string | null>(null);
  // const [driverObjectTA, setDriverObjectTA] = useState<Record<string, any>>({});

  useEffect(() => {
    if (detailAction?.modifiedDataTables?.added_datas)
      setDataCombineds([
        ...[...detailAction?.modifiedDataTables?.added_datas].reverse(),
        ...dataItems,
      ]);
    else {
      setDataCombineds(dataItems);
    }
  }, [detailAction?.modifiedDataTables?.added_datas, dataItems]);

  // const flagUpdated = () => {
  //   if (!detailAction?.modifiedDataTables?.updated_datas) return;
  //   // copy existing list
  //   setDataCombineds((prev) => {
  //     let updatedList = [...prev];

  //     detailAction.modifiedDataTables.updated_datas.forEach((updatedItem) => {
  //       const updatedId = updatedItem[key_table];
  //       const index = updatedList.findIndex(
  //         (item) => item[key_table] === updatedId,
  //       );

  //       if (index !== -1) {
  //         updatedList[index] = updatedItem;
  //       }
  //     });
  //     return updatedList;
  //   });
  // };

  // const flagDeleted = () => {
  //   if (!detailAction?.modifiedDataTables?.deleted_ids) return;
  //   const deletedIds = detailAction.modifiedDataTables.deleted_ids;
  //   setDataCombineds((prev) => {
  //     return prev.map((item) => {
  //       const itemId = item[key_table];
  //       if (deletedIds.includes(itemId)) {
  //         return {
  //           ...item,
  //           flagDel: "deleted",
  //         };
  //       }
  //       return item;
  //     });
  //   });
  // };

  useEffect(() => {
    detailAction?.setMode("add");
    detailAction?.setFormData(detailAction.initialAdder);
    setSelectedRows((prev) => {
      if (prev.length === 0) return prev; // ⛔ tidak rerender
      if (photoViewIndex === null) return [];
      return prev;
    });

    // console.log("KJDFJKSAD")
    // flagDeleted();
    // flagUpdated();
    const deletedIds = detailAction?.modifiedDataTables?.deleted_ids ?? [];
    const updatedDatas = detailAction?.modifiedDataTables?.updated_datas ?? [];

    setDataCombineds((prev) => {
      let next = [...prev];

      // 1️⃣ FLAG DELETED
      next = next.map((item) =>
        deletedIds.includes(item[key_table])
          ? { ...item, flagDel: "deleted" }
          : item,
      );

      // 2️⃣ APPLY UPDATED
      updatedDatas.forEach((updatedItem: any) => {
        const updatedId = updatedItem[key_table];
        const index = next.findIndex((item) => item[key_table] === updatedId);

        if (index !== -1) {
          next[index] = {
            ...next[index], // ⬅️ preserve flagDel
            ...updatedItem,
          };
        }
      });

      return next;
    });
  }, [
    dataItems,
    detailAction?.modifiedDataTables?.added_datas,
    detailAction?.modifiedDataTables?.updated_datas,
    detailAction?.modifiedDataTables?.deleted_ids,
  ]);

  const prevUpdatedRef = useRef<any[]>([]);

  useEffect(() => {
    const prevUpdated = prevUpdatedRef.current ?? [];
    const currentUpdated =
      detailAction?.modifiedDataTables?.updated_datas ?? [];

    if (currentUpdated.length < prevUpdated.length) {
      const prevIds = prevUpdated.map((item) => item[key_table]);
      const currentIds = currentUpdated.map((item) => item[key_table]);
      const removedIds = prevIds.filter((id) => !currentIds.includes(id));
      if (removedIds.length > 0) {
        setDataCombineds((prevCombined) => {
          return prevCombined.map((item) => {
            const itemId = item[key_table];
            if (removedIds.includes(itemId)) {
              const original = dataItems.find((d) => d[key_table] === itemId);
              return original ?? item;
            }
            return item;
          });
        });
      }
    }

    // simpan current sebagai previous
    prevUpdatedRef.current = currentUpdated;
  }, [detailAction?.modifiedDataTables?.updated_datas]);

  const prevDeletedRef = useRef<(number | string)[]>([]);

  useEffect(() => {
    const prevDeleted = prevDeletedRef.current ?? [];
    const currentDeleted = detailAction?.modifiedDataTables?.deleted_ids ?? [];

    // jika jumlah deleted_ids berkurang → ada restore
    if (currentDeleted.length < prevDeleted.length) {
      const restoredIds = prevDeleted.filter(
        (id) => !currentDeleted.includes(id),
      );

      if (restoredIds.length > 0) {
        setDataCombineds((prevCombined) =>
          prevCombined.map((item) => {
            const itemId = item[key_table];

            // hanya hapus flagDel "deleted"
            if (restoredIds.includes(itemId)) {
              const newItem = { ...item };
              delete newItem.flagDel; // ⬅️ hanya ini
              return newItem;
            }

            return item;
          }),
        );
      }
    }

    // simpan current sebagai previous
    prevDeletedRef.current = currentDeleted;
  }, [detailAction?.modifiedDataTables?.deleted_ids]);

  const scrollRef = useRef<HTMLDivElement>(null);

  const [globalLimit, setGlobalLimit] = useState<number>(defaultLimit);
  const [cursorKey, setCursorKey] = useState<number | null>(null); // id terakhir
  const [cursorValue, setCursorValue] = useState<any | null>(null); // value sort terakhir
  const [hasMore, setHasMore] = useState(true); // apakah masih bisa load
  const [isLoading, setIsLoading] = useState(url ? true : false);
  const [isFetchingMore, setIsFetchingMore] = useState(url ? true : false);
  const [error, setError] = useState<string | null>(null);

  const router = useRouter();
  const { showToast, setWrapperClassname } = useToast();

  const { getAPI, deleteAPI, patchAPI, postAPI } = useFetchAPI();

  const abortControllerRef = useRef<AbortController | null>(null);

  const nullCursor = () => {
    setCursorKey(null);
    setCursorValue(null);
  };

  // 🔹 Ambil data dari API
  const fetchData = useCallback(
    async ({
      merge = false,
    }: {
      merge?: boolean;
      callFrom?: string;
    }) => {
      if (!url) {
        setIsLoading(false);
        setIsFetchingMore(false);
        if (table_data) {
          setDataItems(table_data);
        }
        return;
      }
      if (abortControllerRef.current) {
        abortControllerRef.current.abort();
      }
      const controller = new AbortController();
      abortControllerRef.current = controller;
      setIsLoading(!merge);
      setIsFetchingMore(merge);
      setError(null);

      try {
        const queryString =
          Object.keys(filterValues).length > 0
            ? buildQueryString(filterValues)
            : "";
        const params = new URLSearchParams();
        if (queryString) params.append("query", queryString);
        if (merge && cursorKey) params.append("cursor_key", String(cursorKey));
        if (merge && cursorValue)
          params.append("cursor_value", String(cursorValue));
        if (merge && cursorValue)
          params.append("cursor_type", typeof cursorValue);
        if (globalLimit) params.append("limit", String(globalLimit));
        if (sortValues.order_by) {
          params.append("order_by", sortValues.order_by);
        } else {
          params.append("order_by", "created_at");
        }
        if (sortValues.sort) {
          params.append("sort", sortValues.sort);
        } else {
          params.append("sort", "desc");
        }
        const res = await getAPI<any>(`${url}?${params.toString()}`, {
          authToken: "admin",
          signal: controller.signal,
        });
        if (res.code === 200 && res.data) {
          const fetchItems = res.data.items;
          setDataItems((prev) =>
            merge ? [...prev, ...fetchItems] : fetchItems,
          );
          if (fetchItems.length > 0) {
            const lastItem = fetchItems[fetchItems.length - 1];
            setCursorKey(lastItem[`${table_name}_id`] || cursorKey);
            if (sortValues.order_by) {
              setCursorValue(lastItem[sortValues.order_by]);
            } else {
              setCursorValue(lastItem.created_at);
            }
          } else {
            nullCursor();
          }
          setHasMore(fetchItems.length >= globalLimit);
        } else {
          setDataItems([]);
          nullCursor();
          setHasMore(false);
        }
      } catch (err: any) {
        if (err.name === "AbortError") return;
        console.error("Error fetching table data:", err);
        setError("Error fetching table data");
      } finally {
        setIsLoading(false);
        setIsFetchingMore(false);
      }
    },
    [
      url,
      filterValues,
      globalLimit,
      sortValues,
      cursorKey,
      cursorValue,
      table_name,
      table_data,
    ],
  );

  const throttledTrailingFetchData = useThrottleLeadingTrailing(
    (args: { merge?: boolean; callFrom?: string }) => {
      fetchData({ merge: args.merge, callFrom: args.callFrom });
    },
    500,
  );

  const debouncedFetchData = useDebounceFn(
    (args: { merge?: boolean; callFrom?: string }) => {
      fetchData({ merge: args.merge, callFrom: args.callFrom });
    },
    200,
  );

  const didMountFilter = useRef(filters.length === 0 ? true : false);
  const didMountSort = useRef(false);

  useEffect(() => {
    if (!didMountFilter.current) {
      didMountFilter.current = true;
      return;
    }
    if (scrollRef.current) scrollRef.current.scrollTop = 0;
    nullCursor();
    throttledTrailingFetchData({
      merge: false,
      callFrom: "filter",
    });
  }, [filterValues, globalLimit]);

  useEffect(() => {
    if (!didMountSort.current) {
      didMountSort.current = true;
      return;
    }
    if (scrollRef.current) scrollRef.current.scrollTop = 0;
    nullCursor();
    debouncedFetchData({ merge: false, callFrom: "sort" });
  }, [sortValues]);

  const handleScroll = useCallback(async () => {
    if (!scrollRef.current || isFetchingMore || !hasMore) return;

    const { scrollTop, scrollHeight, clientHeight } = scrollRef.current;
    const isNearBottom = scrollTop + clientHeight >= scrollHeight - 100;

    if (isNearBottom) {
      await fetchData({ merge: true, callFrom: "near bottom (cursor)" });
    }
  }, [cursorKey, hasMore, isFetchingMore]);

  // 🧭 Pasang listener scroll
  useEffect(() => {
    const el = scrollRef.current;
    if (!el) return;
    el.addEventListener("scroll", handleScroll);
    return () => el.removeEventListener("scroll", handleScroll);
  }, [handleScroll]);

  // 🔹 Pilih row
  const handleSelectRow = (item: any) => {
    if (isForSelect && onCellClick) {
      onCellClick(item);
    }
    if (multipleMode) {
      const isSelected = selectedRows.some(
        (row) => row[key_table] === item[key_table],
      );
      setSelectedRows(
        isSelected
          ? selectedRows.filter((row) => row[key_table] !== item[key_table])
          : [...selectedRows, item],
      );
    } else {
      const isSelected = selectedRows.some(
        (row) => row[key_table] === item[key_table],
      );
      setSelectedRows(isSelected ? [] : [item]);
    }
  };

  useEffect(() => {
    if (!multipleMode) {
      if (selectedRows.length > 0)
        setSelectedRows([selectedRows[selectedRows.length - 1]]);
    }
  }, [multipleMode]);

  useEffect(() => {
    setActivateData(selectedRows.some((sr) => sr?.is_active == "0"));
    if (selectedRows.length > 0) {
      detailAction?.setFormData(selectedRows[selectedRows.length - 1]);
      detailAction?.setMode("edit");
    } else {
      detailAction?.setFormData(detailAction.initialAdder);
      detailAction?.setMode("add");
    }
  }, [selectedRows]);

  // 🔹 Hapus row
  const handleDelete = async () => {
    try {
      setLoadingAction(true);
      if (selectedRows.length === 0) return;
      let res;
      if (selectedRows.length > 1) {
        const delete_ids = selectedRows.map((sr) => sr[key_table]);
        res = await deleteAPI(
          `${url}/bulk`,
          { authToken: "admin" },
          { delete_ids },
        );
      } else {
        const id = selectedRows[0][key_table];
        res = await deleteAPI(`${url}/${id}`, { authToken: "admin" });
      }

      if (res.code !== 200 && res.code !== 204) {
        throw new Error(res.message || "Failed to delete");
      }

      setSelectedRows([]);
      showToast(3000, "success", "Item berhasil dihapus!");
    } catch (err) {
      console.error(err);
      showToast(3000, "error", "Fitur belum ada, tanyakan ke developer");
    } finally {
      setLoadingAction(false);
      setModalDeleteOpen(false);
      await fetchData({ merge: false, callFrom: "handleDelete" }); // refresh table
    }
  };

  const handleActivate = async () => {
    try {
      setLoadingAction(true);
      if (selectedRows.length === 0) return;

      const activate_ids = selectedRows.map((sr) => sr[key_table]);
      const res = await patchAPI(
        `${url}/is_active`,
        { authToken: "admin" },
        { activate_ids, is_active: activateData ? "1" : "0" },
      );
      if (res.code !== 200 && res.code !== 204) {
        throw new Error(res.message || "Failed to delete");
      }

      setSelectedRows([]);
      showToast(
        3000,
        "success",
        `Item berhasil ${activateData ? "diactivate" : "dideactivate"}!`,
      );
      await fetchData({ merge: false, callFrom: "activate" }); // refresh table
    } catch (err) {
      console.error(err);
      showToast(
        3000,
        "error",
        // `Gagal ${activateData ? "activate" : "deactivate"} item!`
        "Fitur belum ada, tanyakan ke developer",
      );
    } finally {
      setLoadingAction(false);
      setModalActivateOpen(false);

      await fetchData({ merge: false, callFrom: "handleActivate" }); // refresh table
    }
  };

  const handleAssign = useCallback(
    async (
      assignUrl: string,
      driverData: Record<string, any> | undefined,
      doData: Record<string, any> | undefined,
    ) => {
      setIsLoadingAction(true);

      if (!driverData) {
        showToast(3000, "error", "Pilih driver terlebih dahulu!");
        return;
      }

      // const missing: string[] = [];
      // if (!driverData.bank_merk_id) {
      //   missing.push("Bank Merk");
      // }
      // if (!driverData.bank_account_number) {
      //   missing.push("Bank Account Number");
      // }
      // if (!driverData.bank_account_name) {
      //   missing.push("Bank Account Name");
      // }

      // if (missing.length > 0) {
      //   showToast(
      //     3000,
      //     "error",
      //     `Gagal Assign Driver! Data bank driver tidak lengkap, kurang: ${missing.join(
      //       ", ",
      //     )}`,
      //   );
      //   setIsLoadingAction(false);
      //   return;
      // }

      if (!doData) {
        showToast(3000, "error", "Pilih DO terlebih dahulu!");
        // setIsLoadingAction(false);
        return;
      }

      try {
        const payload = {
          delivery_order_id: doData.delivery_order_id,
          truck_id: driverData.truck_id,
          driver_id: driverData.driver_id,
          truck_type_id: driverData.truck_type_id,
        };

        const res = await postAPI<any>(
          assignUrl,
          { authToken: "admin" },
          payload,
        );
        if (res && res.code) {
          if ([200, 201].includes(res.code)) {
            // if (onSuccess) onSuccess(res.data);
            showToast(3000, "success", `${"Assign driver berhasil!"}`);
            // setDataItems((prev) =>
            //   prev.map((item) =>
            //     item[key_table] === res.data[key_table] ? res.data : item,
            //   ),
            // );
          } else {
            showToast(3000, "error", `Gagal Assign Driver! ${res.message}`);
            // setIsLoadingAction(false);
          }
        } else {
          showToast(3000, "error", `Gagal Assign Driver! ${res.message}`);
          // setIsLoadingAction(false);
        }
      } catch (err) {
        console.error("Submit failed", err);
        showToast(3000, "error", `Gagal Assign Driver!`);
        // setIsLoadingAction(false);
      } finally {
        await fetchData({ merge: false, callFrom: "handleAssign" }); // refresh table
        setIsLoadingAction(false);
      }
    },
    [],
  );

  const { adminPayload } = useAuth();
  const delivery_sequence = useMemo(
    () => Number(title.charAt(title.length - 1)),
    [title],
  );

  const dota = useMemo(() => {
    if (![1, 2].includes(delivery_sequence)) return undefined;

    return selectedRows?.[0]?.delivery_order_travel_allowances?.[
      delivery_sequence - 1
    ];
  }, [selectedRows, delivery_sequence]);

  const { isApproveable, isApproveable1 } = useMemo(() => {
    if (![1, 2].includes(delivery_sequence) || !dota) {
      return {
        isApproveable: false,
        isApproveable1: false,
      };
    }

    return isApproveableData({
      approvalData: dota.delivery_order_travel_allowance_approvals,
      adminPayload,
      firstStepAllowed: [3, 4, 5, 6],
      secondStepAllowed: [3, 4],
    });
  }, [dota, adminPayload, delivery_sequence]);

  const handleApproveTA = useCallback(
    async ({
      approveTAUrl,
      dota,
      selectedRows,
      isApproveable1,
    }: {
      approveTAUrl: string;
      dota: any;
      selectedRows: any[];
      isApproveable1: boolean;
    }) => {
      setIsLoadingAction(true);

      try {
        const payload = {
          driver_id: selectedRows[0].driver_id,
          delivery_order_id: selectedRows[0].delivery_order_id,
          delivery_order_travel_allowance_id:
            dota.delivery_order_travel_allowance_id,
          approval_step_id: isApproveable1 ? 1 : 2,
          approval_status_id: isApproveable1 ? 1 : 4,
          delivery_sequence: dota.delivery_sequence,
          approval_notes: null,
        };
        const res = await postAPI<any>(
          approveTAUrl,
          { authToken: "admin" },
          payload,
        );
        if (res && res.code) {
          if ([200, 201].includes(res.code)) {
            // if (onSuccess) onSuccess(res.data);
            showToast(
              3000,
              "success",
              `${"Approve Travel Allowance berhasil!"}`,
            );
          } else {
            showToast(
              3000,
              "error",
              `Gagal Approve Travel Allowance! ${res.message}`,
            );
            // setIsLoadingAction(false);
          }
        } else {
          showToast(
            3000,
            "error",
            `Gagal Approve Travel Allowance! ${res.message}`,
          );
          // setIsLoadingAction(false);
        }
      } catch (err) {
        console.error("Submit failed", err);
        showToast(3000, "error", `Gagal Approve Travel Allowance!`);
        // setIsLoadingAction(false);
      } finally {
        await fetchData({ merge: false, callFrom: "handleApproveTA" }); // refresh table
        setModalApproveTAOpen(null);
        setIsLoadingAction(false);
      }
    },
    [],
  );

  const handleApprovePhoto = useCallback(
    async ({
      approvalBaseTableName,
      payload,
    }: {
      approvalBaseTableName: string;
      payload: Record<string, any>;
    }) => {
      if (!approvalBaseTableName) {
        showToast(3000, "error", "Approval Url tidak ada!");
        return;
      }

      setIsLoadingAction(true);
      try {
        // const payload = {

        //   approval_step_id: 1,
        //   approval_status_id ,
        //   approval_notes: "catatan",
        // };

        const res = await postAPI<any>(
          `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/${approvalBaseTableName}_approval`,
          { authToken: "admin" },
          payload,
        );
        if (res && res.code) {
          if ([200, 201].includes(res.code)) {
            // if (onSuccess) onSuccess(res.data);
            const approval = res.data;

            const baseIdKey = `${approvalBaseTableName}_id`;
            const approvalsKey = `${approvalBaseTableName}_approvals`;

            setDataItems((prev) =>
              prev.map((delivery_order: any) => ({
                ...delivery_order,
                delivery_order_destinations:
                  delivery_order.delivery_order_destinations.map(
                    (dod: any) => ({
                      ...dod,
                      delivery_order_statuses: dod.delivery_order_statuses?.map(
                        (dos: any) => ({
                          ...dos,

                          // NOTE PHOTOS / CARGO PHOTOS dll
                          [`${approvalBaseTableName}s`]: dos[
                            `${approvalBaseTableName}s`
                          ]?.map((item: any) => {
                            // cocokkan ID
                            if (item[baseIdKey] !== approval[baseIdKey]) {
                              return item;
                            }

                            return {
                              ...item,
                              [approvalsKey]: [
                                ...(item[approvalsKey] ?? []),
                                approval, // ⬅️ append
                              ],
                            };
                          }),
                        }),
                      ),
                    }),
                  ),
              })),
            );

            // setPhotoViewData(
            //   dod.delivery_order_statuses.flatMap(
            //     (dos: any) =>
            //       dos.delivery_order_note_photos?.map((np: any) => ({
            //         table_image_id: np.delivery_order_note_photo_id,
            //         image_url: resolveFileUrl(np.delivery_order_note_photo_url),
            //         approval_data: np.delivery_order_note_photo_approvals ?? [],
            //       })) ?? [],
            //   ),
            // );
            setWrapperClassname(3000, "top-20!");
            showToast(
              3000,
              "success",
              `${[1, 4].includes(payload.approval_status_id) ? "Approve" : "Reject"}  berhasil!`,
            );
          } else {
            setWrapperClassname(3000, "top-20!");
            showToast(3000, "error", `Gagal Approve! ${res.message}`);
            // setIsLoadingAction(false);
          }
        } else {
          setWrapperClassname(3000, "top-20!");

          showToast(3000, "error", `Gagal Approve Photo! ${res.message}`);
          // setIsLoadingAction(false);
        }
      } catch (err) {
        console.error("Submit failed", err);
        setWrapperClassname(3000, "top-20!");
        showToast(3000, "error", `Gagal Approve Photo!`);
        await fetchData({ merge: false, callFrom: "handleApprovePhoto" });
        setPhotoViewIndex(null);
      } finally {
        setIsLoadingAction(false);
        // await fetchData({ merge: false, callFrom: "handleApprovePhoto" }); // refresh table
      }
    },
    [],
  );

  // =======================================
  // START hanya untuk delivery order active
  // =======================================
  useEffect(() => {
    if (!dataCombineds || !selectedRows?.length) {
      setPhotoViewData([]);
      return;
    }

    // 1️⃣ FILTER dulu: hanya row yang selected
    const selectedItems = dataCombineds.filter((item: any) =>
      selectedRows.some((row: any) => row[key_table] === item[key_table]),
    );
    // 2️⃣ FLATTEN photo dari selected items

    const photos =
      selectedItems.flatMap(
        (item: any) =>
          item.delivery_order_destinations
            ?.filter(
              (_dod: any, idx: number) => destinationIndexInImageViewer === idx,
            )
            ?.flatMap(
              (dod: any) =>
                dod.delivery_order_statuses?.flatMap(
                  (dos: any) =>
                    dos[`${approvalBaseTableName}s`]?.map((np: any) => ({
                      table_image_id: np[`${approvalBaseTableName}_id`],
                      image_url: resolveFileUrl(
                        np[`${approvalBaseTableName}_url`],
                      ),
                      approval_data:
                        np[`${approvalBaseTableName}_approvals`] ?? [],
                    })) ?? [],
                ) ?? [],
            ) ?? [],
      ) ?? [];

    // 3️⃣ SET STATE SEKALI
    setPhotoViewData(photos);
  }, [dataCombineds, selectedRows, key_table]);

  // =======================================
  // END hanya untuk delivery order active
  // =======================================

  return (
    <>
      <div
        className={`${isForDetail || isForSelect ? "" : "pt-3 px-3"} w-full`}
      >
        <div
          className={`${
            isForDetail || isForSelect
              ? ""
              : "app-scrollbar rounded-md shadow px-4 pt-4 pb-3 border-t-4 border-[#004f7f] overflow-auto h-[calc(100vh-80px)]"
          } bg-white`}
        >
          {/* Header */}
          {!isForSelect && (
            <div className="flex justify-between items-center">
              <h2
                className={`${
                  isForDetail ? "text-lg" : "text-xl"
                } font-bold mb-2`}
              >
                {title}
              </h2>
              <div className="flex flex-row gap-2">
                {isForDetail ? null : (
                  <SelectField
                    id="global_limit"
                    label=""
                    className="w-20"
                    value={globalLimit}
                    uncloseable
                    options={[
                      { label: "10", value: 10 },
                      { label: "25", value: 25 },
                      { label: "50", value: 50 },
                      { label: "100", value: 100 },
                    ]}
                    onChange={(val) => {
                      setGlobalLimit(val as number);
                      if (!isForDetail) {
                        if (val != 10) {
                          setParams({ global_limit: val as string });
                        } else {
                          setParams({ global_limit: "" });
                        }
                      }
                    }}
                  />
                )}
                {url && (
                  <Button
                    id=""
                    onClick={() => {
                      nullCursor();
                      throttledTrailingFetchData({
                        merge: false,
                        callFrom: "reload",
                      });
                    }}
                    variant="gray-outline"
                  >
                    <div
                      className={`${isLoading ? "animate-spin" : ""} pl-[2px]`}
                    >
                      <IoReload size={20} />
                    </div>
                  </Button>
                )}
                {!isForDetail && !isForSelect && isCreateable && !readOnly && (
                  <Button
                    id=""
                    onClick={() =>
                      router.push(
                        `/admin/data/${table_web_url}/create?${searchParams}`,
                      )
                    }
                  >
                    Create
                  </Button>
                )}
              </div>
            </div>
          )}
          {/* Toolbar */}
          <div className="flex flex-col items-start">
            {/* Filter & Sort */}
            {filters.length !== 0 && (
              <FilterFormTable
                defaultFilterValues={defaultFilterValues}
                filters={filters}
                tableFor={tableFor}
                onApplyFilter={(newFilters) => setFilterValues(newFilters)}
                onResetFilter={() => setFilterValues(defaultFilterValues)}
              />
            )}
            {!isForSelect && !readOnly && (
              <TableToolbar
                selectedCount={selectedRows.length}
                onCopy={() => {
                  isForDetail
                    ? detailAction?.setMode("copy")
                    : router.push(
                        `/admin/data/${table_web_url}/${selectedRows[0][key_table]}/copy?${searchParams}`,
                      );
                }}
                onEdit={() => {
                  isForDetail
                    ? detailAction?.setMode("edit")
                    : router.push(
                        `/admin/data/${table_web_url}/${selectedRows[0][key_table]}/edit?${searchParams}`,
                      );
                }}
                onView={() => {
                  isForDetail
                    ? detailAction?.setMode("view")
                    : router.push(
                        `/admin/data/${table_web_url}/${selectedRows[0][key_table]}/view?${searchParams}`,
                      );
                }}
                {...(hasIsActiveFilter && activateData
                  ? { onActivate: () => setModalActivateOpen(true) }
                  : {})}
                {...(hasIsActiveFilter && !activateData
                  ? { onDeactivate: () => setModalActivateOpen(true) }
                  : {})}
                {...(selectedRows.some((row) => row.flagDel !== "deleted")
                  ? {
                      onDelete: () => {
                        if (!isForDetail) {
                          setModalDeleteOpen(true);
                          return;
                        }

                        const deleteIds = selectedRows.map(
                          (sr) => sr[key_table],
                        );
                        deleteIds.forEach((id) => detailAction?.remove(id));
                      },
                    }
                  : {})}
                {...(selectedRows.every((row) => row.flagDel === "deleted")
                  ? {
                      onRestore: () => {
                        if (!isForDetail) return;

                        const restoreIds = selectedRows.map(
                          (sr) => sr[key_table],
                        );
                        restoreIds.forEach((id) => detailAction?.restore(id));
                      },
                    }
                  : {})}
                {...(opendata
                  ? {
                      onOpen: () =>
                        window.open(selectedRows[0][opendata], "_blank"),
                      isOnOpenDisabled:
                        selectedRows.length > 0 && !selectedRows[0][opendata],
                    }
                  : {})}
                {...(addedToolbarButtons?.includes("travel_allowance") &&
                isApproveable
                  ? {
                      onApproveTA: () => setModalApproveTAOpen("approveTA"),
                    }
                  : {})}
                {...(assignTable
                  ? {
                      onAssign: () => {
                        setModalAssignOpen(true);
                      },
                      assignIcon: assignTable.assignIcon,
                    }
                  : {})}
              />
            )}
          </div>

          {/* Table */}
          <div
            className={`overflow-x-auto overflow-y-auto ${
              isForSelect ? "max-h-100" : "max-h-80"
            } app-scrollbar`}
            ref={scrollRef}
          >
            <table className="w-full border-collapse table-fixed">
              <thead className="sticky top-0 z-10">
                <tr>
                  <th className="bg-black h-[1px] w-12"></th>
                  {columns.map((col) => (
                    <th
                      className="bg-black h-[1px] border-black"
                      key={col.key}
                      style={{ width: col.columnLength }}
                    ></th>
                  ))}
                </tr>
                <tr>
                  <th
                    className="p-2 border-x w-12 text-center bg-gray-200 cursor-pointer"
                    onClick={() => setMultipleMode(!multipleMode)}
                  >
                    {multipleMode ? (
                      <FaListCheck
                        size={22}
                        className="mx-auto text-green-500"
                      />
                    ) : (
                      "No"
                    )}
                  </th>
                  {columns.map((col) => (
                    <th
                      key={col.key}
                      className={`p-2 border-x bg-gray-200 ${
                        col.sortable ? "cursor-pointer" : ""
                      }`}
                      style={{ width: col.columnLength }}
                      onClick={() => {
                        if (!col.sortable) return;
                        // urutan toggle: "" → asc → desc → ""
                        let sortData: any = {};
                        if (sortValues.order_by !== col.key) {
                          sortData = { order_by: col.key, sort: "asc" };
                        } else if (sortValues.sort === "asc") {
                          sortData = { order_by: col.key, sort: "desc" };
                        }
                        setSortValues(sortData);
                        if (sortData.order_by && !isForDetail) {
                          setParams(sortData);
                        } else {
                          setParams({ order_by: "", sort: "" });
                        }
                      }}
                    >
                      <div className="flex items-center justify-between w-full">
                        {/* Label di tengah */}
                        <span className="flex-1 text-center ml-4">
                          {col.label}
                        </span>

                        <>
                          <div
                            className={`transition duration-300 flex w-4 ${
                              sortValues.order_by !== col.key
                                ? col.sortable
                                  ? "opacity-10"
                                  : "opacity-0"
                                : ""
                            }`}
                          >
                            <div className="w-[10px] h-4 flex items-start justify-start">
                              <HiArrowLongUp
                                viewBox="4 4 16 16"
                                strokeWidth={1.5}
                                className={`transition duration-300 ${
                                  sortValues.order_by == col.key &&
                                  sortValues.sort === "desc"
                                    ? ""
                                    : "rotate-180"
                                }`}
                              />
                            </div>
                            <div className="text-[7px] flex flex-col gap-[0] pt-[2px] w-[6px] select-none">
                              {col.sortable === "string" && (
                                <>
                                  <div className="leading-[7px]">A</div>
                                  <div className="leading-[7px]">Z</div>
                                </>
                              )}
                              {col.sortable === "number" && (
                                <>
                                  <div className="leading-[7px]">0</div>
                                  <div className="leading-[7px]">9</div>
                                </>
                              )}
                              {col.sortable &&
                                ["date", "boolean", "table_key"].includes(
                                  col.sortable,
                                ) && (
                                  <div className="flex flex-col justify-between h-3">
                                    <div className="h-[1px] bg-black scale-x-50 origin-left"></div>
                                    <div className="h-[1px] bg-black scale-x-75 origin-left"></div>
                                    <div className="h-[1px] bg-black scale-x-100 origin-left"></div>
                                    <div className="h-[1px] bg-black scale-x-[1.25] origin-left"></div>
                                  </div>
                                )}
                            </div>
                          </div>
                        </>
                      </div>
                    </th>
                  ))}
                </tr>
                <tr>
                  <th className="bg-black h-[1px] border-x w-12 "></th>
                  {columns.map((col) => (
                    <th
                      className="bg-black h-[1px] border-r"
                      key={col.key}
                      style={{ width: col.columnLength }}
                    ></th>
                  ))}
                </tr>
              </thead>
              <tbody>
                {isLoading && (
                  <LoadingRow rowCount={5} columnCount={columns.length + 1} />
                )}

                {error && (
                  <tr>
                    <td
                      colSpan={columns.length + 1}
                      className="text-center p-1 text-red-500"
                    >
                      {error}
                    </td>
                  </tr>
                )}
                {/* {detailAction?.modifiedDataTables?.added_datas &&
                  detailAction?.modifiedDataTables?.added_datas.length > 0 &&
                  detailAction?.modifiedDataTables?.added_datas.map((added) => (
                    <div>{added}</div>
                  ))} */}
                {!isLoading &&
                  !error &&
                  (dataCombineds.length > 0 ? (
                    dataCombineds.map((item: any, index: number) => {
                      const isSelected = selectedRows.some(
                        (row) => row[key_table] === item[key_table],
                      );
                      return (
                        <Fragment key={item[key_table]}>
                          <tr
                            title={
                              renderExpandedRow
                                ? "Tampilkan detail"
                                : undefined
                            }
                            className={`cursor-pointer ${
                              item.flagDel == "deleted"
                                ? isSelected
                                  ? "bg-red-200"
                                  : "bg-red-50 hover:bg-red-100"
                                : item.flag == "added"
                                  ? isSelected
                                    ? "bg-blue-200"
                                    : "bg-blue-50 hover:bg-blue-100"
                                  : item.flag == "edited"
                                    ? isSelected
                                      ? "bg-green-200"
                                      : "bg-green-50 hover:bg-green-100"
                                    : isSelected
                                      ? "bg-gray-200"
                                      : "hover:bg-gray-100"
                            }`}
                            onClick={() => {
                              handleSelectRow(item);
                              if (renderExpandedRow) {
                                setExpandedRowKey((current) =>
                                  current === item[key_table]
                                    ? null
                                    : item[key_table],
                                );
                              }
                            }}
                          >
                            <td
                              className={`px-2 py-1 ${
                                index == 0
                                  ? "border-x"
                                  : dataItems.length == index + 1
                                    ? "border-x"
                                    : "border"
                              } text-center align-top`}
                            >
                              {index + 1}
                            </td>
                            {columns.map((col) => (
                              <td
                                key={col.key}
                                className={`px-2 py-1 ${
                                  index == 0
                                    ? "border-x"
                                    : dataItems.length == index + 1
                                      ? "border-x"
                                      : "border"
                                } align-top`}
                              >
                                {col.render ? (
                                  <div
                                    className={`${
                                      isSelected
                                        ? ""
                                        : "line-clamp-1 break-all"
                                    }`}
                                  >
                                    {col.render(
                                      item,
                                      isSelected,
                                      setPhotoViewData,
                                      setPhotoViewIndex,
                                      setapprovalBaseTableName,
                                      setItemInImageViewer,
                                      setDestinationIndexInImageViewer,
                                    )}
                                  </div>
                                ) : (
                                  <div
                                    className={`${
                                      isSelected
                                        ? ""
                                        : "line-clamp-1 break-all"
                                    }`}
                                  >
                                    {item[col.key]}
                                  </div>
                                )}
                              </td>
                            ))}
                          </tr>
                          {renderExpandedRow &&
                            expandedRowKey === item[key_table] && (
                              <tr className="bg-slate-50">
                                <td
                                  colSpan={columns.length + 1}
                                  className="border px-3 py-3"
                                >
                                  {renderExpandedRow(item)}
                                </td>
                              </tr>
                            )}
                        </Fragment>
                      );
                    })
                  ) : (
                    <tr>
                      <td
                        colSpan={columns.length + 1}
                        className="text-center p-1"
                      >
                        data yang dicari tidak ditemukan
                      </td>
                    </tr>
                  ))}

                {isFetchingMore && (
                  <LoadingRow rowCount={5} columnCount={columns.length + 1} />
                )}
              </tbody>
            </table>
          </div>
        </div>
      </div>

      {photoViewIndex != null && (
        <ImageViewer
          images={photoViewData.map((data) => data.image_url)}
          activeIndex={photoViewIndex}
          onClose={() => setPhotoViewIndex(null)}
          itemInImageViewer={itemInImageViewer}
          imageIDs={photoViewData.map((data) => data.table_image_id)}
          approvalBaseTableName={approvalBaseTableName}
          onApprove={
            approvalBaseTableName
              ? async (payload: Record<string, any>) => {
                  handleApprovePhoto({
                    approvalBaseTableName,
                    payload: {
                      ...payload,
                      delivery_order_id: selectedRows[0].delivery_order_id,
                      driver_id: selectedRows[0].driver.app_user_id,
                      [`${approvalBaseTableName}_id`]: payload.table_id,
                    },
                  });
                }
              : undefined
          }
          onReject={
            approvalBaseTableName
              ? async (payload: Record<string, any>) => {
                  handleApprovePhoto({
                    approvalBaseTableName,
                    payload: {
                      ...payload,
                      delivery_order_id: selectedRows[0].delivery_order_id,
                      driver_id: selectedRows[0].driver.app_user_id,
                      [`${approvalBaseTableName}_id`]: payload.table_id,
                    },
                  });
                }
              : undefined
          }
          approvalDataAll={photoViewData.map((data) => data.approval_data)}
        />
      )}

      {/* Modal Delete */}
      <Modal
        isOpen={modalDeleteOpen}
        title="Delete Confirmation"
        confirmText="Delete"
        confirmVariant="red-solid"
        onCancel={() => setModalDeleteOpen(false)}
        onConfirm={handleDelete}
        loading={loadingAction}
      >
        Apakah Anda yakin ingin menghapus item ini? Tindakan ini tidak dapat
        dibatalkan.
      </Modal>

      <Modal
        isOpen={modalActivateOpen}
        title={
          activateData ? "Activate Confirmation" : "Deactivate Confirmation"
        }
        confirmText={activateData ? "Activate" : "Deactivate"}
        confirmVariant="yellow-solid"
        onCancel={() => setModalActivateOpen(false)}
        onConfirm={() => handleActivate()}
        loading={loadingAction}
      >
        Are you sure you want to {activateData ? "activate" : "deactivate"} this
        item?
      </Modal>

      {addedToolbarButtons?.includes("travel_allowance") &&
        (() => {
          if (selectedRows.length !== 1) {
            return;
          }

          if (!dota) {
            return;
          }
          const {
            statusTypeId,
            statusTime,
            destinationName,
            districtName,
            cityName,
            provinceName,
            destinationIndex,
          } = getLastDeliveryOrderStatusTypeIdStatusTime(
            selectedRows[0].delivery_order_destinations,
          );
          return (
            <Modal
              size="3xl"
              isOpen={
                modalApproveTAOpen
                  ? ["approveTA"].includes(modalApproveTAOpen)
                  : false
              }
              title={`Approve Uang Jalan`}
              cancelText="Tutup"
              confirmText="Approve UJ"
              confirmVariant="green-solid"
              onCancel={() => {
                setModalApproveTAOpen(null);
              }}
              loading={isLoadingAction}
              onConfirm={() => {
                if (
                  !selectedRows[0]?.driver.bank_account_name ||
                  !selectedRows[0]?.driver.bank_account_number ||
                  !selectedRows[0]?.driver.bank_merk_id
                ) {
                  showToast(
                    3000,
                    "error",
                    `Gagal Approve Travel Allowance, data bank driver belum ada!`,
                  );
                  fetchData({ merge: false, callFrom: "handleApproveTA" });
                  setModalApproveTAOpen(null);
                  return;
                }
                handleApproveTA({
                  approveTAUrl: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/delivery_order_travel_allowance_approval`,
                  dota,
                  selectedRows,
                  isApproveable1,
                });
                setModalApproveTAOpen(null);
                // setSelectedRows([]);
              }}
            >
              <div className="h-[60vh] space-y-2">
                {/* HEADER */}
                <div className="text-lg font-semibold">
                  {selectedRows[0].delivery_order_number}
                </div>

                {/* BASIC INFO */}
                <div className=" bg-white space-y-2">
                  <InfoRow label="Customer">
                    {selectedRows[0].customer?.customer_name}
                  </InfoRow>

                  <InfoRow label="Driver">
                    {selectedRows[0].driver?.app_user_name}
                  </InfoRow>

                  <InfoRow label="Pengiriman ke">
                    <span className="font-medium">
                      {delivery_sequence} (
                      {delivery_sequence == 1 ? "Satu" : "Dua"})
                    </span>
                  </InfoRow>
                  <InfoRow label="Jumlah">
                    <span className="font-semibold text-emerald-600">
                      {formatCurrencyIDR(
                        selectedRows[0][
                          `travel_allowance_amount_${delivery_sequence}`
                        ],
                      )}
                    </span>
                  </InfoRow>
                </div>

                {/* STATUS */}
                <div className=" space-y-2">
                  <InfoRow label="Status Terakhir">
                    <span className="font-semibold">
                      {DeliveryOrderStatusTypeMap[statusTypeId]?.value}
                    </span>
                    <span className="text-gray-500">
                      {" "}
                      at {formatDateTime(statusTime)}
                    </span>
                  </InfoRow>

                  <InfoRow label="Lokasi">
                    <span>
                      {destinationName}, {districtName ?? "-"},{" "}
                      {cityName ?? "-"}
                    </span>
                  </InfoRow>
                </div>

                {/* BANK INFO */}
                <div className="rounded-sm border bg-white p-3 shadow-sm space-y-2 w-2/3 min-w-sm">
                  <div className="text-base font-medium ">Informasi Bank</div>
                  <>
                    <InfoRow label="Bank">
                      {/* {BankMerkMap[selectedRows[0]?.driver.bank_merk_id]
                        ?.name ?? <div className="text-red-500">belum ada</div>} */}
                      {BankMerkMap[dota?.bank_merk_id].name}
                    </InfoRow>

                    <InfoRow label="No. Rekening">
                      {/* {selectedRows[0]?.driver.bank_account_number ?? (<div className="text-red-500">belum ada</div>)} */}
                      {dota?.bank_account_number}
                    </InfoRow>

                    <InfoRow label="Nama Akun">
                      {/* {selectedRows[0]?.bank_account_name ||
                        selectedRows[0]?.driver?.app_user_name} */}
                      {dota?.bank_account_name}
                    </InfoRow>
                  </>
                </div>
              </div>
            </Modal>
          );
        })()}
      {assignTable && (
        <Modal
          size="7xl"
          isOpen={modalAssignOpen}
          title={assignTable.title}
          cancelText="Tutup"
          confirmText="Pilih"
          loading={isLoadingAction}
          onCancel={() => {
            setAssignData(undefined);
            setModalAssignOpen(false);
          }}
          onConfirm={() => {
            // console.log("assignData:", assignData, selectedRows);
            if (assignTable.table_name === "delivery_order") {
              handleAssign(assignTable.assignUrl, selectedRows[0], assignData);
            } else if (assignTable.table_name === "truck") {
              handleAssign(assignTable.assignUrl, assignData, selectedRows[0]);
            }
            setAssignData(undefined);
            setModalAssignOpen(false);
            // setSelectedRows([]);
          }}
        >
          <div className="h-[60vh]">
            <Table
              url={assignTable.url}
              title={assignTable.title}
              table_name={assignTable.table_name}
              columns={assignTable.columns}
              filters={assignTable.filters}
              default_filter_values={
                assignTable.default_filter_values
                  ? resolveDefaultFilterValues(
                      assignTable.default_filter_values,
                      selectedRows[0],
                    )
                  : undefined
              }
              tableFor="select"
              onCellClick={(row) => {
                setAssignData(row);
              }}
            />
          </div>
        </Modal>
      )}
    </>
  );
}
