import type { FC } from "react";
import type { IconProps } from "./types";

// Satuan Kerja (org_unit): Hierarchy / Organizational branching structure
export const OrgUnitIcon: FC<IconProps> = ({
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
    <rect x="9" y="3" width="6" height="4" rx="1" />
    <rect x="3" y="15" width="6" height="4" rx="1" />
    <rect x="15" y="15" width="6" height="4" rx="1" />
    <path d="M12 7v4" />
    <path d="M6 11h12" />
    <path d="M6 11v4" />
    <path d="M18 11v4" />
  </svg>
);

// Pengguna Sistem (sys_user): User account with access key badge
export const SysUserIcon: FC<IconProps> = ({
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
    <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
    <circle cx="12" cy="7" r="4" />
    <path d="M16 11l2 2 4-4" />
  </svg>
);

// Audit Log (audit_log): Security audit log with fingerprint/activity record
export const AuditLogIcon: FC<IconProps> = ({
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
    <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" />
    <path d="M9 12l2 2 4-4" />
  </svg>
);

