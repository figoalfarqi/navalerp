package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/app/middleware"
	"github.com/figoalfarqi/navalerp/internal/auth"
	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/service"
	"github.com/figoalfarqi/navalerp/pkg/response"
	"github.com/jackc/pgx/v5"
)

type AppUserHandler struct {
	Svc *service.AppUserService
	Cfg *config.Config
}

func NewAppUserHandler(s *service.AppUserService, c *config.Config) *AppUserHandler {
	return &AppUserHandler{Svc: s, Cfg: c}
}

// ==================================================
// POST /api/v1/admin/app_user
// ==================================================
func (h *AppUserHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	var req model.AppUserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}
	callerRoleID, ok := requestRoleID(r)
	if !ok || !prepareUserRoleForCreate(r.URL.Path, callerRoleID, &req) {
		response.JSON(w, http.StatusForbidden, "not allowed to create this user role", nil, nil)
		return
	}

	user, err, formErr := h.Svc.Create(r.Context(), loginID, &req, r.URL.Path)
	if formErr != nil {
		response.JSON(w, http.StatusBadRequest, "validation error", nil, formErr)
		return
	}
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to create user", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusCreated, "user created", user, nil)
}

// ==================================================
// PUT /api/v1/admin/admin/{id}
// PUT /api/v1/admin/driver/{id}
// ==================================================
func (h *AppUserHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)

	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	id := requestID(r)
	if id <= 0 {
		response.JSON(w, http.StatusBadRequest, "invalid id", nil, map[string]string{"id": "invalid"})
		return
	}

	callerRoleID, ok := requestRoleID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	before, err := h.Svc.GetByID(r.Context(), id)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	if !canManageUserTarget(r.URL.Path, callerRoleID, loginID, before.AppUserID, before.AppRoleID) {
		response.JSON(w, http.StatusForbidden, "not allowed to update this user", nil, nil)
		return
	}

	var req model.AppUserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}
	if !prepareUserRoleForUpdate(r.URL.Path, callerRoleID, before.AppRoleID, &req) {
		response.JSON(w, http.StatusForbidden, "not allowed to assign this user role", nil, nil)
		return
	}

	user, err, formErr := h.Svc.Update(r.Context(), loginID, id, &req)
	if formErr != nil {
		response.JSON(w, http.StatusBadRequest, "validation error", nil, formErr)
		return
	}
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to update user", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, "user updated", user, nil)
}

// ==================================================
// DELETE /api/v1/admin/app_user/{id}
// ==================================================
func (h *AppUserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	id := requestID(r)
	if id <= 0 {
		response.JSON(w, http.StatusBadRequest, "invalid id", nil, map[string]string{"id": "invalid"})
		return
	}

	callerRoleID, ok := requestRoleID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	target, err := h.Svc.GetByID(r.Context(), id)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	if !canManageUserTarget(r.URL.Path, callerRoleID, loginID, target.AppUserID, target.AppRoleID) {
		response.JSON(w, http.StatusForbidden, "not allowed to delete this user", nil, nil)
		return
	}

	if err := h.Svc.Delete(r.Context(), loginID, id); err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to delete user", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, "user deleted", map[string]int{
		"app_user_id": id,
	}, nil)
}

