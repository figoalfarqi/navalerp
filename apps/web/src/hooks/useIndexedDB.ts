/* eslint-disable @typescript-eslint/no-explicit-any */
// hooks/useIndexedDB.ts
import { DB_NAME, DB_VERSION, StoreName, STORES } from "@/types/tableIDB";
import { useEffect, useState } from "react";

/* -------------------------
   🔹 Singleton Connection Cache
-------------------------- */
const dbCache: Record<string, IDBDatabase> = {};

// Async function untuk mendapatkan / membuat instance IndexedDB
async function getDBInstance(): Promise<IDBDatabase> {
  if (dbCache[DB_NAME]) return dbCache[DB_NAME];

  return new Promise((resolve, reject) => {
    const request = indexedDB.open(DB_NAME, DB_VERSION);

    request.onupgradeneeded = (event) => {
      const db = (event.target as IDBOpenDBRequest).result;

      Object.values(STORES).forEach((storeName) => {
        if (!db.objectStoreNames.contains(storeName)) {
          db.createObjectStore(storeName, {
            keyPath: `${storeName}_id`,
            autoIncrement: true,
          });
        }
      });
    };

    request.onsuccess = () => {
      const db = request.result;

      db.onclose = () => delete dbCache[DB_NAME];
      db.onversionchange = () => {
        db.close();
        delete dbCache[DB_NAME];
      };

      dbCache[DB_NAME] = db;
      resolve(db);
    };

    request.onerror = () => reject(request.error);
  });
}

/* -------------------------
   🔹 Hook utama
-------------------------- */
interface IDBResult<T> {
  data: T | null;
  error: Error | null;
  loading: boolean;
}

export function useIndexedDB() {
  const [dbReady, setDbReady] = useState(false);
  const [db, setDb] = useState<IDBDatabase | null>(null);

  useEffect(() => {
    let mounted = true;
    getDBInstance()
      .then((database) => {
        if (mounted) {
          setDb(database);
          setDbReady(true);
        }
      })
      .catch(console.error);

    return () => {
      mounted = false;
    };
  }, []);

  // Helper buat buka transaksi
  const getStore = (storeName: StoreName, mode: IDBTransactionMode = "readonly") => {
    if (!db) throw new Error("Database not initialized");
    const tx = db.transaction(storeName, mode);
    return tx.objectStore(storeName);
  };

  // Template executor biar ada data/error/loading
  const execute = async <T>(fn: () => Promise<T>): Promise<IDBResult<T>> => {
    const result: IDBResult<T> = { data: null, error: null, loading: true };
    try {
      const data = await fn();
      result.data = data;
    } catch (err) {
      result.error = err as Error;
    } finally {
      result.loading = false;
    }
    return result;
  };

  /* -------------------------
     CRUD METHODS
  -------------------------- */

  const getIDB: (storeName: StoreName, key?: IDBValidKey) => Promise<IDBResult<any>> = async (storeName: StoreName, key?: IDBValidKey) =>
    execute(async () => {
      return new Promise<any>((resolve, reject) => {
        const store = getStore(storeName, "readonly");
        const request = key ? store.get(key) : store.getAll();
        request.onsuccess = () => resolve(request.result);
        request.onerror = () => reject(request.error);
      });
    });

  const postIDB = async (storeName: StoreName, data: any) =>
    execute(async () => {
      return new Promise<IDBValidKey>((resolve, reject) => {
        const store = getStore(storeName, "readwrite");
        const request = store.add(data);
        request.onsuccess = () => resolve(request.result);
        request.onerror = () => reject(request.error);
      });
    });

  const putIDB = async (storeName: StoreName, data: any) =>
    execute(async () => {
      return new Promise<IDBValidKey>((resolve, reject) => {
        const store = getStore(storeName, "readwrite");
        const request = store.put(data);
        request.onsuccess = () => resolve(request.result);
        request.onerror = () => reject(request.error);
      });
    });

  const patchIDB = async (storeName: StoreName, key: IDBValidKey, partialData: any) =>
    execute(async () => {
      return new Promise<any>((resolve, reject) => {
        const store = getStore(storeName, "readwrite");
        const getReq = store.get(key);

        getReq.onsuccess = () => {
          const existing = getReq.result;
          if (!existing) return reject(new Error("Data not found"));
          const updated = { ...existing, ...partialData };
          const putReq = store.put(updated);
          putReq.onsuccess = () => resolve(updated);
          putReq.onerror = () => reject(putReq.error);
        };
        getReq.onerror = () => reject(getReq.error);
      });
    });

  const deleteIDB = async (storeName: StoreName, key: IDBValidKey) =>
    execute(async () => {
      return new Promise<boolean>((resolve, reject) => {
        const store = getStore(storeName, "readwrite");
        const request = store.delete(key);
        request.onsuccess = () => resolve(true);
        request.onerror = () => reject(request.error);
      });
    });

  const clearIDB = async (storeName: StoreName) =>
    execute(async () => {
      return new Promise<boolean>((resolve, reject) => {
        const store = getStore(storeName, "readwrite");
        const request = store.clear();
        request.onsuccess = () => resolve(true);
        request.onerror = () => reject(request.error);
      });
    });

  const setIDB = async (storeName: StoreName, dataArray: any[]) =>
    execute(async () => {
      return new Promise<boolean>((resolve, reject) => {
        const store = getStore(storeName, "readwrite");
        const clearReq = store.clear();
        clearReq.onsuccess = () => {
          dataArray.forEach((item) => store.put(item));
          resolve(true);
        };
        clearReq.onerror = () => reject(clearReq.error);
      });
    });

  const mergeIDB = async (storeName: StoreName, dataArray: any[]) =>
    execute(async () => {
      return new Promise<boolean>((resolve, reject) => {
        try {
          const store = getStore(storeName, "readwrite");
          dataArray.forEach((item) => store.put(item));
          resolve(true);
        } catch (err) {
          reject(err);
        }
      });
    });

  const deleteDatabase = async (): Promise<boolean> => {
    const idbChannel = new BroadcastChannel("idb-control");

    return new Promise(async (resolve, reject) => {
      try {
        idbChannel.postMessage("force-close");
        await new Promise((res) => setTimeout(res, 300));

        if (db) db.close();
        delete dbCache[DB_NAME];

        const request = indexedDB.deleteDatabase(DB_NAME);

        request.onsuccess = async () => {
          setDb(null);
          setDbReady(false);

          const newDB = await getDBInstance();
          setDb(newDB);
          setDbReady(true);

          resolve(true);
        };

        request.onerror = () => reject(request.error);

        request.onblocked = () => {
          reject(new Error("Database still blocked"));
        };
      } catch (err) {
        reject(err);
      }
    });
  };

  return {
    dbReady,
    getIDB,
    postIDB,
    putIDB,
    patchIDB,
    deleteIDB,
    clearIDB,
    setIDB,
    mergeIDB,
    deleteDatabase,
  };
}
