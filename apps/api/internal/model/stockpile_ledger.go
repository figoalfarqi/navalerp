package model

type StockpileLedger struct {
	StockpileLedgerID  int                  `json:"stockpile_ledger_id"`
	ProjectID          int                  `json:"project_id"`
	StockpileCargoID   int                  `json:"stockpile_cargo_id"`
	ProjectTransportID *int                 `json:"project_transport_id,omitempty"`
	AdjustmentID       *int                 `json:"adjustment_id,omitempty"`
	ReferenceTypeID    int                  `json:"reference_type_id"`
	AmountVolumeCubic  float64              `json:"amount_volume_cubic"`
	BalanceVolumeCubic float64              `json:"balance_volume_cubic"`
	AmountWeightTon    float64              `json:"amount_weight_ton"`
	BalanceWeightTon   float64              `json:"balance_weight_ton"`
	Project            *Project             `json:"project,omitempty"`
	StockpileCargo     *StockpileCargo      `json:"stockpile_cargo,omitempty"`
	ProjectTransport   *ProjectTransport    `json:"project_transport,omitempty"`
	Adjustment         *StockpileAdjustment `json:"adjustment,omitempty"`
	Audit
}

type StockpileLedgerRequest struct {
	ProjectID          int     `json:"project_id" validate:"required,gt=0"`
	StockpileCargoID   int     `json:"stockpile_cargo_id" validate:"required,gt=0"`
	ProjectTransportID *int    `json:"project_transport_id,omitempty" validate:"omitempty,gt=0"`
	AdjustmentID       *int    `json:"adjustment_id,omitempty" validate:"omitempty,gt=0"`
	ReferenceTypeID    int     `json:"reference_type_id" validate:"required,min=1,max=5"`
	AmountVolumeCubic  float64 `json:"amount_volume_cubic"`
	BalanceVolumeCubic float64 `json:"balance_volume_cubic"`
	AmountWeightTon    float64 `json:"amount_weight_ton"`
	BalanceWeightTon   float64 `json:"balance_weight_ton"`
}

type StockpileLedgerResponse = StockpileLedger
