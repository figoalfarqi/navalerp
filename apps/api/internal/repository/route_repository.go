package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RouteRepository struct {
	DB *pgxpool.Pool
}

func NewRouteRepository(db *pgxpool.Pool) *RouteRepository {
	return &RouteRepository{DB: db}
}

// Get retrieves a single route by route_id
func (r *RouteRepository) Get(ctx context.Context, id string) (*model.Route, error) {
	query := `SELECT route_id, route_code, route_name, origin_facility_id, destination_facility_id, distance_nautical_miles, estimated_transit_hours, risk_level, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM log_routes WHERE route_id = $1 AND deleted_at IS NULL`

	var m model.Route
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.RouteId, &m.RouteCode, &m.RouteName, &m.OriginFacilityId, &m.DestinationFacilityId, &m.DistanceNauticalMiles, &m.EstimatedTransitHours, &m.RiskLevel, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated route records
func (r *RouteRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Route, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(route_code ILIKE $%[1]d OR route_name ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM log_routes WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT route_id, route_code, route_name, origin_facility_id, destination_facility_id, distance_nautical_miles, estimated_transit_hours, risk_level, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM log_routes WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Route
	for rows.Next() {
		var m model.Route
		if err := rows.Scan(&m.RouteId, &m.RouteCode, &m.RouteName, &m.OriginFacilityId, &m.DestinationFacilityId, &m.DistanceNauticalMiles, &m.EstimatedTransitHours, &m.RiskLevel, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new route with optional child items
func (r *RouteRepository) Create(ctx context.Context, m *model.Route) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO log_routes (route_code, route_name, origin_facility_id, destination_facility_id, distance_nautical_miles, estimated_transit_hours, risk_level, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7::sea_route_risk_level_type, $8, $9, $10) RETURNING route_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.RouteCode, m.RouteName, m.OriginFacilityId, m.DestinationFacilityId, m.DistanceNauticalMiles, m.EstimatedTransitHours, m.RiskLevel, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing route
func (r *RouteRepository) Update(ctx context.Context, id string, m *model.Route) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE log_routes SET route_code = $1, route_name = $2, origin_facility_id = $3, destination_facility_id = $4, distance_nautical_miles = $5, estimated_transit_hours = $6, risk_level = $7::sea_route_risk_level_type, updated_by = $8, updated_at = CURRENT_TIMESTAMP WHERE route_id = $9`
	_, err = tx.Exec(ctx, updateQuery, m.RouteCode, m.RouteName, m.OriginFacilityId, m.DestinationFacilityId, m.DistanceNauticalMiles, m.EstimatedTransitHours, m.RiskLevel, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes route
func (r *RouteRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE log_routes SET deleted_at = CURRENT_TIMESTAMP WHERE route_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
