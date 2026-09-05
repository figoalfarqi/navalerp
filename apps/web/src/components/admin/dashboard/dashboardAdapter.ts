import type {
  AdminDailyMetric,
  AdminDashboardData,
  AdminDashboardSummary,
  AdminProjectComparison,
  AdminReportRow,
} from "@/types/adminDashboard";

type UnknownRecord = Record<string, unknown>;

function asRecord(value: unknown): UnknownRecord {
  return value && typeof value === "object" && !Array.isArray(value)
    ? (value as UnknownRecord)
    : {};
}

function asArray(value: unknown): UnknownRecord[] {
  return Array.isArray(value) ? value.map(asRecord) : [];
}

function firstValue(record: UnknownRecord, keys: string[]): unknown {
  for (const key of keys) {
    if (record[key] !== undefined && record[key] !== null) return record[key];
  }
  return undefined;
}

function numberValue(record: UnknownRecord, keys: string[]): number {
  const parsed = Number(firstValue(record, keys) ?? 0);
  return Number.isFinite(parsed) ? parsed : 0;
}

function stringValue(record: UnknownRecord, keys: string[]): string {
  const value = firstValue(record, keys);
  return value === undefined || value === null ? "" : String(value);
}

function normalizeSummary(value: unknown): AdminDashboardSummary {
  const record = asRecord(value);
  return {
    project_count: numberValue(record, ["project_count", "total_projects"]),
    transport_count: numberValue(record, [
      "transport_count",
      "total_transport",
      "total_transports",
    ]),
    completed_transport_count: numberValue(record, [
      "completed_transport_count",
      "completed_transports",
      "total_completed",
    ]),
    volume_cubic: numberValue(record, [
      "volume_cubic",
      "total_volume_cubic",
      "total_volume",
    ]),
    weight_ton: numberValue(record, [
      "weight_ton",
      "total_weight_ton",
      "total_weight",
    ]),
    total_income: numberValue(record, ["total_income", "income"]),
    total_expense: numberValue(record, ["total_expense", "expense"]),
    net_profit: numberValue(record, ["net_profit", "profit"]),
  };
}

function normalizeComparison(record: UnknownRecord): AdminProjectComparison {
  const summary = normalizeSummary(record);
  return {
    project_id: numberValue(record, ["project_id"]),
    project_name:
      stringValue(record, ["project_name", "name"]) || "Tanpa nama",
    project_code: stringValue(record, ["project_code", "code"]) || undefined,
    ...summary,
  };
}

function normalizeDaily(record: UnknownRecord): AdminDailyMetric {
  const rawDate = stringValue(record, ["date", "day", "metric_date"]);
  const date = rawDate || "-";
  const parsed = new Date(rawDate);
  const label = Number.isNaN(parsed.getTime())
    ? date
    : parsed.toLocaleDateString("id-ID", {
        day: "2-digit",
        month: "short",
      });
  return {
    date,
    label,
    ...normalizeSummary(record),
  };
}

export function normalizeDashboardResponse(value: unknown): AdminDashboardData {
  const root = asRecord(value);
  const nested = asRecord(root.data);
  const source = Object.keys(nested).length > 0 ? nested : root;
  const projects = asArray(
    firstValue(source, [
      "projects",
      "project_comparison",
      "project_comparisons",
      "comparison",
    ]),
  ).map(normalizeComparison);
  const daily = asArray(
    firstValue(source, ["daily", "daily_metrics", "daily_series", "timeline"]),
  ).map(normalizeDaily);
  const summary = normalizeSummary(
    firstValue(source, ["summary", "totals"]) ?? source,
  );

  return { summary, projects, daily };
}

export function normalizeReportRows(value: unknown): {
  summary: AdminDashboardSummary;
  rows: AdminReportRow[];
} {
  const root = asRecord(value);
  const nested = asRecord(root.data);
  const source = Object.keys(nested).length > 0 ? nested : root;
  const rawRows = asArray(
    firstValue(source, ["items", "rows", "records", "projects", "report"]),
  );
  const rows = rawRows.map((record) => {
    const metrics = normalizeSummary(record);
    return {
      project_id: numberValue(record, ["project_id"]) || undefined,
      project_code:
        stringValue(record, ["project_code", "code"]) || undefined,
      project_name:
        stringValue(record, ["project_name", "name"]) || "Semua Project",
      period_label:
        stringValue(record, [
          "period_label",
          "period",
          "week_label",
          "month_label",
          "date",
        ]) || "-",
      ...metrics,
    };
  });

  return {
    summary: normalizeSummary(
      firstValue(source, ["summary", "totals"]) ?? source,
    ),
    rows,
  };
}
