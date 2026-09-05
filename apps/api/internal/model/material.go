package model

import (
	"time"
)

type Material struct {
	MaterialId string `json:"material_id"`
	MaterialCode string `json:"material_code"`
	Nsn *string `json:"nsn,omitempty"`
	PartNumber *string `json:"part_number,omitempty"`
	OemName *string `json:"oem_name,omitempty"`
	MaterialName string `json:"material_name"`
	Category string `json:"category"`
	Uom string `json:"uom"`
	WeightKg *float64 `json:"weight_kg,omitempty"`
	MinStockLevel *float64 `json:"min_stock_level,omitempty"`
	MaxStockLevel *float64 `json:"max_stock_level,omitempty"`
	ReorderPoint *float64 `json:"reorder_point,omitempty"`
	SafetyStock *float64 `json:"safety_stock,omitempty"`
	ShelfLifeDays *int `json:"shelf_life_days,omitempty"`
	IsControlledItem *bool `json:"is_controlled_item,omitempty"`
	UnitPriceIdr *float64 `json:"unit_price_idr,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	EquipmentLinks []MaterialEquipmentLinks `json:"equipmentlinks,omitempty"`
}

type MaterialEquipmentLinks struct {
	LinkId string `json:"link_id"`
	MaterialId string `json:"material_id"`
	SystemId *string `json:"system_id,omitempty"`
	EquipmentId *string `json:"equipment_id,omitempty"`
	IsMandatorySpare *bool `json:"is_mandatory_spare,omitempty"`
	InterchangeabilityCode *string `json:"interchangeability_code,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type MaterialRequest struct {
	MaterialCode *string `json:"material_code"`
	Nsn *string `json:"nsn"`
	PartNumber *string `json:"part_number"`
	OemName *string `json:"oem_name"`
	MaterialName *string `json:"material_name"`
	Category *string `json:"category"`
	Uom *string `json:"uom"`
	WeightKg *float64 `json:"weight_kg"`
	MinStockLevel *float64 `json:"min_stock_level"`
	MaxStockLevel *float64 `json:"max_stock_level"`
	ReorderPoint *float64 `json:"reorder_point"`
	SafetyStock *float64 `json:"safety_stock"`
	ShelfLifeDays *int `json:"shelf_life_days"`
	IsControlledItem *bool `json:"is_controlled_item"`
	UnitPriceIdr *float64 `json:"unit_price_idr"`
	EquipmentLinks []MaterialEquipmentLinks `json:"equipmentlinks"`
}
