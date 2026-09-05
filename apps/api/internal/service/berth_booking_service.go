package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type BerthBookingService struct {
	Repo *repository.BerthBookingRepository
}

func NewBerthBookingService(repo *repository.BerthBookingRepository) *BerthBookingService {
	return &BerthBookingService{Repo: repo}
}

func (s *BerthBookingService) GetByID(ctx context.Context, id string) (*model.BerthBooking, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *BerthBookingService) List(ctx context.Context, opts model.ListOptions) ([]model.BerthBooking, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *BerthBookingService) Create(ctx context.Context, loginID string, req *model.BerthBookingRequest) (*model.BerthBooking, error, map[string]string) {
	m := &model.BerthBooking{}
	if req.FacilityId != nil { m.FacilityId = *req.FacilityId }
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.BookingPurpose != nil { m.BookingPurpose = *req.BookingPurpose }
	if req.Eta != nil { m.Eta = *req.Eta }
	if req.Etd != nil { m.Etd = *req.Etd }
	m.ActualBerthTime = req.ActualBerthTime
	m.ActualUnberthTime = req.ActualUnberthTime
	m.ShorePowerKwhUsed = req.ShorePowerKwhUsed
	m.FreshWaterTonUsed = req.FreshWaterTonUsed
	m.Status = req.Status
	m.ApprovedByUserId = req.ApprovedByUserId
	m.Remarks = req.Remarks
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *BerthBookingService) Update(ctx context.Context, loginID string, id string, req *model.BerthBookingRequest) (*model.BerthBooking, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.FacilityId != nil { m.FacilityId = *req.FacilityId }
	if req.ShipId != nil { m.ShipId = *req.ShipId }
	if req.BookingPurpose != nil { m.BookingPurpose = *req.BookingPurpose }
	if req.Eta != nil { m.Eta = *req.Eta }
	if req.Etd != nil { m.Etd = *req.Etd }
	if req.ActualBerthTime != nil { m.ActualBerthTime = req.ActualBerthTime }
	if req.ActualUnberthTime != nil { m.ActualUnberthTime = req.ActualUnberthTime }
	if req.ShorePowerKwhUsed != nil { m.ShorePowerKwhUsed = req.ShorePowerKwhUsed }
	if req.FreshWaterTonUsed != nil { m.FreshWaterTonUsed = req.FreshWaterTonUsed }
	if req.Status != nil { m.Status = req.Status }
	if req.ApprovedByUserId != nil { m.ApprovedByUserId = req.ApprovedByUserId }
	if req.Remarks != nil { m.Remarks = req.Remarks }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *BerthBookingService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
