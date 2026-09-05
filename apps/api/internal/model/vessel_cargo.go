package model

import "time"

type VesselCargo struct {
	VesselCargoID        int        `json:"vessel_cargo_id"`
	VesselID             int        `json:"vessel_id"`
	PortID               int        `json:"port_id"`
	CargoTypeID          int        `json:"cargo_type_id"`
	VoyageNumber         *string    `json:"voyage_number,omitempty"`
	BillOfLadingNumber   *string    `json:"bill_of_lading_number,omitempty"`
	ArrivalAt            *time.Time `json:"arrival_at,omitempty"`
	UnloadingStartedAt   *time.Time `json:"unloading_started_at,omitempty"`
	UnloadingCompletedAt *time.Time `json:"unloading_completed_at,omitempty"`
	ManifestVolumeCubic  *float64   `json:"manifest_volume_cubic,omitempty"`
	ManifestWeightTon    *float64   `json:"manifest_weight_ton,omitempty"`
	IsActive             int        `json:"is_active"`
	Vessel               *Vessel    `json:"vessel,omitempty"`
	Port                 *Port      `json:"port,omitempty"`
	CargoType            *CargoType `json:"cargo_type,omitempty"`
	Audit
}

type VesselCargoRequest struct {
	VesselID             int        `json:"vessel_id" validate:"required,gt=0"`
	PortID               int        `json:"port_id" validate:"required,gt=0"`
	CargoTypeID          int        `json:"cargo_type_id" validate:"required,gt=0"`
	VoyageNumber         *string    `json:"voyage_number,omitempty"`
	BillOfLadingNumber   *string    `json:"bill_of_lading_number,omitempty"`
	ArrivalAt            *time.Time `json:"arrival_at,omitempty"`
	UnloadingStartedAt   *time.Time `json:"unloading_started_at,omitempty"`
	UnloadingCompletedAt *time.Time `json:"unloading_completed_at,omitempty"`
	ManifestVolumeCubic  *float64   `json:"manifest_volume_cubic,omitempty"`
	ManifestWeightTon    *float64   `json:"manifest_weight_ton,omitempty"`
	IsActive             *int       `json:"is_active,omitempty" validate:"omitempty,oneof=0 1"`
}
