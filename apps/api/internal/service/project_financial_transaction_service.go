package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/apipml/internal/model"
	"github.com/figoalfarqi/apipml/internal/repository"
)

type ProjectFinancialTransactionService struct {
	Repo *repository.ProjectFinancialTransactionRepository
}

func NewProjectFinancialTransactionService(repo *repository.ProjectFinancialTransactionRepository) *ProjectFinancialTransactionService {
	return &ProjectFinancialTransactionService{Repo: repo}
}

func (s *ProjectFinancialTransactionService) validate(req *model.ProjectFinancialTransactionRequest) (error, map[string]string) {
	if err, fields := validateRequest(req); err != nil {
		return err, fields
	}
	fields := map[string]string{}
	if err := validateDate(&req.TransactionDate); err != nil {
		fields["transaction_date"] = "must use YYYY-MM-DD"
	}
	if err := validateDate(req.DueDate); err != nil {
		fields["due_date"] = "must use YYYY-MM-DD"
	}
	if len(fields) > 0 {
		return errors.New("validation error"), fields
	}
	return nil, nil
}
func (s *ProjectFinancialTransactionService) Create(ctx context.Context, userID int, req *model.ProjectFinancialTransactionRequest) (*model.ProjectFinancialTransaction, error, map[string]string) {
	if err, fields := s.validate(req); err != nil {
		return nil, err, fields
	}
	id, err := s.Repo.Create(ctx, req, userID)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}
func (s *ProjectFinancialTransactionService) Update(ctx context.Context, userID, id int, req *model.ProjectFinancialTransactionRequest) (*model.ProjectFinancialTransaction, error, map[string]string) {
	if err, fields := s.validate(req); err != nil {
		return nil, err, fields
	}
	if err := s.Repo.Update(ctx, id, req, userID); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.GetByID(ctx, id)
	return item, err, nil
}
func (s *ProjectFinancialTransactionService) Delete(ctx context.Context, userID, id int) error {
	return s.Repo.SoftDelete(ctx, id, userID)
}
func (s *ProjectFinancialTransactionService) GetByID(ctx context.Context, id int) (*model.ProjectFinancialTransaction, error) {
	return s.Repo.GetByID(ctx, id)
}
func (s *ProjectFinancialTransactionService) List(ctx context.Context, opts model.ListOptions) ([]model.ProjectFinancialTransaction, error) {
	return s.Repo.List(ctx, normalizeListOptions(opts))
}
