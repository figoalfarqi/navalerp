package model

import "time"

type ApprovalRequest struct {
	EntityType       string `json:"entity_type"`       // "requisition", "purchase_order", "work_order", "goods_receipt", "payment", "cui_alert"
	EntityID         string `json:"entity_id"`         // UUID of the entity
	Action           string `json:"action"`            // "APPROVE" or "REJECT"
	Notes            string `json:"notes"`             // Approval comments/remarks
	DigitalSignature string `json:"digital_signature"` // Optional signature hash / certificate
}

type ApprovalResponse struct {
	Success        bool      `json:"success"`
	EntityType     string    `json:"entity_type"`
	EntityID       string    `json:"entity_id"`
	NewStatus      string    `json:"new_status"`
	ApprovedBy     string    `json:"approved_by"`
	ApprovedAt     time.Time `json:"approved_at"`
	SignatureStamp string    `json:"signature_stamp"`
	Notes          string    `json:"notes"`
}

type PendingApprovalItem struct {
	EntityType    string    `json:"entity_type"`
	EntityID      string    `json:"entity_id"`
	ReferenceNo   string    `json:"reference_no"`
	Title         string    `json:"title"`
	RequestedBy   string    `json:"requested_by"`
	RequestedDate time.Time `json:"requested_date"`
	Amount        *float64  `json:"amount,omitempty"`
	Priority      string    `json:"priority"`
	Status        string    `json:"status"`
}
