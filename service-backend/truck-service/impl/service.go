package impl

import (
	"context"
	"fleet-backend/truck-service/proto"
	"github.com/google/uuid"
)

type Service struct {
	Service *TruckServiceRepository
}

func (s Service) CreateTruck(ctx context.Context, req *proto.Truck) (*proto.Truck, error) {
	truck := &proto.Truck{
		Id:             uuid.New().String(),
		LicensePlate:   req.LicensePlate,
		ClockedInUser:  req.ClockedInUser,
		Miles:          req.Miles,
		FleetCompanyId: req.FleetCompanyId,
		CorporationId:  req.CorporationId,
		RegionId:       req.RegionId,
		DistrictId:     req.DistrictId,
		LocationId:     req.LocationId,
	}

	if err := s.Service.CreateTruck(ctx, truck); err != nil {
		return nil, err
	} else {
		return truck, nil
	}
}

func (s Service) UpdateTruck(ctx context.Context, req *proto.Truck) (*proto.Truck, error) {
	if truck, err := s.Service.UpdateTruck(ctx, req); err != nil {
		return nil, err
	} else {
		return truck, nil
	}
}

func (s Service) GetTruckById(ctx context.Context, id string) (*proto.Truck, error) {
	if result, err := s.Service.GetTruckById(ctx, id); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (s Service) GetAllTrucksByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([]*proto.Truck, error) {
	if results, err := s.Service.GetAllTrucksByFleetCompanyId(ctx, fleetCompanyId); err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func (s Service) ClockIn(ctx context.Context, operation *proto.ClockOperation) (*proto.Truck, error) {
	if truck, err := s.Service.GetTruckById(ctx, operation.TruckId); err != nil {
		return nil, err
	} else {
		if truck.ClockedInUser == "" {
			truck.ClockedInUser = operation.DriverId
			if truck, err := s.Service.UpdateTruck(ctx, truck); err != nil {
				return nil, err
			} else {
				return truck, nil
			}
		}
		return truck, nil
	}
}

func (s Service) ClockOut(ctx context.Context, operation *proto.ClockOperation) (*proto.Truck, error) {
	if truck, err := s.Service.GetTruckById(ctx, operation.TruckId); err != nil {
		return nil, err
	} else {
		if truck.ClockedInUser == operation.DriverId {
			truck.ClockedInUser = ""
			if truck, err := s.Service.UpdateTruck(ctx, truck); err != nil {
				return nil, err
			} else {
				return truck, nil
			}
		}
		return truck, nil
	}
}
