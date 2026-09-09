package model

import (
	"time"
)

type CuiAlert struct {
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

type CuiAlertRequest struct {
	AlertCode *string `json:"alert_code"`
	CuiAssetId *string `json:"cui_asset_id"`
	AlertType *string `json:"alert_type"`
	Severity *string `json:"severity"`
	DetectedAt *time.Time `json:"detected_at"`
	AssignedShipId *string `json:"assigned_ship_id"`
	Status *string `json:"status"`
	AiConfidence *float64 `json:"ai_confidence"`
	RecommendedAction *string `json:"recommended_action"`
	ResolutionNotes *string `json:"resolution_notes"`
	ResolvedAt *time.Time `json:"resolved_at"`
}
