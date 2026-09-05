package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type TenderService struct {
	Repo *repository.TenderRepository
}

func NewTenderService(repo *repository.TenderRepository) *TenderService {
	return &TenderService{Repo: repo}
}

func (s *TenderService) GetByID(ctx context.Context, id string) (*model.Tender, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *TenderService) List(ctx context.Context, opts model.ListOptions) ([]model.Tender, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *TenderService) Create(ctx context.Context, loginID string, req *model.TenderRequest) (*model.Tender, error, map[string]string) {
	m := &model.Tender{}
	if req.TenderNumber != nil { m.TenderNumber = *req.TenderNumber }
	if req.Title != nil { m.Title = *req.Title }
	if req.ProcurementCategory != nil { m.ProcurementCategory = *req.ProcurementCategory }
	if req.EstimatedBudget != nil { m.EstimatedBudget = *req.EstimatedBudget }
	if req.ProcurementMethod != nil { m.ProcurementMethod = *req.ProcurementMethod }
	if req.StartDate != nil { m.StartDate = *req.StartDate }
	if req.ClosingDate != nil { m.ClosingDate = *req.ClosingDate }
	m.Status = req.Status
	m.WinnerVendorId = req.WinnerVendorId
	m.Bids = req.Bids
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *TenderService) Update(ctx context.Context, loginID string, id string, req *model.TenderRequest) (*model.Tender, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.TenderNumber != nil { m.TenderNumber = *req.TenderNumber }
	if req.Title != nil { m.Title = *req.Title }
	if req.ProcurementCategory != nil { m.ProcurementCategory = *req.ProcurementCategory }
	if req.EstimatedBudget != nil { m.EstimatedBudget = *req.EstimatedBudget }
	if req.ProcurementMethod != nil { m.ProcurementMethod = *req.ProcurementMethod }
	if req.StartDate != nil { m.StartDate = *req.StartDate }
	if req.ClosingDate != nil { m.ClosingDate = *req.ClosingDate }
	if req.Status != nil { m.Status = req.Status }
	if req.WinnerVendorId != nil { m.WinnerVendorId = req.WinnerVendorId }
	if req.Bids != nil { m.Bids = req.Bids }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *TenderService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
