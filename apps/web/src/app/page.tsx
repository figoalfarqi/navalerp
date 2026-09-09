"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import Image from "next/image";
import { useAuth } from "@/context/AuthContext";
import { FaAnchor, FaShieldHalved, FaCompass, FaArrowRight } from "@/components/icons";

export default function Home() {
  const router = useRouter();
  const { adminToken, tokenLoaded } = useAuth();

  useEffect(() => {
    if (tokenLoaded && adminToken) {
      router.replace("/admin");
    }
  }, [adminToken, tokenLoaded, router]);

  return (
    <main className="min-h-screen bg-gradient-to-br from-slate-950 via-[#072444] to-[#041224] text-white flex flex-col justify-between">
      {/* Navbar Header */}
      <header className="w-full max-w-7xl mx-auto px-6 py-6 flex items-center justify-between">
        <div className="flex items-center gap-3">
          <Image
            src="/logo.png"
            alt="Naval ERP"
            width={48}
            height={48}
            className="rounded-full shadow-lg"
          />
          <div>
            <h1 className="text-xl font-bold tracking-tight text-white">NAVAL ERP</h1>
            <p className="text-xs text-cyan-300 font-medium tracking-wider uppercase">TNI Angkatan Laut</p>
          </div>
        </div>
        <button
          onClick={() => router.push("/admin/login")}
          className="inline-flex items-center gap-2 px-5 py-2.5 rounded-xl bg-cyan-600 hover:bg-cyan-500 text-white font-medium text-sm transition shadow-lg shadow-cyan-900/30"
        >
          Masuk Portal <FaArrowRight size={14} />
        </button>
      </header>

      {/* Hero Section */}
      <div className="w-full max-w-5xl mx-auto px-6 py-12 flex flex-col items-center text-center">
        <div className="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-cyan-950/80 border border-cyan-500/30 text-cyan-300 text-xs font-semibold tracking-wide uppercase mb-8 backdrop-blur">
          <FaShieldHalved size={14} /> Integrated Defence Resource Planning
        </div>

        <h2 className="text-4xl sm:text-6xl font-extrabold tracking-tight text-white leading-tight">
          Sistem Informasi & Manajemen <br />
          <span className="bg-gradient-to-r from-cyan-400 via-teal-300 to-blue-400 bg-clip-text text-transparent">
            Alutsista Matra Laut
          </span>
        </h2>

        <p className="mt-6 text-base sm:text-lg text-slate-300 max-w-2xl leading-relaxed">
          Platform Enterprise terintegrasi untuk pengelolaan Armada KRI, Kesiapan Operasi Tempur,
          Pemeliharaan & MRO, Logistik Persenjataan, dan Manajemen Personel Militer.
        </p>

        <div className="mt-10 flex flex-col sm:flex-row gap-4 w-full sm:w-auto">
          <button
            onClick={() => router.push("/admin")}
            className="inline-flex items-center justify-center gap-3 px-8 py-4 rounded-xl bg-gradient-to-r from-cyan-500 to-blue-600 hover:from-cyan-400 hover:to-blue-500 text-white font-semibold text-base transition shadow-xl shadow-cyan-900/40"
          >
            Buka Console Naval ERP <FaArrowRight size={16} />
          </button>
          <button
            onClick={() => router.push("/admin/login")}
            className="inline-flex items-center justify-center gap-2 px-8 py-4 rounded-xl bg-white/10 hover:bg-white/15 border border-white/20 text-white font-semibold text-base backdrop-blur transition"
          >
            Login Petugas
          </button>
        </div>

        {/* 3 Pillars */}
        <div className="mt-16 grid grid-cols-1 sm:grid-cols-3 gap-6 w-full text-left">
          <div className="p-6 rounded-2xl bg-white/5 border border-white/10 backdrop-blur">
            <div className="w-12 h-12 rounded-xl bg-cyan-500/20 text-cyan-400 flex items-center justify-center mb-4">
              <FaAnchor size={22} />
            </div>
            <h3 className="text-lg font-bold text-white mb-2">Armada & Kesiapan</h3>
            <p className="text-xs text-slate-300 leading-relaxed">
              Monitoring kesiapan alutsista KRI, sensor, persenjataan, serta status operasional pangkalan dan satuan tempur.
            </p>
          </div>

          <div className="p-6 rounded-2xl bg-white/5 border border-white/10 backdrop-blur">
            <div className="w-12 h-12 rounded-xl bg-blue-500/20 text-blue-400 flex items-center justify-center mb-4">
              <FaCompass size={22} />
            </div>
            <h3 className="text-lg font-bold text-white mb-2">Misi & Operasi</h3>
            <p className="text-xs text-slate-300 leading-relaxed">
              Perencanaan dan pencatatan misi pelayaran, log harian navigasi, pergerakan satuan angkut, dan logistik perbekalan.
            </p>
          </div>

          <div className="p-6 rounded-2xl bg-white/5 border border-white/10 backdrop-blur">
            <div className="w-12 h-12 rounded-xl bg-teal-500/20 text-teal-400 flex items-center justify-center mb-4">
              <FaShieldHalved size={22} />
            </div>
            <h3 className="text-lg font-bold text-white mb-2">MRO & Rantai Pasok</h3>
            <p className="text-xs text-slate-300 leading-relaxed">
              Work order pemeliharaan berkala, docking record fasharkan, pengadaan alpalhankam, dan perputaran suku cadang.
            </p>
          </div>
        </div>
      </div>

      {/* Footer */}
      <footer className="w-full border-t border-white/10 py-6 px-6 text-center text-xs text-slate-400">
        Naval Enterprise Resource Planning System &copy; 2026 TNI Angkatan Laut. Hak Cipta Dilindungi.
      </footer>
    </main>
  );
}
