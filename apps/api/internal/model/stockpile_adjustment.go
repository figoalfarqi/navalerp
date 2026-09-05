package model

type StockpileAdjustment struct {
	StockpileAdjustmentID int               `json:"stockpile_adjustment_id"`
	ProjectID             int               `json:"project_id"`
	StockpileCargoID      int               `json:"stockpile_cargo_id"`
	AdjustmentDate        string            `json:"adjustment_date"`
	ProjectTransportID    *int              `json:"project_transport_id,omitempty"`
	ReferenceTypeID       int               `json:"reference_type_id"`
	AmountVolumeCubic     float64           `json:"amount_volume_cubic"`
	AmountWeightTon       float64           `json:"amount_weight_ton"`
	Reason                *string           `json:"reason,omitempty"`
	Project               *Project          `json:"project,omitempty"`
	StockpileCargo        *StockpileCargo   `json:"stockpile_cargo,omitempty"`
	ProjectTransport      *ProjectTransport `json:"project_transport,omitempty"`
	Audit
}

type StockpileAdjustmentRequest struct {
	ProjectID          int     `json:"project_id" validate:"required,gt=0"`
	StockpileCargoID   int     `json:"stockpile_cargo_id" validate:"required,gt=0"`
	AdjustmentDate     string  `json:"adjustment_date" validate:"required"`
	ProjectTransportID *int    `json:"project_transport_id,omitempty" validate:"omitempty,gt=0"`
	ReferenceTypeID    int     `json:"reference_type_id" validate:"required,min=1,max=5"`
	AmountVolumeCubic  float64 `json:"amount_volume_cubic"`
	AmountWeightTon    float64 `json:"amount_weight_ton"`
	Reason             *string `json:"reason,omitempty"`
}

type StockpileAdjustmentResponse = StockpileAdjustment
