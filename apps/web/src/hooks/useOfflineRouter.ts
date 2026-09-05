// src/hooks/useOfflineRouter.ts
"use client";

import { useRouter } from "next/navigation";
import { useOnlineStatus } from "./useOnlineStatus";
import { canNavigate } from "@/utils/chace";
import { useToast } from "@/components/ToastContext";

type PushOptions = {
    force?: boolean; // bypass offline check
};

export function useOfflineRouter() {
    const router = useRouter();
    const { isOnline } = useOnlineStatus();
    const { showToast } = useToast();

    const routePush = async (url: string, options?: PushOptions) => {
        if (options?.force) {
            router.push(url);
            return;
        }

        if (isOnline) {
            router.push(url);
            return;
        }

        const availableOffline = await canNavigate(url);

        if (!availableOffline) {
            showToast(3000, "error", "Anda sedang offline dan halaman ini tidak tersedia di perangkat");
            return;
        }

        router.push(url);
    };

    return {
        routePush,
        isOnline,
    };
}
