package model

import (
	"time"
)

type FuelBunker struct {
	BunkerId           string     `json:"bunker_id"`
	ShipId             string     `json:"ship_id"`
	ShipName           *string    `json:"ship_name,omitempty"`
	FacilityId         *string    `json:"facility_id,omitempty"`
	FacilityName       *string    `json:"facility_name,omitempty"`
	FuelType           string     `json:"fuel_type"`
	QuantityLiters     float64    `json:"quantity_liters"`
	Density15c         *float64   `json:"density_15c,omitempty"`
	FlowRateLph        *float64   `json:"flow_rate_lph,omitempty"`
	BunkeringStartTime time.Time  `json:"bunkering_start_time"`
	BunkeringEndTime   time.Time  `json:"bunkering_end_time"`
	ReceiptVoucherNo   string     `json:"receipt_voucher_no"`
	AuthorisedByUserId string     `json:"authorised_by_user_id"`
	CreatedBy          *string    `json:"created_by,omitempty"`
	UpdatedBy          *string    `json:"updated_by,omitempty"`
	DeletedBy          *string    `json:"deleted_by"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
	DeletedAt          *time.Time `json:"deleted_at"`
}

type FuelBunkerRequest struct {
	ShipId             *string    `json:"ship_id"`
	FacilityId         *string    `json:"facility_id"`
	FuelType           *string    `json:"fuel_type"`
	QuantityLiters     *float64   `json:"quantity_liters"`
	Density15c         *float64   `json:"density_15c"`
	FlowRateLph        *float64   `json:"flow_rate_lph"`
	BunkeringStartTime *time.Time `json:"bunkering_start_time"`
	BunkeringEndTime   *time.Time `json:"bunkering_end_time"`
	ReceiptVoucherNo   *string    `json:"receipt_voucher_no"`
	AuthorisedByUserId *string    `json:"authorised_by_user_id"`
}
