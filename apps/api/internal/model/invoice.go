package model

import (
	"time"
)

type Invoice struct {
	InvoiceId string `json:"invoice_id"`
	InvoiceNumber string `json:"invoice_number"`
	VendorId string `json:"vendor_id"`
	VendorName *string `json:"vendor_name,omitempty"`
	ContractId *string `json:"contract_id,omitempty"`
	ContractNumber *string `json:"contract_number,omitempty"`
	PoId *string `json:"po_id,omitempty"`
	PoNumber *string `json:"po_number,omitempty"`
	InvoiceDate time.Time `json:"invoice_date"`
	DueDate time.Time `json:"due_date"`
	TaxInvoiceNumber *string `json:"tax_invoice_number,omitempty"`
	Subtotal float64 `json:"subtotal"`
	TaxAmount *float64 `json:"tax_amount,omitempty"`
	TotalAmount *float64 `json:"total_amount,omitempty"`
	VerificationStatus *string `json:"verification_status,omitempty"`
	VerifiedByUserId *string `json:"verified_by_user_id,omitempty"`
	VerifierName *string `json:"verifier_name,omitempty"`
	PaymentStatus *string `json:"payment_status,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type InvoiceRequest struct {
	InvoiceNumber *string `json:"invoice_number"`
	VendorId *string `json:"vendor_id"`
	ContractId *string `json:"contract_id"`
	PoId *string `json:"po_id"`
	InvoiceDate *time.Time `json:"invoice_date"`
	DueDate *time.Time `json:"due_date"`
	TaxInvoiceNumber *string `json:"tax_invoice_number"`
	Subtotal *float64 `json:"subtotal"`
	TaxAmount *float64 `json:"tax_amount"`
	VerificationStatus *string `json:"verification_status"`
	VerifiedByUserId *string `json:"verified_by_user_id"`
	PaymentStatus *string `json:"payment_status"`
}
