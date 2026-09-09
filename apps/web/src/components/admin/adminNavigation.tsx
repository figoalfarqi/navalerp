import type { ComponentType } from "react";
import type { IconProps } from "@/components/icons";
import {
  // Group icons
  OverviewGroupIcon,
  OrganizationGroupIcon,
  FleetGroupIcon,
  LogisticsGroupIcon,
  ProcurementGroupIcon,
  FinanceGroupIcon,
  PersonnelGroupIcon,
  BaseFacilityGroupIcon,
  TransportGroupIcon,
  EdrmsGroupIcon,
  CommandReadinessGroupIcon,
  // 00. Overview
  DashboardIcon,
  ReportIcon,
  // 01. Organisasi & Pengguna
  OrgUnitIcon,
  SysUserIcon,
  AuditLogIcon,
  // 02. Armada Kapal & MRO
  ShipClassIcon,
  WarshipIcon,
  ShipSystemIcon,
  EquipmentIcon,
  PmScheduleIcon,
  FailureReportIcon,
  WorkOrderIcon,
  DockingRecordIcon,
  // 03. Logistik & Pergudangan
  WarehouseIcon,
  MaterialIcon,
  StockBalanceIcon,
  ItemInstanceIcon,
  StockTransferIcon,
  StockAdjustmentIcon,
  // 04. Pengadaan Pertahanan
  VendorIcon,
  RequisitionIcon,
  TenderIcon,
  ContractIcon,
  PurchaseOrderIcon,
  GoodsReceiptIcon,
  // 05. Keuangan & Anggaran
  ChartOfAccountIcon,
  BudgetProgramIcon,
  BudgetCommitmentIcon,
  InvoiceIcon,
  PaymentIcon,
  JournalEntryIcon,
  PlatformTcoIcon,
  // 06. SDM Militer (HCM)
  MilitaryRankIcon,
  MilitaryCorpsIcon,
  QualificationIcon,
  PersonnelIcon,
  CrewAssignmentIcon,
  // 07. Pangkalan & Fasilitas
  BaseFacilityIcon,
  BerthBookingIcon,
  FuelBunkerIcon,
  // 08. Transportasi Militer
  TransportUnitIcon,
  RouteIcon,
  ShipmentIcon,
  // 09. EDRMS Dokumen Digital
  DocumentCategoryIcon,
  DocumentIcon,
  // 10. Komando & Kesiapan
  TheaterIcon,
  MissionIcon,
  DailyLogIcon,
  ReadinessReportIcon,
  ReadinessAlertIcon,
} from "@/components/icons";

export type AdminRoleId = 3 | 4 | 5 | 6;

export const ADMIN_ROLE_NAMES: Record<number, string> = {
  3: "Owner",
  4: "IT Developer",
  5: "Super Admin",
  6: "Admin",
};

export interface AdminNavigationItem {
  name?: string;
  label: string;
  href: string;
  url?: string;
  icon: ComponentType<IconProps>;
  roles: AdminRoleId[];
}

export interface AdminNavigationGroup {
  title?: string;
  label: string;
  icon: ComponentType<IconProps>;
  items: AdminNavigationItem[];
}

const allAdminRoles: AdminRoleId[] = [3, 4, 5, 6];

