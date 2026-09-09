import type { FC } from "react";
import type { IconProps } from "./types";

// 00. Overview Group: Radar / Compass
export const OverviewGroupIcon: FC<IconProps> = ({
  size = 18,
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
    <circle cx="12" cy="12" r="9" />
    <path d="M12 3v18" />
    <path d="M3 12h18" />
    <circle cx="12" cy="12" r="4" />
    <path d="m14.5 9.5-5 5" />
  </svg>
);

// 01. Organisasi & Pengguna: Command HQ / Hierarchy
export const OrganizationGroupIcon: FC<IconProps> = ({
  size = 18,
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
    <path d="M12 2 3 7v6c0 5.25 3.75 9.5 9 10.5 5.25-1 9-5.25 9-10.5V7l-9-5Z" />
    <circle cx="12" cy="9" r="2.5" />
    <path d="M7.5 16c0-2 2-3.5 4.5-3.5s4.5 1.5 4.5 3.5" />
  </svg>
);

// 02. Armada Kapal & MRO: Warship KRI Cruiser
export const FleetGroupIcon: FC<IconProps> = ({
  size = 18,
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
    <path d="M2 17.5 4 19h16l2-1.5-3-6H5l-3 6Z" />
    <path d="M12 5v8" />
    <path d="M9 8h6" />
    <path d="M8 13h8" />
    <path d="M3 21c3-1 6-1 9 0 3-1 6-1 9 0" />
  </svg>
);

// 03. Logistik & Pergudangan: Military Depot & Boxes
export const LogisticsGroupIcon: FC<IconProps> = ({
  size = 18,
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
    <path d="M3 9 12 4l9 5-9 5-9-5Z" />
    <path d="M3 14.5 12 19.5l9-5" />
    <path d="M12 14v5.5" />
    <path d="M3 9v5.5" />
    <path d="M21 9v5.5" />
  </svg>
);

// 04. Pengadaan Pertahanan: Procurement Contract & Shield
export const ProcurementGroupIcon: FC<IconProps> = ({
  size = 18,
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
    <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z" />
    <polyline points="14 2 14 8 20 8" />
    <path d="m9 15 2 2 4-4" />
    <line x1="8" y1="11" x2="13" y2="11" />
  </svg>
);

// 05. Keuangan & Anggaran: Military Finance Vault
export const FinanceGroupIcon: FC<IconProps> = ({
  size = 18,
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
    <rect x="2" y="6" width="20" height="13" rx="2" />
    <circle cx="12" cy="12.5" r="2.5" />
    <path d="M6 10h.01" />
    <path d="M18 15h.01" />
    <path d="M4 6V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v2" />
  </svg>
);

// 06. SDM Militer (HCM): Military Personnel & Ranks
export const PersonnelGroupIcon: FC<IconProps> = ({
  size = 18,
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
    <path d="M16 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" />
    <circle cx="10" cy="7" r="4" />
    <path d="M20 8v6" />
    <path d="M23 11h-6" />
  </svg>
);

// 07. Pangkalan & Fasilitas: Coastal Naval Base / Pier
export const BaseFacilityGroupIcon: FC<IconProps> = ({
  size = 18,
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
    <path d="M3 21h18" />
    <path d="M5 21V9l5-4v16" />
    <path d="M10 9l9 3v9" />
    <line x1="8" y1="13" x2="8" y2="13.01" />
    <line x1="8" y1="17" x2="8" y2="17.01" />
    <line x1="14" y1="15" x2="14" y2="15.01" />
    <line x1="17" y1="15" x2="17" y2="15.01" />
  </svg>
);

// 08. Transportasi Militer: Tactical Transport Convoy
export const TransportGroupIcon: FC<IconProps> = ({
  size = 18,
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
    <path d="M1 4h13v12H1z" />
    <path d="M14 8h4l3 4v4h-7V8z" />
    <circle cx="5.5" cy="18.5" r="2.5" />
    <circle cx="17.5" cy="18.5" r="2.5" />
  </svg>
);

// 09. EDRMS Dokumen Digital: Classified Archive Folder
export const EdrmsGroupIcon: FC<IconProps> = ({
  size = 18,
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
    <path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z" />
    <circle cx="12" cy="14" r="2" />
    <path d="M12 12v-1" />
  </svg>
);

// 10. Komando & Kesiapan: Operations Radar & Target
export const CommandReadinessGroupIcon: FC<IconProps> = ({
  size = 18,
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
    <circle cx="12" cy="12" r="9" />
    <circle cx="12" cy="12" r="5" />
    <line x1="12" y1="2" x2="12" y2="5" />
    <line x1="12" y1="19" x2="12" y2="22" />
    <line x1="2" y1="12" x2="5" y2="12" />
    <line x1="19" y1="12" x2="22" y2="12" />
    <circle cx="12" cy="12" r="1.5" fill="currentColor" />
  </svg>
);

