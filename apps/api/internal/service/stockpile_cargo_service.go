package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type StockpileCargoService struct {
	Repo *repository.StockpileCargoRepository
}

func NewStockpileCargoService(repo *repository.StockpileCargoRepository) *StockpileCargoService {
	return &StockpileCargoService{Repo: repo}
}

func validateStockpileCargo(req *model.StockpileCargoRequest) (error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return err, fields
	}
	fields := map[string]string{}
	if req.CapacityVolumeCubic != nil && req.CurrentVolumeCubic != nil &&
		*req.CurrentVolumeCubic > *req.CapacityVolumeCubic {
		fields["current_volume_cubic"] = "cannot exceed capacity_volume_cubic"
	}
	if req.CapacityWeightTon != nil && req.CurrentWeightTon != nil &&
		*req.CurrentWeightTon > *req.CapacityWeightTon {
		fields["current_weight_ton"] = "cannot exceed capacity_weight_ton"
	}
	if len(fields) > 0 {
		return errors.New("validation error"), fields
	}
	return nil, nil
}

func (s *StockpileCargoService) Create(ctx context.Context, userID int, req *model.StockpileCargoRequest) (*model.StockpileCargo, error, map[string]string) {
	if err, fields := validateStockpileCargo(req); err != nil {
		return nil, err, fields
	}
	id, err := s.Repo.Create(ctx, req, userID)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}

func (s *StockpileCargoService) Update(ctx context.Context, userID, id int, req *model.StockpileCargoRequest) (*model.StockpileCargo, error, map[string]string) {
	if err, fields := validateStockpileCargo(req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}
func (s *StockpileCargoService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, userID, id)
}
func (s *StockpileCargoService) GetByID(ctx context.Context, id int) (*model.StockpileCargo, error) {
	return s.Repo.GetByID(ctx, id)
}
func (s *StockpileCargoService) List(ctx context.Context, opts model.ListOptions) ([]model.StockpileCargo, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}

func toStockpileCargoResp(item *model.StockpileCargo) *model.StockpileCargoResponse {
	return item
}
