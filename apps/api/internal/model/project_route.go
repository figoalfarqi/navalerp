package model

type ProjectRoute struct {
	ProjectRouteID                    int                              `json:"project_route_id"`
	ProjectID                         int                              `json:"project_id"`
	ProjectPattern                    string                           `json:"project_pattern"`
	RouteSequence                     int                              `json:"route_sequence"`
	RouteType                         string                           `json:"route_type"`
	RouteName                         *string                          `json:"route_name,omitempty"`
	DistanceKM                        *float64                         `json:"distance_km,omitempty"`
	TransportServiceUnit              string                           `json:"transport_service_unit"`
	TransportServicePricePerCubic     *float64                         `json:"transport_service_price_per_cubic,omitempty"`
	TransportServicePricePerTon       *float64                         `json:"transport_service_price_per_ton,omitempty"`
	TransportServicePricePerTransport *float64                         `json:"transport_service_price_per_transport,omitempty"`
	TransportCostUnit                 string                           `json:"transport_cost_unit"`
	TransportCostPerCubic             *float64                         `json:"transport_cost_per_cubic,omitempty"`
	TransportCostPerTon               *float64                         `json:"transport_cost_per_ton,omitempty"`
	TransportCostPerTransport         *float64                         `json:"transport_cost_per_transport,omitempty"`
	RoadMoneyPerTransport             float64                          `json:"road_money_per_transport"`
	LoadingCostPerTransport           float64                          `json:"loading_cost_per_transport"`
	UnloadingCostPerTransport         float64                          `json:"unloading_cost_per_transport"`
	FuelCostPerTransport              float64                          `json:"fuel_cost_per_transport"`
	TollCostPerTransport              float64                          `json:"toll_cost_per_transport"`
	OtherIncomePerTransport           float64                          `json:"other_income_per_transport"`
	OtherExpensePerTransport          float64                          `json:"other_expense_per_transport"`
	RouteNote                         *string                          `json:"route_note,omitempty"`
	IsActive                          int                              `json:"is_active"`
	Project                           *ProjectTransportProjectRelation `json:"project,omitempty"`
	Audit
}

type ProjectRouteRequest struct {
	ProjectID                         int      `json:"project_id" validate:"required,gt=0"`
	ProjectPattern                    string   `json:"project_pattern,omitempty"`
	RouteSequence                     int      `json:"route_sequence" validate:"required,gt=0"`
	RouteType                         string   `json:"route_type" validate:"required"`
	RouteName                         *string  `json:"route_name,omitempty"`
	DistanceKM                        *float64 `json:"distance_km,omitempty"`
	TransportServiceUnit              string   `json:"transport_service_unit" validate:"required,oneof=NONE M3 TON TRANSPORT"`
	TransportServicePricePerCubic     *float64 `json:"transport_service_price_per_cubic,omitempty"`
	TransportServicePricePerTon       *float64 `json:"transport_service_price_per_ton,omitempty"`
	TransportServicePricePerTransport *float64 `json:"transport_service_price_per_transport,omitempty"`
	TransportCostUnit                 string   `json:"transport_cost_unit" validate:"required,oneof=NONE M3 TON TRANSPORT"`
	TransportCostPerCubic             *float64 `json:"transport_cost_per_cubic,omitempty"`
	TransportCostPerTon               *float64 `json:"transport_cost_per_ton,omitempty"`
	TransportCostPerTransport         *float64 `json:"transport_cost_per_transport,omitempty"`
	RoadMoneyPerTransport             *float64 `json:"road_money_per_transport,omitempty"`
	LoadingCostPerTransport           *float64 `json:"loading_cost_per_transport,omitempty"`
	UnloadingCostPerTransport         *float64 `json:"unloading_cost_per_transport,omitempty"`
	FuelCostPerTransport              *float64 `json:"fuel_cost_per_transport,omitempty"`
	TollCostPerTransport              *float64 `json:"toll_cost_per_transport,omitempty"`
	OtherIncomePerTransport           *float64 `json:"other_income_per_transport,omitempty"`
	OtherExpensePerTransport          *float64 `json:"other_expense_per_transport,omitempty"`
	RouteNote                         *string  `json:"route_note,omitempty"`
	IsActive                          *int     `json:"is_active,omitempty" validate:"omitempty,oneof=0 1"`
}
