package model

type ProjectTransportPhoto struct {
	ProjectTransportPhotoID  int                              `json:"project_transport_photo_id"`
	ProjectTransportStatusID int                              `json:"project_transport_status_id"`
	PhotoURL                 string                           `json:"photo_url"`
	PhotoTypeID              int                              `json:"photo_type_id"`
	PhotoDescription         *string                          `json:"photo_description,omitempty"`
	ProjectTransportStatus   *ProjectTransportStatusReference `json:"project_transport_status,omitempty"`
	Audit
}

type ProjectTransportPhotoRequest struct {
	ProjectTransportStatusID int     `json:"project_transport_status_id" validate:"required,gt=0"`
	PhotoURL                 string  `json:"photo_url" validate:"required"`
	PhotoTypeID              int     `json:"photo_type_id" validate:"required,gt=0"`
	PhotoDescription         *string `json:"photo_description,omitempty"`
}
