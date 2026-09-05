package model

type Project struct {
	ProjectID                 int                `json:"project_id"`
	ProjectCode               string             `json:"project_code"`
	ProjectName               string             `json:"project_name"`
	RouteType                 string             `json:"route_type"`
	MineID                    *int               `json:"mine_id,omitempty"`
	VesselCargoID             *int               `json:"vessel_cargo_id,omitempty"`
	StockpileCargoID          *int               `json:"stockpile_cargo_id,omitempty"`
	ClientDestinationID       int                `json:"client_destination_id"`
	CargoTypeID               int                `json:"cargo_type_id"`
	ProjectStatus             string             `json:"project_status"`
	StartDate                 *string            `json:"start_date,omitempty"`
	EndDate                   *string            `json:"end_date,omitempty"`
	PlannedVolumeCubic        *float64           `json:"planned_volume_cubic,omitempty"`
	PlannedWeightTon          *float64           `json:"planned_weight_ton,omitempty"`
	VolumeToWeightConversion  *float64           `json:"volume_to_weight_conversion,omitempty"`
	MaterialPurchaseUnit      string             `json:"material_purchase_unit"`
	MaterialBuyPricePerCubic  *float64           `json:"material_buy_price_per_cubic,omitempty"`
	MaterialBuyPricePerTon    *float64           `json:"material_buy_price_per_ton,omitempty"`
	MaterialSaleUnit          string             `json:"material_sale_unit"`
	MaterialSellPricePerCubic *float64           `json:"material_sell_price_per_cubic,omitempty"`
	MaterialSellPricePerTon   *float64           `json:"material_sell_price_per_ton,omitempty"`
	FixedOtherIncome          float64            `json:"fixed_other_income"`
	FixedOtherExpense         float64            `json:"fixed_other_expense"`
	ProjectNote               *string            `json:"project_note,omitempty"`
	IsActive                  int                `json:"is_active"`
	IsDefault                 int                `json:"is_default,omitempty"`
	Mine                      *Mine              `json:"mine,omitempty"`
	VesselCargo               *VesselCargo       `json:"vessel_cargo,omitempty"`
	StockpileCargo            *StockpileCargo    `json:"stockpile_cargo,omitempty"`
	ClientDestination         *ClientDestination `json:"client_destination,omitempty"`
	CargoType                 *CargoType         `json:"cargo_type,omitempty"`
	Audit
}

// CheckerProject is the operational project contract exposed to checker users.
// Pricing and other financial fields intentionally stay in the admin-only Project DTO.
type CheckerProject struct {
	ProjectID           int     `json:"project_id"`
	ProjectCode         string  `json:"project_code"`
	ProjectName         string  `json:"project_name"`
	RouteType           string  `json:"route_type"`
	MineID              *int    `json:"mine_id,omitempty"`
	VesselCargoID       *int    `json:"vessel_cargo_id,omitempty"`
	StockpileCargoID    *int    `json:"stockpile_cargo_id,omitempty"`
	ClientDestinationID int     `json:"client_destination_id"`
	CargoTypeID         int     `json:"cargo_type_id"`
	ProjectStatus       string  `json:"project_status"`
	StartDate           *string `json:"start_date,omitempty"`
	EndDate             *string `json:"end_date,omitempty"`
	ProjectNote         *string `json:"project_note,omitempty"`
	IsActive            int     `json:"is_active"`
	IsDefault           int     `json:"is_default,omitempty"`
}

func NewCheckerProject(project Project) CheckerProject {
	return CheckerProject{
		ProjectID:           project.ProjectID,
		ProjectCode:         project.ProjectCode,
		ProjectName:         project.ProjectName,
		RouteType:           project.RouteType,
		MineID:              project.MineID,
		VesselCargoID:       project.VesselCargoID,
		StockpileCargoID:    project.StockpileCargoID,
		ClientDestinationID: project.ClientDestinationID,
		CargoTypeID:         project.CargoTypeID,
		ProjectStatus:       project.ProjectStatus,
		StartDate:           project.StartDate,
		EndDate:             project.EndDate,
		ProjectNote:         project.ProjectNote,
		IsActive:            project.IsActive,
		IsDefault:           project.IsDefault,
	}
}

type ProjectRequest struct {
	ProjectCode               string   `json:"project_code" validate:"required,max=50"`
	ProjectName               string   `json:"project_name" validate:"required,max=200"`
	RouteType                 string   `json:"route_type" validate:"required,oneof=MINE_CLIENT MINE_STOCKPILE_CLIENT VESSEL_CLIENT VESSEL_STOCKPILE_CLIENT"`
	MineID                    *int     `json:"mine_id,omitempty"`
	VesselCargoID             *int     `json:"vessel_cargo_id,omitempty"`
	StockpileCargoID          *int     `json:"stockpile_cargo_id,omitempty"`
	ClientDestinationID       int      `json:"client_destination_id" validate:"required,gt=0"`
	CargoTypeID               int      `json:"cargo_type_id" validate:"required,gt=0"`
	ProjectStatus             string   `json:"project_status" validate:"required,oneof=DRAFT ACTIVE COMPLETED CANCELLED"`
	StartDate                 *string  `json:"start_date,omitempty"`
	EndDate                   *string  `json:"end_date,omitempty"`
	PlannedVolumeCubic        *float64 `json:"planned_volume_cubic,omitempty"`
	PlannedWeightTon          *float64 `json:"planned_weight_ton,omitempty"`
	VolumeToWeightConversion  *float64 `json:"volume_to_weight_conversion,omitempty"`
	MaterialPurchaseUnit      string   `json:"material_purchase_unit" validate:"required,oneof=NONE M3 TON"`
	MaterialBuyPricePerCubic  *float64 `json:"material_buy_price_per_cubic,omitempty"`
	MaterialBuyPricePerTon    *float64 `json:"material_buy_price_per_ton,omitempty"`
	MaterialSaleUnit          string   `json:"material_sale_unit" validate:"required,oneof=NONE M3 TON"`
	MaterialSellPricePerCubic *float64 `json:"material_sell_price_per_cubic,omitempty"`
	MaterialSellPricePerTon   *float64 `json:"material_sell_price_per_ton,omitempty"`
	FixedOtherIncome          *float64 `json:"fixed_other_income,omitempty"`
	FixedOtherExpense         *float64 `json:"fixed_other_expense,omitempty"`
	ProjectNote               *string  `json:"project_note,omitempty"`
	IsActive                  *int     `json:"is_active,omitempty" validate:"omitempty,oneof=0 1"`
}
