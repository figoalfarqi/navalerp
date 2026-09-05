import type { AdminProjectOption } from "@/types/adminDashboard";

export function extractAdminProjectOptions(
  value: unknown,
): AdminProjectOption[] {
  const root =
    value && typeof value === "object"
      ? (value as Record<string, unknown>)
      : {};
  const nested =
    root.data && typeof root.data === "object"
      ? (root.data as Record<string, unknown>)
      : root;
  const candidates = Array.isArray(nested.items)
    ? nested.items
    : Array.isArray(nested.projects)
      ? nested.projects
      : [];

  return candidates
    .map((item) => item as Record<string, unknown>)
    .map((item) => ({
      project_id: Number(item.project_id),
      project_name: String(item.project_name ?? "Tanpa nama"),
      project_code: item.project_code ? String(item.project_code) : undefined,
    }))
    .filter((item) => Number.isFinite(item.project_id));
}
