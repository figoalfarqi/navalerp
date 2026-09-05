package model

import (
	"time"
)

type StockAdjustment struct {
	AdjustmentId string `json:"adjustment_id"`
	WarehouseId string `json:"warehouse_id"`
	AdjustmentNumber string `json:"adjustment_number"`
	AdjustmentDate time.Time `json:"adjustment_date"`
	ConductedByUserId *string `json:"conducted_by_user_id,omitempty"`
	Reason string `json:"reason"`
	Status *string `json:"status,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Items []StockAdjustmentItems `json:"items,omitempty"`
}

type StockAdjustmentItems struct {
	AdjItemId string `json:"adj_item_id"`
	AdjustmentId string `json:"adjustment_id"`
	MaterialId string `json:"material_id"`
	BookQuantity float64 `json:"book_quantity"`
	PhysicalQuantity float64 `json:"physical_quantity"`
	DifferenceQuantity *float64 `json:"difference_quantity,omitempty"`
	UnitCost *float64 `json:"unit_cost,omitempty"`
	TotalAdjustmentValue *float64 `json:"total_adjustment_value,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type StockAdjustmentRequest struct {
	WarehouseId *string `json:"warehouse_id"`
	AdjustmentNumber *string `json:"adjustment_number"`
	AdjustmentDate *time.Time `json:"adjustment_date"`
	ConductedByUserId *string `json:"conducted_by_user_id"`
	Reason *string `json:"reason"`
	Status *string `json:"status"`
	Remarks *string `json:"remarks"`
	Items []StockAdjustmentItems `json:"items"`
}
