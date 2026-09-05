package model

import (
	"time"
)

type WorkOrder struct {
	WorkOrderId string `json:"work_order_id"`
	FailureReportId *string `json:"failure_report_id,omitempty"`
	PmScheduleId *string `json:"pm_schedule_id,omitempty"`
	EquipmentId string `json:"equipment_id"`
	EquipmentName *string `json:"equipment_name,omitempty"`
	WorkOrderNumber string `json:"work_order_number"`
	WorkOrderType string `json:"work_order_type"`
	Priority string `json:"priority"`
	ScheduledStartDate *time.Time `json:"scheduled_start_date,omitempty"`
	ScheduledEndDate *time.Time `json:"scheduled_end_date,omitempty"`
	ActualStartDate *time.Time `json:"actual_start_date,omitempty"`
	ActualEndDate *time.Time `json:"actual_end_date,omitempty"`
	LeadEngineerUserId *string `json:"lead_engineer_user_id,omitempty"`
	EngineerName *string `json:"engineer_name,omitempty"`
	AssignedFacility *string `json:"assigned_facility,omitempty"`
	Status *string `json:"status,omitempty"`
	TotalLaborHours *float64 `json:"total_labor_hours,omitempty"`
	EstimatedCost *float64 `json:"estimated_cost,omitempty"`
	ActualCost *float64 `json:"actual_cost,omitempty"`
	CompletionNotes *string `json:"completion_notes,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Tasks []WorkOrderTasks `json:"tasks,omitempty"`
	Items []WorkOrderItems `json:"items,omitempty"`
}

type WorkOrderTasks struct {
	TaskId string `json:"task_id"`
	WorkOrderId string `json:"work_order_id"`
	StepNumber int `json:"step_number"`
	TaskDescription string `json:"task_description"`
	EstimatedMinutes *int `json:"estimated_minutes,omitempty"`
	ActualMinutes *int `json:"actual_minutes,omitempty"`
	IsCompleted *bool `json:"is_completed,omitempty"`
	CompletedByUserId *string `json:"completed_by_user_id,omitempty"`
	Notes *string `json:"notes,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type WorkOrderItems struct {
	WoItemId string `json:"wo_item_id"`
	WorkOrderId string `json:"work_order_id"`
	MaterialId string `json:"material_id"`
	QuantityRequired float64 `json:"quantity_required"`
	QuantityIssued *float64 `json:"quantity_issued,omitempty"`
	UnitCost *float64 `json:"unit_cost,omitempty"`
	TotalCost *float64 `json:"total_cost,omitempty"`
	IsCriticalSpare *bool `json:"is_critical_spare,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type WorkOrderRequest struct {
	FailureReportId *string `json:"failure_report_id"`
	PmScheduleId *string `json:"pm_schedule_id"`
	EquipmentId *string `json:"equipment_id"`
	WorkOrderNumber *string `json:"work_order_number"`
	WorkOrderType *string `json:"work_order_type"`
	Priority *string `json:"priority"`
	ScheduledStartDate *time.Time `json:"scheduled_start_date"`
	ScheduledEndDate *time.Time `json:"scheduled_end_date"`
	ActualStartDate *time.Time `json:"actual_start_date"`
	ActualEndDate *time.Time `json:"actual_end_date"`
	LeadEngineerUserId *string `json:"lead_engineer_user_id"`
	AssignedFacility *string `json:"assigned_facility"`
	Status *string `json:"status"`
	TotalLaborHours *float64 `json:"total_labor_hours"`
	EstimatedCost *float64 `json:"estimated_cost"`
	ActualCost *float64 `json:"actual_cost"`
	CompletionNotes *string `json:"completion_notes"`
	Tasks []WorkOrderTasks `json:"tasks"`
	Items []WorkOrderItems `json:"items"`
}
