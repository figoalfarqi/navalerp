"use client";

import MobileBottomNav, {
  MobileNavItem,
} from "@/components/mobile/MobileBottomNav";
import { BiSolidTruck, BiUser } from "@/components/icons";

const driverNavigation: MobileNavItem[] = [
  {
    href: "/driver/transport",
    label: "Perjalanan",
    icon: <BiSolidTruck size={23} />,
    activePrefixes: ["/driver/transport"],
  },
  {
    href: "/driver/account",
    label: "Akun",
    icon: <BiUser size={23} />,
  },
];

export default function BottomNav() {
  return <MobileBottomNav items={driverNavigation} />;
}
