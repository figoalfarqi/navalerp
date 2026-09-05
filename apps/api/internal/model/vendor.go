package model

import "time"

type Vendor struct {
	VendorID          int        `json:"vendor_id"`
	VendorTypeID      int        `json:"vendor_type_id"`
	BankMerkID        *int       `json:"bank_merk_id,omitempty"`
	VendorName        string     `json:"vendor_name"`
	VendorEmail       *string    `json:"vendor_email,omitempty"`
	VendorPhone       *string    `json:"vendor_phone,omitempty"`
	VendorTin         *string    `json:"vendor_tin,omitempty"`
	CityID            *int       `json:"city_id,omitempty"`
	VendorAddress     *string    `json:"vendor_address,omitempty"`
	BankAccountNumber *string    `json:"bank_account_number,omitempty"`
	BankAccountName   *string    `json:"bank_account_name,omitempty"`
	IsActive          int        `json:"is_active"`
	CreatedBy         int        `json:"created_by"`
	UpdatedBy         int        `json:"updated_by"`
	DeletedBy         *int       `json:"deleted_by,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	VendorType *VendorType `json:"vendor_type,omitempty"`
	BankMerk   *BankMerk   `json:"bank_merk,omitempty"`
	City       *City       `json:"city,omitempty"`
	Trucks     []Truck     `json:"trucks,omitempty"`
}

type VendorNullable struct {
	VendorID          *int       `json:"vendor_id"`
	VendorTypeID      *int       `json:"vendor_type_id,omitempty"`
	BankMerkID        *int       `json:"bank_merk_id,omitempty"`
	VendorName        *string    `json:"vendor_name"`
	VendorEmail       *string    `json:"vendor_email,omitempty"`
	VendorPhone       *string    `json:"vendor_phone,omitempty"`
	VendorTin         *string    `json:"vendor_tin,omitempty"`
	CityID            *int       `json:"city_id,omitempty"`
	VendorAddress     *string    `json:"vendor_address,omitempty"`
	BankAccountNumber *string    `json:"bank_account_number,omitempty"`
	BankAccountName   *string    `json:"bank_account_name,omitempty"`
	IsActive          *int       `json:"is_active"`
	CreatedBy         *int       `json:"created_by"`
	UpdatedBy         *int       `json:"updated_by"`
	DeletedBy         *int       `json:"deleted_by,omitempty"`
	CreatedAt         *time.Time `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	VendorType *VendorType `json:"vendor_type,omitempty"`
	BankMerk   *BankMerk   `json:"bank_merk,omitempty"`
	City       *City       `json:"city,omitempty"`
	Trucks     []Truck     `json:"trucks,omitempty"`
}

type VendorRequest struct {
	VendorTypeID      int     `json:"vendor_type_id" validate:"required,gt=0"`
	BankMerkID        *int    `json:"bank_merk_id" validate:"omitempty,gt=0"`
	VendorName        string  `json:"vendor_name" validate:"required,min=2,max=100"`
	VendorEmail       *string `json:"vendor_email,omitempty" validate:"omitempty,email"`
	VendorPhone       *string `json:"vendor_phone,omitempty" validate:"omitempty,max=20"`
	VendorTin         *string `json:"vendor_tin,omitempty"`
	CityID            *int    `json:"city_id" validate:"omitempty,gt=0"`
	VendorAddress     *string `json:"vendor_address,omitempty"`
	BankAccountNumber *string `json:"bank_account_number,omitempty"`
	BankAccountName   *string `json:"bank_account_name,omitempty"`
	IsActive          *int    `json:"is_active,omitempty"`
}

type VendorResponse struct {
	VendorID          int        `json:"vendor_id"`
	VendorTypeID      int        `json:"vendor_type_id"`
	BankMerkID        *int       `json:"bank_merk_id,omitempty"`
	VendorName        string     `json:"vendor_name"`
	VendorEmail       *string    `json:"vendor_email,omitempty"`
	VendorPhone       *string    `json:"vendor_phone,omitempty"`
	VendorTin         *string    `json:"vendor_tin,omitempty"`
	CityID            *int       `json:"city_id,omitempty"`
	VendorAddress     *string    `json:"vendor_address,omitempty"`
	BankAccountNumber *string    `json:"bank_account_number,omitempty"`
	BankAccountName   *string    `json:"bank_account_name,omitempty"`
	IsActive          int        `json:"is_active"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`

	VendorType *VendorTypeResponse `json:"vendor_type,omitempty"`
	BankMerk   *BankMerkResponse   `json:"bank_merk,omitempty"`
	City       *CityResponse       `json:"city,omitempty"`
	Trucks     []TruckResponse     `json:"trucks,omitempty"`
}

func (bm *VendorNullable) ToNotNullable() *Vendor {
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
	createdBy := 0
	if bm.CreatedBy != nil {
		createdBy = *bm.CreatedBy
	}

	updatedBy := 0
	if bm.UpdatedBy != nil {
		updatedBy = *bm.UpdatedBy
	}
	vendorTypeID := 0
	if bm.VendorTypeID != nil {
		vendorTypeID = *bm.VendorTypeID
	}
	return &Vendor{
		VendorID:          *bm.VendorID,
		VendorTypeID:      vendorTypeID,
		BankMerkID:        bm.BankMerkID,
		VendorName:        *bm.VendorName,
		VendorEmail:       bm.VendorEmail,
		VendorPhone:       bm.VendorPhone,
		VendorTin:         bm.VendorTin,
		CityID:            bm.CityID,
		VendorAddress:     bm.VendorAddress,
		BankAccountNumber: bm.BankAccountNumber,
		BankAccountName:   bm.BankAccountName,
		IsActive:          *bm.IsActive,
		CreatedBy:         createdBy,
		UpdatedBy:         updatedBy,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
		DeletedAt:         bm.DeletedAt,
		VendorType:        bm.VendorType,
		BankMerk:          bm.BankMerk,
		City:              bm.City,
		Trucks:            bm.Trucks,
	}
}
