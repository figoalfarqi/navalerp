package model

import (
	"time"
)

type Document struct {
	DocumentId string `json:"document_id"`
	DocumentNumber string `json:"document_number"`
	Title string `json:"title"`
	CategoryId string `json:"category_id"`
	OriginatingUnitId *string `json:"originating_unit_id,omitempty"`
	ClassificationLevel *string `json:"classification_level,omitempty"`
	EffectiveDate time.Time `json:"effective_date"`
	ExpiryDate *time.Time `json:"expiry_date,omitempty"`
	Status *string `json:"status,omitempty"`
	ApprovedByUserId *string `json:"approved_by_user_id,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Versions []DocumentVersions `json:"versions,omitempty"`
	Links []DocumentLinks `json:"links,omitempty"`
}

type DocumentVersions struct {
	VersionId string `json:"version_id"`
	DocumentId string `json:"document_id"`
	VersionNumber string `json:"version_number"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	FileSizeBytes int64 `json:"file_size_bytes"`
	FileHashSha256 string `json:"file_hash_sha256"`
	MimeType string `json:"mime_type"`
	ChangeSummary *string `json:"change_summary,omitempty"`
	UploadedByUserId *string `json:"uploaded_by_user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type DocumentLinks struct {
	LinkId string `json:"link_id"`
	DocumentId string `json:"document_id"`
	EntityType string `json:"entity_type"`
	EntityId string `json:"entity_id"`
	LinkPurpose string `json:"link_purpose"`
	CreatedAt time.Time `json:"created_at"`
}

type DocumentRequest struct {
	DocumentNumber *string `json:"document_number"`
	Title *string `json:"title"`
	CategoryId *string `json:"category_id"`
	OriginatingUnitId *string `json:"originating_unit_id"`
	ClassificationLevel *string `json:"classification_level"`
	EffectiveDate *time.Time `json:"effective_date"`
	ExpiryDate *time.Time `json:"expiry_date"`
	Status *string `json:"status"`
	ApprovedByUserId *string `json:"approved_by_user_id"`
	Versions []DocumentVersions `json:"versions"`
	Links []DocumentLinks `json:"links"`
}
