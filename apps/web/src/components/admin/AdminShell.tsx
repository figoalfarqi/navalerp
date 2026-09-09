"use client";

import type { ReactNode } from "react";
import { useState } from "react";
import Header from "@/components/Header";
import Sidebar from "@/components/sidebar";
import { ADMIN_ROLE_NAMES } from "./adminNavigation";

interface AdminShellProps {
  children: ReactNode;
  roleId: number;
}

export default function AdminShell({ children, roleId }: AdminShellProps) {
  const [mobileSidebarOpen, setMobileSidebarOpen] = useState(false);
  const [desktopCollapsed, setDesktopCollapsed] = useState(false);

  return (
    <div className="min-h-screen bg-slate-100 text-slate-900">
      <Sidebar
        roleId={roleId}
        mobileOpen={mobileSidebarOpen}
        onMobileClose={() => setMobileSidebarOpen(false)}
        desktopCollapsed={desktopCollapsed}
        onDesktopCollapsedChange={setDesktopCollapsed}
      />
      <div
        className={`min-h-screen transition-[padding] duration-300 ${
          desktopCollapsed ? "md:pl-14" : "md:pl-72"
        }`}
      >
        <Header
          onOpenMobileMenu={() => setMobileSidebarOpen(true)}
          roleName={ADMIN_ROLE_NAMES[roleId] ?? "Administrator"}
        />
        <main className="mx-auto w-full max-w-[1800px] p-0">
          {children}
        </main>
      </div>
    </div>
  );
}
