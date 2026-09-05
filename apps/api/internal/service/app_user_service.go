package service

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/figoalfarqi/apipml/internal/auth"
	"github.com/figoalfarqi/apipml/internal/helper"
	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
	"github.com/figoalfarqi/apipml/internal/validation"
)

type AppUserService struct {
	Repo *repository.AppUserRepository
	File *FileUploadService
}

func NewAppUserService(r *repository.AppUserRepository, fus *FileUploadService) *AppUserService {
	return &AppUserService{Repo: r, File: fus}
}

// ======================================================================
// Utility
// ======================================================================
func trimAppUser(req interface{}) {
	switch v := req.(type) {
	case *model.AppUserCreateRequest:
		validation.TrimStrings(
			&v.Username,
			&v.AppUserName,
			v.AppUserPhone,
			v.AppUserPreferredName,
			v.AppUserAddress,
			v.IDCardNumber,
			v.FamilyCardNumber,
			v.DriverLicenseBNumber,
			v.TaxIDNumber,
			v.BpjsNumber,
			v.BankAccountNumber,
			v.BankAccountName,
		)
	case *model.AppUserUpdateRequest:
		validation.TrimStrings(
			v.Username,
			&v.AppUserName,
			v.AppUserPreferredName,
			v.AppUserPhone,
			v.AppUserAddress,
			v.IDCardNumber,
			v.FamilyCardNumber,
			v.DriverLicenseBNumber,
			v.TaxIDNumber,
			v.BpjsNumber,
			v.BankAccountNumber,
			v.BankAccountName,
		)
	}
}

// ======================================================================
// Create App User
// ======================================================================
func (s *AppUserService) Create(ctx context.Context, loginID int, req *model.AppUserCreateRequest, rUrlPath string) (*model.AppUserResponse, error, map[string]string) {
	trimAppUser(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}

	filters := make(map[string]string)
	if helper.IsGetDataAdmin(rUrlPath) {
		filters["app_role_type_id"] = "3"
	} else if helper.IsGetDataDriver(rUrlPath) {
		filters["app_role_type_id"] = "1"
	} else if helper.IsGetDataClientPic(rUrlPath) {
		filters["app_role_type_id"] = "2"
	}

	au, password, err := s.BuildAppUserFromCreateReq(ctx, req, loginID, filters)

	if err != nil {
		return nil, err, nil
	}

	id, err := s.Repo.Create(ctx, au)
	if err != nil {
		return nil, err, nil
	}

	// if req.AppUserPhotoURL != nil {
	// 	baseUrl := "http://localhost:5000"
	// 	folder, filename := s.File.getFolderAndFileName(*req.AppUserPhotoURL, baseUrl)
	// 	err := s.File.FinalizeFile(baseUrl, folder, filename)
	// 	if err != nil {
	// 		return nil, err, nil
	// 	}
	// 	url := baseUrl + "/uploads/" + folder + "/" + filename
	// 	req.AppUserPhotoURL = &url
	// }
	// if req.IDCardPhotoURL != nil {
	// 	baseUrl := "http://localhost:5000"
	// 	folder, filename := s.File.getFolderAndFileName(*req.IDCardPhotoURL, baseUrl)
	// 	err := s.File.FinalizeFile(baseUrl, folder, filename)
	// 	if err != nil {
	// 		return nil, err, nil
	// 	}
	// 	url := baseUrl + "/uploads/" + folder + "/" + filename
	// 	req.IDCardPhotoURL = &url
	// }

	files := map[string]*string{
		"app_user_photo":             req.AppUserPhotoURL,
		"id_card_photo":              req.IDCardPhotoURL,
		"family_card_photo_url":      req.FamilyCardPhotoURL,
		"driver_license_b_photo_url": req.DriverLicenseBPhotoURL,
		"tax_id_photo_url":           req.TaxIDPhotoURL,
		"bpjs_photo_url":             req.BpjsPhotoURL,
	}

	if err := s.File.FinalizeManyFiles(files); err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	res.Password = &password
	return res, err, nil
}

