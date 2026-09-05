export const transportStatusOptions = [
  { label: "1 — Tiba di lokasi asal", value: 1 },
  { label: "2 — Mulai memuat", value: 2 },
  { label: "3 — Selesai memuat", value: 3 },
  { label: "4 — Menuju lokasi tujuan", value: 4 },
  { label: "5 — Tiba di lokasi tujuan", value: 5 },
  { label: "6 — Mulai membongkar", value: 6 },
  { label: "7 — Selesai membongkar", value: 7 },
  { label: "8 — Selesai", value: 8 },
];

export const transportPhotoTypeOptions = [
  { label: "Delivery note", value: 1 },
  { label: "Cargo box / truck", value: 2 },
];

export function getTransportStatusLabel(value: unknown): string {
  const status = transportStatusOptions.find(
    (option) => option.value === Number(value),
  );
  return status?.label ?? "Status belum tersedia";
}

export function getTransportPhotoTypeLabel(value: unknown): string {
  const photoType = transportPhotoTypeOptions.find(
    (option) => option.value === Number(value),
  );
  return photoType?.label ?? "Foto transport";
}
