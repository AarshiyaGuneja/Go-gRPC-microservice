package acceptance_tests

import (
	"context"
	common "fleet-backend/common/proto"
	"fleet-backend/customer-service/proto"
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/emicklei/go-restful/log"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/consul"
	"testing"
)

func TestDriverWorkflow(t *testing.T) {
	client := proto.NewCustomerService("customer-service", client.NewClient(useConsul))
	response, err := client.SignUp(context.Background(), &proto.SignUpRequest{
		Name:             "Aarshiya",
		Password:         "oldpassword",
		Email:            "aarshiya.guneja@allonblock.com",
		FleetCompanyName: "FleetCompany1",
	})
	then.AssertThat(t, err, is.Nil())
	log.Print(response.Driver)
	then.AssertThat(t, response.Driver.Name, is.EqualTo("Aarshiya"))
	response, err = client.GetDriverById(context.Background(), &common.IdRequest{Id: response.Driver.Id})
	log.Print(response.Driver.Name)
	then.AssertThat(t, response.Driver.Name, is.EqualTo("Aarshiya"))
	then.AssertThat(t, response.Driver.Email, is.EqualTo("aarshiya.guneja@allonblock.com"))

	//Test GetDriversByFleetCompanyId
	log.Print("Starting GetDriversByFleetCompanyId")
	usersResponse, err := client.GetDriversByFleetCompanyId(context.Background(), &common.IdRequest{Id: response.Driver.FleetCompanyId})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, usersResponse.Drivers, has.Length(1))
	log.Print("Exiting GetDriversByFleetCompanyId")

	//Test CreateDriver
	createResponse, err := client.CreateDriver(context.Background(), &proto.Driver{
		Name:           "Driver1",
		Password:       "password",
		Email:          "Driver@allonblock.com",
		FleetCompanyId: response.Driver.FleetCompanyId,
	})
	log.Print("Starting CreateDriver")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, createResponse.Driver.Id == "", is.False())
	then.AssertThat(t, createResponse.Driver.FleetCompanyId == "", is.False())
	log.Print("Exiting CreateDriver")

	//Test UpdateDriver
	updateDriverReponse, err := client.UpdateDriver(context.Background(), &proto.Driver{
		Name:           "Driver updated",
		Password:       "new",
		Email:          "driverUpdated@allonblock.com",
		Id:             response.Driver.Id,
		FleetCompanyId: response.Driver.FleetCompanyId,
	})
	log.Print("Starting UpdateDriver")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, updateDriverReponse.Driver.Name, is.EqualTo("Driver updated"))
	then.AssertThat(t, updateDriverReponse.Driver.Password, is.EqualTo("new"))
	then.AssertThat(t, updateDriverReponse.Driver.Email, is.EqualTo("driverUpdated@allonblock.com"))
	log.Print("Exiting UpdateDriver")

}

