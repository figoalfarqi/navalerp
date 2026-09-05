package repository

import (
	"encoding/json"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func binaryFilter(value string) *int {
	parsed, err := strconv.Atoi(value)
	if err != nil || (parsed != 0 && parsed != 1) {
		return nil
	}
	return &parsed
}

func positiveFilter(value string) *int {
	parsed, err := strconv.Atoi(value)
	if err != nil || parsed <= 0 {
		return nil
	}
	return &parsed
}

func scanJSONRow[T any](row pgx.Row) (*T, error) {
	var raw []byte
	if err := row.Scan(&raw); err != nil {
		return nil, err
	}
	var item T
	if err := json.Unmarshal(raw, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func scanJSONRows[T any](rows pgx.Rows) ([]T, error) {
	defer rows.Close()
	items := make([]T, 0)
	for rows.Next() {
		var raw []byte
		if err := rows.Scan(&raw); err != nil {
			return nil, err
		}
		var item T
		if err := json.Unmarshal(raw, &item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}
