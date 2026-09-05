package model

import "time"

type ProjectTransportProjectRelation struct {
	ProjectID   int    `json:"project_id"`
	ProjectCode string `json:"project_code"`
	ProjectName string `json:"project_name"`
}

type ProjectTransportRouteRelation struct {
	ProjectRouteID int     `json:"project_route_id"`
	RouteName      *string `json:"route_name,omitempty"`
	RouteType      string  `json:"route_type"`
}

type ProjectTransportTruckRelation struct {
	TruckID      int    `json:"truck_id"`
	LicensePlate string `json:"license_plate"`
}

type ProjectTransportUserRelation struct {
	AppUserID   int    `json:"app_user_id"`
	AppUserName string `json:"app_user_name"`
	Username    string `json:"username"`
}

type ProjectTransportVendorRelation struct {
	VendorID   int    `json:"vendor_id"`
	VendorName string `json:"vendor_name"`
}

type ProjectTransportLatestStatusRelation struct {
	ProjectTransportStatusTypeID int       `json:"project_transport_status_type_id"`
	StatusTime                   time.Time `json:"status_time"`
}

type ProjectTransportReference struct {
	ProjectTransportID int                              `json:"project_transport_id"`
	TransportNumber    string                           `json:"transport_number"`
	Project            *ProjectTransportProjectRelation `json:"project,omitempty"`
	ProjectRoute       *ProjectTransportRouteRelation   `json:"project_route,omitempty"`
	Truck              *ProjectTransportTruckRelation   `json:"truck,omitempty"`
	Driver             *ProjectTransportUserRelation    `json:"driver,omitempty"`
	TransportVendor    *ProjectTransportVendorRelation  `json:"transport_vendor,omitempty"`
}

type ProjectTransport struct {
	ProjectTransportID                int                                   `json:"project_transport_id"`
	ProjectID                         int                                   `json:"project_id"`
	ProjectRouteID                    int                                   `json:"project_route_id"`
	TransportNumber                   string                                `json:"transport_number"`
	DeliveryNoteNumber                *string                               `json:"delivery_note_number,omitempty"`
	TransportedAt                     time.Time                             `json:"transported_at"`
	TruckID                           *int                                  `json:"truck_id,omitempty"`
	ProjectTruckAssignmentID          *int                                  `json:"project_truck_assignment_id,omitempty"`
	DriverID                          *int                                  `json:"driver_id,omitempty"`
	TransportVendorID                 *int                                  `json:"transport_vendor_id,omitempty"`
	LoadedVolumeCubic                 *float64                              `json:"loaded_volume_cubic,omitempty"`
	LoadedWeightTon                   *float64                              `json:"loaded_weight_ton,omitempty"`
	DeliveredVolumeCubic              *float64                              `json:"delivered_volume_cubic,omitempty"`
	DeliveredWeightTon                *float64                              `json:"delivered_weight_ton,omitempty"`
	PurchaseVolumeCubic               *float64                              `json:"purchase_volume_cubic,omitempty"`
	PurchaseWeightTon                 *float64                              `json:"purchase_weight_ton,omitempty"`
	SaleVolumeCubic                   *float64                              `json:"sale_volume_cubic,omitempty"`
	SaleWeightTon                     *float64                              `json:"sale_weight_ton,omitempty"`
	TransportServiceVolumeCubic       *float64                              `json:"transport_service_volume_cubic,omitempty"`
	TransportServiceWeightTon         *float64                              `json:"transport_service_weight_ton,omitempty"`
	TransportCostVolumeCubic          *float64                              `json:"transport_cost_volume_cubic,omitempty"`
	TransportCostWeightTon            *float64                              `json:"transport_cost_weight_ton,omitempty"`
	VolumeToWeightConversion          *float64                              `json:"volume_to_weight_conversion,omitempty"`
	MaterialPurchaseUnit              string                                `json:"material_purchase_unit"`
	MaterialBuyPricePerCubic          *float64                              `json:"material_buy_price_per_cubic,omitempty"`
	MaterialBuyPricePerTon            *float64                              `json:"material_buy_price_per_ton,omitempty"`
	MaterialPurchaseAmount            float64                               `json:"material_purchase_amount"`
	MaterialSaleUnit                  string                                `json:"material_sale_unit"`
	MaterialSellPricePerCubic         *float64                              `json:"material_sell_price_per_cubic,omitempty"`
	MaterialSellPricePerTon           *float64                              `json:"material_sell_price_per_ton,omitempty"`
	MaterialSaleAmount                float64                               `json:"material_sale_amount"`
	TransportServiceUnit              string                                `json:"transport_service_unit"`
	TransportServicePricePerCubic     *float64                              `json:"transport_service_price_per_cubic,omitempty"`
	TransportServicePricePerTon       *float64                              `json:"transport_service_price_per_ton,omitempty"`
	TransportServicePricePerTransport *float64                              `json:"transport_service_price_per_transport,omitempty"`
	TransportServiceIncomeAmount      float64                               `json:"transport_service_income_amount"`
	TransportCostUnit                 string                                `json:"transport_cost_unit"`
	TransportCostPerCubic             *float64                              `json:"transport_cost_per_cubic,omitempty"`
	TransportCostPerTon               *float64                              `json:"transport_cost_per_ton,omitempty"`
	TransportCostPerTransport         *float64                              `json:"transport_cost_per_transport,omitempty"`
	TransportExpenseAmount            float64                               `json:"transport_expense_amount"`
	RoadMoneyAmount                   float64                               `json:"road_money_amount"`
	LoadingCostAmount                 float64                               `json:"loading_cost_amount"`
	UnloadingCostAmount               float64                               `json:"unloading_cost_amount"`
	FuelCostAmount                    float64                               `json:"fuel_cost_amount"`
	TollCostAmount                    float64                               `json:"toll_cost_amount"`
	OtherIncomeAmount                 float64                               `json:"other_income_amount"`
	OtherExpenseAmount                float64                               `json:"other_expense_amount"`
	IsCompleted                       int                                   `json:"is_completed"`
	TransportNote                     *string                               `json:"transport_note,omitempty"`
	Project                           *ProjectTransportProjectRelation      `json:"project,omitempty"`
	ProjectRoute                      *ProjectTransportRouteRelation        `json:"project_route,omitempty"`
	Truck                             *ProjectTransportTruckRelation        `json:"truck,omitempty"`
	Driver                            *ProjectTransportUserRelation         `json:"driver,omitempty"`
	TransportVendor                   *ProjectTransportVendorRelation       `json:"transport_vendor,omitempty"`
	LatestStatus                      *ProjectTransportLatestStatusRelation `json:"latest_status,omitempty"`
	StatusCount                       int                                   `json:"status_count"`
	PhotoCount                        int                                   `json:"photo_count"`
	Statuses                          []ProjectTransportStatus              `json:"statuses,omitempty"`
	Audit
}

