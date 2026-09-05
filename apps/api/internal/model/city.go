package model

import "time"

type City struct {
	CityID     int        `json:"city_id"`
	ProvinceID int        `json:"province_id"`
	CityName   string     `json:"city_name"`
	IsActive   int        `json:"is_active"`
	CreatedBy  int        `json:"created_by"`
	UpdatedBy  int        `json:"updated_by"`
	DeletedBy  *int       `json:"deleted_by,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	Province *Province `json:"province,omitempty"`
}

type CityNullable struct {
	CityID     *int              `json:"city_id"`
	ProvinceID *int              `json:"province_id"`
	CityName   *string           `json:"city_name"`
	IsActive   *int              `json:"is_active"` // 1 aktif, 0 tidak aktif
	CreatedBy  *int              `json:"created_by"`
	UpdatedBy  *int              `json:"updated_by"`
	DeletedBy  *int              `json:"deleted_by,omitempty"`
	CreatedAt  *time.Time        `json:"created_at"`
	UpdatedAt  *time.Time        `json:"updated_at"`
	DeletedAt  *time.Time        `json:"deleted_at,omitempty"`
	Province   *ProvinceNullable `json:"province,omitempty"`
}

type CityRequest struct {
	ProvinceID int    `json:"province_id" validate:"required,gt=0"`
	CityName   string `json:"city_name" validate:"required,min=2,max=100"`
	IsActive   *int   `json:"is_active,omitempty"`
}

type CityResponse struct {
	CityID     int        `json:"city_id"`
	ProvinceID int        `json:"province_id"`
	CityName   string     `json:"city_name"`
	IsActive   int        `json:"is_active"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	Province *ProvinceResponse `json:"province,omitempty"`
}

func (ct *CityNullable) ToNotNullable() *City {
	if ct == nil {
		return nil
	}

	createdAt := time.Time{}
	if ct.CreatedAt != nil {
		createdAt = *ct.CreatedAt
	}

	updatedAt := time.Time{}
	if ct.UpdatedAt != nil {
		updatedAt = *ct.UpdatedAt
	}

	return &City{
		CityID:     *ct.CityID,
		ProvinceID: *ct.ProvinceID,
		CityName:   *ct.CityName,
		IsActive:   *ct.IsActive,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeletedAt:  ct.DeletedAt,
		Province:   ct.Province.ToNotNullable(),
	}
}
