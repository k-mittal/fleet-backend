package acceptance_tests

import (
	"context"
	"fleet-backend/common"
	proto2 "fleet-backend/common/proto"
	"fleet-backend/customer-service/proto"
	"github.com/corbym/gocrest/has"

	//"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/micro/go-micro/client"
	"testing"

	//"github.com/micro/go-micro/registry/consul"
)

func TestDriver(t *testing.T){
	response, err :=customerServiceClient().SignUp(context.Background(), &proto.SignUpRequest{
		FleetCompanyName:"asd",
		Name:"komal",
		Email:"komal@allonblock.com",
		Password:"asdswqew",
	})
	then.AssertThat(t,err,is.Nil())
	driver := response.Driver
	then.AssertThat(t,driver.Id=="",is.False())
	then.AssertThat(t,driver.Name=="",is.False())
	then.AssertThat(t,driver.Email=="",is.False())
	then.AssertThat(t,driver.Password=="",is.False())


	driverResponse, err :=customerServiceClient().CreateDriver(context.Background(), &proto.Driver{
		Name:"driver",
		Email:"driver@gmail.com",
		Password:"password",
		FleetCompanyId:"5d036710edf563627a833d61",
	})
	driver = driverResponse.Driver
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,driver.Name=="",is.False())
	then.AssertThat(t,driver.Email=="",is.False())
	then.AssertThat(t,driver.Password=="",is.False())
	then.AssertThat(t,driver.FleetCompanyId=="",is.False())


	driver = driverResponse.Driver

	driverResponse, err = customerServiceClient().UpdateDriver(context.Background(),&proto.Driver{
		Id:driver.Id,
		Name:driver.Name,
		Email:"asdfg@gmail.com",
		Password:"asdfghjkl",
		FleetCompanyId:"5d036710edf563627a833d61",
	})
	then.AssertThat(t,err,is.Nil())


	driverResponse, err =customerServiceClient().GetDriverById(context.Background(),&proto2.IdRequest{
		Id:driver.Id,
	})
	then.AssertThat(t,err,is.Nil())


	driverResponse, err =customerServiceClient().CreateDriver(context.Background(), &proto.Driver{
		Name:"driver",
		Email:"driver@gmail.com",
		Password:"password",
		FleetCompanyId:"5d036710edf563627a833d61",
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,driver.Name=="",is.False())
	then.AssertThat(t,driver.Email=="",is.False())
	then.AssertThat(t,driver.Password=="",is.False())
	then.AssertThat(t,driver.FleetCompanyId=="",is.False())


	driverResponses,err := customerServiceClient().GetDriversByFleetCompanyId(context.Background(), &proto2.IdRequest{
		Id:driver.FleetCompanyId,
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,driverResponses.Drivers,has.Length(2))
	//t.Log(driverResponses.Drivers)



	corporationResponse,err := customerServiceClient().CreateCorporation(context.Background(), &proto.Corporation{
		Name:"All On Block",
		FleetCompanyId:driver.FleetCompanyId,
	})
	corporation := corporationResponse.Corporation
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,corporation.Name=="",is.False())
	then.AssertThat(t,corporation.FleetCompanyId=="",is.False())



	corporationResponse,err = customerServiceClient().GetCorporationById(context.Background(), &proto2.IdRequest{
		Id:corporation.Id,
	})
	then.AssertThat(t,err,is.Nil())


	corporationResponses,err := customerServiceClient().GetAllCorporationsByFleetCompanyId(context.Background(), &proto2.IdRequest{
		Id:driver.FleetCompanyId,
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,corporationResponses.Corporations,has.Length(1))
	//t.Log(corporationResponses.Corporations)

	regionResponse,err := customerServiceClient().CreateRegion(context.Background(), &proto.Region{
		Name:"AOB Jaipur",
		CorporationId:corporationResponse.Corporation.Id,
	})
	then.AssertThat(t,err,is.Nil())

	regionResponse,err = customerServiceClient().GetRegionById(context.Background(), &proto2.IdRequest{
		Id:regionResponse.Region.Id,
	})
	then.AssertThat(t,err,is.Nil())

	regionResponses,err := customerServiceClient().GetAllRegionsByCorporationId(context.Background(), &proto2.IdRequest{
		Id:corporationResponse.Corporation.Id,
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,regionResponses.Regions,has.Length(1))

	//t.Log(regionResponses.Regions)

	dristrictResponse, err := customerServiceClient().CreateDistrict(context.Background(), &proto.District{
		Name:"Jaipur",
		RegionId:regionResponse.Region.Id,
	})
	then.AssertThat(t,err,is.Nil())

	dristrictResponse,err = customerServiceClient().GetDistrictById(context.Background(), &proto2.IdRequest{
		Id:dristrictResponse.District.Id,
	})
	then.AssertThat(t,err,is.Nil())


	dristrictResponses,err := customerServiceClient().GetAllDistrictsByRegionId(context.Background(), &proto2.IdRequest{
		Id:regionResponse.Region.Id,
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,dristrictResponses.Districts,has.Length(1))

	//t.Log(dristrictResponses.Districts)


	locationResponse, err := customerServiceClient().CreateLocation(context.Background(), &proto.Location{
		Name:"Jaipur",
		DistrictId:dristrictResponse.District.Id,
	})
	then.AssertThat(t,err,is.Nil())

	locationResponse,err = customerServiceClient().GetLocationById(context.Background(), &proto2.IdRequest{
		Id:locationResponse.Location.Id,
	})
	then.AssertThat(t,err,is.Nil())


	locationResponses,err := customerServiceClient().GetAllLocationsByDistrictId(context.Background(), &proto2.IdRequest{
		Id:dristrictResponse.District.Id,
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,locationResponses.Locations,has.Length(1))
	//t.Log(locationResponses.Locations)

}

func customerServiceClient() proto.CustomerService{
	return proto.NewCustomerService("customer-service",client.NewClient(common.UseConsul))
}


