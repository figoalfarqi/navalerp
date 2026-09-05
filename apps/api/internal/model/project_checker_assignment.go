package model

import "time"

type ProjectCheckerAssignment struct {
	ProjectCheckerAssignmentID int                              `json:"project_checker_assignment_id"`
	ProjectID                  int                              `json:"project_id"`
	CheckerID                  int                              `json:"checker_id"`
	AccessStartedAt            time.Time                        `json:"access_started_at"`
	AccessEndedAt              *time.Time                       `json:"access_ended_at,omitempty"`
	AssignmentNote             *string                          `json:"assignment_note,omitempty"`
	IsActive                   int                              `json:"is_active"`
	IsDefault                  int                              `json:"is_default"`
	Project                    *ProjectTransportProjectRelation `json:"project,omitempty"`
	Checker                    *ProjectTransportUserRelation    `json:"checker,omitempty"`
	Audit
}

type ProjectCheckerAssignmentRequest struct {
	ProjectID       int        `json:"project_id" validate:"required,gt=0"`
	CheckerID       int        `json:"checker_id" validate:"required,gt=0"`
	AccessStartedAt *time.Time `json:"access_started_at,omitempty"`
	AccessEndedAt   *time.Time `json:"access_ended_at,omitempty"`
	AssignmentNote  *string    `json:"assignment_note,omitempty"`
	IsActive        *int       `json:"is_active,omitempty" validate:"omitempty,oneof=0 1"`
	IsDefault       *int       `json:"is_default,omitempty" validate:"omitempty,oneof=0 1"`
}
