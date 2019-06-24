package impl

import (
	"context"
	"fleet-backend/customer-service/common"
	"fleet-backend/customer-service/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type CustomerServiceRepository struct {
	client *mgo.Session
	dbName string
}

const dbName = "customer"
const driverCollection = "drivers"
const clientCollection = "clients"
const corporationCollection = "corporation"
const regionCollection = "region"
const districtCollection = "district"
const locationCollection = "location"

func NewCustomerRepository() (*CustomerServiceRepository, error) {
	if session, err := common.ConnectToMongo();
		err != nil {
		return nil, err
	} else {
		log.Print("Connected to Mongo")
		return &CustomerServiceRepository{
			client: session,
		}, nil
	}
}

func (c *CustomerServiceRepository) driverCollection() *mgo.Collection {
	return c.client.DB(dbName).C(driverCollection)
}

func (c *CustomerServiceRepository) clientCollection() *mgo.Collection {
	return c.client.DB(dbName).C(clientCollection)
}

func (c *CustomerServiceRepository) corporationCollection() *mgo.Collection {
	return c.client.DB(dbName).C(corporationCollection)
}

func (c *CustomerServiceRepository) regionCollection() *mgo.Collection {
	return c.client.DB(dbName).C(regionCollection)
}

func (c *CustomerServiceRepository) districtCollection() *mgo.Collection {
	return c.client.DB(dbName).C(districtCollection)
}

func (c *CustomerServiceRepository) locationCollection() *mgo.Collection {
	return c.client.DB(dbName).C(locationCollection)
}

func (c *CustomerServiceRepository) Close() {
	c.client.Close()
}

//Drivers

func (c *CustomerServiceRepository) SignUp(ctx context.Context, user *proto.Driver) error {
	if err := c.driverCollection().Insert(user); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) AddClient(ctx context.Context, client *proto.FleetCompany) error {
	if err := c.clientCollection().Insert(client); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) CreateDriver(ctx context.Context, user *proto.Driver) error {
	if err := c.driverCollection().Insert(user); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) UpdateDriver(ctx context.Context, user *proto.Driver) (*proto.Driver, error) {
	result := bson.M{"id": user.Id}
	change := bson.M{"$set": bson.M{"name": user.Name, "email": user.Email, "password": user.Password, "fleetcompanyid": user.FleetCompanyId}}
	err := c.driverCollection().Update(result, change)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (c *CustomerServiceRepository) DeleteDriverById(ctx context.Context, user *proto.Driver) error {
	err := c.driverCollection().Remove(
		bson.M{
			"id": bson.M{
				"$eq": user.Id,
			},
		})
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) GetDriverById(ctx context.Context, id string) (*proto.Driver, error) {
	var result *proto.Driver
	if err := c.driverCollection().Find(bson.M{"id": id}).One(&result); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (c *CustomerServiceRepository) GetDriversByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([]*proto.Driver, error) {
	var results [] *proto.Driver
	if err := c.driverCollection().Find(bson.M{"fleetcompanyid": fleetCompanyId}).All(&results); err != nil {
		return nil, err
	} else {
		return results, err
	}
}

// Corporation

func (c *CustomerServiceRepository) CreateCorporation(ctx context.Context, user *proto.Corporation) error {
	if err := c.corporationCollection().Insert(user); err != nil {
		return err
	} else {
		return nil
	}
}

func (c CustomerServiceRepository) GetAllCorporationsByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([]*proto.Corporation, error) {
	var results [] *proto.Corporation
	if err := c.corporationCollection().Find(bson.M{"fleetcompanyid": fleetCompanyId}).All(&results); err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func (c *CustomerServiceRepository) GetCorporationById(ctx context.Context, id string) (*proto.Corporation, error) {
	var result *proto.Corporation
	err := c.corporationCollection().Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// Regions

func (c *CustomerServiceRepository) CreateRegion(ctx context.Context, user *proto.Region) error {
	if err := c.regionCollection().Insert(user); err != nil {
		return err
	} else {
		return nil
	}
}

func (c CustomerServiceRepository) GetAllRegionsByCorporationId(ctx context.Context, corporationId string) ([]*proto.Region, error) {
	var results [] *proto.Region
	if err := c.regionCollection().Find(bson.M{"corporationid": corporationId}).All(&results); err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func (c *CustomerServiceRepository) GetRegionById(ctx context.Context, id string) (*proto.Region, error) {
	var result *proto.Region
	err := c.regionCollection().Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// District

func (c *CustomerServiceRepository) CreateDistrict(ctx context.Context, user *proto.District) error {
	if err := c.districtCollection().Insert(user); err != nil {
		return err
	} else {
		return nil
	}
}

func (c CustomerServiceRepository) GetAllDistrictsByRegionId(ctx context.Context, regionId string) ([]*proto.District, error) {
	var results [] *proto.District
	if err := c.districtCollection().Find(bson.M{"regionid": regionId}).All(&results); err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func (c *CustomerServiceRepository) GetDistrictById(ctx context.Context, id string) (*proto.District, error) {
	var result *proto.District
	err := c.districtCollection().Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// Location

func (c *CustomerServiceRepository) CreateLocation(ctx context.Context, user *proto.Location) error {
	if err := c.locationCollection().Insert(user); err != nil {
		return err
	} else {
		return nil
	}
}

func (c CustomerServiceRepository) GetAllLocationsByDistrictId(ctx context.Context, districtId string) ([]*proto.Location, error) {
	var results [] *proto.Location
	if err := c.locationCollection().Find(bson.M{"districtid": districtId}).All(&results); err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func (c *CustomerServiceRepository) GetLocationById(ctx context.Context, id string) (*proto.Location, error) {
	var result *proto.Location
	err := c.locationCollection().Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
