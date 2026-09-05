package model

import "time"

type StockpileCargo struct {
	StockpileCargoID    int        `json:"stockpile_cargo_id"`
	StockpileID         int        `json:"stockpile_id"`
	CargoTypeID         int        `json:"cargo_type_id"`
	CapacityVolumeCubic *float64   `json:"capacity_volume_cubic,omitempty"`
	CapacityWeightTon   *float64   `json:"capacity_weight_ton,omitempty"`
	CurrentVolumeCubic  float64    `json:"current_volume_cubic"`
	CurrentWeightTon    float64    `json:"current_weight_ton"`
	IsActive            int        `json:"is_active"`
	Stockpile           *Stockpile `json:"stockpile,omitempty"`
	CargoType           *CargoType `json:"cargo_type,omitempty"`
	Audit
}

type StockpileCargoRequest struct {
	StockpileID         int      `json:"stockpile_id" validate:"required,gt=0"`
	CargoTypeID         int      `json:"cargo_type_id" validate:"required,gt=0"`
	CapacityVolumeCubic *float64 `json:"capacity_volume_cubic,omitempty" validate:"omitempty,gte=0"`
	CapacityWeightTon   *float64 `json:"capacity_weight_ton,omitempty" validate:"omitempty,gte=0"`
	CurrentVolumeCubic  *float64 `json:"current_volume_cubic,omitempty" validate:"omitempty,gte=0"`
	CurrentWeightTon    *float64 `json:"current_weight_ton,omitempty" validate:"omitempty,gte=0"`
	IsActive            *int     `json:"is_active,omitempty" validate:"omitempty,oneof=0 1"`
}

type StockpileCargoNullable struct {
	StockpileCargoID    *int       `json:"stockpile_cargo_id"`
	StockpileID         *int       `json:"stockpile_id"`
	CargoTypeID         *int       `json:"cargo_type_id"`
	CapacityVolumeCubic *float64   `json:"capacity_volume_cubic,omitempty"`
	CapacityWeightTon   *float64   `json:"capacity_weight_ton,omitempty"`
	CurrentVolumeCubic  *float64   `json:"current_volume_cubic"`
	CurrentWeightTon    *float64   `json:"current_weight_ton"`
	IsActive            *int       `json:"is_active"`
	CreatedBy           *int       `json:"created_by"`
	UpdatedBy           *int       `json:"updated_by"`
	DeletedBy           *int       `json:"deleted_by,omitempty"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at,omitempty"`
}

func (value *StockpileCargoNullable) ToNotNullable() *StockpileCargo {
	if value == nil || value.StockpileCargoID == nil {
		return nil
	}
	item := &StockpileCargo{
		StockpileCargoID:    *value.StockpileCargoID,
		CapacityVolumeCubic: value.CapacityVolumeCubic,
		CapacityWeightTon:   value.CapacityWeightTon,
		Audit: Audit{
			DeletedBy: value.DeletedBy,
			DeletedAt: value.DeletedAt,
		},
	}
	if value.StockpileID != nil {
		item.StockpileID = *value.StockpileID
	}
	if value.CargoTypeID != nil {
		item.CargoTypeID = *value.CargoTypeID
	}
	if value.CurrentVolumeCubic != nil {
		item.CurrentVolumeCubic = *value.CurrentVolumeCubic
	}
	if value.CurrentWeightTon != nil {
		item.CurrentWeightTon = *value.CurrentWeightTon
	}
	if value.IsActive != nil {
		item.IsActive = *value.IsActive
	}
	if value.CreatedBy != nil {
		item.CreatedBy = *value.CreatedBy
	}
	if value.UpdatedBy != nil {
		item.UpdatedBy = *value.UpdatedBy
	}
	if value.CreatedAt != nil {
		item.CreatedAt = *value.CreatedAt
	}
	if value.UpdatedAt != nil {
		item.UpdatedAt = *value.UpdatedAt
	}
	return item
}

type StockpileCargoResponse = StockpileCargo
