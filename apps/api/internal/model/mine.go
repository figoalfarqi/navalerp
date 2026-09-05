package model

import "time"

type Mine struct {
	MineID        int        `json:"mine_id"`
	MineName      string     `json:"mine_name"`
	CityID        *int       `json:"city_id,omitempty"`
	MineAddress   *string    `json:"mine_address,omitempty"`
	MineLatitude  *float64   `json:"mine_latitude,omitempty"`
	MineLongitude *float64   `json:"mine_longitude,omitempty"`
	MineMapUrl    *string    `json:"mine_map_url,omitempty"`
	IsActive      int        `json:"is_active"`
	CreatedBy     int        `json:"created_by"`
	UpdatedBy     int        `json:"updated_by"`
	DeletedBy     *int       `json:"deleted_by,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	City *City `json:"city,omitempty"`
}

type MineNullable struct {
	MineID        *int       `json:"mine_id"`
	MineName      *string    `json:"mine_name"`
	CityID        *int       `json:"city_id,omitempty"`
	MineAddress   *string    `json:"mine_address,omitempty"`
	MineLatitude  *float64   `json:"mine_latitude,omitempty"`
	MineLongitude *float64   `json:"mine_longitude,omitempty"`
	MineMapUrl    *string    `json:"mine_map_url,omitempty"`
	IsActive      *int       `json:"is_active"`
	CreatedBy     *int       `json:"created_by"`
	UpdatedBy     *int       `json:"updated_by"`
	DeletedBy     *int       `json:"deleted_by,omitempty"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	City          *City      `json:"city,omitempty"`
}

type MineRequest struct {
	MineName      string   `json:"mine_name" validate:"required,min=2,max=200"`
	CityID        *int     `json:"city_id" validate:"omitempty,gt=0"`
	MineAddress   *string  `json:"mine_address,omitempty"`
	MineLatitude  *float64 `json:"mine_latitude,omitempty"`
	MineLongitude *float64 `json:"mine_longitude,omitempty"`
	MineMapUrl    *string  `json:"mine_map_url,omitempty"`
	IsActive      *int     `json:"is_active,omitempty"`
}

type MineResponse struct {
	MineID        int        `json:"mine_id"`
	MineName      string     `json:"mine_name"`
	CityID        *int       `json:"city_id,omitempty"`
	MineAddress   *string    `json:"mine_address,omitempty"`
	MineLatitude  *float64   `json:"mine_latitude,omitempty"`
	MineLongitude *float64   `json:"mine_longitude,omitempty"`
	MineMapUrl    *string    `json:"mine_map_url,omitempty"`
	IsActive      int        `json:"is_active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`

	City *CityResponse `json:"city,omitempty"`
}

func (m *MineNullable) ToNotNullable() *Mine {
	if m == nil {
		return nil
	}

	// Wajib ada
	if m.MineID == nil {
		return nil
	}

	mineName := ""
	if m.MineName != nil {
		mineName = *m.MineName
	}

	isActive := 0
	if m.IsActive != nil {
		isActive = *m.IsActive
	}

	createdBy := 0
	if m.CreatedBy != nil {
		createdBy = *m.CreatedBy
	}

	updatedBy := 0
	if m.UpdatedBy != nil {
		updatedBy = *m.UpdatedBy
	}

	createdAt := time.Time{}
	if m.CreatedAt != nil {
		createdAt = *m.CreatedAt
	}

	updatedAt := time.Time{}
	if m.UpdatedAt != nil {
		updatedAt = *m.UpdatedAt
	}

	return &Mine{
		MineID:        *m.MineID,
		MineName:      mineName,
		CityID:        m.CityID,
		MineAddress:   m.MineAddress,
		MineLatitude:  m.MineLatitude,
		MineLongitude: m.MineLongitude,
		MineMapUrl:    m.MineMapUrl,
		IsActive:      isActive,
		CreatedBy:     createdBy,
		UpdatedBy:     updatedBy,
		DeletedBy:     m.DeletedBy,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     m.DeletedAt,
		City:          m.City,
	}
}
