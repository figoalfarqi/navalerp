package model

import (
	"time"
)

type PlatformTco struct {
	TcoId string `json:"tco_id"`
	ShipId string `json:"ship_id"`
	FiscalYear int `json:"fiscal_year"`
	AcquisitionAmortization *float64 `json:"acquisition_amortization,omitempty"`
	FuelLubeCost *float64 `json:"fuel_lube_cost,omitempty"`
	MroSparepartsCost *float64 `json:"mro_spareparts_cost,omitempty"`
	DockingServicesCost *float64 `json:"docking_services_cost,omitempty"`
	CrewPayrollAllowances *float64 `json:"crew_payroll_allowances,omitempty"`
	ModernizationUpgradesCost *float64 `json:"modernization_upgrades_cost,omitempty"`
	TotalAnnualOperatingCost *float64 `json:"total_annual_operating_cost,omitempty"`
	OperatingHoursSea *float64 `json:"operating_hours_sea,omitempty"`
	CostPerOperatingHour *float64 `json:"cost_per_operating_hour,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type PlatformTcoRequest struct {
	ShipId *string `json:"ship_id"`
	FiscalYear *int `json:"fiscal_year"`
	AcquisitionAmortization *float64 `json:"acquisition_amortization"`
	FuelLubeCost *float64 `json:"fuel_lube_cost"`
	MroSparepartsCost *float64 `json:"mro_spareparts_cost"`
	DockingServicesCost *float64 `json:"docking_services_cost"`
	CrewPayrollAllowances *float64 `json:"crew_payroll_allowances"`
	ModernizationUpgradesCost *float64 `json:"modernization_upgrades_cost"`
	OperatingHoursSea *float64 `json:"operating_hours_sea"`
	CostPerOperatingHour *float64 `json:"cost_per_operating_hour"`
	Remarks *string `json:"remarks"`
}
