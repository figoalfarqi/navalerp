package model

import (
	"time"
)

type Personnel struct {
	PersonnelId string `json:"personnel_id"`
	Nrp string `json:"nrp"`
	FullName string `json:"full_name"`
	RankId string `json:"rank_id"`
	RankName *string `json:"rank_name,omitempty"`
	CorpsId string `json:"corps_id"`
	CorpsName *string `json:"corps_name,omitempty"`
	CurrentUnitId string `json:"current_unit_id"`
	UnitName *string `json:"unit_name,omitempty"`
	CurrentPosition string `json:"current_position"`
	BirthPlace *string `json:"birth_place,omitempty"`
	BirthDate time.Time `json:"birth_date"`
	Gender *string `json:"gender,omitempty"`
	BloodType *string `json:"blood_type,omitempty"`
	Religion *string `json:"religion,omitempty"`
	EducationLevel *string `json:"education_level,omitempty"`
	ServiceEntryDate time.Time `json:"service_entry_date"`
	UserId *string `json:"user_id,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	ServiceRecords []ServiceRecords `json:"servicerecords,omitempty"`
	Qualifications []PersonnelQualifications `json:"qualifications,omitempty"`
	MedicalReadiness []MedicalReadiness `json:"medicalreadiness,omitempty"`
}

type ServiceRecords struct {
	RecordId string `json:"record_id"`
	PersonnelId string `json:"personnel_id"`
	OrderLetterNumber string `json:"order_letter_number"`
	AssignmentType string `json:"assignment_type"`
	FromUnitId *string `json:"from_unit_id,omitempty"`
	ToUnitId string `json:"to_unit_id"`
	PositionTitle string `json:"position_title"`
	StartDate time.Time `json:"start_date"`
	EndDate *time.Time `json:"end_date,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type PersonnelQualifications struct {
	PersonnelQualId string `json:"personnel_qual_id"`
	PersonnelId string `json:"personnel_id"`
	QualificationId string `json:"qualification_id"`
	CertificateNumber *string `json:"certificate_number,omitempty"`
	ObtainedDate time.Time `json:"obtained_date"`
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type MedicalReadiness struct {
	MedicalId string `json:"medical_id"`
	PersonnelId string `json:"personnel_id"`
	ExaminationDate time.Time `json:"examination_date"`
	StakesCategory string `json:"stakes_category"`
	PhysicalFitnessScore *float64 `json:"physical_fitness_score,omitempty"`
	VisionStatus *string `json:"vision_status,omitempty"`
	DentalStatus *string `json:"dental_status,omitempty"`
	CardioStatus *string `json:"cardio_status,omitempty"`
	GeneralHealthStatus string `json:"general_health_status"`
	DoctorRemarks *string `json:"doctor_remarks,omitempty"`
	ValidUntil time.Time `json:"valid_until"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type PersonnelRequest struct {
	Nrp *string `json:"nrp"`
	FullName *string `json:"full_name"`
	RankId *string `json:"rank_id"`
	CorpsId *string `json:"corps_id"`
	CurrentUnitId *string `json:"current_unit_id"`
	CurrentPosition *string `json:"current_position"`
	BirthPlace *string `json:"birth_place"`
	BirthDate *time.Time `json:"birth_date"`
	Gender *string `json:"gender"`
	BloodType *string `json:"blood_type"`
	Religion *string `json:"religion"`
	EducationLevel *string `json:"education_level"`
	ServiceEntryDate *time.Time `json:"service_entry_date"`
	UserId *string `json:"user_id"`
	Status *string `json:"status"`
	ServiceRecords []ServiceRecords `json:"servicerecords"`
	Qualifications []PersonnelQualifications `json:"qualifications"`
	MedicalReadiness []MedicalReadiness `json:"medicalreadiness"`
}
