/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useEffect } from "react";
import { useIndexedDB } from "./useIndexedDB";
import { STORES } from "@/types/tableIDB";
import { useFetchAPI } from "./useFetchAPI";

export default function useOnlineSync() {
    const { dbReady, getIDB } = useIndexedDB();
    const { postAPI } = useFetchAPI();
    useEffect(() => {
        async function syncData() {
            if (dbReady) {
                const drafts = await getIDB(STORES.PENDING_POSTS);
                if (drafts.data.length > 0) {
                    for (const item of drafts.data) {
                        try {
                            await postAPI<any>("/api/sync", { authToken: "none" },
                                JSON.stringify(item))
                        }
                        catch (err) {
                            console.error("Sync failed, will retry later", err);
                        }
                    }
                }
            }
        }

        if (navigator.onLine) syncData();
        window.addEventListener("online", syncData);

        return () => window.removeEventListener("online", syncData);
    }, [dbReady, getIDB, postAPI]);
}


/*
harusnya tidak ada PENDING_POSTS, 
yang benar adalah di setiap tabel dalam IDB memiliki 
"status": "pending_add" | "pending_edit" | "pending_delete" | "synced"
jadi ketika kembali online harus mengecek semua data tabelnya 


kemudian untuk form yang belum di submit oleh user
memiliki IDB juga, yaitu forms_not_submited dengan id adalah nama formnya seperti edit_irregular_cost
ketika create  ada tulisan, anda memiliki draft pembuatan, ada 2 tombol
iya lanjutkan draft, tidak buat baru
begitupun ketika edit
*/
