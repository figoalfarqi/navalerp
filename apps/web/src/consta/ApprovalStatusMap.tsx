export const ApprovalStatusMap: Record<
  number,
  { value: string; color: string; bgColor: string }
> = {
  // PENDIG
  0: { value: "Pending", color: "#374151", bgColor: "#E5E7EB" }, // grey
  // APPROVE
  1: { value: "Approve", color: "#166534", bgColor: "#DCFCE7" }, // green
  4: { value: "Approve", color: "#14532D", bgColor: "#BBF7D0" }, // green dark

  // PENDING
  2: { value: "Pending", color: "#B45309", bgColor: "#FFEDD5" }, // amber
  5: { value: "Pending", color: "#92400E", bgColor: "#FED7AA" }, // amber dark

  // REJECT
  3: { value: "Reject", color: "#B91C1C", bgColor: "#FEE2E2" }, // red
  6: { value: "Reject", color: "#7F1D1D", bgColor: "#FECACA" }, // red dark
};
