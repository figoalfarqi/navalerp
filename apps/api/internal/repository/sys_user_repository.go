package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SysUserRepository struct {
	DB *pgxpool.Pool
}

func NewSysUserRepository(db *pgxpool.Pool) *SysUserRepository {
	return &SysUserRepository{DB: db}
}

// Get retrieves a single sys_user by user_id
func (r *SysUserRepository) Get(ctx context.Context, id string) (*model.SysUser, error) {
	query := `SELECT user_id, unit_id, username, password_hash, full_name, email, phone, military_id, rank_title, department, role, is_active, last_login_at, failed_login_attempts, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at, auth_version FROM sys_users WHERE user_id = $1 AND deleted_at IS NULL`

	var m model.SysUser
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.UserId, &m.UnitId, &m.Username, &m.PasswordHash, &m.FullName, &m.Email, &m.Phone, &m.MilitaryId, &m.RankTitle, &m.Department, &m.Role, &m.IsActive, &m.LastLoginAt, &m.FailedLoginAttempts, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt, &m.AuthVersion)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated sys_user records
func (r *SysUserRepository) List(ctx context.Context, opts model.ListOptions) ([]model.SysUser, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(username ILIKE $%[1]d OR password_hash ILIKE $%[1]d OR full_name ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM sys_users WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT user_id, unit_id, username, password_hash, full_name, email, phone, military_id, rank_title, department, role, is_active, last_login_at, failed_login_attempts, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at, auth_version FROM sys_users WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.SysUser
	for rows.Next() {
		var m model.SysUser
		if err := rows.Scan(&m.UserId, &m.UnitId, &m.Username, &m.PasswordHash, &m.FullName, &m.Email, &m.Phone, &m.MilitaryId, &m.RankTitle, &m.Department, &m.Role, &m.IsActive, &m.LastLoginAt, &m.FailedLoginAttempts, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt, &m.AuthVersion); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new sys_user with optional child items
func (r *SysUserRepository) Create(ctx context.Context, m *model.SysUser) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO sys_users (unit_id, username, password_hash, full_name, email, phone, military_id, rank_title, department, role, is_active, last_login_at, failed_login_attempts, created_by, updated_by, deleted_by, auth_version) VALUES ($1, $2, $3, $4, $5, $6, $7, $8::military_rank_type, $9::department_type, $10::user_role_type, $11, $12, $13, $14, $15, $16, $17) RETURNING user_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.UnitId, m.Username, m.PasswordHash, m.FullName, m.Email, m.Phone, m.MilitaryId, m.RankTitle, m.Department, m.Role, m.IsActive, m.LastLoginAt, m.FailedLoginAttempts, m.CreatedBy, m.UpdatedBy, m.DeletedBy, m.AuthVersion).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing sys_user
func (r *SysUserRepository) Update(ctx context.Context, id string, m *model.SysUser) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE sys_users SET unit_id = $1, username = $2, password_hash = $3, full_name = $4, email = $5, phone = $6, military_id = $7, rank_title = $8::military_rank_type, department = $9::department_type, role = $10::user_role_type, is_active = $11, last_login_at = $12, failed_login_attempts = $13, updated_by = $14, updated_at = CURRENT_TIMESTAMP, auth_version = $15 WHERE user_id = $16`
	_, err = tx.Exec(ctx, updateQuery, m.UnitId, m.Username, m.PasswordHash, m.FullName, m.Email, m.Phone, m.MilitaryId, m.RankTitle, m.Department, m.Role, m.IsActive, m.LastLoginAt, m.FailedLoginAttempts, m.UpdatedBy, m.AuthVersion, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes sys_user
func (r *SysUserRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE sys_users SET deleted_at = CURRENT_TIMESTAMP WHERE user_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
