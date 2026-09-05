package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type BudgetProgramService struct {
	Repo *repository.BudgetProgramRepository
}

func NewBudgetProgramService(repo *repository.BudgetProgramRepository) *BudgetProgramService {
	return &BudgetProgramService{Repo: repo}
}

func (s *BudgetProgramService) GetByID(ctx context.Context, id string) (*model.BudgetProgram, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *BudgetProgramService) List(ctx context.Context, opts model.ListOptions) ([]model.BudgetProgram, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *BudgetProgramService) Create(ctx context.Context, loginID string, req *model.BudgetProgramRequest) (*model.BudgetProgram, error, map[string]string) {
	m := &model.BudgetProgram{}
	if req.FiscalYear != nil { m.FiscalYear = *req.FiscalYear }
	if req.DipaNumber != nil { m.DipaNumber = *req.DipaNumber }
	if req.ProgramCode != nil { m.ProgramCode = *req.ProgramCode }
	if req.ProgramName != nil { m.ProgramName = *req.ProgramName }
	if req.TotalBudget != nil { m.TotalBudget = *req.TotalBudget }
	if req.ResponsibleUnitId != nil { m.ResponsibleUnitId = *req.ResponsibleUnitId }
	m.Status = req.Status
	m.Allocations = req.Allocations
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *BudgetProgramService) Update(ctx context.Context, loginID string, id string, req *model.BudgetProgramRequest) (*model.BudgetProgram, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.FiscalYear != nil { m.FiscalYear = *req.FiscalYear }
	if req.DipaNumber != nil { m.DipaNumber = *req.DipaNumber }
	if req.ProgramCode != nil { m.ProgramCode = *req.ProgramCode }
	if req.ProgramName != nil { m.ProgramName = *req.ProgramName }
	if req.TotalBudget != nil { m.TotalBudget = *req.TotalBudget }
	if req.ResponsibleUnitId != nil { m.ResponsibleUnitId = *req.ResponsibleUnitId }
	if req.Status != nil { m.Status = req.Status }
	if req.Allocations != nil { m.Allocations = req.Allocations }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *BudgetProgramService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
