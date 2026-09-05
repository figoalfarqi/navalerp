package model

import "time"

type Province struct {
	ProvinceID       int        `json:"province_id"`
	ProvinceName     string     `json:"province_name"`
	ProvinceRealName *string    `json:"province_real_name,omitempty"`
	IsActive         int        `json:"is_active"`
	CreatedBy        int        `json:"created_by"`
	UpdatedBy        int        `json:"updated_by"`
	DeletedBy        *int       `json:"deleted_by,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	Cities []City `json:"cities,omitempty"`
}

type ProvinceNullable struct {
	ProvinceID       *int       `json:"province_id"`
	ProvinceName     *string    `json:"province_name"`
	ProvinceRealName *string    `json:"province_real_name"`
	IsActive         *int       `json:"is_active"`
	CreatedBy        *int       `json:"created_by"`
	UpdatedBy        *int       `json:"updated_by"`
	DeletedBy        *int       `json:"deleted_by,omitempty"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

type ProvinceRequest struct {
	ProvinceName     string  `json:"province_name" validate:"required,min=2,max=100"`
	ProvinceRealName *string `json:"province_real_name,omitempty"`
	IsActive         *int    `json:"is_active,omitempty"`
}

type ProvinceResponse struct {
	ProvinceID       int        `json:"province_id"`
	ProvinceName     string     `json:"province_name"`
	ProvinceRealName *string    `json:"province_real_name,omitempty"`
	IsActive         int        `json:"is_active"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	Cities []CityResponse `json:"cities,omitempty"`
}

func (ct *ProvinceNullable) ToNotNullable() *Province {
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
	return &Province{
		ProvinceID:       *ct.ProvinceID,
		ProvinceName:     *ct.ProvinceName,
		ProvinceRealName: ct.ProvinceRealName,
		IsActive:         *ct.IsActive,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		DeletedAt:        ct.DeletedAt,
	}
}
