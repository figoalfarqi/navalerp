package helper

import (
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"
)

// ParseIsActiveFilter digunakan untuk menghasilkan filter is_active universal
// return:
// - string value untuk dipakai di filter map ("" berarti tidak difilter)
// - string error jika is_active invalid
func ParseIsActiveFilter(params url.Values) (string, string) {
	raw := params.Get("is_active")

	if raw == "" {
		// default → hanya aktif
		return "1", ""
	}

	switch raw {
	case "0":
		return "0", ""
	case "1":
		return "1", ""
	case "2":
		// tampilkan semua → tidak difilter
		return "", ""
	default:
		return "", "is_active must be 0, 1, 2, or empty"
	}
}

// ParseQueryFilter universal untuk ?query=client_name=AAA&city_id=2
func ParseQueryFilter(params url.Values, allowedKeys []string) map[string]string {
	result := make(map[string]string)

	raw := params.Get("query")
	if raw == "" {
		return result
	}

	decoded, _ := url.QueryUnescape(raw)
	pairs := strings.Split(decoded, "&")

	for _, p := range pairs {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			continue
		}

		key := kv[0]
		value, _ := url.QueryUnescape(kv[1])

		for _, allowed := range allowedKeys {
			if key == allowed {
				result[key] = value
			}
		}
	}

	if slices.Contains(allowedKeys, "is_active") {
		if v, err := strconv.Atoi(result["is_active"]); err != nil || v < 0 || v > 2 {
			result["is_active"] = "1"
		} else if result["is_active"] == "2" {
			result["is_active"] = ""
		}
	}

	if slices.Contains(allowedKeys, "app_user_status_id") {
		if v, err := strconv.Atoi(result["app_user_status_id"]); err != nil || v < -1 || (v > 3 && v != 99) {
			result["app_user_status_id"] = "1"
		} else if result["app_user_status_id"] == "99" {
			result["app_user_status_id"] = ""
		}
	}

	return result
}

// NormalizeTimeFilters menormalkan semua filter waktu agar UTC-safe,
// dengan logika “awal hari” untuk *_after dan “akhir hari” untuk *_before.
// Contoh hasil: created_at_before=2025-11-18T16:59:59Z (UTC, dari +7)
// NormalizeTimeFilters menormalkan filter waktu menjadi UTC-safe.
// Param:
// - filters: map filter input dari query
// - extraTimeKeys: filter tambahan spesifik tabel ("key" -> "after"/"before")
func NormalizeTimeFilters(filters map[string]string, extraTimeKeys map[string]string) (map[string]string, error) {
	if filters == nil {
		return nil, nil
	}

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, fmt.Errorf("failed to load timezone: %v", err)
	}

	// Default filter yang selalu ada:
	defaultTimeKeys := map[string]string{
		"created_at_after":  "after",
		"created_at_before": "before",
		"updated_at_after":  "after",
		"updated_at_before": "before",
	}

	// Gabungkan default + tambahan
	timeKeys := map[string]string{}
	for k, v := range defaultTimeKeys {
		timeKeys[k] = v
	}
	for k, v := range extraTimeKeys {
		timeKeys[k] = v
	}

	for key, mode := range timeKeys {
		if val, ok := filters[key]; ok {
			t, err := time.Parse(time.RFC3339, val)
			if err != nil {
				return nil, fmt.Errorf("invalid time format for %s: %v", key, err)
			}

			t = t.In(loc)

			if mode == "after" {
				t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
			} else if mode == "before" {
				t = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999_000_000, loc)
			}

			filters[key] = t.UTC().Format(time.RFC3339)
		}
	}

	return filters, nil
}

// ParseQuerySort mengambil order_by dan sort_type dari query,
// memvalidasi berdasarkan allowedColumns,
// dan memberikan nilai default: created_at DESC
func ParseQuerySort(
	params url.Values,
	allowedColumns []string,
) (orderBy string, sort string) {

	// --- Ambil dari query ---
	orderBy = strings.TrimSpace(params.Get("order_by"))
	sort = strings.ToLower(strings.TrimSpace(params.Get("sort")))

	// --- Validasi order_by ---
	if orderBy == "" || !slices.Contains(allowedColumns, orderBy) {
		orderBy = "created_at"
	}

	// --- Validasi sort_type ---
	if sort != "asc" && sort != "desc" {
		sort = "desc"
	}

	return orderBy, sort
}

type CursorData struct {
	Value interface{} // bisa string, int, float64, time.Time
	Key   *int        // selalu int (primary key)
	Type  string      // "string" | "number" | "date"
}

