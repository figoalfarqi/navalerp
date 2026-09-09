import type { FC } from "react";
import type { IconProps } from "./types";

// 11. CUI Bawah Laut (CuiGroupIcon): Gelombang laut & sensor sonar kedalaman bawah laut
export const CuiGroupIcon: FC<IconProps> = ({
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
    <path d="M2 7c2 1.5 4 1.5 6 0s4-1.5 6 0 4 1.5 6 0" />
    <path d="M2 12c2 1.5 4 1.5 6 0s4-1.5 6 0 4 1.5 6 0" />
    <path d="M2 17c2 1.5 4 1.5 6 0s4-1.5 6 0 4 1.5 6 0" />
    <path d="M12 3v13" />
    <circle cx="12" cy="19" r="1.5" fill="currentColor" />
  </svg>
);

// Pusat Komando CUI (CuiCommandCenterIcon): Radar koordinat maritim
export const CuiCommandCenterIcon: FC<IconProps> = ({
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
    <circle cx="12" cy="12" r="9" />
    <path d="M12 3v18" />
    <path d="M3 12h18" />
    <circle cx="12" cy="12" r="2.5" fill="currentColor" />
  </svg>
);

// Aset Bawah Laut (CuiAssetIcon): Kabel & pipa transmisi dasar laut
export const CuiAssetIcon: FC<IconProps> = ({
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
    <path d="M13 10V3L4 14h7v7l9-11h-7z" />
  </svg>
);

// Log Sensor Telemetri CUI (CuiMonitoringLogIcon): Sinyal akustik/seismik
export const CuiMonitoringLogIcon: FC<IconProps> = ({
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
    <path d="M3 12h4l3 8 4-16 3 8h4" />
  </svg>
);

// Peringatan Anomali CUI (CuiAlertIcon): Indikator bahaya segitiga militer
export const CuiAlertIcon: FC<IconProps> = ({
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

// Inspeksi Bawah Laut (CuiInspectionIcon): Verifikasi ROV / log checklist
export const CuiInspectionIcon: FC<IconProps> = ({
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
    <path d="M9 5H7a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2" />
    <rect width="8" height="4" x="8" y="3" rx="1" />
    <path d="m9 14 2 2 4-4" />
  </svg>
);

