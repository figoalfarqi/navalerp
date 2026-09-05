export interface AdminProjectOption {
  project_id: number;
  project_name: string;
  project_code?: string;
}

export interface AdminDashboardSummary {
  project_count: number;
  transport_count: number;
  completed_transport_count: number;
  volume_cubic: number;
  weight_ton: number;
  total_income: number;
  total_expense: number;
  net_profit: number;
}

export interface AdminProjectComparison {
  project_id: number;
  project_name: string;
  project_code?: string;
  transport_count: number;
  completed_transport_count: number;
  volume_cubic: number;
  weight_ton: number;
  total_income: number;
  total_expense: number;
  net_profit: number;
}

export interface AdminDailyMetric {
  date: string;
  label: string;
  transport_count: number;
  completed_transport_count: number;
  volume_cubic: number;
  weight_ton: number;
  total_income: number;
  total_expense: number;
  net_profit: number;
}

export interface AdminDashboardData {
  summary: AdminDashboardSummary;
  projects: AdminProjectComparison[];
  daily: AdminDailyMetric[];
}

export interface AdminReportRow {
  project_id?: number;
  project_code?: string;
  project_name: string;
  period_label: string;
  transport_count: number;
  completed_transport_count: number;
  volume_cubic: number;
  weight_ton: number;
  total_income: number;
  total_expense: number;
  net_profit: number;
}
