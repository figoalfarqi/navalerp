package model

type NavalReportItem struct {
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Status   string  `json:"status"`
	Metric   float64 `json:"metric"`
	Notes    string  `json:"notes"`
}

type ReportResponse struct {
	Period    string            `json:"period"`
	Generated string            `json:"generated"`
	Summary   DashboardSummary  `json:"summary"`
	Items     []NavalReportItem `json:"items"`
}
