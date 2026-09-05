import type { IconType } from "react-icons";
import {
  FaBoxesStacked,
  FaChartLine,
  FaCity,
  FaClipboardList,
  FaFileInvoiceDollar,
  FaGear,
  FaMapLocationDot,
  FaRoute,
  FaShip,
  FaTruck,
  FaUsersGear,
  FaWarehouse,
} from "react-icons/fa6";
import { HiOutlineDocumentReport } from "react-icons/hi";
import { LuBanknote, LuSettings2 } from "react-icons/lu";

export type AdminRoleId = 3 | 4 | 5 | 6;

export const ADMIN_ROLE_NAMES: Record<number, string> = {
  3: "Owner",
  4: "IT Developer",
  5: "Super Admin",
  6: "Admin",
};

export interface AdminNavigationItem {
  label: string;
  href: string;
  icon: IconType;
  roles: AdminRoleId[];
}

export interface AdminNavigationGroup {
  label: string;
  icon: IconType;
  items: AdminNavigationItem[];
}

const allAdminRoles: AdminRoleId[] = [3, 4, 5, 6];
const operationalRoles: AdminRoleId[] = [5, 6];
const systemRoles: AdminRoleId[] = [4, 5];
const userManagementRoles: AdminRoleId[] = [4, 5, 6];

