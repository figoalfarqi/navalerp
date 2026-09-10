"use client";

import Image from "next/image";
import { usePathname, useRouter } from "next/navigation";
import { useMemo, useState } from "react";
import {
  ChevronDownIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
  CloseIcon,
} from "@/components/icons";
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
      <div
        className={`flex h-16 shrink-0 items-center border-b border-white/10 ${
          desktopCollapsed ? "md:px-0 md:justify-center px-4" : "px-4"
        }`}
      >
        <button
          type="button"
          onClick={() => navigate("/admin")}
          className={`flex min-w-0 items-center gap-3 text-left cursor-pointer select-none ${
            desktopCollapsed ? "md:justify-center md:w-full" : ""
          }`}
          aria-label="Buka dashboard"
          title={desktopCollapsed ? "Naval ERP - TNI Angkatan Laut" : undefined}
        >
          <Image
            src="/logo.webp"
            alt="Naval ERP"
            width={32}
            height={32}
            className="shrink-0 object-contain cursor-pointer"
          />
          <span
            className={`min-w-0 cursor-pointer ${desktopCollapsed ? "md:hidden" : ""}`}
          >
            <span className="block truncate text-base font-bold tracking-wide text-white cursor-pointer">
              NAVAL ERP
            </span>
            <span className="block truncate text-[11px] text-sky-300 cursor-pointer">
              TNI Angkatan Laut
            </span>
          </span>
        </button>
        <button
          type="button"
          onClick={onMobileClose}
          className="ml-auto rounded-lg p-2 text-white/80 hover:bg-white/10 md:hidden cursor-pointer"
          aria-label="Tutup menu"
        >
          <CloseIcon size={20} className="cursor-pointer" />
        </button>
      </div>

      <nav
        className={`h-[calc(100vh-4rem-2.5rem)] overflow-y-auto py-3 custom-scrollbar ${
          desktopCollapsed ? "md:px-1.5 px-3" : "px-3"
        }`}
      >
        <div className={`space-y-1.5 ${desktopCollapsed ? "md:space-y-2" : ""}`}>
          {groups.map((group) => {
            const GroupIcon = group.icon;
            const open = isGroupOpen(group);
            const isGroupActive = group.items.some((item) =>
              isItemActive(pathname, item.href),
            );

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
                  className={`flex items-center text-sm font-medium transition cursor-pointer select-none ${
                    desktopCollapsed
                      ? "md:h-10 md:w-10 md:mx-auto md:justify-center md:p-0 md:rounded-xl w-full rounded-xl px-3 py-2.5"
                      : "w-full rounded-xl px-3 py-2.5"
                  } ${
                    isGroupActive
                      ? desktopCollapsed
                        ? "md:bg-white md:text-[#081d38] md:shadow-md bg-white/12 text-white"
                        : "bg-white/12 text-white"
                      : "text-slate-200/90 hover:bg-white/10 hover:text-white"
                  }`}
                  title={desktopCollapsed ? group.label : undefined}
                >
                  <GroupIcon size={20} className="shrink-0 cursor-pointer" />
                  <span
                    className={`contents cursor-pointer ${
                      desktopCollapsed ? "md:hidden" : ""
                    }`}
                  >
                    <span className="ml-3 flex-1 text-left cursor-pointer">
                      {group.label}
                    </span>
                    <ChevronDownIcon
                      size={13}
                      className={`transition-transform cursor-pointer ${
                        open ? "rotate-180" : ""
                      }`}
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
                          className={`flex w-full items-center rounded-xl px-3 py-2 text-left text-sm transition cursor-pointer select-none ${
                            active
                              ? "bg-white text-[#081d38] shadow-sm font-semibold"
                              : "text-slate-300 hover:bg-white/10 hover:text-white"
                          }`}
                        >
                          <ItemIcon size={16} className="shrink-0 cursor-pointer" />
                          <span className="ml-3 truncate cursor-pointer">{item.label}</span>
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
        className={`hidden border-t border-white/10 text-slate-300 transition hover:bg-white/10 hover:text-white md:flex items-center justify-center cursor-pointer select-none shrink-0 ${
          desktopCollapsed
            ? "h-10 w-full"
            : "h-10 w-full gap-2 px-3 text-xs font-medium"
        }`}
        aria-label={desktopCollapsed ? "Perbesar sidebar" : "Kecilkan sidebar"}
        title={desktopCollapsed ? "Perbesar sidebar" : "Kecilkan menu"}
      >
        {desktopCollapsed ? (
          <ChevronRightIcon size={16} className="cursor-pointer" />
        ) : (
          <>
            <ChevronLeftIcon size={14} className="cursor-pointer" />
            <span className="cursor-pointer">Kecilkan menu</span>
          </>
        )}
      </button>
    </>
  );

  return (
    <>
      {mobileOpen && (
        <button
          type="button"
          aria-label="Tutup menu"
          onClick={onMobileClose}
          className="fixed inset-0 z-40 bg-slate-950/50 backdrop-blur-[1px] transition-opacity md:hidden cursor-pointer"
        />
      )}
      <aside
        className={`fixed inset-y-0 left-0 z-50 w-72 overflow-hidden bg-gradient-to-b from-[#081d38] via-[#0b2447] to-[#051428] shadow-2xl transition-[width,transform] duration-300 md:translate-x-0 ${
          desktopCollapsed ? "md:w-14" : "md:w-72"
        } ${
          mobileOpen
            ? "translate-x-0 pointer-events-auto"
            : "-translate-x-full pointer-events-none md:pointer-events-auto"
        }`}
      >
        {sidebarBody}
      </aside>
    </>
  );
}
