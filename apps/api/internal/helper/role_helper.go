package helper

import (
	"fmt"
	"strings"
)

const (
	AppRoleDriver     = 1
	AppRoleChecker    = 2
	AppRoleOwner      = 3
	AppRoleITDev      = 4
	AppRoleSuperAdmin = 5
	AppRoleAdmin      = 6
)

func GetAppRoleIDsByRoleTypeName(roleName string) []int {
	switch roleName {
	case "admin":
		return []int{AppRoleOwner, AppRoleITDev, AppRoleSuperAdmin, AppRoleAdmin}
	case "admin_dashboard":
		return []int{AppRoleOwner, AppRoleITDev, AppRoleSuperAdmin, AppRoleAdmin}
	case "admin_report":
		return []int{AppRoleOwner, AppRoleITDev, AppRoleSuperAdmin, AppRoleAdmin}
	case "admin_technical":
		return []int{AppRoleITDev, AppRoleSuperAdmin}
	case "admin_operational":
		return []int{AppRoleSuperAdmin, AppRoleAdmin}
	case "driver":
		return []int{AppRoleDriver}
	case "checker":
		return []int{AppRoleChecker}
	case "allrole":
		return []int{
			AppRoleDriver,
			AppRoleChecker,
			AppRoleOwner,
			AppRoleITDev,
			AppRoleSuperAdmin,
			AppRoleAdmin,
		}
	default:
		return []int{}
	}
}

// CanonicalAppRoleID maps role names from both the legacy and current schemas
// to the stable role IDs used by JWT claims and authorization rules.
func CanonicalAppRoleID(roleName string) (int, bool) {
	normalized := strings.NewReplacer(
		" ", "",
		"_", "",
		"-", "",
	).Replace(strings.ToLower(strings.TrimSpace(roleName)))

	switch normalized {
	case "driver":
		return AppRoleDriver, true
	case "checker":
		return AppRoleChecker, true
	case "owner":
		return AppRoleOwner, true
	case "itdev", "itdeveloper":
		return AppRoleITDev, true
	case "superadmin":
		return AppRoleSuperAdmin, true
	case "admin", "systemadmin":
		return AppRoleAdmin, true
	default:
		return 0, false
	}
}

// AppRoleName returns the canonical, user-facing role name for a stable role ID.
func AppRoleName(roleID int) (string, bool) {
	switch roleID {
	case AppRoleDriver:
		return "driver", true
	case AppRoleChecker:
		return "checker", true
	case AppRoleOwner:
		return "owner", true
	case AppRoleITDev:
		return "itdev", true
	case AppRoleSuperAdmin:
		return "superadmin", true
	case AppRoleAdmin:
		return "admin", true
	default:
		return "", false
	}
}

func FormatAppRole(roleID int) string {
	roleName, ok := AppRoleName(roleID)
	if !ok {
		return fmt.Sprintf("%d (unknown)", roleID)
	}
	return fmt.Sprintf("%d (%s)", roleID, roleName)
}

func FormatAppRoles(roleIDs []int) string {
	formatted := make([]string, 0, len(roleIDs))
	for _, roleID := range roleIDs {
		formatted = append(formatted, FormatAppRole(roleID))
	}
	return "[" + strings.Join(formatted, ", ") + "]"
}