type ProjectTransportRequest struct {
	ProjectID                   int        `json:"project_id" validate:"required,gt=0"`
	ProjectRouteID              int        `json:"project_route_id" validate:"required,gt=0"`
	TransportNumber             *string    `json:"transport_number,omitempty"`
	DeliveryNoteNumber          *string    `json:"delivery_note_number,omitempty"`
	TransportedAt               *time.Time `json:"transported_at,omitempty"`
	TruckID                     *int       `json:"truck_id,omitempty"`
	ProjectTruckAssignmentID    *int       `json:"project_truck_assignment_id,omitempty"`
	DriverID                    *int       `json:"driver_id,omitempty"`
	TransportVendorID           *int       `json:"transport_vendor_id,omitempty"`
	LoadedVolumeCubic           *float64   `json:"loaded_volume_cubic,omitempty"`
	LoadedWeightTon             *float64   `json:"loaded_weight_ton,omitempty"`
	DeliveredVolumeCubic        *float64   `json:"delivered_volume_cubic,omitempty"`
	DeliveredWeightTon          *float64   `json:"delivered_weight_ton,omitempty"`
	PurchaseVolumeCubic         *float64   `json:"purchase_volume_cubic,omitempty"`
	PurchaseWeightTon           *float64   `json:"purchase_weight_ton,omitempty"`
	SaleVolumeCubic             *float64   `json:"sale_volume_cubic,omitempty"`
	SaleWeightTon               *float64   `json:"sale_weight_ton,omitempty"`
	TransportServiceVolumeCubic *float64   `json:"transport_service_volume_cubic,omitempty"`
	TransportServiceWeightTon   *float64   `json:"transport_service_weight_ton,omitempty"`
	TransportCostVolumeCubic    *float64   `json:"transport_cost_volume_cubic,omitempty"`
	TransportCostWeightTon      *float64   `json:"transport_cost_weight_ton,omitempty"`
	RoadMoneyAmount             *float64   `json:"road_money_amount,omitempty"`
	LoadingCostAmount           *float64   `json:"loading_cost_amount,omitempty"`
	UnloadingCostAmount         *float64   `json:"unloading_cost_amount,omitempty"`
	FuelCostAmount              *float64   `json:"fuel_cost_amount,omitempty"`
	TollCostAmount              *float64   `json:"toll_cost_amount,omitempty"`
	OtherIncomeAmount           *float64   `json:"other_income_amount,omitempty"`
	OtherExpenseAmount          *float64   `json:"other_expense_amount,omitempty"`
	IsCompleted                 *int       `json:"is_completed,omitempty" validate:"omitempty,oneof=0 1"`
	TransportNote               *string    `json:"transport_note,omitempty"`
}

// DriverProjectTransport deliberately excludes every price and financial field.
type DriverProjectTransport struct {
	ProjectTransportID   int                      `json:"project_transport_id"`
	ProjectID            int                      `json:"project_id"`
	ProjectCode          string                   `json:"project_code"`
	ProjectName          string                   `json:"project_name"`
	ProjectRouteID       int                      `json:"project_route_id"`
	RouteName            *string                  `json:"route_name,omitempty"`
	RouteType            string                   `json:"route_type"`
	OriginName           string                   `json:"origin_name"`
	DestinationName      string                   `json:"destination_name"`
	DistanceKM           *float64                 `json:"distance_km,omitempty"`
	TransportNumber      string                   `json:"transport_number"`
	TransportedAt        time.Time                `json:"transported_at"`
	TruckID              *int                     `json:"truck_id,omitempty"`
	LicensePlate         *string                  `json:"license_plate,omitempty"`
	LoadedVolumeCubic    *float64                 `json:"loaded_volume_cubic,omitempty"`
	LoadedWeightTon      *float64                 `json:"loaded_weight_ton,omitempty"`
	DeliveredVolumeCubic *float64                 `json:"delivered_volume_cubic,omitempty"`
	DeliveredWeightTon   *float64                 `json:"delivered_weight_ton,omitempty"`
	IsCompleted          int                      `json:"is_completed"`
	Statuses             []ProjectTransportStatus `json:"statuses"`
}
