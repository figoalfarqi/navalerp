export const DeliveryOrderStatusTypeMap: Record<
  number,
  { value: string; color: string; bgColor: string }
> = {
  "-1": { value: "-", color: "#6B7280", bgColor: "#F3F4F6" }, // gray soft

  0: { value: "Pending", color: "#374151", bgColor: "#E5E7EB" }, // gray
  1: { value: "Assigned", color: "#1D4ED8", bgColor: "#DBEAFE" }, // blue
  2: { value: "Going to", color: "#1E40AF", bgColor: "#E0E7FF" }, // indigo
  3: { value: "Arrived at", color: "#0369A1", bgColor: "#E0F2FE" }, // sky
  4: { value: "Loading", color: "#0E7490", bgColor: "#CFFAFE" }, // cyan
  5: { value: "Loaded", color: "#047857", bgColor: "#D1FAE5" }, // emerald
  6: { value: "Going to", color: "#15803D", bgColor: "#DCFCE7" }, // green
  7: { value: "Arrived at", color: "#4D7C0F", bgColor: "#ECFCCB" }, // lime
  8: { value: "Unloading", color: "#A16207", bgColor: "#FEF9C3" }, // yellow
  9: { value: "Unloaded", color: "#B45309", bgColor: "#FFEDD5" }, // amber
  10: { value: "Completed", color: "#166534", bgColor: "#DCFCE7" }, // green dark
  11: { value: "Canceled", color: "#B91C1C", bgColor: "#FEE2E2" }, // red
};

export function transformDeliveryOrderStatusTypeMap() {
  return Object.entries(DeliveryOrderStatusTypeMap)
    .filter(([key]) => key !== "-1")
    .map(([key, item]) => ({
      value: Number(key),
      label: item.value,
    }));
}
