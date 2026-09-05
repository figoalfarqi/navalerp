package model

type ReportProject struct {
	ProjectID   int    `json:"project_id"`
	ProjectCode string `json:"project_code"`
	ProjectName string `json:"project_name"`
	PeriodLabel string `json:"period_label"`
	DashboardSummary
}

type ReportResponse struct {
	Period                string                        `json:"period"`
	PeriodLabel           string                        `json:"period_label"`
	DateFrom              string                        `json:"date_from"`
	DateTo                string                        `json:"date_to"`
	Summary               DashboardSummary              `json:"summary"`
	Projects              []ReportProject               `json:"projects"`
	Transports            []ProjectTransport            `json:"transports"`
	FinancialTransactions []ProjectFinancialTransaction `json:"financial_transactions"`
}
