"use client";

import type { ReactNode } from "react";
import { useEffect } from "react";
import { usePathname, useRouter } from "next/navigation";
import { useAuth } from "@/context/AuthContext";
import DesktopPageLoader from "@/components/admin/DesktopPageLoader";
import AdminShell from "@/components/admin/AdminShell";
import {
  canAccessAdminPath,
  isAdminRoleId,
} from "@/components/admin/adminNavigation";

export default function AdminLayout({ children }: { children: ReactNode }) {
  const { adminToken, adminPayload, tokenLoaded, clearAdminToken } = useAuth();
  const router = useRouter();
  const pathname = usePathname();
  const isLoginPage = pathname === "/admin/login";
  const roleId = adminPayload?.app_role_id ?? 0;

  useEffect(() => {
    if (!tokenLoaded) return;

    if (!adminToken) {
      if (!isLoginPage) router.replace("/admin/login");
      return;
    }

    if (!isAdminRoleId(roleId)) {
      void clearAdminToken();
      router.replace("/admin/login");
      return;
    }

    if (isLoginPage) {
      router.replace("/admin");
      return;
    }

    if (!canAccessAdminPath(roleId, pathname)) {
      router.replace("/admin");
    }
  }, [
    adminToken,
    clearAdminToken,
    isLoginPage,
    pathname,
    roleId,
    router,
    tokenLoaded,
  ]);

  if (!tokenLoaded) return <DesktopPageLoader />;
  if (isLoginPage) return <>{children}</>;
  if (!adminToken || !isAdminRoleId(roleId)) return <DesktopPageLoader />;

  return <AdminShell roleId={roleId}>{children}</AdminShell>;
}
