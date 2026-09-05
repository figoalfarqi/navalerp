package model

import (
	"time"
)

type ReadinessAlert struct {
	AlertId string `json:"alert_id"`
	ShipId string `json:"ship_id"`
	EquipmentId *string `json:"equipment_id,omitempty"`
	Severity string `json:"severity"`
	AlertType string `json:"alert_type"`
	AlertMessage string `json:"alert_message"`
	IsAcknowledged *bool `json:"is_acknowledged,omitempty"`
	AcknowledgedByUserId *string `json:"acknowledged_by_user_id,omitempty"`
	AcknowledgedAt *time.Time `json:"acknowledged_at,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type ReadinessAlertRequest struct {
	ShipId *string `json:"ship_id"`
	EquipmentId *string `json:"equipment_id"`
	Severity *string `json:"severity"`
	AlertType *string `json:"alert_type"`
	AlertMessage *string `json:"alert_message"`
	IsAcknowledged *bool `json:"is_acknowledged"`
	AcknowledgedByUserId *string `json:"acknowledged_by_user_id"`
	AcknowledgedAt *time.Time `json:"acknowledged_at"`
}
