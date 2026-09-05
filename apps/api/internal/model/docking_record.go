package model

import (
	"time"
)

type DockingRecord struct {
	DockingId string `json:"docking_id"`
	ShipId string `json:"ship_id"`
	ShipyardName string `json:"shipyard_name"`
	DockingType string `json:"docking_type"`
	EntryDate time.Time `json:"entry_date"`
	ScheduledExitDate time.Time `json:"scheduled_exit_date"`
	ActualExitDate *time.Time `json:"actual_exit_date,omitempty"`
	SeaTrialPassed *bool `json:"sea_trial_passed,omitempty"`
	ClassificationSurveyor *string `json:"classification_surveyor,omitempty"`
	CertificateNumber *string `json:"certificate_number,omitempty"`
	TotalDockingCost *float64 `json:"total_docking_cost,omitempty"`
	DockingSummary *string `json:"docking_summary,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type DockingRecordRequest struct {
	ShipId *string `json:"ship_id"`
	ShipyardName *string `json:"shipyard_name"`
	DockingType *string `json:"docking_type"`
	EntryDate *time.Time `json:"entry_date"`
	ScheduledExitDate *time.Time `json:"scheduled_exit_date"`
	ActualExitDate *time.Time `json:"actual_exit_date"`
	SeaTrialPassed *bool `json:"sea_trial_passed"`
	ClassificationSurveyor *string `json:"classification_surveyor"`
	CertificateNumber *string `json:"certificate_number"`
	TotalDockingCost *float64 `json:"total_docking_cost"`
	DockingSummary *string `json:"docking_summary"`
}
