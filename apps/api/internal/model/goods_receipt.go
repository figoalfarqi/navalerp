package model

import (
	"time"
)

type GoodsReceipt struct {
	ReceiptId string `json:"receipt_id"`
	ReceiptNumber string `json:"receipt_number"`
	PoId string `json:"po_id"`
	PoNumber *string `json:"po_number,omitempty"`
	WarehouseId string `json:"warehouse_id"`
	WarehouseName *string `json:"warehouse_name,omitempty"`
	ReceivedDate time.Time `json:"received_date"`
	DeliveryOrderNumber *string `json:"delivery_order_number,omitempty"`
	InspectedByUserId string `json:"inspected_by_user_id"`
	InspectorName *string `json:"inspector_name,omitempty"`
	InspectionPassed *bool `json:"inspection_passed,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Items []GoodsReceiptItems `json:"items,omitempty"`
}

type GoodsReceiptItems struct {
	ReceiptItemId string `json:"receipt_item_id"`
	ReceiptId string `json:"receipt_id"`
	PoItemId string `json:"po_item_id"`
	MaterialId string `json:"material_id"`
	QuantityReceived float64 `json:"quantity_received"`
	QuantityAccepted float64 `json:"quantity_accepted"`
	QuantityRejected *float64 `json:"quantity_rejected,omitempty"`
	RejectionReason *string `json:"rejection_reason,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type GoodsReceiptRequest struct {
	ReceiptNumber *string `json:"receipt_number"`
	PoId *string `json:"po_id"`
	WarehouseId *string `json:"warehouse_id"`
	ReceivedDate *time.Time `json:"received_date"`
	DeliveryOrderNumber *string `json:"delivery_order_number"`
	InspectedByUserId *string `json:"inspected_by_user_id"`
	InspectionPassed *bool `json:"inspection_passed"`
	Remarks *string `json:"remarks"`
	Items []GoodsReceiptItems `json:"items"`
}
