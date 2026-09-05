// "use client";

// import { useEffect, useCallback, useState } from "react";
// import { authTokenType, useFetchAPI } from "@/hooks/useFetchAPI";

// const VAPID_PUBLIC_KEY =
//     process.env.NEXT_PUBLIC_VAPID_PUBLIC_KEY!;

// function urlBase64ToUint8Array(base64String: string) {
//     const padding = "=".repeat((4 - (base64String.length % 4)) % 4);
//     const base64 = (base64String + padding)
//         .replace(/-/g, "+")
//         .replace(/_/g, "/");

//     const rawData = atob(base64);
//     return Uint8Array.from([...rawData].map((c) => c.charCodeAt(0)));
// }

// export function usePushNotification(topic: string, authTokenType: authTokenType) {
//     const { postAPI } = useFetchAPI();
//     const [isSupported, setIsSupported] = useState(false);
//     const [isSubscribed, setIsSubscribed] = useState(false);

//     // register SW
//     useEffect(() => {
//         if ("serviceWorker" in navigator && "PushManager" in window) {
//             navigator.serviceWorker.register("/service-worker.js");
//             setIsSupported(true);
//         }
//     }, []);

//     // check current subscription
//     useEffect(() => {
//         if (!isSupported) return;

//         (async () => {
//             const reg = await navigator.serviceWorker.ready;
//             const sub = await reg.pushManager.getSubscription();
//             setIsSubscribed(!!sub);
//         })();
//     }, [isSupported]);

//     const subscribe = useCallback(async () => {
//         if (!isSupported) return;

//         if (Notification.permission === "denied") {
//             throw new Error("Notification permission denied");
//         }

//         if (Notification.permission !== "granted") {
//             await Notification.requestPermission();
//         }

//         const reg = await navigator.serviceWorker.ready;
//         let sub = await reg.pushManager.getSubscription();

//         if (!sub) {
//             sub = await reg.pushManager.subscribe({
//                 userVisibleOnly: true,
//                 applicationServerKey: urlBase64ToUint8Array(VAPID_PUBLIC_KEY),
//             });
//         }

//         await postAPI(
//             `${process.env.NEXT_PUBLIC_API_BASE_URL}/push/subscribe`,
//             { authToken: authTokenType },
//             {
//                 endpoint: sub.endpoint,
//                 keys: {
//                     p256dh: btoa(
//                         String.fromCharCode(...new Uint8Array(sub.getKey("p256dh")!))
//                     ),
//                     auth: btoa(
//                         String.fromCharCode(...new Uint8Array(sub.getKey("auth")!))
//                     ),
//                 },
//                 topic,
//                 user_agent: navigator.userAgent,
//             }
//         );

//         setIsSubscribed(true);
//     }, [isSupported, postAPI, topic]);

//     const unsubscribe = useCallback(async () => {
//         if (!isSupported) return;

//         const reg = await navigator.serviceWorker.ready;
//         const sub = await reg.pushManager.getSubscription();
//         if (!sub) return;

//         await postAPI(
//             `${process.env.NEXT_PUBLIC_API_BASE_URL}/push/unsubscribe`,
//             { authToken: authTokenType },
//             { endpoint: sub.endpoint }
//         );

//         await sub.unsubscribe();
//         setIsSubscribed(false);
//     }, [isSupported, postAPI]);

//     return {
//         isSupported,
//         isSubscribed,
//         subscribe,
//         unsubscribe,
//     };
// }


"use client";

import { useEffect, useCallback, useState } from "react";
import { authTokenType, useFetchAPI } from "@/hooks/useFetchAPI";
import { useAuth } from "@/context/AuthContext";

const VAPID_PUBLIC_KEY = process.env.NEXT_PUBLIC_VAPID_PUBLIC_KEY!;

/* ----------------------------- utils ----------------------------- */

function urlBase64ToUint8Array(base64String: string) {
    const padding = "=".repeat((4 - (base64String.length % 4)) % 4);
    const base64 = (base64String + padding)
        .replace(/-/g, "+")
        .replace(/_/g, "/");

    const rawData = atob(base64);
    return Uint8Array.from([...rawData].map((c) => c.charCodeAt(0)));
}

/** persistent per browser profile */
function getDeviceId(): string {
    const key = "push_device_id";
    let id = localStorage.getItem(key);

    if (!id) {
        id = crypto.randomUUID();
        localStorage.setItem(key, id);
    }
    return id;
}

function getHumanPlatform(ua: string) {

    if (/Windows NT 10.0/.test(ua)) return "Windows";
    if (/Mac OS X/.test(ua)) return "macOS";
    if (/Android/.test(ua)) return "Android";
    if (/iPhone|iPad/.test(ua)) return "iOS";
    if (/Linux/.test(ua)) return "Linux";

    return "Unknown";
}

function getDeviceInfo() {
    const ua = navigator.userAgent;

    const browser =
        ua.includes("Edg") ? "Edge" :
            ua.includes("Chrome") ? "Chrome" :
                ua.includes("Firefox") ? "Firefox" :
                    ua.includes("Safari") ? "Safari" :
                        "Unknown";

    return {
        device_id: getDeviceId(),
        browser,
        platform: getHumanPlatform(ua),
        user_agent: ua,
    };
}

