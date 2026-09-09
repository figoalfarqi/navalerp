import type { FC } from "react";
import type { IconProps } from "./types";

// Fasilitas Pangkalan (base_facility): Naval base infrastructure & docks
export const BaseFacilityIcon: FC<IconProps> = ({
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
    <path d="M2 20h20" />
    <path d="M4 20V8l7-4 7 4v12" />
    <path d="M9 13h6" />
    <path d="M9 17h6" />
    <circle cx="11" cy="7" r="1" />
  </svg>
);

// Penjadwalan Sandar (berth_booking): Berth mooring & pier scheduling
export const BerthBookingIcon: FC<IconProps> = ({
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
    <path d="M3 19h18" />
    <path d="M5 19v-6a2 2 0 0 1 2-2h1" />
    <path d="M8 11V7a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4" />
    <path d="M16 11h1a2 2 0 0 1 2 2v6" />
    <circle cx="12" cy="11" r="2" />
  </svg>
);

// Bunker BBM (fuel_bunker): Naval fuel bunker & refueling depot
export const FuelBunkerIcon: FC<IconProps> = ({
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
    <path d="M3 22V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v17" />
    <path d="M15 9h3a2 2 0 0 1 2 2v4a2 2 0 0 0 2 2" />
    <path d="M22 17v-4" />
    <rect x="6" y="6" width="6" height="4" rx="1" />
    <path d="M1 22h16" />
  </svg>
);

