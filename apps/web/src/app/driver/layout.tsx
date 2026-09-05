"use client";

import { ReactNode, useEffect } from "react";
import { useRouter, usePathname } from "next/navigation";
import { useAuth } from "@/context/AuthContext";
import BottomNav from "@/components/driver/BottomNav";

export default function Page({ children }: { children: ReactNode }) {
  const { driverToken, driverPayload, tokenLoaded, clearDriverToken } =
    useAuth();
  const router = useRouter();
  const pathname = usePathname();
  const hasDriverAccess =
    Boolean(driverToken) && driverPayload.app_role_id === 1;

  useEffect(() => {
    if (!tokenLoaded) return;
    if (!hasDriverAccess && pathname !== "/driver/login") {
      void clearDriverToken();
      router.replace("/driver/login");
    } else if (!hasDriverAccess && driverToken) {
      void clearDriverToken();
    } else if (hasDriverAccess && pathname === "/driver/login") {
      router.replace("/driver/transport");
    }
  }, [
    clearDriverToken,
    driverPayload.app_role_id,
    driverToken,
    hasDriverAccess,
    pathname,
    router,
    tokenLoaded,
  ]);

  if (
    !tokenLoaded ||
    (!hasDriverAccess && pathname !== "/driver/login") ||
    (hasDriverAccess && pathname === "/driver/login")
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
      {pathname !== "/driver/login" && <BottomNav />}
    </div>
  );
}
