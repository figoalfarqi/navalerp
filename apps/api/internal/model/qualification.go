package model

import (
	"time"
)

type Qualification struct {
	QualificationId string `json:"qualification_id"`
	QualificationCode string `json:"qualification_code"`
	QualificationName string `json:"qualification_name"`
	QualificationCategory string `json:"qualification_category"`
	IssuingInstitution string `json:"issuing_institution"`
	ValidityYears *int `json:"validity_years,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type QualificationRequest struct {
	QualificationCode *string `json:"qualification_code"`
	QualificationName *string `json:"qualification_name"`
	QualificationCategory *string `json:"qualification_category"`
	IssuingInstitution *string `json:"issuing_institution"`
	ValidityYears *int `json:"validity_years"`
	Description *string `json:"description"`
}
