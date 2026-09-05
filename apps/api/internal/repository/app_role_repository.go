package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/figoalfarqi/apipml/internal/helper"
	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AppRoleRepository struct {
	DB *pgxpool.Pool
}

func NewAppRoleRepository(db *pgxpool.Pool) *AppRoleRepository {
	return &AppRoleRepository{DB: db}
}

// ==================================================
// Create
// ==================================================
func (r *AppRoleRepository) Create(ctx context.Context, m *model.AppRole) (int, error) {
	query := `
		INSERT INTO app_role (
			app_role_type_id,
			app_role_name,
			app_role_description,
			is_active,
			created_by,
			updated_by
		) VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING app_role_id
	`

	var id int
	err := r.DB.QueryRow(ctx, query,
		m.AppRoleTypeID,
		m.AppRoleName,
		m.AppRoleDescription,
		m.IsActive,
		m.CreatedBy,
		m.UpdatedBy,
	).Scan(&id)

	return id, err
}

// ==================================================
// Update (Dynamic)
// ==================================================
func (r *AppRoleRepository) Update(ctx context.Context, id int, m *model.AppRole) error {

	setIsActive := ""
	args := []interface{}{
		m.AppRoleTypeID,      // $1
		m.AppRoleName,        // $2
		m.AppRoleDescription, // $3
	}

	argPos := 4

	if m.IsActive != -1 { // -1 sentinel → tidak update kolom is_active
		setIsActive = fmt.Sprintf(", is_active = $%d", argPos)
		args = append(args, m.IsActive)
		argPos++
	}

	args = append(args,
		m.UpdatedBy, // updated_by → $argPos
		time.Now(),  // updated_at  → $argPos+1
		id,          // where       → $argPos+2
	)

	query := fmt.Sprintf(`
		UPDATE app_role
		SET app_role_type_id = $1,
			app_role_name = $2,
			app_role_description = $3
			%s,
			updated_by = $%d,
			updated_at = $%d
		WHERE app_role_id = $%d AND deleted_at IS NULL
	`,
		setIsActive,
		argPos,   // updated_by
		argPos+1, // updated_at
		argPos+2, // where
	)

	_, err := r.DB.Exec(ctx, query, args...)
	return err
}

// ==================================================
// Soft Delete
// ==================================================
func (r *AppRoleRepository) SoftDelete(ctx context.Context, deletedBy, id int) error {
	query := `
		UPDATE app_role
		SET deleted_at = $1,
			deleted_by = $2,
			is_active = 0,
			updated_at = $1,
			updated_by = $2
		WHERE app_role_id = $3 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(ctx, query, time.Now(), deletedBy, id)
	return err
}

// ==================================================
// Get By ID
// ==================================================
func (r *AppRoleRepository) GetByID(ctx context.Context, id int) (*model.AppRole, error) {
	query := `
		SELECT
			app_role_id,
			app_role_type_id,
			app_role_name,
			app_role_description,
			is_active,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at
		FROM app_role
		WHERE app_role_id = $1 AND deleted_at IS NULL
	`

	var m model.AppRole

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&m.AppRoleID,
		&m.AppRoleTypeID,
		&m.AppRoleName,
		&m.AppRoleDescription,
		&m.IsActive,
		&m.CreatedBy,
		&m.UpdatedBy,
		&m.DeletedBy,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &m, nil
}

// ==================================================
// List (Cursor Pagination + Filters)
// ==================================================
func (r *AppRoleRepository) List(ctx context.Context,
	cursorValue interface{}, cursorKey *int, limit int, filters map[string]string,
	orderBy, sort string) ([]model.AppRole, error) {
	tableKey := "app_role_id"
	baseQuery := `
		SELECT
			app_role_id,
			app_role_type_id,
			app_role_name,
			app_role_description,
			is_active,
			created_by,
			updated_by,
			deleted_by,
			created_at,
			updated_at,
			deleted_at
		FROM app_role
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	argPos := 1

	// cursor
	cursorClause, cursorArgs, argPos := helper.BuildCursorCondition(orderBy, tableKey, sort, cursorValue, cursorKey, argPos)
	baseQuery += cursorClause
	args = append(args, cursorArgs...)

	// LIKE filters
	likeFilters := map[string]string{
		"app_role_name":        "app_role_name ILIKE $%d",
		"app_role_description": "app_role_description ILIKE $%d",
	}

	for key, clause := range likeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, "%"+v+"%")
			argPos++
		}
	}

	// exact match filters
	exactFilters := []string{"app_role_type_id", "is_active", "created_by", "updated_by"}

	for _, key := range exactFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += fmt.Sprintf(" AND %s = $%d", key, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// time filters (>= <=)
	timeFilters := map[string]string{
		"created_at_after":  "created_at >= $%d",
		"created_at_before": "created_at <= $%d",
		"updated_at_after":  "updated_at >= $%d",
		"updated_at_before": "updated_at <= $%d",
	}

	for key, clause := range timeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// final order + limit
	baseQuery += fmt.Sprintf(" ORDER BY %s %s, %s %s LIMIT $%d", orderBy, sort, tableKey, sort, argPos)
	args = append(args, limit)

	rows, err := r.DB.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.AppRole

	for rows.Next() {
		var m model.AppRole
		if err := rows.Scan(
			&m.AppRoleID,
			&m.AppRoleTypeID,
			&m.AppRoleName,
			&m.AppRoleDescription,
			&m.IsActive,
			&m.CreatedBy,
			&m.UpdatedBy,
			&m.DeletedBy,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.DeletedAt,
		); err != nil {
			return nil, err
		}
		list = append(list, m)
	}

	return list, rows.Err()
}
