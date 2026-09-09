import type { FC } from "react";
import type { IconProps } from "./types";

// Pangkat Militer (military_rank): Military chevrons insignia
export const MilitaryRankIcon: FC<IconProps> = ({
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
    <polyline points="6 9 12 4 18 9" />
    <polyline points="6 14 12 9 18 14" />
    <polyline points="6 19 12 14 18 19" />
  </svg>
);

// Korps Militer (military_corps): Naval corps anchor emblem
export const MilitaryCorpsIcon: FC<IconProps> = ({
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
    <circle cx="12" cy="5" r="3" />
    <line x1="12" y1="22" x2="12" y2="8" />
    <path d="M5 12H2a10 10 0 0 0 20 0h-3" />
    <line x1="8" y1="12" x2="16" y2="12" />
  </svg>
);

// Kualifikasi & Brevet (qualification): Qualification wings / medal
export const QualificationIcon: FC<IconProps> = ({
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
    <path d="M12 15l-3 6 4-2 4 2-3-6" />
    <circle cx="12" cy="9" r="6" />
    <polyline points="12 6 12 9 14 10" />
  </svg>
);

// Prajurit TNI AL (personnel): Sailor officer profile
export const PersonnelIcon: FC<IconProps> = ({
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
    <path d="M18 21a6 6 0 0 0-12 0" />
    <circle cx="12" cy="11" r="4" />
    <path d="M6 7h12l-2-4H8L6 7Z" />
  </svg>
);

// Awak KRI (crew_assignment): Crew roster / assignment list
export const CrewAssignmentIcon: FC<IconProps> = ({
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
    <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2" />
    <circle cx="9" cy="7" r="4" />
    <path d="M22 21v-2a4 4 0 0 0-3-3.87" />
    <path d="M16 3.13a4 4 0 0 1 0 7.75" />
  </svg>
);

