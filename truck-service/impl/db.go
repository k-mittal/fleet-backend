package impl

import (
	"context"
	"fleet-backend/common"
	"fleet-backend/truck-service/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName ="customer"
const truckCollection = "truck"

type TruckServiceRepository struct {
	dbName              string
	fleetcompanySession *mgo.Session
}

func NewTruckServiceRepository() (*TruckServiceRepository, error) {
	if session, err := common.ConnectToMongo(); err != nil {
		return nil, err
	} else {
		return &TruckServiceRepository{
			fleetcompanySession: session,
		}, err
	}
}

func (c *TruckServiceRepository) Close() {
	c.fleetcompanySession.Close()
}

func (c *TruckServiceRepository) AddTruck(context context.Context, truck *proto.Truck) error {
	if err := c.truckCollection().Insert(truck); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *TruckServiceRepository) truckCollection() *mgo.Collection {
	return c.fleetcompanySession.DB(dbName).C(truckCollection)
}

func (c *TruckServiceRepository) UpdateTruck(ctx context.Context, truck *proto.Truck) (*proto.Truck, error) {
	colQurier := bson.M{"id":truck.Id}
	change := bson.M{"$set":bson.M{"licensePlate":truck.LicensePlate,"clockedInUser":truck.ClockedInUser,"miles":truck.Miles,"fleetCompanyId":truck.FleetCompanyId,"corporationId":truck.CorporationId,"regionId":truck.RegionId,"districtId":truck.DistrictId,"locationId":truck.LocationId}}
	if err := c.truckCollection().Update(colQurier,change); err != nil{
		return nil,err
	} else {
		return truck, nil
	}
}

func (c *TruckServiceRepository) GetTruckById(ctx context.Context, truckId string) (*proto.Truck, error) {
	truck := &proto.Truck{}
	if err := c.truckCollection().Find(bson.M{"id":truckId}).One(truck); err != nil{
		return nil, err
	} else {
		return truck,nil
	}
}

func (c *TruckServiceRepository) GetAllTrucksByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([] *proto.Truck, error) {
	var trucks [] *proto.Truck
	if err := c.truckCollection().Find(bson.M{"id":fleetCompanyId}).All(&trucks); err !=nil{
		return nil, err
	} else {
		return trucks, nil
	}
}

func (c *TruckServiceRepository) ClockIn(ctx context.Context, operation *proto.ClockOperation) (*proto.Truck, error) {
	if truck,err := c.GetTruckById(ctx, operation.TruckId); err != nil{
		return nil,err
	} else {
		if truck.ClockedInUser == "" {
			truck.ClockedInUser= operation.DriverId
			if truck,err := c.UpdateTruck(ctx,truck); err != nil{
				return nil,err
			} else {
				return truck,nil
			}
		} else {
			return truck,nil
		}
	}
}

func (c *TruckServiceRepository) ClockOut(ctx context.Context, operation *proto.ClockOperation) (*proto.Truck, error) {
	if truck, err := c.GetTruckById(ctx,operation.TruckId); err != nil{
		return nil,err
	} else {
		if truck.ClockedInUser == operation.DriverId {
			truck.ClockedInUser=""
			if truck,err := c.UpdateTruck(ctx,truck); err != nil{
				return nil,err
			} else {
				return truck,nil
			}
		} else {
			return truck, nil
		}
	}
}