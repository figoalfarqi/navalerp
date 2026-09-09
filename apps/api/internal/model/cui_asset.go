package model

import (
	"time"
)

type CuiAsset struct {
	CuiAssetId string `json:"cui_asset_id"`
	AssetCode string `json:"asset_code"`
	AssetName string `json:"asset_name"`
	AssetType string `json:"asset_type"`
	OperatorName string `json:"operator_name"`
	TheaterId *string `json:"theater_id,omitempty"`
	DepthMeters *float64 `json:"depth_meters,omitempty"`
	LengthKm *float64 `json:"length_km,omitempty"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	StartCoordinates *string `json:"start_coordinates,omitempty"`
	EndCoordinates *string `json:"end_coordinates,omitempty"`
	Status string `json:"status"`
	HealthScore int `json:"health_score"`
	ProtectionPriority string `json:"protection_priority"`
	LastInspectedAt *time.Time `json:"last_inspected_at,omitempty"`
	NextInspectionDue *time.Time `json:"next_inspection_due,omitempty"`
	Notes *string `json:"notes,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	MonitoringLogs []MonitoringLogs `json:"monitoringlogs,omitempty"`
	Alerts []Alerts `json:"alerts,omitempty"`
	Inspections []Inspections `json:"inspections,omitempty"`
}

type MonitoringLogs struct {
	LogId string `json:"log_id"`
	CuiAssetId string `json:"cui_asset_id"`
	SensorCode string `json:"sensor_code"`
	SensorType string `json:"sensor_type"`
	LogTime time.Time `json:"log_time"`
	MetricValue float64 `json:"metric_value"`
	MetricUnit string `json:"metric_unit"`
	Status string `json:"status"`
	VesselProximityMmsi *string `json:"vessel_proximity_mmsi,omitempty"`
	AnomalyScore *float64 `json:"anomaly_score,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Alerts struct {
	AlertId string `json:"alert_id"`
	AlertCode string `json:"alert_code"`
	CuiAssetId string `json:"cui_asset_id"`
	AlertType string `json:"alert_type"`
	Severity string `json:"severity"`
	DetectedAt time.Time `json:"detected_at"`
	AssignedShipId *string `json:"assigned_ship_id,omitempty"`
	Status string `json:"status"`
	AiConfidence float64 `json:"ai_confidence"`
	RecommendedAction *string `json:"recommended_action,omitempty"`
	ResolutionNotes *string `json:"resolution_notes,omitempty"`
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Inspections struct {
	InspectionId string `json:"inspection_id"`
	InspectionNumber string `json:"inspection_number"`
	CuiAssetId string `json:"cui_asset_id"`
	ShipId *string `json:"ship_id,omitempty"`
	InspectionDate time.Time `json:"inspection_date"`
	InspectorOfficerId *string `json:"inspector_officer_id,omitempty"`
	Method string `json:"method"`
	ConditionRating string `json:"condition_rating"`
	Findings *string `json:"findings,omitempty"`
	RemedialActionRequired bool `json:"remedial_action_required"`
	NextInspectionDate *time.Time `json:"next_inspection_date,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CuiAssetRequest struct {
	AssetCode *string `json:"asset_code"`
	AssetName *string `json:"asset_name"`
	AssetType *string `json:"asset_type"`
	OperatorName *string `json:"operator_name"`
	TheaterId *string `json:"theater_id"`
	DepthMeters *float64 `json:"depth_meters"`
	LengthKm *float64 `json:"length_km"`
	Latitude *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	StartCoordinates *string `json:"start_coordinates"`
	EndCoordinates *string `json:"end_coordinates"`
	Status *string `json:"status"`
	HealthScore *int `json:"health_score"`
	ProtectionPriority *string `json:"protection_priority"`
	LastInspectedAt *time.Time `json:"last_inspected_at"`
	NextInspectionDue *time.Time `json:"next_inspection_due"`
	Notes *string `json:"notes"`
	MonitoringLogs []MonitoringLogs `json:"monitoringlogs"`
	Alerts []Alerts `json:"alerts"`
	Inspections []Inspections `json:"inspections"`
}
