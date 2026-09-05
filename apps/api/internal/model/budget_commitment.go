package model

import (
	"time"
)

type BudgetCommitment struct {
	CommitmentId string `json:"commitment_id"`
	CommitmentNumber string `json:"commitment_number"`
	AllocationId string `json:"allocation_id"`
	ContractId *string `json:"contract_id,omitempty"`
	PoId *string `json:"po_id,omitempty"`
	WorkOrderId *string `json:"work_order_id,omitempty"`
	CommittedAmount float64 `json:"committed_amount"`
	CommitmentDate time.Time `json:"commitment_date"`
	Status *string `json:"status,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type BudgetCommitmentRequest struct {
	CommitmentNumber *string `json:"commitment_number"`
	AllocationId *string `json:"allocation_id"`
	ContractId *string `json:"contract_id"`
	PoId *string `json:"po_id"`
	WorkOrderId *string `json:"work_order_id"`
	CommittedAmount *float64 `json:"committed_amount"`
	CommitmentDate *time.Time `json:"commitment_date"`
	Status *string `json:"status"`
}
