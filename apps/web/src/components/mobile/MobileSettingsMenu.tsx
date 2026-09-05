"use client";

import BackButton from "@/components/mobile/BackButton";
import { useRouter } from "next/navigation";
import { FiChevronRight, FiHelpCircle, FiInfo } from "react-icons/fi";

type MobileRole = "checker" | "driver";

export default function MobileSettingsMenu({ role }: { role: MobileRole }) {
  const router = useRouter();
  const items = [
    {
      label: "Bantuan",
      description: "Panduan penggunaan sesuai peran",
      href: `/${role}/account/settings/support`,
      icon: <FiHelpCircle size={21} />,
    },
    {
      label: "Tentang aplikasi",
      description: "Informasi aplikasi dan versi",
      href: `/${role}/account/settings/about`,
      icon: <FiInfo size={21} />,
    },
  ];

  return (
    <main className="mx-auto min-h-dvh max-w-md bg-slate-50 px-4 pb-28 pt-5">
      <BackButton url={`/${role}/account`} />
      <div className="mb-5 mt-4">
        <p className="text-xs font-semibold uppercase tracking-[0.15em] text-blue-600">
          Akun
        </p>
        <h1 className="mt-1 text-2xl font-bold text-slate-900">Pengaturan</h1>
      </div>

      <section className="overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm">
        {items.map((item, index) => (
          <button
            key={item.href}
            type="button"
            onClick={() => router.push(item.href)}
            className={`flex min-h-[72px] w-full items-center gap-3 px-4 text-left transition active:bg-slate-50 ${
              index < items.length - 1 ? "border-b border-slate-100" : ""
            }`}
          >
            <span className="text-blue-600">{item.icon}</span>
            <span className="min-w-0 flex-1">
              <span className="block text-sm font-semibold text-slate-800">
                {item.label}
              </span>
              <span className="mt-0.5 block text-xs text-slate-500">
                {item.description}
              </span>
            </span>
            <FiChevronRight size={18} className="text-slate-400" />
          </button>
        ))}
      </section>
    </main>
  );
}
