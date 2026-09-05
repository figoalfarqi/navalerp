package model

import "time"

type BankMerk struct {
	BankMerkID          int        `json:"bank_merk_id"`
	BankMerkName        string     `json:"bank_merk_name"`
	BankMerkDescription *string    `json:"bank_merk_description,omitempty"`
	IsActive            int        `json:"is_active"`
	CreatedBy           int        `json:"created_by"`
	UpdatedBy           int        `json:"updated_by"`
	DeletedBy           *int       `json:"deleted_by,omitempty"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at,omitempty"`
}

type BankMerkNullable struct {
	BankMerkID          *int       `json:"bank_merk_id"`
	BankMerkName        *string    `json:"bank_merk_name"`
	BankMerkDescription *string    `json:"bank_merk_description,omitempty"`
	IsActive            *int       `json:"is_active"`
	CreatedBy           *int       `json:"created_by"`
	UpdatedBy           *int       `json:"updated_by"`
	DeletedBy           *int       `json:"deleted_by,omitempty"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at,omitempty"`
}

type BankMerkRequest struct {
	BankMerkName        string  `json:"bank_merk_name" validate:"required,min=2,max=50"`
	BankMerkDescription *string `json:"bank_merk_description,omitempty"`
	IsActive            *int    `json:"is_active,omitempty"`
}

type BankMerkResponse struct {
	BankMerkID          int        `json:"bank_merk_id"`
	BankMerkName        string     `json:"bank_merk_name"`
	BankMerkDescription *string    `json:"bank_merk_description,omitempty"`
	IsActive            int        `json:"is_active"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at,omitempty"`
}

func (bm *BankMerkNullable) ToNotNullable() *BankMerk {
	if bm == nil {
		return nil
	}
	createdAt := time.Time{}
	if bm.CreatedAt != nil {
		createdAt = *bm.CreatedAt
	}

	updatedAt := time.Time{}
	if bm.UpdatedAt != nil {
		updatedAt = *bm.UpdatedAt
	}
	return &BankMerk{
		BankMerkID:          *bm.BankMerkID,
		BankMerkName:        *bm.BankMerkName,
		BankMerkDescription: bm.BankMerkDescription,
		IsActive:            *bm.IsActive,
		CreatedAt:           createdAt,
		UpdatedAt:           updatedAt,
		DeletedAt:           bm.DeletedAt,
	}
}
