package model

import (
	"time"
)

type Shipment struct {
	ShipmentId string `json:"shipment_id"`
	ManifestNumber string `json:"manifest_number"`
	RouteId string `json:"route_id"`
	TransportUnitId string `json:"transport_unit_id"`
	OriginWarehouseId string `json:"origin_warehouse_id"`
	OriginWarehouseName *string `json:"origin_warehouse_name,omitempty"`
	DestinationWarehouseId string `json:"destination_warehouse_id"`
	DestinationWarehouseName *string `json:"destination_warehouse_name,omitempty"`
	DepartureDate time.Time `json:"departure_date"`
	ArrivalDate *time.Time `json:"arrival_date,omitempty"`
	EscortSecurityLevel *string `json:"escort_security_level,omitempty"`
	Status *string `json:"status,omitempty"`
	AuthorizedByUserId *string `json:"authorized_by_user_id,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Items []ShipmentItems `json:"items,omitempty"`
}

type ShipmentItems struct {
	ShipmentItemId string `json:"shipment_item_id"`
	ShipmentId string `json:"shipment_id"`
	MaterialId string `json:"material_id"`
	QuantityDispatched float64 `json:"quantity_dispatched"`
	QuantityReceived *float64 `json:"quantity_received,omitempty"`
	PackagingType string `json:"packaging_type"`
	WeightKg *float64 `json:"weight_kg,omitempty"`
	Notes *string `json:"notes,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type ShipmentRequest struct {
	ManifestNumber *string `json:"manifest_number"`
	RouteId *string `json:"route_id"`
	TransportUnitId *string `json:"transport_unit_id"`
	OriginWarehouseId *string `json:"origin_warehouse_id"`
	DestinationWarehouseId *string `json:"destination_warehouse_id"`
	DepartureDate *time.Time `json:"departure_date"`
	ArrivalDate *time.Time `json:"arrival_date"`
	EscortSecurityLevel *string `json:"escort_security_level"`
	Status *string `json:"status"`
	AuthorizedByUserId *string `json:"authorized_by_user_id"`
	Remarks *string `json:"remarks"`
	Items []ShipmentItems `json:"items"`
}
