package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProvinceRepository struct {
	DB *pgxpool.Pool
}

func NewProvinceRepository(db *pgxpool.Pool) *ProvinceRepository {
	return &ProvinceRepository{DB: db}
}

// ==================================================
// Create Province
// ==================================================
func (r *ProvinceRepository) Create(ctx context.Context, p *model.Province) (int, error) {
	query := `
        INSERT INTO province (
            province_name,
            province_real_name,
            is_active,
            created_by,
            updated_by,
            created_at,
            updated_at
        ) VALUES ($1,$2,$3,$4,$5,$6,$7)
        RETURNING province_id
    `

	var id int
	err := r.DB.QueryRow(ctx, query,
		p.ProvinceName,
		p.ProvinceRealName,
		p.IsActive,
		p.CreatedBy,
		p.UpdatedBy,
		p.CreatedAt,
		p.UpdatedAt,
	).Scan(&id)

	return id, err
}

// ==================================================
// Update Province
// ==================================================
func (r *ProvinceRepository) Update(ctx context.Context, id int, p *model.Province) error {

	// conditional is_active update
	setIsActive := ""
	args := []interface{}{
		p.ProvinceName,     // $1
		p.ProvinceRealName, // $2
	}
	argPos := len(args) + 1

	if p.IsActive != -1 { // pola sama dengan client
		setIsActive = fmt.Sprintf(", is_active = $%d", argPos)
		args = append(args, p.IsActive)
		argPos++
	}

	args = append(args,
		p.UpdatedBy, // $argPos
		p.UpdatedAt, // $argPos+1
		id,          // $argPos+2
	)

	query := fmt.Sprintf(`
        UPDATE province
        SET province_name = $1,
		 	province_real_name = $2
            %s,
            updated_by = $%d,
            updated_at = $%d
        WHERE province_id = $%d AND deleted_at IS NULL
    `,
		setIsActive,
		argPos,   // updated_by
		argPos+1, // updated_at
		argPos+2, // id
	)

	_, err := r.DB.Exec(ctx, query, args...)
	return err
}

// ==================================================
// Soft Delete Province
// ==================================================
func (r *ProvinceRepository) SoftDelete(ctx context.Context, deletedBy, id int) error {

	query := `
        UPDATE province
        SET deleted_at = $1,
            deleted_by = $2,
            updated_at = $1,
            updated_by = $2,
            is_active = 0
        WHERE province_id = $3 AND deleted_at IS NULL
    `

	_, err := r.DB.Exec(ctx, query, time.Now(), deletedBy, id)
	return err
}

// ==================================================
// Get By ID
// ==================================================
func (r *ProvinceRepository) GetByID(ctx context.Context, id int) (*model.Province, error) {
	query := `
        SELECT
            province_id,
            province_name,
            province_real_name,
            is_active,
            created_by,
            updated_by,
            deleted_by,
            created_at,
            updated_at,
            deleted_at
        FROM province
        WHERE province_id = $1 AND deleted_at IS NULL
    `

	var p model.Province

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&p.ProvinceID,
		&p.ProvinceName,
		&p.ProvinceRealName,
		&p.IsActive,
		&p.CreatedBy,
		&p.UpdatedBy,
		&p.DeletedBy,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &p, nil
}

// ==================================================
// List Province (cursor pagination + filters)
// ==================================================
func (r *ProvinceRepository) List(ctx context.Context,
	cursorValue interface{}, cursorKey *int, limit int, filters map[string]string,
	orderBy, sort string) ([]model.Province, error) {
	tableKey := "province_id"
	parsedOrderBy := orderBy
	baseQuery := `
        SELECT
            province_id,
            province_name,
            province_real_name,
            is_active,
            created_by,
            updated_by,
            deleted_by,
            created_at,
            updated_at,
            deleted_at
        FROM province
        WHERE deleted_at IS NULL
    `

	args := []interface{}{}
	argPos := 1

	cursorClause, cursorArgs, argPos := helper.BuildCursorCondition(parsedOrderBy, tableKey, sort, cursorValue, cursorKey, argPos)
	baseQuery += cursorClause
	args = append(args, cursorArgs...)

	likeFilters := map[string]string{
		"province_name":      "province_name ILIKE $%d",
		"province_real_name": "province_real_name ILIKE $%d",
	}

	for key, clause := range likeFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += " AND " + fmt.Sprintf(clause, argPos)
			args = append(args, "%"+v+"%")
			argPos++
		}
	}

	// Exact (include is_active)
	exactFilters := []string{
		"is_active",
		"created_by",
		"updated_by",
	}
	for _, key := range exactFilters {
		if v, ok := filters[key]; ok && v != "" {
			baseQuery += fmt.Sprintf(" AND %s = $%d", key, argPos)
			args = append(args, v)
			argPos++
		}
	}

	// Time range
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

	// Order + Limit
	baseQuery += fmt.Sprintf(" ORDER BY %s %s, %s %s LIMIT $%d", parsedOrderBy, sort, tableKey, sort, argPos)
	args = append(args, limit)
	rows, err := r.DB.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []model.Province

	for rows.Next() {
		var p model.Province
		if err := rows.Scan(
			&p.ProvinceID,
			&p.ProvinceName,
			&p.ProvinceRealName,
			&p.IsActive,
			&p.CreatedBy,
			&p.UpdatedBy,
			&p.DeletedBy,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.DeletedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, p)
	}

	return out, rows.Err()
}
