"use client";

import Button from "@/components/form/Button";
import Modal from "@/components/Modal";
import { useAppObjectStore } from "@/hooks/useAppObjectStore";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { BsPerson } from "react-icons/bs";
import { FiChevronRight, FiLock, FiLogOut, FiSettings, FiUser } from "react-icons/fi";

type MobileRole = "checker" | "driver";

type MobileAccountPageProps = {
  role: MobileRole;
  user: {
    name: string;
    username: string;
    photoUrl?: string;
  };
  clearToken: () => Promise<void>;
};

const roleLabels: Record<MobileRole, string> = {
  checker: "Checker",
  driver: "Driver",
};

export default function MobileAccountPage({
  role,
  user,
  clearToken,
}: MobileAccountPageProps) {
  const router = useRouter();
  const { clearAllAppObjectStore } = useAppObjectStore();
  const [isLogoutOpen, setIsLogoutOpen] = useState(false);
  const [isLoggingOut, setIsLoggingOut] = useState(false);

  const menuItems = [
    {
      label: "Detail profil",
      href: `/${role}/account/profile`,
      icon: <FiUser size={20} />,
    },
    {
      label: "Ganti password",
      href: `/${role}/account/password`,
      icon: <FiLock size={20} />,
    },
    {
      label: "Pengaturan",
      href: `/${role}/account/settings`,
      icon: <FiSettings size={20} />,
    },
  ];

  const handleLogout = async () => {
    setIsLoggingOut(true);
    await Promise.all([clearToken(), clearAllAppObjectStore()]);
    router.replace(`/${role}/login`);
  };

  const normalizedPhotoUrl = user.photoUrl
    ? user.photoUrl.startsWith("/")
      ? user.photoUrl
      : `/app_user/${user.photoUrl}`
    : "";

  return (
    <main className="mx-auto min-h-dvh max-w-md bg-slate-50 px-4 pb-28 pt-5">
      <div className="mb-5">
        <p className="text-xs font-semibold uppercase tracking-[0.15em] text-blue-600">
          Akun {roleLabels[role]}
        </p>
        <h1 className="mt-1 text-2xl font-bold text-slate-900">Profil saya</h1>
      </div>

      <section className="rounded-2xl border border-slate-200 bg-white p-5 text-center shadow-sm">
        {normalizedPhotoUrl ? (
          <Image
            src={normalizedPhotoUrl}
            alt={`Foto ${roleLabels[role]}`}
            width={88}
            height={88}
            className="mx-auto h-[88px] w-[88px] rounded-full border-4 border-blue-100 object-cover"
          />
        ) : (
          <div className="mx-auto flex h-[88px] w-[88px] items-center justify-center rounded-full border-4 border-blue-100 bg-blue-50">
            <BsPerson size={48} className="text-blue-600" />
          </div>
        )}
        <h2 className="mt-4 text-lg font-bold text-slate-900">
          {user.name || roleLabels[role]}
        </h2>
        <p className="mt-0.5 text-sm text-slate-500">{user.username || "-"}</p>
      </section>

      <section className="mt-4 overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm">
        {menuItems.map((item, index) => (
          <button
            key={item.href}
            type="button"
            onClick={() => router.push(item.href)}
            className={`flex min-h-14 w-full items-center gap-3 px-4 text-left text-sm font-medium text-slate-700 transition active:bg-slate-50 ${
              index < menuItems.length - 1 ? "border-b border-slate-100" : ""
            }`}
          >
            <span className="text-blue-600">{item.icon}</span>
            <span className="flex-1">{item.label}</span>
            <FiChevronRight size={18} className="text-slate-400" />
          </button>
        ))}
      </section>

      <Button
        id={`${role}_logout`}
        onClick={() => setIsLogoutOpen(true)}
        variant="red-ghost"
        wrapperClassName="mt-4 w-full"
        className="min-h-12 w-full rounded-xl border border-red-100 bg-white"
      >
        <FiLogOut size={20} />
        <span>Keluar</span>
      </Button>

      <Modal
        isOpen={isLogoutOpen}
        title="Keluar dari akun?"
        confirmText="Keluar"
        cancelText="Batal"
        confirmVariant="red-solid"
        loading={isLoggingOut}
        onConfirm={() => void handleLogout()}
        onCancel={() => setIsLogoutOpen(false)}
      >
        {role === "checker"
          ? "Sesi akan diakhiri. Pilihan posisi tetap tersimpan pada perangkat untuk mencegah reset jeda lokasi."
          : "Sesi akun pada perangkat ini akan diakhiri."}
      </Modal>
    </main>
  );
}
