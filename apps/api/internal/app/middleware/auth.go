package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/figoalfarqi/apipml/config"
	"github.com/figoalfarqi/apipml/internal/auth"
	"github.com/figoalfarqi/apipml/internal/helper"
	"github.com/figoalfarqi/apipml/pkg/database"
	"github.com/figoalfarqi/apipml/pkg/response"
	"github.com/jackc/pgx/v5"
)

type ctxKey string

const (
	CtxAppUserID ctxKey = "app_user_id"
	CtxUsername  ctxKey = "username"
	CtxAppRoleID ctxKey = "app_role_id"
)

func AuthJWT(cfg *config.Config, allowedRoles []int, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := strings.Fields(r.Header.Get("Authorization"))
		if len(authorization) != 2 || !strings.EqualFold(authorization[0], "Bearer") {
			response.JSON(w, http.StatusUnauthorized, "missing bearer token", nil, map[string]string{"authorization": "required"})
			return
		}

		claims, err := auth.ParseToken(cfg.JWTKey, authorization[1])
		if err != nil {
			response.JSON(w, http.StatusUnauthorized, "invalid token", nil, map[string]string{"error": err.Error()})
			return
		}

		pool := database.GetPool()
		if pool == nil {
			response.JSON(w, http.StatusServiceUnavailable, "authentication service unavailable", nil, nil)
			return
		}
		var currentUsername, currentRoleName string
		var currentStatusID, currentAuthVersion int
		err = pool.QueryRow(r.Context(), `
			SELECT u.username,u.app_user_status_id,u.auth_version,r.app_role_name
			FROM app_user u
			JOIN app_role r ON r.app_role_id=u.app_role_id AND r.deleted_at IS NULL
			WHERE u.app_user_id=$1 AND u.deleted_at IS NULL`,
			claims.AppUserID,
		).Scan(
			&currentUsername,
			&currentStatusID,
			&currentAuthVersion,
			&currentRoleName,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			response.JSON(w, http.StatusUnauthorized, "user session is no longer active", nil, nil)
			return
		}
		if err != nil {
			response.JSON(w, http.StatusServiceUnavailable, "authentication service unavailable", nil, nil)
			return
		}
		currentRoleID, ok := helper.CanonicalAppRoleID(currentRoleName)
		if !ok {
			response.JSON(w, http.StatusUnauthorized, "user role is no longer active", nil, nil)
			return
		}
		if currentStatusID != 1 ||
			currentAuthVersion != claims.AuthVersion ||
			currentUsername != claims.Username ||
			currentRoleID != claims.AppRoleID {
			response.JSON(w, http.StatusUnauthorized, "user session is no longer active", nil, nil)
			return
		}

		effectiveAllowedRoles := allowedRolesForPath(r.URL.Path, allowedRoles)
		allowed := containsRole(effectiveAllowedRoles, currentRoleID)
		if !allowed {
			msg := fmt.Sprintf(
				"forbidden: %s not allowed, only for %s",
				helper.FormatAppRole(currentRoleID),
				helper.FormatAppRoles(effectiveAllowedRoles),
			)
			response.JSON(w, http.StatusForbidden, msg, nil, nil)
			return
		}

		ctx := context.WithValue(r.Context(), CtxAppUserID, claims.AppUserID)
		ctx = context.WithValue(ctx, CtxUsername, currentUsername)
		ctx = context.WithValue(ctx, CtxAppRoleID, currentRoleID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func isRoleAllowed(path string, currentRoleID int, allowedRoles []int) bool {
	return containsRole(allowedRolesForPath(path, allowedRoles), currentRoleID)
}

func allowedRolesForPath(path string, configuredRoles []int) []int {
	if !strings.HasPrefix(path, "/api/v1/admin/") {
		return configuredRoles
	}

	// Owner and IT developer may perform every action exposed through the
	// protected admin API.
	roles := make([]int, 0, len(configuredRoles)+2)
	roles = append(roles, helper.AppRoleOwner, helper.AppRoleITDev)
	for _, roleID := range configuredRoles {
		if roleID != helper.AppRoleOwner && roleID != helper.AppRoleITDev {
			roles = append(roles, roleID)
		}
	}
	return roles
}

func containsRole(roles []int, currentRoleID int) bool {
	for _, roleID := range roles {
		if currentRoleID == roleID {
			return true
		}
	}
	return false
}