export const ADMIN_NAVIGATION: AdminNavigationGroup[] = [
  {
    label: "Overview",
    icon: FaChartLine,
    items: [
      {
        label: "Dashboard",
        href: "/admin",
        icon: FaChartLine,
        roles: allAdminRoles,
      },
      {
        label: "Laporan",
        href: "/admin/report",
        icon: HiOutlineDocumentReport,
        roles: [3, 5, 6],
      },
    ],
  },
  {
    label: "Project & Operasional",
    icon: FaRoute,
    items: [
      {
        label: "Project",
        href: "/admin/data/project",
        icon: FaRoute,
        roles: operationalRoles,
      },
      {
        label: "Rute Project",
        href: "/admin/data/project_route",
        icon: FaMapLocationDot,
        roles: operationalRoles,
      },
      {
        label: "Checker Project",
        href: "/admin/data/project_checker_assignment",
        icon: FaUsersGear,
        roles: operationalRoles,
      },
      {
        label: "Truck Project",
        href: "/admin/data/project_truck_assignment",
        icon: FaTruck,
        roles: operationalRoles,
      },
      {
        label: "Transport",
        href: "/admin/data/project_transport",
        icon: FaTruck,
        roles: operationalRoles,
      },
      {
        label: "Transaksi Finansial",
        href: "/admin/data/project_financial_transaction",
        icon: FaFileInvoiceDollar,
        roles: operationalRoles,
      },
    ],
  },
  {
    label: "Lokasi & Cargo",
    icon: FaMapLocationDot,
    items: [
      {
        label: "Provinsi",
        href: "/admin/data/province",
        icon: FaCity,
        roles: operationalRoles,
      },
      {
        label: "Kota",
        href: "/admin/data/city",
        icon: FaCity,
        roles: operationalRoles,
      },
      {
        label: "Tambang",
        href: "/admin/data/mine",
        icon: FaMapLocationDot,
        roles: operationalRoles,
      },
      {
        label: "Pelabuhan",
        href: "/admin/data/port",
        icon: FaMapLocationDot,
        roles: operationalRoles,
      },
      {
        label: "Stockpile",
        href: "/admin/data/stockpile",
        icon: FaWarehouse,
        roles: operationalRoles,
      },
      {
        label: "Cargo Type",
        href: "/admin/data/cargo_type",
        icon: FaBoxesStacked,
        roles: operationalRoles,
      },
      {
        label: "Stockpile Cargo",
        href: "/admin/data/stockpile_cargo",
        icon: FaBoxesStacked,
        roles: operationalRoles,
      },
      {
        label: "Adjustment",
        href: "/admin/data/stockpile_adjustment",
        icon: LuSettings2,
        roles: operationalRoles,
      },
      {
        label: "Ledger",
        href: "/admin/data/stockpile_ledger",
        icon: FaClipboardList,
        roles: operationalRoles,
      },
    ],
  },
  {
    label: "Armada & Vendor",
    icon: FaTruck,
    items: [
      {
        label: "Vendor Type",
        href: "/admin/data/vendor_type",
        icon: FaUsersGear,
        roles: operationalRoles,
      },
      {
        label: "Vendor",
        href: "/admin/data/vendor",
        icon: FaUsersGear,
        roles: operationalRoles,
      },
      {
        label: "Truck Type",
        href: "/admin/data/truck_type",
        icon: FaTruck,
        roles: operationalRoles,
      },
      {
        label: "Truck Merk",
        href: "/admin/data/truck_merk",
        icon: FaTruck,
        roles: operationalRoles,
      },
      {
        label: "Truck",
        href: "/admin/data/truck",
        icon: FaTruck,
        roles: operationalRoles,
      },
      {
        label: "Vessel",
        href: "/admin/data/vessel",
        icon: FaShip,
        roles: operationalRoles,
      },
      {
        label: "Vessel Cargo",
        href: "/admin/data/vessel_cargo",
        icon: FaShip,
        roles: operationalRoles,
      },
    ],
  },
  {
    label: "Client",
    icon: FaUsersGear,
    items: [
      {
        label: "Client",
        href: "/admin/data/client",
        icon: FaUsersGear,
        roles: operationalRoles,
      },
      {
        label: "Tujuan Client",
        href: "/admin/data/client_destination",
        icon: FaMapLocationDot,
        roles: operationalRoles,
      },
    ],
  },
  {
    label: "Pengguna & Sistem",
    icon: FaGear,
    items: [
      {
        label: "Semua Pengguna",
        href: "/admin/data/app_user",
        icon: FaUsersGear,
        roles: [4, 5],
      },
      {
        label: "Admin",
        href: "/admin/data/admin",
        icon: FaUsersGear,
        roles: userManagementRoles,
      },
      {
        label: "Checker",
        href: "/admin/data/checker",
        icon: FaUsersGear,
        roles: userManagementRoles,
      },
      {
        label: "Driver",
        href: "/admin/data/driver",
        icon: FaUsersGear,
        roles: userManagementRoles,
      },
      {
        label: "Role",
        href: "/admin/data/app_role",
        icon: FaUsersGear,
        roles: systemRoles,
      },
      {
        label: "App Setting",
        href: "/admin/data/app_setting",
        icon: LuSettings2,
        roles: systemRoles,
      },
      {
        label: "Bank",
        href: "/admin/data/bank_merk",
        icon: LuBanknote,
        roles: operationalRoles,
      },
      {
        label: "Pengaturan",
        href: "/admin/settings",
        icon: FaGear,
        roles: allAdminRoles,
      },
    ],
  },
];

export function isAdminRoleId(value: number): value is AdminRoleId {
  return [3, 4, 5, 6].includes(value);
}

export function getAdminNavigation(roleId: number): AdminNavigationGroup[] {
  if (!isAdminRoleId(roleId)) return [];

  const canViewEverything = [3, 4, 5].includes(roleId);
  return ADMIN_NAVIGATION.map((group) => ({
    ...group,
    items: group.items.filter(
      (item) => canViewEverything || item.roles.includes(roleId),
    ),
  })).filter((group) => group.items.length > 0);
}

export function canAccessAdminPath(roleId: number, pathname: string): boolean {
  if (pathname === "/admin/login") return true;
  if (!isAdminRoleId(roleId)) return false;
  if ([3, 4, 5].includes(roleId)) return true;

  return getAdminNavigation(roleId).some((group) =>
    group.items.some(
      (item) =>
        pathname === item.href ||
        (item.href !== "/admin" && pathname.startsWith(`${item.href}/`)),
    ),
  );
}
