package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ContractService struct {
	Repo *repository.ContractRepository
}

func NewContractService(repo *repository.ContractRepository) *ContractService {
	return &ContractService{Repo: repo}
}

func (s *ContractService) GetByID(ctx context.Context, id string) (*model.Contract, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *ContractService) List(ctx context.Context, opts model.ListOptions) ([]model.Contract, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *ContractService) Create(ctx context.Context, loginID string, req *model.ContractRequest) (*model.Contract, error, map[string]string) {
	m := &model.Contract{}
	m.TenderId = req.TenderId
	if req.ContractNumber != nil { m.ContractNumber = *req.ContractNumber }
	if req.VendorId != nil { m.VendorId = *req.VendorId }
	if req.ContractTitle != nil { m.ContractTitle = *req.ContractTitle }
	if req.ContractValue != nil { m.ContractValue = *req.ContractValue }
	m.Currency = req.Currency
	if req.StartDate != nil { m.StartDate = *req.StartDate }
	if req.EndDate != nil { m.EndDate = *req.EndDate }
	m.ProcurementMethod = req.ProcurementMethod
	m.WarrantyPeriodMonths = req.WarrantyPeriodMonths
	m.TotClauseSummary = req.TotClauseSummary
	m.Status = req.Status
	m.Amendments = req.Amendments
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ContractService) Update(ctx context.Context, loginID string, id string, req *model.ContractRequest) (*model.Contract, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.TenderId != nil { m.TenderId = req.TenderId }
	if req.ContractNumber != nil { m.ContractNumber = *req.ContractNumber }
	if req.VendorId != nil { m.VendorId = *req.VendorId }
	if req.ContractTitle != nil { m.ContractTitle = *req.ContractTitle }
	if req.ContractValue != nil { m.ContractValue = *req.ContractValue }
	if req.Currency != nil { m.Currency = req.Currency }
	if req.StartDate != nil { m.StartDate = *req.StartDate }
	if req.EndDate != nil { m.EndDate = *req.EndDate }
	if req.ProcurementMethod != nil { m.ProcurementMethod = req.ProcurementMethod }
	if req.WarrantyPeriodMonths != nil { m.WarrantyPeriodMonths = req.WarrantyPeriodMonths }
	if req.TotClauseSummary != nil { m.TotClauseSummary = req.TotClauseSummary }
	if req.Status != nil { m.Status = req.Status }
	if req.Amendments != nil { m.Amendments = req.Amendments }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ContractService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
