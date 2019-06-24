package impl

import (
	"context"
	"fleet-backend/truck-service/common"
	"fleet-backend/truck-service/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type TruckServiceRepository struct {
	client *mgo.Session
	dbName string
}

const dbName = "truck"
const truckCollection = "trucks"

func (c *TruckServiceRepository) truckCollection() *mgo.Collection {
	return c.client.DB(dbName).C(truckCollection)
}

func NewTruckRepository() (*TruckServiceRepository, error) {
	if session, err := common.ConnectToMongo();
		err != nil {
		return nil, err
	} else {
		log.Print("Connected to Mongo")
		return &TruckServiceRepository{
			client: session,
			dbName: "truck-db",
		}, nil
	}
}

func (c *TruckServiceRepository) Close() {
	c.client.Close()
}

func (c *TruckServiceRepository) CreateTruck(ctx context.Context, user *proto.Truck) error {
	if err := c.truckCollection().Insert(user); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *TruckServiceRepository) UpdateTruck(ctx context.Context, user *proto.Truck) (*proto.Truck, error) {
	result := bson.M{"id": user.Id}
	change := bson.M{"$set": bson.M{"licenseplate": user.LicensePlate, "clockedinuser": user.ClockedInUser, "miles": user.Miles, "fleetcompanyid": user.FleetCompanyId, "corporationid": user.CorporationId, "regionid": user.RegionId, "districtid": user.DistrictId, "locationid": user.LocationId}}
	err := c.truckCollection().Update(result, change)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (c *TruckServiceRepository) GetTruckById(ctx context.Context, id string) (*proto.Truck, error) {
	var result *proto.Truck
	if err := c.truckCollection().Find(bson.M{"id": id}).One(&result); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (c *TruckServiceRepository) GetAllTrucksByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([]*proto.Truck, error) {
	var results [] *proto.Truck
	if err := c.truckCollection().Find(bson.M{"fleetcompanyid": fleetCompanyId}).All(&results); err != nil {
		return nil, err
	} else {
		return results, err
	}
}
