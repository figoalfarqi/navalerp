package service

import (
	"errors"
	"strings"
	"time"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/validation"
)

func validateRequest(req any) (error, map[string]string) {
	if err := validation.ValidateStruct(req); err != nil {
		return errors.New("validation error"), validation.ValidationErrors(err)
	}
	return nil, nil
}

func normalizeListOptions(opts model.ListOptions) model.ListOptions {
	if opts.Limit <= 0 {
		opts.Limit = 20
	}
	if opts.Limit > 10000 {
		opts.Limit = 10000
	}
	if opts.Offset < 0 {
		opts.Offset = 0
	}
	opts.Search = strings.TrimSpace(opts.Search)
	return opts
}

func validateDate(value *string) error {
	if value == nil || strings.TrimSpace(*value) == "" {
		return nil
	}
	trimmed := strings.TrimSpace(*value)
	if _, err := time.Parse("2006-01-02", trimmed); err == nil {
		*value = trimmed
		return nil
	}
	parsed, err := time.Parse(time.RFC3339, trimmed)
	if err != nil {
		return err
	}
	jakarta := time.FixedZone("Asia/Jakarta", 7*60*60)
	normalized := parsed.In(jakarta).Format("2006-01-02")
	*value = normalized
	return nil
}
