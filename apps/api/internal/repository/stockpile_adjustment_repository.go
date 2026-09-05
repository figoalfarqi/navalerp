package repository

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockpileAdjustmentRepository struct{ DB *pgxpool.Pool }

func NewStockpileAdjustmentRepository(db *pgxpool.Pool) *StockpileAdjustmentRepository {
	return &StockpileAdjustmentRepository{DB: db}
}

type cargoBalance struct {
	volume, weight float64
	capacityVolume *float64
	capacityWeight *float64
}

func lockCargoBalance(ctx context.Context, tx pgx.Tx, id int) (*cargoBalance, error) {
	var row cargoBalance
	err := tx.QueryRow(ctx, `
		SELECT current_volume_cubic,current_weight_ton,capacity_volume_cubic,capacity_weight_ton
		FROM stockpile_cargo
		WHERE stockpile_cargo_id=$1 AND deleted_at IS NULL AND is_active=1
		FOR UPDATE`, id).Scan(&row.volume, &row.weight, &row.capacityVolume, &row.capacityWeight)
	return &row, err
}

func validateCargoBalance(row *cargoBalance, volume, weight float64) error {
	if volume < 0 || weight < 0 {
		return errors.New("stockpile balance cannot be negative")
	}
	if row.capacityVolume != nil && volume > *row.capacityVolume {
		return errors.New("stockpile volume exceeds capacity")
	}
	if row.capacityWeight != nil && weight > *row.capacityWeight {
		return errors.New("stockpile weight exceeds capacity")
	}
	return nil
}

func updateCargoBalance(ctx context.Context, tx pgx.Tx, id int, volume, weight float64, userID int) error {
	_, err := tx.Exec(ctx, `
		UPDATE stockpile_cargo
		SET current_volume_cubic=$1,current_weight_ton=$2,updated_by=$3,updated_at=CURRENT_TIMESTAMP
		WHERE stockpile_cargo_id=$4 AND deleted_at IS NULL`, volume, weight, userID, id)
	return err
}

func insertStockpileLedger(
	ctx context.Context,
	tx pgx.Tx,
	req *model.StockpileAdjustmentRequest,
	adjustmentID int,
	amountVolume, amountWeight, balanceVolume, balanceWeight float64,
	stockpileCargoID, userID int,
) error {
	_, err := tx.Exec(ctx, `
		INSERT INTO stockpile_ledger (
			project_id,stockpile_cargo_id,project_transport_id,adjustment_id,reference_type_id,
			amount_volume_cubic,balance_volume_cubic,amount_weight_ton,balance_weight_ton,
			created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$10)`,
		req.ProjectID, stockpileCargoID, req.ProjectTransportID, adjustmentID, req.ReferenceTypeID,
		amountVolume, balanceVolume, amountWeight, balanceWeight, userID)
	return err
}

func (r *StockpileAdjustmentRepository) Create(ctx context.Context, req *model.StockpileAdjustmentRequest, userID int) (int, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	balance, err := lockCargoBalance(ctx, tx, req.StockpileCargoID)
	if err != nil {
		return 0, err
	}
	newVolume := balance.volume + req.AmountVolumeCubic
	newWeight := balance.weight + req.AmountWeightTon
	if err := validateCargoBalance(balance, newVolume, newWeight); err != nil {
		return 0, err
	}
	var id int
	err = tx.QueryRow(ctx, `
		INSERT INTO stockpile_adjustment (
			project_id,stockpile_cargo_id,adjustment_date,project_transport_id,
			reference_type_id,amount_volume_cubic,amount_weight_ton,reason,created_by,updated_by
		) VALUES ($1,$2,$3::date,$4,$5,$6,$7,$8,$9,$9)
		RETURNING stockpile_adjustment_id`,
		req.ProjectID, req.StockpileCargoID, req.AdjustmentDate, req.ProjectTransportID,
		req.ReferenceTypeID, req.AmountVolumeCubic, req.AmountWeightTon, req.Reason, userID).Scan(&id)
	if err != nil {
		return 0, err
	}
	if err := updateCargoBalance(ctx, tx, req.StockpileCargoID, newVolume, newWeight, userID); err != nil {
		return 0, err
	}
	if err := insertStockpileLedger(ctx, tx, req, id, req.AmountVolumeCubic, req.AmountWeightTon, newVolume, newWeight, req.StockpileCargoID, userID); err != nil {
		return 0, err
	}
	return id, tx.Commit(ctx)
}

