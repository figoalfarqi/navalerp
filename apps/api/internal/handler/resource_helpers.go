package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/figoalfarqi/apipml/internal/app/middleware"
	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/pkg/response"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func requestUserID(r *http.Request) (int, bool) {
	id, ok := r.Context().Value(middleware.CtxAppUserID).(int)
	return id, ok && id > 0
}

func requestRoleID(r *http.Request) (int, bool) {
	id, ok := r.Context().Value(middleware.CtxAppRoleID).(int)
	return id, ok && id > 0
}

func requestID(r *http.Request) int {
	value := r.PathValue("id")
	if value == "" {
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(parts) > 0 {
			value = parts[len(parts)-1]
		}
	}
	id, _ := strconv.Atoi(value)
	return id
}

func decodeRequest[T any](w http.ResponseWriter, r *http.Request) (*T, bool) {
	var req T
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid json format", nil, map[string]string{"body": err.Error()})
		return nil, false
	}
	return &req, true
}

func queryValues(r *http.Request) url.Values {
	values := r.URL.Query()
	nested, err := url.ParseQuery(values.Get("query"))
	if err == nil {
		for key, value := range nested {
			if values.Get(key) == "" {
				values[key] = value
			}
		}
	}
	return values
}

func intPointer(value string) *int {
	id, err := strconv.Atoi(value)
	if err != nil || id <= 0 {
		return nil
	}
	return &id
}

func binaryIntPointer(value string) *int {
	id, err := strconv.Atoi(value)
	if err != nil || (id != 0 && id != 1) {
		return nil
	}
	return &id
}

