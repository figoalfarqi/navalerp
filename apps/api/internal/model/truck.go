package model

import "time"

type Truck struct {
	TruckID           int        `json:"truck_id"`
	TruckTypeID       *int       `json:"truck_type_id,omitempty"`
	TruckMerkID       *int       `json:"truck_merk_id,omitempty"`
	DriverID          *int       `json:"driver_id,omitempty"`
	VendorID          *int       `json:"vendor_id,omitempty"`
	LicensePlate      string     `json:"license_plate"`
	OwnershipStatusID int        `json:"ownership_status_id"`
	ProductionYear    *int       `json:"production_year,omitempty"`
	NumberOfTires     *int       `json:"number_of_tires,omitempty"`
	IsActive          int        `json:"is_active"`
	CreatedBy         int        `json:"created_by"`
	UpdatedBy         int        `json:"updated_by"`
	DeletedBy         *int       `json:"deleted_by,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	TruckType *TruckType `json:"truck_type,omitempty"`
	TruckMerk *TruckMerk `json:"truck_merk,omitempty"`
	Driver    *AppUser   `json:"driver,omitempty"`
	Vendor    *Vendor    `json:"vendor,omitempty"`
}

type TruckRequest struct {
	TruckTypeID       *int   `json:"truck_type_id" validate:"omitempty,gt=0"`
	TruckMerkID       *int   `json:"truck_merk_id" validate:"omitempty,gt=0"`
	DriverID          *int   `json:"driver_id" validate:"omitempty,gt=0"`
	VendorID          *int   `json:"vendor_id" validate:"omitempty,gt=0"`
	LicensePlate      string `json:"license_plate" validate:"required,min=1,max=20"`
	OwnershipStatusID int    `json:"ownership_status_id" validate:"required,gt=0"`
	ProductionYear    *int   `json:"production_year,omitempty" validate:"omitempty,gt=1900"`
	NumberOfTires     *int   `json:"number_of_tires,omitempty" validate:"omitempty,gt=0"`
	IsActive          *int   `json:"is_active,omitempty"`
}

type TruckResponse struct {
	TruckID           int        `json:"truck_id"`
	TruckTypeID       *int       `json:"truck_type_id,omitempty"`
	TruckMerkID       *int       `json:"truck_merk_id,omitempty"`
	DriverID          *int       `json:"driver_id,omitempty"`
	VendorID          *int       `json:"vendor_id,omitempty"`
	LicensePlate      string     `json:"license_plate"`
	OwnershipStatusID int        `json:"ownership_status_id"`
	ProductionYear    *int       `json:"production_year,omitempty"`
	NumberOfTires     *int       `json:"number_of_tires,omitempty"`
	IsActive          int        `json:"is_active"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`

	TruckType *TruckTypeResponse `json:"truck_type,omitempty"`
	TruckMerk *TruckMerkResponse `json:"truck_merk,omitempty"`
	Driver    *AppUserResponse   `json:"driver,omitempty"`
	Vendor    *VendorResponse    `json:"vendor,omitempty"`
}
