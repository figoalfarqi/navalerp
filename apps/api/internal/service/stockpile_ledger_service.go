package service

import (
	"context"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type StockpileLedgerService struct {
	Repo *repository.StockpileLedgerRepository
}

func NewStockpileLedgerService(repo *repository.StockpileLedgerRepository) *StockpileLedgerService {
	return &StockpileLedgerService{Repo: repo}
}
func (s *StockpileLedgerService) GetByID(ctx context.Context, id int) (*model.StockpileLedger, error) {
	return s.Repo.GetByID(ctx, id)
}
func (s *StockpileLedgerService) List(ctx context.Context, opts model.ListOptions) ([]model.StockpileLedger, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}
