package repository

import (
	"context"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AppSettingRepository struct{ DB *pgxpool.Pool }

func NewAppSettingRepository(db *pgxpool.Pool) *AppSettingRepository {
	return &AppSettingRepository{DB: db}
}

func (r *AppSettingRepository) Create(ctx context.Context, req *model.AppSettingRequest, userID int) (int, error) {
	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO app_setting (
			app_setting_key, app_setting_value, app_setting_description,
			is_active, created_by, updated_by
		) VALUES ($1,$2,$3,$4,$5,$5)
		RETURNING app_setting_id`,
		req.AppSettingKey, req.AppSettingValue, req.AppSettingDescription, isActive, userID,
	).Scan(&id)
	return id, err
}

func (r *AppSettingRepository) Update(ctx context.Context, id int, req *model.AppSettingRequest, userID int) error {
	isActive := 1
	if req.IsActive != nil {
		isActive = *req.IsActive
	}
	_, err := r.DB.Exec(ctx, `
		UPDATE app_setting SET
			app_setting_key=$1, app_setting_value=$2,
			app_setting_description=$3, is_active=$4,
			updated_by=$5, updated_at=CURRENT_TIMESTAMP
		WHERE app_setting_id=$6 AND deleted_at IS NULL`,
		req.AppSettingKey, req.AppSettingValue, req.AppSettingDescription, isActive, userID, id,
	)
	return err
}

func (r *AppSettingRepository) SoftDelete(ctx context.Context, id, userID int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE app_setting
		SET deleted_by=$1, deleted_at=CURRENT_TIMESTAMP, updated_by=$1, updated_at=CURRENT_TIMESTAMP
		WHERE app_setting_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

func (r *AppSettingRepository) GetByID(ctx context.Context, id int) (*model.AppSetting, error) {
	return scanJSONRow[model.AppSetting](r.DB.QueryRow(ctx,
		`SELECT to_jsonb(s) FROM app_setting s WHERE app_setting_id=$1 AND deleted_at IS NULL`, id))
}

func (r *AppSettingRepository) List(ctx context.Context, opts model.ListOptions, checkerOnly bool) ([]model.AppSetting, error) {
	query := `
		SELECT to_jsonb(s)
		FROM app_setting s
		WHERE s.deleted_at IS NULL
		  AND ($1::int IS NULL OR s.app_setting_id < $1)
		  AND (
			NOT $2::boolean OR s.app_setting_key IN (
				'checker.location_cooldown_minutes',
				'checker.truck_cooldown_minutes'
			)
		  )
		  AND ($3::smallint IS NULL OR s.is_active=$3)
		  AND ($4='' OR s.app_setting_key ILIKE '%' || $4 || '%'
			OR COALESCE(s.app_setting_description,'') ILIKE '%' || $4 || '%')
		ORDER BY s.app_setting_id DESC
		LIMIT $5 OFFSET $6`
	rows, err := r.DB.Query(ctx, query, opts.CursorKey, checkerOnly, opts.IsActive, opts.Search, opts.Limit, opts.Offset)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.AppSetting](rows)
}
