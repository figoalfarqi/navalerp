package model

import (
	"time"
)

type FailureReport struct {
	ReportId string `json:"report_id"`
	EquipmentId string `json:"equipment_id"`
	EquipmentName *string `json:"equipment_name,omitempty"`
	ReportedByUserId string `json:"reported_by_user_id"`
	ReporterName *string `json:"reporter_name,omitempty"`
	ReportNumber string `json:"report_number"`
	IncidentDate time.Time `json:"incident_date"`
	Severity string `json:"severity"`
	FailureMode *string `json:"failure_mode,omitempty"`
	Description string `json:"description"`
	OperationalImpact *string `json:"operational_impact,omitempty"`
	ImmediateActionTaken *string `json:"immediate_action_taken,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type FailureReportRequest struct {
	EquipmentId *string `json:"equipment_id"`
	ReportedByUserId *string `json:"reported_by_user_id"`
	ReportNumber *string `json:"report_number"`
	IncidentDate *time.Time `json:"incident_date"`
	Severity *string `json:"severity"`
	FailureMode *string `json:"failure_mode"`
	Description *string `json:"description"`
	OperationalImpact *string `json:"operational_impact"`
	ImmediateActionTaken *string `json:"immediate_action_taken"`
	Status *string `json:"status"`
}
