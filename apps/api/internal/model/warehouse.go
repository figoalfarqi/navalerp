package model

import (
	"time"
)

type Warehouse struct {
	WarehouseId string `json:"warehouse_id"`
	UnitId string `json:"unit_id"`
	WarehouseCode string `json:"warehouse_code"`
	WarehouseName string `json:"warehouse_name"`
	WarehouseType string `json:"warehouse_type"`
	CapacityM3 *float64 `json:"capacity_m3,omitempty"`
	ManagerUserId *string `json:"manager_user_id,omitempty"`
	LocationAddress *string `json:"location_address,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Locations []StorageLocations `json:"locations,omitempty"`
}

type StorageLocations struct {
	LocationId string `json:"location_id"`
	WarehouseId string `json:"warehouse_id"`
	ZoneName string `json:"zone_name"`
	Aisle *string `json:"aisle,omitempty"`
	Rack *string `json:"rack,omitempty"`
	Shelf *string `json:"shelf,omitempty"`
	BinCode string `json:"bin_code"`
	CapacityKg *float64 `json:"capacity_kg,omitempty"`
	IsHazardousZone *bool `json:"is_hazardous_zone,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type WarehouseRequest struct {
	UnitId *string `json:"unit_id"`
	WarehouseCode *string `json:"warehouse_code"`
	WarehouseName *string `json:"warehouse_name"`
	WarehouseType *string `json:"warehouse_type"`
	CapacityM3 *float64 `json:"capacity_m3"`
	ManagerUserId *string `json:"manager_user_id"`
	LocationAddress *string `json:"location_address"`
	IsActive *bool `json:"is_active"`
	Locations []StorageLocations `json:"locations"`
}
