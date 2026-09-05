package model

import "time"

// Audit is embedded by database-backed API resources.
type Audit struct {
	CreatedBy int        `json:"created_by"`
	UpdatedBy int        `json:"updated_by"`
	DeletedBy *int       `json:"deleted_by,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ListResult[T any] struct {
	Items  []T `json:"items"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListOptions struct {
	Limit           int
	Offset          int
	CursorKey       *int
	Sort            string
	Search          string
	ProjectID       *int
	TransportID     *int
	DriverID        *int
	CheckerID       *int
	TruckID         *int
	VesselID        *int
	PortID          *int
	CargoTypeID     *int
	StatusTypeID    *int
	IsActive        *int
	IsCompleted     *int
	IsFraud         *int
	IsPosted        *int
	RouteType       string
	TransactionKind string
	Filters         map[string]string
	DateFrom        *time.Time
	DateTo          *time.Time
}
