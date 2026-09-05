package model

import "time"

type ProjectTruckAssignmentTruckTypeRelation struct {
	TruckTypeID   int    `json:"truck_type_id"`
	TruckTypeName string `json:"truck_type_name"`
}

type ProjectTruckAssignmentTruckMerkRelation struct {
	TruckMerkID   int    `json:"truck_merk_id"`
	TruckMerkName string `json:"truck_merk_name"`
}

type ProjectTruckAssignmentTruckRelation struct {
	TruckID      int                                      `json:"truck_id"`
	LicensePlate string                                   `json:"license_plate"`
	Driver       *ProjectTransportUserRelation            `json:"driver,omitempty"`
	Vendor       *ProjectTransportVendorRelation          `json:"vendor,omitempty"`
	TruckType    *ProjectTruckAssignmentTruckTypeRelation `json:"truck_type,omitempty"`
	TruckMerk    *ProjectTruckAssignmentTruckMerkRelation `json:"truck_merk,omitempty"`
}

type ProjectTruckAssignment struct {
	ProjectTruckAssignmentID int                                  `json:"project_truck_assignment_id"`
	ProjectID                int                                  `json:"project_id"`
	TruckID                  int                                  `json:"truck_id"`
	AssignmentStartedAt      time.Time                            `json:"assignment_started_at"`
	AssignmentEndedAt        *time.Time                           `json:"assignment_ended_at,omitempty"`
	AssignmentNote           *string                              `json:"assignment_note,omitempty"`
	IsActive                 int                                  `json:"is_active"`
	Project                  *ProjectTransportProjectRelation     `json:"project,omitempty"`
	Truck                    *ProjectTruckAssignmentTruckRelation `json:"truck,omitempty"`
	Audit
}

type ProjectTruckAssignmentRequest struct {
	ProjectID           int        `json:"project_id" validate:"required,gt=0"`
	TruckID             int        `json:"truck_id" validate:"required,gt=0"`
	AssignmentStartedAt *time.Time `json:"assignment_started_at,omitempty"`
	AssignmentEndedAt   *time.Time `json:"assignment_ended_at,omitempty"`
	AssignmentNote      *string    `json:"assignment_note,omitempty"`
	IsActive            *int       `json:"is_active,omitempty" validate:"omitempty,oneof=0 1"`
}
