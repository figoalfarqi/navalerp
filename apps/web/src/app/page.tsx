"use client";

import React, { useState } from "react";
import Link from "next/link";
import Image from "next/image";
import {
  CuiGroupIcon,
  WarshipIcon,
  WorkOrderIcon,
  WarehouseIcon,
  PurchaseOrderIcon,
  PersonnelIcon,
  FuelBunkerIcon,
  BaseFacilityIcon,
  ChartOfAccountIcon,
  DashboardIcon,
  CheckIcon,
} from "@/components/icons";
import { useAuth } from "@/context/AuthContext";
import { isAdminRoleId } from "@/components/admin/adminNavigation";

export default function NavalErpLandingPage() {
  const { adminToken, adminPayload, tokenLoaded } = useAuth();
  const isAdmin = tokenLoaded && !!adminToken && isAdminRoleId(adminPayload?.app_role_id ?? 0);
  const portalHref = isAdmin ? "/admin" : "/admin/login";

  const [activeVessel, setActiveVessel] = useState<number>(0);
  const [activeMroStep, setActiveMroStep] = useState<number>(1);
  const [activeProcStep, setActiveProcStep] = useState<number>(1);
  const [demoModalOpen, setDemoModalOpen] = useState<boolean>(false);
  const [demoSubmitted, setDemoSubmitted] = useState<boolean>(false);
  const [demoForm, setDemoForm] = useState({
    name: "",
    rank: "",
    organization: "TNI Angkatan Laut (TNI AL)",
    email: "",
    phone: "",
    notes: "",
  });

  const vessels = [
    {
      name: "KRI Raden Eddy Martadinata-331",
      classType: "SIGMA 10514 Frigate",
      status: "OPERATIONAL",
      readiness: 94,
      nextServiceDays: 27,
      engineHours: "8,421",
      fuelLevel: 76,
      subsystems: [
        { name: "Main Propulsion Diesel (2x MTU 20V 8000)", status: "OK" },
        { name: "Diesel Generators (4x 735 kW)", status: "OK" },
        { name: "Thales SMART-S Mk2 3D Radar", status: "WARNING" },
        { name: "Combat Management System (TACTICOS)", status: "OK" },
      ],
      documents: [
        { title: "Sertifikat Kelaikan Laut", date: "Valid s/d Des 2026" },
        { title: "Klasifikasi Biro Klasifikasi Indonesia (BKI)", date: "Inspeksi Tahunan Terverifikasi" },
        { title: "Surat Tanda Kebangsaan Kapal Perang", date: "Mabesal Terdaftar" },
      ],
    },
    {
      name: "KRI I Gusti Ngurah Rai-332",
      classType: "SIGMA 10514 Frigate",
      status: "OPERATIONAL",
      readiness: 97,
      nextServiceDays: 45,
      engineHours: "6,890",
      fuelLevel: 88,
      subsystems: [
        { name: "Main Propulsion Diesel (2x MTU 20V 8000)", status: "OK" },
        { name: "Diesel Generators (4x 735 kW)", status: "OK" },
        { name: "Thales SMART-S Mk2 3D Radar", status: "OK" },
        { name: "Combat Management System (TACTICOS)", status: "OK" },
      ],
      documents: [
        { title: "Sertifikat Kelaikan Laut", date: "Valid s/d Jan 2027" },
        { title: "Klasifikasi BKI & BV", date: "Lengkap" },
        { title: "Surat Tanda Kebangsaan", date: "Mabesal Terdaftar" },
      ],
    },
    {
      name: "KRI Alugoro-405",
      classType: "Nagapasa-Class Submarine",
      status: "PATROL_READY",
      readiness: 91,
      nextServiceDays: 14,
      engineHours: "5,120",
      fuelLevel: 82,
      subsystems: [
        { name: "MTU 12V 493 Diesel-Electric Engines", status: "OK" },
        { name: "Subsea Battery Energy Bank", status: "OK" },
        { name: "CSU 90 Active/Passive Sonar Suite", status: "OK" },
        { name: "Torpedo Launch Tube Hydraulics", status: "OK" },
      ],
      documents: [
        { title: "Sertifikat Uji Tekan Lambung (Diving Test)", date: "Lolos Uji Kedalaman Maksimal" },
        { title: "Klasifikasi Kapal Selam Tempur", date: "Satsel Koarmada II" },
        { title: "Sertifikasi Sistem Oksigen Darurat", date: "Teruji" },
      ],
    },
  ];

  const mroSteps = [
    { step: "PLAN", title: "Rencana & Penjadwalan", desc: "Perencanaan pemeliharaan berkala (PMS), kalkulasi jam kerja mesin, dan alokasi suku cadang." },
    { step: "WORK ORDER", title: "Penerbitan SPK", desc: "Penerbitan Surat Perintah Kerja (WO) teknisi pangkalan/galangan lengkap dengan instruksi keselamatan." },
    { step: "INSPECTION", title: "Inspeksi Teknis", desc: "Pemeriksaan nondestruktif (NDT), pengukuran keausan komponen, dan pembongkaran terencana." },
    { step: "REPAIR", title: "Perbaikan & Penggantian", desc: "Tindakan perbaikan mesin, overhaul sistem navigasi, atau rekondisi lambung oleh teknisi bersertifikat." },
    { step: "TEST", title: "Uji Fungsi & Sea Trial", desc: "Pengujian darat, uji tekan sistem, dan uji berlayar (Sea Trial) untuk verifikasi kelaikan operasi." },
    { step: "CLOSE", title: "Penutupan & Sertifikasi", desc: "Pembaruan riwayat alutsista, digital sign-off komandan, dan pemutakhiran status kesiapan tempur." },
  ];

  const procSteps = [
    { step: "1. REQUEST", title: "Kebutuhan Material", desc: "Pengajuan bekal & suku cadang oleh Satuan/KRI secara digital." },
    { step: "2. APPROVAL", title: "Otorisasi Komando", desc: "Persetujuan berjenjang dengan tanda tangan digital & verifikasi anggaran." },
    { step: "3. TENDER", title: "Pengadaan / RFQ", desc: "Proses lelang transparan atau penunjukan penyedia alpalhankam terdaftar." },
    { step: "4. SUPPLIER", title: "Kontrak Rekanan", desc: "Evaluasi kapabilitas industri pertahanan dan pemenuhan klausul garansi." },
    { step: "5. PO", title: "Pesanan Pembelian", desc: "Penerbitan Purchase Order otomatis terhubung ke sistem anggaran Mabesal." },
    { step: "6. DELIVERY", title: "Pengiriman Satuan", desc: "Pelacakan transit ekspedisi militer darat, laut, dan udara ke pangkalan." },
    { step: "7. INSPECTION", title: "Uji Terima (FAT/SAT)", desc: "Pemeriksaan mutu fisik, kuantitas, dan kesesuaian spesifikasi teknis militer." },
    { step: "8. PAYMENT", title: "Pembayaran 3-Way", desc: "Pencairan dana setelah verifikasi 3-Way Match (PO + Berita Acara + Invoice)." },
  ];

  const currentVessel = vessels[activeVessel];

  const handleDemoSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setDemoSubmitted(true);
    setTimeout(() => {
      setDemoSubmitted(false);
      setDemoModalOpen(false);
    }, 2800);
  };

  return (
    <div className="min-h-screen bg-[#051428] text-slate-100 font-sans selection:bg-cyan-500 selection:text-white">
      {/* ── Top Navigation Bar ────────────────────────────────────────────── */}
      <header className="sticky top-0 z-50 bg-[#051428]/95 backdrop-blur-md border-b border-blue-900/50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-16 flex items-center justify-between">
          <Link href="/" className="flex items-center gap-3 group">
            <div className="relative w-9 h-9 rounded-lg bg-gradient-to-br from-[#0b2447] to-[#040e1c] border border-cyan-500/40 flex items-center justify-center shadow-md group-hover:border-cyan-400 transition">
              <Image src="/logo.webp" alt="Naval ERP" width={24} height={24} className="object-contain" />
            </div>
            <div className="flex items-center gap-2">
              <span className="text-base font-black tracking-wider text-white font-mono">NAVAL ERP™</span>
              <span className="px-1.5 py-0.5 text-[9px] font-bold bg-cyan-950/80 text-cyan-400 border border-cyan-800/80 rounded font-mono">
                TNI AL
              </span>
            </div>
          </Link>

          {/* Nav menu links */}
          <nav className="hidden lg:flex items-center gap-7 text-xs font-semibold text-slate-300 tracking-wider uppercase">
            <a href="#platform" className="hover:text-cyan-400 transition">Platform</a>
            <a href="#fleet" className="hover:text-cyan-400 transition">Armada & MRO</a>
            <a href="#logistics" className="hover:text-cyan-400 transition">Rantai Pasok</a>
            <a href="#cui" className="hover:text-cyan-400 transition text-cyan-400 flex items-center gap-1.5">
              <span className="w-1.5 h-1.5 rounded-full bg-cyan-400 animate-pulse" />
              CUI Bawah Laut
            </a>
            <a href="#security" className="hover:text-cyan-400 transition">Keamanan</a>
          </nav>

          {/* Action CTAs */}
          <div className="flex items-center gap-3">
            <button
              onClick={() => setDemoModalOpen(true)}
              className="hidden sm:inline-flex items-center px-3.5 py-2 rounded-lg bg-cyan-600 hover:bg-cyan-500 text-white text-xs font-bold tracking-wider uppercase transition shadow-md shadow-cyan-950/50"
            >
              Request Demo
            </button>
            <Link
              href={portalHref}
              className="inline-flex items-center gap-1.5 px-3.5 py-2 rounded-lg bg-[#0b2447] hover:bg-[#12366b] border border-blue-700/60 text-cyan-200 text-xs font-bold tracking-wider uppercase transition"
            >
              Portal Komando →
            </Link>
          </div>
        </div>
      </header>

      {/* ── 1. HOME / HERO SECTION ─────────────────────────────────────────── */}
      <section className="relative pt-12 pb-20 overflow-hidden border-b border-blue-900/40">
        {/* Ambient Glows */}
        <div className="absolute top-1/4 left-1/2 -translate-x-1/2 w-[800px] h-[350px] bg-cyan-500/10 rounded-full blur-3xl pointer-events-none" />
        <div className="absolute top-0 right-0 w-[400px] h-[400px] bg-blue-600/10 rounded-full blur-3xl pointer-events-none" />

        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center relative z-10">
          <div className="inline-flex items-center gap-2 px-3.5 py-1 rounded-full bg-[#081d38] border border-cyan-500/30 text-cyan-300 text-xs font-mono tracking-widest uppercase mb-6 shadow-inner">
            <span className="w-2 h-2 rounded-full bg-cyan-400 animate-ping" />
            TNI ANGKATAN LAUT • MARITIME DEFENSE PLATFORM
          </div>

          <h1 className="text-4xl sm:text-5xl lg:text-6xl font-extrabold tracking-tight text-white leading-tight">
            NAVAL ERP™ <br />
            <span className="bg-gradient-to-r from-cyan-400 via-sky-300 to-blue-400 bg-clip-text text-transparent">
              The Digital Backbone of Maritime Operations
            </span>
          </h1>

          <p className="mt-4 max-w-2xl mx-auto text-sm sm:text-base text-slate-300 leading-relaxed">
            Platform terintegrasi kesiapan armada, rantai pasok pertahanan, dan pengamanan infrastruktur bawah laut nasional.
          </p>

          {/* Hero CTAs */}
          <div className="mt-10 flex flex-col sm:flex-row items-center justify-center gap-4">
            <button
              onClick={() => setDemoModalOpen(true)}
              className="w-full sm:w-auto px-8 py-4 rounded-xl bg-gradient-to-r from-cyan-500 to-blue-600 hover:from-cyan-400 hover:to-blue-500 text-white font-bold text-sm uppercase tracking-wider transition shadow-xl shadow-cyan-950/80 cursor-pointer"
            >
              [ REQUEST DEMO ]
            </button>
            <a
              href="#platform"
              className="w-full sm:w-auto px-8 py-4 rounded-xl bg-[#081d38] hover:bg-[#0b2447] border border-blue-800/80 text-cyan-200 font-bold text-sm uppercase tracking-wider transition"
            >
              [ EXPLORE PLATFORM ]
            </a>
          </div>

          {/* 4 Key Metrics Bar */}
          <div className="mt-16 grid grid-cols-2 lg:grid-cols-4 gap-4 text-left">
            <div className="bg-[#081d38]/80 backdrop-blur border border-blue-900/60 rounded-xl p-6 shadow-lg">
              <div className="text-xs font-mono text-cyan-400 uppercase tracking-wider">Fleet Assets</div>
              <div className="text-3xl lg:text-4xl font-black text-white mt-2 font-mono">1,000+</div>
              <p className="text-xs text-slate-400 mt-2 leading-relaxed">Configurable naval & maritime assets</p>
            </div>

            <div className="bg-[#081d38]/80 backdrop-blur border border-blue-900/60 rounded-xl p-6 shadow-lg">
              <div className="text-xs font-mono text-emerald-400 uppercase tracking-wider">Fleet Readiness</div>
              <div className="text-3xl lg:text-4xl font-black text-emerald-400 mt-2 font-mono">94%</div>
              <p className="text-xs text-slate-400 mt-2 leading-relaxed">Real-time telemetry & readiness monitoring</p>
            </div>

            <div className="bg-[#081d38]/80 backdrop-blur border border-blue-900/60 rounded-xl p-6 shadow-lg">
              <div className="text-xs font-mono text-amber-400 uppercase tracking-wider">Maintenance (MRO)</div>
              <div className="text-3xl lg:text-4xl font-black text-amber-400 mt-2 font-mono">Predictive</div>
              <p className="text-xs text-slate-400 mt-2 leading-relaxed">Planned lifecycle & AI failure forecasting</p>
            </div>

            <div className="bg-[#081d38]/80 backdrop-blur border border-blue-900/60 rounded-xl p-6 shadow-lg">
              <div className="text-xs font-mono text-sky-400 uppercase tracking-wider">Infrastructure</div>
              <div className="text-3xl lg:text-4xl font-black text-sky-400 mt-2 font-mono">CUI Visibility</div>
              <p className="text-xs text-slate-400 mt-2 leading-relaxed">428 Critical Underwater Assets under watch</p>
            </div>
          </div>
        </div>
      </section>

      {/* ── 2. PLATFORM: ONE PLATFORM, MULTIPLE FUNCTIONS ───────────────────── */}
      <section id="platform" className="py-24 border-b border-blue-900/40 bg-[#040e1c]">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center max-w-3xl mx-auto mb-16">
            <div className="text-xs font-mono uppercase tracking-widest text-cyan-400 mb-2">MODULAR ARCHITECTURE</div>
            <h2 className="text-3xl sm:text-5xl font-extrabold text-white tracking-tight">
              One Platform. Multiple Maritime Functions.
            </h2>
            <p className="mt-4 text-slate-300 text-sm sm:text-base leading-relaxed">
              NAVAL ERP integrates the major business, engineering, and operational-support functions of a naval or maritime organization in an interoperable single ecosystem.
            </p>
          </div>

          {/* 10 Modules Interactive Table / Cards */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-4">
            {[
              { icon: WarshipIcon, title: "01. Fleet Management", func: "Vessel & asset lifecycle, engine hours, speed trial & hull integrity.", href: "#fleet" },
              { icon: WorkOrderIcon, title: "02. Maintenance & MRO", func: "Planned & corrective maintenance, dry-docking & shipyard operations.", href: "#maintenance" },
              { icon: WarehouseIcon, title: "03. Logistics & Supply", func: "Inventory, engineering components, ammunition & warehouse network.", href: "#logistics" },
              { icon: PurchaseOrderIcon, title: "04. Procurement", func: "Purchasing, defense tenders, vendor evaluation & contract tracking.", href: "#procurement" },
              { icon: PersonnelIcon, title: "05. Personnel & Crew", func: "Crew qualifications, rank, sea duties & tactical training readiness.", href: "#personnel" },
              { icon: FuelBunkerIcon, title: "06. Fuel & Energy", func: "Bunker inventory, bunker barge logistics & real-time fuel consumption.", href: "#fleet" },
              { icon: BaseFacilityIcon, title: "07. Base & Facilities", func: "Pier berths, naval docks, maintenance hangars & shore utilities.", href: "/admin/data/base_facility" },
              { icon: ChartOfAccountIcon, title: "08. Finance & Budget", func: "Program budgeting, 3-way matching, payments & cost per steaming hour.", href: "#procurement" },
              { icon: CuiGroupIcon, title: "09. Underwater CUI", func: "Submarine cables, pipelines, landing stations & acoustic anomalies.", href: "#cui" },
              { icon: DashboardIcon, title: "10. AI Analytics", func: "Readiness prediction, demand forecasting & executive intelligence.", href: "#ai-command" },
            ].map((m, idx) => (
              <a
                key={idx}
                href={m.href}
                className="bg-[#081d38] hover:bg-[#0b2447] border border-blue-900/60 hover:border-cyan-500/60 rounded-xl p-5 transition flex flex-col justify-between group shadow-lg"
              >
                <div>
                  <div className="w-10 h-10 rounded-lg bg-cyan-950 border border-cyan-800/60 text-cyan-400 flex items-center justify-center mb-3 group-hover:scale-105 transition">
                    <m.icon className="w-5 h-5" />
                  </div>
                  <h3 className="font-bold text-sm text-white group-hover:text-cyan-300 transition font-mono">
                    {m.title}
                  </h3>
                  <p className="mt-2 text-xs text-slate-400 leading-relaxed">{m.func}</p>
                </div>
                <div className="mt-4 pt-3 border-t border-blue-900/40 text-[11px] font-mono text-cyan-400 flex items-center justify-between">
                  <span>Lihat Detail</span>
                  <span>→</span>
                </div>
              </a>
            ))}
          </div>
        </div>
      </section>

      {/* ── 3. FLEET MANAGEMENT & DIGITAL VESSEL PROFILE ───────────────────── */}
      <section id="fleet" className="py-24 border-b border-blue-900/40 bg-[#051428]">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col lg:flex-row lg:items-end justify-between mb-12 gap-6">
            <div>
              <div className="text-xs font-mono uppercase tracking-widest text-cyan-400 mb-2">CENTRALIZED ASSET PROFILE</div>
              <h2 className="text-3xl sm:text-4xl font-extrabold text-white tracking-tight">
                Fleet Readiness at a Glance
              </h2>
              <p className="mt-2 text-slate-300 text-sm max-w-2xl">
                NAVAL ERP provides a centralized digital profile for every warship, submarine, and maritime auxiliary asset.
              </p>
            </div>

            {/* Vessel Selector Tabs */}
            <div className="flex flex-wrap gap-2">
              {vessels.map((v, i) => (
                <button
                  key={i}
                  onClick={() => setActiveVessel(i)}
                  className={`px-3.5 py-2 rounded-lg text-xs font-mono font-bold transition ${
                    activeVessel === i
                      ? "bg-cyan-600 text-white shadow-md shadow-cyan-900/50"
                      : "bg-[#081d38] text-slate-400 hover:text-white border border-blue-900"
                  }`}
                >
                  {v.name.split("-")[0]}
                </button>
              ))}
            </div>
          </div>

          {/* Interactive Digital Vessel Card */}
          <div className="bg-[#081d38] border border-blue-900/70 rounded-2xl p-6 sm:p-8 shadow-2xl relative overflow-hidden">
            <div className="flex flex-col md:flex-row md:items-center justify-between border-b border-blue-900/50 pb-6 gap-4">
              <div>
                <span className="text-[11px] font-mono text-cyan-400 uppercase tracking-wider">KRI / DIGITAL VESSEL PROFILE</span>
                <h3 className="text-2xl font-black text-white mt-1 font-mono">{currentVessel.name}</h3>
                <p className="text-xs text-slate-400 font-mono mt-0.5">{currentVessel.classType} • Satuan Kapal Eskorta Koarmada II</p>
              </div>
              <div className="flex items-center gap-3">
                <span className="px-3 py-1.5 rounded-md text-xs font-mono font-bold bg-emerald-950 text-emerald-300 border border-emerald-700 flex items-center gap-1.5">
                  <span className="w-2 h-2 rounded-full bg-emerald-400 animate-pulse" />
                  STATUS: {currentVessel.status}
                </span>
                <Link
                  href="/admin/data/ship"
                  className="px-3.5 py-1.5 rounded-md text-xs font-mono font-semibold bg-[#0b2447] text-cyan-300 hover:bg-[#12366b] border border-blue-700 transition"
                >
                  Buka di Admin →
                </Link>
              </div>
            </div>

            {/* Vessel Key Telemetry Gauges */}
            <div className="grid grid-cols-2 sm:grid-cols-4 gap-4 my-6">
              <div className="bg-[#051428] p-4 rounded-xl border border-blue-900/40">
                <div className="text-[11px] font-mono text-slate-400 uppercase">READINESS SCORE</div>
                <div className="text-3xl font-black text-emerald-400 font-mono mt-1">{currentVessel.readiness}%</div>
                <div className="text-[10px] text-slate-400 mt-1">Siap Tempur & Patroli</div>
              </div>
              <div className="bg-[#051428] p-4 rounded-xl border border-blue-900/40">
                <div className="text-[11px] font-mono text-slate-400 uppercase">NEXT SERVICE DUE</div>
                <div className="text-3xl font-black text-cyan-300 font-mono mt-1">{currentVessel.nextServiceDays} DAYS</div>
                <div className="text-[10px] text-slate-400 mt-1">Pemeliharaan Terencana</div>
              </div>
              <div className="bg-[#051428] p-4 rounded-xl border border-blue-900/40">
                <div className="text-[11px] font-mono text-slate-400 uppercase">ENGINE HOURS</div>
                <div className="text-3xl font-black text-white font-mono mt-1">{currentVessel.engineHours}</div>
                <div className="text-[10px] text-slate-400 mt-1">Total Jam Operasional</div>
              </div>
              <div className="bg-[#051428] p-4 rounded-xl border border-blue-900/40">
                <div className="text-[11px] font-mono text-slate-400 uppercase">FUEL LEVEL (BUNKER)</div>
                <div className="text-3xl font-black text-amber-400 font-mono mt-1">{currentVessel.fuelLevel}%</div>
                <div className="text-[10px] text-slate-400 mt-1">Ketahanan Jelajah 3,800 NM</div>
              </div>
            </div>

            {/* Subsystems & Documents 2-Column */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6 pt-4 border-t border-blue-900/50">
              <div className="space-y-3">
                <div className="text-xs font-mono font-bold text-cyan-300 uppercase tracking-wider flex items-center gap-2">
                  <WorkOrderIcon className="w-4 h-4" />
                  STATUS SISTEM TEKNIS & SENSOR (MAINTENANCE)
                </div>
                <div className="space-y-2">
                  {currentVessel.subsystems.map((s, idx) => (
                    <div
                      key={idx}
                      className="bg-[#051428] p-3 rounded-lg border border-blue-900/40 text-xs flex items-center justify-between font-mono"
                    >
                      <span className="text-slate-200">{s.name}</span>
                      <span
                        className={`px-2 py-0.5 rounded text-[10px] font-bold ${
                          s.status === "OK"
                            ? "bg-emerald-950 text-emerald-400 border border-emerald-800"
                            : "bg-amber-950 text-amber-400 border border-amber-800"
                        }`}
                      >
                        {s.status === "OK" ? "✓ NORMAL" : "⚠ PERHATIAN"}
                      </span>
                    </div>
                  ))}
                </div>
              </div>

              <div className="space-y-3">
                <div className="text-xs font-mono font-bold text-cyan-300 uppercase tracking-wider flex items-center gap-2">
                  <CheckIcon className="w-4 h-4" />
                  DOKUMEN & SERTIFIKASI MILITER RESMI
                </div>
                <div className="space-y-2">
                  {currentVessel.documents.map((d, idx) => (
                    <div
                      key={idx}
                      className="bg-[#051428] p-3 rounded-lg border border-blue-900/40 text-xs flex items-center justify-between font-mono"
                    >
                      <div>
                        <div className="font-semibold text-white">{d.title}</div>
                        <div className="text-[10px] text-slate-400 mt-0.5">{d.date}</div>
                      </div>
                      <span className="text-emerald-400 font-bold">✓ VERIFIED</span>
                    </div>
                  ))}
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* ── 4. MAINTENANCE & SHIPYARD LIFECYCLE ────────────────────────────── */}
      <section id="maintenance" className="py-24 border-b border-blue-900/40 bg-[#040e1c]">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center max-w-3xl mx-auto mb-16">
            <div className="text-xs font-mono uppercase tracking-widest text-cyan-400 mb-2">MAINTENANCE & SHIPYARD</div>
            <h2 className="text-3xl sm:text-4xl font-extrabold text-white tracking-tight">
              From Reactive Maintenance to Predictive Readiness
            </h2>
            <p className="mt-4 text-slate-300 text-sm leading-relaxed">
              Siklus penuh pemeliharaan kapal perang, dry-docking, inventaris suku cadang, dan analitik probabilitas kegagalan dengan bantuan AI.
            </p>
          </div>

          {/* Interactive Stepper: PLAN → WORK ORDER → INSPECTION → REPAIR → TEST → CLOSE */}
          <div className="grid grid-cols-2 md:grid-cols-6 gap-2 mb-8">
            {mroSteps.map((s, idx) => (
              <button
                key={idx}
                onClick={() => setActiveMroStep(idx)}
                className={`p-3 rounded-xl border text-left transition ${
                  activeMroStep === idx
                    ? "bg-[#0b2447] border-cyan-400 shadow-lg"
                    : "bg-[#081d38] border-blue-900 hover:border-blue-700"
                }`}
              >
                <div className="text-[10px] font-mono text-cyan-400 font-bold">TAHAP 0{idx + 1}</div>
                <div className="font-mono font-black text-sm text-white mt-1">{s.step}</div>
                <div className="text-[11px] text-slate-400 truncate mt-0.5">{s.title}</div>
              </button>
            ))}
          </div>

          {/* Active MRO Step Detail Card */}
          <div className="bg-[#081d38] border border-blue-900/60 rounded-2xl p-6 sm:p-8">
            <div className="flex flex-col md:flex-row md:items-center justify-between gap-4 border-b border-blue-900/40 pb-4">
              <div>
                <span className="text-xs font-mono text-cyan-400">FASE: {mroSteps[activeMroStep].step}</span>
                <h3 className="text-xl font-bold text-white mt-0.5">{mroSteps[activeMroStep].title}</h3>
              </div>
              <Link
                href="/admin/data/work_order"
                className="px-4 py-2 bg-blue-700 hover:bg-blue-600 text-white rounded-lg text-xs font-bold font-mono transition"
              >
                Buka Modul Work Order →
              </Link>
            </div>
            <p className="text-slate-300 text-sm leading-relaxed mt-4">
              {mroSteps[activeMroStep].desc}
            </p>

            {/* AI Maintenance Callout (as specified in plan goal.txt) */}
            <div className="mt-6 p-4 rounded-xl bg-[#051428] border border-cyan-500/40 flex items-start gap-4">
              <div className="w-9 h-9 rounded-lg bg-cyan-950 border border-cyan-800 text-cyan-400 flex items-center justify-center shrink-0">
                🤖
              </div>
              <div>
                <div className="text-xs font-mono font-bold text-cyan-400 uppercase">AI MAINTENANCE DECISION SUPPORT</div>
                <p className="text-xs text-slate-200 mt-1 italic leading-relaxed">
                  &ldquo;Equipment Main Propulsion Diesel MTU 20V memiliki probabilitas 84% membutuhkan inspeksi injektor bahan bakar dalam siklus 45 hari mendatang berdasarkan tren suhu pembuangan gas silinder.&rdquo;
                </p>
                <div className="text-[10px] text-slate-400 mt-1">
                  *Sistem memberikan rekomendasi berbasis data; keputusan akhir eksekusi tetap berada di bawah otorisasi perwira teknik yang berwenang.
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* ── 5. LOGISTICS & SUPPLY CHAIN ───────────────────────────────────── */}
      <section id="logistics" className="py-24 border-b border-blue-900/40 bg-[#051428]">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center max-w-3xl mx-auto mb-16">
            <div className="text-xs font-mono uppercase tracking-widest text-cyan-400 mb-2">SUPPLY CHAIN & INVENTORY</div>
            <h2 className="text-3xl sm:text-4xl font-extrabold text-white tracking-tight">
              Know What You Have. Know Where It Is. Know When You Need It.
            </h2>
            <p className="mt-4 text-slate-300 text-sm leading-relaxed">
              Visibilitas menyeluruh suku cadang alutsista, amunisi, bahan bakar minyak (BBM), dan perbekalan umum dari pabrikan hingga ke lambung kapal perang di laut lepas.
            </p>
          </div>

          {/* Supply Chain Flow Pipeline */}
          <div className="bg-[#081d38] border border-blue-900/60 rounded-2xl p-6 sm:p-8 mb-8">
            <div className="text-xs font-mono text-cyan-400 uppercase mb-4">PIPELINE RANTAI PASOK MATRA LAUT</div>
            <div className="grid grid-cols-2 md:grid-cols-6 gap-3 text-center">
              {[
                { step: "01", title: "SUPPLIER", desc: "Industri Pertahanan" },
                { step: "02", title: "PROCUREMENT", desc: "Disadal Mabesal" },
                { step: "03", title: "CENTRAL WAREHOUSE", desc: "Arsenal / Dopusbektim" },
                { step: "04", title: "REGIONAL DEPOT", desc: "Fasharkan / Bekang" },
                { step: "05", title: "NAVAL BASE", desc: "Lantamal / Lanal" },
                { step: "06", title: "WARSHIP (VESSEL)", desc: "KRI Satuan Tugas" },
              ].map((p, idx) => (
                <div key={idx} className="bg-[#051428] border border-blue-900/50 rounded-xl p-3.5 relative">
                  <div className="text-[10px] font-mono text-slate-500 font-bold">{p.step}</div>
                  <div className="text-xs font-mono font-bold text-white mt-1">{p.title}</div>
                  <div className="text-[10px] text-cyan-300 mt-1">{p.desc}</div>
                </div>
              ))}
            </div>

            {/* AI Supply Forecasting Box */}
            <div className="mt-6 pt-6 border-t border-blue-900/50 grid grid-cols-1 md:grid-cols-2 gap-4 text-xs font-mono">
              <div className="bg-[#0b2447]/60 p-4 rounded-xl border border-blue-900">
                <div className="text-cyan-300 font-bold mb-1">PREDIKSI KEBUTUHAN BEBEKAL (AI FORECAST)</div>
                <p className="text-slate-300 leading-relaxed">
                  Algoritma memperhitungkan intensitas patroli laut ALKI I/II/III untuk memperkirakan titik pemesanan ulang (Reorder Point) suku cadang kritis 60 hari sebelum stock-out.
                </p>
              </div>
              <div className="bg-[#0b2447]/60 p-4 rounded-xl border border-blue-900">
                <div className="text-emerald-300 font-bold mb-1">INVENTARIS MULTI-PANGKALAN</div>
                <p className="text-slate-300 leading-relaxed">
                  Sinkronisasi stok antargudang: Pangkalan Utama Koarmada I (Jakarta), Koarmada II (Surabaya), dan Koarmada III (Sorong) dengan tracking mutasi bekal real-time.
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* ── 6. PROCUREMENT & FINANCE ───────────────────────────────────────── */}
      <section id="procurement" className="py-24 border-b border-blue-900/40 bg-[#040e1c]">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center max-w-3xl mx-auto mb-16">
            <div className="text-xs font-mono uppercase tracking-widest text-cyan-400 mb-2">PROCUREMENT & DIGITAL APPROVAL</div>
            <h2 className="text-3xl sm:text-4xl font-extrabold text-white tracking-tight">
              Transparent Maritime Procurement
            </h2>
            <p className="mt-4 text-slate-300 text-sm leading-relaxed">
              Menghubungkan kebutuhan operasional Satuan langsung ke proses lelang, kontrak, berita acara penerimaan, hingga verifikasi pembayaran dengan stempel digital.
            </p>
          </div>

          {/* Stepper Pipeline */}
          <div className="grid grid-cols-2 sm:grid-cols-4 lg:grid-cols-8 gap-2 mb-8">
            {procSteps.map((p, idx) => (
              <button
                key={idx}
                onClick={() => setActiveProcStep(idx)}
                className={`p-2.5 rounded-xl border text-center transition ${
                  activeProcStep === idx
                    ? "bg-[#0b2447] border-cyan-400 shadow-md"
                    : "bg-[#081d38] border-blue-900 hover:border-blue-700"
                }`}
              >
                <div className="text-[10px] font-mono text-cyan-300 font-bold">{p.step.split(" ")[1]}</div>
                <div className="text-[10px] text-slate-400 mt-0.5 truncate">{p.title}</div>
              </button>
            ))}
          </div>

          <div className="bg-[#081d38] border border-blue-900/60 rounded-2xl p-6 sm:p-8">
            <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-4 border-b border-blue-900/40 pb-4">
              <div>
                <span className="text-xs font-mono text-cyan-400">{procSteps[activeProcStep].step}</span>
                <h3 className="text-xl font-bold text-white mt-0.5">{procSteps[activeProcStep].title}</h3>
              </div>
              <div className="flex items-center gap-2">
                <span className="px-3 py-1 rounded bg-emerald-950 border border-emerald-700 text-emerald-300 font-mono text-xs">
                  ✓ Digital Stamp Ready
                </span>
                <Link
                  href="/admin/data/requisition"
                  className="px-3 py-1 bg-cyan-700 hover:bg-cyan-600 text-white rounded font-mono text-xs transition"
                >
                  Buka Pengadaan →
                </Link>
              </div>
            </div>
            <p className="text-slate-300 text-sm leading-relaxed mt-4">
              {procSteps[activeProcStep].desc}
            </p>

            {/* Core Controls Checklist */}
            <div className="grid grid-cols-2 md:grid-cols-4 gap-3 mt-6 pt-6 border-t border-blue-900/40 text-xs font-mono">
              <div className="flex items-center gap-2 text-slate-300">
                <span className="text-cyan-400 font-bold">✓</span> Digital Approval Workflow
              </div>
              <div className="flex items-center gap-2 text-slate-300">
                <span className="text-cyan-400 font-bold">✓</span> Complete Audit Trail
              </div>
              <div className="flex items-center gap-2 text-slate-300">
                <span className="text-cyan-400 font-bold">✓</span> 3-Way Matching Verifier
              </div>
              <div className="flex items-center gap-2 text-slate-300">
                <span className="text-cyan-400 font-bold">✓</span> Real-Time Budget Control
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* ── 7. PERSONNEL & TRAINING READINESS ───────────────────────────────── */}
      <section id="personnel" className="py-24 border-b border-blue-900/40 bg-[#051428]">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center max-w-3xl mx-auto mb-16">
            <div className="text-xs font-mono uppercase tracking-widest text-cyan-400 mb-2">HUMAN CAPITAL MANAGEMENT</div>
            <h2 className="text-3xl sm:text-4xl font-extrabold text-white tracking-tight">
              People Readiness
            </h2>
            <p className="mt-4 text-slate-300 text-sm leading-relaxed">
              Pengelolaan kualifikasi prajurit, penugasan awak kapal perang (DSP/TOP KRI), status sertifikasi selam/navigasi, dan kesiapan personel tempur.
            </p>
          </div>

          {/* Personnel Readiness Dashboard (as specified in plan goal.txt) */}
          <div className="bg-[#081d38] border border-blue-900/60 rounded-2xl p-6 sm:p-8">
            <div className="flex items-center justify-between border-b border-blue-900/40 pb-4 mb-6">
              <div>
                <span className="text-xs font-mono text-cyan-400 uppercase">PERSONNEL READINESS OVERVIEW</span>
                <div className="text-2xl font-black text-white font-mono mt-1">TOTAL PRAJURIT AKTIF: 8,420</div>
              </div>
              <Link
                href="/admin/data/personnel"
                className="px-4 py-2 rounded-lg bg-[#0b2447] text-cyan-300 hover:bg-[#12366b] border border-blue-700 text-xs font-mono transition"
              >
                Kelola Personel →
              </Link>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
              <div className="bg-[#051428] p-5 rounded-xl border border-blue-900/50">
                <div className="text-xs font-mono text-slate-400">QUALIFIED FOR SEA DUTY</div>
                <div className="text-4xl font-black text-emerald-400 font-mono mt-2">92%</div>
                <div className="w-full bg-slate-800 rounded-full h-2 mt-3 overflow-hidden">
                  <div className="bg-emerald-500 h-full rounded-full" style={{ width: "92%" }} />
                </div>
              </div>

              <div className="bg-[#051428] p-5 rounded-xl border border-blue-900/50">
                <div className="text-xs font-mono text-slate-400">TRAINING CURRENT</div>
                <div className="text-4xl font-black text-cyan-400 font-mono mt-2">88%</div>
                <div className="w-full bg-slate-800 rounded-full h-2 mt-3 overflow-hidden">
                  <div className="bg-cyan-500 h-full rounded-full" style={{ width: "88%" }} />
                </div>
              </div>

              <div className="bg-[#051428] p-5 rounded-xl border border-blue-900/50">
                <div className="text-xs font-mono text-slate-400">CERTIFICATION CURRENT</div>
                <div className="text-4xl font-black text-sky-400 font-mono mt-2">96%</div>
                <div className="w-full bg-slate-800 rounded-full h-2 mt-3 overflow-hidden">
                  <div className="bg-sky-500 h-full rounded-full" style={{ width: "96%" }} />
                </div>
              </div>
            </div>

            <div className="grid grid-cols-2 gap-4 text-xs font-mono bg-[#051428] p-4 rounded-xl border border-blue-900/40">
              <div>
                <span className="text-amber-400 font-bold">143 PERSONEL</span> Sertifikasi Segera Kedaluwarsa (Expiring Soon &lt; 30 Hari)
              </div>
              <div>
                <span className="text-rose-400 font-bold">217 PERSONEL</span> Membutuhkan Pelatihan Penyegaran / Refresher Course
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* ── 8. CUI — CRITICAL UNDERWATER INFRASTRUCTURE ─────────────────────── */}
      <section id="cui" className="py-24 border-b border-blue-900/40 bg-[#040e1c]">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center max-w-3xl mx-auto mb-16">
            <div className="text-xs font-mono uppercase tracking-widest text-cyan-400 mb-2">SIGNATURE MODULE</div>
            <h2 className="text-3xl sm:text-4xl font-extrabold text-white tracking-tight">
              CUI — Critical Underwater Infrastructure
            </h2>
            <p className="mt-4 text-slate-300 text-sm leading-relaxed">
              Inventaris digital nasional aset bawah laut strategis: Kabel Komunikasi Submarine (SKKL), Pipa Gas Subsea, Fasilitas Energi Lepas Pantai, dan Deteksi Anomali AI.
            </p>
          </div>

          {/* CUI Dashboard Stats (428, 391, 17, 6) */}
          <div className="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
            <div className="bg-[#081d38] border border-blue-900/60 rounded-xl p-5 shadow-lg">
              <div className="text-xs font-mono text-slate-400 uppercase">CRITICAL ASSETS</div>
              <div className="text-3xl font-black text-white font-mono mt-1">428</div>
              <div className="text-[11px] text-cyan-300 mt-1">Aset Bawah Laut Nasional</div>
            </div>
            <div className="bg-[#081d38] border border-blue-900/60 rounded-xl p-5 shadow-lg">
              <div className="text-xs font-mono text-slate-400 uppercase">ACTIVE MONITORING</div>
              <div className="text-3xl font-black text-emerald-400 font-mono mt-1">391</div>
              <div className="text-[11px] text-emerald-300 mt-1">Sensor Telemetri Aktif</div>
            </div>
            <div className="bg-[#081d38] border border-blue-900/60 rounded-xl p-5 shadow-lg">
              <div className="text-xs font-mono text-slate-400 uppercase">INSPECTION REQUIRED</div>
              <div className="text-3xl font-black text-amber-400 font-mono mt-1">17</div>
              <div className="text-[11px] text-amber-300 mt-1">Jadwal Penyelaman / ROV</div>
            </div>
            <div className="bg-[#081d38] border border-blue-900/60 rounded-xl p-5 shadow-lg">
              <div className="text-xs font-mono text-slate-400 uppercase">ALERTS (ANOMALI)</div>
              <div className="text-3xl font-black text-rose-400 font-mono mt-1">6</div>
              <div className="text-[11px] text-rose-300 mt-1">Korelasi Sensor AIS</div>
            </div>
          </div>

          {/* CUI Map and AI Anomaly Preview Box */}
          <div className="bg-[#081d38] border border-blue-900/60 rounded-2xl p-6 sm:p-8 flex flex-col md:flex-row gap-6 items-center justify-between">
            <div className="space-y-3 max-w-xl">
              <div className="flex items-center gap-2 text-cyan-400 text-xs font-mono">
                <CuiGroupIcon className="w-5 h-5" />
                <span>Pusat Komando CUI Interaktif</span>
              </div>
              <h3 className="text-xl font-bold text-white">
                Pemantauan Geografis & Deteksi Dini Ancaman Sabotase Bawah Laut
              </h3>
              <p className="text-xs text-slate-300 leading-relaxed">
                Platform mengkorelasikan data sensor akustik dasar laut, sensor medan magnet, pelacakan AIS kapal asing, dan riwayat perawatan untuk memprioritaskan investigasi KRI.
              </p>
            </div>
            <div className="flex flex-col sm:flex-row gap-3 w-full md:w-auto">
              <Link
                href="/admin/cui"
                className="px-6 py-3.5 bg-cyan-600 hover:bg-cyan-500 text-white font-mono font-bold text-xs uppercase tracking-wider rounded-xl transition text-center shadow-lg shadow-cyan-900/40"
              >
                Buka Pusat Komando CUI →
              </Link>
            </div>
          </div>
        </div>
      </section>

      {/* ── 9. AI COMMAND & EXECUTIVE ANALYTICS ────────────────────────────── */}
      <section id="ai-command" className="py-24 border-b border-blue-900/40 bg-[#051428]">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center max-w-3xl mx-auto mb-16">
            <div className="text-xs font-mono uppercase tracking-widest text-cyan-400 mb-2">EXECUTIVE INTELLIGENCE</div>
            <h2 className="text-3xl sm:text-4xl font-extrabold text-white tracking-tight">
              From Data to Decisions
            </h2>
            <p className="mt-4 text-slate-300 text-sm leading-relaxed">
              Lapisan Artificial Intelligence terintegrasi di atas seluruh modul ERP, menyajikan ringkasan eksekutif bagi para pimpinan komando tinggi TNI AL.
            </p>
          </div>

          {/* Executive Dashboard (as specified in plan goal.txt: 91%, 87%, 94%, 96%, 89%, 76%) */}
          <div className="bg-[#081d38] border border-blue-900/60 rounded-2xl p-6 sm:p-8">
            <div className="text-xs font-mono text-cyan-400 uppercase mb-6 flex items-center justify-between">
              <span>EXECUTIVE DASHBOARD • 6 KEY DEFENSE INDICATORS</span>
              <span className="text-slate-400">Live Synthesis</span>
            </div>

            <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4">
              {[
                { title: "Fleet Readiness", val: "91%", color: "text-emerald-400", bar: "w-[91%] bg-emerald-500" },
                { title: "Maintenance", val: "87%", color: "text-cyan-400", bar: "w-[87%] bg-cyan-500" },
                { title: "Logistics", val: "94%", color: "text-sky-400", bar: "w-[94%] bg-sky-500" },
                { title: "Personnel", val: "96%", color: "text-emerald-400", bar: "w-[96%] bg-emerald-500" },
                { title: "Infrastructure", val: "89%", color: "text-amber-400", bar: "w-[89%] bg-amber-500" },
                { title: "Budget Util.", val: "76%", color: "text-blue-400", bar: "w-[76%] bg-blue-500" },
              ].map((ind, idx) => (
                <div key={idx} className="bg-[#051428] p-4 rounded-xl border border-blue-900/40">
                  <div className="text-[11px] font-mono text-slate-400">{ind.title}</div>
                  <div className={`text-2xl font-black font-mono mt-1 ${ind.color}`}>{ind.val}</div>
                  <div className="w-full bg-slate-800 rounded-full h-1.5 mt-2 overflow-hidden">
                    <div className={`h-full rounded-full ${ind.bar}`} />
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* ── 10. SECURITY & ARCHITECTURE ────────────────────────────────────── */}
      <section id="security" className="py-24 border-b border-blue-900/40 bg-[#040e1c]">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center max-w-3xl mx-auto mb-16">
            <div className="text-xs font-mono uppercase tracking-widest text-cyan-400 mb-2">SECURE BY DESIGN</div>
            <h2 className="text-3xl sm:text-4xl font-extrabold text-white tracking-tight">
              Military-Grade Security & Architecture
            </h2>
            <p className="mt-4 text-slate-300 text-sm leading-relaxed">
              Arsitektur berstandar pertahanan nasional dengan segmentasi jaringan ketat, enkripsi end-to-end, dan audit trail tak terhapus.
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
            {[
              { label: "Identity", title: "Role-Based Access Control", desc: "Hierarki hak akses berjenjang (Mabesal, Pangkoarmada, Komandan KRI, Logistik, Teknisi)." },
              { label: "Data", title: "Encryption at Rest & Transit", desc: "Enkripsi data rahasia dengan standar AES-256 dan TLS 1.3 terkunci." },
              { label: "Audit", title: "Complete Activity Logging", desc: "Seluruh aksi transaksi dicatat ke dalam tabel sys_audit_logs secara permanen." },
              { label: "Network", title: "Controlled Connectivity", desc: "Mendukung operasi On-Premise, Air-Gapped Intranet TNI AL, dan Sinkronisasi Satelit." },
              { label: "Availability", title: "Redundancy & DR", desc: "Arsitektur multi-cluster aktif-pasif menjamin ketersediaan 99.99% di masa krisis." },
              { label: "Governance", title: "Classification Policies", desc: "Otorisasi berlapis sesuai klasifikasi dokumen (Biasa, Rahasia, Sangat Rahasia)." },
            ].map((s, idx) => (
              <div key={idx} className="bg-[#081d38] border border-blue-900/60 rounded-xl p-5 space-y-2">
                <span className="text-[10px] font-mono font-bold text-cyan-400 uppercase tracking-widest">{s.label}</span>
                <h3 className="font-bold text-sm text-white font-mono">{s.title}</h3>
                <p className="text-xs text-slate-400 leading-relaxed">{s.desc}</p>
              </div>
            ))}
          </div>

          {/* Architecture Visual Diagram */}
          <div className="bg-[#081d38] border border-blue-900/60 rounded-2xl p-6 sm:p-8 font-mono text-center text-xs space-y-4">
            <div className="text-slate-400 uppercase tracking-wider font-bold">SYSTEM ARCHITECTURE TOPOLOGY</div>
            <div className="inline-block p-2 rounded bg-cyan-950 text-cyan-300 border border-cyan-800">
              USERS: WEB PORTAL • MOBILE APP (KRI TABLET)
            </div>
            <div className="text-slate-500">↓ (JWT Authenticated TLS 1.3)</div>
            <div className="inline-block p-2 rounded bg-blue-950 text-blue-300 border border-blue-800">
              API GATEWAY (Go REST Microservices)
            </div>
            <div className="text-slate-500">↓</div>
            <div className="flex flex-wrap items-center justify-center gap-2">
              <span className="p-2 bg-[#051428] rounded border border-blue-900 text-slate-200">ERP CORE (MRO / LOG / PROC / FIN)</span>
              <span className="p-2 bg-[#051428] rounded border border-blue-900 text-slate-200">GIS & MARITIME MAP ENGINE</span>
              <span className="p-2 bg-[#051428] rounded border border-blue-900 text-slate-200">CUI SENSOR TELEMETRY</span>
            </div>
            <div className="text-slate-500">↓</div>
            <div className="inline-block p-2 rounded bg-emerald-950 text-emerald-300 border border-emerald-800">
              POSTGRESQL 16 HIGH AVAILABILITY DATA PLATFORM + AI ANALYTICS ENGINE
            </div>
          </div>
        </div>
      </section>

      {/* ── FOOTER ─────────────────────────────────────────────────────────── */}
      <footer className="py-12 bg-[#030b14] border-t border-blue-900/50 text-xs text-slate-400">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
          <div className="flex flex-col md:flex-row items-center justify-between gap-6">
            <div className="flex items-center gap-3">
              <Image src="/logo.webp" alt="Naval ERP" width={36} height={36} className="object-contain" />
              <div>
                <span className="font-bold font-mono text-white text-sm">NAVAL ERP™</span>
                <p className="text-[11px] text-cyan-400">TNI Angkatan Laut • Jalesveva Jayamahe</p>
              </div>
            </div>
            <div className="flex items-center gap-6 font-mono text-xs text-slate-400">
              <a href="#platform" className="hover:text-white transition">Platform</a>
              <a href="#cui" className="hover:text-white transition">CUI</a>
              <a href="#security" className="hover:text-white transition">Security</a>
              <Link href="/admin/login" className="text-cyan-400 hover:text-white transition">Petugas Login</Link>
            </div>
          </div>

          <div className="pt-6 border-t border-blue-900/30 text-center text-slate-500 text-[11px] leading-relaxed">
            NAVAL ERP™ is an integrated Maritime Enterprise Management Platform combining ERP, Fleet Asset Management, Maintenance, Logistics, GIS, CUI Management and AI Analytics in a single secure digital ecosystem.
            <br />
            © {new Date().getFullYear()} Markas Besar Tentara Nasional Indonesia Angkatan Laut (MABESAL). Hak Cipta Dilindungi Undang-Undang.
          </div>
        </div>
      </footer>

      {/* ── INTERACTIVE REQUEST DEMO MODAL ─────────────────────────────────── */}
      {demoModalOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm">
          <div className="bg-[#081d38] border border-blue-800/80 rounded-2xl max-w-lg w-full p-6 text-slate-100 shadow-2xl relative">
            <button
              onClick={() => setDemoModalOpen(false)}
              className="absolute top-4 right-4 text-slate-400 hover:text-white text-lg font-mono"
            >
              ✕
            </button>

            <div className="mb-4">
              <span className="text-[10px] font-mono text-cyan-400 uppercase tracking-wider">MARITIME PLATFORM DEMO</span>
              <h3 className="text-xl font-bold text-white mt-0.5">Permintaan Presentasi & Uji Coba</h3>
              <p className="text-xs text-slate-400 mt-1">
                Jadwalkan demonstrasi teknis untuk satuan operasional atau instansi maritim Anda.
              </p>
            </div>

            {demoSubmitted ? (
              <div className="py-8 text-center space-y-3">
                <div className="w-12 h-12 rounded-full bg-emerald-950 border border-emerald-500 text-emerald-400 flex items-center justify-center mx-auto text-xl font-bold">
                  ✓
                </div>
                <h4 className="text-base font-bold text-white font-mono">Permintaan Berhasil Dikirim!</h4>
                <p className="text-xs text-slate-300">
                  Tim teknis kami akan menghubungi perwira penghubung Anda dalam 1x24 jam kerja.
                </p>
              </div>
            ) : (
              <form onSubmit={handleDemoSubmit} className="space-y-3.5 text-xs font-mono">
                <div>
                  <label className="block text-slate-300 mb-1">Nama Lengkap & Pangkat *</label>
                  <input
                    type="text"
                    required
                    value={demoForm.name}
                    onChange={(e) => setDemoForm({ ...demoForm, name: e.target.value })}
                    placeholder="Contoh: Kolonel Laut (T) ..."
                    className="w-full p-2.5 rounded-lg bg-[#051428] border border-blue-900 text-white focus:outline-none focus:border-cyan-500"
                  />
                </div>

                <div>
                  <label className="block text-slate-300 mb-1">Satuan / Instansi *</label>
                  <select
                    value={demoForm.organization}
                    onChange={(e) => setDemoForm({ ...demoForm, organization: e.target.value })}
                    className="w-full p-2.5 rounded-lg bg-[#051428] border border-blue-900 text-white focus:outline-none focus:border-cyan-500"
                  >
                    <option value="TNI AL - Mabesal">TNI AL - Mabesal / Kotama</option>
                    <option value="TNI AL - Koarmada I">TNI AL - Koarmada I</option>
                    <option value="TNI AL - Koarmada II">TNI AL - Koarmada II</option>
                    <option value="TNI AL - Koarmada III">TNI AL - Koarmada III</option>
                    <option value="Bakamla RI">Bakamla RI / Coast Guard</option>
                    <option value="Kemenhub / Hubla">Kementerian Perhubungan (Ditjen Hubla)</option>
                    <option value="Galangan / BUMN Pertahanan">PT PAL Indonesia / BUMN Pertahanan</option>
                    <option value="Operator Energi Lepas Pantai">Operator Infrastruktur Maritim / Migas</option>
                  </select>
                </div>

                <div className="grid grid-cols-2 gap-2">
                  <div>
                    <label className="block text-slate-300 mb-1">Email Dinas / Kantor *</label>
                    <input
                      type="email"
                      required
                      value={demoForm.email}
                      onChange={(e) => setDemoForm({ ...demoForm, email: e.target.value })}
                      placeholder="nama@tnial.mil.id"
                      className="w-full p-2.5 rounded-lg bg-[#051428] border border-blue-900 text-white focus:outline-none focus:border-cyan-500"
                    />
                  </div>
                  <div>
                    <label className="block text-slate-300 mb-1">No. Kontak / WA *</label>
                    <input
                      type="tel"
                      required
                      value={demoForm.phone}
                      onChange={(e) => setDemoForm({ ...demoForm, phone: e.target.value })}
                      placeholder="+62 8..."
                      className="w-full p-2.5 rounded-lg bg-[#051428] border border-blue-900 text-white focus:outline-none focus:border-cyan-500"
                    />
                  </div>
                </div>

                <div>
                  <label className="block text-slate-300 mb-1">Fokus Kebutuhan / Catatan</label>
                  <textarea
                    rows={2}
                    value={demoForm.notes}
                    onChange={(e) => setDemoForm({ ...demoForm, notes: e.target.value })}
                    placeholder="Contoh: Integrasi MRO KRI dan pemantauan kabel bawah laut Natuna..."
                    className="w-full p-2.5 rounded-lg bg-[#051428] border border-blue-900 text-white focus:outline-none focus:border-cyan-500"
                  />
                </div>

                <div className="flex justify-end gap-2 pt-2 border-t border-blue-900/50">
                  <button
                    type="button"
                    onClick={() => setDemoModalOpen(false)}
                    className="px-4 py-2 bg-[#051428] text-slate-400 hover:text-white rounded-lg transition"
                  >
                    Batal
                  </button>
                  <button
                    type="submit"
                    className="px-5 py-2 bg-cyan-600 hover:bg-cyan-500 text-white font-bold rounded-lg transition shadow-md shadow-cyan-950"
                  >
                    Kirim Permintaan Demo
                  </button>
                </div>
              </form>
            )}
          </div>
        </div>
      )}
    </div>
  );
}
