package model

import "time"

type ProjectFinancialTransaction struct {
	ProjectFinancialTransactionID int               `json:"project_financial_transaction_id"`
	ProjectID                     int               `json:"project_id"`
	ProjectTransportID            *int              `json:"project_transport_id,omitempty"`
	TransactionNumber             string            `json:"transaction_number"`
	TransactionKind               string            `json:"transaction_kind"`
	TransactionDate               string            `json:"transaction_date"`
	DueDate                       *string           `json:"due_date,omitempty"`
	PaidAt                        *time.Time        `json:"paid_at,omitempty"`
	ClientID                      *int              `json:"client_id,omitempty"`
	VendorID                      *int              `json:"vendor_id,omitempty"`
	ReferenceNumber               *string           `json:"reference_number,omitempty"`
	MaterialSaleIncome            float64           `json:"material_sale_income"`
	TransportServiceIncome        float64           `json:"transport_service_income"`
	OtherIncome                   float64           `json:"other_income"`
	MaterialPurchaseExpense       float64           `json:"material_purchase_expense"`
	TransportExpense              float64           `json:"transport_expense"`
	RoadMoneyExpense              float64           `json:"road_money_expense"`
	LoadingExpense                float64           `json:"loading_expense"`
	UnloadingExpense              float64           `json:"unloading_expense"`
	FuelExpense                   float64           `json:"fuel_expense"`
	TollExpense                   float64           `json:"toll_expense"`
	OtherExpense                  float64           `json:"other_expense"`
	TransactionNote               *string           `json:"transaction_note,omitempty"`
	IsPosted                      int               `json:"is_posted"`
	Project                       *Project          `json:"project,omitempty"`
	ProjectTransport              *ProjectTransport `json:"project_transport,omitempty"`
	Client                        *Client           `json:"client,omitempty"`
	Vendor                        *Vendor           `json:"vendor,omitempty"`
	Audit
}

type ProjectFinancialTransactionRequest struct {
	ProjectID               int        `json:"project_id" validate:"required,gt=0"`
	ProjectTransportID      *int       `json:"project_transport_id,omitempty"`
	TransactionNumber       string     `json:"transaction_number" validate:"required,max=100"`
	TransactionKind         string     `json:"transaction_kind" validate:"required,oneof=RECEIVABLE RECEIPT PAYABLE PAYMENT ADJUSTMENT"`
	TransactionDate         string     `json:"transaction_date" validate:"required"`
	DueDate                 *string    `json:"due_date,omitempty"`
	PaidAt                  *time.Time `json:"paid_at,omitempty"`
	ClientID                *int       `json:"client_id,omitempty"`
	VendorID                *int       `json:"vendor_id,omitempty"`
	ReferenceNumber         *string    `json:"reference_number,omitempty"`
	MaterialSaleIncome      *float64   `json:"material_sale_income,omitempty"`
	TransportServiceIncome  *float64   `json:"transport_service_income,omitempty"`
	OtherIncome             *float64   `json:"other_income,omitempty"`
	MaterialPurchaseExpense *float64   `json:"material_purchase_expense,omitempty"`
	TransportExpense        *float64   `json:"transport_expense,omitempty"`
	RoadMoneyExpense        *float64   `json:"road_money_expense,omitempty"`
	LoadingExpense          *float64   `json:"loading_expense,omitempty"`
	UnloadingExpense        *float64   `json:"unloading_expense,omitempty"`
	FuelExpense             *float64   `json:"fuel_expense,omitempty"`
	TollExpense             *float64   `json:"toll_expense,omitempty"`
	OtherExpense            *float64   `json:"other_expense,omitempty"`
	TransactionNote         *string    `json:"transaction_note,omitempty"`
	IsPosted                *int       `json:"is_posted,omitempty" validate:"omitempty,oneof=0 1"`
}
