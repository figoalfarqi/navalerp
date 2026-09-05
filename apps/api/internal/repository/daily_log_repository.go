package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DailyLogRepository struct {
	DB *pgxpool.Pool
}

func NewDailyLogRepository(db *pgxpool.Pool) *DailyLogRepository {
	return &DailyLogRepository{DB: db}
}

// Get retrieves a single daily_log by log_id
func (r *DailyLogRepository) Get(ctx context.Context, id string) (*model.DailyLog, error) {
	query := `SELECT log_id, ship_id, log_date, latitude, longitude, heading_degrees, speed_knots, sea_state, weather_condition, fuel_remaining_liters, fresh_water_remaining_tons, tactical_summary, logged_by_user_id, created_at FROM ops_daily_logs WHERE log_id = $1`

	var m model.DailyLog
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.LogId, &m.ShipId, &m.LogDate, &m.Latitude, &m.Longitude, &m.HeadingDegrees, &m.SpeedKnots, &m.SeaState, &m.WeatherCondition, &m.FuelRemainingLiters, &m.FreshWaterRemainingTons, &m.TacticalSummary, &m.LoggedByUserId, &m.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated daily_log records
func (r *DailyLogRepository) List(ctx context.Context, opts model.ListOptions) ([]model.DailyLog, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(weather_condition ILIKE $%[1]d OR tactical_summary ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM ops_daily_logs WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT log_id, ship_id, log_date, latitude, longitude, heading_degrees, speed_knots, sea_state, weather_condition, fuel_remaining_liters, fresh_water_remaining_tons, tactical_summary, logged_by_user_id, created_at FROM ops_daily_logs WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.DailyLog
	for rows.Next() {
		var m model.DailyLog
		if err := rows.Scan(&m.LogId, &m.ShipId, &m.LogDate, &m.Latitude, &m.Longitude, &m.HeadingDegrees, &m.SpeedKnots, &m.SeaState, &m.WeatherCondition, &m.FuelRemainingLiters, &m.FreshWaterRemainingTons, &m.TacticalSummary, &m.LoggedByUserId, &m.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new daily_log with optional child items
func (r *DailyLogRepository) Create(ctx context.Context, m *model.DailyLog) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO ops_daily_logs (ship_id, log_date, latitude, longitude, heading_degrees, speed_knots, sea_state, weather_condition, fuel_remaining_liters, fresh_water_remaining_tons, tactical_summary, logged_by_user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING log_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.ShipId, m.LogDate, m.Latitude, m.Longitude, m.HeadingDegrees, m.SpeedKnots, m.SeaState, m.WeatherCondition, m.FuelRemainingLiters, m.FreshWaterRemainingTons, m.TacticalSummary, m.LoggedByUserId).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing daily_log
func (r *DailyLogRepository) Update(ctx context.Context, id string, m *model.DailyLog) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE ops_daily_logs SET ship_id = $1, log_date = $2, latitude = $3, longitude = $4, heading_degrees = $5, speed_knots = $6, sea_state = $7, weather_condition = $8, fuel_remaining_liters = $9, fresh_water_remaining_tons = $10, tactical_summary = $11, logged_by_user_id = $12 WHERE log_id = $13`
	_, err = tx.Exec(ctx, updateQuery, m.ShipId, m.LogDate, m.Latitude, m.Longitude, m.HeadingDegrees, m.SpeedKnots, m.SeaState, m.WeatherCondition, m.FuelRemainingLiters, m.FreshWaterRemainingTons, m.TacticalSummary, m.LoggedByUserId, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes daily_log
func (r *DailyLogRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM ops_daily_logs WHERE log_id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
