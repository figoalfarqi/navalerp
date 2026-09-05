"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { useAuth } from "@/context/AuthContext";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import {
  FaShip,
  FaCompass,
  FaWrench,
  FaBoxesStacked,
  FaUsers,
  FaShieldHalved,
  FaArrowRight,
  FaPlus,
  FaSitemap,
  FaMoneyBillWave,
  FaTruckFast,
  FaFileContract,
  FaArrowRotateRight,
} from "react-icons/fa6";

interface MetricState {
  ships: number;
  missions: number;
  workOrders: number;
  materials: number;
  personnel: number;
  readinessReports: number;
}

export default function AdminDashboardPage() {
  const { adminPayload } = useAuth();
  const { getAPI } = useFetchAPI();
  const [metrics, setMetrics] = useState<MetricState>({
    ships: 0,
    missions: 0,
    workOrders: 0,
    materials: 0,
    personnel: 0,
    readinessReports: 0,
  });
  const [loading, setLoading] = useState(true);

  const fetchDashboardData = async () => {
    setLoading(true);
    try {
      const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
      const opts = { authToken: "admin" as const };

      const [shipsRes, missionsRes, woRes, matRes, persRes, repRes] = await Promise.allSettled([
        getAPI<any>(`${baseUrl}/admin/ship?limit=1`, opts),
        getAPI<any>(`${baseUrl}/admin/mission?limit=1`, opts),
        getAPI<any>(`${baseUrl}/admin/work_order?limit=1`, opts),
        getAPI<any>(`${baseUrl}/admin/material?limit=1`, opts),
        getAPI<any>(`${baseUrl}/admin/personnel?limit=1`, opts),
        getAPI<any>(`${baseUrl}/admin/readiness_report?limit=1`, opts),
      ]);

      setMetrics({
        ships: shipsRes.status === "fulfilled" && shipsRes.value?.data ? (shipsRes.value.data.total ?? 0) : 0,
        missions: missionsRes.status === "fulfilled" && missionsRes.value?.data ? (missionsRes.value.data.total ?? 0) : 0,
        workOrders: woRes.status === "fulfilled" && woRes.value?.data ? (woRes.value.data.total ?? 0) : 0,
        materials: matRes.status === "fulfilled" && matRes.value?.data ? (matRes.value.data.total ?? 0) : 0,
        personnel: persRes.status === "fulfilled" && persRes.value?.data ? (persRes.value.data.total ?? 0) : 0,
        readinessReports: repRes.status === "fulfilled" && repRes.value?.data ? (repRes.value.data.total ?? 0) : 0,
      });
    } catch (err) {
      console.error("Failed to load dashboard metrics", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchDashboardData();
  }, []);

  const kpis = [
    {
      title: "Armada KRI",
      count: metrics.ships,
      desc: "Kapal perang terdaftar",
      href: "/admin/data/ship",
      icon: FaShip,
      color: "from-blue-600 to-cyan-600",
      textColor: "text-cyan-400",
    },
    {
      title: "Operasi & Misi",
      count: metrics.missions,
      desc: "Misi pelayaran aktif",
      href: "/admin/data/mission",
      icon: FaCompass,
      color: "from-teal-600 to-emerald-600",
      textColor: "text-teal-400",
    },
    {
      title: "Pemeliharaan (MRO)",
      count: metrics.workOrders,
      desc: "Work order perawatan & docking",
      href: "/admin/data/work_order",
      icon: FaWrench,
      color: "from-amber-600 to-orange-600",
      textColor: "text-amber-400",
    },
    {
      title: "Logistik & Suku Cadang",
      count: metrics.materials,
      desc: "Katalog material terdaftar",
      href: "/admin/data/material",
      icon: FaBoxesStacked,
      color: "from-indigo-600 to-blue-600",
      textColor: "text-indigo-400",
    },
    {
      title: "Personel Prajurit",
      count: metrics.personnel,
      desc: "Prajurit & perwira TNI AL",
      href: "/admin/data/personnel",
      icon: FaUsers,
      color: "from-sky-600 to-blue-700",
      textColor: "text-sky-400",
    },
    {
      title: "Kesiapan Tempur",
      count: metrics.readinessReports,
      desc: "Laporan kesiapan alutsista",
      href: "/admin/data/readiness_report",
      icon: FaShieldHalved,
      color: "from-rose-600 to-red-600",
      textColor: "text-rose-400",
    },
  ];

  const modules = [
    {
      name: "Core & Master Data",
      code: "MOD-01",
      desc: "Satuan organisasi, pangkalan TNI AL, fasharkan, theater operasi, kepangkatan & korps.",
      href: "/admin/data/org_unit",
      icon: FaSitemap,
      badge: "7 Tabel",
    },
    {
      name: "Armada & Platform KRI",
      code: "MOD-02",
      desc: "Kelas kapal, armada KRI, sistem kapal, alutsista senjata, sensor & mesin penggerak.",
      href: "/admin/data/ship",
      icon: FaShip,
      badge: "8 Tabel",
    },
    {
      name: "Rantai Pasok & Gudang",
      code: "MOD-03",
      desc: "Katalog material alpalhankam, serial instance, saldo gudang, mutasi & penyesuaian stok.",
      href: "/admin/data/material",
      icon: FaBoxesStacked,
      badge: "8 Tabel",
    },
    {
      name: "Pemeliharaan & MRO",
      code: "MOD-04",
      desc: "Work order perawatan, jadwal pemeliharaan preventif, laporan kerusakan, docking record.",
      href: "/admin/data/work_order",
      icon: FaWrench,
      badge: "8 Tabel",
    },
    {
      name: "Operasi & Kesiapan",
      code: "MOD-05",
      desc: "Misi pelayaran tempur, penugasan kru KRI, log harian navigasi, bunker BBM, peringatan dini.",
      href: "/admin/data/mission",
      icon: FaCompass,
      badge: "9 Tabel",
    },
    {
      name: "Pengadaan & Vendor",
      code: "MOD-06",
      desc: "Pengadaan alutsista, tender pertahanan, kontrak mitra, purchase order, penerimaan barang.",
      href: "/admin/data/purchase_order",
      icon: FaFileContract,
      badge: "8 Tabel",
    },
    {
      name: "Anggaran & Keuangan",
      code: "MOD-07",
      desc: "Program DIPA anggaran, komitmen dana, bagan akun standar (COA), jurnal umum, faktur & TCO alutsista.",
      href: "/admin/data/budget_program",
      icon: FaMoneyBillWave,
      badge: "9 Tabel",
    },
    {
      name: "Logistik & Pergerakan",
      code: "MOD-08",
      desc: "Satuan angkut darat/laut, rute pelayaran & pergerakan logistik pangkalan.",
      href: "/admin/data/transport_unit",
      icon: FaTruckFast,
      badge: "5 Tabel",
    },
    {
      name: "Personel Militer",
      code: "MOD-09",
      desc: "Data prajurit matra laut, kualifikasi/brevet, riwayat penugasan operasi, rekam medis kelautan.",
      href: "/admin/data/personnel",
      icon: FaUsers,
      badge: "4 Tabel",
    },
    {
      name: "Sistem & Keamanan",
      code: "MOD-10",
      desc: "Akun pengguna sistem, audit log aktivitas transaksi, manajemen dokumen & arsip digital.",
      href: "/admin/data/sys_user",
      icon: FaShieldHalved,
      badge: "5 Tabel",
    },
  ];

  return (
    <div className="space-y-6 p-4 sm:p-6 lg:p-8">
      {/* Header Banner */}
      <section className="relative overflow-hidden rounded-2xl bg-gradient-to-r from-[#062444] via-[#09426f] to-[#0b8fa5] p-6 text-white shadow-xl sm:p-8">
        <div className="relative z-10 flex flex-col justify-between gap-4 md:flex-row md:items-center">
          <div>
            <div className="inline-flex items-center gap-2 rounded-full bg-cyan-950/60 px-3 py-1 text-xs font-semibold tracking-wider text-cyan-300 uppercase backdrop-blur border border-cyan-400/20">
              <span className="h-2 w-2 rounded-full bg-emerald-400 animate-pulse" />
              SISTEM KOMANDO NAVAL ERP AKTIF
            </div>
            <h1 className="mt-3 text-2xl font-extrabold sm:text-3xl lg:text-4xl tracking-tight">
              Pusat Komando & Operasional Alutsista
            </h1>
            <p className="mt-2 max-w-2xl text-sm text-cyan-100/90 leading-relaxed">
              Monitoring kesiapan tempur armada KRI, integrasi logistik perbekalan, pemeliharaan alpalhankam,
              dan manajemen personel pertahanan maritim TNI AL.
            </p>
          </div>

          <div className="flex flex-col sm:flex-row items-start sm:items-center gap-3">
            <button
              onClick={fetchDashboardData}
              className="inline-flex items-center gap-2 rounded-xl bg-white/10 hover:bg-white/20 border border-white/20 px-4 py-2.5 text-xs sm:text-sm font-medium backdrop-blur transition"
            >
              <FaArrowRotateRight className={loading ? "animate-spin" : ""} />
              Refresh Data
            </button>
            <div className="rounded-xl bg-black/20 border border-white/10 px-4 py-2 text-xs backdrop-blur">
              <span className="block text-cyan-200 font-semibold">User: {adminPayload?.app_user_name || "Super Admin"}</span>
              <span className="text-white/70">Role: {adminPayload?.username || "admin"}</span>
            </div>
          </div>
        </div>
      </section>

      {/* KPI Metrics Cards */}
      <section>
        <div className="flex items-center justify-between mb-4">
          <h2 className="text-lg font-bold text-slate-800 tracking-tight">
            Ringkasan Alutsista & Operasi
          </h2>
          <span className="text-xs text-slate-500 font-medium">Data Terintegrasi Database Real-time</span>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
          {kpis.map((kpi) => {
            const Icon = kpi.icon;
            return (
              <Link
                key={kpi.title}
                href={kpi.href}
                className="group relative overflow-hidden rounded-2xl border border-slate-200 bg-white p-5 shadow-sm transition hover:shadow-md hover:border-cyan-500/50"
              >
                <div className="flex items-start justify-between">
                  <div>
                    <p className="text-xs font-semibold uppercase tracking-wider text-slate-500">
                      {kpi.title}
                    </p>
                    <div className="mt-2 flex items-baseline gap-2">
                      <span className="text-3xl font-extrabold text-slate-900 tracking-tight">
                        {loading ? "..." : kpi.count}
                      </span>
                    </div>
                    <p className="mt-1 text-xs text-slate-500">{kpi.desc}</p>
                  </div>
                  <div className={`rounded-xl bg-gradient-to-br ${kpi.color} p-3.5 text-white shadow-md shadow-cyan-900/20 group-hover:scale-105 transition-transform`}>
                    <Icon size={20} />
                  </div>
                </div>
                <div className="mt-4 flex items-center gap-1.5 text-xs font-semibold text-cyan-700 group-hover:text-cyan-800">
                  <span>Lihat Selengkapnya</span>
                  <FaArrowRight size={11} className="transition-transform group-hover:translate-x-1" />
                </div>
              </Link>
            );
          })}
        </div>
      </section>

      {/* Quick Actions */}
      <section className="rounded-2xl border border-slate-200 bg-white p-5 sm:p-6 shadow-sm">
        <h2 className="text-base font-bold text-slate-800 mb-4 tracking-tight">
          Aksi Cepat Operasional
        </h2>
        <div className="flex flex-wrap gap-3">
          <Link
            href="/admin/data/ship/create"
            className="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl bg-slate-900 hover:bg-slate-800 text-white text-xs font-semibold shadow transition"
          >
            <FaPlus size={12} /> Tambah Armada KRI
          </Link>
          <Link
            href="/admin/data/work_order/create"
            className="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl bg-amber-600 hover:bg-amber-500 text-white text-xs font-semibold shadow transition"
          >
            <FaPlus size={12} /> Buat Work Order MRO
          </Link>
          <Link
            href="/admin/data/mission/create"
            className="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl bg-teal-600 hover:bg-teal-500 text-white text-xs font-semibold shadow transition"
          >
            <FaPlus size={12} /> Rencana Misi Baru
          </Link>
          <Link
            href="/admin/data/personnel/create"
            className="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl bg-blue-600 hover:bg-blue-500 text-white text-xs font-semibold shadow transition"
          >
            <FaPlus size={12} /> Input Prajurit Baru
          </Link>
          <Link
            href="/admin/data/material/create"
            className="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl bg-cyan-700 hover:bg-cyan-600 text-white text-xs font-semibold shadow transition"
          >
            <FaPlus size={12} /> Registrasi Material
          </Link>
        </div>
      </section>

      {/* 10 Naval ERP Modules Grid */}
      <section>
        <div className="flex items-center justify-between mb-4">
          <div>
            <h2 className="text-lg font-bold text-slate-800 tracking-tight">
              10 Modul Naval ERP Terintegrasi
            </h2>
            <p className="text-xs text-slate-500">Akses langsung ke setiap subsistem data arsitektur Naval ERP</p>
          </div>
          <span className="text-xs font-semibold px-2.5 py-1 rounded-full bg-cyan-100 text-cyan-800 border border-cyan-200">
            Total 72 Tabel Aktif
          </span>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-4">
          {modules.map((mod) => {
            const Icon = mod.icon;
            return (
              <Link
                key={mod.code}
                href={mod.href}
                className="group flex flex-col justify-between rounded-2xl border border-slate-200 bg-white p-5 shadow-sm transition hover:shadow-md hover:border-cyan-500/50 hover:bg-slate-50/50"
              >
                <div>
                  <div className="flex items-center justify-between mb-3">
                    <span className="text-[11px] font-bold tracking-wider uppercase px-2.5 py-0.5 rounded-full bg-slate-100 text-slate-600 border border-slate-200">
                      {mod.code}
                    </span>
                    <span className="text-[11px] font-semibold text-cyan-700 bg-cyan-50 px-2 py-0.5 rounded-full border border-cyan-100">
                      {mod.badge}
                    </span>
                  </div>
                  <div className="flex items-center gap-3 mb-2">
                    <div className="w-9 h-9 rounded-xl bg-cyan-100 text-cyan-700 flex items-center justify-center shrink-0 group-hover:bg-cyan-600 group-hover:text-white transition-colors">
                      <Icon size={18} />
                    </div>
                    <h3 className="text-sm font-bold text-slate-800 group-hover:text-cyan-700 transition-colors">
                      {mod.name}
                    </h3>
                  </div>
                  <p className="text-xs text-slate-500 leading-relaxed line-clamp-2">
                    {mod.desc}
                  </p>
                </div>
                <div className="mt-4 pt-3 border-t border-slate-100 flex items-center justify-between text-xs font-semibold text-slate-600 group-hover:text-cyan-700">
                  <span>Buka Modul</span>
                  <FaArrowRight size={11} className="transition-transform group-hover:translate-x-1" />
                </div>
              </Link>
            );
          })}
        </div>
      </section>
    </div>
  );
}
