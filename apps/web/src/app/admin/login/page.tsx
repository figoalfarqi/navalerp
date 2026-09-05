/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { useRouter } from "next/navigation";
import Login from "@/components/login/Login";
import { useAuth } from "@/context/AuthContext";

export default function Page() {
  const router = useRouter();
  const { setAdminToken } = useAuth();

  const loginSucces = async (res: any) => {
    if (res?.data?.token) {
      await setAdminToken(res.data.token);
      router.push("/admin");
    }
  };

  return <Login onSuccess={loginSucces} role="admin" />;
}
