package model

import (
	"time"
)

type ShipSystem struct {
	SystemId string `json:"system_id"`
	ShipId string `json:"ship_id"`
	ShipName *string `json:"ship_name,omitempty"`
	ParentSystemId *string `json:"parent_system_id,omitempty"`
	SystemCode string `json:"system_code"`
	SystemName string `json:"system_name"`
	SystemCategory string `json:"system_category"`
	SystemLevel string `json:"system_level"`
	Description *string `json:"description,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ShipSystemRequest struct {
	ShipId *string `json:"ship_id"`
	ParentSystemId *string `json:"parent_system_id"`
	SystemCode *string `json:"system_code"`
	SystemName *string `json:"system_name"`
	SystemCategory *string `json:"system_category"`
	SystemLevel *string `json:"system_level"`
	Description *string `json:"description"`
}
