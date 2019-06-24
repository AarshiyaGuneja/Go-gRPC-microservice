package impl

import (
	"context"
	common "fleet-backend/common/proto"
	"fleet-backend/customer-service/proto"
	"github.com/google/uuid"
)

type Service struct {
	Service *CustomerServiceRepository
}

// Drivers

func (s Service) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.Driver, error) {
	client := &proto.FleetCompany{
		Id:   uuid.New().String(),
		Name: req.FleetCompanyName,
	}

	if err := s.Service.AddClient(ctx, client); err != nil {
		return nil, err
	}

	driver := &proto.Driver{
		Name:           req.Name,
		Id:             uuid.New().String(),
		Email:          req.Email,
		Password:       req.Password,
		FleetCompanyId: client.Id,
	}

	if err := s.Service.CreateDriver(ctx, driver); err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

func (s Service) CreateDriver(ctx context.Context, req *proto.Driver) (*proto.Driver, error) {
	driver := &proto.Driver{
		Name:           req.Name,
		Id:             uuid.New().String(),
		Email:          req.Email,
		FleetCompanyId: req.FleetCompanyId,
		Password:       req.Password,
	}

	if err := s.Service.CreateDriver(ctx, driver); err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

func (s Service) UpdateDriver(ctx context.Context, req *proto.Driver) (*proto.Driver, error) {
	if driver, err := s.Service.UpdateDriver(ctx, req); err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

func (s Service) DeleteDriverById(ctx context.Context, req *common.IdRequest) (*proto.Driver, error) {
	driver := &proto.Driver{
		Id: req.Id,
	}

	if err := s.Service.DeleteDriverById(ctx, driver); err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

func (s Service) GetDriverById(ctx context.Context, id string) (*proto.Driver, error) {
	if result, err := s.Service.GetDriverById(ctx, id); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (s Service) GetDriversByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([]*proto.Driver, error) {
	if results, err := s.Service.GetDriversByFleetCompanyId(ctx, fleetCompanyId); err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

// Corporation

func (s Service) CreateCorporation(ctx context.Context, req *proto.Corporation) (*proto.Corporation, error) {
	fleetcompany := &proto.FleetCompany{
		Id:   uuid.New().String(),
		Name: req.Name,
	}

	if err := s.Service.AddClient(ctx, fleetcompany); err != nil {
		return nil, err
	}

	corp := &proto.Corporation{
		Name:           req.Name,
		Id:             uuid.New().String(),
		FleetCompanyId: fleetcompany.Id,
	}

	if err := s.Service.CreateCorporation(ctx, corp); err != nil {
		return nil, err
	} else {
		return corp, err
	}
}

func (s Service) GetAllCorporationsByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([]*proto.Corporation, error) {
	if results, err := s.Service.GetAllCorporationsByFleetCompanyId(ctx, fleetCompanyId); err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func (s Service) GetCorporationById(ctx context.Context, id string) (*proto.Corporation, error) {
	if result, err := s.Service.GetCorporationById(ctx, id); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// Regions

func (s Service) CreateRegion(ctx context.Context, req *proto.Region) (*proto.Region, error) {
	region := &proto.Region{
		Name:          req.Name,
		Id:            uuid.New().String(),
		CorporationId: req.CorporationId,
	}

	if err := s.Service.CreateRegion(ctx, region); err != nil {
		return nil, err
	} else {
		return region, nil
	}
}

func (s Service) GetAllRegionsByCorporationId(ctx context.Context, corporationId string) ([]*proto.Region, error) {
	if results, err := s.Service.GetAllRegionsByCorporationId(ctx, corporationId); err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func (s Service) GetRegionById(ctx context.Context, id string) (*proto.Region, error) {
	if result, err := s.Service.GetRegionById(ctx, id); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// District

func (s Service) CreateDistrict(ctx context.Context, req *proto.District) (*proto.District, error) {
	district := &proto.District{
		Name:     req.Name,
		Id:       uuid.New().String(),
		RegionId: req.RegionId,
	}

	if err := s.Service.CreateDistrict(ctx, district); err != nil {
		return nil, err
	} else {
		return district, nil
	}
}

func (s Service) GetAllDistrictsByRegionId(ctx context.Context, regionId string) ([]*proto.District, error) {
	if results, err := s.Service.GetAllDistrictsByRegionId(ctx, regionId); err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func (s Service) GetDistrictById(ctx context.Context, id string) (*proto.District, error) {
	if result, err := s.Service.GetDistrictById(ctx, id); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// Location

func (s Service) CreateLocation(ctx context.Context, req *proto.Location) (*proto.Location, error) {
	location := &proto.Location{
		Name:       req.Name,
		Id:         uuid.New().String(),
		DistrictId: req.DistrictId,
	}

	if err := s.Service.CreateLocation(ctx, location); err != nil {
		return nil, err
	} else {
		return location, nil
	}
}

func (s Service) GetAllLocationsByDistrictId(ctx context.Context, districtId string) ([]*proto.Location, error) {
	if results, err := s.Service.GetAllLocationsByDistrictId(ctx, districtId); err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func (s Service) GetLocationById(ctx context.Context, id string) (*proto.Location, error) {
	if result, err := s.Service.GetLocationById(ctx, id); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