// ======================================================================
// Update App User
// ======================================================================
func (s *AppUserService) Update(ctx context.Context, loginID, id int, req *model.AppUserUpdateRequest) (*model.AppUserResponse, error, map[string]string) {
	trimAppUser(req)

	if err := validation.ValidateStruct(req); err != nil {
		return nil, errors.New("validation error"), validation.ValidationErrors(err)
	}
	if req.Password != nil {
		passwordHashed, err := auth.HashPassword(*req.Password)
		if err != nil {
			return nil, err, nil
		}
		req.Password = &passwordHashed
	}

	beforeRes, err := s.GetByID(ctx, id)

	if err != nil {
		return nil, err, nil
	}

	// if beforeRes.AppUserPhotoURL != nil && !helper.StrPtrEqual(req.AppUserPhotoURL, beforeRes.AppUserPhotoURL) {
	// 	baseUrl := "http://localhost:5000"
	// 	folder, filename := s.File.getFolderAndFileName(*beforeRes.AppUserPhotoURL, baseUrl)
	// 	err := s.File.DeleteFile(baseUrl, folder, filename)
	// 	fmt.Println("ERROR DELETE FILE", err)
	// }

	// if req.AppUserPhotoURL != nil && !helper.StrPtrEqual(req.AppUserPhotoURL, beforeRes.AppUserPhotoURL) {
	// 	baseUrl := "http://localhost:5000"
	// 	folder, filename := s.File.getFolderAndFileName(*req.AppUserPhotoURL, baseUrl)
	// 	err := s.File.FinalizeFile(baseUrl, folder, filename)
	// 	if err != nil {
	// 		return nil, err, nil
	// 	}
	// 	url := baseUrl + "/uploads/" + folder + "/" + filename
	// 	req.AppUserPhotoURL = &url
	// }

	// if beforeRes.IDCardPhotoURL != nil && !helper.StrPtrEqual(req.IDCardPhotoURL, beforeRes.IDCardPhotoURL) {
	// 	baseUrl := "http://localhost:5000"
	// 	folder, filename := s.File.getFolderAndFileName(*beforeRes.IDCardPhotoURL, baseUrl)
	// 	err := s.File.DeleteFile(baseUrl, folder, filename)
	// 	fmt.Println("ERROR DELETE FILE", err)
	// }

	// if req.IDCardPhotoURL != nil && !helper.StrPtrEqual(req.IDCardPhotoURL, beforeRes.IDCardPhotoURL) {
	// 	baseUrl := "http://localhost:5000"
	// 	folder, filename := s.File.getFolderAndFileName(*req.IDCardPhotoURL, baseUrl)
	// 	err := s.File.FinalizeFile(baseUrl, folder, filename)
	// 	if err != nil {
	// 		return nil, err, nil
	// 	}
	// 	url := baseUrl + "/uploads/" + folder + "/" + filename
	// 	req.IDCardPhotoURL = &url
	// }

	// update AppUserPhotoURL
	err = s.File.UpdateFile(beforeRes.AppUserPhotoURL, &req.AppUserPhotoURL)
	if err != nil {
		return nil, err, nil
	}

	// update IDCardPhotoURL
	err = s.File.UpdateFile(beforeRes.IDCardPhotoURL, &req.IDCardPhotoURL)
	if err != nil {
		return nil, err, nil
	}

	// update FamilyCardPhotoURL
	err = s.File.UpdateFile(beforeRes.FamilyCardPhotoURL, &req.FamilyCardPhotoURL)
	if err != nil {
		return nil, err, nil
	}

	// update DriverLicenseBPhotoURL
	err = s.File.UpdateFile(beforeRes.DriverLicenseBPhotoURL, &req.DriverLicenseBPhotoURL)
	if err != nil {
		return nil, err, nil
	}

	// update TaxIDPhotoURL
	err = s.File.UpdateFile(beforeRes.TaxIDPhotoURL, &req.TaxIDPhotoURL)
	if err != nil {
		return nil, err, nil
	}

	// update BpjsPhotoURL
	err = s.File.UpdateFile(beforeRes.BpjsPhotoURL, &req.BpjsPhotoURL)
	if err != nil {
		return nil, err, nil
	}

	uau := s.BuildAppUserFromUpdateReq(req, loginID)

	if err := s.Repo.Update(ctx, id, uau); err != nil {
		return nil, err, nil
	}

	res, err := s.GetByID(ctx, id)
	return res, err, nil
}

