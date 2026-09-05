package model

import (
	"time"
)

type ShipClass struct {
	ClassId string `json:"class_id"`
	ClassCode string `json:"class_code"`
	ClassName string `json:"class_name"`
	Category string `json:"category"`
	Specifications any `json:"specifications,omitempty"`
	Builder *string `json:"builder,omitempty"`
	TotalBuilt *int `json:"total_built,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ShipClassRequest struct {
	ClassCode *string `json:"class_code"`
	ClassName *string `json:"class_name"`
	Category *string `json:"category"`
	Specifications any `json:"specifications"`
	Builder *string `json:"builder"`
	TotalBuilt *int `json:"total_built"`
}
