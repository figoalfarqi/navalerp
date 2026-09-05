"use client";

import React, {
  createContext,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";
import { useIndexedDB } from "../hooks/useIndexedDB";
import { DriverPayload } from "@/types/driverPayload";
import { decodeJWT } from "@/utils/decodeJWT";
import { AdminPayload } from "@/types/adminPayload";
import { CheckerPayload } from "@/types/checkerPayload";

interface AuthState {
  tokenLoaded: boolean;
  driverToken: string | null;
  driverPayload: DriverPayload;
  adminToken: string | null;
  adminPayload: AdminPayload;
  checkerToken: string | null;
  checkerPayload: CheckerPayload;
  setDriverToken: (token: string) => Promise<void>;
  clearDriverToken: () => Promise<void>;
  setAdminToken: (token: string) => Promise<void>;
  clearAdminToken: () => Promise<void>;
  setCheckerToken: (token: string) => Promise<void>;
  clearCheckerToken: () => Promise<void>;
}

interface Item {
  auth_id: string;
  value: string;
}
const initialDriverState: DriverPayload = {
  app_user_id: 0,
  username: "08123456789",
  app_user_name: "nama driver",
  app_user_photo_url: "",
  app_role_id: 0,
  app_user_status_id: 1,
  work_status_id: 1,
};

const initialAdminState: AdminPayload = {
  app_user_id: 0,
  username: "08123456789",
  app_user_name: "nama admin",
  app_user_photo_url: "",
  app_role_id: 0,
  app_user_status_id: 1,
};

const initialCheckerState: CheckerPayload = {
  app_user_id: 0,
  username: "08123456789",
  app_user_name: "nama checker",
  app_user_photo_url: "",
  app_role_id: 0,
  app_user_status_id: 1,
};

const AuthContext = createContext<AuthState | undefined>(undefined);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const { dbReady, getIDB, putIDB, deleteIDB } = useIndexedDB();
  const [tokenLoaded, setTokenLoaded] = useState(false);
  const [driverToken, setDriverTokenState] = useState<string | null>(null);
  const [adminToken, setAdminTokenState] = useState<string | null>(null);
  const [checkerToken, setCheckerTokenState] = useState<string | null>(null);

  const driverPayload: DriverPayload = useMemo(() => {
    if (!driverToken) return initialDriverState;
    const decoded = decodeJWT<DriverPayload>(driverToken);
    if (!decoded) return initialDriverState;
    return { ...decoded, work_status_id: 1 };
  }, [driverToken]);

  const adminPayload: AdminPayload = useMemo(() => {
    if (!adminToken) return initialAdminState;
    const decoded = decodeJWT<AdminPayload>(adminToken);
    if (!decoded) return initialAdminState;
    return { ...decoded, work_status_id: 1 };
  }, [adminToken]);

  const checkerPayload: CheckerPayload = useMemo(() => {
    if (!checkerToken) return initialCheckerState;
    const decoded = decodeJWT<CheckerPayload>(checkerToken);
    if (!decoded) return initialCheckerState;
    return { ...decoded, work_status_id: 1 };
  }, [checkerToken]);

  // Set Driver Token
  const setDriverToken = async (token: string) => {
    await putIDB("auth", { auth_id: "driver_token", value: token });
    setDriverTokenState(token);
  };

  // Clear Driver Token
  const clearDriverToken = async () => {
    await deleteIDB("auth", "driver_token");
    setDriverTokenState(null);
  };

  // Set Admin Token
  const setAdminToken = async (token: string) => {
    await putIDB("auth", { auth_id: "admin_token", value: token });
    setAdminTokenState(token);
  };

  // Clear Admin Token
  const clearAdminToken = async () => {
    await deleteIDB("auth", "admin_token");
    setAdminTokenState(null);
  };

  // Set Checker Token
  const setCheckerToken = async (token: string) => {
    await putIDB("auth", { auth_id: "checker_token", value: token });
    setCheckerTokenState(token);
  };

  // Clear Checker Token
  const clearCheckerToken = async () => {
    await deleteIDB("auth", "checker_token");
    setCheckerTokenState(null);
  };

  // Load semua token dari DB saat startup
  useEffect(() => {
    if (!dbReady) return;

    let isMounted = true;

    (async () => {
      try {
        const all = await getIDB("auth");

        if (!isMounted) return;

        const storedItems: Item[] = Array.isArray(all.data) ? all.data : [];
        const driver =
          storedItems.find((item) => item.auth_id === "driver_token")?.value ??
          null;

        const admin =
          storedItems.find((item) => item.auth_id === "admin_token")?.value ??
          null;

        const checker =
          storedItems.find((item) => item.auth_id === "checker_token")?.value ??
          null;

        setDriverTokenState(driver);
        setAdminTokenState(admin);
        setCheckerTokenState(checker);
      } catch {
        if (!isMounted) return;
        setDriverTokenState(null);
        setAdminTokenState(null);
        setCheckerTokenState(null);
      } finally {
        if (isMounted) setTokenLoaded(true);
      }
    })();

    return () => {
      isMounted = false;
    };
  }, [dbReady]);

  return (
    <AuthContext.Provider
      value={{
        tokenLoaded,
        driverToken,
        driverPayload,
        adminToken,
        adminPayload,
        checkerToken,
        checkerPayload,
        setDriverToken,
        clearDriverToken,
        setAdminToken,
        clearAdminToken,
        setCheckerToken,
        clearCheckerToken,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export function useAuth() {
  const ctx = useContext(AuthContext);
  if (!ctx) throw new Error("useAuth must be used within AuthProvider");
  return ctx;
}
