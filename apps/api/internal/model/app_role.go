package model

import "time"

type AppRole struct {
	AppRoleID          int        `json:"app_role_id"`
	AppRoleTypeID      int        `json:"app_role_type_id"`
	AppRoleName        string     `json:"app_role_name"`
	AppRoleDescription *string    `json:"app_role_description,omitempty"`
	IsActive           int        `json:"is_active"`
	CreatedBy          int        `json:"created_by"`
	UpdatedBy          int        `json:"updated_by"`
	DeletedBy          *int       `json:"deleted_by,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
}

type AppRoleRequest struct {
	AppRoleTypeID      int     `json:"app_role_type_id" validate:"required,numeric"`
	AppRoleName        string  `json:"app_role_name" validate:"required,min=2,max=50"`
	AppRoleDescription *string `json:"app_role_description,omitempty"`
	IsActive           *int    `json:"is_active,omitempty"`
}

type AppRoleResponse struct {
	AppRoleID          int        `json:"app_role_id"`
	AppRoleTypeID      int        `json:"app_role_type_id"`
	AppRoleName        string     `json:"app_role_name"`
	AppRoleDescription *string    `json:"app_role_description,omitempty"`
	IsActive           int        `json:"is_active"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
}
