// "use client";

// import { useEffect } from "react";
// import useOnlineSync from "@/hooks/useOnlineSync";

// export default function LayoutClientSync() {
//   useOnlineSync();

//   /* ===============================
//      REGISTER + UPDATE SERVICE WORKER
//      =============================== */
//   useEffect(() => {
//     if (!("serviceWorker" in navigator)) return;

//     let messageHandler: ((e: MessageEvent) => void) | null = null;

//     (async () => {
//       try {
//         const reg = await navigator.serviceWorker.register("/service-worker.js");
//         console.log("✅ Service Worker registered");

//         // paksa cek update
//         await reg.update();

//         // listen message dari SW
//         messageHandler = (event: MessageEvent) => {
//           if (event.data?.type === "update-idb") {
//             console.log("Update dari SW:", event.data.payload);
//           }
//         };

//         navigator.serviceWorker.addEventListener(
//           "message",
//           messageHandler
//         );
//       } catch (err) {
//         console.error("❌ SW registration failed:", err);
//       }
//     })();

//     return () => {
//       if (messageHandler) {
//         navigator.serviceWorker.removeEventListener(
//           "message",
//           messageHandler
//         );
//       }
//     };
//   }, []);

//   /* ===============================
//      STORAGE PERSISTENCE (OPTIONAL)
//      =============================== */
//   useEffect(() => {
//     if ("storage" in navigator && "persist" in navigator.storage) {
//       navigator.storage.persist().then((granted) => {
//         console.log("Storage persisted:", granted);
//       });
//     }
//   }, []);

//   return null;
// }