func parseListOptions(r *http.Request) model.ListOptions {
	q := queryValues(r)
	limit, _ := strconv.Atoi(q.Get("limit"))
	if limit <= 0 {
		limit = 20
	}
	if limit > 1000 {
		limit = 1000
	}
	offset, _ := strconv.Atoi(q.Get("offset"))
	if offset < 0 {
		offset = 0
	}
	sort := strings.ToLower(q.Get("sort"))
	if sort != "asc" {
		sort = "desc"
	}
	filters := make(map[string]string, len(q))
	for key, values := range q {
		if len(values) > 0 && values[0] != "" {
			filters[key] = values[0]
		}
	}
	return model.ListOptions{
		Limit:     limit,
		Offset:    offset,
		CursorKey: intPointer(q.Get("cursor_key")),
		Sort:      sort,
		Search: firstNonEmpty(
			q.Get("search"),
			q.Get("q"),
			q.Get("project_name"),
			q.Get("project_code"),
			q.Get("app_setting_key"),
			q.Get("port_name"),
			q.Get("vessel_name"),
		),
		ProjectID:       intPointer(q.Get("project_id")),
		TransportID:     intPointer(q.Get("project_transport_id")),
		DriverID:        intPointer(q.Get("driver_id")),
		CheckerID:       intPointer(q.Get("checker_id")),
		TruckID:         intPointer(q.Get("truck_id")),
		VesselID:        intPointer(q.Get("vessel_id")),
		PortID:          intPointer(q.Get("port_id")),
		CargoTypeID:     intPointer(q.Get("cargo_type_id")),
		StatusTypeID:    intPointer(q.Get("project_transport_status_type_id")),
		IsActive:        binaryIntPointer(q.Get("is_active")),
		IsCompleted:     binaryIntPointer(q.Get("is_completed")),
		IsFraud:         binaryIntPointer(q.Get("is_fraud")),
		IsPosted:        binaryIntPointer(q.Get("is_posted")),
		RouteType:       q.Get("route_type"),
		TransactionKind: q.Get("transaction_kind"),
		Filters:         filters,
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func parseDateRange(r *http.Request, defaultToday bool) (*time.Time, *time.Time, error) {
	q := queryValues(r)
	dateValue := firstNonEmpty(
		q.Get("date"),
		q.Get("transported_at"),
		q.Get("status_time"),
		q.Get("transaction_date"),
		q.Get("adjustment_date"),
		q.Get("created_at"),
	)
	after := firstNonEmpty(
		q.Get("transported_at_after"),
		q.Get("status_time_after"),
		q.Get("transaction_date_after"),
		q.Get("adjustment_date_after"),
		q.Get("date_after"),
		q.Get("created_at_after"),
	)
	before := firstNonEmpty(
		q.Get("transported_at_before"),
		q.Get("status_time_before"),
		q.Get("transaction_date_before"),
		q.Get("adjustment_date_before"),
		q.Get("date_before"),
		q.Get("created_at_before"),
	)
	// All operational dates are interpreted in Western Indonesia Time.
	location := time.FixedZone("Asia/Jakarta", 7*60*60)
	if dateValue != "" {
		day, err := time.ParseInLocation("2006-01-02", dateValue, location)
		if err != nil {
			return nil, nil, err
		}
		to := day.AddDate(0, 0, 1)
		return &day, &to, nil
	}
	var from, to *time.Time
	if after != "" {
		value, err := time.ParseInLocation("2006-01-02", after, location)
		if err != nil {
			return nil, nil, err
		}
		from = &value
	}
	if before != "" {
		value, err := time.ParseInLocation("2006-01-02", before, location)
		if err != nil {
			return nil, nil, err
		}
		value = value.AddDate(0, 0, 1)
		to = &value
	}
	if defaultToday && from == nil && to == nil {
		now := time.Now()
		day := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
		next := day.AddDate(0, 0, 1)
		from, to = &day, &next
	}
	return from, to, nil
}

func writeResourceError(w http.ResponseWriter, err error, fields map[string]string) {
	if fields != nil {
		response.JSON(w, http.StatusBadRequest, "validation error", nil, fields)
		return
	}
	if errors.Is(err, pgx.ErrNoRows) {
		response.JSON(w, http.StatusNotFound, "data not found", nil, nil)
		return
	}
	switch err.Error() {
	case "checker is not assigned to this project":
		response.JSON(w, http.StatusForbidden, "checker is not assigned to this project", nil, nil)
		return
	case "project truck assignment is inactive or outside its access period":
		response.JSON(w, http.StatusConflict, "project truck assignment is not available", nil, nil)
		return
	case "checker_id is not an active checker":
		response.JSON(w, http.StatusBadRequest, "checker_id is not an active checker", nil, nil)
		return
	case "stockpile balance cannot be negative",
		"stockpile volume exceeds capacity",
		"stockpile weight exceeds capacity":
		response.JSON(w, http.StatusConflict, err.Error(), nil, nil)
		return
	}
	var databaseError *pgconn.PgError
	if errors.As(err, &databaseError) {
		switch databaseError.Code {
		case "23505":
			response.JSON(w, http.StatusConflict, "data already exists", nil, map[string]string{"constraint": databaseError.ConstraintName})
			return
		case "23503":
			response.JSON(w, http.StatusConflict, "data is still referenced or its reference is invalid", nil, map[string]string{"constraint": databaseError.ConstraintName})
			return
		case "23502", "23514", "22P02":
			response.JSON(w, http.StatusBadRequest, "data does not satisfy database rules", nil, map[string]string{"constraint": databaseError.ConstraintName})
			return
		}
	}
	response.JSON(w, http.StatusInternalServerError, "request failed", nil, nil)
}

func handleCreate[Req, Resp any](
	w http.ResponseWriter,
	r *http.Request,
	fn func(context.Context, int, *Req) (*Resp, error, map[string]string),
) {
	userID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	req, ok := decodeRequest[Req](w, r)
	if !ok {
		return
	}
	item, err, fields := fn(r.Context(), userID, req)
	if err != nil {
		writeResourceError(w, err, fields)
		return
	}
	response.JSON(w, http.StatusCreated, "created", item, nil)
}

func handleUpdate[Req, Resp any](
	w http.ResponseWriter,
	r *http.Request,
	fn func(context.Context, int, int, *Req) (*Resp, error, map[string]string),
) {
	userID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	id := requestID(r)
	if id <= 0 {
		response.JSON(w, http.StatusBadRequest, "invalid id", nil, map[string]string{"id": "invalid"})
		return
	}
	req, ok := decodeRequest[Req](w, r)
	if !ok {
		return
	}
	item, err, fields := fn(r.Context(), userID, id, req)
	if err != nil {
		writeResourceError(w, err, fields)
		return
	}
	response.JSON(w, http.StatusOK, "updated", item, nil)
}

func handleDelete(
	w http.ResponseWriter,
	r *http.Request,
	fn func(context.Context, int, int) error,
) {
	userID, ok := requestUserID(r)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
		return
	}
	id := requestID(r)
	if id <= 0 {
		response.JSON(w, http.StatusBadRequest, "invalid id", nil, map[string]string{"id": "invalid"})
		return
	}
	if err := fn(r.Context(), userID, id); err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "deleted", map[string]int{"id": id}, nil)
}

func handleGet[Resp any](
	w http.ResponseWriter,
	r *http.Request,
	get func(context.Context, int) (*Resp, error),
	list func(context.Context, model.ListOptions) ([]Resp, error),
) {
	if id := requestID(r); id > 0 {
		item, err := get(r.Context(), id)
		if err != nil {
			writeResourceError(w, err, nil)
			return
		}
		response.JSON(w, http.StatusOK, "ok", item, nil)
		return
	}
	opts := parseListOptions(r)
	items, err := list(r.Context(), opts)
	if err != nil {
		writeResourceError(w, err, nil)
		return
	}
	response.JSON(w, http.StatusOK, "ok", model.ListResult[Resp]{Items: items, Limit: opts.Limit, Offset: opts.Offset}, nil)
}
