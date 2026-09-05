package model

import "time"

type TruckMerk struct {
	TruckMerkID   int        `json:"truck_merk_id"`
	TruckMerkName string     `json:"truck_merk_name"`
	IsActive      int        `json:"is_active"`
	CreatedBy     int        `json:"created_by"`
	UpdatedBy     int        `json:"updated_by"`
	DeletedBy     *int       `json:"deleted_by,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	Trucks []Truck `json:"trucks,omitempty"`
}

type TruckMerkRequest struct {
	TruckMerkName string `json:"truck_merk_name" validate:"required,min=2,max=100"`
	IsActive      *int   `json:"is_active,omitempty"`
}

type TruckMerkResponse struct {
	TruckMerkID   int        `json:"truck_merk_id"`
	TruckMerkName string     `json:"truck_merk_name"`
	IsActive      int        `json:"is_active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`

	Trucks []TruckResponse `json:"trucks,omitempty"`
}
