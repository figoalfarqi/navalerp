package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApprovalService struct {
	DB *pgxpool.Pool
}

func NewApprovalService(db *pgxpool.Pool) *ApprovalService {
	return &ApprovalService{DB: db}
}

func (s *ApprovalService) ProcessApproval(ctx context.Context, loginID string, req *model.ApprovalRequest) (*model.ApprovalResponse, error) {
	if req.EntityType == "" || req.EntityID == "" || req.Action == "" {
		return nil, errors.New("entity_type, entity_id, and action are required")
	}

	action := strings.ToUpper(req.Action)
	if action == "RESOLVE" {
		action = "APPROVE"
	} else if action == "DISMISS" {
		action = "REJECT"
	}
	if action != "APPROVE" && action != "REJECT" {
		return nil, errors.New("action must be APPROVE or REJECT")
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var targetTable string
	var newStatus string
	var oldStatus string

	switch req.EntityType {
	case "requisition", "proc_requisitions":
		targetTable = "proc_requisitions"
		if action == "APPROVE" {
			newStatus = "APPROVED"
		} else {
			newStatus = "REJECTED"
		}
		_ = tx.QueryRow(ctx, "SELECT COALESCE(approval_status::text, '') FROM proc_requisitions WHERE requisition_id = $1", req.EntityID).Scan(&oldStatus)
		query := `UPDATE proc_requisitions 
		          SET approval_status = $1, approved_by_user_id = $2, approved_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP 
		          WHERE requisition_id = $3`
		tag, err := tx.Exec(ctx, query, newStatus, loginID, req.EntityID)
		if err != nil {
			return nil, fmt.Errorf("failed to update requisition: %w", err)
		}
		if tag.RowsAffected() == 0 {
			return nil, errors.New("requisition not found")
		}

	case "purchase_order", "proc_purchase_orders":
		targetTable = "proc_purchase_orders"
		if action == "APPROVE" {
			newStatus = "APPROVED"
		} else {
			newStatus = "REJECTED"
		}
		_ = tx.QueryRow(ctx, "SELECT COALESCE(status::text, '') FROM proc_purchase_orders WHERE po_id = $1", req.EntityID).Scan(&oldStatus)
		query := `UPDATE proc_purchase_orders 
		          SET status = $1, updated_at = CURRENT_TIMESTAMP, updated_by = $2 
		          WHERE po_id = $3`
		tag, err := tx.Exec(ctx, query, newStatus, loginID, req.EntityID)
		if err != nil {
			return nil, fmt.Errorf("failed to update purchase order: %w", err)
		}
		if tag.RowsAffected() == 0 {
			return nil, errors.New("purchase order not found")
		}

	case "work_order", "mro_work_orders":
		targetTable = "mro_work_orders"
		if action == "APPROVE" {
			newStatus = "APPROVED"
		} else {
			newStatus = "CANCELLED"
		}
		_ = tx.QueryRow(ctx, "SELECT COALESCE(status::text, '') FROM mro_work_orders WHERE work_order_id = $1", req.EntityID).Scan(&oldStatus)
		query := `UPDATE mro_work_orders 
		          SET status = $1, updated_at = CURRENT_TIMESTAMP 
		          WHERE work_order_id = $2`
		tag, err := tx.Exec(ctx, query, newStatus, req.EntityID)
		if err != nil {
			return nil, fmt.Errorf("failed to update work order: %w", err)
		}
		if tag.RowsAffected() == 0 {
			return nil, errors.New("work order not found")
		}

	case "goods_receipt", "proc_goods_receipts":
		targetTable = "proc_goods_receipts"
		passed := (action == "APPROVE")
		if passed {
			newStatus = "VERIFIED"
		} else {
			newStatus = "REJECTED"
		}
		_ = tx.QueryRow(ctx, "SELECT CASE WHEN inspection_passed THEN 'PASSED' ELSE 'FAILED' END FROM proc_goods_receipts WHERE receipt_id = $1", req.EntityID).Scan(&oldStatus)
		query := `UPDATE proc_goods_receipts 
		          SET inspection_passed = $1, updated_at = CURRENT_TIMESTAMP 
		          WHERE receipt_id = $2`
		tag, err := tx.Exec(ctx, query, passed, req.EntityID)
		if err != nil {
			return nil, fmt.Errorf("failed to update goods receipt: %w", err)
		}
		if tag.RowsAffected() == 0 {
			return nil, errors.New("goods receipt not found")
		}

	case "payment", "fin_payments":
		targetTable = "fin_payments"
		if action == "APPROVE" {
			newStatus = "COMPLETED"
		} else {
			newStatus = "CANCELLED"
		}
		_ = tx.QueryRow(ctx, "SELECT COALESCE(payment_status::text, '') FROM fin_payments WHERE payment_id = $1", req.EntityID).Scan(&oldStatus)
		query := `UPDATE fin_payments 
		          SET payment_status = $1, updated_at = CURRENT_TIMESTAMP 
		          WHERE payment_id = $2`
		tag, err := tx.Exec(ctx, query, newStatus, req.EntityID)
		if err != nil {
			return nil, fmt.Errorf("failed to update payment: %w", err)
		}
		if tag.RowsAffected() == 0 {
			return nil, errors.New("payment not found")
		}

	case "cui_alert", "cui_alerts":
		targetTable = "cui_alerts"
		if action == "APPROVE" {
			newStatus = "RESOLVED"
		} else {
			newStatus = "FALSE_ALARM"
		}
		_ = tx.QueryRow(ctx, "SELECT COALESCE(status::text, '') FROM cui_alerts WHERE alert_id = $1", req.EntityID).Scan(&oldStatus)
		query := `UPDATE cui_alerts 
		          SET status = $1, updated_at = CURRENT_TIMESTAMP, updated_by = $2, resolved_at = CURRENT_TIMESTAMP 
		          WHERE alert_id = $3`
		tag, err := tx.Exec(ctx, query, newStatus, loginID, req.EntityID)
		if err != nil {
			return nil, fmt.Errorf("failed to update cui alert: %w", err)
		}
		if tag.RowsAffected() == 0 {
			return nil, errors.New("cui alert not found")
		}

	default:
		return nil, fmt.Errorf("unsupported entity type: %s", req.EntityType)
	}

	// Generate digital cryptographic verification stamp
	now := time.Now()
	rawSign := fmt.Sprintf("NAVALERP|%s|%s|%s|%s|%s|%d", req.EntityType, req.EntityID, action, loginID, newStatus, now.Unix())
	hash := sha256.Sum256([]byte(rawSign))
	stamp := fmt.Sprintf("TNI-AL-SIG-%s-%s", now.Format("20060102"), strings.ToUpper(hex.EncodeToString(hash[:])[:12]))

	// Record in sys_audit_logs
	auditNewValues, _ := json.Marshal(map[string]any{
		"action":            action,
		"new_status":        newStatus,
		"notes":             req.Notes,
		"signature_stamp":   stamp,
		"digital_signature": req.DigitalSignature,
		"timestamp":         now,
	})
	auditOldValues, _ := json.Marshal(map[string]any{
		"old_status": oldStatus,
	})

	var userUUID *string
	if len(loginID) == 36 {
		userUUID = &loginID
	}

	_, auditErr := tx.Exec(ctx, `
		INSERT INTO sys_audit_logs (user_id, action, entity_table, entity_id, old_values, new_values, created_at)
		VALUES ($1, 'APPROVAL'::audit_action_type, $2, $3, $4::jsonb, $5::jsonb, CURRENT_TIMESTAMP)
	`, userUUID, targetTable, req.EntityID, string(auditOldValues), string(auditNewValues))
	if auditErr != nil {
		fmt.Printf("Audit log insert warning: %v\n", auditErr)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("transaction commit failed: %w", err)
	}

	return &model.ApprovalResponse{
		Success:        true,
		EntityType:     req.EntityType,
		EntityID:       req.EntityID,
		NewStatus:      newStatus,
		ApprovedBy:     loginID,
		ApprovedAt:     now,
		SignatureStamp: stamp,
		Notes:          req.Notes,
	}, nil
}

func (s *ApprovalService) GetPendingApprovals(ctx context.Context) ([]model.PendingApprovalItem, error) {
	items := make([]model.PendingApprovalItem, 0)

	// 1. Requisitions pending
	rows, err := s.DB.Query(ctx, `
		SELECT requisition_id::text, requisition_number, COALESCE(justification, 'Pengadaan Bebekal/Material'), 
		       'Staf Perbekalan', requested_date, total_estimated_cost::float8, priority::text, approval_status::text
		FROM proc_requisitions 
		WHERE approval_status = 'PENDING' AND deleted_at IS NULL
		ORDER BY requested_date DESC LIMIT 10
	`)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var it model.PendingApprovalItem
			it.EntityType = "requisition"
			if scanErr := rows.Scan(&it.EntityID, &it.ReferenceNo, &it.Title, &it.RequestedBy, &it.RequestedDate, &it.Amount, &it.Priority, &it.Status); scanErr == nil {
				items = append(items, it)
			}
		}
	}

	// 2. Work Orders pending
	woRows, err := s.DB.Query(ctx, `
		SELECT work_order_id::text, work_order_number, 'Perintah Kerja Pemeliharaan KRI', 
		       COALESCE(assigned_facility, 'Fasharkan Surabaya'), scheduled_start_date, estimated_cost::float8, priority::text, status::text
		FROM mro_work_orders 
		WHERE status = 'DRAFT' AND deleted_at IS NULL
		ORDER BY scheduled_start_date DESC LIMIT 10
	`)
	if err == nil {
		defer woRows.Close()
		for woRows.Next() {
			var it model.PendingApprovalItem
			it.EntityType = "work_order"
			if scanErr := woRows.Scan(&it.EntityID, &it.ReferenceNo, &it.Title, &it.RequestedBy, &it.RequestedDate, &it.Amount, &it.Priority, &it.Status); scanErr == nil {
				items = append(items, it)
			}
		}
	}

	// 3. CUI Alerts active / investigating
	alertRows, err := s.DB.Query(ctx, `
		SELECT alert_id::text, alert_code, alert_type::text, 
		       'Pusat Komando Bawah Laut', detected_at, 0.0::float8, severity::text, status::text
		FROM cui_alerts 
		WHERE status IN ('ACTIVE', 'INVESTIGATING', 'DISPATCHED') AND deleted_at IS NULL
		ORDER BY detected_at DESC LIMIT 10
	`)
	if err == nil {
		defer alertRows.Close()
		for alertRows.Next() {
			var it model.PendingApprovalItem
			it.EntityType = "cui_alert"
			if scanErr := alertRows.Scan(&it.EntityID, &it.ReferenceNo, &it.Title, &it.RequestedBy, &it.RequestedDate, &it.Amount, &it.Priority, &it.Status); scanErr == nil {
				items = append(items, it)
			}
		}
	}

	return items, nil
}
