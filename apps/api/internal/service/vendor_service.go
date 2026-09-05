package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type VendorService struct {
	Repo *repository.VendorRepository
}

func NewVendorService(repo *repository.VendorRepository) *VendorService {
	return &VendorService{Repo: repo}
}

func (s *VendorService) GetByID(ctx context.Context, id string) (*model.Vendor, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *VendorService) List(ctx context.Context, opts model.ListOptions) ([]model.Vendor, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *VendorService) Create(ctx context.Context, loginID string, req *model.VendorRequest) (*model.Vendor, error, map[string]string) {
	m := &model.Vendor{}
	if req.VendorCode != nil { m.VendorCode = *req.VendorCode }
	if req.VendorName != nil { m.VendorName = *req.VendorName }
	m.TaxNumber = req.TaxNumber
	m.SecurityClearanceLevel = req.SecurityClearanceLevel
	m.DefenceIndustryLicenseNo = req.DefenceIndustryLicenseNo
	m.Country = req.Country
	m.ContactPerson = req.ContactPerson
	m.Email = req.Email
	m.Phone = req.Phone
	m.BankAccountName = req.BankAccountName
	m.BankAccountNo = req.BankAccountNo
	m.BankName = req.BankName
	m.PerformanceRating = req.PerformanceRating
	m.IsApproved = req.IsApproved
	m.Ratings = req.Ratings
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *VendorService) Update(ctx context.Context, loginID string, id string, req *model.VendorRequest) (*model.Vendor, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.VendorCode != nil { m.VendorCode = *req.VendorCode }
	if req.VendorName != nil { m.VendorName = *req.VendorName }
	if req.TaxNumber != nil { m.TaxNumber = req.TaxNumber }
	if req.SecurityClearanceLevel != nil { m.SecurityClearanceLevel = req.SecurityClearanceLevel }
	if req.DefenceIndustryLicenseNo != nil { m.DefenceIndustryLicenseNo = req.DefenceIndustryLicenseNo }
	if req.Country != nil { m.Country = req.Country }
	if req.ContactPerson != nil { m.ContactPerson = req.ContactPerson }
	if req.Email != nil { m.Email = req.Email }
	if req.Phone != nil { m.Phone = req.Phone }
	if req.BankAccountName != nil { m.BankAccountName = req.BankAccountName }
	if req.BankAccountNo != nil { m.BankAccountNo = req.BankAccountNo }
	if req.BankName != nil { m.BankName = req.BankName }
	if req.PerformanceRating != nil { m.PerformanceRating = req.PerformanceRating }
	if req.IsApproved != nil { m.IsApproved = req.IsApproved }
	if req.Ratings != nil { m.Ratings = req.Ratings }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *VendorService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
