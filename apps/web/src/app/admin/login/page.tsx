/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import Login from "@/components/login/Login";
import { useAuth } from "@/context/AuthContext";

export default function Page() {
  const { setAdminToken } = useAuth();

  const loginSucces = (res: any) => {
    if (res?.data?.token) {
      setAdminToken(res.data.token);
    }
  };

  return <Login onSuccess={loginSucces} role="admin" />;
}
