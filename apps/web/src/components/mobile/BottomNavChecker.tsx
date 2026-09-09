"use client";

import MobileBottomNav, {
  MobileNavItem,
} from "@/components/mobile/MobileBottomNav";
import { BiSolidTruck, BiUser } from "@/components/icons";
import { FiActivity, FiGrid } from "@/components/icons";

const checkerNavigation: MobileNavItem[] = [
  {
    href: "/checker/transport",
    label: "Lapor",
    icon: <BiSolidTruck size={22} />,
    activePrefixes: ["/checker/transport"],
  },
  {
    href: "/checker/active",
    label: "Aktif",
    icon: <FiActivity size={21} />,
  },
  {
    href: "/checker/more",
    label: "Lainnya",
    icon: <FiGrid size={21} />,
  },
  {
    href: "/checker/account",
    label: "Akun",
    icon: <BiUser size={22} />,
  },
];

export default function BottomNavChecker() {
  return <MobileBottomNav items={checkerNavigation} />;
}