// ======================================================================
// Delete App User (Soft Delete)
// ======================================================================
func (s *AppUserService) Delete(ctx context.Context, loginID, id int) error {
	return s.Repo.SoftDelete(ctx, loginID, id)
}

// ======================================================================
// Get App User By ID
// ======================================================================
func (s *AppUserService) GetByID(ctx context.Context, id int) (*model.AppUserResponse, error) {
	u, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toAppUserResp(u), nil
}

// ======================================================================
// List App Users (Cursor + Filters)
// ======================================================================
func (s *AppUserService) List(
	ctx context.Context,
	cursorValue interface{}, cursorKey *int,
	limit int,
	filters map[string]string,
	orderBy, sort string,
) ([]model.AppUserResponse, error) {
	// Normalisasi time filter
	filters, err := helper.NormalizeTimeFilters(filters, nil)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.List(ctx, cursorValue, cursorKey, limit, filters, orderBy, sort)
	if err != nil {
		return nil, err
	}

	resp := make([]model.AppUserResponse, 0, len(items))
	for _, u := range items {
		resp = append(resp, *toAppUserResp(&u))
	}

	return resp, nil
}

// ======================================================================
// Mapping Entity → Response DTO
// ======================================================================
func toAppUserResp(u *model.AppUser) *model.AppUserResponse {
	if u == nil {
		return nil
	}

	var bankMerkResp *model.BankMerkResponse
	var cityResp *model.CityResponse
	var appRRoleResp *model.AppRoleResponse
	if u.BankMerk != nil {
		bankMerkResp = toBankMerkResp(u.BankMerk)
	}
	if u.City != nil {
		cityResp = toCityResp(u.City)
	}
	if u.AppRole != nil {
		appRRoleResp = toAppRoleResp(u.AppRole)
	}

	return &model.AppUserResponse{
		AppUserID:              u.AppUserID,
		AppRoleID:              u.AppRoleID,
		ClientID:               u.ClientID,
		BankMerkID:             u.BankMerkID,
		Username:               u.Username,
		AppUserStatusID:        u.AppUserStatusID,
		AppUserName:            u.AppUserName,
		AppUserPreferredName:   u.AppUserPreferredName,
		AppUserPhone:           u.AppUserPhone,
		CityID:                 u.CityID,
		AppUserAddress:         u.AppUserAddress,
		AppUserPhotoURL:        u.AppUserPhotoURL,
		IDCardPhotoURL:         u.IDCardPhotoURL,
		IDCardNumber:           u.IDCardNumber,
		FamilyCardPhotoURL:     u.FamilyCardPhotoURL,
		FamilyCardNumber:       u.FamilyCardNumber,
		DriverLicenseBPhotoURL: u.DriverLicenseBPhotoURL,
		DriverLicenseBNumber:   u.DriverLicenseBNumber,
		DriverLicenseBExpiry:   u.DriverLicenseBExpiry,
		TaxIDPhotoURL:          u.TaxIDPhotoURL,
		TaxIDNumber:            u.TaxIDNumber,
		BpjsPhotoURL:           u.BpjsPhotoURL,
		BpjsNumber:             u.BpjsNumber,
		BankAccountNumber:      u.BankAccountNumber,
		BankAccountName:        u.BankAccountName,
		SalaryPercentage:       u.SalaryPercentage,
		CreatedBy:              u.CreatedBy,
		UpdatedBy:              u.UpdatedBy,
		CreatedAt:              u.CreatedAt,
		UpdatedAt:              u.UpdatedAt,
		DeletedAt:              u.DeletedAt,
		City:                   cityResp,
		AppRole:                appRRoleResp,
		BankMerk:               bankMerkResp,
	}
}

func generateUsername(appUserName string) (string, error) {

	// lowercase & hapus spasi
	base := strings.ToLower(appUserName)
	base = strings.ReplaceAll(base, " ", "")
	digitstring, err := helper.GenerateNDigitString(6)
	if err != nil {
		return "GenerateNDigitStringFailed", err
	}
	return fmt.Sprintf("%s%s", base, digitstring), nil
}