func (r *StockpileAdjustmentRepository) Update(ctx context.Context, id int, req *model.StockpileAdjustmentRequest, userID int) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var old model.StockpileAdjustmentRequest
	err = tx.QueryRow(ctx, `
		SELECT project_id,stockpile_cargo_id,adjustment_date::text,project_transport_id,
			reference_type_id,amount_volume_cubic,amount_weight_ton,reason
		FROM stockpile_adjustment
		WHERE stockpile_adjustment_id=$1 AND deleted_at IS NULL FOR UPDATE`, id).Scan(
		&old.ProjectID, &old.StockpileCargoID, &old.AdjustmentDate, &old.ProjectTransportID,
		&old.ReferenceTypeID, &old.AmountVolumeCubic, &old.AmountWeightTon, &old.Reason)
	if err != nil {
		return err
	}

	ids := []int{old.StockpileCargoID}
	if req.StockpileCargoID != old.StockpileCargoID {
		ids = append(ids, req.StockpileCargoID)
	}
	sort.Ints(ids)
	balances := map[int]*cargoBalance{}
	for _, cargoID := range ids {
		balance, lockErr := lockCargoBalance(ctx, tx, cargoID)
		if lockErr != nil {
			return lockErr
		}
		balances[cargoID] = balance
	}

	oldBalance := balances[old.StockpileCargoID]
	newBalance := balances[req.StockpileCargoID]
	oldVolume, oldWeight := oldBalance.volume, oldBalance.weight
	newVolume, newWeight := newBalance.volume, newBalance.weight
	if req.StockpileCargoID == old.StockpileCargoID {
		newVolume += req.AmountVolumeCubic - old.AmountVolumeCubic
		newWeight += req.AmountWeightTon - old.AmountWeightTon
		if err := validateCargoBalance(newBalance, newVolume, newWeight); err != nil {
			return err
		}
	} else {
		oldVolume -= old.AmountVolumeCubic
		oldWeight -= old.AmountWeightTon
		newVolume += req.AmountVolumeCubic
		newWeight += req.AmountWeightTon
		if err := validateCargoBalance(oldBalance, oldVolume, oldWeight); err != nil {
			return err
		}
		if err := validateCargoBalance(newBalance, newVolume, newWeight); err != nil {
			return err
		}
	}

	_, err = tx.Exec(ctx, `
		UPDATE stockpile_adjustment SET project_id=$1,stockpile_cargo_id=$2,adjustment_date=$3::date,
			project_transport_id=$4,reference_type_id=$5,amount_volume_cubic=$6,
			amount_weight_ton=$7,reason=$8,updated_by=$9,updated_at=CURRENT_TIMESTAMP
		WHERE stockpile_adjustment_id=$10 AND deleted_at IS NULL`,
		req.ProjectID, req.StockpileCargoID, req.AdjustmentDate, req.ProjectTransportID,
		req.ReferenceTypeID, req.AmountVolumeCubic, req.AmountWeightTon, req.Reason, userID, id)
	if err != nil {
		return err
	}
	if req.StockpileCargoID == old.StockpileCargoID {
		if err := updateCargoBalance(ctx, tx, req.StockpileCargoID, newVolume, newWeight, userID); err != nil {
			return err
		}
		deltaVolume := req.AmountVolumeCubic - old.AmountVolumeCubic
		deltaWeight := req.AmountWeightTon - old.AmountWeightTon
		if deltaVolume != 0 || deltaWeight != 0 {
			if err := insertStockpileLedger(ctx, tx, req, id,
				deltaVolume, deltaWeight,
				newVolume, newWeight, req.StockpileCargoID, userID); err != nil {
				return err
			}
		}
	} else {
		if err := updateCargoBalance(ctx, tx, old.StockpileCargoID, oldVolume, oldWeight, userID); err != nil {
			return err
		}
		if err := updateCargoBalance(ctx, tx, req.StockpileCargoID, newVolume, newWeight, userID); err != nil {
			return err
		}
		reversal := *req
		reversal.ProjectID = old.ProjectID
		reversal.ProjectTransportID = old.ProjectTransportID
		reversal.ReferenceTypeID = old.ReferenceTypeID
		if err := insertStockpileLedger(ctx, tx, &reversal, id, -old.AmountVolumeCubic, -old.AmountWeightTon, oldVolume, oldWeight, old.StockpileCargoID, userID); err != nil {
			return err
		}
		if err := insertStockpileLedger(ctx, tx, req, id, req.AmountVolumeCubic, req.AmountWeightTon, newVolume, newWeight, req.StockpileCargoID, userID); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

func (r *StockpileAdjustmentRepository) SoftDelete(ctx context.Context, userID, id int) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	var req model.StockpileAdjustmentRequest
	err = tx.QueryRow(ctx, `
		SELECT project_id,stockpile_cargo_id,adjustment_date::text,project_transport_id,
			reference_type_id,amount_volume_cubic,amount_weight_ton,reason
		FROM stockpile_adjustment
		WHERE stockpile_adjustment_id=$1 AND deleted_at IS NULL FOR UPDATE`, id).Scan(
		&req.ProjectID, &req.StockpileCargoID, &req.AdjustmentDate, &req.ProjectTransportID,
		&req.ReferenceTypeID, &req.AmountVolumeCubic, &req.AmountWeightTon, &req.Reason)
	if err != nil {
		return err
	}
	balance, err := lockCargoBalance(ctx, tx, req.StockpileCargoID)
	if err != nil {
		return err
	}
	newVolume, newWeight := balance.volume-req.AmountVolumeCubic, balance.weight-req.AmountWeightTon
	if err := validateCargoBalance(balance, newVolume, newWeight); err != nil {
		return err
	}
	if err := updateCargoBalance(ctx, tx, req.StockpileCargoID, newVolume, newWeight, userID); err != nil {
		return err
	}
	_, err = tx.Exec(ctx, `
		UPDATE stockpile_adjustment SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP,
			updated_by=$1,updated_at=CURRENT_TIMESTAMP
		WHERE stockpile_adjustment_id=$2 AND deleted_at IS NULL`, userID, id)
	if err != nil {
		return err
	}
	if err := insertStockpileLedger(ctx, tx, &req, id, -req.AmountVolumeCubic, -req.AmountWeightTon, newVolume, newWeight, req.StockpileCargoID, userID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

const stockpileAdjustmentJSON = `
	to_jsonb(a) || jsonb_build_object(
		'project',to_jsonb(p),
		'stockpile_cargo',to_jsonb(sc) || jsonb_build_object(
			'stockpile',to_jsonb(s),'cargo_type',to_jsonb(ct)
		),
		'project_transport',CASE WHEN pt.project_transport_id IS NULL THEN NULL ELSE to_jsonb(pt) END
	)`

func (r *StockpileAdjustmentRepository) GetByID(ctx context.Context, id int) (*model.StockpileAdjustment, error) {
	return scanJSONRow[model.StockpileAdjustment](r.DB.QueryRow(ctx, `
		SELECT `+stockpileAdjustmentJSON+`
		FROM stockpile_adjustment a
		JOIN project p ON p.project_id=a.project_id
		JOIN stockpile_cargo sc ON sc.stockpile_cargo_id=a.stockpile_cargo_id
		JOIN stockpile s ON s.stockpile_id=sc.stockpile_id
		JOIN cargo_type ct ON ct.cargo_type_id=sc.cargo_type_id
		LEFT JOIN project_transport pt ON pt.project_transport_id=a.project_transport_id
		WHERE a.stockpile_adjustment_id=$1 AND a.deleted_at IS NULL`, id))
}

func (r *StockpileAdjustmentRepository) List(ctx context.Context, opts model.ListOptions) ([]model.StockpileAdjustment, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+stockpileAdjustmentJSON+`
		FROM stockpile_adjustment a
		JOIN project p ON p.project_id=a.project_id
		JOIN stockpile_cargo sc ON sc.stockpile_cargo_id=a.stockpile_cargo_id
		JOIN stockpile s ON s.stockpile_id=sc.stockpile_id
		JOIN cargo_type ct ON ct.cargo_type_id=sc.cargo_type_id
		LEFT JOIN project_transport pt ON pt.project_transport_id=a.project_transport_id
		WHERE a.deleted_at IS NULL
		  AND ($1::int IS NULL OR a.stockpile_adjustment_id < $1)
		  AND ($2::int IS NULL OR a.project_id=$2)
		  AND ($3::int IS NULL OR a.stockpile_cargo_id=$3)
		  AND ($4::smallint IS NULL OR a.reference_type_id=$4)
		  AND ($5::date IS NULL OR a.adjustment_date >= $5::date)
		  AND ($6::date IS NULL OR a.adjustment_date < $6::date)
		ORDER BY a.stockpile_adjustment_id DESC LIMIT $7 OFFSET $8`,
		opts.CursorKey, opts.ProjectID, positiveFilter(opts.Filters["stockpile_cargo_id"]),
		positiveFilter(opts.Filters["reference_type_id"]), opts.DateFrom, opts.DateTo,
		opts.Limit, opts.Offset)
	if err != nil {
		return nil, fmt.Errorf("list stockpile adjustments: %w", err)
	}
	return scanJSONRows[model.StockpileAdjustment](rows)
}
