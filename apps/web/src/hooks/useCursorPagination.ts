/* eslint-disable @typescript-eslint/no-explicit-any */
import { buildQueryString } from "@/utils/queryString";
import { useCallback, useRef, useState } from "react";
import { authTokenType, useFetchAPI } from "./useFetchAPI";

interface FetchParams {
    merge?: boolean;
    callFrom?: string;
}

interface UseCursorPaginationProps {
    url: string;
    tableName: string;
    authToken: authTokenType;
    filterValues?: Record<string, any>;
    sortValues?: {
        order_by?: string;
        sort?: "asc" | "desc";
    };
    limit?: number;
}

export function useCursorPagination<T>({
    url,
    tableName,
    authToken,
    filterValues = {},
    sortValues = {},
    limit = 10,
}: UseCursorPaginationProps) {
    const abortControllerRef = useRef<AbortController | null>(null);

    const [data, setData] = useState<T[]>([]);
    const [isLoading, setIsLoading] = useState(false);
    const [isFetchingMore, setIsFetchingMore] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const [cursorKey, setCursorKey] = useState<any>(null);
    const [cursorValue, setCursorValue] = useState<any>(null);
    const [hasMore, setHasMore] = useState(true);

    const cursorRef = useRef<{
        key: any;
        value: any;
    }>({ key: null, value: null });

    const nullCursor = () => {
        cursorRef.current = { key: null, value: null };
        setCursorKey(null);
        setCursorValue(null);
    };

    const { getAPI } = useFetchAPI();

    const fetchData = useCallback(
        async ({ merge = false }: FetchParams = {}) => {
            // abort previous request
            if (abortControllerRef.current) {
                abortControllerRef.current.abort();
            }

            const controller = new AbortController();
            abortControllerRef.current = controller;

            setIsLoading(!merge);
            setIsFetchingMore(merge);
            setError(null);

            try {
                const params = new URLSearchParams();

                Object.entries(filterValues).forEach(([key, value]) => {
                    if (value !== undefined && value !== null && value !== "") {
                        if (Array.isArray(value)) {
                            value.forEach((v) => params.append(key, String(v)));
                        } else {
                            params.append(key, String(value));
                        }
                    }
                });
                if (Object.keys(filterValues).length > 0) {
                    params.append("query", buildQueryString(filterValues));
                }

                if (merge && cursorRef.current.key) {
                    params.append("cursor_key", String(cursorRef.current.key));
                    params.append("cursor_value", String(cursorRef.current.value));
                    params.append("cursor_type", typeof cursorRef.current.value);
                }

                params.append("limit", String(limit));
                params.append("order_by", sortValues.order_by || "created_at");
                params.append("sort", sortValues.sort || "desc");

                const res = await getAPI<any>(`${url}?${params.toString()}`, {
                    authToken,
                    signal: controller.signal,
                });

                if (res.code === 200 && res.data) {
                    const items: T[] = res.data.items ?? [];

                    setData((prev) => (merge ? [...prev, ...items] : items));

                    if (items.length > 0) {
                        const lastItem: any = items[items.length - 1];

                        const nextKey = lastItem[`${tableName}_id`];
                        const nextValue = sortValues.order_by
                            ? lastItem[sortValues.order_by]
                            : lastItem.created_at;

                        cursorRef.current = {
                            key: nextKey,
                            value: nextValue,
                        };

                        setCursorKey(nextKey);
                        setCursorValue(nextValue);
                    } else {
                        nullCursor();
                    }

                    setHasMore(items.length >= limit);
                } else {
                    setData([]);
                    nullCursor();
                    setHasMore(false);
                }
            } catch (err: any) {
                if (err.name === "AbortError") return;
                console.error(err);
                setError("Error fetching table data");
            } finally {
                setIsLoading(false);
                setIsFetchingMore(false);
            }
        },
        [url, filterValues, limit, sortValues, cursorKey, cursorValue, tableName]
    );

    return {
        data,
        isLoading,
        isFetchingMore,
        error,
        hasMore,
        fetchData,
        reset: () => {
            setData([]);
            nullCursor();
            setHasMore(true);
        },
    };
}
