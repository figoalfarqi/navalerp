package model

import (
	"time"
)

type Ship struct {
	ShipId                   string     `json:"ship_id"`
	ClassId                  string     `json:"class_id"`
	ClassName                *string    `json:"class_name,omitempty"`
	AssignedUnitId           string     `json:"assigned_unit_id"`
	UnitName                 *string    `json:"unit_name,omitempty"`
	HullNumber               string     `json:"hull_number"`
	ShipName                 string     `json:"ship_name"`
	CallSign                 *string    `json:"call_sign,omitempty"`
	CommissionDate           *time.Time `json:"commission_date,omitempty"`
	HomePort                 *string    `json:"home_port,omitempty"`
	LengthM                  *float64   `json:"length_m,omitempty"`
	BeamM                    *float64   `json:"beam_m,omitempty"`
	DraftM                   *float64   `json:"draft_m,omitempty"`
	DisplacementTons         *float64   `json:"displacement_tons,omitempty"`
	MaxSpeedKnots            *float64   `json:"max_speed_knots,omitempty"`
	CruiseRangeNm            *float64   `json:"cruise_range_nm,omitempty"`
	CrewCapacity             *int       `json:"crew_capacity,omitempty"`
	FuelCapacityLiters       *float64   `json:"fuel_capacity_liters,omitempty"`
	FreshWaterCapacityLiters *float64   `json:"fresh_water_capacity_liters,omitempty"`
	Status                   *string    `json:"status,omitempty"`
	CurrentReadinessStatus   *string    `json:"current_readiness_status,omitempty"`
	CreatedBy                *string    `json:"created_by,omitempty"`
	UpdatedBy                *string    `json:"updated_by,omitempty"`
	DeletedBy                *string    `json:"deleted_by"`
	CreatedAt                time.Time  `json:"created_at"`
	UpdatedAt                *time.Time `json:"updated_at,omitempty"`
	DeletedAt                *time.Time `json:"deleted_at"`
}

type ShipRequest struct {
	ClassId                  *string    `json:"class_id"`
	AssignedUnitId           *string    `json:"assigned_unit_id"`
	HullNumber               *string    `json:"hull_number"`
	ShipName                 *string    `json:"ship_name"`
	CallSign                 *string    `json:"call_sign"`
	CommissionDate           *time.Time `json:"commission_date"`
	HomePort                 *string    `json:"home_port"`
	LengthM                  *float64   `json:"length_m"`
	BeamM                    *float64   `json:"beam_m"`
	DraftM                   *float64   `json:"draft_m"`
	DisplacementTons         *float64   `json:"displacement_tons"`
	MaxSpeedKnots            *float64   `json:"max_speed_knots"`
	CruiseRangeNm            *float64   `json:"cruise_range_nm"`
	CrewCapacity             *int       `json:"crew_capacity"`
	FuelCapacityLiters       *float64   `json:"fuel_capacity_liters"`
	FreshWaterCapacityLiters *float64   `json:"fresh_water_capacity_liters"`
	Status                   *string    `json:"status"`
	CurrentReadinessStatus   *string    `json:"current_readiness_status"`
}
