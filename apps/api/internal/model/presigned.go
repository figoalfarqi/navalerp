package model

// PresignedRequest dipakai oleh FE untuk meminta presigned URL
type PresignedRequest struct {
	Folder   string `json:"folder" validate:"required"`
	Filename string `json:"filename" validate:"required"`
	Action   string `json:"action" validate:"required,oneof=upload delete"`
}
