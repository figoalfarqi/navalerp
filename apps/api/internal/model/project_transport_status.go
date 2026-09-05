package model

import "time"

type ProjectTransportStatus struct {
	ProjectTransportStatusID     int                        `json:"project_transport_status_id"`
	ProjectTransportID           int                        `json:"project_transport_id"`
	ProjectTransportStatusTypeID int                        `json:"project_transport_status_type_id"`
	StatusTime                   time.Time                  `json:"status_time"`
	CargoBoxLength               *float64                   `json:"cargo_box_length,omitempty"`
	CargoBoxWidth                *float64                   `json:"cargo_box_width,omitempty"`
	CargoBoxHeight               *float64                   `json:"cargo_box_height,omitempty"`
	CargoVolumeCubic             *float64                   `json:"cargo_volume_cubic,omitempty"`
	CargoWeightTon               *float64                   `json:"cargo_weight_ton,omitempty"`
	ProjectTransportStatusNote   *string                    `json:"project_transport_status_note,omitempty"`
	IsFraud                      int                        `json:"is_fraud"`
	IsActive                     int                        `json:"is_active"`
	ProjectTransport             *ProjectTransportReference `json:"project_transport,omitempty"`
	Photos                       []ProjectTransportPhoto    `json:"photos,omitempty"`
	Audit
}

type ProjectTransportStatusReference struct {
	ProjectTransportStatusID     int                        `json:"project_transport_status_id"`
	ProjectTransportStatusTypeID int                        `json:"project_transport_status_type_id"`
	StatusTime                   time.Time                  `json:"status_time"`
	ProjectTransport             *ProjectTransportReference `json:"project_transport,omitempty"`
}

type ProjectTransportStatusRequest struct {
	ProjectTransportID           int        `json:"project_transport_id" validate:"required,gt=0"`
	ProjectTransportStatusTypeID int        `json:"project_transport_status_type_id" validate:"required,gt=0"`
	StatusTime                   *time.Time `json:"status_time,omitempty"`
	CargoBoxLength               *float64   `json:"cargo_box_length,omitempty"`
	CargoBoxWidth                *float64   `json:"cargo_box_width,omitempty"`
	CargoBoxHeight               *float64   `json:"cargo_box_height,omitempty"`
	CargoVolumeCubic             *float64   `json:"cargo_volume_cubic,omitempty"`
	CargoWeightTon               *float64   `json:"cargo_weight_ton,omitempty"`
	ProjectTransportStatusNote   *string    `json:"project_transport_status_note,omitempty"`
	IsFraud                      *int       `json:"is_fraud,omitempty" validate:"omitempty,oneof=0 1"`
	IsActive                     *int       `json:"is_active,omitempty" validate:"omitempty,oneof=0 1"`
}