// ==================================================
// GET /api/v1/admin/admin OR /{id}
// GET /api/v1/admin/driver OR /{id}
// ==================================================
func (h *AppUserHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}

	callerRoleID, ok := requestRoleID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	id := requestID(r)

	// GET by ID
	if id > 0 {
		user, err := h.Svc.GetByID(r.Context(), id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				response.JSON(w, http.StatusNotFound, "user not found", nil, nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, "failed to fetch user", nil, map[string]string{
				"error": err.Error(),
			})
			return
		}
		if !canReadUserTarget(r.URL.Path, callerRoleID, loginID, user.AppUserID, user.AppRoleID) {
			response.JSON(w, http.StatusForbidden, "not allowed to view this user", nil, nil)
			return
		}

		response.JSON(w, http.StatusOK, "ok", user, nil)
		return
	}

	// GET list
	queryParams := queryValues(r)
	limit := helper.AtoiSafeDefault(queryParams.Get("limit"), 10)
	if limit > 1000 {
		limit = 1000
	}
	cursorData, err := helper.ParseCursorParams(queryParams)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid cursor", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}
	baseAllowed := []string{
		"username", "app_user_name",
		"app_user_preferred_name",
		"app_user_phone",
		"app_user_address",
		"id_card_number",
		"driver_license_b_number",
		"driver_license_b_expiry",
		"tax_id_number",
		"bpjs_number",
		"bank_account_number",
		"bank_account_name",
		"salary_percentage",
		"app_role_id",
		"app_role_type_id",
		"bank_merk_id",
		"city_id",
		"app_user_status_id",
		"client_id",
		"created_by", "updated_by",
	}

	filterOnly := []string{
		"created_at_after",
		"created_at_before",
		"updated_at_after",
		"updated_at_before",
		"usernameEXACT",
		"app_user_idNOT",
	}

	sortOnly := []string{
		"created_at",
		"updated_at",
	}

	allowedFilters := append(baseAllowed, filterOnly...)
	allowedSorts := append(baseAllowed, sortOnly...)

	filters := helper.ParseQueryFilter(queryParams, allowedFilters)
	orderBy, sort := helper.ParseQuerySort(queryParams, allowedSorts)

	if helper.IsGetDataAdmin(r.URL.Path) {
		filters["app_role_type_id"] = "3"
		if callerRoleID == helper.AppRoleAdmin {
			filters["app_role_id"] = "6"
		}
	} else if helper.IsGetDataDriver(r.URL.Path) {
		filters["app_role_type_id"] = "1"
	} else if helper.IsGetDataChecker(r.URL.Path) {
		filters["app_role_type_id"] = "2"
	}

	var cursorValue interface{} = nil
	var cursorKey *int = nil
	if cursorData != nil {
		cursorValue = cursorData.Value
		cursorKey = cursorData.Key
	}

	items, err := h.Svc.List(r.Context(), cursorValue, cursorKey, limit, filters, orderBy, sort)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to list users", nil, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, "ok", map[string]interface{}{
		"items":        items,
		"limit":        limit,
		"cursor_value": cursorValue,
		"cursor_key":   cursorKey,
		"filters":      filters,
		"order_by":     orderBy,
		"sort":         sort,
	}, nil)
}

func userResource(path string) string {
	switch {
	case strings.Contains(path, "/admin/driver"):
		return "driver"
	case strings.Contains(path, "/admin/checker"):
		return "checker"
	case strings.Contains(path, "/admin/admin"):
		return "admin"
	case strings.Contains(path, "/admin/app_user"):
		return "app_user"
	case strings.Contains(path, "/driver/driver"):
		return "driver_self"
	case strings.Contains(path, "/checker/checker"):
		return "checker_self"
	default:
		return ""
	}
}

func prepareUserRoleForCreate(path string, callerRoleID int, req *model.AppUserCreateRequest) bool {
	resource := userResource(path)
	if callerRoleID == helper.AppRoleOwner ||
		callerRoleID == helper.AppRoleITDev {
		switch resource {
		case "driver":
			req.AppRoleID = helper.AppRoleDriver
			return true
		case "checker":
			req.AppRoleID = helper.AppRoleChecker
			return true
		case "admin":
			return req.AppRoleID >= helper.AppRoleOwner &&
				req.AppRoleID <= helper.AppRoleAdmin
		case "app_user":
			_, knownRole := helper.AppRoleName(req.AppRoleID)
			return knownRole
		}
	}

	switch resource {
	case "driver":
		req.AppRoleID = helper.AppRoleDriver
		return callerRoleID == helper.AppRoleSuperAdmin || callerRoleID == helper.AppRoleAdmin
	case "checker":
		req.AppRoleID = helper.AppRoleChecker
		return callerRoleID == helper.AppRoleSuperAdmin || callerRoleID == helper.AppRoleAdmin
	case "admin":
		if callerRoleID == helper.AppRoleAdmin {
			req.AppRoleID = helper.AppRoleAdmin
			return true
		}
		if callerRoleID == helper.AppRoleSuperAdmin {
			return req.AppRoleID >= helper.AppRoleOwner && req.AppRoleID <= helper.AppRoleAdmin
		}
		return callerRoleID == helper.AppRoleITDev &&
			req.AppRoleID >= helper.AppRoleOwner &&
			req.AppRoleID <= helper.AppRoleAdmin &&
			req.AppRoleID != helper.AppRoleSuperAdmin
	case "app_user":
		if callerRoleID == helper.AppRoleSuperAdmin {
			return req.AppRoleID >= helper.AppRoleDriver && req.AppRoleID <= helper.AppRoleAdmin
		}
		return callerRoleID == helper.AppRoleITDev &&
			req.AppRoleID >= helper.AppRoleDriver && req.AppRoleID <= helper.AppRoleAdmin &&
			req.AppRoleID != helper.AppRoleSuperAdmin
	default:
		return false
	}
}

