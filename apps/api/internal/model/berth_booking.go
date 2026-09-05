package model

import (
	"time"
)

type BerthBooking struct {
	BookingId string `json:"booking_id"`
	FacilityId string `json:"facility_id"`
	ShipId string `json:"ship_id"`
	BookingPurpose string `json:"booking_purpose"`
	Eta time.Time `json:"eta"`
	Etd time.Time `json:"etd"`
	ActualBerthTime *time.Time `json:"actual_berth_time,omitempty"`
	ActualUnberthTime *time.Time `json:"actual_unberth_time,omitempty"`
	ShorePowerKwhUsed *float64 `json:"shore_power_kwh_used,omitempty"`
	FreshWaterTonUsed *float64 `json:"fresh_water_ton_used,omitempty"`
	Status *string `json:"status,omitempty"`
	ApprovedByUserId *string `json:"approved_by_user_id,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type BerthBookingRequest struct {
	FacilityId *string `json:"facility_id"`
	ShipId *string `json:"ship_id"`
	BookingPurpose *string `json:"booking_purpose"`
	Eta *time.Time `json:"eta"`
	Etd *time.Time `json:"etd"`
	ActualBerthTime *time.Time `json:"actual_berth_time"`
	ActualUnberthTime *time.Time `json:"actual_unberth_time"`
	ShorePowerKwhUsed *float64 `json:"shore_power_kwh_used"`
	FreshWaterTonUsed *float64 `json:"fresh_water_ton_used"`
	Status *string `json:"status"`
	ApprovedByUserId *string `json:"approved_by_user_id"`
	Remarks *string `json:"remarks"`
}
