package helper

import (
	"fmt"
	"reflect"
	"time"
)

func Ternary[T any](condition bool, vTrue, vFalse T) T {
	if condition {
		return vTrue
	}
	return vFalse
}

func TernaryPtr(v *int, defaultVal int) int {
	if v == nil {
		return defaultVal
	}
	return *v
}

func StrPtr(s string) *string {
	return &s
}

func StrPtrEqual(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func PtrOrZero[T ~int | ~int64 | ~float64](v *T) T {
	var zero T
	if v == nil {
		return zero
	}
	return *v
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	return v.Kind() == reflect.Ptr && v.IsNil()
}

// ==================================================
// Helper Functions
// ==================================================
func FloatToString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

func GetOrDefault(val *float64, def float64) float64 {
	if val != nil {
		return *val
	}
	return def
}

func DerefInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func DerefFloat64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func DerefString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func DerefTime(v *time.Time) time.Time {
	if v == nil {
		return time.Time{}
	}
	return *v
}
