package model

import (
	"time"
)

type CuiInspection struct {
	InspectionId           string     `json:"inspection_id"`
	InspectionNumber       string     `json:"inspection_number"`
	CuiAssetId             string     `json:"cui_asset_id"`
	ShipId                 *string    `json:"ship_id,omitempty"`
	InspectionDate         time.Time  `json:"inspection_date"`
	InspectorOfficerId     *string    `json:"inspector_officer_id,omitempty"`
	Method                 string     `json:"method"`
	ConditionRating        string     `json:"condition_rating"`
	Findings               *string    `json:"findings,omitempty"`
	RemedialActionRequired bool       `json:"remedial_action_required"`
	NextInspectionDate     *time.Time `json:"next_inspection_date,omitempty"`
	CreatedBy              *string    `json:"created_by,omitempty"`
	UpdatedBy              *string    `json:"updated_by,omitempty"`
	DeletedBy              *string    `json:"deleted_by"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
	DeletedAt              *time.Time `json:"deleted_at"`

	// Joined Display Fields
	AssetName     *string `json:"asset_name,omitempty"`
	AssetCode     *string `json:"asset_code,omitempty"`
	ShipName      *string `json:"ship_name,omitempty"`
	HullNumber    *string `json:"hull_number,omitempty"`
	InspectorName *string `json:"inspector_name,omitempty"`
	InspectorNrp  *string `json:"inspector_nrp,omitempty"`
}

type CuiInspectionRequest struct {
	InspectionNumber       *string    `json:"inspection_number"`
	CuiAssetId             *string    `json:"cui_asset_id"`
	ShipId                 *string    `json:"ship_id"`
	InspectionDate         *time.Time `json:"inspection_date"`
	InspectorOfficerId     *string    `json:"inspector_officer_id"`
	Method                 *string    `json:"method"`
	ConditionRating        *string    `json:"condition_rating"`
	Findings               *string    `json:"findings"`
	RemedialActionRequired *bool      `json:"remedial_action_required"`
	NextInspectionDate     *time.Time `json:"next_inspection_date"`
}
