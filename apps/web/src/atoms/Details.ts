// import { DetailsMap } from "@/hooks/useDetails";
// import { atom } from "jotai";

// export const detailsAtom = atom<DetailsMap<any>>({});

// const createInitialDetailsMap = <T,>(
//     keys: string[],
//     defaultFormData?: Record<string, Record<string, any>>
// ): DetailsMap<T> => {
//     const obj: DetailsMap<T> = {};

//     keys.forEach((key) => {
//         obj[key] = {
//             mode: "edit",
//             formData: defaultFormData?.[key] ?? {},
//             initialAdder: defaultFormData?.[key] ?? {},
//             initialDataTables: [],
//             modifiedDataTables: {
//                 added_datas: [],
//                 updated_datas: [],
//                 deleted_ids: [],
//             },
//             localIdCounter: 1,
//         };
//     });

//     return obj;
// };


// export const initDetailsAtom = atom(
//     null,
//     (
//         get,
//         set,
//         payload: {
//             keys: string[];
//             defaultFormData?: Record<string, Record<string, any>>;
//         }
//     ) => {
//         const current = get(detailsAtom);
//         const currentKeys = Object.keys(current);

//         const same =
//             currentKeys.length === payload.keys.length &&
//             payload.keys.every((k) => currentKeys.includes(k));

//         if (same) return;

//         set(
//             detailsAtom,
//             createInitialDetailsMap(payload.keys, payload.defaultFormData)
//         );
//     }
// );
