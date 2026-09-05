package model

import (
	"time"
)

type AuditLog struct {
	LogId string `json:"log_id"`
	UserId *string `json:"user_id,omitempty"`
	Action string `json:"action"`
	EntityTable string `json:"entity_table"`
	EntityId *string `json:"entity_id,omitempty"`
	OldValues any `json:"old_values,omitempty"`
	NewValues any `json:"new_values,omitempty"`
	IpAddress *string `json:"ip_address,omitempty"`
	UserAgent *string `json:"user_agent,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type AuditLogRequest struct {
	UserId *string `json:"user_id"`
	Action *string `json:"action"`
	EntityTable *string `json:"entity_table"`
	EntityId *string `json:"entity_id"`
	OldValues any `json:"old_values"`
	NewValues any `json:"new_values"`
	IpAddress *string `json:"ip_address"`
	UserAgent *string `json:"user_agent"`
}