export const ADMIN_NAVIGATION: AdminNavigationGroup[] = [
  {
    label: "Overview",
    title: "Overview",
    icon: OverviewGroupIcon,
    items: [
      {
        label: "Dashboard",
        name: "Dashboard",
        href: "/admin",
        url: "/admin",
        icon: DashboardIcon,
        roles: allAdminRoles,
      },
      {
        label: "Laporan",
        name: "Laporan",
        href: "/admin/report",
        url: "/admin/report",
        icon: ReportIcon,
        roles: allAdminRoles,
      },
    ],
  },
  {
    label: "01. Organisasi & Pengguna",
    title: "01. Organisasi & Pengguna",
    icon: OrganizationGroupIcon,
    items: [
      {
        label: "Satuan Kerja",
        name: "Satuan Kerja",
        href: "/admin/data/org_unit",
        url: "/admin/data/org_unit",
        icon: OrgUnitIcon,
        roles: allAdminRoles,
      },
      {
        label: "Pengguna Sistem",
        name: "Pengguna Sistem",
        href: "/admin/data/sys_user",
        url: "/admin/data/sys_user",
        icon: SysUserIcon,
        roles: allAdminRoles,
      },
      {
        label: "Audit Log",
        name: "Audit Log",
        href: "/admin/data/audit_log",
        url: "/admin/data/audit_log",
        icon: AuditLogIcon,
        roles: allAdminRoles,
      },
    ],
  },
  {
    label: "02. Armada Kapal & MRO",
    title: "02. Armada Kapal & MRO",
    icon: FleetGroupIcon,
    items: [
      {
        label: "Kelas Kapal",
        name: "Kelas Kapal",
        href: "/admin/data/ship_class",
        url: "/admin/data/ship_class",
        icon: ShipClassIcon,
        roles: allAdminRoles,
      },
      {
        label: "Kapal Perang KRI",
        name: "Kapal Perang KRI",
        href: "/admin/data/ship",
        url: "/admin/data/ship",
        icon: WarshipIcon,
        roles: allAdminRoles,
      },
      {
        label: "Sistem Kapal",
        name: "Sistem Kapal",
        href: "/admin/data/ship_system",
        url: "/admin/data/ship_system",
        icon: ShipSystemIcon,
        roles: allAdminRoles,
      },
      {
        label: "Peralatan Mesin",
        name: "Peralatan Mesin",
        href: "/admin/data/equipment",
        url: "/admin/data/equipment",
        icon: EquipmentIcon,
        roles: allAdminRoles,
      },
      {
        label: "Jadwal PMS",
        name: "Jadwal PMS",
        href: "/admin/data/pm_schedule",
        url: "/admin/data/pm_schedule",
        icon: PmScheduleIcon,
        roles: allAdminRoles,
      },
      {
        label: "Laporan Kerusakan",
        name: "Laporan Kerusakan",
        href: "/admin/data/failure_report",
        url: "/admin/data/failure_report",
        icon: FailureReportIcon,
        roles: allAdminRoles,
      },
      {
        label: "Perintah Kerja MRO",
        name: "Perintah Kerja MRO",
        href: "/admin/data/work_order",
        url: "/admin/data/work_order",
        icon: WorkOrderIcon,
        roles: allAdminRoles,
      },
      {
        label: "Riwayat Docking",
        name: "Riwayat Docking",
        href: "/admin/data/docking_record",
        url: "/admin/data/docking_record",
        icon: DockingRecordIcon,
        roles: allAdminRoles,
      },
    ],
  },
  {
    label: "03. Logistik & Pergudangan",
    title: "03. Logistik & Pergudangan",
    icon: LogisticsGroupIcon,
    items: [
      {
        label: "Gudang Militer",
        name: "Gudang Militer",
        href: "/admin/data/warehouse",
        url: "/admin/data/warehouse",
        icon: WarehouseIcon,
        roles: allAdminRoles,
      },
      {
        label: "Material & Suku Cadang",
        name: "Material & Suku Cadang",
        href: "/admin/data/material",
        url: "/admin/data/material",
        icon: MaterialIcon,
        roles: allAdminRoles,
      },
      {
        label: "Saldo Stok",
        name: "Saldo Stok",
        href: "/admin/data/stock_balance",
        url: "/admin/data/stock_balance",
        icon: StockBalanceIcon,
        roles: allAdminRoles,
      },
      {
        label: "Nomor Seri / Batch",
        name: "Nomor Seri / Batch",
        href: "/admin/data/item_instance",
        url: "/admin/data/item_instance",
        icon: ItemInstanceIcon,
        roles: allAdminRoles,
      },
      {
        label: "Transfer Bebekal",
        name: "Transfer Bebekal",
        href: "/admin/data/stock_transfer",
        url: "/admin/data/stock_transfer",
        icon: StockTransferIcon,
        roles: allAdminRoles,
      },
      {
        label: "Penyesuaian Stok",
        name: "Penyesuaian Stok",
        href: "/admin/data/stock_adjustment",
        url: "/admin/data/stock_adjustment",
        icon: StockAdjustmentIcon,
        roles: allAdminRoles,
      },
    ],
  },
  {
    label: "04. Pengadaan Pertahanan",
    title: "04. Pengadaan Pertahanan",
    icon: ProcurementGroupIcon,
    items: [
      {
        label: "Rekanan Industri",
        name: "Rekanan Industri",
        href: "/admin/data/vendor",
        url: "/admin/data/vendor",
        icon: VendorIcon,
        roles: allAdminRoles,
      },
      {
        label: "Permintaan Pengadaan",
        name: "Permintaan Pengadaan",
        href: "/admin/data/requisition",
        url: "/admin/data/requisition",
        icon: RequisitionIcon,
        roles: allAdminRoles,
      },
      {
        label: "Tender & Lelang",
        name: "Tender & Lelang",
        href: "/admin/data/tender",
        url: "/admin/data/tender",
        icon: TenderIcon,
        roles: allAdminRoles,
      },
      {
        label: "Kontrak Militer",
        name: "Kontrak Militer",
        href: "/admin/data/contract",
        url: "/admin/data/contract",
        icon: ContractIcon,
        roles: allAdminRoles,
      },
      {
        label: "Purchase Order",
        name: "Purchase Order",
        href: "/admin/data/purchase_order",
        url: "/admin/data/purchase_order",
        icon: PurchaseOrderIcon,
        roles: allAdminRoles,
      },
      {
        label: "Penerimaan BAPHP",
        name: "Penerimaan BAPHP",
        href: "/admin/data/goods_receipt",
        url: "/admin/data/goods_receipt",
        icon: GoodsReceiptIcon,
        roles: allAdminRoles,
      },
    ],
  },
  {
    label: "05. Keuangan & Anggaran",
    title: "05. Keuangan & Anggaran",
    icon: FinanceGroupIcon,
    items: [
      {
        label: "Bagan Akun (COA)",
        name: "Bagan Akun (COA)",
        href: "/admin/data/chart_of_account",
        url: "/admin/data/chart_of_account",
        icon: ChartOfAccountIcon,
        roles: allAdminRoles,
      },
      {
        label: "Program DIPA",
        name: "Program DIPA",
        href: "/admin/data/budget_program",
        url: "/admin/data/budget_program",
        icon: BudgetProgramIcon,
        roles: allAdminRoles,
      },
      {
        label: "Komitmen Anggaran",
        name: "Komitmen Anggaran",
        href: "/admin/data/budget_commitment",
        url: "/admin/data/budget_commitment",
        icon: BudgetCommitmentIcon,
        roles: allAdminRoles,
      },
      {
        label: "Tagihan Rekanan",
        name: "Tagihan Rekanan",
        href: "/admin/data/invoice",
        url: "/admin/data/invoice",
        icon: InvoiceIcon,
        roles: allAdminRoles,
      },
      {
        label: "Pembayaran SP2D",
        name: "Pembayaran SP2D",
        href: "/admin/data/payment",
        url: "/admin/data/payment",
        icon: PaymentIcon,
        roles: allAdminRoles,
      },
      {
        label: "Jurnal Akuntansi",
        name: "Jurnal Akuntansi",
        href: "/admin/data/journal_entry",
        url: "/admin/data/journal_entry",
        icon: JournalEntryIcon,
        roles: allAdminRoles,
      },
      {
        label: "Total Cost of Ownership",
        name: "Total Cost of Ownership",
        href: "/admin/data/platform_tco",
        url: "/admin/data/platform_tco",
        icon: PlatformTcoIcon,
        roles: allAdminRoles,
      },
    ],
  },
  {
    label: "06. SDM Militer (HCM)",
    title: "06. SDM Militer (HCM)",
    icon: PersonnelGroupIcon,
    items: [
      {
        label: "Pangkat Militer",
        name: "Pangkat Militer",
        href: "/admin/data/military_rank",
        url: "/admin/data/military_rank",
        icon: MilitaryRankIcon,
        roles: allAdminRoles,
      },
      {
        label: "Korps Militer",
        name: "Korps Militer",
        href: "/admin/data/military_corps",
        url: "/admin/data/military_corps",
        icon: MilitaryCorpsIcon,
        roles: allAdminRoles,
      },
      {
        label: "Kualifikasi & Brevet",
        name: "Kualifikasi & Brevet",
        href: "/admin/data/qualification",
        url: "/admin/data/qualification",
        icon: QualificationIcon,
        roles: allAdminRoles,
      },
      {
        label: "Prajurit TNI AL",
        name: "Prajurit TNI AL",
        href: "/admin/data/personnel",
        url: "/admin/data/personnel",
        icon: PersonnelIcon,
        roles: allAdminRoles,
      },
      {
        label: "Awak KRI",
        name: "Awak KRI",
        href: "/admin/data/crew_assignment",
        url: "/admin/data/crew_assignment",
        icon: CrewAssignmentIcon,
        roles: allAdminRoles,
      },
    ],
  },
  {
    label: "07. Pangkalan & Fasilitas",
    title: "07. Pangkalan & Fasilitas",
    icon: BaseFacilityGroupIcon,
    items: [
      {
        label: "Fasilitas Pangkalan",
        name: "Fasilitas Pangkalan",
        href: "/admin/data/base_facility",
        url: "/admin/data/base_facility",
        icon: BaseFacilityIcon,
        roles: allAdminRoles,
      },
      {
        label: "Penjadwalan Sandar",
        name: "Penjadwalan Sandar",
        href: "/admin/data/berth_booking",
        url: "/admin/data/berth_booking",
        icon: BerthBookingIcon,
        roles: allAdminRoles,
      },
      {
        label: "Bunker BBM",
        name: "Bunker BBM",
        href: "/admin/data/fuel_bunker",
        url: "/admin/data/fuel_bunker",
        icon: FuelBunkerIcon,
        roles: allAdminRoles,
      },
    ],
  },
  {
    label: "08. Transportasi Militer",
    title: "08. Transportasi Militer",
    icon: TransportGroupIcon,
    items: [
      {
        label: "Armada Transportasi",
        name: "Armada Transportasi",
        href: "/admin/data/transport_unit",
        url: "/admin/data/transport_unit",
        icon: TransportUnitIcon,
        roles: allAdminRoles,
      },
      {
        label: "Rute Pelayaran",
        name: "Rute Pelayaran",
        href: "/admin/data/route",
        url: "/admin/data/route",
        icon: RouteIcon,
        roles: allAdminRoles,
      },
      {
        label: "Pengiriman Konvoi",
        name: "Pengiriman Konvoi",
        href: "/admin/data/shipment",
        url: "/admin/data/shipment",
        icon: ShipmentIcon,
        roles: allAdminRoles,
      },
    ],
  },
  {
    label: "09. EDRMS Dokumen Digital",
    title: "09. EDRMS Dokumen Digital",
    icon: EdrmsGroupIcon,
    items: [
      {
        label: "Kategori Dokumen",
        name: "Kategori Dokumen",
        href: "/admin/data/document_category",
        url: "/admin/data/document_category",
        icon: DocumentCategoryIcon,
        roles: allAdminRoles,
      },
      {
        label: "Dokumen Militer",
        name: "Dokumen Militer",
        href: "/admin/data/document",
        url: "/admin/data/document",
        icon: DocumentIcon,
        roles: allAdminRoles,
      },
    ],
  },
  {
    label: "10. Komando & Kesiapan",
    title: "10. Komando & Kesiapan",
    icon: CommandReadinessGroupIcon,
    items: [
      {
        label: "Teater Operasi",
        name: "Teater Operasi",
        href: "/admin/data/theater",
        url: "/admin/data/theater",
        icon: TheaterIcon,
        roles: allAdminRoles,
      },
      {
        label: "Misi Tempur",
        name: "Misi Tempur",
        href: "/admin/data/mission",
        url: "/admin/data/mission",
        icon: MissionIcon,
        roles: allAdminRoles,
      },
      {
        label: "Log Harian KRI",
        name: "Log Harian KRI",
        href: "/admin/data/daily_log",
        url: "/admin/data/daily_log",
        icon: DailyLogIcon,
        roles: allAdminRoles,
      },
      {
        label: "Kesiapan KRI",
        name: "Kesiapan KRI",
        href: "/admin/data/readiness_report",
        url: "/admin/data/readiness_report",
        icon: ReadinessReportIcon,
        roles: allAdminRoles,
      },
      {
        label: "Peringatan Dini",
        name: "Peringatan Dini",
        href: "/admin/data/readiness_alert",
        url: "/admin/data/readiness_alert",
        icon: ReadinessAlertIcon,
        roles: allAdminRoles,
      },
    ],
  },
];

export const adminNavigationGroups = ADMIN_NAVIGATION;

export function isAdminRoleId(value: number): value is AdminRoleId {
  return [3, 4, 5, 6].includes(value);
}

export function getAdminNavigation(roleId: number): AdminNavigationGroup[] {
  if (!isAdminRoleId(roleId)) return [];
  return ADMIN_NAVIGATION;
}

export function canAccessAdminPath(roleId: number, pathname: string): boolean {
  if (pathname === "/admin/login") return true;
  if (!isAdminRoleId(roleId)) return false;
  return true;
}
