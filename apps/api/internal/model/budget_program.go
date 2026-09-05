package model

import (
	"time"
)

type BudgetProgram struct {
	ProgramId string `json:"program_id"`
	FiscalYear int `json:"fiscal_year"`
	DipaNumber string `json:"dipa_number"`
	ProgramCode string `json:"program_code"`
	ProgramName string `json:"program_name"`
	TotalBudget float64 `json:"total_budget"`
	ResponsibleUnitId string `json:"responsible_unit_id"`
	Status *string `json:"status,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Allocations []BudgetAllocations `json:"allocations,omitempty"`
}

type BudgetAllocations struct {
	AllocationId string `json:"allocation_id"`
	ProgramId string `json:"program_id"`
	ActivityCode string `json:"activity_code"`
	ActivityName string `json:"activity_name"`
	TargetUnitId string `json:"target_unit_id"`
	ShipId *string `json:"ship_id,omitempty"`
	AccountId string `json:"account_id"`
	AllocatedAmount float64 `json:"allocated_amount"`
	AbsorbedAmount *float64 `json:"absorbed_amount,omitempty"`
	RemainingAmount *float64 `json:"remaining_amount,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type BudgetProgramRequest struct {
	FiscalYear *int `json:"fiscal_year"`
	DipaNumber *string `json:"dipa_number"`
	ProgramCode *string `json:"program_code"`
	ProgramName *string `json:"program_name"`
	TotalBudget *float64 `json:"total_budget"`
	ResponsibleUnitId *string `json:"responsible_unit_id"`
	Status *string `json:"status"`
	Allocations []BudgetAllocations `json:"allocations"`
}
