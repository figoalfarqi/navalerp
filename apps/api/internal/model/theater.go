package model

import (
	"time"
)

type Theater struct {
	TheaterId string `json:"theater_id"`
	TheaterCode string `json:"theater_code"`
	TheaterName string `json:"theater_name"`
	ResponsibleCommandUnitId string `json:"responsible_command_unit_id"`
	ThreatLevel *string `json:"threat_level,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type TheaterRequest struct {
	TheaterCode *string `json:"theater_code"`
	TheaterName *string `json:"theater_name"`
	ResponsibleCommandUnitId *string `json:"responsible_command_unit_id"`
	ThreatLevel *string `json:"threat_level"`
	Description *string `json:"description"`
}
