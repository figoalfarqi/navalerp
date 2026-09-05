"use client";

import { ReactNode, useEffect } from "react";
import { useRouter, usePathname } from "next/navigation";
import { useAuth } from "@/context/AuthContext";
import BottomNavChecker from "@/components/mobile/BottomNavChecker";

export default function Page({ children }: { children: ReactNode }) {
  const {
    checkerToken,
    checkerPayload,
    tokenLoaded,
    clearCheckerToken,
  } = useAuth();
  const router = useRouter();
  const pathname = usePathname();
  const hasCheckerAccess =
    Boolean(checkerToken) && checkerPayload.app_role_id === 2;

  useEffect(() => {
    if (!tokenLoaded) return;
    if (!hasCheckerAccess && pathname !== "/checker/login") {
      void clearCheckerToken();
      router.replace("/checker/login");
    } else if (!hasCheckerAccess && checkerToken) {
      void clearCheckerToken();
    } else if (hasCheckerAccess && pathname === "/checker/login") {
      router.replace("/checker/position");
    }
  }, [
    checkerPayload.app_role_id,
    checkerToken,
    clearCheckerToken,
    hasCheckerAccess,
    pathname,
    router,
    tokenLoaded,
  ]);

  if (
    !tokenLoaded ||
    (!hasCheckerAccess && pathname !== "/checker/login") ||
    (hasCheckerAccess && pathname === "/checker/login")
  ) {
    return (
      <div className="flex justify-center items-center h-screen">
        Loading...
      </div>
    );
  }

  return (
    <div className="min-h-dvh bg-slate-100">
      {children}
      {pathname !== "/checker/login" && pathname !== "/checker/position" && (
        <BottomNavChecker />
      )}
    </div>
  );
}
