package model

type DashboardSummary struct {
	ProjectCount            int     `json:"project_count"`
	TransportCount          int     `json:"transport_count"`
	CompletedTransportCount int     `json:"completed_transport_count"`
	VolumeCubic             float64 `json:"volume_cubic"`
	WeightTon               float64 `json:"weight_ton"`
	TotalIncome             float64 `json:"total_income"`
	TotalExpense            float64 `json:"total_expense"`
	NetProfit               float64 `json:"net_profit"`
}

type DashboardProject struct {
	ProjectID   int    `json:"project_id"`
	ProjectCode string `json:"project_code"`
	ProjectName string `json:"project_name"`
	DashboardSummary
}

type DashboardDaily struct {
	Date string `json:"date"`
	DashboardSummary
}

type DashboardResponse struct {
	Summary  DashboardSummary   `json:"summary"`
	Projects []DashboardProject `json:"projects"`
	Daily    []DashboardDaily   `json:"daily"`
}