func TestCorporation(t *testing.T) {
	//Test CreateCorporation
	client := proto.NewCustomerService("customer-service", client.NewClient(useConsul))
	response, err := client.CreateCorporation(context.Background(), &proto.Corporation{
		Name: "AllOnBlock Corporation",
	})
	log.Print("Starting CreateCorporation")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, response.Corporation.Id == "", is.False())
	then.AssertThat(t, response.Corporation.FleetCompanyId == "", is.False())
	then.AssertThat(t, response.Corporation.Name, is.EqualTo("AllOnBlock Corporation"))
	log.Print("Exiting CreateCorporation")

	//Test GetAllCorporationsByFleetCompanyId
	log.Print("Starting GetAllCorporationsByFleetCompanyId")
	log.Print(response.Corporation)
	corporationsResponse, err := client.GetAllCorporationsByFleetCompanyId(context.Background(), &common.IdRequest{Id: response.Corporation.FleetCompanyId})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, corporationsResponse.Corporations, has.Length(1))
	log.Print("Exiting GetAllCorporationsByFleetCompanyId")

	//Test GetCorporationById
	log.Print("Starting GetCorporationById")

	response, err = client.GetCorporationById(context.Background(), &common.IdRequest{Id: response.Corporation.Id})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, response.Corporation.Name, is.EqualTo("AllOnBlock Corporation"))
	log.Print("Exiting GetCorporationById")

	//Test CreateRegion
	createRegionResponse, err := client.CreateRegion(context.Background(), &proto.Region{
		Name:          "North",
		CorporationId: response.Corporation.Id,
	})
	log.Print("Starting CreateRegion")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, createRegionResponse.Region.Id == "", is.False())
	then.AssertThat(t, createRegionResponse.Region.CorporationId == "", is.False())
	then.AssertThat(t, createRegionResponse.Region.Name == "North", is.True())
	log.Print("Exiting CreateRegion")

	//Test GetAllRegionsByCorporationId
	log.Print("Starting GetAllRegionsByCorporationId")
	regionsByCorporationIdResponse, err := client.GetAllRegionsByCorporationId(context.Background(), &common.IdRequest{Id: createRegionResponse.Region.CorporationId})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, regionsByCorporationIdResponse.Regions, has.Length(1))
	log.Print("Exiting GetAllRegionsByCorporationId")

	//Test GetRegionById
	log.Print("Starting GetRegionById")
	getRegionByIdResponse, err := client.GetRegionById(context.Background(), &common.IdRequest{Id: createRegionResponse.Region.Id})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, getRegionByIdResponse.Region.Name, is.EqualTo("North"))
	log.Print("Exiting GetRegionById")

	//Test CreateDistrict
	createDistrictResponse, err := client.CreateDistrict(context.Background(), &proto.District{
		Name:     "Jaipur",
		RegionId: createRegionResponse.Region.Id,
	})
	log.Print("Starting CreateDistrict")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, createDistrictResponse.District.Id == "", is.False())
	then.AssertThat(t, createDistrictResponse.District.RegionId == "", is.False())
	then.AssertThat(t, createDistrictResponse.District.Name == "Jaipur", is.True())
	log.Print("Exiting CreateDistrict")

	//Test GetAllDistrictsByRegionId
	log.Print("Starting GetAllDistrictsByRegionId")
	districtByDistrictIdResponse, err := client.GetAllDistrictsByRegionId(context.Background(), &common.IdRequest{Id: createDistrictResponse.District.RegionId})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, districtByDistrictIdResponse.Districts, has.Length(1))
	log.Print("Exiting GetAllDistrictsByRegionId")

	//Test GetDistrictById
	log.Print("Starting GetDistrictById")
	getDistrictByIdResponse, err := client.GetDistrictById(context.Background(), &common.IdRequest{Id: createDistrictResponse.District.Id})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, getDistrictByIdResponse.District.Name, is.EqualTo("Jaipur"))
	log.Print("Exiting GetDistrictById")

	//Test CreateLocation
	createLocationResponse, err := client.CreateLocation(context.Background(), &proto.Location{
		Name:       "Mansarovar",
		DistrictId: getDistrictByIdResponse.District.Id,
	})
	log.Print("Starting CreateLocation")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, createLocationResponse.Location.Id == "", is.False())
	then.AssertThat(t, createLocationResponse.Location.DistrictId == "", is.False())
	then.AssertThat(t, createLocationResponse.Location.Name == "Mansarovar", is.True())
	log.Print("Exiting CreateLocation")

	//Test GetAllLocationsByDistrictId
	log.Print("Starting GetAllLocationsByDistrictId")
	locationByDistrictIdResponse, err := client.GetAllLocationsByDistrictId(context.Background(), &common.IdRequest{Id: createLocationResponse.Location.DistrictId})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, locationByDistrictIdResponse.Locations, has.Length(1))
	log.Print("Exiting GetAllLocationsByDistrictId")

	//Test GetLocationById
	log.Print("Starting GetLocationById")
	GetLocationByIdResponse, err := client.GetLocationById(context.Background(), &common.IdRequest{Id: createLocationResponse.Location.Id})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, GetLocationByIdResponse.Location.Name, is.EqualTo("Mansarovar"))
	log.Print("Exiting GetLocationById")

}

var useConsul = func(options *client.Options) {
	options.Registry = consul.NewRegistry()
}
