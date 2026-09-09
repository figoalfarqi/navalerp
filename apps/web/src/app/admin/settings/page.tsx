"use client";

import { useRouter } from "next/navigation";
import {
  FaChevronRight,
  FaSliders,
  FaUsers,
  FaUsersGear,
} from "@/components/icons";
import { useAuth } from "@/context/AuthContext";

export default function AdminSettingsPage() {
  const router = useRouter();
  const { adminPayload } = useAuth();
  const roleId = adminPayload?.app_role_id ?? 0;
  const cards = [
    {
      title: roleId === 6 ? "Pengguna Admin" : "Semua pengguna",
      description:
        roleId === 6
          ? "Kelola akun admin operasional, checker, dan driver secara terpisah."
          : "Kelola akun berdasarkan role dan kewenangannya.",
      href: roleId === 6 ? "/admin/data/admin" : "/admin/data/app_user",
      icon: FaUsers,
      visible: [4, 5, 6].includes(roleId),
    },
    {
      title: "Role aplikasi",
      description: "Atur role Owner, IT Dev, Super Admin, Admin, Checker, dan Driver.",
      href: "/admin/data/app_role",
      icon: FaUsersGear,
      visible: [4, 5].includes(roleId),
    },
    {
      title: "Konfigurasi aplikasi",
      description: "Kelola nilai konfigurasi global yang tersimpan di app_setting.",
      href: "/admin/data/app_setting",
      icon: FaSliders,
      visible: [4, 5].includes(roleId),
    },
  ].filter((card) => card.visible);

  return (
    <div className="space-y-5">
      <section className="rounded-2xl border border-slate-200 bg-white p-5 shadow-sm sm:p-7">
        <p className="text-xs font-semibold uppercase tracking-[0.18em] text-blue-600">
          Administration
        </p>
        <h1 className="mt-1 text-2xl font-semibold text-slate-900">
          Pengaturan
        </h1>
        <p className="mt-2 max-w-2xl text-sm text-slate-500">
          Pengaturan ditampilkan sesuai kewenangan role yang sedang masuk.
          Data konfigurasi, role, dan pengguna tetap dikelola pada CRUD
          tabelnya masing-masing.
        </p>
      </section>

      {cards.length > 0 && (
        <section className="grid gap-3 md:grid-cols-2 xl:grid-cols-3">
          {cards.map((card) => {
            const Icon = card.icon;
            return (
              <button
                key={card.href}
                type="button"
                onClick={() => router.push(card.href)}
                className="group flex items-center gap-4 rounded-2xl border border-slate-200 bg-white p-5 text-left shadow-sm transition hover:-translate-y-0.5 hover:border-blue-200 hover:shadow-md cursor-pointer"
              >
                <span className="grid h-12 w-12 shrink-0 place-items-center rounded-xl bg-blue-50 text-blue-700">
                  <Icon size={20} />
                </span>
                <span className="min-w-0 flex-1">
                  <span className="block font-semibold text-slate-900">
                    {card.title}
                  </span>
                  <span className="mt-1 block text-xs leading-relaxed text-slate-500">
                    {card.description}
                  </span>
                </span>
                <FaChevronRight className="text-slate-300 transition group-hover:translate-x-1 group-hover:text-blue-500" />
              </button>
            );
          })}
        </section>
      )}
    </div>
  );
}
