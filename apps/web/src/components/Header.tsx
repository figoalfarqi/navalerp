"use client";

import { useAuth } from "@/context/AuthContext";
import { usePathname, useRouter } from "next/navigation";
import { useEffect, useMemo, useRef, useState } from "react";
import {
  FaAngleRight,
  FaBars,
  FaGear,
  FaRightFromBracket,
  FaUser,
} from "react-icons/fa6";

interface HeaderProps {
  onOpenMobileMenu: () => void;
  roleName: string;
}

export default function Header({
  onOpenMobileMenu,
  roleName,
}: HeaderProps) {
  const router = useRouter();
  const pathname = usePathname();
  const { adminPayload, clearAdminToken } = useAuth();
  const [userMenuOpen, setUserMenuOpen] = useState(false);
  const userMenuRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const closeWhenOutside = (event: MouseEvent) => {
      if (
        userMenuRef.current &&
        !userMenuRef.current.contains(event.target as Node)
      ) {
        setUserMenuOpen(false);
      }
    };
    document.addEventListener("mousedown", closeWhenOutside);
    return () => document.removeEventListener("mousedown", closeWhenOutside);
  }, []);

  const breadcrumbs = useMemo(() => {
    const segments = pathname.split("/").filter(Boolean);
    return segments
      .filter((segment) => !/^\d+$/.test(segment))
      .map((segment, index, filtered) => ({
        label:
          segment === "admin"
            ? "Dashboard"
            : segment
                .replace(/_/g, " ")
                .replace(/\b\w/g, (letter) => letter.toUpperCase()),
        href:
          segment === "admin"
            ? "/admin"
            : `/${filtered.slice(0, index + 1).join("/")}`,
      }));
  }, [pathname]);

  const logout = async () => {
    await clearAdminToken();
    router.replace("/admin/login");
  };

  return (
    <>
      <header className="sticky top-0 z-30 border-b border-slate-200/80 bg-white/95 px-3 py-1 shadow-sm backdrop-blur sm:px-5 lg:px-7">
        <div className="mx-auto flex min-h-11 max-w-[1800px] items-center justify-between gap-3">
          <div className="flex min-w-0 items-center gap-3">
            <button
              type="button"
              onClick={onOpenMobileMenu}
              className="rounded-xl border border-slate-200 p-2.5 text-slate-600 hover:bg-slate-50 md:hidden"
              aria-label="Buka menu"
            >
              <FaBars />
            </button>
            <nav className="hidden min-w-0 items-center sm:flex">
              {breadcrumbs.map((breadcrumb, index) => (
                <span
                  key={breadcrumb.href}
                  className="flex min-w-0 items-center"
                >
                  {index > 0 && (
                    <FaAngleRight
                      className="mx-2 shrink-0 text-slate-300"
                      size={12}
                    />
                  )}
                  <button
                    type="button"
                    onClick={() => router.push(breadcrumb.href)}
                    className={`truncate text-sm ${
                      index === breadcrumbs.length - 1
                        ? "font-semibold text-slate-800"
                        : "text-slate-500 hover:text-blue-600"
                    }`}
                  >
                    {breadcrumb.label}
                  </button>
                </span>
              ))}
            </nav>
            <span className="truncate font-semibold text-slate-800 sm:hidden">
              {breadcrumbs.at(-1)?.label ?? "Dashboard"}
            </span>
          </div>

          <div className="flex shrink-0 items-center gap-2">
            <div ref={userMenuRef} className="relative">
              <button
                type="button"
                onClick={() => setUserMenuOpen((open) => !open)}
                className="flex items-center gap-2 text-left hover:bg-slate-50"
              >
                <span className="grid h-7 w-7 place-items-center rounded-lg bg-blue-100 text-blue-700">
                  <FaUser size={13} />
                </span>
                <span className="hidden max-w-40 sm:block">
                  <span className="block truncate text-xs font-semibold text-slate-800">
                    {adminPayload?.app_user_name || adminPayload?.username}
                  </span>
                  <span className="block truncate text-[11px] text-slate-500">
                    {roleName}
                  </span>
                </span>
              </button>
              {userMenuOpen && (
                <div className="absolute right-0 mt-2 w-56 overflow-hidden rounded-xl border border-slate-200 bg-white p-1.5 shadow-xl">
                  <button
                    type="button"
                    onClick={() => {
                      setUserMenuOpen(false);
                      router.push("/admin/settings");
                    }}
                    className="flex w-full items-center gap-3 rounded-lg px-3 py-2.5 text-sm text-slate-700 hover:bg-slate-100"
                  >
                    <FaGear className="text-slate-400" />
                    Pengaturan
                  </button>
                  <button
                    type="button"
                    onClick={logout}
                    className="flex w-full items-center gap-3 rounded-lg px-3 py-2.5 text-sm text-red-600 hover:bg-red-50"
                  >
                    <FaRightFromBracket />
                    Keluar
                  </button>
                </div>
              )}
            </div>
          </div>
        </div>
      </header>

    </>
  );
}
