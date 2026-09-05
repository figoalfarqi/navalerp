package model

type Vessel struct {
	VesselID           int     `json:"vessel_id"`
	VendorID           *int    `json:"vendor_id,omitempty"`
	VesselName         string  `json:"vessel_name"`
	IMONumber          *string `json:"imo_number,omitempty"`
	RegistrationNumber *string `json:"registration_number,omitempty"`
	IsActive           int     `json:"is_active"`
	Vendor             *Vendor `json:"vendor,omitempty"`
	Audit
}

type VesselRequest struct {
	VendorID           *int    `json:"vendor_id,omitempty" validate:"omitempty,gt=0"`
	VesselName         string  `json:"vessel_name" validate:"required,max=200"`
	IMONumber          *string `json:"imo_number,omitempty"`
	RegistrationNumber *string `json:"registration_number,omitempty"`
	IsActive           *int    `json:"is_active,omitempty" validate:"omitempty,oneof=0 1"`
}
