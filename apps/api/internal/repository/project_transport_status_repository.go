package repository

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectTransportStatusRepository struct{ DB *pgxpool.Pool }

const (
	automaticStockpileReceiptReason = "Automatic stockpile receipt from completed transport unloading"
	automaticStockpileReleaseReason = "Automatic stockpile release from loaded outgoing transport"
)

type transportStockKey struct {
	transportID int
	statusType  int
}

type automaticStockMovement struct {
	adjustmentID     int
	projectID        int
	stockpileCargoID int
	transportID      int
	referenceType    int
	volume           float64
	weight           float64
	statusTime       time.Time
	reason           string
}

type automaticStockMovementChange struct {
	old *automaticStockMovement
	new *automaticStockMovement
}

func NewProjectTransportStatusRepository(db *pgxpool.Pool) *ProjectTransportStatusRepository {
	return &ProjectTransportStatusRepository{DB: db}
}

func (r *ProjectTransportStatusRepository) HasStatus(ctx context.Context, transportID, statusType int) (bool, error) {
	return r.HasStatusExcept(ctx, transportID, statusType, nil)
}

func (r *ProjectTransportStatusRepository) HasStatusExcept(
	ctx context.Context,
	transportID, statusType int,
	excludedStatusID *int,
) (bool, error) {
	var exists bool
	err := r.DB.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM project_transport_status
			WHERE project_transport_id=$1 AND project_transport_status_type_id=$2
			  AND deleted_at IS NULL AND is_active=1
			  AND ($3::int IS NULL OR project_transport_status_id<>$3)
		)`, transportID, statusType, excludedStatusID).Scan(&exists)
	return exists, err
}

func (r *ProjectTransportStatusRepository) Create(ctx context.Context, req *model.ProjectTransportStatusRequest, userID int) (int, error) {
	statusTime := time.Now()
	if req.StatusTime != nil {
		statusTime = *req.StatusTime
	}
	fraud, active := 0, 1
	if req.IsFraud != nil {
		fraud = *req.IsFraud
	}
	if req.IsActive != nil {
		active = *req.IsActive
	}
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)
	if err := lockTransportsForStatus(ctx, tx, req.ProjectTransportID); err != nil {
		return 0, err
	}
	before, err := loadAutomaticStockMovements(ctx, tx, []int{req.ProjectTransportID})
	if err != nil {
		return 0, err
	}
	var id int
	err = tx.QueryRow(ctx, `
		INSERT INTO project_transport_status (
			project_transport_id,project_transport_status_type_id,status_time,
			cargo_box_length,cargo_box_width,cargo_box_height,cargo_volume_cubic,
			cargo_weight_ton,project_transport_status_note,is_fraud,is_active,created_by,updated_by
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$12)
		RETURNING project_transport_status_id`,
		req.ProjectTransportID, req.ProjectTransportStatusTypeID, statusTime,
		req.CargoBoxLength, req.CargoBoxWidth, req.CargoBoxHeight,
		req.CargoVolumeCubic, req.CargoWeightTon, req.ProjectTransportStatusNote,
		fraud, active, userID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	if err := enforceDestinationFraud(ctx, tx, req.ProjectTransportID, userID); err != nil {
		return 0, err
	}
	if err := syncTransportSnapshot(ctx, tx, req.ProjectTransportID, userID); err != nil {
		return 0, err
	}
	after, err := desiredAutomaticStockMovements(ctx, tx, []int{req.ProjectTransportID})
	if err != nil {
		return 0, err
	}
	if err := reconcileAutomaticStockMovements(ctx, tx, before, after, userID); err != nil {
		return 0, err
	}
	return id, tx.Commit(ctx)
}

func automaticStockMovementIdentity(statusType int) (int, string, bool) {
	switch statusType {
	case 4:
		return 2, automaticStockpileReleaseReason, true
	case 7:
		return 1, automaticStockpileReceiptReason, true
	default:
		return 0, "", false
	}
}

func uniqueTransportIDs(ids ...int) []int {
	seen := make(map[int]struct{}, len(ids))
	result := make([]int, 0, len(ids))
	for _, id := range ids {
		if id <= 0 {
			continue
		}
		if _, exists := seen[id]; exists {
			continue
		}
		seen[id] = struct{}{}
		result = append(result, id)
	}
	sort.Ints(result)
	return result
}

func lockTransportsForStatus(ctx context.Context, tx pgx.Tx, ids ...int) error {
	for _, transportID := range uniqueTransportIDs(ids...) {
		var lockedID int
		err := tx.QueryRow(ctx, `
			SELECT project_transport_id
			FROM project_transport
			WHERE project_transport_id=$1 AND deleted_at IS NULL
			FOR UPDATE`, transportID).Scan(&lockedID)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadAutomaticStockMovements(
	ctx context.Context,
	tx pgx.Tx,
	transportIDs []int,
) (map[transportStockKey]*automaticStockMovement, error) {
	result := make(map[transportStockKey]*automaticStockMovement)
	for _, transportID := range uniqueTransportIDs(transportIDs...) {
		for _, statusType := range []int{4, 7} {
			referenceType, reason, _ := automaticStockMovementIdentity(statusType)
			var movement automaticStockMovement
			err := tx.QueryRow(ctx, `
				SELECT stockpile_adjustment_id,project_id,stockpile_cargo_id,
					project_transport_id,reference_type_id,amount_volume_cubic,
					amount_weight_ton,adjustment_date::timestamptz,reason
				FROM stockpile_adjustment
				WHERE project_transport_id=$1 AND reference_type_id=$2
				  AND reason=$3 AND deleted_at IS NULL
				ORDER BY stockpile_adjustment_id DESC
				LIMIT 1
				FOR UPDATE`,
				transportID, referenceType, reason,
			).Scan(
				&movement.adjustmentID,
				&movement.projectID,
				&movement.stockpileCargoID,
				&movement.transportID,
				&movement.referenceType,
				&movement.volume,
				&movement.weight,
				&movement.statusTime,
				&movement.reason,
			)
			if errors.Is(err, pgx.ErrNoRows) {
				continue
			}
			if err != nil {
				return nil, err
			}
			result[transportStockKey{transportID: transportID, statusType: statusType}] = &movement
		}
	}
	return result, nil
}

func desiredAutomaticStockMovements(
	ctx context.Context,
	tx pgx.Tx,
	transportIDs []int,
) (map[transportStockKey]*automaticStockMovement, error) {
	result := make(map[transportStockKey]*automaticStockMovement)
	for _, transportID := range uniqueTransportIDs(transportIDs...) {
		for _, statusType := range []int{4, 7} {
			referenceType, reason, _ := automaticStockMovementIdentity(statusType)
			var (
				projectID        int
				stockpileCargoID *int
				routeType        string
				statusTime       time.Time
				volume           float64
				weight           float64
			)
			err := tx.QueryRow(ctx, `
				SELECT pt.project_id,p.stockpile_cargo_id,pr.route_type,s.status_time,
					CASE WHEN $2=7
						THEN COALESCE(pt.delivered_volume_cubic,pt.loaded_volume_cubic,0)
						ELSE COALESCE(pt.loaded_volume_cubic,0)
					END,
					CASE WHEN $2=7
						THEN COALESCE(pt.delivered_weight_ton,pt.loaded_weight_ton,0)
						ELSE COALESCE(pt.loaded_weight_ton,0)
					END
				FROM project_transport_status s
				JOIN project_transport pt ON pt.project_transport_id=s.project_transport_id
				JOIN project p ON p.project_id=pt.project_id
				JOIN project_route pr ON pr.project_route_id=pt.project_route_id
				WHERE s.project_transport_id=$1
				  AND s.project_transport_status_type_id=$2
				  AND s.is_active=1 AND s.deleted_at IS NULL
				  AND pt.deleted_at IS NULL`,
				transportID, statusType,
			).Scan(&projectID, &stockpileCargoID, &routeType, &statusTime, &volume, &weight)
			if errors.Is(err, pgx.ErrNoRows) {
				continue
			}
			if err != nil {
				return nil, err
			}
			if stockpileCargoID == nil || (volume == 0 && weight == 0) {
				continue
			}
			sign := 1.0
			switch {
			case routeType == "SOURCE_TO_STOCKPILE" && statusType == 7:
			case routeType == "STOCKPILE_TO_CLIENT" && statusType == 4:
				sign = -1
			default:
				continue
			}
			result[transportStockKey{transportID: transportID, statusType: statusType}] = &automaticStockMovement{
				projectID:        projectID,
				stockpileCargoID: *stockpileCargoID,
				transportID:      transportID,
				referenceType:    referenceType,
				volume:           sign * volume,
				weight:           sign * weight,
				statusTime:       statusTime,
				reason:           reason,
			}
		}
	}
	return result, nil
}

func automaticStockMovementsEqual(left, right *automaticStockMovement) bool {
	if left == nil || right == nil {
		return left == right
	}
	leftDate := left.statusTime.Format("2006-01-02")
	rightDate := right.statusTime.Format("2006-01-02")
	return left.projectID == right.projectID &&
		left.stockpileCargoID == right.stockpileCargoID &&
		left.transportID == right.transportID &&
		left.referenceType == right.referenceType &&
		left.volume == right.volume &&
		left.weight == right.weight &&
		leftDate == rightDate &&
		left.reason == right.reason
}

func reconcileAutomaticStockMovements(
	ctx context.Context,
	tx pgx.Tx,
	before map[transportStockKey]*automaticStockMovement,
	after map[transportStockKey]*automaticStockMovement,
	userID int,
) error {
	keySet := make(map[transportStockKey]struct{}, len(before)+len(after))
	for key := range before {
		keySet[key] = struct{}{}
	}
	for key := range after {
		keySet[key] = struct{}{}
	}
	keys := make([]transportStockKey, 0, len(keySet))
	for key := range keySet {
		if automaticStockMovementsEqual(before[key], after[key]) {
			continue
		}
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].transportID == keys[j].transportID {
			return keys[i].statusType < keys[j].statusType
		}
		return keys[i].transportID < keys[j].transportID
	})
	if len(keys) == 0 {
		return nil
	}

	cargoIDSet := make(map[int]struct{})
	for _, key := range keys {
		if movement := before[key]; movement != nil {
			cargoIDSet[movement.stockpileCargoID] = struct{}{}
		}
		if movement := after[key]; movement != nil {
			cargoIDSet[movement.stockpileCargoID] = struct{}{}
		}
	}
	cargoIDs := make([]int, 0, len(cargoIDSet))
	for cargoID := range cargoIDSet {
		cargoIDs = append(cargoIDs, cargoID)
	}
	sort.Ints(cargoIDs)
	balances := make(map[int]*cargoBalance, len(cargoIDs))
	finalVolume := make(map[int]float64, len(cargoIDs))
	finalWeight := make(map[int]float64, len(cargoIDs))
	for _, cargoID := range cargoIDs {
		balance, err := lockCargoBalance(ctx, tx, cargoID)
		if err != nil {
			return err
		}
		balances[cargoID] = balance
		finalVolume[cargoID] = balance.volume
		finalWeight[cargoID] = balance.weight
	}
	for _, key := range keys {
		if movement := before[key]; movement != nil {
			finalVolume[movement.stockpileCargoID] -= movement.volume
			finalWeight[movement.stockpileCargoID] -= movement.weight
		}
		if movement := after[key]; movement != nil {
			finalVolume[movement.stockpileCargoID] += movement.volume
			finalWeight[movement.stockpileCargoID] += movement.weight
		}
	}
	for _, cargoID := range cargoIDs {
		if err := validateCargoBalance(
			balances[cargoID],
			finalVolume[cargoID],
			finalWeight[cargoID],
		); err != nil {
			return err
		}
	}

	runningVolume := make(map[int]float64, len(cargoIDs))
	runningWeight := make(map[int]float64, len(cargoIDs))
	for _, cargoID := range cargoIDs {
		runningVolume[cargoID] = balances[cargoID].volume
		runningWeight[cargoID] = balances[cargoID].weight
	}
	for _, key := range keys {
		if movement := before[key]; movement != nil {
			runningVolume[movement.stockpileCargoID] -= movement.volume
			runningWeight[movement.stockpileCargoID] -= movement.weight
			_, err := tx.Exec(ctx, `
				UPDATE stockpile_adjustment
				SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP,
					updated_by=$1,updated_at=CURRENT_TIMESTAMP
				WHERE stockpile_adjustment_id=$2 AND deleted_at IS NULL`,
				userID, movement.adjustmentID)
			if err != nil {
				return err
			}
			_, err = tx.Exec(ctx, `
				INSERT INTO stockpile_ledger (
					project_id,stockpile_cargo_id,project_transport_id,adjustment_id,
					reference_type_id,amount_volume_cubic,balance_volume_cubic,
					amount_weight_ton,balance_weight_ton,created_by,updated_by
				) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$10)`,
				movement.projectID,
				movement.stockpileCargoID,
				movement.transportID,
				movement.adjustmentID,
				movement.referenceType,
				-movement.volume,
				runningVolume[movement.stockpileCargoID],
				-movement.weight,
				runningWeight[movement.stockpileCargoID],
				userID,
			)
			if err != nil {
				return err
			}
		}
		if movement := after[key]; movement != nil {
			runningVolume[movement.stockpileCargoID] += movement.volume
			runningWeight[movement.stockpileCargoID] += movement.weight
			err := tx.QueryRow(ctx, `
				INSERT INTO stockpile_adjustment (
					project_id,stockpile_cargo_id,adjustment_date,project_transport_id,
					reference_type_id,amount_volume_cubic,amount_weight_ton,reason,
					created_by,updated_by
				) VALUES ($1,$2,$3::date,$4,$5,$6,$7,$8,$9,$9)
				RETURNING stockpile_adjustment_id`,
				movement.projectID,
				movement.stockpileCargoID,
				movement.statusTime,
				movement.transportID,
				movement.referenceType,
				movement.volume,
				movement.weight,
				movement.reason,
				userID,
			).Scan(&movement.adjustmentID)
			if err != nil {
				return err
			}
			_, err = tx.Exec(ctx, `
				INSERT INTO stockpile_ledger (
					project_id,stockpile_cargo_id,project_transport_id,adjustment_id,
					reference_type_id,amount_volume_cubic,balance_volume_cubic,
					amount_weight_ton,balance_weight_ton,created_by,updated_by
				) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$10)`,
				movement.projectID,
				movement.stockpileCargoID,
				movement.transportID,
				movement.adjustmentID,
				movement.referenceType,
				movement.volume,
				runningVolume[movement.stockpileCargoID],
				movement.weight,
				runningWeight[movement.stockpileCargoID],
				userID,
			)
			if err != nil {
				return err
			}
		}
	}
	for _, cargoID := range cargoIDs {
		if err := updateCargoBalance(
			ctx,
			tx,
			cargoID,
			finalVolume[cargoID],
			finalWeight[cargoID],
			userID,
		); err != nil {
			return err
		}
	}
	return nil
}

func syncTransportSnapshot(ctx context.Context, tx pgx.Tx, transportID, userID int) error {
	command, err := tx.Exec(ctx, `
		WITH measurements AS (
			SELECT
				(
					SELECT s.cargo_volume_cubic
					FROM project_transport_status s
					WHERE s.project_transport_id=$1
					  AND s.project_transport_status_type_id BETWEEN 1 AND 4
					  AND s.cargo_volume_cubic IS NOT NULL
					  AND s.is_active=1 AND s.deleted_at IS NULL
					ORDER BY s.status_time DESC,s.project_transport_status_id DESC
					LIMIT 1
				) AS loaded_volume,
				(
					SELECT s.cargo_weight_ton
					FROM project_transport_status s
					WHERE s.project_transport_id=$1
					  AND s.project_transport_status_type_id BETWEEN 1 AND 4
					  AND s.cargo_weight_ton IS NOT NULL
					  AND s.is_active=1 AND s.deleted_at IS NULL
					ORDER BY s.status_time DESC,s.project_transport_status_id DESC
					LIMIT 1
				) AS loaded_weight,
				(
					SELECT s.cargo_volume_cubic
					FROM project_transport_status s
					WHERE s.project_transport_id=$1
					  AND s.project_transport_status_type_id BETWEEN 5 AND 8
					  AND s.cargo_volume_cubic IS NOT NULL
					  AND s.is_active=1 AND s.deleted_at IS NULL
					ORDER BY s.status_time DESC,s.project_transport_status_id DESC
					LIMIT 1
				) AS delivered_volume,
				(
					SELECT s.cargo_weight_ton
					FROM project_transport_status s
					WHERE s.project_transport_id=$1
					  AND s.project_transport_status_type_id BETWEEN 5 AND 8
					  AND s.cargo_weight_ton IS NOT NULL
					  AND s.is_active=1 AND s.deleted_at IS NULL
					ORDER BY s.status_time DESC,s.project_transport_status_id DESC
					LIMIT 1
				) AS delivered_weight,
				(
					SELECT s.cargo_volume_cubic
					FROM project_transport_status s
					WHERE s.project_transport_id=$1
					  AND s.cargo_volume_cubic IS NOT NULL
					  AND s.is_active=1 AND s.deleted_at IS NULL
					ORDER BY s.status_time DESC,s.project_transport_status_id DESC
					LIMIT 1
				) AS latest_volume,
				(
					SELECT s.cargo_weight_ton
					FROM project_transport_status s
					WHERE s.project_transport_id=$1
					  AND s.cargo_weight_ton IS NOT NULL
					  AND s.is_active=1 AND s.deleted_at IS NULL
					ORDER BY s.status_time DESC,s.project_transport_status_id DESC
					LIMIT 1
				) AS latest_weight,
				EXISTS (
					SELECT 1
					FROM project_transport_status s
					WHERE s.project_transport_id=$1
					  AND s.project_transport_status_type_id=8
					  AND s.is_active=1 AND s.deleted_at IS NULL
				) AS completed
		)
		UPDATE project_transport pt
		SET loaded_volume_cubic=m.loaded_volume,
			loaded_weight_ton=m.loaded_weight,
			delivered_volume_cubic=m.delivered_volume,
			delivered_weight_ton=m.delivered_weight,
			purchase_volume_cubic=CASE
				WHEN pr.route_type IN ('SOURCE_TO_CLIENT','SOURCE_TO_STOCKPILE')
				THEN m.loaded_volume
				ELSE NULL
			END,
			purchase_weight_ton=CASE
				WHEN pr.route_type IN ('SOURCE_TO_CLIENT','SOURCE_TO_STOCKPILE')
				THEN m.loaded_weight
				ELSE NULL
			END,
			sale_volume_cubic=CASE
				WHEN pr.route_type IN ('SOURCE_TO_CLIENT','STOCKPILE_TO_CLIENT')
				THEN m.delivered_volume
				ELSE NULL
			END,
			sale_weight_ton=CASE
				WHEN pr.route_type IN ('SOURCE_TO_CLIENT','STOCKPILE_TO_CLIENT')
				THEN m.delivered_weight
				ELSE NULL
			END,
			transport_service_volume_cubic=m.latest_volume,
			transport_service_weight_ton=m.latest_weight,
			transport_cost_volume_cubic=m.latest_volume,
			transport_cost_weight_ton=m.latest_weight,
			is_completed=CASE WHEN m.completed THEN 1 ELSE 0 END,
			updated_by=$2,
			updated_at=CURRENT_TIMESTAMP
		FROM project_route pr
		CROSS JOIN measurements m
		WHERE pt.project_transport_id=$1
		  AND pr.project_route_id=pt.project_route_id
		  AND pt.deleted_at IS NULL`,
		transportID, userID)
	if err != nil {
		return err
	}
	if command.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func enforceDestinationFraud(ctx context.Context, tx pgx.Tx, transportID, userID int) error {
	_, err := tx.Exec(ctx, `
		UPDATE project_transport_status destination
		SET is_fraud=1,
			project_transport_status_note=CASE
				WHEN BTRIM(COALESCE(destination.project_transport_status_note,''))=''
				THEN 'Arrived at destination without an origin arrival status'
				ELSE destination.project_transport_status_note
			END,
			updated_by=$2,
			updated_at=CURRENT_TIMESTAMP
		WHERE destination.project_transport_id=$1
		  AND destination.project_transport_status_type_id=5
		  AND destination.is_active=1
		  AND destination.deleted_at IS NULL
		  AND NOT EXISTS (
			SELECT 1
			FROM project_transport_status origin
			WHERE origin.project_transport_id=destination.project_transport_id
			  AND origin.project_transport_status_type_id=1
			  AND origin.is_active=1
			  AND origin.deleted_at IS NULL
		  )`,
		transportID, userID)
	return err
}

func (r *ProjectTransportStatusRepository) Update(ctx context.Context, id int, req *model.ProjectTransportStatusRequest, userID int) error {
	statusTime := time.Now()
	if req.StatusTime != nil {
		statusTime = *req.StatusTime
	}
	fraud, active := 0, 1
	if req.IsFraud != nil {
		fraud = *req.IsFraud
	}
	if req.IsActive != nil {
		active = *req.IsActive
	}
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	var oldTransportID int
	err = tx.QueryRow(ctx, `
		SELECT project_transport_id
		FROM project_transport_status
		WHERE project_transport_status_id=$1 AND deleted_at IS NULL`,
		id).Scan(&oldTransportID)
	if err != nil {
		return err
	}
	transportIDs := uniqueTransportIDs(oldTransportID, req.ProjectTransportID)
	if err := lockTransportsForStatus(ctx, tx, transportIDs...); err != nil {
		return err
	}
	var lockedTransportID int
	err = tx.QueryRow(ctx, `
		SELECT project_transport_id
		FROM project_transport_status
		WHERE project_transport_status_id=$1 AND deleted_at IS NULL
		FOR UPDATE`, id).Scan(&lockedTransportID)
	if err != nil {
		return err
	}
	if lockedTransportID != oldTransportID {
		oldTransportID = lockedTransportID
		transportIDs = uniqueTransportIDs(oldTransportID, req.ProjectTransportID)
		if err := lockTransportsForStatus(ctx, tx, transportIDs...); err != nil {
			return err
		}
	}
	before, err := loadAutomaticStockMovements(ctx, tx, transportIDs)
	if err != nil {
		return err
	}
	command, err := tx.Exec(ctx, `
		UPDATE project_transport_status SET
			project_transport_id=$1,project_transport_status_type_id=$2,status_time=$3,
			cargo_box_length=$4,cargo_box_width=$5,cargo_box_height=$6,cargo_volume_cubic=$7,
			cargo_weight_ton=$8,project_transport_status_note=$9,is_fraud=$10,is_active=$11,
			updated_by=$12,updated_at=CURRENT_TIMESTAMP
		WHERE project_transport_status_id=$13 AND deleted_at IS NULL`,
		req.ProjectTransportID, req.ProjectTransportStatusTypeID, statusTime,
		req.CargoBoxLength, req.CargoBoxWidth, req.CargoBoxHeight,
		req.CargoVolumeCubic, req.CargoWeightTon, req.ProjectTransportStatusNote,
		fraud, active, userID, id)
	if err != nil {
		return err
	}
	if command.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	for _, transportID := range transportIDs {
		if err := enforceDestinationFraud(ctx, tx, transportID, userID); err != nil {
			return err
		}
		if err := syncTransportSnapshot(ctx, tx, transportID, userID); err != nil {
			return err
		}
	}
	after, err := desiredAutomaticStockMovements(ctx, tx, transportIDs)
	if err != nil {
		return err
	}
	if err := reconcileAutomaticStockMovements(ctx, tx, before, after, userID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (r *ProjectTransportStatusRepository) SoftDelete(ctx context.Context, id, userID int) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	var transportID int
	err = tx.QueryRow(ctx, `
		SELECT project_transport_id
		FROM project_transport_status
		WHERE project_transport_status_id=$1 AND deleted_at IS NULL`,
		id).Scan(&transportID)
	if err != nil {
		return err
	}
	if err := lockTransportsForStatus(ctx, tx, transportID); err != nil {
		return err
	}
	var lockedTransportID int
	err = tx.QueryRow(ctx, `
		SELECT project_transport_id
		FROM project_transport_status
		WHERE project_transport_status_id=$1 AND deleted_at IS NULL
		FOR UPDATE`, id).Scan(&lockedTransportID)
	if err != nil {
		return err
	}
	if lockedTransportID != transportID {
		transportID = lockedTransportID
		if err := lockTransportsForStatus(ctx, tx, transportID); err != nil {
			return err
		}
	}
	before, err := loadAutomaticStockMovements(ctx, tx, []int{transportID})
	if err != nil {
		return err
	}
	command, err := tx.Exec(ctx, `
		UPDATE project_transport_status
		SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP,updated_by=$1,updated_at=CURRENT_TIMESTAMP
		WHERE project_transport_status_id=$2 AND deleted_at IS NULL`, userID, id)
	if err != nil {
		return err
	}
	if command.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	_, err = tx.Exec(ctx, `
		UPDATE project_transport_photo
		SET deleted_by=$1,deleted_at=CURRENT_TIMESTAMP,
			updated_by=$1,updated_at=CURRENT_TIMESTAMP
		WHERE project_transport_status_id=$2 AND deleted_at IS NULL`,
		userID, id)
	if err != nil {
		return err
	}
	if err := enforceDestinationFraud(ctx, tx, transportID, userID); err != nil {
		return err
	}
	if err := syncTransportSnapshot(ctx, tx, transportID, userID); err != nil {
		return err
	}
	after, err := desiredAutomaticStockMovements(ctx, tx, []int{transportID})
	if err != nil {
		return err
	}
	if err := reconcileAutomaticStockMovements(ctx, tx, before, after, userID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

const projectTransportStatusJSON = `
	to_jsonb(s) || jsonb_build_object(
		'project_transport',` + projectTransportReferenceJSON + `
	)`

func (r *ProjectTransportStatusRepository) GetByID(ctx context.Context, id int, checkerID *int) (*model.ProjectTransportStatus, error) {
	return scanJSONRow[model.ProjectTransportStatus](r.DB.QueryRow(ctx, `
		SELECT `+projectTransportStatusJSON+`
		FROM project_transport_status s
		JOIN project_transport pt ON pt.project_transport_id=s.project_transport_id
		JOIN project p ON p.project_id=pt.project_id
		JOIN project_route pr ON pr.project_route_id=pt.project_route_id
		LEFT JOIN truck t ON t.truck_id=pt.truck_id
		LEFT JOIN app_user d ON d.app_user_id=pt.driver_id AND d.deleted_at IS NULL
		LEFT JOIN vendor tv ON tv.vendor_id=pt.transport_vendor_id AND tv.deleted_at IS NULL
		WHERE s.project_transport_status_id=$1 AND s.deleted_at IS NULL AND pt.deleted_at IS NULL
		  AND (
			$2::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment ca
				WHERE ca.project_id=pt.project_id AND ca.checker_id=$2
				  AND ca.is_active=1 AND ca.deleted_at IS NULL
				  AND ca.access_started_at <= CURRENT_TIMESTAMP
				  AND (ca.access_ended_at IS NULL OR ca.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )`, id, checkerID))
}

func (r *ProjectTransportStatusRepository) List(ctx context.Context, opts model.ListOptions, transportID *int) ([]model.ProjectTransportStatus, error) {
	rows, err := r.DB.Query(ctx, `
		SELECT `+projectTransportStatusJSON+`
		FROM project_transport_status s
		JOIN project_transport pt ON pt.project_transport_id=s.project_transport_id
		JOIN project p ON p.project_id=pt.project_id
		JOIN project_route pr ON pr.project_route_id=pt.project_route_id
		LEFT JOIN truck t ON t.truck_id=pt.truck_id
		LEFT JOIN app_user d ON d.app_user_id=pt.driver_id AND d.deleted_at IS NULL
		LEFT JOIN vendor tv ON tv.vendor_id=pt.transport_vendor_id AND tv.deleted_at IS NULL
		WHERE s.deleted_at IS NULL AND pt.deleted_at IS NULL
		  AND ($1::int IS NULL OR s.project_transport_status_id < $1)
		  AND ($2::int IS NULL OR s.project_transport_id=$2)
		  AND ($3::int IS NULL OR pt.project_id=$3)
		  AND ($7::smallint IS NULL OR s.project_transport_status_type_id=$7)
		  AND ($8::smallint IS NULL OR s.is_fraud=$8)
		  AND ($9::smallint IS NULL OR s.is_active=$9)
		  AND ($10::timestamptz IS NULL OR s.status_time >= $10)
		  AND ($11::timestamptz IS NULL OR s.status_time < $11)
		  AND (
			$4::int IS NULL OR EXISTS (
				SELECT 1 FROM project_checker_assignment ca
				WHERE ca.project_id=pt.project_id AND ca.checker_id=$4
				  AND ca.is_active=1 AND ca.deleted_at IS NULL
				  AND ca.access_started_at <= CURRENT_TIMESTAMP
				  AND (ca.access_ended_at IS NULL OR ca.access_ended_at >= CURRENT_TIMESTAMP)
			)
		  )
		ORDER BY s.project_transport_status_id DESC LIMIT $5 OFFSET $6`,
		opts.CursorKey, transportID, opts.ProjectID, opts.CheckerID, opts.Limit, opts.Offset,
		opts.StatusTypeID, opts.IsFraud, opts.IsActive, opts.DateFrom, opts.DateTo)
	if err != nil {
		return nil, err
	}
	return scanJSONRows[model.ProjectTransportStatus](rows)
}