func ParseCursorParams(query url.Values) (*CursorData, error) {
	cursorValueStr := query.Get("cursor_value")
	cursorKeyStr := query.Get("cursor_key")
	cursorType := query.Get("cursor_type")

	// Tidak ada cursor — return nil (first page)
	if cursorKeyStr == "" {
		return nil, nil
	}
	// Cursor Key → harus integer
	cursorKey, err := strconv.Atoi(cursorKeyStr)
	if err != nil {
		return nil, fmt.Errorf("invalid cursor_key, must be integer: %w", err)
	}

	if cursorValueStr == "" || cursorType == "" {
		return &CursorData{
			Value: nil,
			Key:   &cursorKey,
			Type:  cursorType,
		}, nil
	}

	var cursorValue interface{}

	switch cursorType {
	case "number":
		num, err := strconv.ParseFloat(cursorValueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid cursor_value for number: %w", err)
		}
		cursorValue = num

	case "string":
		cursorValue = cursorValueStr

	case "date":
		// ISO8601 recommended
		t, err := time.Parse(time.RFC3339, cursorValueStr)
		if err != nil {
			return nil, fmt.Errorf("invalid cursor_value for date, must be RFC3339: %w", err)
		}
		cursorValue = t

	default:
		return nil, fmt.Errorf("invalid cursor_type: %s (allowed: string|number|date)", cursorType)
	}

	return &CursorData{
		Value: cursorValue,
		Key:   &cursorKey,
		Type:  cursorType,
	}, nil
}

func BuildCursorCondition(orderBy, tableKey, sort string, cursorValue, cursorKey interface{}, argPos int) (string, []interface{}, int) {

	// Jika tidak ada cursor → tidak tambahkan kondisi apa pun
	if IsNil(cursorKey) {
		return "", nil, argPos
	}

	cursorKeyInt, ok := cursorKey.(*int)
	if !ok || cursorKeyInt == nil {
		return "", nil, argPos
	}

	var condition string

	// fmt.Println("AAAAAAAAAAA", tableKey)

	if IsNil(cursorValue) {
		// fmt.Println("BBBBBBBBB", tableKey)

		if sort == "asc" {
			condition = fmt.Sprintf(`
				AND (
					(%s IS NULL AND %s > $%d)
				)
			`, orderBy, tableKey, argPos)
			// fmt.Println("CCCCCCCCCC", fmt.Sprintf(`
			// 	AND (
			// 		(%s IS NOT NULL)
			// 		OR (%s IS NULL AND %s > $%d)
			// 	)
			// `, orderBy, orderBy, tableKey, argPos), *cursorKeyInt)

		} else {
			condition = fmt.Sprintf(`
				AND (
					(%s IS NOT NULL) OR
					(%s IS NULL AND %s < $%d)
				)
			`, orderBy, orderBy, tableKey, argPos)
			// fmt.Println("DDDDDDDDDD", fmt.Sprintf(`
			// 	AND (
			// 		(%s IS NULL AND %s < $%d)
			// 	)
			// `, orderBy, tableKey, argPos), *cursorKeyInt)

		}
		args := []interface{}{*cursorKeyInt}
		argPos += 1

		return condition, args, argPos
	}

	if sort == "asc" {
		condition = fmt.Sprintf(`
            AND (
                (%s > $%d)
                OR (%s = $%d AND %s > $%d)
            )
        `, orderBy, argPos, orderBy, argPos, tableKey, argPos+1)
		// fmt.Println("EEEEEE", fmt.Sprintf(`
		//     AND (
		//         (%s > $%d)
		//         OR (%s = $%d AND %s > $%d)
		//     )
		// `, orderBy, argPos, orderBy, argPos, tableKey, argPos+1), cursorValue, *cursorKeyInt)
	} else { // desc
		condition = fmt.Sprintf(`
            AND (
                (%s < $%d)
                OR (%s = $%d AND %s < $%d)
            )
        `, orderBy, argPos, orderBy, argPos, tableKey, argPos+1)
		// fmt.Println("FFFFFFF", fmt.Sprintf(`
		//     AND (
		//         (%s < $%d)
		//         OR (%s = $%d AND %s < $%d)
		//     )
		// `, orderBy, argPos, orderBy, argPos, tableKey, argPos+1), cursorValue, *cursorKeyInt)
	}

	args := []interface{}{cursorValue, *cursorKeyInt}
	argPos += 2

	return condition, args, argPos
}

func AddILIKEFilter(baseQuery *string, args *[]interface{}, argPos *int, column string, value string) {
	if value != "" {
		*baseQuery += fmt.Sprintf(" AND %s ILIKE $%d", column, *argPos)
		*args = append(*args, "%"+value+"%")
		*argPos++
	}
}

func AddRangeFilter(baseQuery *string, args *[]interface{}, argPos *int, column, minVal, maxVal string) {
	if minVal != "" {
		*baseQuery += fmt.Sprintf(" AND %s >= $%d", column, *argPos)
		*args = append(*args, minVal)
		*argPos++
	}
	if maxVal != "" {
		*baseQuery += fmt.Sprintf(" AND %s <= $%d", column, *argPos)
		*args = append(*args, maxVal)
		*argPos++
	}
}

func AddExactFilter(baseQuery *string, args *[]interface{}, argPos *int, column, value string) {
	if value != "" {
		*baseQuery += fmt.Sprintf(" AND %s = $%d", column, *argPos)
		*args = append(*args, value)
		*argPos++
	}
}
