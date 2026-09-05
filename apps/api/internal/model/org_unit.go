package model

import (
	"time"
)

type OrgUnit struct {
	UnitId string `json:"unit_id"`
	ParentUnitId *string `json:"parent_unit_id,omitempty"`
	ParentUnitName *string `json:"parent_unit_name,omitempty"`
	UnitCode string `json:"unit_code"`
	UnitName string `json:"unit_name"`
	UnitType string `json:"unit_type"`
	Description *string `json:"description,omitempty"`
	CommandLevel *int `json:"command_level,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Address *string `json:"address,omitempty"`
	Phone *string `json:"phone,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type OrgUnitRequest struct {
	ParentUnitId *string `json:"parent_unit_id"`
	UnitCode *string `json:"unit_code"`
	UnitName *string `json:"unit_name"`
	UnitType *string `json:"unit_type"`
	Description *string `json:"description"`
	CommandLevel *int `json:"command_level"`
	Latitude *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Address *string `json:"address"`
	Phone *string `json:"phone"`
	IsActive *bool `json:"is_active"`
}
