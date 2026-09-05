package model

import (
	"time"
)

type DocumentCategory struct {
	CategoryId string `json:"category_id"`
	CategoryCode string `json:"category_code"`
	CategoryName string `json:"category_name"`
	RetentionYears *int `json:"retention_years,omitempty"`
	ConfidentialityLevel *string `json:"confidentiality_level,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type DocumentCategoryRequest struct {
	CategoryCode *string `json:"category_code"`
	CategoryName *string `json:"category_name"`
	RetentionYears *int `json:"retention_years"`
	ConfidentialityLevel *string `json:"confidentiality_level"`
	Description *string `json:"description"`
}
