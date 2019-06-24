package acceptance_tests

import (
	"context"
	common "fleet-backend/common/proto"
	"fleet-backend/truck-service/proto"
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"testing"
)

var UseConsul = func(options *client.Options) {
	options.Registry = consul.NewRegistry()
}

func truckServiceClient() proto.TruckService {
	return proto.NewTruckService("truck-service", client.NewClient(UseConsul))
}

func TestTruck(t *testing.T) {
	//Test SignUp
	createTruckResponse, err := truckServiceClient().CreateTruck(context.Background(), &proto.Truck{
		LicensePlate:   "RJ14-AG-0000",
		ClockedInUser:  "User-1",
		Miles:          1000,
		FleetCompanyId: uuid.New().String(),
		CorporationId:  uuid.New().String(),
		RegionId:       uuid.New().String(),
		DistrictId:     uuid.New().String(),
		LocationId:     uuid.New().String(),
	})

	then.AssertThat(t, err, is.Nil())

	truck := createTruckResponse.Truck
	then.AssertThat(t, truck.Id == "", is.False())
	then.AssertThat(t, truck.LicensePlate, is.EqualTo("RJ14-AG-0000"))
	then.AssertThat(t, truck.ClockedInUser, is.EqualTo("User-1"))
	then.AssertThat(t, truck.Miles == 1000, is.True())
	then.AssertThat(t, truck.FleetCompanyId == "", is.False())
	then.AssertThat(t, truck.CorporationId == "", is.False())
	then.AssertThat(t, truck.RegionId == "", is.False())
	then.AssertThat(t, truck.DistrictId == "", is.False())
	then.AssertThat(t, truck.LocationId == "", is.False())

	//Test UpdateTruck
	updateTruckResponse, err := truckServiceClient().UpdateTruck(context.Background(), &proto.Truck{
		Id:             truck.Id,
		LicensePlate:   "RJ14-AG-1111",
		ClockedInUser:  "",
		Miles:          2000,
		FleetCompanyId: truck.FleetCompanyId,
		CorporationId:  truck.CorporationId,
		RegionId:       truck.RegionId,
		DistrictId:     truck.DistrictId,
		LocationId:     truck.LocationId,
	})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, updateTruckResponse.Truck.LicensePlate, is.EqualTo("RJ14-AG-1111"))
	then.AssertThat(t, updateTruckResponse.Truck.ClockedInUser, is.EqualTo(""))
	then.AssertThat(t, updateTruckResponse.Truck.Miles == 2000, is.True())
	then.AssertThat(t, updateTruckResponse.Truck.FleetCompanyId, is.EqualTo(truck.FleetCompanyId))
	then.AssertThat(t, updateTruckResponse.Truck.CorporationId, is.EqualTo(truck.CorporationId))
	then.AssertThat(t, updateTruckResponse.Truck.RegionId, is.EqualTo(truck.RegionId))
	then.AssertThat(t, updateTruckResponse.Truck.DistrictId, is.EqualTo(truck.DistrictId))
	then.AssertThat(t, updateTruckResponse.Truck.LocationId, is.EqualTo(truck.LocationId))

	//Test GetTruckById
	getTruckByIdResponse, err := truckServiceClient().GetTruckById(context.Background(), &common.IdRequest{Id: updateTruckResponse.Truck.Id})
	then.AssertThat(t, err, is.Nil())
	truck = getTruckByIdResponse.Truck
	then.AssertThat(t, truck.LicensePlate, is.EqualTo("RJ14-AG-1111"))
	then.AssertThat(t, truck.ClockedInUser, is.EqualTo(""))

	//Test GetDriversByFleetCompanyId
	trucksByFleetCompanyIdResponse, err := truckServiceClient().GetAllTrucksByFleetCompanyId(context.Background(), &common.IdRequest{Id: truck.FleetCompanyId})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, trucksByFleetCompanyIdResponse.Trucks, has.Length(1))

	//Test ClockIn
	clockInResponse, err := truckServiceClient().ClockIn(context.Background(), &proto.ClockOperation{
		DriverId: "Driver-1",
		TruckId:  truck.Id,
	})
	log.Print(clockInResponse)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, clockInResponse.Truck.Id, is.EqualTo(truck.Id))
	then.AssertThat(t, clockInResponse.Truck.ClockedInUser, is.EqualTo("Driver-1"))

	//Test ClockOut
	clockOutResponse, err := truckServiceClient().ClockOut(context.Background(), &proto.ClockOperation{
		DriverId: "Driver-1",
		TruckId:  truck.Id,
	})
	log.Print(clockOutResponse)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, clockOutResponse.Truck.Id, is.EqualTo(truck.Id))
	then.AssertThat(t, clockOutResponse.Truck.ClockedInUser, is.EqualTo(""))
}
