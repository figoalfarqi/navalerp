package model

import (
	"time"
)

type Tender struct {
	TenderId string `json:"tender_id"`
	TenderNumber string `json:"tender_number"`
	Title string `json:"title"`
	ProcurementCategory string `json:"procurement_category"`
	EstimatedBudget float64 `json:"estimated_budget"`
	ProcurementMethod string `json:"procurement_method"`
	StartDate time.Time `json:"start_date"`
	ClosingDate time.Time `json:"closing_date"`
	Status *string `json:"status,omitempty"`
	WinnerVendorId *string `json:"winner_vendor_id,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Bids []TenderBids `json:"bids,omitempty"`
}

type TenderBids struct {
	BidId string `json:"bid_id"`
	TenderId string `json:"tender_id"`
	VendorId string `json:"vendor_id"`
	BidAmount float64 `json:"bid_amount"`
	SubmissionDate time.Time `json:"submission_date"`
	TechnicalScore *float64 `json:"technical_score,omitempty"`
	CommercialScore *float64 `json:"commercial_score,omitempty"`
	IsWinner *bool `json:"is_winner,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type TenderRequest struct {
	TenderNumber *string `json:"tender_number"`
	Title *string `json:"title"`
	ProcurementCategory *string `json:"procurement_category"`
	EstimatedBudget *float64 `json:"estimated_budget"`
	ProcurementMethod *string `json:"procurement_method"`
	StartDate *time.Time `json:"start_date"`
	ClosingDate *time.Time `json:"closing_date"`
	Status *string `json:"status"`
	WinnerVendorId *string `json:"winner_vendor_id"`
	Bids []TenderBids `json:"bids"`
}
