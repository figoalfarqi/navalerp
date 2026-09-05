package model

import (
	"time"
)

type MilitaryCorps struct {
	CorpsId string `json:"corps_id"`
	CorpsCode string `json:"corps_code"`
	CorpsName string `json:"corps_name"`
	Description *string `json:"description,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type MilitaryCorpsRequest struct {
	CorpsCode *string `json:"corps_code"`
	CorpsName *string `json:"corps_name"`
	Description *string `json:"description"`
}
