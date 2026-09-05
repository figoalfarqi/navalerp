package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ChartOfAccountService struct {
	Repo *repository.ChartOfAccountRepository
}

func NewChartOfAccountService(repo *repository.ChartOfAccountRepository) *ChartOfAccountService {
	return &ChartOfAccountService{Repo: repo}
}

func (s *ChartOfAccountService) GetByID(ctx context.Context, id string) (*model.ChartOfAccount, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *ChartOfAccountService) List(ctx context.Context, opts model.ListOptions) ([]model.ChartOfAccount, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *ChartOfAccountService) Create(ctx context.Context, loginID string, req *model.ChartOfAccountRequest) (*model.ChartOfAccount, error, map[string]string) {
	m := &model.ChartOfAccount{}
	if req.AccountCode != nil { m.AccountCode = *req.AccountCode }
	if req.AccountName != nil { m.AccountName = *req.AccountName }
	if req.AccountType != nil { m.AccountType = *req.AccountType }
	m.ParentAccountId = req.ParentAccountId
	m.IsActive = req.IsActive
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ChartOfAccountService) Update(ctx context.Context, loginID string, id string, req *model.ChartOfAccountRequest) (*model.ChartOfAccount, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.AccountCode != nil { m.AccountCode = *req.AccountCode }
	if req.AccountName != nil { m.AccountName = *req.AccountName }
	if req.AccountType != nil { m.AccountType = *req.AccountType }
	if req.ParentAccountId != nil { m.ParentAccountId = req.ParentAccountId }
	if req.IsActive != nil { m.IsActive = req.IsActive }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ChartOfAccountService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
