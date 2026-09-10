/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useState } from "react";
import Link from "next/link";
import Button from "../form/Button";
import TextField from "../form/TextField";
import Image from "next/image";
import { useToast } from "../ToastContext";
import LoaderDots from "../form/LoaderDots";
import { FiUser } from "@/components/icons";
import { GoLock } from "@/components/icons";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { useIsMobile } from "@/hooks/useIsMobile";

export default function Login({
  role,
  onSuccess,
}: {
  role: "driver" | "admin" | "customer" | "checker";
  onSuccess?: (res: any) => void;
}) {
  const { isMobile } = useIsMobile();
  const { postAPI } = useFetchAPI();
  const [formData, setFormData] = useState<{ [key: string]: string | number }>({
    username: "",
    password: "",
  });
  const [loadingSubmit, setLoadingSubmit] = useState(false);
  const url = `${process.env.NEXT_PUBLIC_API_BASE_URL}/${role}/login`;
  const { showToast } = useToast();

  const handleSubmit = async (e: React.FormEvent) => {
    setLoadingSubmit(true);
    e.preventDefault();
    try {
      const res = await postAPI<any>(url, { authToken: "none" }, formData);
      if (res.code === 200) {
        if (onSuccess) onSuccess(res);
        showToast(3000, "success", "Selamat Datang di Sistem Naval ERP!");
      } else {
        showToast(3000, "error", res.message || "Gagal Melakukan Login!");
      }
    } catch (err) {
      console.error("Submit failed", err);
      showToast(3000, "error", "Gagal Melakukan Login!");
    } finally {
      setLoadingSubmit(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-950 via-[#071a33] to-slate-950 px-4 sm:px-6 lg:px-8">
      <div className="w-full max-w-sm sm:max-w-md bg-white p-6 sm:p-8 rounded-2xl shadow-[0_20px_50px_rgba(0,30,70,0.4)] border-t-4 border-[#0a2540]">
        <div className="w-full">
          <div className="flex items-center justify-between mb-4">
            <Link
              href="/#security"
              className="inline-flex items-center gap-1.5 text-xs font-semibold text-slate-500 hover:text-blue-700 transition cursor-pointer"
            >
              ← Kembali ke Beranda
            </Link>
            <span className="text-[10px] font-mono text-cyan-800 bg-cyan-50 px-2 py-0.5 rounded border border-cyan-200">
              TNI AL SECURE
            </span>
          </div>

          <div className="flex justify-center mb-4">
            <Image
              src="/logo.webp"
              alt="Naval ERP"
              width={isMobile ? 100 : 120}
              height={isMobile ? 100 : 120}
              priority
              className="object-contain"
            />
          </div>
          <h2 className="text-2xl font-bold text-gray-800 text-center tracking-tight">
            Naval ERP Portal
          </h2>
          <p className="text-gray-500 text-center text-xs sm:text-sm mt-1">
            Sistem Informasi Operasi & Alutsista TNI AL
          </p>

          <form onSubmit={handleSubmit} className="mt-6 space-y-4">
            <div>
              <TextField
                id="username_login"
                icon={
                  <FiUser className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                }
                placeholder="Username (e.g. admin)"
                border="border-gray-300"
                value={formData["username"] || ""}
                onChange={(val) =>
                  setFormData({ ...formData, username: val })
                }
                className="w-full pl-11 text-base"
              />
            </div>

            <div>
              <TextField
                id="password_login"
                icon={
                  <GoLock className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                }
                type="password"
                isPasswordShowable
                placeholder="Password (e.g. Password123!)"
                border="border-gray-300"
                value={formData["password"] || ""}
                onChange={(val) =>
                  setFormData({ ...formData, password: val })
                }
                className="w-full pl-11 text-base"
              />
            </div>

            <div className="w-full pt-2">
              <Button
                id="login"
                type="submit"
                variant="blue-dkl"
                size="xl"
                wrapperClassName="!block"
                className="w-full h-11 text-base font-semibold bg-[#0a2540] hover:bg-[#081d38] text-white shadow-md transition"
              >
                {loadingSubmit ? <LoaderDots /> : "Masuk Sistem"}
              </Button>
            </div>
          </form>

          <div className="mt-6 text-center border-t border-slate-100 pt-4">
            <Link
              href="/#security"
              className="inline-flex items-center gap-1.5 text-xs font-medium text-slate-500 hover:text-blue-700 transition"
            >
              ← Kembali ke Halaman Utama
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
