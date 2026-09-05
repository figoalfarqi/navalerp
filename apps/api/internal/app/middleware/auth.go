package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/auth"
	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/pkg/database"
	"github.com/figoalfarqi/navalerp/pkg/response"
	"github.com/jackc/pgx/v5"
)

type ctxKey string

const (
	CtxAppUserID ctxKey = "app_user_id"
	CtxUserID    ctxKey = "user_id"
	CtxUsername  ctxKey = "username"
	CtxAppRoleID ctxKey = "app_role_id"
	CtxRole      ctxKey = "role"
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
		var currentUserID, currentUsername, currentRole string
		var currentAuthVersion int
		var currentIsActive bool
		err = pool.QueryRow(r.Context(), `
			SELECT user_id, username, role, auth_version, is_active
			FROM sys_users
			WHERE (user_id::text = $1 OR username = $2) AND deleted_at IS NULL`,
			claims.UserID, claims.Username,
		).Scan(
			&currentUserID,
			&currentUsername,
			&currentRole,
			&currentAuthVersion,
			&currentIsActive,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			response.JSON(w, http.StatusUnauthorized, "user session is no longer active", nil, nil)
			return
		}
		if err != nil {
			response.JSON(w, http.StatusServiceUnavailable, "authentication service unavailable", nil, nil)
			return
		}
		if !currentIsActive || (claims.AuthVersion > 0 && currentAuthVersion != claims.AuthVersion) {
			response.JSON(w, http.StatusUnauthorized, "user session is no longer active", nil, nil)
			return
		}

		ctx := context.WithValue(r.Context(), CtxUserID, currentUserID)
		ctx = context.WithValue(ctx, CtxAppUserID, claims.AppUserID)
		ctx = context.WithValue(ctx, CtxUsername, currentUsername)
		ctx = context.WithValue(ctx, CtxRole, currentRole)
		ctx = context.WithValue(ctx, CtxAppRoleID, claims.AppRoleID)

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
