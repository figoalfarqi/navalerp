package helper

import "testing"

func TestCanonicalAppRoleID(t *testing.T) {
	tests := []struct {
		name     string
		roleName string
		wantID   int
		wantOK   bool
	}{
		{name: "driver", roleName: "Driver", wantID: AppRoleDriver, wantOK: true},
		{name: "checker", roleName: "Checker", wantID: AppRoleChecker, wantOK: true},
		{name: "owner", roleName: "owner", wantID: AppRoleOwner, wantOK: true},
		{name: "it dev legacy", roleName: "IT Developer", wantID: AppRoleITDev, wantOK: true},
		{name: "it dev current", roleName: "itdev", wantID: AppRoleITDev, wantOK: true},
		{name: "super admin", roleName: "Super Admin", wantID: AppRoleSuperAdmin, wantOK: true},
		{name: "admin", roleName: "admin", wantID: AppRoleAdmin, wantOK: true},
		{name: "system admin legacy", roleName: "System Admin", wantID: AppRoleAdmin, wantOK: true},
		{name: "unsupported role", roleName: "client", wantID: 0, wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, gotOK := CanonicalAppRoleID(tt.roleName)
			if gotID != tt.wantID || gotOK != tt.wantOK {
				t.Fatalf(
					"CanonicalAppRoleID(%q) = (%d, %t), want (%d, %t)",
					tt.roleName,
					gotID,
					gotOK,
					tt.wantID,
					tt.wantOK,
				)
			}
		})
	}
}

func TestFormatAppRole(t *testing.T) {
	if got := FormatAppRole(AppRoleOwner); got != "3 (owner)" {
		t.Fatalf("FormatAppRole(owner) = %q, want %q", got, "3 (owner)")
	}
	if got := FormatAppRoles([]int{AppRoleSuperAdmin, AppRoleAdmin}); got != "[5 (superadmin), 6 (admin)]" {
		t.Fatalf(
			"FormatAppRoles(superadmin, admin) = %q, want %q",
			got,
			"[5 (superadmin), 6 (admin)]",
		)
	}
	if got := FormatAppRole(99); got != "99 (unknown)" {
		t.Fatalf("FormatAppRole(unknown) = %q, want %q", got, "99 (unknown)")
	}
}
