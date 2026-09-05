package model

import (
	"time"
)

type Contract struct {
	ContractId string `json:"contract_id"`
	TenderId *string `json:"tender_id,omitempty"`
	TenderNumber *string `json:"tender_number,omitempty"`
	ContractNumber string `json:"contract_number"`
	VendorId string `json:"vendor_id"`
	VendorName *string `json:"vendor_name,omitempty"`
	ContractTitle string `json:"contract_title"`
	ContractValue float64 `json:"contract_value"`
	Currency *string `json:"currency,omitempty"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	ProcurementMethod *string `json:"procurement_method,omitempty"`
	WarrantyPeriodMonths *int `json:"warranty_period_months,omitempty"`
	TotClauseSummary *string `json:"tot_clause_summary,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Amendments []ContractAmendments `json:"amendments,omitempty"`
}

type ContractAmendments struct {
	AmendmentId string `json:"amendment_id"`
	ContractId string `json:"contract_id"`
	AmendmentNumber string `json:"amendment_number"`
	AmendmentDate time.Time `json:"amendment_date"`
	Description string `json:"description"`
	AdditionalValue *float64 `json:"additional_value,omitempty"`
	ExtendedEndDate *time.Time `json:"extended_end_date,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type ContractRequest struct {
	TenderId *string `json:"tender_id"`
	ContractNumber *string `json:"contract_number"`
	VendorId *string `json:"vendor_id"`
	ContractTitle *string `json:"contract_title"`
	ContractValue *float64 `json:"contract_value"`
	Currency *string `json:"currency"`
	StartDate *time.Time `json:"start_date"`
	EndDate *time.Time `json:"end_date"`
	ProcurementMethod *string `json:"procurement_method"`
	WarrantyPeriodMonths *int `json:"warranty_period_months"`
	TotClauseSummary *string `json:"tot_clause_summary"`
	Status *string `json:"status"`
	Amendments []ContractAmendments `json:"amendments"`
}
