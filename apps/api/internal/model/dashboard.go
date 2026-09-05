package model

type DashboardSummary struct {
	TotalShips            int     `json:"total_ships"`
	ActiveShips           int     `json:"active_ships"`
	FullyMissionCapable   int     `json:"fully_mission_capable"`
	TotalPersonnel        int     `json:"total_personnel"`
	TotalWarehouses       int     `json:"total_warehouses"`
	ActiveMissions        int     `json:"active_missions"`
	OpenWorkOrders        int     `json:"open_work_orders"`
	OverallReadinessScore float64 `json:"overall_readiness_score"`
}

type DashboardShipReadiness struct {
	ShipID          string  `json:"ship_id"`
	ShipName        string  `json:"ship_name"`
	HullNumber      string  `json:"hull_number"`
	ReadinessStatus string  `json:"readiness_status"`
	CompositeScore  float64 `json:"composite_score"`
}

type DashboardResponse struct {
	Summary DashboardSummary         `json:"summary"`
	Ships   []DashboardShipReadiness `json:"ships"`
}
