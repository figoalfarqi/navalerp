package model

import (
	"time"
)

type DailyLog struct {
	LogId                   string    `json:"log_id"`
	ShipId                  string    `json:"ship_id"`
	ShipName                *string   `json:"ship_name,omitempty"`
	LogDate                 time.Time `json:"log_date"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	HeadingDegrees          *int      `json:"heading_degrees,omitempty"`
	SpeedKnots              float64   `json:"speed_knots"`
	SeaState                *int      `json:"sea_state,omitempty"`
	WeatherCondition        *string   `json:"weather_condition,omitempty"`
	FuelRemainingLiters     float64   `json:"fuel_remaining_liters"`
	FreshWaterRemainingTons float64   `json:"fresh_water_remaining_tons"`
	TacticalSummary         *string   `json:"tactical_summary,omitempty"`
	LoggedByUserId          string    `json:"logged_by_user_id"`
	LoggerName              *string   `json:"logger_name,omitempty"`
	CreatedAt               time.Time `json:"created_at"`
}

type DailyLogRequest struct {
	ShipId                  *string    `json:"ship_id"`
	LogDate                 *time.Time `json:"log_date"`
	Latitude                *float64   `json:"latitude"`
	Longitude               *float64   `json:"longitude"`
	HeadingDegrees          *int       `json:"heading_degrees"`
	SpeedKnots              *float64   `json:"speed_knots"`
	SeaState                *int       `json:"sea_state"`
	WeatherCondition        *string    `json:"weather_condition"`
	FuelRemainingLiters     *float64   `json:"fuel_remaining_liters"`
	FreshWaterRemainingTons *float64   `json:"fresh_water_remaining_tons"`
	TacticalSummary         *string    `json:"tactical_summary"`
	LoggedByUserId          *string    `json:"logged_by_user_id"`
}
