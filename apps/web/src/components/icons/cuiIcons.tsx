import React from "react";
import type { IconProps } from "./types";

export function CuiGroupIcon({ className = "w-5 h-5", ...props }: IconProps) {
  return (
    <svg className={className} fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.75} {...props}>
      <path strokeLinecap="round" strokeLinejoin="round" d="M2 12c2.5-3 5.5-3 8 0s5.5 3 8 0 4-2 4-2" />
      <path strokeLinecap="round" strokeLinejoin="round" d="M2 17c2.5-3 5.5-3 8 0s5.5 3 8 0 4-2 4-2" />
      <path strokeLinecap="round" strokeLinejoin="round" d="M12 3v6m0 0l-3-3m3 3l3-3" />
      <circle cx="12" cy="9" r="1.5" fill="currentColor" />
    </svg>
  );
}

export function CuiAssetIcon({ className = "w-5 h-5", ...props }: IconProps) {
  return (
    <svg className={className} fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.75} {...props}>
      <path strokeLinecap="round" strokeLinejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z" />
    </svg>
  );
}

export function CuiMonitoringLogIcon({ className = "w-5 h-5", ...props }: IconProps) {
  return (
    <svg className={className} fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.75} {...props}>
      <path strokeLinecap="round" strokeLinejoin="round" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
    </svg>
  );
}

export function CuiAlertIcon({ className = "w-5 h-5", ...props }: IconProps) {
  return (
    <svg className={className} fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.75} {...props}>
      <path strokeLinecap="round" strokeLinejoin="round" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
    </svg>
  );
}

export function CuiInspectionIcon({ className = "w-5 h-5", ...props }: IconProps) {
  return (
    <svg className={className} fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.75} {...props}>
      <path strokeLinecap="round" strokeLinejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4" />
    </svg>
  );
}

export function CuiCommandCenterIcon({ className = "w-5 h-5", ...props }: IconProps) {
  return (
    <svg className={className} fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.75} {...props}>
      <circle cx="12" cy="12" r="9" />
      <path strokeLinecap="round" strokeLinejoin="round" d="M12 3v18M3 12h18" />
      <circle cx="12" cy="12" r="3" fill="currentColor" />
    </svg>
  );
}

