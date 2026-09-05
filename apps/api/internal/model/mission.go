package model

import (
	"time"
)

type Mission struct {
	MissionId string `json:"mission_id"`
	TheaterId string `json:"theater_id"`
	MissionCode string `json:"mission_code"`
	MissionName string `json:"mission_name"`
	MissionType string `json:"mission_type"`
	StartDate time.Time `json:"start_date"`
	EndDate *time.Time `json:"end_date,omitempty"`
	CommandingOfficerUserId *string `json:"commanding_officer_user_id,omitempty"`
	MissionStatus *string `json:"mission_status,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	AssignedShips []MissionShipAssignments `json:"assignedships,omitempty"`
}

type MissionShipAssignments struct {
	AssignmentId string `json:"assignment_id"`
	MissionId string `json:"mission_id"`
	ShipId string `json:"ship_id"`
	TacticalCallsign *string `json:"tactical_callsign,omitempty"`
	RoleInTaskForce string `json:"role_in_task_force"`
	JoinedDate time.Time `json:"joined_date"`
	ReleasedDate *time.Time `json:"released_date,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type MissionRequest struct {
	TheaterId *string `json:"theater_id"`
	MissionCode *string `json:"mission_code"`
	MissionName *string `json:"mission_name"`
	MissionType *string `json:"mission_type"`
	StartDate *time.Time `json:"start_date"`
	EndDate *time.Time `json:"end_date"`
	CommandingOfficerUserId *string `json:"commanding_officer_user_id"`
	MissionStatus *string `json:"mission_status"`
	AssignedShips []MissionShipAssignments `json:"assignedships"`
}
