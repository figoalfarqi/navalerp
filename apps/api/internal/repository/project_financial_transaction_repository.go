package repository

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectFinancialTransactionRepository struct{ DB *pgxpool.Pool }

const projectFinancialTransactionJSON = `
	to_jsonb(f) || jsonb_build_object(
		'project',to_jsonb(p),
		'project_transport',CASE
			WHEN pt.project_transport_id IS NULL THEN NULL
			ELSE to_jsonb(pt) || jsonb_build_object('project',to_jsonb(p))
		END,
		'client',CASE
			WHEN client.client_id IS NULL THEN NULL
			ELSE to_jsonb(client) || jsonb_build_object(
				'city',CASE WHEN client_city.city_id IS NULL THEN NULL ELSE to_jsonb(client_city) END
			)
		END,
		'vendor',CASE
			WHEN vendor.vendor_id IS NULL THEN NULL
			ELSE to_jsonb(vendor) || jsonb_build_object(
				'vendor_type',CASE WHEN vt.vendor_type_id IS NULL THEN NULL ELSE to_jsonb(vt) END,
				'city',CASE WHEN vendor_city.city_id IS NULL THEN NULL ELSE to_jsonb(vendor_city) END
			)
		END
	)`

func NewProjectFinancialTransactionRepository(db *pgxpool.Pool) *ProjectFinancialTransactionRepository {
	return &ProjectFinancialTransactionRepository{DB: db}
}