func (s *AppUserService) BuildAppUserFromCreateReq(ctx context.Context, req *model.AppUserCreateRequest, loginID int, filters map[string]string) (*model.AppUser, string, error) {
	username := req.Username
	if strings.TrimSpace(username) == "" {
		usernameGenerated, err := generateUsername(req.AppUserName)
		if err != nil {
			return nil, "", err
		}
		username = usernameGenerated
	}

	password := req.Password
	if strings.TrimSpace(password) == "dutakasih123456" {
		digitstring, err := helper.GenerateNDigitString(6)
		if err != nil {
			return nil, "", err
		}
		password = digitstring
	}
	passwordHashed, err := auth.HashPassword(password)
	if err != nil {
		return nil, "", err
	}

	return &model.AppUser{
		AppRoleID:              req.AppRoleID,
		ClientID:               req.ClientID,
		BankMerkID:             req.BankMerkID,
		Username:               username,
		Password:               passwordHashed,
		AppUserStatusID:        *req.AppUserStatusID,
		AppUserName:            req.AppUserName,
		AppUserPreferredName:   req.AppUserPreferredName,
		AppUserPhone:           req.AppUserPhone,
		CityID:                 req.CityID,
		AppUserAddress:         req.AppUserAddress,
		AppUserPhotoURL:        req.AppUserPhotoURL,
		IDCardPhotoURL:         req.IDCardPhotoURL,
		IDCardNumber:           req.IDCardNumber,
		FamilyCardPhotoURL:     req.FamilyCardPhotoURL,
		FamilyCardNumber:       req.FamilyCardNumber,
		DriverLicenseBPhotoURL: req.DriverLicenseBPhotoURL,
		DriverLicenseBNumber:   req.DriverLicenseBNumber,
		DriverLicenseBExpiry:   req.DriverLicenseBExpiry,
		TaxIDPhotoURL:          req.TaxIDPhotoURL,
		TaxIDNumber:            req.TaxIDNumber,
		BpjsPhotoURL:           req.BpjsPhotoURL,
		BpjsNumber:             req.BpjsNumber,
		BankAccountNumber:      req.BankAccountNumber,
		BankAccountName:        req.BankAccountName,
		SalaryPercentage:       req.SalaryPercentage,
		CreatedBy:              loginID,
		UpdatedBy:              loginID,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}, password, nil
}

func (s *AppUserService) BuildAppUserFromUpdateReq(req *model.AppUserUpdateRequest, loginID int) *model.AppUser {
	return &model.AppUser{
		AppRoleID:              pointerValueOr(req.AppRoleID, -1),
		ClientID:               req.ClientID,
		BankMerkID:             req.BankMerkID,
		Username:               stringPointerValueOr(req.Username, ""),
		Password:               stringPointerValueOr(req.Password, ""),
		AppUserName:            req.AppUserName,
		AppUserPreferredName:   req.AppUserPreferredName,
		AppUserPhone:           req.AppUserPhone,
		CityID:                 req.CityID,
		AppUserAddress:         req.AppUserAddress,
		AppUserPhotoURL:        req.AppUserPhotoURL,
		AppUserStatusID:        pointerValueOr(req.AppUserStatusID, -1),
		IDCardPhotoURL:         req.IDCardPhotoURL,
		IDCardNumber:           req.IDCardNumber,
		FamilyCardPhotoURL:     req.FamilyCardPhotoURL,
		FamilyCardNumber:       req.FamilyCardNumber,
		DriverLicenseBPhotoURL: req.DriverLicenseBPhotoURL,
		DriverLicenseBNumber:   req.DriverLicenseBNumber,
		DriverLicenseBExpiry:   req.DriverLicenseBExpiry,
		TaxIDPhotoURL:          req.TaxIDPhotoURL,
		TaxIDNumber:            req.TaxIDNumber,
		BpjsPhotoURL:           req.BpjsPhotoURL,
		BpjsNumber:             req.BpjsNumber,
		BankAccountNumber:      req.BankAccountNumber,
		BankAccountName:        req.BankAccountName,
		SalaryPercentage:       req.SalaryPercentage,
		UpdatedBy:              loginID,
		UpdatedAt:              time.Now(),
	}
}

