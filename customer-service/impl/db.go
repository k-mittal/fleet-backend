package impl

import (
	"context"
	"fleet-backend/common"
	"fleet-backend/customer-service/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName  = "customer"
const driverCollection= "drivers"
const fleetcompanyCollection ="fleetcompany"
const corporationCollection  = "corporation"
const regionCollection = "region"
const districtCollection = "district"
const locationCollection ="location"

type CustomerServiceRepository struct {
	fleetcompanySession *mgo.Session
	dbName string
}

func NewCustomerServiceRepository() (*CustomerServiceRepository, error) {
	if session, err := common.ConnectToMongo(); err != nil {
		return nil, err
	} else {
		return &CustomerServiceRepository{
			fleetcompanySession: session,
		}, err
	}
}


func (c *CustomerServiceRepository) AddFleetCompany(ctx context.Context, fleetcompany *proto.FleetCompany) error{
	if err:=  c.fleetcompanyCollection().Insert(fleetcompany); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) AddDriver(ctx context.Context, driver *proto.Driver) error{
	if err:=  c.driverCollection().Insert(driver); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) UpdateDriver(ctx context.Context, driver *proto.Driver) (*proto.Driver,error) {
	colQuerier := bson.M{"id":driver.Id}
	change := bson.M{"$set": bson.M{"name":driver.Name,"email":driver.Email,"password":driver.Password,"fleetcompanyid":driver.FleetCompanyId}}
	if err:=c.driverCollection().Update(colQuerier,change); err !=nil{
		return nil,err
	} else {
		return driver, nil
	}
}

func (c *CustomerServiceRepository) GetDriverById(ctx context.Context, driveId string) (*proto.Driver,error) {
	driver := &proto.Driver{}
	if err := c.driverCollection().Find(bson.M{"id":driveId}).One(driver); err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

func (c *CustomerServiceRepository) GetDriversByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([]*proto.Driver,error) {
	var drivers [] *proto.Driver
	if err := c.driverCollection().Find(bson.M{"fleetcompanyid":fleetCompanyId}).All(&drivers); err != nil {
		return nil, err
	} else {
		return drivers, nil
	}
}

func (c *CustomerServiceRepository) driverCollection() *mgo.Collection{
	return c.fleetcompanySession.DB(dbName).C(driverCollection)
}

func (c *CustomerServiceRepository) fleetcompanyCollection() *mgo.Collection{
	return c.fleetcompanySession.DB(dbName).C(fleetcompanyCollection)
}

func (c *CustomerServiceRepository) Close(){
	c.fleetcompanySession.Close()
}


//Corporation
func (c *CustomerServiceRepository) AddCorporation(ctx context.Context, corporation *proto.Corporation) error {
	if err := c.corporationCollection().Insert(corporation); err !=nil {
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) GetAllCorporationsByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([] *proto.Corporation, error) {
	var corporations [] *proto.Corporation
	if err := c.corporationCollection().Find(bson.M{"fleetcompanyid":fleetCompanyId}).All(&corporations); err != nil{
		return nil, err
	} else {
		return corporations, nil
	}
}



func (c *CustomerServiceRepository) corporationCollection() *mgo.Collection {
	return c.fleetcompanySession.DB(dbName).C(corporationCollection)
}

func (c *CustomerServiceRepository) GetCorporationById(ctx context.Context, corporationId string) (*proto.Corporation, error) {
	corporation := &proto.Corporation{}
	if err := c.corporationCollection().Find(bson.M{"id":corporationId}).One(&corporation); err !=nil{
		return nil, err
	} else {
		return corporation, nil
	}
}


//Region

func (c *CustomerServiceRepository) RegionCollection() *mgo.Collection {
	return c.fleetcompanySession.DB(dbName).C(regionCollection)
}

func (c *CustomerServiceRepository) AddRegion(ctx context.Context, region *proto.Region) error {
	if err := c.RegionCollection().Insert(region); err != nil{
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) GetAllRegionsByCorporationId(ctx context.Context, corporationId string) ([] *proto.Region, error) {
	var regions [] *proto.Region
	if err := c.RegionCollection().Find(bson.M{"corporationid":corporationId}).All(&regions); err != nil{
		return nil, err
	} else {
		return regions, nil
	}
}

func (c *CustomerServiceRepository) GetRegionById(ctx context.Context, regionId string) (*proto.Region, error) {
	region := &proto.Region{}
	if err := c.RegionCollection().Find(bson.M{"id":regionId}).One(region); err != nil{
		return nil, err
	} else {
		return region, nil
	}
}


//District

func (c *CustomerServiceRepository) AddDistrict(ctx context.Context, district *proto.District) error {
	if err := c.DistrictCollecion().Insert(district); err != nil{
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) DistrictCollecion() *mgo.Collection {
	return c.fleetcompanySession.DB(dbName).C(districtCollection)
}

func (c *CustomerServiceRepository) GetAllDistrictByRegionId(ctx context.Context, regionId string) ([] *proto.District, error) {
	var districts [] *proto.District
	if err := c.DistrictCollecion().Find(bson.M{"regionid":regionId}).All(&districts); err != nil{
		return nil,err
	} else {
		return districts, nil
	}
}

func (c *CustomerServiceRepository) GetDistrictById(ctx context.Context, districtId string) (*proto.District, error) {
	district := &proto.District{}
	if err := c.DistrictCollecion().Find(bson.M{"id":districtId}).One(district); err != nil{
		return nil,err
	} else {
		return district, nil
	}
}


//Location

func (c *CustomerServiceRepository) AddLocation(ctx context.Context, location *proto.Location) error {
	if err := c.LocationCollecion().Insert(location); err != nil{
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) LocationCollecion() *mgo.Collection {
	return c.fleetcompanySession.DB(dbName).C(locationCollection)
}

func (c *CustomerServiceRepository) GetAllLocationsByDistrictId(ctx context.Context, districtId string) ([] *proto.Location, error) {
	var locations [] *proto.Location
	if err := c.LocationCollecion().Find(bson.M{"districtid":districtId}).All(&locations); err != nil{
		return nil,err
	} else {
		return locations, nil
	}
}

func (c *CustomerServiceRepository) GetLocationById(ctx context.Context, locationId string) (*proto.Location, error) {
	location := &proto.Location{}
	if err := c.LocationCollecion().Find(bson.M{"id":locationId}).One(location); err != nil{
		return nil,err
	} else {
		return location, nil
	}
}