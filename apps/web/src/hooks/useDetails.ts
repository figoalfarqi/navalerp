// import { detailsAtom, initDetailsAtom } from "@/atoms/Details";
// import { useAtom, useSetAtom } from "jotai";
import { useEffect, useState } from "react";

/* ================================
   TYPES 
================================ */

export type Mode = "edit" | "copy" | "view" | "add";

export interface ModifiedDataTables<T = any> {
    added_datas: T[];
    updated_datas: T[];
    deleted_ids: (string | number)[];
}

export interface DetailState<T = any> {
    mode: Mode;
    initialAdder: Record<string, T>;
    formData: Record<string, T>;
    initialDataTables: T[];
    modifiedDataTables: ModifiedDataTables<T>;
    localIdCounter: number; // ---> counter ID untuk tabel tersebut
}

export type DetailsMap<T = any> = Record<string, DetailState<T>>;


/* ================================
   HOOK
================================ */

export function useDetails<T = any>(
    keys: string[],
    defaultFormData: Record<string,Record<string, T>> = {}
) {
    const [details, setDetails] = useState<Record<string, DetailState<T>>>(() => {
        const initialState: Record<string, DetailState<T>> = {};

        keys.forEach((key) => {
            const defaultData = defaultFormData[key] ?? ({} as Record<string, T>);

            initialState[key] = {
                mode: "edit",
                formData: defaultData,
                initialAdder: defaultData,
                initialDataTables: [],
                modifiedDataTables: {
                    added_datas: [],
                    updated_datas: [],
                    deleted_ids: [],
                },
                localIdCounter: 1,
            };
        });

        return initialState;
    });

    // const initDetails = useSetAtom(initDetailsAtom);

    // useEffect(() => {
    //     initDetails({ keys, defaultFormData });
    // }, [keys]);

    /* =====================================
       Helper: Generate local ID
       ID = 999999 + counter
    ===================================== */
    const generateLocalId = (tableState: DetailState<T>) => {
        const counter = tableState.localIdCounter;

        return 9_999 * 1_000 + counter;
        // contoh 9999000 + counter → sangat kecil kemungkinan bentrok.
    };


    /* ===============================
     SET MODE 
    ================================= */
    const setMode = (key: string, mode: Mode) => {
        setDetails((prev) => {
            if (prev[key]?.mode === mode) return prev;
            return {
                ...prev,
                [key]: { ...prev[key], mode },
            };
        });
    };


    /* ===============================
       SET INITIAL DATA TABLES
    ================================= */
    const setInitialDataTables = (key: string, data: T[]) => {
        setDetails((prev) => ({
            ...prev,
            [key]: { ...prev[key], initialDataTables: data },
        }));
    };

    /* -----------------------------------
      SET FORM DATA
    ----------------------------------- */
    const setFormData = (key: string, data: Record<string, any>) =>
        setDetails((prev) => ({
            ...prev,
            [key]: { ...prev[key], formData: data },
        }));


    /* =====================================
       ADD — data CREATE
       - Generate custom ID -> {table_name}_id
    ===================================== */
    const add = (key: string, data: T, table_name?: string) => {
        setDetails((prev) => {
            const tableState = prev[key];
            const idName = `${table_name ?? key}_id`; // contoh: customer_product_id

            const newId = generateLocalId(tableState);

            const newData = {
                ...data,
                [idName]: newId,
            };

            return {
                ...prev,
                [key]: {
                    ...tableState,
                    modifiedDataTables: {
                        ...tableState.modifiedDataTables,
                        added_datas: [...tableState.modifiedDataTables.added_datas, { ...newData, flag: 'added' }],
                    },
                    localIdCounter: tableState.localIdCounter + 1, // increment counter
                },
            };
        });
    };

    /* =====================================
       UPDATE
       - Data added → update berdasarkan {table}_id
       - Data updated → update berdasarkan {table}_id
    ===================================== */
    const update = (key: string, data: T, table_name?: string) => {
        const key_id = `${table_name ?? key}_id`;
        const key_val = (data as any)[key_id];

        setDetails((prev) => {
            const { added_datas, updated_datas } = prev[key].modifiedDataTables;

            /* --- Update added_datas --- */
            const addedIndex = added_datas.findIndex(
                (item) => (item as any)[key_id] === key_val
            );
            if (addedIndex !== -1) {
                const newAdded = [...added_datas];
                newAdded[addedIndex] = { ...data, flag: 'added' };

                return {
                    ...prev,
                    [key]: {
                        ...prev[key],
                        modifiedDataTables: {
                            ...prev[key].modifiedDataTables,
                            added_datas: newAdded,
                        },
                    },
                };
            }

            /* --- Update updated_datas based on DB {table}_id --- */
            const updatedIndex = updated_datas.findIndex(
                (item) => (item as any)[key_id] === key_val
            );
            if (updatedIndex !== -1) {
                const newUpdated = [...updated_datas];
                newUpdated[updatedIndex] = { ...data, flag: 'edited' };

                return {
                    ...prev,
                    [key]: {
                        ...prev[key],
                        modifiedDataTables: {
                            ...prev[key].modifiedDataTables,
                            updated_datas: newUpdated,
                        },
                    },
                };
            }
            /* --- Not found → push to updated_datas --- */
            return {
                ...prev,
                [key]: {
                    ...prev[key],
                    modifiedDataTables: {
                        ...prev[key].modifiedDataTables,
                        updated_datas: [...updated_datas, { ...data, flag: 'edited' }],
                    },
                },
            };
        });
    };

    /* =====================================
       REMOVE
       - Remove from added_datas using {table}_id
       - Remove from updated_datas using real ID
    ===================================== */
    const remove = (key: string, id: number, table_name?: string) => {
        const key_id = `${table_name ?? key}_id`;

        setDetails((prev) => {
            const mod = prev[key].modifiedDataTables;

            const isAdded = mod.added_datas.some(
                (item) => (item as any)[key_id] === id
            );

            const isUpdated = mod.updated_datas.some(
                (item) => (item as any)[key_id] === id
            );

            // if (isAdded) {
            //     return {
            //         ...prev,
            //         [key]: {
            //             ...prev[key],
            //             modifiedDataTables: {
            //                 ...mod,
            //                 added_datas: mod.added_datas.filter(
            //                     (d) => (d as any)[key_id] !== id
            //                 ),
            //             },
            //         },
            //     };
            // }
            // if (isUpdated) {
            //     return {
            //         ...prev,
            //         [key]: {
            //             ...prev[key],
            //             modifiedDataTables: {
            //                 ...mod,
            //                 updated_datas: mod.updated_datas.filter((d) => (d as any)[key_id] !== id),
            //                 deleted_ids: [...mod.deleted_ids, id],
            //             },
            //         },
            //     };
            // }
            return {
                ...prev,
                [key]: {
                    ...prev[key],
                    modifiedDataTables: {
                        ...mod,
                        deleted_ids: [...mod.deleted_ids, id],
                    },
                },
            };
        });
    };

    /* =====================================
       RESET / RESTORE → tetap sama
    ===================================== */

    const resetItem = (key: string, item: T, table_name?: string) => {
        const key_id = `${table_name ?? key}_id`;
        const key_val = (item as any)[key_id];

        setDetails((prev) => {
            const mod = prev[key].modifiedDataTables;

            return {
                ...prev,
                [key]: {
                    ...prev[key],
                    modifiedDataTables: {
                        ...mod,
                        added_datas: mod.added_datas.filter(
                            (d) => (d as any)[key_id] !== key_val
                        ),
                        updated_datas: mod.updated_datas.filter((d) => (d as any)[key_id] !== key_val),
                    },
                },
            };
        });
    };

    const restore = (key: string, id: number | string) => {
        setDetails((prev) => ({
            ...prev,
            [key]: {
                ...prev[key],
                modifiedDataTables: {
                    ...prev[key].modifiedDataTables,
                    deleted_ids: prev[key].modifiedDataTables.deleted_ids.filter(
                        (del_id) => del_id !== id
                    ),
                },
            },
        }));
    };

    const resetTable = (key: string) => {
        setDetails((prev) => ({
            ...prev,
            [key]: {
                ...prev[key],
                modifiedDataTables: {
                    added_datas: [],
                    updated_datas: [],
                    deleted_ids: [],
                },
                localIdCounter: 1,
            },
        }));
    };

    return {
        details,
        setDetails,
        setMode,
        setInitialDataTables,
        setFormData,
        add,
        update,
        remove,
        resetItem,
        restore,
        resetTable,
    };
}
