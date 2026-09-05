package model

import "time"

type TruckType struct {
	TruckTypeID          int        `json:"truck_type_id"`
	TruckTypeName        string     `json:"truck_type_name"`
	TruckTypeDescription *string    `json:"truck_type_description,omitempty"`
	TruckBoxLength       float64    `json:"truck_box_length"`
	TruckBoxWidth        float64    `json:"truck_box_width"`
	TruckBoxHeight       float64    `json:"truck_box_height"`
	TruckCapacity        float64    `json:"truck_capacity"`
	IsActive             int        `json:"is_active"`
	CreatedBy            int        `json:"created_by"`
	UpdatedBy            int        `json:"updated_by"`
	DeletedBy            *int       `json:"deleted_by,omitempty"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	Trucks []Truck `json:"trucks,omitempty"`
}

type TruckTypeRequest struct {
	TruckTypeName        string  `json:"truck_type_name" validate:"required,min=2,max=100"`
	TruckTypeDescription *string `json:"truck_type_description,omitempty"`
	TruckBoxLength       float64 `json:"truck_box_length" validate:"required,gt=0"`
	TruckBoxWidth        float64 `json:"truck_box_width" validate:"required,gt=0"`
	TruckBoxHeight       float64 `json:"truck_box_height" validate:"required,gt=0"`
	TruckCapacity        float64 `json:"truck_capacity" validate:"required,gt=0"`
	IsActive             *int    `json:"is_active,omitempty"`
}

type TruckTypeResponse struct {
	TruckTypeID          int        `json:"truck_type_id"`
	TruckTypeName        string     `json:"truck_type_name"`
	TruckTypeDescription *string    `json:"truck_type_description,omitempty"`
	TruckBoxLength       float64    `json:"truck_box_length"`
	TruckBoxWidth        float64    `json:"truck_box_width"`
	TruckBoxHeight       float64    `json:"truck_box_height"`
	TruckCapacity        float64    `json:"truck_capacity"`
	IsActive             int        `json:"is_active"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty"`

	Trucks []TruckResponse `json:"trucks,omitempty"`
}
