package model

import (
	"time"
)

type ReadinessReport struct {
	SnapshotId string `json:"snapshot_id"`
	ShipId string `json:"ship_id"`
	ShipName *string `json:"ship_name,omitempty"`
	SnapshotTimestamp time.Time `json:"snapshot_timestamp"`
	ReadinessCategory string `json:"readiness_category"`
	MroReadinessScore float64 `json:"mro_readiness_score"`
	PersonnelManningScore float64 `json:"personnel_manning_score"`
	LogisticsSupplyScore float64 `json:"logistics_supply_score"`
	CompositeReadinessIndex *float64 `json:"composite_readiness_index,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type ReadinessReportRequest struct {
	ShipId *string `json:"ship_id"`
	SnapshotTimestamp *time.Time `json:"snapshot_timestamp"`
	ReadinessCategory *string `json:"readiness_category"`
	MroReadinessScore *float64 `json:"mro_readiness_score"`
	PersonnelManningScore *float64 `json:"personnel_manning_score"`
	LogisticsSupplyScore *float64 `json:"logistics_supply_score"`
	Remarks *string `json:"remarks"`
}
