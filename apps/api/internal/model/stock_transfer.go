package model

import (
	"time"
)

type StockTransfer struct {
	TransferId          string               `json:"transfer_id"`
	TransferNumber      string               `json:"transfer_number"`
	FromWarehouseId     string               `json:"from_warehouse_id"`
	SourceWarehouseName *string              `json:"source_warehouse_name,omitempty"`
	ToWarehouseId       string               `json:"to_warehouse_id"`
	DestWarehouseName   *string              `json:"dest_warehouse_name,omitempty"`
	MovementType        string               `json:"movement_type"`
	ScheduledDeparture  *time.Time           `json:"scheduled_departure,omitempty"`
	ActualDeparture     *time.Time           `json:"actual_departure,omitempty"`
	ScheduledArrival    *time.Time           `json:"scheduled_arrival,omitempty"`
	ActualArrival       *time.Time           `json:"actual_arrival,omitempty"`
	TransporterUnit     *string              `json:"transporter_unit,omitempty"`
	Status              *string              `json:"status,omitempty"`
	CreatedBy           *string              `json:"created_by,omitempty"`
	UpdatedBy           *string              `json:"updated_by,omitempty"`
	DeletedBy           *string              `json:"deleted_by"`
	CreatedAt           time.Time            `json:"created_at"`
	UpdatedAt           *time.Time           `json:"updated_at,omitempty"`
	DeletedAt           *time.Time           `json:"deleted_at"`
	Items               []StockTransferItems `json:"items,omitempty"`
}

type StockTransferItems struct {
	TransferItemId     string     `json:"transfer_item_id"`
	TransferId         string     `json:"transfer_id"`
	MaterialId         string     `json:"material_id"`
	MaterialName       *string    `json:"material_name,omitempty"`
	MaterialCode       *string    `json:"material_code,omitempty"`
	QuantityShipped    float64    `json:"quantity_shipped"`
	QuantityReceived   *float64   `json:"quantity_received,omitempty"`
	ConditionOnReceipt *string    `json:"condition_on_receipt,omitempty"`
	CreatedBy          *string    `json:"created_by,omitempty"`
	UpdatedBy          *string    `json:"updated_by,omitempty"`
	DeletedBy          *string    `json:"deleted_by"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
	DeletedAt          *time.Time `json:"deleted_at"`
}

type StockTransferRequest struct {
	TransferNumber     *string              `json:"transfer_number"`
	FromWarehouseId    *string              `json:"from_warehouse_id"`
	ToWarehouseId      *string              `json:"to_warehouse_id"`
	MovementType       *string              `json:"movement_type"`
	ScheduledDeparture *time.Time           `json:"scheduled_departure"`
	ActualDeparture    *time.Time           `json:"actual_departure"`
	ScheduledArrival   *time.Time           `json:"scheduled_arrival"`
	ActualArrival      *time.Time           `json:"actual_arrival"`
	TransporterUnit    *string              `json:"transporter_unit"`
	Status             *string              `json:"status"`
	Items              []StockTransferItems `json:"items"`
}
