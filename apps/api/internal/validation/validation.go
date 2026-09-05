package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Map error -> ramah FE
func ValidationErrors(err error) map[string]string {
	out := map[string]string{}
	if err == nil {
		return out
	}

	for _, e := range err.(validator.ValidationErrors) {
		field := toSnake(e.Field())
		switch e.Tag() {
		case "required":
			out[field] = "field is required"
		case "min":
			out[field] = fmt.Sprintf("value is too short, min %s characters", e.Param())
		case "max":
			out[field] = fmt.Sprintf("value is too long, max %s characters", e.Param())
		case "alphanum":
			out[field] = "only alphanumeric allowed"
		case "oneof":
			out[field] = fmt.Sprintf("invalid option, must be one of [%s]", e.Param())
		case "gt":
			out[field] = fmt.Sprintf("must be greater than %s", e.Param())
		default:
			out[field] = "invalid value"
		}
	}
	return out
}

func toSnake(s string) string {
	var b strings.Builder
	runes := []rune(s)
	for i, r := range runes {
		if i > 0 && r >= 'A' && r <= 'Z' {
			// Jangan tambah underscore jika huruf kapital sebelumnya juga kapital
			prev := runes[i-1]
			if !(prev >= 'A' && prev <= 'Z') {
				b.WriteByte('_')
			}
		}
		b.WriteRune(r)
	}
	return strings.ToLower(b.String())
}

func ValidateStruct(i interface{}) error {
	return validate.Struct(i)
}

// ===== Trimming helpers =====
func TrimStrings(ss ...*string) {
	for _, p := range ss {
		if p != nil {
			*p = strings.TrimSpace(*p)
		}
	}
}
