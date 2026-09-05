package model

import "time"

// ===== Entity =====
type AppUser struct {
	AppUserID              int        `json:"app_user_id"`
	AppRoleID              int        `json:"app_role_id"`
	ClientID               *int       `json:"client_id,omitempty"`
	BankMerkID             *int       `json:"bank_merk_id,omitempty"`
	Username               string     `json:"username"`
	Password               string     `json:"-"` // hashed, tidak di-expose
	AppUserStatusID        int        `json:"app_user_status_id"`
	AppUserName            string     `json:"app_user_name"`
	AppUserPreferredName   *string    `json:"app_user_preferred_name,omitempty"`
	AppUserPhone           *string    `json:"app_user_phone,omitempty"`
	CityID                 *int       `json:"city_id,omitempty"`
	AppUserAddress         *string    `json:"app_user_address,omitempty"`
	AppUserPhotoURL        *string    `json:"app_user_photo_url,omitempty"`
	IDCardPhotoURL         *string    `json:"id_card_photo_url,omitempty"`
	IDCardNumber           *string    `json:"id_card_number,omitempty"`
	FamilyCardPhotoURL     *string    `json:"family_card_photo_url,omitempty"`
	FamilyCardNumber       *string    `json:"family_card_number,omitempty"`
	DriverLicenseBPhotoURL *string    `json:"driver_license_b_photo_url,omitempty"`
	DriverLicenseBNumber   *string    `json:"driver_license_b_number,omitempty"`
	DriverLicenseBExpiry   *time.Time `json:"driver_license_b_expiry,omitempty"`
	TaxIDPhotoURL          *string    `json:"tax_id_photo_url,omitempty"`
	TaxIDNumber            *string    `json:"tax_id_number,omitempty"`
	BpjsPhotoURL           *string    `json:"bpjs_photo_url,omitempty"`
	BpjsNumber             *string    `json:"bpjs_number,omitempty"`
	BankAccountNumber      *string    `json:"bank_account_number,omitempty"`
	BankAccountName        *string    `json:"bank_account_name,omitempty"`
	SalaryPercentage       *float64   `json:"salary_percentage,omitempty"`
	CreatedBy              int        `json:"created_by"`
	UpdatedBy              int        `json:"updated_by"`
	DeletedBy              *int       `json:"deleted_by,omitempty"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
	DeletedAt              *time.Time `json:"deleted_at,omitempty"`
	LoginRoleName          string     `json:"-"`

	City     *City     `json:"city,omitempty"`
	AppRole  *AppRole  `json:"app_role,omitempty"`
	BankMerk *BankMerk `json:"bank_merk,omitempty"`
}

// ===== Create Request =====
type AppUserCreateRequest struct {
	AppRoleID              int        `json:"app_role_id" validate:"gt=0"`
	ClientID               *int       `json:"client_id" validate:"omitempty,gt=0"`
	BankMerkID             *int       `json:"bank_merk_id" validate:"omitempty,gt=0"`
	Username               string     `json:"username" validate:"required,min=4,max=200"`
	Password               string     `json:"password" validate:"required,min=6"`
	AppUserStatusID        *int       `json:"app_user_status_id" validate:"required,oneof=0 1 2"`
	AppUserName            string     `json:"app_user_name" validate:"required,min=3,max=100"`
	AppUserPreferredName   *string    `json:"app_user_preferred_name,omitempty"`
	AppUserPhone           *string    `json:"app_user_phone" validate:"omitempty,required,max=20"`
	CityID                 *int       `json:"city_id" validate:"omitempty,gt=0"`
	AppUserAddress         *string    `json:"app_user_address" validate:"omitempty"`
	AppUserPhotoURL        *string    `json:"app_user_photo_url" validate:"omitempty"`
	IDCardPhotoURL         *string    `json:"id_card_photo_url,omitempty"`
	IDCardNumber           *string    `json:"id_card_number,omitempty"`
	FamilyCardPhotoURL     *string    `json:"family_card_photo_url,omitempty"`
	FamilyCardNumber       *string    `json:"family_card_number,omitempty"`
	DriverLicenseBPhotoURL *string    `json:"driver_license_b_photo_url,omitempty"`
	DriverLicenseBNumber   *string    `json:"driver_license_b_number,omitempty"`
	DriverLicenseBExpiry   *time.Time `json:"driver_license_b_expiry,omitempty"`
	TaxIDPhotoURL          *string    `json:"tax_id_photo_url,omitempty"`
	TaxIDNumber            *string    `json:"tax_id_number,omitempty"`
	BpjsPhotoURL           *string    `json:"bpjs_photo_url,omitempty"`
	BpjsNumber             *string    `json:"bpjs_number,omitempty"`
	BankAccountNumber      *string    `json:"bank_account_number,omitempty"`
	BankAccountName        *string    `json:"bank_account_name,omitempty"`
	SalaryPercentage       *float64   `json:"salary_percentage,omitempty"`
}

// ===== Update Request =====
type AppUserUpdateRequest struct {
	AppRoleID              *int       `json:"app_role_id" validate:"omitempty,gt=0"`
	ClientID               *int       `json:"client_id" validate:"omitempty,gt=0"`
	BankMerkID             *int       `json:"bank_merk_id" validate:"omitempty,gt=0"`
	Username               *string    `json:"username,omitempty" validate:"omitempty,min=4,max=200"`
	Password               *string    `json:"password,omitempty" validate:"omitempty,min=6"`
	AppUserName            string     `json:"app_user_name" validate:"min=3,max=100"`
	AppUserPreferredName   *string    `json:"app_user_preferred_name,omitempty"`
	AppUserPhone           *string    `json:"app_user_phone" validate:"omitempty,max=20"`
	CityID                 *int       `json:"city_id" validate:"omitempty,gt=0"`
	AppUserAddress         *string    `json:"app_user_address,omitempty"`
	AppUserPhotoURL        *string    `json:"app_user_photo_url,omitempty"`
	AppUserStatusID        *int       `json:"app_user_status_id,omitempty" validate:"omitempty,oneof=0 1 2"`
	IDCardPhotoURL         *string    `json:"id_card_photo_url,omitempty"`
	IDCardNumber           *string    `json:"id_card_number,omitempty"`
	FamilyCardPhotoURL     *string    `json:"family_card_photo_url,omitempty"`
	FamilyCardNumber       *string    `json:"family_card_number,omitempty"`
	DriverLicenseBPhotoURL *string    `json:"driver_license_b_photo_url,omitempty"`
	DriverLicenseBNumber   *string    `json:"driver_license_b_number,omitempty"`
	DriverLicenseBExpiry   *time.Time `json:"driver_license_b_expiry,omitempty"`
	TaxIDPhotoURL          *string    `json:"tax_id_photo_url,omitempty"`
	TaxIDNumber            *string    `json:"tax_id_number,omitempty"`
	BpjsPhotoURL           *string    `json:"bpjs_photo_url,omitempty"`
	BpjsNumber             *string    `json:"bpjs_number,omitempty"`
	BankAccountNumber      *string    `json:"bank_account_number,omitempty"`
	BankAccountName        *string    `json:"bank_account_name,omitempty"`
	SalaryPercentage       *float64   `json:"salary_percentage,omitempty"`
}

type AppUserUpdateRequestForClient struct {
	AppUserID int `json:"app_user_id" validate:"required,gt=0"`
	// AppRoleID              *int       `json:"app_role_id" validate:"omitempty,gt=0"`
	ClientID             *int    `json:"client_id" validate:"omitempty,gt=0"`
	BankMerkID           *int    `json:"bank_merk_id" validate:"omitempty,gt=0"`
	AppUserName          string  `json:"app_user_name" validate:"min=3,max=100"`
	AppUserPreferredName *string `json:"app_user_preferred_name,omitempty"`
	AppUserPhone         *string `json:"app_user_phone" validate:"omitempty,max=20"`
	CityID               *int    `json:"city_id" validate:"omitempty,gt=0"`
	AppUserAddress       *string `json:"app_user_address,omitempty"`
	AppUserPhotoURL      *string `json:"app_user_photo_url,omitempty"`
	// AppUserStatusID        *int       `json:"app_user_status_id,omitempty"`
	IDCardPhotoURL         *string    `json:"id_card_photo_url,omitempty"`
	IDCardNumber           *string    `json:"id_card_number,omitempty"`
	FamilyCardPhotoURL     *string    `json:"family_card_photo_url,omitempty"`
	FamilyCardNumber       *string    `json:"family_card_number,omitempty"`
	DriverLicenseBPhotoURL *string    `json:"driver_license_b_photo_url,omitempty"`
	DriverLicenseBNumber   *string    `json:"driver_license_b_number,omitempty"`
	DriverLicenseBExpiry   *time.Time `json:"driver_license_b_expiry,omitempty"`
	TaxIDPhotoURL          *string    `json:"tax_id_photo_url,omitempty"`
	TaxIDNumber            *string    `json:"tax_id_number,omitempty"`
	BpjsPhotoURL           *string    `json:"bpjs_photo_url,omitempty"`
	BpjsNumber             *string    `json:"bpjs_number,omitempty"`
	BankAccountNumber      *string    `json:"bank_account_number,omitempty"`
	BankAccountName        *string    `json:"bank_account_name,omitempty"`
	SalaryPercentage       *float64   `json:"salary_percentage,omitempty"`
}

// ===== Response =====
type AppUserResponse struct {
	AppUserID              int        `json:"app_user_id"`
	AppRoleID              int        `json:"app_role_id,omitempty"`
	ClientID               *int       `json:"client_id,omitempty"`
	BankMerkID             *int       `json:"bank_merk_id,omitempty"`
	Username               string     `json:"username"`
	Password               *string    `json:"password,omitempty"`
	AppUserStatusID        int        `json:"app_user_status_id"`
	AppUserName            string     `json:"app_user_name"`
	AppUserPreferredName   *string    `json:"app_user_preferred_name,omitempty"`
	AppUserPhone           *string    `json:"app_user_phone,omitempty"`
	CityID                 *int       `json:"city_id,omitempty"`
	AppUserAddress         *string    `json:"app_user_address,omitempty"`
	AppUserPhotoURL        *string    `json:"app_user_photo_url,omitempty"`
	IDCardPhotoURL         *string    `json:"id_card_photo_url,omitempty"`
	IDCardNumber           *string    `json:"id_card_number,omitempty"`
	FamilyCardPhotoURL     *string    `json:"family_card_photo_url,omitempty"`
	FamilyCardNumber       *string    `json:"family_card_number,omitempty"`
	DriverLicenseBPhotoURL *string    `json:"driver_license_b_photo_url,omitempty"`
	DriverLicenseBNumber   *string    `json:"driver_license_b_number,omitempty"`
	DriverLicenseBExpiry   *time.Time `json:"driver_license_b_expiry,omitempty"`
	TaxIDPhotoURL          *string    `json:"tax_id_photo_url,omitempty"`
	TaxIDNumber            *string    `json:"tax_id_number,omitempty"`
	BpjsPhotoURL           *string    `json:"bpjs_photo_url,omitempty"`
	BpjsNumber             *string    `json:"bpjs_number,omitempty"`
	BankAccountNumber      *string    `json:"bank_account_number,omitempty"`
	BankAccountName        *string    `json:"bank_account_name,omitempty"`
	SalaryPercentage       *float64   `json:"salary_percentage,omitempty"`
	CreatedBy              int        `json:"created_by"`
	UpdatedBy              int        `json:"updated_by"`
	DeletedBy              *int       `json:"deleted_by,omitempty"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
	DeletedAt              *time.Time `json:"deleted_at,omitempty"`

	City     *CityResponse     `json:"city,omitempty"`
	AppRole  *AppRoleResponse  `json:"app_role,omitempty"`
	BankMerk *BankMerkResponse `json:"bank_merk,omitempty"`
}

// ===== Password Change =====
type AppUserChangePassword struct {
	OldPassword       string `json:"old_password" validate:"required"`
	NewPassword       string `json:"new_password" validate:"required"`
	RetypeNewPassword string `json:"retype_new_password" validate:"required"`
}

type AppUserResetPassword struct {
	AppUserID int `json:"app_user_id"`
}

// ===== Auth =====
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
