package model

import (
	"time"
)

type Requisition struct {
	RequisitionId      string             `json:"requisition_id"`
	RequisitionNumber  string             `json:"requisition_number"`
	OriginUnitId       string             `json:"origin_unit_id"`
	UnitName           *string            `json:"unit_name,omitempty"`
	WorkOrderId        *string            `json:"work_order_id,omitempty"`
	Priority           string             `json:"priority"`
	RequestedDate      time.Time          `json:"requested_date"`
	RequiredByDate     *time.Time         `json:"required_by_date,omitempty"`
	ApprovalStatus     *string            `json:"approval_status,omitempty"`
	ApprovedByUserId   *string            `json:"approved_by_user_id,omitempty"`
	ApproverName       *string            `json:"approver_name,omitempty"`
	ApprovedAt         *time.Time         `json:"approved_at,omitempty"`
	TotalEstimatedCost *float64           `json:"total_estimated_cost,omitempty"`
	Justification      *string            `json:"justification,omitempty"`
	CreatedBy          *string            `json:"created_by,omitempty"`
	UpdatedBy          *string            `json:"updated_by,omitempty"`
	DeletedBy          *string            `json:"deleted_by"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          *time.Time         `json:"updated_at,omitempty"`
	DeletedAt          *time.Time         `json:"deleted_at"`
	Items              []RequisitionItems `json:"items,omitempty"`
}

type RequisitionItems struct {
	ReqItemId           string     `json:"req_item_id"`
	RequisitionId       string     `json:"requisition_id"`
	MaterialId          string     `json:"material_id"`
	MaterialName        *string    `json:"material_name,omitempty"`
	MaterialCode        *string    `json:"material_code,omitempty"`
	Quantity            float64    `json:"quantity"`
	EstimatedUnitPrice  float64    `json:"estimated_unit_price"`
	EstimatedTotalPrice *float64   `json:"estimated_total_price,omitempty"`
	Notes               *string    `json:"notes,omitempty"`
	CreatedBy           *string    `json:"created_by,omitempty"`
	UpdatedBy           *string    `json:"updated_by,omitempty"`
	DeletedBy           *string    `json:"deleted_by"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
	DeletedAt           *time.Time `json:"deleted_at"`
}

type RequisitionRequest struct {
	RequisitionNumber  *string            `json:"requisition_number"`
	OriginUnitId       *string            `json:"origin_unit_id"`
	WorkOrderId        *string            `json:"work_order_id"`
	Priority           *string            `json:"priority"`
	RequestedDate      *time.Time         `json:"requested_date"`
	RequiredByDate     *time.Time         `json:"required_by_date"`
	ApprovalStatus     *string            `json:"approval_status"`
	ApprovedByUserId   *string            `json:"approved_by_user_id"`
	ApprovedAt         *time.Time         `json:"approved_at"`
	TotalEstimatedCost *float64           `json:"total_estimated_cost"`
	Justification      *string            `json:"justification"`
	Items              []RequisitionItems `json:"items"`
}
