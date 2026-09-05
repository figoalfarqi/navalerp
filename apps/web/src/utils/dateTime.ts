// export function formatDateTime(dateString: string) {
//   const date = new Date(dateString);
//   return new Intl.DateTimeFormat("id-ID", {
//     day: "2-digit",
//     month: "2-digit",
//     year: "numeric",
//     hour: "2-digit",
//     minute: "2-digit",
//     hour12: false,
//     timeZone: "Asia/Jakarta", // biar sesuai WIB
//   }).format(date);
// }



export const formatDateTime = (value: string | Date) => {
  if (!value) return "";
  const d = new Date(value);

  const dd = String(d.getDate()).padStart(2, "0");
  const mm = String(d.getMonth() + 1).padStart(2, "0");
  const yyyy = d.getFullYear();

  const hh = String(d.getHours()).padStart(2, "0");
  const min = String(d.getMinutes()).padStart(2, "0");

  return `${dd}/${mm}/${yyyy} ${hh}:${min}`;
};

export function formatDate(dateString: string | Date) {
  const d = new Date(dateString);
  const dd = String(d.getDate()).padStart(2, "0");
  const mm = String(d.getMonth() + 1).padStart(2, "0");
  const yyyy = d.getFullYear();
  return `${dd}/${mm}/${yyyy}`;
}

export function formatTime(dateString: string) {
  const d = new Date(dateString);
  const hh = String(d.getHours()).padStart(2, "0");
  const min = String(d.getMinutes()).padStart(2, "0");
  return `${hh}:${min}`;
}


// utils/dateFormat.ts
export function pad(num: number): string {
  return num.toString().padStart(2, "0");
}

// 📆 1️⃣ ke format ISO-like tanpa geser timezone
export function toLocalISOString(date: Date): string {
  const year = date.getFullYear();
  const month = pad(date.getMonth() + 1);
  const day = pad(date.getDate());
  const hours = pad(date.getHours());
  const minutes = pad(date.getMinutes());
  const seconds = pad(date.getSeconds());
  const millis = String(date.getMilliseconds()).padStart(3, "0");
  return `${year}-${month}-${day}T${hours}:${minutes}:${seconds}.${millis}`;
}

// 📅 2️⃣ ke format tanggal saja
export function toLocalDateString(date: Date): string {
  const year = date.getFullYear();
  const month = pad(date.getMonth() + 1);
  const day = pad(date.getDate());
  return `${year}-${month}-${day}`;
}

// ⏰ 3️⃣ ke format waktu saja
export function toLocalTimeString(date: Date): string {
  const hours = pad(date.getHours());
  const minutes = pad(date.getMinutes());
  const seconds = pad(date.getSeconds());
  return `${hours}:${minutes}:${seconds}`;
}


// utils/duration.ts
export function minutesToDuration(totalMinutes: number) {
  const days = Math.floor(totalMinutes / (24 * 60));
  const hours = Math.floor((totalMinutes % (24 * 60)) / 60);
  const minutes = totalMinutes % 60;

  return { days, hours, minutes };
}

export function durationToMinutes(days: number, hours: number, minutes: number) {
  return days * 24 * 60 + hours * 60 + minutes;
}

export function formatDuration(days: number, hours: number, minutes: number) {
  const parts: string[] = [];

  if (days > 0) parts.push(`${days} hari`);
  if (hours > 0) parts.push(`${hours} jam`);
  if (minutes > 0 || parts.length === 0) parts.push(`${minutes} menit`);

  return parts.join(" ");
}

export function addDays(date: Date, days: number) {
  const d = new Date(date);
  d.setDate(d.getDate() + days);
  return d;
};

export const addMinutes = (date: Date, minutes: number) => {
  const d = new Date(date);
  d.setMinutes(d.getMinutes() + minutes);
  return d.toISOString();
};

const ZERO_DATE = "0001-01-01";

export const getDayStatus = (completedAt?: string) => {
  if (!completedAt) return "default";

  if (completedAt.startsWith(ZERO_DATE)) return "default";

  const completedDate = new Date(completedAt);
  if (isNaN(completedDate.getTime())) return "default";

  const today = new Date();

  // normalize ke tanggal saja (hapus jam)
  const todayDate = new Date(
    today.getFullYear(),
    today.getMonth(),
    today.getDate()
  );

  const completedOnlyDate = new Date(
    completedDate.getFullYear(),
    completedDate.getMonth(),
    completedDate.getDate()
  );

  const diffDays =
    (completedOnlyDate.getTime() - todayDate.getTime()) /
    (1000 * 60 * 60 * 24);

  if (diffDays < 0) return "overdue";
  if (diffDays === 0) return "today";
  if (diffDays === 1) return "tomorrow";

  return "default";
};
