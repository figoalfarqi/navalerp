package model

type Port struct {
	PortID        int      `json:"port_id"`
	PortName      string   `json:"port_name"`
	CityID        *int     `json:"city_id,omitempty"`
	PortAddress   *string  `json:"port_address,omitempty"`
	PortLatitude  *float64 `json:"port_latitude,omitempty"`
	PortLongitude *float64 `json:"port_longitude,omitempty"`
	PortMapURL    *string  `json:"port_map_url,omitempty"`
	IsActive      int      `json:"is_active"`
	City          *City    `json:"city,omitempty"`
	Audit
}

type PortRequest struct {
	PortName      string   `json:"port_name" validate:"required,max=200"`
	CityID        *int     `json:"city_id,omitempty" validate:"omitempty,gt=0"`
	PortAddress   *string  `json:"port_address,omitempty"`
	PortLatitude  *float64 `json:"port_latitude,omitempty"`
	PortLongitude *float64 `json:"port_longitude,omitempty"`
	PortMapURL    *string  `json:"port_map_url,omitempty"`
	IsActive      *int     `json:"is_active,omitempty" validate:"omitempty,oneof=0 1"`
}
