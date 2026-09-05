package model

import (
	"encoding/json"
	"time"
)

type Client struct {
	ClientID            int                 `json:"client_id"`
	ClientName          string              `json:"client_name"`
	ClientEmail         *string             `json:"client_email,omitempty"`
	ClientTin           *string             `json:"client_tin,omitempty"`
	NumberOfDayUntilDue *int                `json:"number_of_day_until_due,omitempty"`
	CityID              *int                `json:"city_id,omitempty"`
	ClientAddress       *string             `json:"client_address,omitempty"`
	OperatingHours      json.RawMessage     `json:"operating_hours,omitempty"`
	IsActive            int                 `json:"is_active"`
	CreatedBy           int                 `json:"created_by"`
	UpdatedBy           int                 `json:"updated_by"`
	DeletedBy           *int                `json:"deleted_by,omitempty"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
	DeletedAt           *time.Time          `json:"deleted_at,omitempty"`
	City                *City               `json:"city,omitempty"`
	ClientDestinations  []ClientDestination `json:"client_destinations,omitempty"`
	ClientPics          []AppUser           `json:"client_pics,omitempty"`
}

type ClientNullable struct {
	ClientID            *int             `json:"client_id"`
	ClientName          *string          `json:"client_name"`
	ClientEmail         *string          `json:"client_email,omitempty"`
	ClientTin           *string          `json:"client_tin,omitempty"`
	NumberOfDayUntilDue *int             `json:"number_of_day_until_due,omitempty"`
	CityID              *int             `json:"city_id,omitempty"`
	ClientAddress       *string          `json:"client_address,omitempty"`
	OperatingHours      *json.RawMessage `json:"operating_hours,omitempty"`
	IsActive            *int             `json:"is_active"`
	CreatedBy           *int             `json:"created_by"`
	UpdatedBy           *int             `json:"updated_by"`
	DeletedBy           *int             `json:"deleted_by,omitempty"`
	CreatedAt           *time.Time       `json:"created_at"`
	UpdatedAt           *time.Time       `json:"updated_at"`
	DeletedAt           *time.Time       `json:"deleted_at,omitempty"`

	City               *City               `json:"city,omitempty"`
	ClientDestinations []ClientDestination `json:"client_destinations,omitempty"`
	ClientPics         []AppUser           `json:"client_pics,omitempty"`
}

type ClientRequest struct {
	ClientName          string          `json:"client_name" validate:"required,min=2,max=100"`
	ClientEmail         *string         `json:"client_email,omitempty" validate:"omitempty,email"`
	ClientTin           *string         `json:"client_tin,omitempty"`
	NumberOfDayUntilDue *int            `json:"number_of_day_until_due,omitempty" validate:"omitempty,gt=0"`
	CityID              *int            `json:"city_id" validate:"omitempty,gt=0"`
	ClientAddress       *string         `json:"client_address,omitempty"`
	OperatingHours      json.RawMessage `json:"operating_hours,omitempty"`
	IsActive            *int            `json:"is_active,omitempty"`
}

type ClientResponse struct {
	ClientID            int             `json:"client_id"`
	ClientName          string          `json:"client_name"`
	ClientEmail         *string         `json:"client_email,omitempty"`
	ClientTin           *string         `json:"client_tin,omitempty"`
	NumberOfDayUntilDue *int            `json:"number_of_day_until_due,omitempty"`
	CityID              *int            `json:"city_id,omitempty"`
	ClientAddress       *string         `json:"client_address,omitempty"`
	OperatingHours      json.RawMessage `json:"operating_hours,omitempty"`
	IsActive            int             `json:"is_active"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"updated_at"`
	DeletedAt           *time.Time      `json:"deleted_at,omitempty"`

	City               *CityResponse               `json:"city,omitempty"`
	ClientDestinations []ClientDestinationResponse `json:"client_destinations,omitempty"`
	ClientPics         []AppUserResponse           `json:"client_pics,omitempty"`
}

func (c *ClientNullable) ToNotNullable() *Client {
	if c == nil {
		return nil
	}

	// Primary key wajib ada
	if c.ClientID == nil {
		return nil
	}

	name := ""
	if c.ClientName != nil {
		name = *c.ClientName
	}

	isActive := 0
	if c.IsActive != nil {
		isActive = *c.IsActive
	}

	createdBy := 0
	if c.CreatedBy != nil {
		createdBy = *c.CreatedBy
	}

	updatedBy := 0
	if c.UpdatedBy != nil {
		updatedBy = *c.UpdatedBy
	}

	createdAt := time.Time{}
	if c.CreatedAt != nil {
		createdAt = *c.CreatedAt
	}

	updatedAt := time.Time{}
	if c.UpdatedAt != nil {
		updatedAt = *c.UpdatedAt
	}

	var operatingHours json.RawMessage
	if c.OperatingHours != nil {
		operatingHours = *c.OperatingHours
	}

	return &Client{
		ClientID:            *c.ClientID,
		ClientName:          name,
		ClientEmail:         c.ClientEmail,
		ClientTin:           c.ClientTin,
		NumberOfDayUntilDue: c.NumberOfDayUntilDue,
		CityID:              c.CityID,
		ClientAddress:       c.ClientAddress,
		OperatingHours:      operatingHours,
		IsActive:            isActive,
		CreatedBy:           createdBy,
		UpdatedBy:           updatedBy,
		DeletedBy:           c.DeletedBy,
		CreatedAt:           createdAt,
		UpdatedAt:           updatedAt,
		DeletedAt:           c.DeletedAt,
		City:                c.City,
		ClientDestinations:  c.ClientDestinations,
		ClientPics:          c.ClientPics,
	}
}
