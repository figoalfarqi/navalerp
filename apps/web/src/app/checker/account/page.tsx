"use client";

import MobileAccountPage from "@/components/mobile/MobileAccountPage";
import { useAuth } from "@/context/AuthContext";

export default function Page() {
  const { checkerPayload, clearCheckerToken } = useAuth();

  return (
    <MobileAccountPage
      role="checker"
      user={{
        name: checkerPayload.app_user_name,
        username: checkerPayload.username,
        photoUrl: checkerPayload.app_user_photo_url,
      }}
      clearToken={clearCheckerToken}
    />
  );
}
