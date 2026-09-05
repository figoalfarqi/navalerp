package model

import (
	"time"
)

type Payment struct {
	PaymentId string `json:"payment_id"`
	PaymentReferenceNo string `json:"payment_reference_no"`
	SppNumber *string `json:"spp_number,omitempty"`
	SpmNumber *string `json:"spm_number,omitempty"`
	InvoiceId string `json:"invoice_id"`
	PaymentDate time.Time `json:"payment_date"`
	AmountPaid float64 `json:"amount_paid"`
	PaymentMethod *string `json:"payment_method,omitempty"`
	BankSourceAccount *string `json:"bank_source_account,omitempty"`
	AuthorisedByUserId *string `json:"authorised_by_user_id,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type PaymentRequest struct {
	PaymentReferenceNo *string `json:"payment_reference_no"`
	SppNumber *string `json:"spp_number"`
	SpmNumber *string `json:"spm_number"`
	InvoiceId *string `json:"invoice_id"`
	PaymentDate *time.Time `json:"payment_date"`
	AmountPaid *float64 `json:"amount_paid"`
	PaymentMethod *string `json:"payment_method"`
	BankSourceAccount *string `json:"bank_source_account"`
	AuthorisedByUserId *string `json:"authorised_by_user_id"`
}
