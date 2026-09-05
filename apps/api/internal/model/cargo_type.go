package model

import (
	"time"

	"github.com/figoalfarqi/navalerp/internal/helper"
)

type CargoType struct {
	CargoTypeID          int        `json:"cargo_type_id"`
	CargoTypeName        string     `json:"cargo_type_name"`
	CargoTypeDescription *string    `json:"cargo_type_description,omitempty"`
	CargoTypeGrade       string     `json:"cargo_type_grade"`
	IsActive             int        `json:"is_active"`
	CreatedBy            int        `json:"created_by"`
	UpdatedBy            int        `json:"updated_by"`
	DeletedBy            *int       `json:"deleted_by,omitempty"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	StockpileCargos []StockpileCargo `json:"stockpile_cargos,omitempty"`
}

type CargoTypeNullable struct {
	CargoTypeID          *int       `json:"cargo_type_id"`
	CargoTypeName        *string    `json:"cargo_type_name"`
	CargoTypeDescription *string    `json:"cargo_type_description,omitempty"`
	CargoTypeGrade       *string    `json:"cargo_type_grade"`
	IsActive             *int       `json:"is_active"`
	CreatedBy            *int       `json:"created_by"`
	UpdatedBy            *int       `json:"updated_by"`
	DeletedBy            *int       `json:"deleted_by,omitempty"`
	CreatedAt            *time.Time `json:"created_at"`
	UpdatedAt            *time.Time `json:"updated_at"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	StockpileCargos []StockpileCargoNullable `json:"stockpile_cargos,omitempty"`
}

type CargoTypeRequest struct {
	CargoTypeName        string  `json:"cargo_type_name" validate:"required,min=2,max=100"`
	CargoTypeDescription *string `json:"cargo_type_description,omitempty"`
	CargoTypeGrade       string  `json:"cargo_type_grade" validate:"required,min=1,max=100"`
	IsActive             *int    `json:"is_active,omitempty"`
}

type CargoTypeResponse struct {
	CargoTypeID          int        `json:"cargo_type_id"`
	CargoTypeName        string     `json:"cargo_type_name"`
	CargoTypeDescription *string    `json:"cargo_type_description,omitempty"`
	CargoTypeGrade       string     `json:"cargo_type_grade"`
	IsActive             int        `json:"is_active"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty"`

	StockpileCargos []StockpileCargoResponse `json:"stockpile_cargos,omitempty"`
}

func (ct *CargoTypeNullable) ToNotNullable() *CargoType {
	if ct == nil {
		return nil
	}

	createdAt := time.Time{}
	if ct.CreatedAt != nil {
		createdAt = *ct.CreatedAt
	}

	updatedAt := time.Time{}
	if ct.UpdatedAt != nil {
		updatedAt = *ct.UpdatedAt
	}

	stockpileCargos := make([]StockpileCargo, 0)
	for _, sc := range ct.StockpileCargos {
		if converted := sc.ToNotNullable(); converted != nil {
			stockpileCargos = append(stockpileCargos, *converted)
		}
	}

	return &CargoType{
		CargoTypeID:          helper.DerefInt(ct.CargoTypeID),
		CargoTypeName:        helper.DerefString(ct.CargoTypeName),
		CargoTypeDescription: ct.CargoTypeDescription,
		CargoTypeGrade:       helper.DerefString(ct.CargoTypeGrade),
		IsActive:             helper.DerefInt(ct.IsActive),
		CreatedBy:            helper.DerefInt(ct.CreatedBy),
		UpdatedBy:            helper.DerefInt(ct.UpdatedBy),
		DeletedBy:            ct.DeletedBy,
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
		DeletedAt:            ct.DeletedAt,
		StockpileCargos:      stockpileCargos,
	}
}
