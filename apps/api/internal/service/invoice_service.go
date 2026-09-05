package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type InvoiceService struct {
	Repo *repository.InvoiceRepository
}

func NewInvoiceService(repo *repository.InvoiceRepository) *InvoiceService {
	return &InvoiceService{Repo: repo}
}

func (s *InvoiceService) GetByID(ctx context.Context, id string) (*model.Invoice, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *InvoiceService) List(ctx context.Context, opts model.ListOptions) ([]model.Invoice, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *InvoiceService) Create(ctx context.Context, loginID string, req *model.InvoiceRequest) (*model.Invoice, error, map[string]string) {
	m := &model.Invoice{}
	if req.InvoiceNumber != nil { m.InvoiceNumber = *req.InvoiceNumber }
	if req.VendorId != nil { m.VendorId = *req.VendorId }
	m.ContractId = req.ContractId
	m.PoId = req.PoId
	if req.InvoiceDate != nil { m.InvoiceDate = *req.InvoiceDate }
	if req.DueDate != nil { m.DueDate = *req.DueDate }
	m.TaxInvoiceNumber = req.TaxInvoiceNumber
	if req.Subtotal != nil { m.Subtotal = *req.Subtotal }
	m.TaxAmount = req.TaxAmount
	m.VerificationStatus = req.VerificationStatus
	m.VerifiedByUserId = req.VerifiedByUserId
	m.PaymentStatus = req.PaymentStatus
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *InvoiceService) Update(ctx context.Context, loginID string, id string, req *model.InvoiceRequest) (*model.Invoice, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.InvoiceNumber != nil { m.InvoiceNumber = *req.InvoiceNumber }
	if req.VendorId != nil { m.VendorId = *req.VendorId }
	if req.ContractId != nil { m.ContractId = req.ContractId }
	if req.PoId != nil { m.PoId = req.PoId }
	if req.InvoiceDate != nil { m.InvoiceDate = *req.InvoiceDate }
	if req.DueDate != nil { m.DueDate = *req.DueDate }
	if req.TaxInvoiceNumber != nil { m.TaxInvoiceNumber = req.TaxInvoiceNumber }
	if req.Subtotal != nil { m.Subtotal = *req.Subtotal }
	if req.TaxAmount != nil { m.TaxAmount = req.TaxAmount }
	if req.VerificationStatus != nil { m.VerificationStatus = req.VerificationStatus }
	if req.VerifiedByUserId != nil { m.VerifiedByUserId = req.VerifiedByUserId }
	if req.PaymentStatus != nil { m.PaymentStatus = req.PaymentStatus }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *InvoiceService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
