package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuditLogRepository struct {
	DB *pgxpool.Pool
}

func NewAuditLogRepository(db *pgxpool.Pool) *AuditLogRepository {
	return &AuditLogRepository{DB: db}
}

// Get retrieves a single audit_log by log_id
func (r *AuditLogRepository) Get(ctx context.Context, id string) (*model.AuditLog, error) {
	query := `SELECT t.log_id, t.user_id, COALESCE(j_usr.full_name, ''), t.action, t.entity_table, t.entity_id, t.old_values, t.new_values, t.ip_address, t.user_agent, t.created_at
	FROM sys_audit_logs t
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.user_id
	WHERE t.log_id = $1`

	var m model.AuditLog
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.LogId, &m.UserId, &m.UserName, &m.Action, &m.EntityTable, &m.EntityId, &m.OldValues, &m.NewValues, &m.IpAddress, &m.UserAgent, &m.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// List retrieves paginated audit_log records
func (r *AuditLogRepository) List(ctx context.Context, opts model.ListOptions) ([]model.AuditLog, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(entity_table ILIKE $%[1]d OR entity_id ILIKE $%[1]d OR ip_address ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM sys_audit_logs t WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf(`SELECT t.log_id, t.user_id, COALESCE(j_usr.full_name, ''), t.action, t.entity_table, t.entity_id, t.old_values, t.new_values, t.ip_address, t.user_agent, t.created_at
	FROM sys_audit_logs t
	LEFT JOIN sys_users j_usr ON j_usr.user_id = t.user_id
	WHERE %s ORDER BY t.created_at DESC LIMIT $%d OFFSET $%d`, whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.AuditLog
	for rows.Next() {
		var m model.AuditLog
		if err := rows.Scan(&m.LogId, &m.UserId, &m.UserName, &m.Action, &m.EntityTable, &m.EntityId, &m.OldValues, &m.NewValues, &m.IpAddress, &m.UserAgent, &m.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new audit_log with optional child items
func (r *AuditLogRepository) Create(ctx context.Context, m *model.AuditLog) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO sys_audit_logs (user_id, action, entity_table, entity_id, old_values, new_values, ip_address, user_agent) VALUES ($1, $2::audit_action_type, $3, $4, $5, $6, $7, $8) RETURNING log_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.UserId, m.Action, m.EntityTable, m.EntityId, m.OldValues, m.NewValues, m.IpAddress, m.UserAgent).Scan(&newID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing audit_log
func (r *AuditLogRepository) Update(ctx context.Context, id string, m *model.AuditLog) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE sys_audit_logs SET user_id = $1, action = $2::audit_action_type, entity_table = $3, entity_id = $4, old_values = $5, new_values = $6, ip_address = $7, user_agent = $8 WHERE log_id = $9`
	_, err = tx.Exec(ctx, updateQuery, m.UserId, m.Action, m.EntityTable, m.EntityId, m.OldValues, m.NewValues, m.IpAddress, m.UserAgent, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes audit_log
func (r *AuditLogRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM sys_audit_logs t WHERE log_id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
