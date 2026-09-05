package model

import "time"

type VendorType struct {
	VendorTypeID          int        `json:"vendor_type_id"`
	VendorTypeName        string     `json:"vendor_type_name"`
	VendorTypeDescription *string    `json:"vendor_type_description,omitempty"`
	IsActive              int        `json:"is_active"`
	CreatedBy             int        `json:"created_by"`
	UpdatedBy             int        `json:"updated_by"`
	DeletedBy             *int       `json:"deleted_by,omitempty"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	DeletedAt             *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	Vendors []Vendor `json:"vendors,omitempty"`
}

type VendorTypeRequest struct {
	VendorTypeName        string  `json:"vendor_type_name" validate:"required,min=2,max=100"`
	VendorTypeDescription *string `json:"vendor_type_description,omitempty"`
	IsActive              *int    `json:"is_active,omitempty"`
}

type VendorTypeResponse struct {
	VendorTypeID          int        `json:"vendor_type_id"`
	VendorTypeName        string     `json:"vendor_type_name"`
	VendorTypeDescription *string    `json:"vendor_type_description,omitempty"`
	IsActive              int        `json:"is_active"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	DeletedAt             *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	Vendors []VendorResponse `json:"vendors,omitempty"`
}
