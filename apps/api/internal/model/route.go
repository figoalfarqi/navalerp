package model

import (
	"time"
)

type Route struct {
	RouteId string `json:"route_id"`
	RouteCode string `json:"route_code"`
	RouteName string `json:"route_name"`
	OriginFacilityId string `json:"origin_facility_id"`
	DestinationFacilityId string `json:"destination_facility_id"`
	DistanceNauticalMiles float64 `json:"distance_nautical_miles"`
	EstimatedTransitHours float64 `json:"estimated_transit_hours"`
	RiskLevel *string `json:"risk_level,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type RouteRequest struct {
	RouteCode *string `json:"route_code"`
	RouteName *string `json:"route_name"`
	OriginFacilityId *string `json:"origin_facility_id"`
	DestinationFacilityId *string `json:"destination_facility_id"`
	DistanceNauticalMiles *float64 `json:"distance_nautical_miles"`
	EstimatedTransitHours *float64 `json:"estimated_transit_hours"`
	RiskLevel *string `json:"risk_level"`
}
