import type { FC } from "react";
import type { IconProps } from "./types";

// Kelas Kapal (ship_class): Blueprint vessel classification
export const ShipClassIcon: FC<IconProps> = ({
  size = 16,
  className = "",
  ...props
}) => (
  <svg
    width={size}
    height={size}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="1.8"
    strokeLinecap="round"
    strokeLinejoin="round"
    className={className}
    {...props}
  >
    <path d="M3 6h18v12H3z" />
    <path d="M3 10h18" />
    <path d="M10 6v12" />
    <path d="M7 14l3-4 3 4" />
  </svg>
);

// Kapal Perang KRI (ship): Naval combatant KRI frigate
export const WarshipIcon: FC<IconProps> = ({
  size = 16,
  className = "",
  ...props
}) => (
  <svg
    width={size}
    height={size}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="1.8"
    strokeLinecap="round"
    strokeLinejoin="round"
    className={className}
    {...props}
  >
    <path d="M2 17.5 4 19h16l2-1.5-2.5-6H4.5L2 17.5Z" />
    <path d="M7 11.5V7l4-2v6.5" />
    <path d="M14 11.5V5l3 2v4.5" />
    <path d="M3 21h18" />
  </svg>
);

// Sistem Kapal (ship_system): Interconnected onboard systems & network
export const ShipSystemIcon: FC<IconProps> = ({
  size = 16,
  className = "",
  ...props
}) => (
  <svg
    width={size}
    height={size}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="1.8"
    strokeLinecap="round"
    strokeLinejoin="round"
    className={className}
    {...props}
  >
    <circle cx="12" cy="12" r="3" />
    <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 1 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 1 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 1 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 1 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z" />
  </svg>
);

// Peralatan Mesin (equipment): Marine engines & heavy technical gear
export const EquipmentIcon: FC<IconProps> = ({
  size = 16,
  className = "",
  ...props
}) => (
  <svg
    width={size}
    height={size}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="1.8"
    strokeLinecap="round"
    strokeLinejoin="round"
    className={className}
    {...props}
  >
    <rect x="3" y="7" width="18" height="13" rx="2" />
    <path d="M7 7V4a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v3" />
    <circle cx="12" cy="13.5" r="2.5" />
    <path d="M12 11v-1" />
    <path d="M12 17v-1" />
    <path d="M14.5 13.5h1" />
    <path d="M8.5 13.5h1" />
  </svg>
);

// Jadwal PMS (pm_schedule): Preventive maintenance schedule clock
export const PmScheduleIcon: FC<IconProps> = ({
  size = 16,
  className = "",
  ...props
}) => (
  <svg
    width={size}
    height={size}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="1.8"
    strokeLinecap="round"
    strokeLinejoin="round"
    className={className}
    {...props}
  >
    <rect x="3" y="4" width="18" height="18" rx="2" />
    <line x1="16" y1="2" x2="16" y2="6" />
    <line x1="8" y1="2" x2="8" y2="6" />
    <line x1="3" y1="10" x2="21" y2="10" />
    <circle cx="12" cy="16" r="3" />
    <polyline points="12 14.5 12 16 13.5 16" />
  </svg>
);

// Laporan Kerusakan (failure_report): Technical damage alert
export const FailureReportIcon: FC<IconProps> = ({
  size = 16,
  className = "",
  ...props
}) => (
  <svg
    width={size}
    height={size}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="1.8"
    strokeLinecap="round"
    strokeLinejoin="round"
    className={className}
    {...props}
  >
    <path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3Z" />
    <line x1="12" y1="9" x2="12" y2="13" />
    <line x1="12" y1="17" x2="12.01" y2="17" />
  </svg>
);

// Perintah Kerja MRO (work_order): Technical maintenance work order
export const WorkOrderIcon: FC<IconProps> = ({
  size = 16,
  className = "",
  ...props
}) => (
  <svg
    width={size}
    height={size}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="1.8"
    strokeLinecap="round"
    strokeLinejoin="round"
    className={className}
    {...props}
  >
    <path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2" />
    <rect x="8" y="2" width="8" height="4" rx="1" />
    <path d="m9 14 2 2 4-4" />
  </svg>
);

// Riwayat Docking (docking_record): Drydock basin & ship support
export const DockingRecordIcon: FC<IconProps> = ({
  size = 16,
  className = "",
  ...props
}) => (
  <svg
    width={size}
    height={size}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="1.8"
    strokeLinecap="round"
    strokeLinejoin="round"
    className={className}
    {...props}
  >
    <path d="M3 17h18" />
    <path d="M5 21h14" />
    <path d="M5 17v4" />
    <path d="M12 17v4" />
    <path d="M19 17v4" />
    <path d="M5 13l7-8 7 8" />
    <path d="M12 5v8" />
  </svg>
);