func prepareUserRoleForUpdate(path string, callerRoleID, currentRoleID int, req *model.AppUserUpdateRequest) bool {
	resource := userResource(path)
	if callerRoleID == helper.AppRoleOwner ||
		callerRoleID == helper.AppRoleITDev {
		if req.AppRoleID == nil {
			return resource == "driver" ||
				resource == "checker" ||
				resource == "admin" ||
				resource == "app_user"
		}

		targetRoleID := *req.AppRoleID
		switch resource {
		case "driver":
			return targetRoleID == helper.AppRoleDriver
		case "checker":
			return targetRoleID == helper.AppRoleChecker
		case "admin":
			return targetRoleID >= helper.AppRoleOwner &&
				targetRoleID <= helper.AppRoleAdmin
		case "app_user":
			_, knownRole := helper.AppRoleName(targetRoleID)
			return knownRole
		default:
			return false
		}
	}

	expectedRoleID := currentRoleID
	switch resource {
	case "driver", "driver_self":
		expectedRoleID = helper.AppRoleDriver
	case "checker", "checker_self":
		expectedRoleID = helper.AppRoleChecker
	case "admin":
		if callerRoleID == helper.AppRoleAdmin {
			expectedRoleID = helper.AppRoleAdmin
		}
		if callerRoleID == helper.AppRoleITDev && currentRoleID == helper.AppRoleSuperAdmin {
			return false
		}
	case "app_user":
		if callerRoleID != helper.AppRoleSuperAdmin && callerRoleID != helper.AppRoleITDev {
			return false
		}
	}
	if req.AppRoleID == nil {
		return currentRoleID == expectedRoleID ||
			(userResource(path) == "admin" && callerRoleID != helper.AppRoleAdmin)
	}
	targetRoleID := *req.AppRoleID
	if userResource(path) == "admin" {
		if callerRoleID != helper.AppRoleITDev && callerRoleID != helper.AppRoleSuperAdmin && callerRoleID != helper.AppRoleAdmin {
			return false
		}
		if callerRoleID == helper.AppRoleAdmin {
			return targetRoleID == helper.AppRoleAdmin
		}
		if callerRoleID == helper.AppRoleITDev {
			return targetRoleID >= helper.AppRoleOwner &&
				targetRoleID <= helper.AppRoleAdmin &&
				targetRoleID != helper.AppRoleSuperAdmin
		}
		return targetRoleID >= helper.AppRoleOwner && targetRoleID <= helper.AppRoleAdmin
	}
	if userResource(path) == "app_user" && callerRoleID == helper.AppRoleITDev {
		return targetRoleID != helper.AppRoleSuperAdmin
	}
	return targetRoleID == expectedRoleID
}

