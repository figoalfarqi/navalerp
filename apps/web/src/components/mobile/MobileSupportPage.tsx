"use client";

import BackButton from "@/components/mobile/BackButton";
import { FiAlertCircle, FiBookOpen, FiUserCheck } from "react-icons/fi";

type MobileRole = "checker" | "driver";

const roleSteps: Record<MobileRole, string[]> = {
  checker: [
    "Pilih proyek dan lokasi kerja pada halaman Posisi.",
    "Buka menu Lapor, pilih kendaraan, lalu isi data yang tersedia.",
    "Pastikan status dan foto sudah benar sebelum mengirim laporan.",
  ],
  driver: [
    "Buka menu Perjalanan untuk melihat data angkutan.",
    "Pilih tanggal; tampilan awal otomatis menggunakan tanggal hari ini.",
    "Ketuk salah satu perjalanan untuk melihat rute dan detail status.",
  ],
};

export default function MobileSupportPage({ role }: { role: MobileRole }) {
  return (
    <main className="mx-auto min-h-dvh max-w-md bg-slate-50 px-4 pb-28 pt-5">
      <BackButton url={`/${role}/account/settings`} />
      <div className="mb-5 mt-4">
        <p className="text-xs font-semibold uppercase tracking-[0.15em] text-blue-600">
          Panduan
        </p>
        <h1 className="mt-1 text-2xl font-bold text-slate-900">Bantuan</h1>
      </div>

      <section className="rounded-2xl border border-slate-200 bg-white p-5 shadow-sm">
        <div className="flex items-center gap-3">
          <span className="rounded-xl bg-blue-50 p-3 text-blue-600">
            <FiBookOpen size={22} />
          </span>
          <div>
            <h2 className="font-bold text-slate-900">Panduan singkat</h2>
            <p className="text-xs text-slate-500">
              Alur utama untuk pengguna {role}
            </p>
          </div>
        </div>

        <ol className="mt-5 space-y-4">
          {roleSteps[role].map((step, index) => (
            <li key={step} className="flex gap-3 text-sm leading-6 text-slate-600">
              <span className="flex h-7 w-7 shrink-0 items-center justify-center rounded-full bg-blue-600 text-xs font-bold text-white">
                {index + 1}
              </span>
              <span>{step}</span>
            </li>
          ))}
        </ol>
      </section>

      <section className="mt-4 space-y-3 rounded-2xl border border-amber-200 bg-amber-50 p-5 text-sm text-amber-950">
        <p className="flex items-start gap-2 font-semibold">
          <FiAlertCircle className="mt-0.5 shrink-0" size={18} />
          Data tidak tampil atau gagal dikirim?
        </p>
        <p className="leading-6">
          Periksa koneksi, proyek penugasan, dan tanggal yang dipilih. Jangan
          menghapus data aplikasi untuk melewati pembatasan operasional.
        </p>
      </section>

      <section className="mt-4 flex items-start gap-3 rounded-2xl border border-slate-200 bg-white p-5 text-sm text-slate-600 shadow-sm">
        <FiUserCheck className="mt-0.5 shrink-0 text-blue-600" size={20} />
        <p className="leading-6">
          Jika kendala berlanjut, hubungi administrator atau supervisor proyek
          dengan menyertakan nama proyek, kendaraan, waktu, dan pesan error.
        </p>
      </section>
    </main>
  );
}
