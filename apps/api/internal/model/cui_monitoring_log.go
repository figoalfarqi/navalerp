package model

import (
	"time"
)

type CuiMonitoringLog struct {
	LogId               string     `json:"log_id"`
	CuiAssetId          string     `json:"cui_asset_id"`
	SensorCode          string     `json:"sensor_code"`
	SensorType          string     `json:"sensor_type"`
	LogTime             time.Time  `json:"log_time"`
	MetricValue         float64    `json:"metric_value"`
	MetricUnit          string     `json:"metric_unit"`
	Status              string     `json:"status"`
	VesselProximityMmsi *string    `json:"vessel_proximity_mmsi,omitempty"`
	AnomalyScore        *float64   `json:"anomaly_score,omitempty"`
	Description         *string    `json:"description,omitempty"`
	CreatedBy           *string    `json:"created_by,omitempty"`
	UpdatedBy           *string    `json:"updated_by,omitempty"`
	DeletedBy           *string    `json:"deleted_by"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
	DeletedAt           *time.Time `json:"deleted_at"`

	// Joined Display Fields
	AssetName *string `json:"asset_name,omitempty"`
	AssetCode *string `json:"asset_code,omitempty"`
}

type CuiMonitoringLogRequest struct {
	CuiAssetId          *string    `json:"cui_asset_id"`
	SensorCode          *string    `json:"sensor_code"`
	SensorType          *string    `json:"sensor_type"`
	LogTime             *time.Time `json:"log_time"`
	MetricValue         *float64   `json:"metric_value"`
	MetricUnit          *string    `json:"metric_unit"`
	Status              *string    `json:"status"`
	VesselProximityMmsi *string    `json:"vessel_proximity_mmsi"`
	AnomalyScore        *float64   `json:"anomaly_score"`
	Description         *string    `json:"description"`
}