func money(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func (r *ProjectFinancialTransactionRepository) Create(ctx context.Context, req *model.ProjectFinancialTransactionRequest, userID int) (int, error) {
	posted := 0
	if req.IsPosted != nil {
		posted = *req.IsPosted
	}
	var id int
	err := r.DB.QueryRow(ctx, `
		INSERT INTO project_financial_transaction (
			project_id,project_transport_id,transaction_number,transaction_kind,
			transaction_date,due_date,paid_at,client_id,vendor_id,reference_number,
			material_sale_income,transport_service_income,other_income,
			material_purchase_expense,transport_expense,road_money_expense,
			loading_expense,unloading_expense,fuel_expense,toll_expense,other_expense,
			transaction_note,is_posted,created_by,updated_by
		) VALUES (
			$1,$2,$3,$4,$5::date,$6::date,$7,$8,$9,$10,$11,$12,$13,$14,$15,
			$16,$17,$18,$19,$20,$21,$22,$23,$24,$24
		) RETURNING project_financial_transaction_id`,
		req.ProjectID, req.ProjectTransportID, req.TransactionNumber, req.TransactionKind,
		req.TransactionDate, req.DueDate, req.PaidAt, req.ClientID, req.VendorID,
		req.ReferenceNumber, money(req.MaterialSaleIncome), money(req.TransportServiceIncome),
		money(req.OtherIncome), money(req.MaterialPurchaseExpense), money(req.TransportExpense),
		money(req.RoadMoneyExpense), money(req.LoadingExpense), money(req.UnloadingExpense),
		money(req.FuelExpense), money(req.TollExpense), money(req.OtherExpense),
		req.TransactionNote, posted, userID,
	).Scan(&id)
	return id, err
}

func (r *ProjectFinancialTransactionRepository) Update(ctx context.Context, id int, req *model.ProjectFinancialTransactionRequest, userID int) error {
	posted := 0
	if req.IsPosted != nil {
		posted = *req.IsPosted
	}
	_, err := r.DB.Exec(ctx, `
		UPDATE project_financial_transaction SET
			project_id=$1,project_transport_id=$2,transaction_number=$3,transaction_kind=$4,
			transaction_date=$5::date,due_date=$6::date,paid_at=$7,client_id=$8,vendor_id=$9,
			reference_number=$10,material_sale_income=$11,transport_service_income=$12,
			other_income=$13,material_purchase_expense=$14,transport_expense=$15,
			road_money_expense=$16,loading_expense=$17,unloading_expense=$18,
			fuel_expense=$19,toll_expense=$20,other_expense=$21,transaction_note=$22,
			is_posted=$23,updated_by=$24,updated_at=CURRENT_TIMESTAMP
		WHERE project_financial_transaction_id=$25 AND deleted_at IS NULL`,
		req.ProjectID, req.ProjectTransportID, req.TransactionNumber, req.TransactionKind,
		req.TransactionDate, req.DueDate, req.PaidAt, req.ClientID, req.VendorID,
		req.ReferenceNumber, money(req.MaterialSaleIncome), money(req.TransportServiceIncome),
		money(req.OtherIncome), money(req.MaterialPurchaseExpense), money(req.TransportExpense),
		money(req.RoadMoneyExpense), money(req.LoadingExpense), money(req.UnloadingExpense),
		money(req.FuelExpense), money(req.TollExpense), money(req.OtherExpense),
		req.TransactionNote, posted, userID, id)
	return err
}

func (r *ProjectFinancialTransactionRepository) SoftDelete(ctx context.Context, id, userID int) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE project_financial_transaction
		SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP,updated_by=$1,updated_at=CURRENT_TIMESTAMP
		WHERE project_financial_transaction_id=$2 AND deleted_at IS NULL`, userID, id)
	return err
}

func (r *ProjectFinancialTransactionRepository) GetByID(ctx context.Context, id int) (*model.ProjectFinancialTransaction, error) {
	return scanJSONRow[model.ProjectFinancialTransaction](r.DB.QueryRow(ctx, `
		SELECT `+projectFinancialTransactionJSON+`
		FROM project_financial_transaction f
		JOIN project p ON p.project_id=f.project_id AND p.deleted_at IS NULL
		LEFT JOIN project_transport pt ON pt.project_transport_id=f.project_transport_id AND pt.deleted_at IS NULL
		LEFT JOIN client client ON client.client_id=f.client_id AND client.deleted_at IS NULL
		LEFT JOIN city client_city ON client_city.city_id=client.city_id AND client_city.deleted_at IS NULL
		LEFT JOIN vendor vendor ON vendor.vendor_id=f.vendor_id AND vendor.deleted_at IS NULL
		LEFT JOIN vendor_type vt ON vt.vendor_type_id=vendor.vendor_type_id AND vt.deleted_at IS NULL
		LEFT JOIN city vendor_city ON vendor_city.city_id=vendor.city_id AND vendor_city.deleted_at IS NULL
		WHERE f.project_financial_transaction_id=$1 AND f.deleted_at IS NULL`, id))
}

func (r *ProjectFinancialTransactionRepository) List(ctx context.Context, opts model.ListOptions) ([]model.ProjectFinancialTransaction, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+projectFinancialTransactionJSON+`
		FROM project_financial_transaction f
		JOIN project p ON p.project_id=f.project_id AND p.deleted_at IS NULL
		LEFT JOIN project_transport pt ON pt.project_transport_id=f.project_transport_id AND pt.deleted_at IS NULL
		LEFT JOIN client client ON client.client_id=f.client_id AND client.deleted_at IS NULL
		LEFT JOIN city client_city ON client_city.city_id=client.city_id AND client_city.deleted_at IS NULL
		LEFT JOIN vendor vendor ON vendor.vendor_id=f.vendor_id AND vendor.deleted_at IS NULL
		LEFT JOIN vendor_type vt ON vt.vendor_type_id=vendor.vendor_type_id AND vt.deleted_at IS NULL
		LEFT JOIN city vendor_city ON vendor_city.city_id=vendor.city_id AND vendor_city.deleted_at IS NULL
		WHERE f.deleted_at IS NULL
		  AND ($1::int IS NULL OR f.project_financial_transaction_id < $1)
		  AND ($2::int IS NULL OR f.project_id=$2)
		  AND ($3::date IS NULL OR f.transaction_date >= $3::date)
		  AND ($4::date IS NULL OR f.transaction_date < $4::date)
		  AND ($7='' OR f.transaction_kind=$7)
		  AND ($8::smallint IS NULL OR f.is_posted=$8)
		  AND ($9='' OR f.transaction_number ILIKE '%' || $9 || '%')
		ORDER BY f.project_financial_transaction_id DESC LIMIT $5 OFFSET $6`,
		opts.CursorKey, opts.ProjectID, opts.DateFrom, opts.DateTo, opts.Limit, opts.Offset,
		opts.TransactionKind, opts.IsPosted, opts.Filters["transaction_number"])
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.ProjectFinancialTransaction](rows)
}
