"use client";

import MobileProfileForm from "@/components/mobile/MobileProfileForm";
import { useAuth } from "@/context/AuthContext";

export default function CheckerProfilePage() {
  const { checkerPayload } = useAuth();
  return (
    <MobileProfileForm role="checker" userId={checkerPayload.app_user_id} />
  );
}
