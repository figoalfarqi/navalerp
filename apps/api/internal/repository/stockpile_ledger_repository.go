package repository

import (
	"context"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockpileLedgerRepository struct{ DB *pgxpool.Pool }

func NewStockpileLedgerRepository(db *pgxpool.Pool) *StockpileLedgerRepository {
	return &StockpileLedgerRepository{DB: db}
}

const stockpileLedgerJSON = `
	to_jsonb(l) || jsonb_build_object(
		'project',to_jsonb(p),
		'stockpile_cargo',to_jsonb(sc) || jsonb_build_object(
			'stockpile',to_jsonb(s),'cargo_type',to_jsonb(ct)
		),
		'project_transport',CASE WHEN pt.project_transport_id IS NULL THEN NULL ELSE to_jsonb(pt) END,
		'adjustment',CASE WHEN adjustment.stockpile_adjustment_id IS NULL THEN NULL ELSE to_jsonb(adjustment) END
	)`

func (r *StockpileLedgerRepository) GetByID(ctx context.Context, id int) (*model.StockpileLedger, error) {
	return scanJSONRow[model.StockpileLedger](r.DB.QueryRow(ctx, `
		SELECT `+stockpileLedgerJSON+`
		FROM stockpile_ledger l
		JOIN project p ON p.project_id=l.project_id
		JOIN stockpile_cargo sc ON sc.stockpile_cargo_id=l.stockpile_cargo_id
		JOIN stockpile s ON s.stockpile_id=sc.stockpile_id
		JOIN cargo_type ct ON ct.cargo_type_id=sc.cargo_type_id
		LEFT JOIN project_transport pt ON pt.project_transport_id=l.project_transport_id
		LEFT JOIN stockpile_adjustment adjustment ON adjustment.stockpile_adjustment_id=l.adjustment_id AND adjustment.deleted_at IS NULL
		WHERE l.stockpile_ledger_id=$1 AND l.deleted_at IS NULL`, id))
}

func (r *StockpileLedgerRepository) List(ctx context.Context, opts model.ListOptions) ([]model.StockpileLedger, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+stockpileLedgerJSON+`
		FROM stockpile_ledger l
		JOIN project p ON p.project_id=l.project_id
		JOIN stockpile_cargo sc ON sc.stockpile_cargo_id=l.stockpile_cargo_id
		JOIN stockpile s ON s.stockpile_id=sc.stockpile_id
		JOIN cargo_type ct ON ct.cargo_type_id=sc.cargo_type_id
		LEFT JOIN project_transport pt ON pt.project_transport_id=l.project_transport_id
		LEFT JOIN stockpile_adjustment adjustment ON adjustment.stockpile_adjustment_id=l.adjustment_id AND adjustment.deleted_at IS NULL
		WHERE l.deleted_at IS NULL
		  AND ($1::int IS NULL OR l.stockpile_ledger_id < $1)
		  AND ($2::int IS NULL OR l.project_id=$2)
		  AND ($3::int IS NULL OR l.stockpile_cargo_id=$3)
		  AND ($4::smallint IS NULL OR l.reference_type_id=$4)
		  AND ($5::timestamptz IS NULL OR l.created_at >= $5)
		  AND ($6::timestamptz IS NULL OR l.created_at < $6)
		ORDER BY l.stockpile_ledger_id DESC LIMIT $7 OFFSET $8`,
		opts.CursorKey, opts.ProjectID, positiveFilter(opts.Filters["stockpile_cargo_id"]),
		positiveFilter(opts.Filters["reference_type_id"]), opts.DateFrom, opts.DateTo,
		opts.Limit, opts.Offset)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.StockpileLedger](rows)
}
