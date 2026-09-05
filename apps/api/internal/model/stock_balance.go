package model

import (
	"time"
)

type StockBalance struct {
	BalanceId string `json:"balance_id"`
	WarehouseId string `json:"warehouse_id"`
	LocationId *string `json:"location_id,omitempty"`
	MaterialId string `json:"material_id"`
	QuantityOnHand float64 `json:"quantity_on_hand"`
	QuantityReserved float64 `json:"quantity_reserved"`
	QuantityInTransit float64 `json:"quantity_in_transit"`
	QuantityAvailable *float64 `json:"quantity_available,omitempty"`
	LastCountDate *time.Time `json:"last_count_date,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type StockBalanceRequest struct {
	WarehouseId *string `json:"warehouse_id"`
	LocationId *string `json:"location_id"`
	MaterialId *string `json:"material_id"`
	QuantityOnHand *float64 `json:"quantity_on_hand"`
	QuantityReserved *float64 `json:"quantity_reserved"`
	QuantityInTransit *float64 `json:"quantity_in_transit"`
	LastCountDate *time.Time `json:"last_count_date"`
}