func pointerValueOr(value *int, fallback int) int {
	if value == nil {
		return fallback
	}
	return *value
}

func stringPointerValueOr(value *string, fallback string) string {
	if value == nil {
		return fallback
	}
	return *value
}

func (s *AppUserService) BuildAppUserFromUpdateReqClient(req *model.AppUserUpdateRequestForClient, loginID int) *model.AppUser {
	return &model.AppUser{
		AppRoleID:              -1,
		ClientID:               req.ClientID,
		BankMerkID:             req.BankMerkID,
		AppUserName:            req.AppUserName,
		AppUserPreferredName:   req.AppUserPreferredName,
		AppUserPhone:           req.AppUserPhone,
		CityID:                 req.CityID,
		AppUserAddress:         req.AppUserAddress,
		AppUserPhotoURL:        req.AppUserPhotoURL,
		AppUserStatusID:        -1,
		IDCardPhotoURL:         req.IDCardPhotoURL,
		IDCardNumber:           req.IDCardNumber,
		FamilyCardPhotoURL:     req.FamilyCardPhotoURL,
		FamilyCardNumber:       req.FamilyCardNumber,
		DriverLicenseBPhotoURL: req.DriverLicenseBPhotoURL,
		DriverLicenseBNumber:   req.DriverLicenseBNumber,
		DriverLicenseBExpiry:   req.DriverLicenseBExpiry,
		TaxIDPhotoURL:          req.TaxIDPhotoURL,
		TaxIDNumber:            req.TaxIDNumber,
		BpjsPhotoURL:           req.BpjsPhotoURL,
		BpjsNumber:             req.BpjsNumber,
		BankAccountNumber:      req.BankAccountNumber,
		BankAccountName:        req.BankAccountName,
		SalaryPercentage:       req.SalaryPercentage,
		UpdatedBy:              loginID,
		UpdatedAt:              time.Now(),
	}
}

// ======================================================================
// Login & Password Management
// ======================================================================

func (s *AppUserService) Login(ctx context.Context, username, password string, appRoleIDs []int) (*model.AppUser, error) {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	u, err := s.Repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if u.AppUserStatusID != 1 {
		return nil, errors.New("invalid credentials")
	}
	canonicalRoleID, ok := helper.CanonicalAppRoleID(u.LoginRoleName)
	if !ok || !slices.Contains(appRoleIDs, canonicalRoleID) {
		return nil, errors.New("invalid role")
	}

	if err := auth.CheckPassword(u.Password, password); err != nil {
		return nil, errors.New("invalid credentials")
	}
	// The database role ID may differ between legacy and current schemas.
	// Tokens consistently use the canonical role ID.
	u.AppRoleID = canonicalRoleID
	return u, nil
}

func (s *AppUserService) ChangePassword(ctx context.Context, id int, req *model.AppUserChangePassword) error {
	password, err := s.Repo.GetPasswordByID(ctx, id)
	if err != nil {
		return err
	}
	if err := auth.CheckPassword(*password, req.OldPassword); err != nil {
		return errors.New("old password is incorrect")
	}
	if req.NewPassword != req.RetypeNewPassword {
		return errors.New("new password and retype do not match")
	}
	hashed, err := auth.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	return s.Repo.UpdatePassword(ctx, id, hashed)
}

func (s *AppUserService) ResetPassword(ctx context.Context, req *model.AppUserResetPassword) error {
	hashed, err := auth.HashPassword("dutakasih")
	if err != nil {
		return err
	}
	return s.Repo.UpdatePassword(ctx, req.AppUserID, hashed)
}

func (s *AppUserService) GetAuthVersion(ctx context.Context, id int) (int, error) {
	return s.Repo.GetAuthVersion(ctx, id)
}

func (s *AppUserService) InvalidateAuth(ctx context.Context, id int) error {
	return s.Repo.InvalidateAuth(ctx, id)
}
