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
} from "@/components/icons";

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

  const displayUserName = useMemo(() => {
    const raw = adminPayload?.username || adminPayload?.app_user_name || "Admin";
    return raw
      .replace(/[._]/g, " ")
      .replace(/\b\w/g, (char) => char.toUpperCase());
  }, [adminPayload?.username, adminPayload?.app_user_name]);

  const breadcrumbs = useMemo(() => {
    const segments = pathname.split("/").filter(Boolean);
    const isIdSegment = (segment: string, index: number, allSegments: string[]): boolean => {
      if (/^\d+$/.test(segment)) return true;
      if (/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/.test(segment)) return true;
      if (/^[0-9a-fA-F]{24,}$/.test(segment)) return true;
      if (index < allSegments.length - 1) {
        const nextSegment = allSegments[index + 1].toLowerCase();
        if (["edit", "view", "copy", "detail", "delete"].includes(nextSegment)) {
          return true;
        }
      }
      return false;
    };

    const result: Array<{ label: string; href: string }> = [];
    let accumulatedPath = "";

    segments.forEach((segment, index) => {
      accumulatedPath += `/${segment}`;

      if (isIdSegment(segment, index, segments)) {
        return;
      }

      const isLast = index === segments.length - 1;
      const href = isLast
        ? pathname
        : segment === "admin"
          ? "/admin"
          : accumulatedPath;

      const label =
        segment === "admin"
          ? "Dashboard"
          : segment
              .replace(/_/g, " ")
              .replace(/\b\w/g, (letter) => letter.toUpperCase());

      result.push({ label, href });
    });

    return result;
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
              className="rounded-xl border border-slate-200 p-2.5 text-slate-600 hover:bg-slate-50 md:hidden cursor-pointer"
              aria-label="Buka menu"
            >
              <FaBars />
            </button>
            <nav className="hidden min-w-0 items-center sm:flex">
              {breadcrumbs.map((breadcrumb, index) => (
                <span
                  key={`${breadcrumb.href}-${index}`}
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
                    className={`truncate text-sm cursor-pointer ${
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
                className="flex items-center gap-2 text-left hover:bg-slate-50 cursor-pointer"
              >
                <span className="grid h-7 w-7 place-items-center rounded-lg bg-blue-100 text-blue-700">
                  <FaUser size={13} />
                </span>
                <span className="hidden max-w-48 sm:block">
                  <span className="block truncate text-xs font-semibold text-slate-800">
                    {displayUserName}
                  </span>
                  <span className="block truncate text-[11px] text-slate-500">
                    {roleName}
                  </span>
                </span>
              </button>
              {userMenuOpen && (
                <div className="absolute right-0 mt-2 w-56 overflow-hidden rounded-xl border border-slate-200 bg-white p-1.5 shadow-xl">
                  <div className="border-b border-slate-100 px-3 py-2">
                    <p className="text-xs font-semibold text-slate-800">{displayUserName}</p>
                    <p className="text-[11px] text-slate-400">{roleName}</p>
                  </div>
                  <button
                    type="button"
                    onClick={() => {
                      setUserMenuOpen(false);
                      router.push("/admin/settings");
                    }}
                    className="flex w-full items-center gap-3 rounded-lg px-3 py-2.5 text-sm text-slate-700 hover:bg-slate-100 cursor-pointer"
                  >
                    <FaGear className="text-slate-400" />
                    Pengaturan
                  </button>
                  <button
                    type="button"
                    onClick={logout}
                    className="flex w-full items-center gap-3 rounded-lg px-3 py-2.5 text-sm text-red-600 hover:bg-red-50 cursor-pointer"
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
