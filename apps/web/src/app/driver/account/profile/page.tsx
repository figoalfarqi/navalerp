"use client";

import MobileProfileForm from "@/components/mobile/MobileProfileForm";
import { useAuth } from "@/context/AuthContext";

export default function DriverProfilePage() {
  const { driverPayload } = useAuth();
  return <MobileProfileForm role="driver" userId={driverPayload.app_user_id} />;
}
