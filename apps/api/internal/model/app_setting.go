package model

import "encoding/json"

type AppSetting struct {
	AppSettingID          int             `json:"app_setting_id"`
	AppSettingKey         string          `json:"app_setting_key"`
	AppSettingValue       json.RawMessage `json:"app_setting_value"`
	AppSettingDescription *string         `json:"app_setting_description,omitempty"`
	IsActive              int             `json:"is_active"`
	Audit
}

type AppSettingRequest struct {
	AppSettingKey         string          `json:"app_setting_key" validate:"required,min=2,max=100"`
	AppSettingValue       json.RawMessage `json:"app_setting_value" validate:"required"`
	AppSettingDescription *string         `json:"app_setting_description,omitempty"`
	IsActive              *int            `json:"is_active,omitempty" validate:"omitempty,oneof=0 1"`
}
