// /* eslint-disable @typescript-eslint/no-explicit-any */
// "use client";

// import { useEffect } from "react";
// import { useSyncExternalStore } from "react";
// import { useIndexedDB } from "./useIndexedDB";
// import { StoreName, STORES } from "@/utils/tableIDB";

// // volatile memory
// const store = new Map<string, any[]>();
// const listeners = new Map<string, Set<() => void>>();

// function subscribe(key: string, cb: () => void) {
//     if (!listeners.has(key)) listeners.set(key, new Set());
//     listeners.get(key)!.add(cb);
//     return () => listeners.get(key)!.delete(cb);
// }

// function notify(key: string) {
//     listeners.get(key)?.forEach((cb) => cb());
// }

// export function useGlobalTable(key: StoreName) {
//     const { dbReady, getIDB, postIDB, putIDB, patchIDB, deleteIDB, clearIDB } =
//         useIndexedDB();

//     // Ambil data awal dari IDB kalau volatile kosong
//     useEffect(() => {
//         if (!dbReady) return;
//         if (!store.has(key)) {
//             (async () => {
//                 const result = await getIDB(key);
//                 const data = Array.isArray(result) ? result : result ? [result] : [];
//                 store.set(key, data);
//                 notify(key);
//             })();
//         }
//     }, [dbReady, key, getIDB]);

//     // Data yang subscribe
//     const data = useSyncExternalStore(
//         (cb) => subscribe(key, cb),
//         () => store.get(key) ?? []
//     );

//     // CRUD API
//     const actions = {
//         getGlobalTable: (id?: number) => {
//             const all = store.get(key) ?? [];
//             return id ? all.find((d: any) => d.id === id) : all;
//         },
//         postGlobalTable: async (item: any) => {
//             await postIDB(key, item);
//             const all = [...(store.get(key) ?? []), item];
//             store.set(key, all);
//             notify(key);
//         },
//         putGlobalTable: async (item: any) => {
//             await putIDB(key, item);
//             const all = (store.get(key) ?? []).map((d: any) =>
//                 d.id === item.id ? item : d
//             );
//             store.set(key, all);
//             notify(key);
//         },
//         patchGlobalTable: async (id: number, partial: any) => {
//             await patchIDB(key, id, partial);
//             const all = (store.get(key) ?? []).map((d: any) =>
//                 d.id === id ? { ...d, ...partial } : d
//             );
//             store.set(key, all);
//             notify(key);
//         },
//         deleteGlobalTable: async (id: number) => {
//             await deleteIDB(key, id);
//             const all = (store.get(key) ?? []).filter((d: any) => d.id !== id);
//             store.set(key, all);
//             notify(key);
//         },
//         clearGlobalTable: async () => {
//             await clearIDB(key);
//             store.set(key, []);
//             notify(key);
//         },
//     };

//     return [data, actions] as const;
// }



// // function UsersList() {
// //   const [users, { postGlobalTable, deleteGlobalTable }] = useGlobalTable("users");

// //   return (
// //     <div>
// //       <button
// //         onClick={() =>
// //           postGlobalTable({ id: Date.now(), name: "User " + Date.now() })
// //         }
// //       >
// //         Tambah User
// //       </button>
// //       <ul>
// //         {users.map((u: any) => (
// //           <li key={u.id}>
// //             {u.name}{" "}
// //             <button onClick={() => deleteGlobalTable(u.id)}>Hapus</button>
// //           </li>
// //         ))}
// //       </ul>
// //     </div>
// //   );
// // }
