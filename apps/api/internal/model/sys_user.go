package model

import (
	"time"
)

type SysUser struct {
	UserId string `json:"user_id"`
	UnitId string `json:"unit_id"`
	UnitName *string `json:"unit_name,omitempty"`
	Username string `json:"username"`
	PasswordHash string `json:"password_hash"`
	FullName string `json:"full_name"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
	MilitaryId *string `json:"military_id,omitempty"`
	RankTitle *string `json:"rank_title,omitempty"`
	Department *string `json:"department,omitempty"`
	Role string `json:"role"`
	IsActive *bool `json:"is_active,omitempty"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	FailedLoginAttempts *int `json:"failed_login_attempts,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	AuthVersion int `json:"auth_version"`
}

type SysUserRequest struct {
	UnitId *string `json:"unit_id"`
	Username *string `json:"username"`
	PasswordHash *string `json:"password_hash"`
	FullName *string `json:"full_name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	MilitaryId *string `json:"military_id"`
	RankTitle *string `json:"rank_title"`
	Department *string `json:"department"`
	Role *string `json:"role"`
	IsActive *bool `json:"is_active"`
	LastLoginAt *time.Time `json:"last_login_at"`
	FailedLoginAttempts *int `json:"failed_login_attempts"`
	AuthVersion *int `json:"auth_version"`
}
