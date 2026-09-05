package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type BudgetCommitmentService struct {
	Repo *repository.BudgetCommitmentRepository
}

func NewBudgetCommitmentService(repo *repository.BudgetCommitmentRepository) *BudgetCommitmentService {
	return &BudgetCommitmentService{Repo: repo}
}

func (s *BudgetCommitmentService) GetByID(ctx context.Context, id string) (*model.BudgetCommitment, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *BudgetCommitmentService) List(ctx context.Context, opts model.ListOptions) ([]model.BudgetCommitment, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *BudgetCommitmentService) Create(ctx context.Context, loginID string, req *model.BudgetCommitmentRequest) (*model.BudgetCommitment, error, map[string]string) {
	m := &model.BudgetCommitment{}
	if req.CommitmentNumber != nil { m.CommitmentNumber = *req.CommitmentNumber }
	if req.AllocationId != nil { m.AllocationId = *req.AllocationId }
	m.ContractId = req.ContractId
	m.PoId = req.PoId
	m.WorkOrderId = req.WorkOrderId
	if req.CommittedAmount != nil { m.CommittedAmount = *req.CommittedAmount }
	if req.CommitmentDate != nil { m.CommitmentDate = *req.CommitmentDate }
	m.Status = req.Status
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *BudgetCommitmentService) Update(ctx context.Context, loginID string, id string, req *model.BudgetCommitmentRequest) (*model.BudgetCommitment, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.CommitmentNumber != nil { m.CommitmentNumber = *req.CommitmentNumber }
	if req.AllocationId != nil { m.AllocationId = *req.AllocationId }
	if req.ContractId != nil { m.ContractId = req.ContractId }
	if req.PoId != nil { m.PoId = req.PoId }
	if req.WorkOrderId != nil { m.WorkOrderId = req.WorkOrderId }
	if req.CommittedAmount != nil { m.CommittedAmount = *req.CommittedAmount }
	if req.CommitmentDate != nil { m.CommitmentDate = *req.CommitmentDate }
	if req.Status != nil { m.Status = req.Status }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *BudgetCommitmentService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
