package model

import (
	"time"
)

type PurchaseOrder struct {
	PoId string `json:"po_id"`
	PoNumber string `json:"po_number"`
	ContractId *string `json:"contract_id,omitempty"`
	VendorId string `json:"vendor_id"`
	IssuingUnitId string `json:"issuing_unit_id"`
	OrderDate time.Time `json:"order_date"`
	DeliveryDeadline *time.Time `json:"delivery_deadline,omitempty"`
	DestinationWarehouseId *string `json:"destination_warehouse_id,omitempty"`
	TotalAmount float64 `json:"total_amount"`
	TaxAmount *float64 `json:"tax_amount,omitempty"`
	GrandTotal *float64 `json:"grand_total,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Items []PurchaseOrderItems `json:"items,omitempty"`
}

type PurchaseOrderItems struct {
	PoItemId string `json:"po_item_id"`
	PoId string `json:"po_id"`
	MaterialId string `json:"material_id"`
	Quantity float64 `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	TotalPrice *float64 `json:"total_price,omitempty"`
	Notes *string `json:"notes,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type PurchaseOrderRequest struct {
	PoNumber *string `json:"po_number"`
	ContractId *string `json:"contract_id"`
	VendorId *string `json:"vendor_id"`
	IssuingUnitId *string `json:"issuing_unit_id"`
	OrderDate *time.Time `json:"order_date"`
	DeliveryDeadline *time.Time `json:"delivery_deadline"`
	DestinationWarehouseId *string `json:"destination_warehouse_id"`
	TotalAmount *float64 `json:"total_amount"`
	TaxAmount *float64 `json:"tax_amount"`
	Status *string `json:"status"`
	Items []PurchaseOrderItems `json:"items"`
}
