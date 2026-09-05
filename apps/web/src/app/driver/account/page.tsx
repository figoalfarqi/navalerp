"use client";

import MobileAccountPage from "@/components/mobile/MobileAccountPage";
import { useAuth } from "@/context/AuthContext";

export default function Page() {
  const { driverPayload, clearDriverToken } = useAuth();

  return (
    <MobileAccountPage
      role="driver"
      user={{
        name: driverPayload.app_user_name,
        username: driverPayload.username,
        photoUrl: driverPayload.app_user_photo_url,
      }}
      clearToken={clearDriverToken}
    />
  );
}
