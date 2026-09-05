package model

import (
	"time"
)

type Equipment struct {
	EquipmentId string `json:"equipment_id"`
	SystemId string `json:"system_id"`
	SerialNumber string `json:"serial_number"`
	EquipmentTag *string `json:"equipment_tag,omitempty"`
	EquipmentName string `json:"equipment_name"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	ModelNumber *string `json:"model_number,omitempty"`
	CountryOfOrigin *string `json:"country_of_origin,omitempty"`
	InstallationDate *time.Time `json:"installation_date,omitempty"`
	TotalOperatingHours *float64 `json:"total_operating_hours,omitempty"`
	DesignLifeHours *float64 `json:"design_life_hours,omitempty"`
	CriticalityLevel *string `json:"criticality_level,omitempty"`
	HealthStatus *string `json:"health_status,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Parameters []EquipmentParameters `json:"parameters,omitempty"`
}

type EquipmentParameters struct {
	ParamId string `json:"param_id"`
	EquipmentId string `json:"equipment_id"`
	RecordedAt time.Time `json:"recorded_at"`
	Rpm *float64 `json:"rpm,omitempty"`
	TemperatureCelsius *float64 `json:"temperature_celsius,omitempty"`
	PressureBar *float64 `json:"pressure_bar,omitempty"`
	VibrationLevel *float64 `json:"vibration_level,omitempty"`
	OilPressureBar *float64 `json:"oil_pressure_bar,omitempty"`
	RunningHoursSnapshot *float64 `json:"running_hours_snapshot,omitempty"`
	StatusFlag *string `json:"status_flag,omitempty"`
	RecordedByUserId *string `json:"recorded_by_user_id,omitempty"`
}

type EquipmentRequest struct {
	SystemId *string `json:"system_id"`
	SerialNumber *string `json:"serial_number"`
	EquipmentTag *string `json:"equipment_tag"`
	EquipmentName *string `json:"equipment_name"`
	Manufacturer *string `json:"manufacturer"`
	ModelNumber *string `json:"model_number"`
	CountryOfOrigin *string `json:"country_of_origin"`
	InstallationDate *time.Time `json:"installation_date"`
	TotalOperatingHours *float64 `json:"total_operating_hours"`
	DesignLifeHours *float64 `json:"design_life_hours"`
	CriticalityLevel *string `json:"criticality_level"`
	HealthStatus *string `json:"health_status"`
	Parameters []EquipmentParameters `json:"parameters"`
}
