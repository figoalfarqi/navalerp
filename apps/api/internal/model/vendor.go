package model

import (
	"time"
)

type Vendor struct {
	VendorId string `json:"vendor_id"`
	VendorCode string `json:"vendor_code"`
	VendorName string `json:"vendor_name"`
	TaxNumber *string `json:"tax_number,omitempty"`
	SecurityClearanceLevel *string `json:"security_clearance_level,omitempty"`
	DefenceIndustryLicenseNo *string `json:"defence_industry_license_no,omitempty"`
	Country *string `json:"country,omitempty"`
	ContactPerson *string `json:"contact_person,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
	BankAccountName *string `json:"bank_account_name,omitempty"`
	BankAccountNo *string `json:"bank_account_no,omitempty"`
	BankName *string `json:"bank_name,omitempty"`
	PerformanceRating *float64 `json:"performance_rating,omitempty"`
	IsApproved *bool `json:"is_approved,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Ratings []VendorRatings `json:"ratings,omitempty"`
}

type VendorRatings struct {
	RatingId string `json:"rating_id"`
	VendorId string `json:"vendor_id"`
	EvaluationDate time.Time `json:"evaluation_date"`
	EvaluatorUserId *string `json:"evaluator_user_id,omitempty"`
	QualityScore float64 `json:"quality_score"`
	DeliveryTimeScore float64 `json:"delivery_time_score"`
	ServiceScore float64 `json:"service_score"`
	PriceScore float64 `json:"price_score"`
	OverallScore float64 `json:"overall_score"`
	Remarks *string `json:"remarks,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type VendorRequest struct {
	VendorCode *string `json:"vendor_code"`
	VendorName *string `json:"vendor_name"`
	TaxNumber *string `json:"tax_number"`
	SecurityClearanceLevel *string `json:"security_clearance_level"`
	DefenceIndustryLicenseNo *string `json:"defence_industry_license_no"`
	Country *string `json:"country"`
	ContactPerson *string `json:"contact_person"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	BankAccountName *string `json:"bank_account_name"`
	BankAccountNo *string `json:"bank_account_no"`
	BankName *string `json:"bank_name"`
	PerformanceRating *float64 `json:"performance_rating"`
	IsApproved *bool `json:"is_approved"`
	Ratings []VendorRatings `json:"ratings"`
}
