/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useState } from "react";
import Button from "../form/Button";
import TextField from "../form/TextField";
import Image from "next/image";
import { useToast } from "../ToastContext";
import LoaderDots from "../form/LoaderDots";
import { FiUser } from "react-icons/fi";
import { GoLock } from "react-icons/go";
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
    // username: role == "admin" ? "08123" : role=="driver" ? "driver1" : "custpic1",
    // password: "dutakasih",
    username: "",
    password: "",
  });
  const [loadingSumbit, setLoadingSumbit] = useState(false);
  const url = `${process.env.NEXT_PUBLIC_API_BASE_URL}/${role}/login`;
  const { showToast } = useToast();
  const handleSubmit = async (e: React.FormEvent) => {
    setLoadingSumbit(true);
    e.preventDefault();
    try {
      const res = await postAPI<any>(url, { authToken: "none" }, formData);
      if (res.code === 200) {
        if (onSuccess) onSuccess(res);
        showToast(3000, "success", "Selamat Datang, Semangat bekerja!");
      } else {
        showToast(3000, "error", "Gagal Melakukan Login!");
      }
    } catch (err) {
      console.error("Submit failed", err);
      showToast(3000, "error", "Gagal Melakukan Login!");
    } finally {
      setLoadingSumbit(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-200 via-gray-100 to-white px-4 sm:px-6 lg:px-8">
      <div
        className="w-full max-w-sm sm:max-w-md bg-white p-4 sm:p-8 rounded-xl shadow-[0_8px_30px_rgba(0,48,101,0.5)]
 border-t-4 border-[#004f7f]"
      >
        <div className="w-full max-w-md bg-white rounded-2xl p-2">
          <div className="flex justify-center mb-4">
            <Image
              src="/logo-pml.png"
              alt="App Logo"
              width={isMobile ? 200: 240}
              height={isMobile ? 200: 240}
            />
          </div>
          <h2 className="text-2xl font-bold text-gray-800 text-center">
            Selamat Datang {role.slice(0, 1).toUpperCase()}
            {role.slice(1)}
          </h2>
          <p className="text-gray-500 text-center">
            Silakan login untuk melanjutkan
          </p>

          {/* Form */}
          <form onSubmit={handleSubmit} className="mt-6 space-y-4">
            <div>
              <TextField
                id={"username_login"}
                icon={
                  <FiUser className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-7 h-7" />
                }
                placeholder="Username"
                border="border-gray-300"
                value={formData["username"] || ""}
                onChange={(val) =>
                  setFormData({ ...formData, ["username"]: val })
                }
                className="w-full pl-11 text-xl"
              />
            </div>

            <div>
              <TextField
                id={"password_login"}
                icon={
                  <GoLock className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-7 h-7" />
                }
                type="password"
                isPasswordShowable
                placeholder="Password"
                border="border-gray-300"
                value={formData["password"] || ""}
                onChange={(val) =>
                  setFormData({ ...formData, ["password"]: val })
                }
                className="w-full pl-11 text-xl"
              />
            </div>
            <div className="w-full">
              <Button
                id={"login"}
                type="submit"
                variant="blue-dkl"
                size="xl"
                wrapperClassName="!block"
                className="w-full h-11"
              >
                {loadingSumbit ? <LoaderDots /> : "Login"}
              </Button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}
