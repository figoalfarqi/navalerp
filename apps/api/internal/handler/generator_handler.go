package handler

import (
	"net/http"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/pkg/response"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EntityNumberConfig struct {
	Table  string
	Column string
	Prefix string
}

var entityNumberConfigs = map[string]EntityNumberConfig{
	"requisition":       {"proc_requisitions", "requisition_number", "PR"},
	"purchase_order":    {"proc_purchase_orders", "po_number", "PO"},
	"goods_receipt":     {"proc_goods_receipts", "receipt_number", "BAPHP"},
	"stock_transfer":    {"inv_stock_transfers", "transfer_number", "TRF"},
	"stock_adjustment":  {"inv_stock_adjustments", "adjustment_number", "ADJ"},
	"work_order":        {"mro_work_orders", "work_order_number", "WO"},
	"journal_entry":     {"fin_journal_entries", "entry_number", "JRN"},
	"shipment":          {"log_shipments", "manifest_number", "MNF"},
	"tender":            {"proc_tenders", "tender_number", "TND"},
	"contract":          {"proc_contracts", "contract_number", "KTR"},
	"invoice":           {"fin_invoices", "invoice_number", "INV"},
	"payment":           {"fin_payments", "payment_reference_no", "SP2D"},
	"budget_commitment": {"fin_budget_commitments", "commitment_number", "KOM"},
	"failure_report":    {"mro_failure_reports", "report_number", "REP"},
	"mission":           {"ops_missions", "mission_code", "MSN"},
	"document":          {"doc_documents", "document_number", "DOC"},
	"daily_log":         {"ops_daily_logs", "report_number", "LOG"},
	"readiness_report":  {"ops_readiness_reports", "report_number", "RR"},
	"readiness_alert":   {"ops_readiness_alerts", "alert_code", "ALT"},
}

type GeneratorHandler struct {
	DB *pgxpool.Pool
}

func NewGeneratorHandler(db *pgxpool.Pool) *GeneratorHandler {
	return &GeneratorHandler{DB: db}
}

func (h *GeneratorHandler) GenerateNumber(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	entity := strings.ToLower(strings.TrimSpace(q.Get("entity")))
	prefix := strings.ToUpper(strings.TrimSpace(q.Get("prefix")))

	cfg, ok := entityNumberConfigs[entity]
	if !ok {
		// Fallback: check if prefix matches any known entity
		for _, v := range entityNumberConfigs {
			if v.Prefix == prefix {
				cfg = v
				ok = true
				break
			}
		}
	}

	if !ok {
		// If custom table/column passed, allow if safe
		table := strings.ToLower(strings.TrimSpace(q.Get("table")))
		column := strings.ToLower(strings.TrimSpace(q.Get("column")))
		if table != "" && column != "" && prefix != "" {
			cfg = EntityNumberConfig{
				Table:  table,
				Column: column,
				Prefix: prefix,
			}
			ok = true
		}
	}

	if !ok {
		response.JSON(w, http.StatusBadRequest, "Invalid entity or prefix for number generation", nil, nil)
		return
	}

	nextNum, err := helper.GenerateNextNumber(r.Context(), h.DB, cfg.Table, cfg.Column, cfg.Prefix)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.JSON(w, http.StatusOK, "Number generated successfully", map[string]interface{}{
		"number": nextNum,
		"prefix": cfg.Prefix,
		"entity": entity,
	}, nil)
}
