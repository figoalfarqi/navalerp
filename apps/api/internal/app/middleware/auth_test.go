package middleware

import (
	"testing"

	"github.com/figoalfarqi/apipml/internal/helper"
)

func TestIsRoleAllowed(t *testing.T) {
	operationalRoles := []int{helper.AppRoleSuperAdmin, helper.AppRoleAdmin}

	tests := []struct {
		name          string
		path          string
		currentRoleID int
		allowedRoles  []int
		want          bool
	}{
		{
			name:          "owner may mutate all admin resources",
			path:          "/api/v1/admin/truck_type/2",
			currentRoleID: helper.AppRoleOwner,
			allowedRoles:  operationalRoles,
			want:          true,
		},
		{
			name:          "itdev may mutate all admin resources",
			path:          "/api/v1/admin/truck_type/2",
			currentRoleID: helper.AppRoleITDev,
			allowedRoles:  operationalRoles,
			want:          true,
		},
		{
			name:          "owner bypass is limited to admin API",
			path:          "/api/v1/checker/project_transport",
			currentRoleID: helper.AppRoleOwner,
			allowedRoles:  []int{helper.AppRoleChecker},
			want:          false,
		},
		{
			name:          "itdev bypass is limited to admin API",
			path:          "/api/v1/checker/project_transport",
			currentRoleID: helper.AppRoleITDev,
			allowedRoles:  []int{helper.AppRoleChecker},
			want:          false,
		},
		{
			name:          "configured admin remains allowed",
			path:          "/api/v1/admin/truck_type/2",
			currentRoleID: helper.AppRoleAdmin,
			allowedRoles:  operationalRoles,
			want:          true,
		},
		{
			name:          "checker remains forbidden from admin mutation",
			path:          "/api/v1/admin/truck_type/2",
			currentRoleID: helper.AppRoleChecker,
			allowedRoles:  operationalRoles,
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRoleAllowed(tt.path, tt.currentRoleID, tt.allowedRoles); got != tt.want {
				t.Fatalf("isRoleAllowed() = %t, want %t", got, tt.want)
			}
		})
	}
}

func TestAllowedRolesForAdminPathIncludesNamedFullAccessOverrides(t *testing.T) {
	got := helper.FormatAppRoles(allowedRolesForPath(
		"/api/v1/admin/truck_type/2",
		[]int{helper.AppRoleSuperAdmin, helper.AppRoleAdmin},
	))
	want := "[3 (owner), 4 (itdev), 5 (superadmin), 6 (admin)]"
	if got != want {
		t.Fatalf("formatted effective roles = %q, want %q", got, want)
	}
}