/* ----------------------------- hook ----------------------------- */

export function usePushNotification(
    topic: string,
    authTokenType: authTokenType
) {
    const { postAPI, getAPI, patchAPI } = useFetchAPI();

    const [isSupported, setIsSupported] = useState(false);
    const [isSubscribed, setIsSubscribed] = useState(false);

    const { driverPayload, adminPayload } = useAuth();

    /* -------- register SW -------- */
    useEffect(() => {
        if (
            typeof window !== "undefined" &&
            "serviceWorker" in navigator &&
            "PushManager" in window
        ) {
            navigator.serviceWorker.register("/service-worker.js");
            setIsSupported(true);
        }
    }, []);

    /* -------- check subscription -------- */
    // useEffect(() => {
    //     if (!isSupported) return;

    //     (async () => {
    //         const reg = await navigator.serviceWorker.ready;
    //         const sub = await reg.pushManager.getSubscription();
    //         console.log("AAAAAAAAAAAAHHHHHHHHHHHHH", sub)
    //         setIsSubscribed(Boolean(sub));
    //     })();
    // }, [isSupported]);

    /* ---------------- subscribe ---------------- */
    const subscribe = useCallback(async () => {
        if (!isSupported) return;

        // if (Notification.permission === "denied") {
        //     throw new Error("Notification permission denied");
        // }

        if (Notification.permission === "default") {
            await Notification.requestPermission();
        }

        if (Notification.permission !== "granted") {
            return;
        }

        const reg = await navigator.serviceWorker.ready;
        let sub = await reg.pushManager.getSubscription();

        if (!sub) {
            sub = await reg.pushManager.subscribe({
                userVisibleOnly: true,
                applicationServerKey: urlBase64ToUint8Array(VAPID_PUBLIC_KEY),
            });
        }

        const res = await postAPI<any>(
            `${process.env.NEXT_PUBLIC_API_BASE_URL}/push/subscribe`,
            { authToken: authTokenType },
            {
                endpoint: sub.endpoint,
                keys: {
                    p256dh: btoa(
                        String.fromCharCode(...new Uint8Array(sub.getKey("p256dh")!))
                    ),
                    auth: btoa(
                        String.fromCharCode(...new Uint8Array(sub.getKey("auth")!))
                    ),
                },
                topic,
                ...getDeviceInfo(),
            }
        );
        if (res.code && [200, 201].includes(res.code)) {
            setIsSubscribed(true);
        } else {
            setIsSubscribed(false);
        }

    }, [isSupported, postAPI, topic, authTokenType]);

    /* ---------------- unsubscribe current device ---------------- */
    const unsubscribeCurrentDevice = useCallback(async () => {
        if (!isSupported) return;

        const reg = await navigator.serviceWorker.ready;
        const sub = await reg.pushManager.getSubscription();
        if (!sub) return;

        await postAPI(
            `${process.env.NEXT_PUBLIC_API_BASE_URL}/push/unsubscribe`,
            { authToken: authTokenType },
            {
                endpoint: sub.endpoint,
                device_id: getDeviceId(),
            }
        );

        await sub.unsubscribe();
        setIsSubscribed(false);
    }, [isSupported, postAPI, authTokenType]);

    /* ---------------- unsubscribe other device ---------------- */
    const unsubscribeOtherDevice = useCallback(
        async (deviceId: string) => {
            await postAPI(
                `${process.env.NEXT_PUBLIC_API_BASE_URL}/push/unsubscribe_device`,
                { authToken: authTokenType },
                { device_id: deviceId, app_user_id: authTokenType === "admin" ? adminPayload?.app_user_id : driverPayload?.app_user_id }
            );
        },
        [postAPI, authTokenType]
    );

    /* ---------------- get active devices ---------------- */
    const getDevices = useCallback(async () => {
        return await getAPI<any>(
            `${process.env.NEXT_PUBLIC_API_BASE_URL}/push/device`,
            { authToken: authTokenType }
        );
    }, [getAPI, authTokenType]);

    /* ---------------- update device label ---------------- */
    const updateDeviceLabel = useCallback(
        async (deviceId: string, label: string) => {
            await patchAPI<any>(
                `${process.env.NEXT_PUBLIC_API_BASE_URL}/push/device_label`,
                { authToken: authTokenType },
                {
                    device_id: deviceId,
                    device_label: label,
                }
            );
        },
        [patchAPI, authTokenType]
    );

    const sendTestPush = useCallback(async () => {
        await postAPI(
            `${process.env.NEXT_PUBLIC_API_BASE_URL}/push/send_to_target`,
            { authToken: authTokenType },
            {},
        );
    }, [postAPI, authTokenType])

    return {
        isSupported,
        isSubscribed,
        setIsSubscribed,
        getDeviceId,
        subscribe,
        unsubscribeCurrentDevice,
        unsubscribeOtherDevice,
        getDevices,
        updateDeviceLabel,
        sendTestPush,
    };
}


// TODO: melihat berapa device, dan bisa mematikan notifikasi
// cara pakai
// const { subscribe } = usePushNotification("order");
// await subscribe();


// Field	Untuk apa
// endpoint	alamat push
// p256dh	enkripsi
// auth	enkripsi
// device_id	identitas browser
// browser	UI info
// platform	UI info
// user_agent	debug