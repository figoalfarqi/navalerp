package service

import (
	"context"
	"errors"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ShipmentService struct {
	Repo *repository.ShipmentRepository
}

func NewShipmentService(repo *repository.ShipmentRepository) *ShipmentService {
	return &ShipmentService{Repo: repo}
}

func (s *ShipmentService) GetByID(ctx context.Context, id string) (*model.Shipment, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.Repo.Get(ctx, id)
}

func (s *ShipmentService) List(ctx context.Context, opts model.ListOptions) ([]model.Shipment, int, error) {
	return s.Repo.List(ctx, opts)
}

func (s *ShipmentService) Create(ctx context.Context, loginID string, req *model.ShipmentRequest) (*model.Shipment, error, map[string]string) {
	m := &model.Shipment{}
	if req.ManifestNumber != nil { m.ManifestNumber = *req.ManifestNumber }
	if req.RouteId != nil { m.RouteId = *req.RouteId }
	if req.TransportUnitId != nil { m.TransportUnitId = *req.TransportUnitId }
	if req.OriginWarehouseId != nil { m.OriginWarehouseId = *req.OriginWarehouseId }
	if req.DestinationWarehouseId != nil { m.DestinationWarehouseId = *req.DestinationWarehouseId }
	if req.DepartureDate != nil { m.DepartureDate = *req.DepartureDate }
	m.ArrivalDate = req.ArrivalDate
	m.EscortSecurityLevel = req.EscortSecurityLevel
	m.Status = req.Status
	m.AuthorizedByUserId = req.AuthorizedByUserId
	m.Remarks = req.Remarks
	m.Items = req.Items
	id, err := s.Repo.Create(ctx, m)
	if err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ShipmentService) Update(ctx context.Context, loginID string, id string, req *model.ShipmentRequest) (*model.Shipment, error, map[string]string) {
	m, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err, nil
	}
	if req.ManifestNumber != nil { m.ManifestNumber = *req.ManifestNumber }
	if req.RouteId != nil { m.RouteId = *req.RouteId }
	if req.TransportUnitId != nil { m.TransportUnitId = *req.TransportUnitId }
	if req.OriginWarehouseId != nil { m.OriginWarehouseId = *req.OriginWarehouseId }
	if req.DestinationWarehouseId != nil { m.DestinationWarehouseId = *req.DestinationWarehouseId }
	if req.DepartureDate != nil { m.DepartureDate = *req.DepartureDate }
	if req.ArrivalDate != nil { m.ArrivalDate = req.ArrivalDate }
	if req.EscortSecurityLevel != nil { m.EscortSecurityLevel = req.EscortSecurityLevel }
	if req.Status != nil { m.Status = req.Status }
	if req.AuthorizedByUserId != nil { m.AuthorizedByUserId = req.AuthorizedByUserId }
	if req.Remarks != nil { m.Remarks = req.Remarks }
	if req.Items != nil { m.Items = req.Items }
	if err := s.Repo.Update(ctx, id, m); err != nil {
		return nil, err, nil
	}
	item, err := s.Repo.Get(ctx, id)
	return item, err, nil
}

func (s *ShipmentService) Delete(ctx context.Context, loginID string, id string) error {
	return s.Repo.Delete(ctx, id)
}
