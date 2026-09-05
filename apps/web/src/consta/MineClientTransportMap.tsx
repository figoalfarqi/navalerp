export const MineClientTransportStatusMap: Record<
  number,
  {
    label: string;
    textColor: string;
    bgColor: string;
  }
> = {
  1: {
    label: "Arrived at Mine",
    textColor: "#1E40AF",
    bgColor: "#DBEAFE",
  },
  2: {
    label: "Loading at Mine",
    textColor: "#B45309",
    bgColor: "#FEF3C7",
  },
  3: {
    label: "Loaded at Mine",
    textColor: "#065F46",
    bgColor: "#D1FAE5",
  },
  4: {
    label: "Going to Client",
    textColor: "#6B21A8",
    bgColor: "#EDE9FE",
  },
  5: {
    label: "Arrived at Client",
    textColor: "#0E7490",
    bgColor: "#CFFAFE",
  },
  6: {
    label: "Unloading at Client",
    textColor: "#9A3412",
    bgColor: "#FED7AA",
  },
  7: {
    label: "Unloaded at Client",
    textColor: "#166534",
    bgColor: "#DCFCE7",
  },
  8: {
    label: "Completed",
    textColor: "#111827",
    bgColor: "#E5E7EB",
  },
};