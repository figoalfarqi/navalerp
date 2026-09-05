package model

import (
	"time"
)

type TransportUnit struct {
	TransportUnitId string `json:"transport_unit_id"`
	UnitCode string `json:"unit_code"`
	UnitName string `json:"unit_name"`
	TransportType string `json:"transport_type"`
	CargoCapacityTons float64 `json:"cargo_capacity_tons"`
	FuelCapacityLiters *float64 `json:"fuel_capacity_liters,omitempty"`
	OperatingUnitId string `json:"operating_unit_id"`
	Status *string `json:"status,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type TransportUnitRequest struct {
	UnitCode *string `json:"unit_code"`
	UnitName *string `json:"unit_name"`
	TransportType *string `json:"transport_type"`
	CargoCapacityTons *float64 `json:"cargo_capacity_tons"`
	FuelCapacityLiters *float64 `json:"fuel_capacity_liters"`
	OperatingUnitId *string `json:"operating_unit_id"`
	Status *string `json:"status"`
}