func canManageUserTarget(path string, callerRoleID, callerUserID, targetUserID, targetRoleID int) bool {
	if (callerRoleID == helper.AppRoleOwner ||
		callerRoleID == helper.AppRoleITDev) &&
		strings.HasPrefix(path, "/api/v1/admin/") {
		return true
	}

	switch userResource(path) {
	case "driver_self":
		return callerRoleID == helper.AppRoleDriver && callerUserID == targetUserID && targetRoleID == helper.AppRoleDriver
	case "checker_self":
		return callerRoleID == helper.AppRoleChecker && callerUserID == targetUserID && targetRoleID == helper.AppRoleChecker
	case "driver":
		return targetRoleID == helper.AppRoleDriver &&
			(callerRoleID == helper.AppRoleSuperAdmin || callerRoleID == helper.AppRoleAdmin)
	case "checker":
		return targetRoleID == helper.AppRoleChecker &&
			(callerRoleID == helper.AppRoleSuperAdmin || callerRoleID == helper.AppRoleAdmin)
	case "admin":
		if callerRoleID == helper.AppRoleAdmin {
			return targetRoleID == helper.AppRoleAdmin
		}
		return callerRoleID == helper.AppRoleSuperAdmin ||
			(callerRoleID == helper.AppRoleITDev && targetRoleID != helper.AppRoleSuperAdmin)
	case "app_user":
		return callerRoleID == helper.AppRoleSuperAdmin ||
			(callerRoleID == helper.AppRoleITDev && targetRoleID != helper.AppRoleSuperAdmin)
	default:
		return false
	}
}

func canReadUserTarget(path string, callerRoleID, callerUserID, targetUserID, targetRoleID int) bool {
	if userResource(path) == "driver_self" || userResource(path) == "checker_self" {
		return callerUserID == targetUserID && callerRoleID == targetRoleID
	}
	if strings.HasPrefix(path, "/api/v1/admin/") &&
		(callerRoleID == helper.AppRoleOwner || callerRoleID == helper.AppRoleITDev) {
		return true
	}
	return canManageUserTarget(path, callerRoleID, callerUserID, targetUserID, targetRoleID)
}

// GET /api/v1/auth/driver/login
// GET /api/v1/auth/admin/login
// GET /api/v1/auth/checker/login
func (h *AppUserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	role, _ := helper.GetIDAndRoleFromPath(r.URL.Path, "login")

	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}
	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Password) == "" {
		response.JSON(w, http.StatusBadRequest, "validation error", nil, map[string]string{"username": "required", "password": "required"})
		return
	}

	appRoleIDs := helper.GetAppRoleIDsByRoleTypeName(role)
	user, err := h.Svc.Login(r.Context(), req.Username, req.Password, appRoleIDs)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, "invalid username or password", nil, map[string]string{"credentials": "invalid"})
		return
	}

	authVersion, err := h.Svc.GetAuthVersion(r.Context(), user.AppUserID)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to create session", nil, nil)
		return
	}
	tok, err := auth.GenerateToken(
		h.Cfg.JWTKey,
		user.AppUserID,
		user.Username,
		user.AppRoleID,
		authVersion,
		24*time.Hour,
	)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to generate token", nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "login success", model.LoginResponse{Token: tok}, nil)
}

// POST /api/v1/auth/logout
func (h *AppUserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		response.JSON(w, http.StatusBadRequest, "authorization header missing", nil, nil)
		return
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		response.JSON(w, http.StatusBadRequest, "invalid authorization header format", nil, nil)
		return
	}

	loginID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	if err := h.Svc.InvalidateAuth(r.Context(), loginID); err != nil {
		response.JSON(w, http.StatusInternalServerError, "failed to invalidate session", nil, nil)
		return
	}

	response.JSON(w, http.StatusOK, "logout success", nil, nil)
}

// PUT /api/v1/app_user/change_password
func (h *AppUserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}
	// idAppUser, ok := r.Context().Value("app_user_id").(int) // sesuaikan context key
	loginID, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	// roleTypeID, ok := r.Context().Value("role_type_id").(int) // sesuaikan context key
	if !ok || loginID <= 0 {
		response.JSON(w, http.StatusUnauthorized, "user not found in context", nil, nil)
		return
	}

	var req model.AppUserChangePassword
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}
	if err := h.Svc.ChangePassword(r.Context(), loginID, &req); err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "password changed successfully", nil, nil)
}

// PUT /api/v1/app_user/change_password
func (h *AppUserHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil, nil)
		return
	}

	var req model.AppUserResetPassword
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"error": err.Error()})
		return
	}
	if err := h.Svc.ResetPassword(r.Context(), &req); err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}
	response.JSON(w, http.StatusOK, "password changed successfully", nil, nil)
}
