package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type PaymentService struct {
	Repo *repository.PaymentRepository
}

func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{Repo: repo}
}

func (s *PaymentService) GetByID(ctx context.Context, id string) (*model.Payment, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *PaymentService) List(ctx context.Context, opts model.ListOptions) ([]model.Payment, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *PaymentService) Create(ctx context.Context, loginID string, req *model.PaymentRequest) (*model.Payment, error, map[string]string) {
	m := &model.Payment{}
	if req.PaymentReferenceNo != nil { m.PaymentReferenceNo = *req.PaymentReferenceNo }
	m.SppNumber = req.SppNumber
	m.SpmNumber = req.SpmNumber
	if req.InvoiceId != nil { m.InvoiceId = *req.InvoiceId }
	if req.PaymentDate != nil { m.PaymentDate = *req.PaymentDate }
	if req.AmountPaid != nil { m.AmountPaid = *req.AmountPaid }
	m.PaymentMethod = req.PaymentMethod
	m.BankSourceAccount = req.BankSourceAccount
	m.AuthorisedByUserId = req.AuthorisedByUserId
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *PaymentService) Update(ctx context.Context, loginID string, id string, req *model.PaymentRequest) (*model.Payment, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.PaymentReferenceNo != nil { m.PaymentReferenceNo = *req.PaymentReferenceNo }
	if req.SppNumber != nil { m.SppNumber = req.SppNumber }
	if req.SpmNumber != nil { m.SpmNumber = req.SpmNumber }
	if req.InvoiceId != nil { m.InvoiceId = *req.InvoiceId }
	if req.PaymentDate != nil { m.PaymentDate = *req.PaymentDate }
	if req.AmountPaid != nil { m.AmountPaid = *req.AmountPaid }
	if req.PaymentMethod != nil { m.PaymentMethod = req.PaymentMethod }
	if req.BankSourceAccount != nil { m.BankSourceAccount = req.BankSourceAccount }
	if req.AuthorisedByUserId != nil { m.AuthorisedByUserId = req.AuthorisedByUserId }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *PaymentService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
