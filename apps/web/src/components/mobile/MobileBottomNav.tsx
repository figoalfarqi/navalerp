"use client";

import { usePathname, useRouter } from "next/navigation";
import { ReactNode } from "react";

export type MobileNavItem = {
  href: string;
  label: string;
  icon: ReactNode;
  activePrefixes?: string[];
};

export default function MobileBottomNav({
  items,
}: {
  items: MobileNavItem[];
}) {
  const pathname = usePathname();
  const router = useRouter();

  return (
    <nav
      aria-label="Navigasi utama"
      className="fixed bottom-0 left-1/2 z-40 flex w-full max-w-md -translate-x-1/2 items-start justify-around border-t border-slate-200 bg-white/95 px-2 pt-2 shadow-[0_-8px_30px_rgba(15,23,42,0.08)] backdrop-blur"
      style={{ paddingBottom: "calc(0.5rem + env(safe-area-inset-bottom))" }}
    >
      {items.map((item) => {
        const prefixes = item.activePrefixes ?? [item.href];
        const isActive = prefixes.some((prefix) =>
          pathname.startsWith(prefix),
        );

        return (
          <button
            key={item.href}
            type="button"
            onClick={() => router.push(item.href)}
            aria-current={isActive ? "page" : undefined}
            className={`flex min-h-12 min-w-16 flex-col items-center justify-center gap-0.5 rounded-xl px-2 text-[11px] font-medium transition ${
              isActive
                ? "bg-blue-50 text-blue-600"
                : "text-slate-400 active:bg-slate-100"
            }`}
          >
            {item.icon}
            <span>{item.label}</span>
          </button>
        );
      })}
    </nav>
  );
}
