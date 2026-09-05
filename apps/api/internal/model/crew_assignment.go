package model

import (
	"time"
)

type CrewAssignment struct {
	AssignmentId string `json:"assignment_id"`
	ShipId string `json:"ship_id"`
	PersonnelId string `json:"personnel_id"`
	CrewRole string `json:"crew_role"`
	Department string `json:"department"`
	WatchBillDuty *string `json:"watch_bill_duty,omitempty"`
	AssignedDate time.Time `json:"assigned_date"`
	RelievedDate *time.Time `json:"relieved_date,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Allowances []SeaDutyAllowances `json:"allowances,omitempty"`
}

type SeaDutyAllowances struct {
	AllowanceId string `json:"allowance_id"`
	PersonnelId string `json:"personnel_id"`
	ShipId string `json:"ship_id"`
	MissionName string `json:"mission_name"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	DaysAtSea int `json:"days_at_sea"`
	DailyAllowanceRate float64 `json:"daily_allowance_rate"`
	TotalAllowance *float64 `json:"total_allowance,omitempty"`
	PaymentStatus *string `json:"payment_status,omitempty"`
	PaymentReferenceNo *string `json:"payment_reference_no,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type CrewAssignmentRequest struct {
	ShipId *string `json:"ship_id"`
	PersonnelId *string `json:"personnel_id"`
	CrewRole *string `json:"crew_role"`
	Department *string `json:"department"`
	WatchBillDuty *string `json:"watch_bill_duty"`
	AssignedDate *time.Time `json:"assigned_date"`
	RelievedDate *time.Time `json:"relieved_date"`
	IsActive *bool `json:"is_active"`
	Allowances []SeaDutyAllowances `json:"allowances"`
}
