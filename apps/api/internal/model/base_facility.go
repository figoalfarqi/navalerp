package model

import (
	"time"
)

type BaseFacility struct {
	FacilityId string `json:"facility_id"`
	BaseUnitId string `json:"base_unit_id"`
	FacilityCode string `json:"facility_code"`
	FacilityName string `json:"facility_name"`
	FacilityType string `json:"facility_type"`
	LengthMeters *float64 `json:"length_meters,omitempty"`
	DraftDepthMeters *float64 `json:"draft_depth_meters,omitempty"`
	MaxDisplacementTonnage *float64 `json:"max_displacement_tonnage,omitempty"`
	HasShorePower *bool `json:"has_shore_power,omitempty"`
	HasFreshWater *bool `json:"has_fresh_water,omitempty"`
	HasFuelBunkerLine *bool `json:"has_fuel_bunker_line,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Maintenances []FacilityMaintenances `json:"maintenances,omitempty"`
}

type FacilityMaintenances struct {
	MaintId string `json:"maint_id"`
	FacilityId string `json:"facility_id"`
	MaintenanceType string `json:"maintenance_type"`
	StartDate time.Time `json:"start_date"`
	EndDate *time.Time `json:"end_date,omitempty"`
	Cost *float64 `json:"cost,omitempty"`
	PerformedBy string `json:"performed_by"`
	Status *string `json:"status,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type BaseFacilityRequest struct {
	BaseUnitId *string `json:"base_unit_id"`
	FacilityCode *string `json:"facility_code"`
	FacilityName *string `json:"facility_name"`
	FacilityType *string `json:"facility_type"`
	LengthMeters *float64 `json:"length_meters"`
	DraftDepthMeters *float64 `json:"draft_depth_meters"`
	MaxDisplacementTonnage *float64 `json:"max_displacement_tonnage"`
	HasShorePower *bool `json:"has_shore_power"`
	HasFreshWater *bool `json:"has_fresh_water"`
	HasFuelBunkerLine *bool `json:"has_fuel_bunker_line"`
	Status *string `json:"status"`
	Maintenances []FacilityMaintenances `json:"maintenances"`
}
