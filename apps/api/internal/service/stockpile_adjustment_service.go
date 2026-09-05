package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type StockpileAdjustmentService struct {
	Repo *repository.StockpileAdjustmentRepository
}

func NewStockpileAdjustmentService(
	repo *repository.StockpileAdjustmentRepository,
	_ *repository.StockpileCargoRepository,
	_ *repository.StockpileLedgerRepository,
	_ *StockpileLedgerService,
) *StockpileAdjustmentService {
	return &StockpileAdjustmentService{Repo: repo}
}

func validateStockpileAdjustment(req *model.StockpileAdjustmentRequest) (error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return err, fields
	}
	fields := map[string]string{}
	req.AdjustmentDate = strings.TrimSpace(req.AdjustmentDate)
	if _, err := time.Parse("2006-01-02", req.AdjustmentDate); err != nil {
		fields["adjustment_date"] = "must use YYYY-MM-DD"
	}
	if req.AmountVolumeCubic == 0 && req.AmountWeightTon == 0 {
		fields["amount_volume_cubic"] = "volume or weight amount is required"
		fields["amount_weight_ton"] = "volume or weight amount is required"
	}
	if len(fields) > 0 {
		return errors.New("validation error"), fields
	}
	return nil, nil
}

func (s *StockpileAdjustmentService) Create(ctx context.Context, userID int, req *model.StockpileAdjustmentRequest) (*model.StockpileAdjustment, error, map[string]string) {
	if err, fields := validateStockpileAdjustment(req); err != nil {
		return nil, err, fields
	}
	id, err := s.Repo.Create(ctx, req, userID)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}
func (s *StockpileAdjustmentService) Update(ctx context.Context, userID, id int, req *model.StockpileAdjustmentRequest) (*model.StockpileAdjustment, error, map[string]string) {
	if err, fields := validateStockpileAdjustment(req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}
func (s *StockpileAdjustmentService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, userID, id)
}
func (s *StockpileAdjustmentService) GetByID(ctx context.Context, id int) (*model.StockpileAdjustment, error) {
	return s.Repo.GetByID(ctx, id)
}
func (s *StockpileAdjustmentService) List(ctx context.Context, opts model.ListOptions) ([]model.StockpileAdjustment, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}
