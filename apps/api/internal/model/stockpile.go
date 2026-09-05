package model

import (
	"time"

	"github.com/figoalfarqi/apipml/internal/helper"
)

type Stockpile struct {
	StockpileID        int        `json:"stockpile_id"`
	StockpileName      string     `json:"stockpile_name"`
	CityID             *int       `json:"city_id,omitempty"`
	StockpileAddress   *string    `json:"stockpile_address,omitempty"`
	StockpileLatitude  *float64   `json:"stockpile_latitude,omitempty"`
	StockpileLongitude *float64   `json:"stockpile_longitude,omitempty"`
	StockpileMapUrl    *string    `json:"stockpile_map_url,omitempty"`
	IsActive           int        `json:"is_active"`
	CreatedBy          int        `json:"created_by"`
	UpdatedBy          int        `json:"updated_by"`
	DeletedBy          *int       `json:"deleted_by,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	City            *City            `json:"city,omitempty"`
	StockpileCargos []StockpileCargo `json:"stockpile_cargos,omitempty"`
}

type StockpileNullable struct {
	StockpileID        *int       `json:"stockpile_id"`
	StockpileName      *string    `json:"stockpile_name"`
	CityID             *int       `json:"city_id,omitempty"`
	StockpileAddress   *string    `json:"stockpile_address,omitempty"`
	StockpileLatitude  *float64   `json:"stockpile_latitude,omitempty"`
	StockpileLongitude *float64   `json:"stockpile_longitude,omitempty"`
	StockpileMapUrl    *string    `json:"stockpile_map_url,omitempty"`
	IsActive           *int       `json:"is_active"`
	CreatedBy          *int       `json:"created_by"`
	UpdatedBy          *int       `json:"updated_by"`
	DeletedBy          *int       `json:"deleted_by,omitempty"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`

	// Optional joins
	City            *CityNullable            `json:"city,omitempty"`
	StockpileCargos []StockpileCargoNullable `json:"stockpile_cargos,omitempty"`
}

type StockpileRequest struct {
	StockpileName      string   `json:"stockpile_name" validate:"required,min=2,max=200"`
	CityID             *int     `json:"city_id" validate:"omitempty,gt=0"`
	StockpileAddress   *string  `json:"stockpile_address,omitempty"`
	StockpileLatitude  *float64 `json:"stockpile_latitude,omitempty"`
	StockpileLongitude *float64 `json:"stockpile_longitude,omitempty"`
	StockpileMapUrl    *string  `json:"stockpile_map_url,omitempty"`
	IsActive           *int     `json:"is_active,omitempty"`
}

type StockpileResponse struct {
	StockpileID        int        `json:"stockpile_id"`
	StockpileName      string     `json:"stockpile_name"`
	CityID             *int       `json:"city_id,omitempty"`
	StockpileAddress   *string    `json:"stockpile_address,omitempty"`
	StockpileLatitude  *float64   `json:"stockpile_latitude,omitempty"`
	StockpileLongitude *float64   `json:"stockpile_longitude,omitempty"`
	StockpileMapUrl    *string    `json:"stockpile_map_url,omitempty"`
	IsActive           int        `json:"is_active"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`

	City            *CityResponse            `json:"city,omitempty"`
	StockpileCargos []StockpileCargoResponse `json:"stockpile_cargos,omitempty"`
}

func (s *StockpileNullable) ToNotNullable() *Stockpile {
	if s == nil {
		return nil
	}

	createdAt := time.Time{}
	if s.CreatedAt != nil {
		createdAt = *s.CreatedAt
	}

	updatedAt := time.Time{}
	if s.UpdatedAt != nil {
		updatedAt = *s.UpdatedAt
	}

	stockpileCargos := make([]StockpileCargo, 0)
	for _, sc := range s.StockpileCargos {
		if converted := sc.ToNotNullable(); converted != nil {
			stockpileCargos = append(stockpileCargos, *converted)
		}
	}

	return &Stockpile{
		StockpileID:        helper.DerefInt(s.StockpileID),
		StockpileName:      helper.DerefString(s.StockpileName),
		CityID:             s.CityID,
		StockpileAddress:   s.StockpileAddress,
		StockpileLatitude:  s.StockpileLatitude,
		StockpileLongitude: s.StockpileLongitude,
		StockpileMapUrl:    s.StockpileMapUrl,
		IsActive:           helper.DerefInt(s.IsActive),
		CreatedBy:          helper.DerefInt(s.CreatedBy),
		UpdatedBy:          helper.DerefInt(s.UpdatedBy),
		DeletedBy:          s.DeletedBy,
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
		DeletedAt:          s.DeletedAt,
		City:               s.City.ToNotNullable(),
		StockpileCargos:    stockpileCargos,
	}
}
