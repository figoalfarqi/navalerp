"use client";

import BackButton from "@/components/mobile/BackButton";
import { FiCheckCircle, FiShield } from "@/components/icons";

type MobileRole = "checker" | "driver";

export default function MobileAboutPage({ role }: { role: MobileRole }) {
  return (
    <main className="mx-auto min-h-dvh max-w-md bg-slate-50 px-4 pb-28 pt-5">
      <BackButton url={`/${role}/account/settings`} />
      <div className="mb-5 mt-4">
        <p className="text-xs font-semibold uppercase tracking-[0.15em] text-blue-600">
          Informasi
        </p>
        <h1 className="mt-1 text-2xl font-bold text-slate-900">
          Tentang aplikasi
        </h1>
      </div>

      <section className="rounded-2xl border border-slate-200 bg-white p-5 shadow-sm">
        <div className="flex h-14 w-14 items-center justify-center rounded-2xl bg-blue-600 text-xs font-black text-white">
          NAVAL
        </div>
        <h2 className="mt-4 text-lg font-bold text-slate-900">
          Naval ERP
        </h2>
        <p className="mt-2 text-sm leading-6 text-slate-600">
          Aplikasi operasional untuk pencatatan perjalanan angkutan dan
          pemantauan data proyek sesuai akses pengguna.
        </p>

        <dl className="mt-5 divide-y divide-slate-100 border-y border-slate-100">
          <div className="flex items-center justify-between py-3 text-sm">
            <dt className="text-slate-500">Versi</dt>
            <dd className="font-semibold text-slate-800">0.1.0</dd>
          </div>
          <div className="flex items-center justify-between py-3 text-sm">
            <dt className="text-slate-500">Mode pengguna</dt>
            <dd className="font-semibold capitalize text-slate-800">{role}</dd>
          </div>
        </dl>

        <div className="mt-5 space-y-3 text-sm text-slate-600">
          <p className="flex items-start gap-2">
            <FiShield className="mt-0.5 shrink-0 text-blue-600" size={18} />
            Data hanya tersedia sesuai akun dan proyek yang ditugaskan.
          </p>
          <p className="flex items-start gap-2">
            <FiCheckCircle
              className="mt-0.5 shrink-0 text-emerald-600"
              size={18}
            />
            Gunakan versi aplikasi terbaru agar alur operasional tetap sesuai.
          </p>
        </div>
      </section>
    </main>
  );
}
