package model

import (
	"time"
)

type ItemInstance struct {
	InstanceId string `json:"instance_id"`
	WarehouseId string `json:"warehouse_id"`
	LocationId *string `json:"location_id,omitempty"`
	MaterialId string `json:"material_id"`
	BatchNumber *string `json:"batch_number,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	LotNumber *string `json:"lot_number,omitempty"`
	ExpiryDate *time.Time `json:"expiry_date,omitempty"`
	ManufacturedDate *time.Time `json:"manufactured_date,omitempty"`
	Condition *string `json:"condition,omitempty"`
	InspectionDueDate *time.Time `json:"inspection_due_date,omitempty"`
	Quantity float64 `json:"quantity"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ItemInstanceRequest struct {
	WarehouseId *string `json:"warehouse_id"`
	LocationId *string `json:"location_id"`
	MaterialId *string `json:"material_id"`
	BatchNumber *string `json:"batch_number"`
	SerialNumber *string `json:"serial_number"`
	LotNumber *string `json:"lot_number"`
	ExpiryDate *time.Time `json:"expiry_date"`
	ManufacturedDate *time.Time `json:"manufactured_date"`
	Condition *string `json:"condition"`
	InspectionDueDate *time.Time `json:"inspection_due_date"`
	Quantity *float64 `json:"quantity"`
}
