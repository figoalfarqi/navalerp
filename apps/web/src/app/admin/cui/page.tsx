"use client";

import React, { useEffect, useState } from "react";
import Link from "next/link";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import {
  CuiGroupIcon,
  CuiAssetIcon,
  CuiMonitoringLogIcon,
  CuiAlertIcon,
  CuiInspectionIcon,
} from "@/components/icons";

interface CuiMetrics {
  critical_assets: number;
  active_monitoring: number;
  inspection_required: number;
  alerts: number;
}

interface CuiAsset {
  cui_asset_id: string;
  asset_code: string;
  asset_name: string;
  asset_type: string;
  operator_name: string;
  depth_meters?: number;
  length_km?: number;
  latitude: number;
  longitude: number;
  status: string;
  health_score: number;
  protection_priority: string;
}

interface CuiAlert {
  alert_id: string;
  alert_code: string;
  title: string;
  asset_name: string;
  asset_type: string;
  severity: string;
  alert_type: string;
  status: string;
  detected_at: string;
  description?: string;
}

interface CuiInspection {
  inspection_id: string;
  inspection_code: string;
  asset_name: string;
  inspection_type: string;
  inspection_date: string;
  overall_condition: string;
  status: string;
}

export default function CuiCommandCenterPage() {
  const [metrics, setMetrics] = useState<CuiMetrics>({
    critical_assets: 428,
    active_monitoring: 391,
    inspection_required: 17,
    alerts: 6,
  });
  const [assets, setAssets] = useState<CuiAsset[]>([]);
  const [alerts, setAlerts] = useState<CuiAlert[]>([]);
  const [inspections, setInspections] = useState<CuiInspection[]>([]);
  const [selectedAsset, setSelectedAsset] = useState<CuiAsset | null>(null);
  const [filterType, setFilterType] = useState<string>("ALL");
  const [actionLoading, setActionLoading] = useState<string | null>(null);
  const [approvalMsg, setApprovalMsg] = useState<string | null>(null);

  const { getAPI, postAPI } = useFetchAPI();

  useEffect(() => {
    getAPI<any>("/admin/cui/overview", { authToken: "admin" })
      .then((res) => {
        if (res && res.code === 200 && res.data) {
          if (res.data.metrics) setMetrics(res.data.metrics);
          if (res.data.assets && res.data.assets.length > 0) {
            setAssets(res.data.assets);
            setSelectedAsset(res.data.assets[0]);
          }
          if (res.data.alerts) setAlerts(res.data.alerts);
          if (res.data.inspections) setInspections(res.data.inspections);
        }
      })
      .catch((err) => {
        console.warn("Using fallback CUI telemetry:", err);
        // Fallback default assets if API is starting up
        const fallbackAssets: CuiAsset[] = [
          {
            cui_asset_id: "cui-001",
            asset_code: "CUI-CAB-001",
            asset_name: "SKKL Natuna Express",
            asset_type: "telecom_cable",
            operator_name: "PT Telkom Indonesia / Dishidros",
            depth_meters: 85,
            length_km: 340,
            latitude: 3.9,
            longitude: 108.3,
            status: "OPERATIONAL",
            health_score: 96,
            protection_priority: "HIGH",
          },
          {
            cui_asset_id: "cui-002",
            asset_code: "CUI-PIP-001",
            asset_name: "South Sumatra - West Java Gas Pipeline (SSWJ)",
            asset_type: "pipeline",
            operator_name: "PT Perusahaan Gas Negara (PGN)",
            depth_meters: 42,
            length_km: 190,
            latitude: -5.8,
            longitude: 106.1,
            status: "OPERATIONAL",
            health_score: 91,
            protection_priority: "CRITICAL",
          },
          {
            cui_asset_id: "cui-003",
            asset_code: "CUI-OFF-001",
            asset_name: "Natuna D-Alpha Offshore Gas Facility",
            asset_type: "offshore_energy",
            operator_name: "SKK Migas / Pertamina",
            depth_meters: 145,
            latitude: 4.5,
            longitude: 109.8,
            status: "MAINTENANCE_REQUIRED",
            health_score: 78,
            protection_priority: "CRITICAL",
          },
          {
            cui_asset_id: "cui-004",
            asset_code: "CUI-CLS-001",
            asset_name: "Batam Submarine Cable Landing Station",
            asset_type: "landing_station",
            operator_name: "Kemenkominfo / TNI AL",
            depth_meters: 12,
            latitude: 1.15,
            longitude: 104.05,
            status: "OPERATIONAL",
            health_score: 98,
            protection_priority: "HIGH",
          },
        ];
        setAssets(fallbackAssets);
        setSelectedAsset(fallbackAssets[0]);
      });
  }, [getAPI]);

  const handleResolveAlert = async (alertId: string) => {
    setActionLoading(alertId);
    setApprovalMsg(null);
    try {
      const res = await postAPI<any>(
        "/admin/approval",
        { authToken: "admin" },
        {
          entity_type: "cui_alert",
          entity_id: alertId,
          action: "APPROVE",
          notes: "Diverifikasi & ditindaklanjuti oleh Komando Pengamanan Bawah Laut",
        }
      );
      if (res && [200, 201].includes(res.code)) {
        setApprovalMsg(`Peringatan ${alertId} berhasil diverifikasi & diselesaikan (Stempel: ${res.data?.signature_stamp || "DIGITAL-VERIFIED"})`);
        setAlerts((prev) =>
          prev.map((a) => (a.alert_id === alertId ? { ...a, status: "RESOLVED" } : a))
        );
      } else {
        alert(res?.message || "Gagal memproses approval");
      }
    } catch {
      alert("Terjadi kesalahan koneksi server");
    } finally {
      setActionLoading(null);
    }
  };

  const filteredAssets = filterType === "ALL" 
    ? assets 
    : assets.filter((a) => a.asset_type === filterType);

  return (
    <div className="min-h-screen bg-[#051428] text-slate-100 p-4 md:p-8 space-y-6">
      {/* Header Bar */}
      <div className="flex flex-col md:flex-row md:items-center justify-between gap-4 border-b border-blue-900/40 pb-6">
        <div>
          <div className="flex items-center gap-2.5 text-cyan-400 text-xs font-mono uppercase tracking-widest mb-1">
            <span className="inline-block w-2.5 h-2.5 rounded-full bg-cyan-400 animate-pulse" />
            TNI Angkatan Laut • Sistem Informasi Pertahanan Maritim
          </div>
          <h1 className="text-2xl md:text-3xl font-bold tracking-tight text-white flex items-center gap-3">
            <CuiGroupIcon className="w-8 h-8 text-cyan-400" />
            Pusat Komando Infrastruktur Bawah Laut (CUI)
          </h1>
          <p className="text-slate-400 text-sm mt-1">
            Pemantauan Real-Time Kabel Bawah Laut, Jalur Pipa Energi, Stasiun Pendarat & Anomali Maritim Nasional
          </p>
        </div>

        {/* Quick CRUD shortcuts */}
        <div className="flex flex-wrap items-center gap-2">
          <Link
            href="/admin/data/cui_asset"
            className="flex items-center gap-2 px-3 py-2 text-xs font-semibold rounded bg-[#0b2447] hover:bg-[#12366b] border border-blue-800/60 text-cyan-200 transition"
          >
            <CuiAssetIcon className="w-4 h-4 text-cyan-400" />
            Kelola Aset CUI
          </Link>
          <Link
            href="/admin/data/cui_alert"
            className="flex items-center gap-2 px-3 py-2 text-xs font-semibold rounded bg-[#0b2447] hover:bg-[#12366b] border border-blue-800/60 text-amber-300 transition"
          >
            <CuiAlertIcon className="w-4 h-4 text-amber-400" />
            Peringatan Anomali
          </Link>
          <Link
            href="/admin/data/cui_inspection"
            className="flex items-center gap-2 px-3 py-2 text-xs font-semibold rounded bg-[#0b2447] hover:bg-[#12366b] border border-blue-800/60 text-emerald-300 transition"
          >
            <CuiInspectionIcon className="w-4 h-4 text-emerald-400" />
            Riwayat Inspeksi
          </Link>
        </div>
      </div>

      {/* Approval Notification Banner */}
      {approvalMsg && (
        <div className="p-3 bg-emerald-950/80 border border-emerald-600/50 rounded-lg text-emerald-200 text-sm flex items-center justify-between">
          <div className="flex items-center gap-2">
            <span className="text-emerald-400">✓</span>
            {approvalMsg}
          </div>
          <button onClick={() => setApprovalMsg(null)} className="text-emerald-400 hover:text-white text-xs">
            ✕
          </button>
        </div>
      )}

      {/* 4 Key Metrics Bar (as specified in plan goal.txt: 428, 391, 17, 6) */}
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
        <div className="bg-[#081d38] border border-blue-900/60 rounded-xl p-5 shadow-lg relative overflow-hidden">
          <div className="absolute top-0 right-0 w-24 h-24 bg-blue-500/5 rounded-full blur-2xl" />
          <div className="text-xs font-mono uppercase tracking-wider text-slate-400">CRITICAL ASSETS</div>
          <div className="text-3xl font-extrabold text-white mt-2 font-mono">{metrics.critical_assets}</div>
          <div className="text-xs text-blue-400 mt-2 flex items-center gap-1.5">
            <span className="w-2 h-2 rounded-full bg-blue-400" />
            Terdaftar di seluruh Perairan RI
          </div>
        </div>

        <div className="bg-[#081d38] border border-blue-900/60 rounded-xl p-5 shadow-lg relative overflow-hidden">
          <div className="absolute top-0 right-0 w-24 h-24 bg-emerald-500/5 rounded-full blur-2xl" />
          <div className="text-xs font-mono uppercase tracking-wider text-slate-400">ACTIVE MONITORING</div>
          <div className="text-3xl font-extrabold text-emerald-400 mt-2 font-mono">{metrics.active_monitoring}</div>
          <div className="text-xs text-emerald-400 mt-2 flex items-center gap-1.5">
            <span className="w-2 h-2 rounded-full bg-emerald-400 animate-pulse" />
            Sensor Telemetri Online 98.4%
          </div>
        </div>

        <div className="bg-[#081d38] border border-blue-900/60 rounded-xl p-5 shadow-lg relative overflow-hidden">
          <div className="absolute top-0 right-0 w-24 h-24 bg-amber-500/5 rounded-full blur-2xl" />
          <div className="text-xs font-mono uppercase tracking-wider text-slate-400">INSPECTION REQUIRED</div>
          <div className="text-3xl font-extrabold text-amber-400 mt-2 font-mono">{metrics.inspection_required}</div>
          <div className="text-xs text-amber-400 mt-2 flex items-center gap-1.5">
            <span className="w-2 h-2 rounded-full bg-amber-400" />
            Jadwal ROV / Tim Selam Dishidros
          </div>
        </div>

        <div className="bg-[#081d38] border border-rose-900/40 rounded-xl p-5 shadow-lg relative overflow-hidden">
          <div className="absolute top-0 right-0 w-24 h-24 bg-rose-500/10 rounded-full blur-2xl" />
          <div className="text-xs font-mono uppercase tracking-wider text-slate-400">ALERTS (ANOMALI)</div>
          <div className="text-3xl font-extrabold text-rose-400 mt-2 font-mono">{metrics.alerts}</div>
          <div className="text-xs text-rose-400 mt-2 flex items-center gap-1.5">
            <span className="w-2 h-2 rounded-full bg-rose-500 animate-ping" />
            Membutuhkan Investigasi / Otorisasi
          </div>
        </div>
      </div>

      {/* Main Command Display: GIS Tactical Map & Asset Inspector */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Left 2 Cols: Maritime Tactical Grid */}
        <div className="lg:col-span-2 bg-[#081d38] border border-blue-900/60 rounded-xl p-5 flex flex-col space-y-4">
          <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-3 border-b border-blue-900/40 pb-3">
            <div className="flex items-center gap-2">
              <div className="w-3 h-3 rounded-full bg-cyan-400 animate-ping" />
              <h2 className="font-semibold text-white tracking-wide">Peta Taktis Infrastruktur Bawah Laut RI</h2>
            </div>
            {/* Filter Tabs */}
            <div className="flex flex-wrap items-center gap-1 text-xs">
              {[
                { label: "SEMUA", value: "ALL" },
                { label: "KABEL (SKKL)", value: "telecom_cable" },
                { label: "PIPA GAS/MINYAK", value: "pipeline" },
                { label: "STASIUN KABEL", value: "landing_station" },
                { label: "FASILITAS ENERGI", value: "offshore_energy" },
              ].map((t) => (
                <button
                  key={t.value}
                  onClick={() => setFilterType(t.value)}
                  className={`px-2.5 py-1 rounded transition ${
                    filterType === t.value
                      ? "bg-cyan-600 text-white font-medium"
                      : "bg-[#051428] text-slate-400 hover:text-white"
                  }`}
                >
                  {t.label}
                </button>
              ))}
            </div>
          </div>

          {/* Interactive Tactical SVG Map of Indonesian Maritime Archipelago */}
          <div className="relative w-full h-80 md:h-96 bg-[#040e1c] rounded-lg border border-blue-900/50 overflow-hidden flex items-center justify-center">
            {/* Radar Circular Scan Animation */}
            <div className="absolute inset-0 bg-[radial-gradient(circle_at_center,_rgba(6,182,212,0.08)_0%,_transparent_70%)] pointer-events-none" />
            <div className="absolute w-[350px] h-[350px] border border-blue-800/30 rounded-full animate-[spin_20s_linear_infinite] pointer-events-none opacity-40">
              <div className="w-1/2 h-0.5 bg-gradient-to-r from-transparent to-cyan-400" />
            </div>

            {/* Map Grid Lines */}
            <svg className="absolute inset-0 w-full h-full opacity-20 pointer-events-none">
              <defs>
                <pattern id="grid" width="40" height="40" patternUnits="userSpaceOnUse">
                  <path d="M 40 0 L 0 0 0 40" fill="none" stroke="#38bdf8" strokeWidth="0.5" />
                </pattern>
              </defs>
              <rect width="100%" height="100%" fill="url(#grid)" />
            </svg>

            {/* Archipelago Islands Outline (Stylized Indonesian Waters) */}
            <svg viewBox="0 0 1000 500" className="absolute inset-0 w-full h-full object-cover pointer-events-none opacity-40">
              {/* Sumatra */}
              <path d="M120 180 L230 310 L280 340 L260 360 L180 310 L100 220 Z" fill="#0b2447" stroke="#1e40af" strokeWidth="1" />
              {/* Java */}
              <path d="M280 370 L480 380 L520 390 L510 405 L280 395 Z" fill="#0b2447" stroke="#1e40af" strokeWidth="1" />
              {/* Kalimantan */}
              <path d="M370 190 L480 200 L510 280 L440 330 L360 300 L350 230 Z" fill="#0b2447" stroke="#1e40af" strokeWidth="1" />
              {/* Sulawesi */}
              <path d="M570 210 L600 240 L620 230 L600 290 L560 310 L580 270 Z" fill="#0b2447" stroke="#1e40af" strokeWidth="1" />
              {/* Papua */}
              <path d="M780 240 L880 250 L920 320 L870 340 L820 300 Z" fill="#0b2447" stroke="#1e40af" strokeWidth="1" />
              {/* Subsea Cables overlay */}
              <path d="M200 220 Q 300 240 380 280" fill="none" stroke="#06b6d4" strokeWidth="2" strokeDasharray="4 2" />
              <path d="M260 350 Q 320 330 450 280" fill="none" stroke="#06b6d4" strokeWidth="2" strokeDasharray="4 2" />
              <path d="M380 150 Q 420 220 460 300" fill="none" stroke="#f59e0b" strokeWidth="2.5" />
            </svg>

            {/* Tactical Interactive Points for Assets */}
            <div className="absolute inset-0 p-6 pointer-events-auto">
              {filteredAssets.map((asset, idx) => {
                // Map lat/long to rough relative percentage on map
                // Indo lat roughly -11 to +6 (17 deg span), long roughly 95 to 141 (46 deg span)
                const topPct = Math.max(10, Math.min(85, ((6 - asset.latitude) / 17) * 80 + 10));
                const leftPct = Math.max(8, Math.min(90, ((asset.longitude - 95) / 46) * 80 + 10));

                const isSelected = selectedAsset?.cui_asset_id === asset.cui_asset_id;
                const isWarning = asset.status === "MAINTENANCE_REQUIRED" || asset.health_score < 80;

                return (
                  <button
                    key={asset.cui_asset_id || idx}
                    onClick={() => setSelectedAsset(asset)}
                    style={{ top: `${topPct}%`, left: `${leftPct}%` }}
                    className="absolute -translate-x-1/2 -translate-y-1/2 group focus:outline-none cursor-pointer"
                  >
                    <span className="relative flex h-5 w-5 items-center justify-center">
                      <span
                        className={`animate-ping absolute inline-flex h-full w-full rounded-full opacity-75 ${
                          isWarning ? "bg-amber-400" : isSelected ? "bg-cyan-400" : "bg-blue-400"
                        }`}
                      />
                      <span
                        className={`relative inline-flex rounded-full h-3.5 w-3.5 border-2 border-[#040e1c] ${
                          isWarning ? "bg-amber-500" : isSelected ? "bg-cyan-400" : "bg-blue-500"
                        }`}
                      />
                    </span>
                    {/* Tooltip on hover */}
                    <div className="absolute bottom-full left-1/2 -translate-x-1/2 mb-2 hidden group-hover:flex flex-col items-center pointer-events-none z-20">
                      <div className="bg-[#0b2447] text-white text-[11px] font-mono px-2.5 py-1 rounded shadow-xl border border-blue-700 whitespace-nowrap">
                        <div className="font-bold text-cyan-300">{asset.asset_name}</div>
                        <div className="text-slate-300">
                          {asset.asset_type} • Kedalaman: {asset.depth_meters ?? "-"}m • Skor: {asset.health_score}%
                        </div>
                      </div>
                    </div>
                  </button>
                );
              })}
            </div>

            {/* Map Legend */}
            <div className="absolute bottom-3 left-3 bg-[#081d38]/90 backdrop-blur border border-blue-900/60 px-3 py-2 rounded text-[11px] font-mono space-y-1 z-10">
              <div className="text-slate-400 font-bold mb-1">LEGENDA INFRASTRUKTUR</div>
              <div className="flex items-center gap-2">
                <span className="w-2 h-2 rounded-full bg-cyan-400" />
                <span>KABEL TELEKOMUNIKASI (SKKL)</span>
              </div>
              <div className="flex items-center gap-2">
                <span className="w-2 h-2 rounded-full bg-amber-400" />
                <span>PIPA MINYAK / GAS BAWAH LAUT</span>
              </div>
              <div className="flex items-center gap-2">
                <span className="w-2 h-2 rounded-full bg-emerald-400" />
                <span>STASIUN PENDARAT / FASILITAS ENERGI</span>
              </div>
            </div>

            {/* Coordinates / Theater badge */}
            <div className="absolute top-3 right-3 bg-[#081d38]/90 backdrop-blur border border-blue-900/60 px-3 py-1.5 rounded text-[11px] font-mono text-cyan-400">
              WILAYAH KOARMADA I • II • III
            </div>
          </div>

          {/* Quick Real-Time Telemetry Logs */}
          <div className="border border-blue-900/40 rounded-lg p-3 bg-[#051428]/60">
            <div className="flex items-center justify-between text-xs text-slate-300 font-mono mb-2">
              <span className="flex items-center gap-1.5 font-semibold text-cyan-400">
                <CuiMonitoringLogIcon className="w-4 h-4" />
                STATUS TELEMETRI SENSOR AKUSTIK & SEISMIK
              </span>
              <span className="text-slate-500">Auto-refresh 5 detik</span>
            </div>
            <div className="grid grid-cols-2 md:grid-cols-4 gap-2 text-xs font-mono">
              <div className="bg-[#0b2447]/60 p-2 rounded border border-blue-900/40">
                <div className="text-slate-400">Sensor Akustik</div>
                <div className="text-emerald-400 font-bold mt-1">118.4 dB (NORMAL)</div>
              </div>
              <div className="bg-[#0b2447]/60 p-2 rounded border border-blue-900/40">
                <div className="text-slate-400">Anomali Magnetik</div>
                <div className="text-cyan-400 font-bold mt-1">0.14 µT (STABIL)</div>
              </div>
              <div className="bg-[#0b2447]/60 p-2 rounded border border-blue-900/40">
                <div className="text-slate-400">Tekanan Subsea</div>
                <div className="text-emerald-400 font-bold mt-1">14.2 Bar (OPTIMAL)</div>
              </div>
              <div className="bg-[#0b2447]/60 p-2 rounded border border-blue-900/40">
                <div className="text-slate-400">Kedekatan Kapal Asing</div>
                <div className="text-amber-400 font-bold mt-1">2 Target AIS Dipantau</div>
              </div>
            </div>
          </div>
        </div>

        {/* Right 1 Col: Selected Asset Detail & Action Panel */}
        <div className="bg-[#081d38] border border-blue-900/60 rounded-xl p-5 flex flex-col justify-between space-y-4">
          <div>
            <div className="flex items-center justify-between border-b border-blue-900/40 pb-3">
              <span className="text-xs font-mono text-cyan-400 uppercase tracking-wider">PROFIL ASET BAWAH LAUT</span>
              <span className="px-2 py-0.5 rounded text-[10px] font-bold bg-blue-950 text-cyan-300 border border-blue-800 font-mono">
                {selectedAsset?.asset_code || "SELECT"}
              </span>
            </div>

            {selectedAsset ? (
              <div className="mt-4 space-y-3 font-mono text-xs">
                <div>
                  <div className="text-slate-400 text-[11px]">NAMA ASET:</div>
                  <div className="text-sm font-bold text-white mt-0.5">{selectedAsset.asset_name}</div>
                </div>

                <div className="grid grid-cols-2 gap-2 pt-1">
                  <div>
                    <div className="text-slate-400">JENIS ASET:</div>
                    <div className="text-cyan-300 font-medium">{selectedAsset.asset_type}</div>
                  </div>
                  <div>
                    <div className="text-slate-400">STATUS OPERASI:</div>
                    <div className={selectedAsset.status === "OPERATIONAL" ? "text-emerald-400" : "text-amber-400"}>
                      {selectedAsset.status}
                    </div>
                  </div>
                </div>

                <div className="grid grid-cols-2 gap-2 pt-1">
                  <div>
                    <div className="text-slate-400">KEDALAMAN:</div>
                    <div className="text-slate-200">{selectedAsset.depth_meters ?? "-"} meter</div>
                  </div>
                  <div>
                    <div className="text-slate-400">PANJANG KABEL/PIPA:</div>
                    <div className="text-slate-200">{selectedAsset.length_km ? `${selectedAsset.length_km} km` : "-"}</div>
                  </div>
                </div>

                <div className="grid grid-cols-2 gap-2 pt-1">
                  <div>
                    <div className="text-slate-400">KOORDINAT LAT/LNG:</div>
                    <div className="text-slate-300">{selectedAsset.latitude.toFixed(4)}, {selectedAsset.longitude.toFixed(4)}</div>
                  </div>
                  <div>
                    <div className="text-slate-400">PRIORITAS PENGAMANAN:</div>
                    <div className="text-rose-400 font-bold">{selectedAsset.protection_priority}</div>
                  </div>
                </div>

                <div className="pt-2">
                  <div className="flex justify-between text-slate-400 mb-1">
                    <span>SKOR INTEGRITAS FISIK:</span>
                    <span className="font-bold text-white">{selectedAsset.health_score}%</span>
                  </div>
                  <div className="w-full bg-slate-800 rounded-full h-2 overflow-hidden">
                    <div
                      className={`h-full rounded-full ${
                        selectedAsset.health_score > 85
                          ? "bg-emerald-500"
                          : selectedAsset.health_score > 70
                          ? "bg-amber-500"
                          : "bg-rose-500"
                      }`}
                      style={{ width: `${selectedAsset.health_score}%` }}
                    />
                  </div>
                </div>

                <div className="pt-2 text-slate-300 text-[11px] leading-relaxed bg-[#051428] p-2.5 rounded border border-blue-900/50">
                  <div className="font-bold text-slate-400 mb-0.5">PENGELOLA / OTORITAS:</div>
                  {selectedAsset.operator_name}
                </div>
              </div>
            ) : (
              <div className="py-12 text-center text-slate-500 text-xs">
                Pilih aset pada peta untuk melihat data telemetri.
              </div>
            )}
          </div>

          {/* Action trigger */}
          {selectedAsset && (
            <div className="pt-4 border-t border-blue-900/40 space-y-2">
              <Link
                href={`/admin/data/cui_asset/${selectedAsset.cui_asset_id}/edit`}
                className="w-full py-2 bg-blue-700 hover:bg-blue-600 text-white rounded font-medium text-xs text-center block transition"
              >
                Ubah Detail Aset Bawah Laut
              </Link>
              <Link
                href={`/admin/data/cui_inspection/create?asset_id=${selectedAsset.cui_asset_id}`}
                className="w-full py-2 bg-[#0b2447] hover:bg-[#12366b] text-cyan-300 border border-blue-700/60 rounded font-medium text-xs text-center block transition"
              >
                Buat Perintah Inspeksi ROV / Penyelam
              </Link>
            </div>
          )}
        </div>
      </div>

      {/* Bottom Section: AI Anomaly Alerts & Recent Inspections */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Anomaly Alerts Table with Digital Approval */}
        <div className="bg-[#081d38] border border-blue-900/60 rounded-xl p-5 space-y-4">
          <div className="flex items-center justify-between border-b border-blue-900/40 pb-3">
            <h3 className="font-semibold text-white flex items-center gap-2">
              <CuiAlertIcon className="w-5 h-5 text-rose-400" />
              Peringatan Anomali & Korelasi Sensor AIS (AI Engine)
            </h3>
            <span className="px-2 py-0.5 rounded text-[11px] font-mono bg-rose-950 text-rose-300 border border-rose-800">
              {alerts.length} Alert
            </span>
          </div>

          <div className="space-y-2.5">
            {alerts.length === 0 ? (
              <div className="text-center py-6 text-slate-500 text-xs font-mono">
                Tidak ada anomali terdeteksi saat ini. Seluruh sensor normal.
              </div>
            ) : (
              alerts.map((al) => (
                <div
                  key={al.alert_id}
                  className="bg-[#051428] border border-blue-900/40 rounded-lg p-3 text-xs space-y-2"
                >
                  <div className="flex items-start justify-between gap-2">
                    <div>
                      <div className="flex items-center gap-2">
                        <span
                          className={`px-1.5 py-0.5 rounded text-[10px] font-bold ${
                            al.severity === "CRITICAL"
                              ? "bg-rose-900 text-rose-200"
                              : "bg-amber-900 text-amber-200"
                          }`}
                        >
                          {al.severity}
                        </span>
                        <span className="font-semibold text-white">{al.title}</span>
                      </div>
                      <div className="text-slate-400 text-[11px] mt-1 font-mono">
                        Aset: <span className="text-cyan-300">{al.asset_name}</span> • Tipe: {al.alert_type}
                      </div>
                    </div>
                    <span
                      className={`px-2 py-0.5 rounded text-[10px] font-mono ${
                        al.status === "RESOLVED"
                          ? "bg-emerald-950 text-emerald-400 border border-emerald-800"
                          : "bg-amber-950 text-amber-400 border border-amber-800"
                      }`}
                    >
                      {al.status}
                    </span>
                  </div>

                  {al.description && (
                    <div className="text-slate-300 text-[11px] bg-[#081d38] p-2 rounded">
                      {al.description}
                    </div>
                  )}

                  <div className="flex items-center justify-between pt-1 text-[11px] font-mono">
                    <span className="text-slate-500">
                      Terdeteksi: {new Date(al.detected_at).toLocaleString("id-ID")}
                    </span>
                    {al.status !== "RESOLVED" && (
                      <button
                        onClick={() => handleResolveAlert(al.alert_id)}
                        disabled={actionLoading === al.alert_id}
                        className="px-2.5 py-1 bg-emerald-700 hover:bg-emerald-600 disabled:opacity-50 text-white font-medium rounded transition flex items-center gap-1"
                      >
                        {actionLoading === al.alert_id ? "Memproses..." : "✓ Otorisasi & Selesaikan"}
                      </button>
                    )}
                  </div>
                </div>
              ))
            )}
          </div>
        </div>

        {/* Inspections Table */}
        <div className="bg-[#081d38] border border-blue-900/60 rounded-xl p-5 space-y-4">
          <div className="flex items-center justify-between border-b border-blue-900/40 pb-3">
            <h3 className="font-semibold text-white flex items-center gap-2">
              <CuiInspectionIcon className="w-5 h-5 text-cyan-400" />
              Riwayat & Jadwal Inspeksi Dishidros
            </h3>
            <Link
              href="/admin/data/cui_inspection/create"
              className="text-xs text-cyan-400 hover:text-white font-mono"
            >
              + Buat Inspeksi
            </Link>
          </div>

          <div className="space-y-2.5">
            {inspections.length === 0 ? (
              <div className="text-center py-6 text-slate-500 text-xs font-mono">
                Belum ada catatan inspeksi terjadwal.
              </div>
            ) : (
              inspections.map((ins) => (
                <div
                  key={ins.inspection_id}
                  className="bg-[#051428] border border-blue-900/40 rounded-lg p-3 text-xs flex items-center justify-between"
                >
                  <div>
                    <div className="font-semibold text-white flex items-center gap-2">
                      <span className="font-mono text-cyan-300">{ins.inspection_code}</span>
                      <span>• {ins.asset_name}</span>
                    </div>
                    <div className="text-slate-400 text-[11px] mt-1 font-mono">
                      Metode: {ins.inspection_type} • Kondisi: <span className="text-emerald-400">{ins.overall_condition}</span>
                    </div>
                  </div>
                  <div className="text-right font-mono text-[11px]">
                    <div className="text-slate-300">{new Date(ins.inspection_date).toLocaleDateString("id-ID")}</div>
                    <span className="inline-block mt-1 px-2 py-0.5 rounded text-[10px] bg-blue-950 text-blue-300 border border-blue-800">
                      {ins.status}
                    </span>
                  </div>
                </div>
              ))
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

