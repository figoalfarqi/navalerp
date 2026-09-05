package model

import (
	"time"
)

type PmSchedule struct {
	PmId string `json:"pm_id"`
	EquipmentId string `json:"equipment_id"`
	EquipmentName *string `json:"equipment_name,omitempty"`
	PmCode string `json:"pm_code"`
	PmTitle string `json:"pm_title"`
	IntervalHours *int `json:"interval_hours,omitempty"`
	IntervalDays *int `json:"interval_days,omitempty"`
	LastPerformedAt *time.Time `json:"last_performed_at,omitempty"`
	NextDueAt *time.Time `json:"next_due_at,omitempty"`
	TaskInstructions string `json:"task_instructions"`
	EstimatedDurationHours *float64 `json:"estimated_duration_hours,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type PmScheduleRequest struct {
	EquipmentId *string `json:"equipment_id"`
	PmCode *string `json:"pm_code"`
	PmTitle *string `json:"pm_title"`
	IntervalHours *int `json:"interval_hours"`
	IntervalDays *int `json:"interval_days"`
	LastPerformedAt *time.Time `json:"last_performed_at"`
	NextDueAt *time.Time `json:"next_due_at"`
	TaskInstructions *string `json:"task_instructions"`
	EstimatedDurationHours *float64 `json:"estimated_duration_hours"`
	IsActive *bool `json:"is_active"`
}
