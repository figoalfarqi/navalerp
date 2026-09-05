package model

import (
	"encoding/json"
	"time"
)

type ClientDestination struct {
	ClientDestinationID        int             `json:"client_destination_id"`
	ClientID                   *int            `json:"client_id,omitempty"`
	ClientDestinationName      string          `json:"client_destination_name"`
	CityID                     *int            `json:"city_id,omitempty"`
	ClientDestinationAddress   *string         `json:"client_destination_address,omitempty"`
	ClientDestinationLatitude  *float64        `json:"client_destination_latitude,omitempty"`
	ClientDestinationLongitude *float64        `json:"client_destination_longitude,omitempty"`
	ClientDestinationMapUrl    *string         `json:"client_destination_map_url,omitempty"`
	OperatingHours             json.RawMessage `json:"operating_hours,omitempty"`
	IsActive                   int             `json:"is_active"`
	CreatedBy                  int             `json:"created_by"`
	UpdatedBy                  int             `json:"updated_by"`
	DeletedBy                  *int            `json:"deleted_by,omitempty"`
	CreatedAt                  time.Time       `json:"created_at"`
	UpdatedAt                  time.Time       `json:"updated_at"`
	DeletedAt                  *time.Time      `json:"deleted_at,omitempty"`
	Client                     *Client         `json:"client,omitempty"`
	City                       *City           `json:"city,omitempty"`
}

type ClientDestinationNullable struct {
	ClientDestinationID        *int             `json:"client_destination_id"`
	ClientID                   *int             `json:"client_id,omitempty"`
	ClientDestinationName      *string          `json:"client_destination_name"`
	CityID                     *int             `json:"city_id,omitempty"`
	ClientDestinationAddress   *string          `json:"client_destination_address,omitempty"`
	ClientDestinationLatitude  *float64         `json:"client_destination_latitude,omitempty"`
	ClientDestinationLongitude *float64         `json:"client_destination_longitude,omitempty"`
	ClientDestinationMapUrl    *string          `json:"client_destination_map_url,omitempty"`
	OperatingHours             *json.RawMessage `json:"operating_hours,omitempty"`
	IsActive                   *int             `json:"is_active"`
	CreatedBy                  *int             `json:"created_by"`
	UpdatedBy                  *int             `json:"updated_by"`
	DeletedBy                  *int             `json:"deleted_by,omitempty"`
	CreatedAt                  *time.Time       `json:"created_at"`
	UpdatedAt                  *time.Time       `json:"updated_at"`
	DeletedAt                  *time.Time       `json:"deleted_at,omitempty"`

	Client *Client `json:"client,omitempty"`
	City   *City   `json:"city,omitempty"`
}

type ClientDestinationRequest struct {
	ClientID                   *int            `json:"client_id" validate:"omitempty,gt=0"`
	ClientDestinationName      string          `json:"client_destination_name" validate:"required,min=2,max=100"`
	CityID                     *int            `json:"city_id" validate:"omitempty,gt=0"`
	ClientDestinationAddress   *string         `json:"client_destination_address,omitempty"`
	ClientDestinationLatitude  *float64        `json:"client_destination_latitude,omitempty"`
	ClientDestinationLongitude *float64        `json:"client_destination_longitude,omitempty"`
	ClientDestinationMapUrl    *string         `json:"client_destination_map_url,omitempty"`
	OperatingHours             json.RawMessage `json:"operating_hours,omitempty"`
	IsActive                   *int            `json:"is_active,omitempty"`
}

type ClientDestinationResponse struct {
	ClientDestinationID        int             `json:"client_destination_id"`
	ClientID                   *int            `json:"client_id,omitempty"`
	ClientDestinationName      string          `json:"client_destination_name"`
	CityID                     *int            `json:"city_id,omitempty"`
	ClientDestinationAddress   *string         `json:"client_destination_address,omitempty"`
	ClientDestinationLatitude  *float64        `json:"client_destination_latitude,omitempty"`
	ClientDestinationLongitude *float64        `json:"client_destination_longitude,omitempty"`
	ClientDestinationMapUrl    *string         `json:"client_destination_map_url,omitempty"`
	OperatingHours             json.RawMessage `json:"operating_hours,omitempty"`
	IsActive                   int             `json:"is_active"`
	CreatedAt                  time.Time       `json:"created_at"`
	UpdatedAt                  time.Time       `json:"updated_at"`
	DeletedAt                  *time.Time      `json:"deleted_at,omitempty"`

	Client *ClientResponse `json:"client,omitempty"`
	City   *CityResponse   `json:"city,omitempty"`
}

func (cd *ClientDestinationNullable) ToNotNullable() *ClientDestination {
	if cd == nil {
		return nil
	}

	// Primary key wajib ada
	if cd.ClientDestinationID == nil {
		return nil
	}

	createdAt := time.Time{}
	if cd.CreatedAt != nil {
		createdAt = *cd.CreatedAt
	}

	updatedAt := time.Time{}
	if cd.UpdatedAt != nil {
		updatedAt = *cd.UpdatedAt
	}

	name := ""
	if cd.ClientDestinationName != nil {
		name = *cd.ClientDestinationName
	}

	isActive := 0
	if cd.IsActive != nil {
		isActive = *cd.IsActive
	}

	createdBy := 0
	if cd.CreatedBy != nil {
		createdBy = *cd.CreatedBy
	}

	updatedBy := 0
	if cd.UpdatedBy != nil {
		updatedBy = *cd.UpdatedBy
	}

	var operatingHours json.RawMessage
	if cd.OperatingHours != nil {
		operatingHours = *cd.OperatingHours
	}

	return &ClientDestination{
		ClientDestinationID:        *cd.ClientDestinationID,
		ClientID:                   cd.ClientID,
		ClientDestinationName:      name,
		CityID:                     cd.CityID,
		ClientDestinationAddress:   cd.ClientDestinationAddress,
		ClientDestinationLatitude:  cd.ClientDestinationLatitude,
		ClientDestinationLongitude: cd.ClientDestinationLongitude,
		ClientDestinationMapUrl:    cd.ClientDestinationMapUrl,
		OperatingHours:             operatingHours,
		IsActive:                   isActive,
		CreatedBy:                  createdBy,
		UpdatedBy:                  updatedBy,
		DeletedBy:                  cd.DeletedBy,
		CreatedAt:                  createdAt,
		UpdatedAt:                  updatedAt,
		DeletedAt:                  cd.DeletedAt,
		Client:                     cd.Client,
		City:                       cd.City,
	}
}
