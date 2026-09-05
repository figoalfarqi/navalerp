package model

import (
	"time"
)

type ChartOfAccount struct {
	AccountId string `json:"account_id"`
	AccountCode string `json:"account_code"`
	AccountName string `json:"account_name"`
	AccountType string `json:"account_type"`
	ParentAccountId *string `json:"parent_account_id,omitempty"`
	ParentAccountName *string `json:"parent_account_name,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ChartOfAccountRequest struct {
	AccountCode *string `json:"account_code"`
	AccountName *string `json:"account_name"`
	AccountType *string `json:"account_type"`
	ParentAccountId *string `json:"parent_account_id"`
	IsActive *bool `json:"is_active"`
}
