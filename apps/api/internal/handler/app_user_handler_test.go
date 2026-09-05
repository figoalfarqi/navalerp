package handler

import (
	"testing"

	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/internal/model"
)

func TestOwnerAndITDevCanManageEveryAdminUserType(t *testing.T) {
	for _, callerRoleID := range []int{
		helper.AppRoleOwner,
		helper.AppRoleITDev,
	} {
		t.Run(helper.FormatAppRole(callerRoleID), func(t *testing.T) {
			createDriver := &model.AppUserCreateRequest{AppRoleID: helper.AppRoleAdmin}
			if !prepareUserRoleForCreate(
				"/api/v1/admin/driver",
				callerRoleID,
				createDriver,
			) {
				t.Fatal("full-access role should be allowed to create a driver")
			}
			if createDriver.AppRoleID != helper.AppRoleDriver {
				t.Fatalf(
					"driver endpoint assigned role %d, want %d",
					createDriver.AppRoleID,
					helper.AppRoleDriver,
				)
			}

			createAnyUser := &model.AppUserCreateRequest{AppRoleID: helper.AppRoleSuperAdmin}
			if !prepareUserRoleForCreate(
				"/api/v1/admin/app_user",
				callerRoleID,
				createAnyUser,
			) {
				t.Fatal("full-access role should be allowed to create every known user role")
			}

			targetRole := helper.AppRoleITDev
			if !prepareUserRoleForUpdate(
				"/api/v1/admin/app_user/9",
				callerRoleID,
				helper.AppRoleAdmin,
				&model.AppUserUpdateRequest{AppRoleID: &targetRole},
			) {
				t.Fatal("full-access role should be allowed to assign every known role")
			}

			if !canManageUserTarget(
				"/api/v1/admin/app_user/9",
				callerRoleID,
				7,
				9,
				helper.AppRoleSuperAdmin,
			) {
				t.Fatal("full-access role should be allowed to manage every admin user target")
			}
		})
	}
}

func TestFullAdminOverrideDoesNotApplyToSelfServiceRoutes(t *testing.T) {
	for _, callerRoleID := range []int{
		helper.AppRoleOwner,
		helper.AppRoleITDev,
	} {
		if canManageUserTarget(
			"/api/v1/driver/driver/5",
			callerRoleID,
			7,
			5,
			helper.AppRoleDriver,
		) {
			t.Fatal("full admin override must not apply to driver self-service routes")
		}
	}
}
