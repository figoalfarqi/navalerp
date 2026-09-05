"use client";

import Image from "next/image";
import { usePathname, useRouter } from "next/navigation";
import { useMemo, useState } from "react";
import {
  FaAngleDown,
  FaAngleLeft,
  FaAngleRight,
  FaXmark,
} from "react-icons/fa6";
import {
  getAdminNavigation,
  type AdminNavigationGroup,
} from "@/components/admin/adminNavigation";

interface SidebarProps {
  roleId: number;
  mobileOpen: boolean;
  onMobileClose: () => void;
  desktopCollapsed: boolean;
  onDesktopCollapsedChange: (collapsed: boolean) => void;
}

function isItemActive(pathname: string, href: string) {
  return (
    pathname === href ||
    (href !== "/admin" && pathname.startsWith(`${href}/`))
  );
}

export default function Sidebar({
  roleId,
  mobileOpen,
  onMobileClose,
  desktopCollapsed,
  onDesktopCollapsedChange,
}: SidebarProps) {
  const pathname = usePathname();
  const router = useRouter();
  const groups = useMemo(() => getAdminNavigation(roleId), [roleId]);
  const [expandedGroups, setExpandedGroups] = useState<Record<string, boolean>>(
    {},
  );

  const isGroupOpen = (group: AdminNavigationGroup) =>
    expandedGroups[group.label] ??
    group.items.some((item) => isItemActive(pathname, item.href));

  const navigate = (href: string) => {
    router.push(href);
    onMobileClose();
  };

  const sidebarBody = (
    <>
      <div className="flex h-20 items-center border-b border-white/10 px-4">
        <button
          type="button"
          onClick={() => navigate("/admin")}
          className="flex min-w-0 items-center gap-3 text-left"
          aria-label="Buka dashboard"
        >
          <Image
            src="/logo.png"
            alt="Naval ERP"
            width={42}
            height={42}
            className="shrink-0 object-contain"
          />
          <span
            className={`min-w-0 ${desktopCollapsed ? "md:hidden" : ""}`}
          >
              <span className="block truncate text-lg font-bold tracking-wide text-white">
                NAVAL ERP
              </span>
              <span className="block truncate text-xs text-cyan-200">
                TNI Angkatan Laut
              </span>
          </span>
        </button>
        <button
          type="button"
          onClick={onMobileClose}
          className="ml-auto rounded-lg p-2 text-white/80 hover:bg-white/10 md:hidden"
          aria-label="Tutup menu"
        >
          <FaXmark size={20} />
        </button>
      </div>

      <nav className="h-[calc(100vh-5rem-3.5rem)] overflow-y-auto px-3 py-4 custom-scrollbar">
        <div className="space-y-2">
          {groups.map((group) => {
            const GroupIcon = group.icon;
            const open = isGroupOpen(group);
            return (
              <section key={group.label}>
                <button
                  type="button"
                  onClick={() => {
                    if (desktopCollapsed) {
                      onDesktopCollapsedChange(false);
                    }
                    setExpandedGroups((current) => ({
                      ...current,
                      [group.label]: !open,
                    }));
                  }}
                  className={`flex w-full items-center rounded-xl px-3 py-2.5 text-sm font-medium transition ${
                    group.items.some((item) =>
                      isItemActive(pathname, item.href),
                    )
                      ? "bg-white/12 text-white"
                      : "text-cyan-50/85 hover:bg-white/10 hover:text-white"
                  }`}
                  title={desktopCollapsed ? group.label : undefined}
                >
                  <GroupIcon size={19} className="shrink-0" />
                  <span
                    className={`contents ${desktopCollapsed ? "md:hidden" : ""}`}
                  >
                      <span className="ml-3 flex-1 text-left">
                        {group.label}
                      </span>
                      <FaAngleDown
                        size={13}
                        className={`transition-transform ${open ? "rotate-180" : ""}`}
                      />
                  </span>
                </button>

                {open && (
                  <div
                    className={`mt-1 space-y-1 pl-3 ${
                      desktopCollapsed ? "md:hidden" : ""
                    }`}
                  >
                    {group.items.map((item) => {
                      const ItemIcon = item.icon;
                      const active = isItemActive(pathname, item.href);
                      return (
                        <button
                          key={item.href}
                          type="button"
                          onClick={() => navigate(item.href)}
                          className={`flex w-full items-center rounded-xl px-3 py-2 text-left text-sm transition ${
                            active
                              ? "bg-white text-[#125aaa] shadow-sm"
                              : "text-cyan-50/80 hover:bg-white/10 hover:text-white"
                          }`}
                        >
                          <ItemIcon size={16} className="shrink-0" />
                          <span className="ml-3 truncate">{item.label}</span>
                        </button>
                      );
                    })}
                  </div>
                )}
              </section>
            );
          })}
        </div>
      </nav>

      <button
        type="button"
        onClick={() => onDesktopCollapsedChange(!desktopCollapsed)}
        className="hidden h-14 w-full items-center justify-center gap-2 border-t border-white/10 text-sm text-cyan-50/80 transition hover:bg-white/10 hover:text-white md:flex"
        aria-label={desktopCollapsed ? "Perbesar sidebar" : "Kecilkan sidebar"}
      >
        {desktopCollapsed ? <FaAngleRight /> : <FaAngleLeft />}
        {!desktopCollapsed && <span>Kecilkan menu</span>}
      </button>
    </>
  );

  return (
    <>
      <button
        type="button"
        aria-label="Tutup menu"
        onClick={onMobileClose}
        className={`fixed inset-0 z-40 bg-slate-950/50 backdrop-blur-[1px] transition-opacity md:hidden ${
          mobileOpen
            ? "pointer-events-auto opacity-100"
            : "pointer-events-none opacity-0"
        }`}
      />
      <aside
        className={`fixed inset-y-0 left-0 z-50 w-72 overflow-hidden bg-gradient-to-b from-[#0b8fa5] via-[#1269ae] to-[#163d7a] shadow-2xl transition-transform duration-300 md:translate-x-0 ${
          desktopCollapsed ? "md:w-16" : "md:w-72"
        } ${mobileOpen ? "translate-x-0" : "-translate-x-full"}`}
      >
        {sidebarBody}
      </aside>
    </>
  );
}
