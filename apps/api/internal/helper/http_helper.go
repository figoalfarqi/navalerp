package helper

import (
	"net/http"
	"strconv"
	"strings"
)

// FromPathID mengambil ID integer dari path URL setelah prefix tertentu.
// Contoh: path="/api/v1/admin/city/123", prefix="/api/v1/admin/city/" → return 123
// Jika tidak valid, return 0.
func FromPathID(path, prefix string) int {
	if !strings.HasPrefix(path, prefix) {
		return 0
	}
	idStr := strings.TrimPrefix(path, prefix)
	idStr = strings.TrimSpace(idStr)
	if idStr == "" || strings.Contains(idStr, "/") { // hindari path seperti /123/extra
		return 0
	}
	id, _ := strconv.Atoi(idStr)
	return id
}

var RolePrefixes = []string{
	"/api/v1/admin/",
	"/api/v1/driver/",
	"/api/v1/client/",
	"/api/v1/checker/",
}

// GetIDAndRoleFromPath
// Contoh path:
// /api/v1/admin/city/12 → ("admin", 12)
// /api/v1/driver/client/55 → ("driver", 55)
func GetIDAndRoleFromPath(path string, resource string) (string, int) {
	for _, prefix := range RolePrefixes {
		fullPrefix := prefix + resource
		if strings.Contains(path, fullPrefix) {
			role := strings.TrimPrefix(prefix, "/api/v1/")
			role = strings.TrimSuffix(role, "/")
			id := FromPathID(path, fullPrefix+"/")
			return role, id
		}
	}
	return "", 0
}

// FromPathString mengambil string dari path URL setelah prefix tertentu.
// Contoh:
// path="/api/v1/admin/city/jakarta", prefix="/api/v1/admin/city/" → "jakarta"
// Jika tidak valid, return "".
func FromPathString(path, prefix string) string {
	if !strings.HasPrefix(path, prefix) {
		return ""
	}

	val := strings.TrimPrefix(path, prefix)
	val = strings.TrimSpace(val)

	// Hindari path seperti /jakarta/extra
	if val == "" || strings.Contains(val, "/") {
		return ""
	}

	return val
}

// GetStringAndRoleFromPath
// Contoh path:
// /api/v1/admin/city/jakarta → ("admin", "jakarta")
// /api/v1/driver/order/active → ("driver", "active")
func GetStringAndRoleFromPath(path string, resource string) (string, string) {
	for _, prefix := range RolePrefixes {
		fullPrefix := prefix + resource
		if strings.HasPrefix(path, fullPrefix+"/") {
			role := strings.TrimPrefix(prefix, "/api/v1/")
			role = strings.TrimSuffix(role, "/")

			val := FromPathString(path, fullPrefix+"/")
			return role, val
		}
	}
	return "", ""
}
func IsGetDataAdmin(path string) bool {
	return strings.Contains(path, "admin/admin")
}

func IsGetDataDriver(path string) bool {
	return strings.Contains(path, "admin/driver")
}

func IsGetDataClientPic(path string) bool {
	return strings.Contains(path, "admin/client_pic")
}

func IsGetDataChecker(path string) bool {
	return strings.Contains(path, "admin/checker")
}

// ParsePaging mengambil parameter limit dan offset dari query URL request.
// Default limit = 20 jika tidak valid, max 100. Default offset = 0 jika negatif.
func ParsePaging(r *http.Request) (limit, offset int) {
	q := r.URL.Query()
	limit, _ = strconv.Atoi(q.Get("limit"))
	offset, _ = strconv.Atoi(q.Get("offset"))

	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}
	return limit, offset
}

func AtoiSafe(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func AtoiSafeDefault(s string, def int) int {
	n := AtoiSafe(s)
	if n == 0 {
		return def
	}
	return n
}

// FromContextID mengambil nilai ID (int) dari context dengan key tertentu.
// Return nilai ID dan bool apakah ID ditemukan.
// func FromContextID(ctx context.Context, key interface{}) (int, bool) {
// 	v := ctx.Value(key)
// 	if v == nil {
// 		return 0, false
// 	}
// 	id, ok := v.(int)
// 	return id, ok
// }
