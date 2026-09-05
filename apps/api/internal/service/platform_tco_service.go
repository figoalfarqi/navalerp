package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type PlatformTcoService struct {
	Repo *repository.PlatformTcoRepository
}

func NewPlatformTcoService(repo *repository.PlatformTcoRepository) *PlatformTcoService {
	return &PlatformTcoService{Repo: repo}
}

func (s *PlatformTcoService) GetByID(ctx context.Context, id string) (*model.PlatformTco, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *PlatformTcoService) List(ctx context.Context, opts model.ListOptions) ([]model.PlatformTco, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *PlatformTcoService) Create(ctx context.Context, loginID string, req *model.PlatformTcoRequest) (*model.PlatformTco, error, map[string]string) {
	m := &model.PlatformTco{}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.FiscalYear != nil { m.FiscalYear = *req.FiscalYear }
	m.AcquisitionAmortization = req.AcquisitionAmortization
	m.FuelLubeCost = req.FuelLubeCost
	m.MroSparepartsCost = req.MroSparepartsCost
	m.DockingServicesCost = req.DockingServicesCost
	m.CrewPayrollAllowances = req.CrewPayrollAllowances
	m.ModernizationUpgradesCost = req.ModernizationUpgradesCost
	m.OperatingHoursSea = req.OperatingHoursSea
	m.CostPerOperatingHour = req.CostPerOperatingHour
	m.Remarks = req.Remarks
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *PlatformTcoService) Update(ctx context.Context, loginID string, id string, req *model.PlatformTcoRequest) (*model.PlatformTco, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.FiscalYear != nil { m.FiscalYear = *req.FiscalYear }
	if req.AcquisitionAmortization != nil { m.AcquisitionAmortization = req.AcquisitionAmortization }
	if req.FuelLubeCost != nil { m.FuelLubeCost = req.FuelLubeCost }
	if req.MroSparepartsCost != nil { m.MroSparepartsCost = req.MroSparepartsCost }
	if req.DockingServicesCost != nil { m.DockingServicesCost = req.DockingServicesCost }
	if req.CrewPayrollAllowances != nil { m.CrewPayrollAllowances = req.CrewPayrollAllowances }
	if req.ModernizationUpgradesCost != nil { m.ModernizationUpgradesCost = req.ModernizationUpgradesCost }
	if req.OperatingHoursSea != nil { m.OperatingHoursSea = req.OperatingHoursSea }
	if req.CostPerOperatingHour != nil { m.CostPerOperatingHour = req.CostPerOperatingHour }
	if req.Remarks != nil { m.Remarks = req.Remarks }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *PlatformTcoService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
