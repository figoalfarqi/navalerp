"use client";

import { usePathname, useRouter, useSearchParams } from "next/navigation";

export function useQueryParams() {
    const router = useRouter();
    const pathname = usePathname();
    const searchParams = useSearchParams();

    // 🔹 Ambil 1 param
    const getParam = (key: string): string | undefined => {
        return searchParams.get(key) ?? undefined;
    };

    const getAllParams = (): Record<string, string> => {
        const entries = Array.from(searchParams.entries());
        return Object.fromEntries(entries);
    };

    // 🔹 Set atau hapus 1 param
    const setParam = (key: string, value: string) => {
        const params = new URLSearchParams(searchParams.toString());

        if (value === "") {
            params.delete(key);
        } else {
            params.set(key, value);
        }

        const queryString = params.toString();
        router.replace(`${pathname}${queryString ? `?${queryString}` : ""}`, { scroll: false });
    };

    // 🔹 Set beberapa param sekaligus
    const setParams = (obj: Record<string, string>) => {
        const params = new URLSearchParams(searchParams.toString());

        Object.entries(obj).forEach(([k, v]) => {
            if (v === "") {
                params.delete(k);
            } else {
                params.set(k, v);
            }
        });

        const queryString = params.toString();
        router.replace(`${pathname}${queryString ? `?${queryString}` : ""}`, { scroll: false });
    };

    const clearParams = () => {
        router.replace(pathname, { scroll: false });
    };

    return {
        getParam,
        getAllParams,
        setParam,
        setParams,
        clearParams,
        searchParams,
    };
}
