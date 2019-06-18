package acceptance_tests

import (
	"context"
	"fleet-backend/common"
	proto2 "fleet-backend/common/proto"
	"fleet-backend/customer-service/proto"
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
	t.Log(driverResponse.Driver)

	driver = driverResponse.Driver

	driverResponse, err = customerServiceClient().UpdateDriver(context.Background(),&proto.Driver{
		Id:driver.Id,
		Name:driver.Name,
		Email:"asdfg@gmail.com",
		Password:"asdfghjkl",
		FleetCompanyId:"5d036710edf563627a833d61",
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(driverResponse.Driver)

	driverResponse, err =customerServiceClient().GetDriverById(context.Background(),&proto2.IdRequest{
		Id:driver.Id,
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(driverResponse.Driver)

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
	t.Log(driverResponse.Driver)

	driverResponses,err := customerServiceClient().GetDriversByFleetCompanyId(context.Background(), &proto2.IdRequest{
		Id:driver.FleetCompanyId,
	})
	then.AssertThat(t,err,is.Nil())
	//then.AssertThat(t,err,has.Length(2))
	t.Log(driverResponses.Drivers)



	corporationResponse,err := customerServiceClient().CreateCorporation(context.Background(), &proto.Corporation{
		Name:"All On Block",
		FleetCompanyId:driver.FleetCompanyId,
	})
	corporation := corporationResponse.Corporation
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,corporation.Name=="",is.False())
	then.AssertThat(t,corporation.FleetCompanyId=="",is.False())
	t.Log(corporationResponse.Corporation)


	corporationResponse,err = customerServiceClient().GetCorporationById(context.Background(), &proto2.IdRequest{
		Id:corporation.Id,
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(corporationResponse.Corporation)

	corporationResponses,err := customerServiceClient().GetAllCorporationsByFleetCompanyId(context.Background(), &proto2.IdRequest{
		Id:driver.FleetCompanyId,
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(corporationResponses.Corporations)

	regionResponse,err := customerServiceClient().CreateRegion(context.Background(), &proto.Region{
		Name:"AOB Jaipur",
		CorporationId:corporationResponse.Corporation.Id,
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(regionResponse.Region)
	regionResponse,err = customerServiceClient().GetRegionById(context.Background(), &proto2.IdRequest{
		Id:regionResponse.Region.Id,
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(regionResponse.Region)

	regionResponses,err := customerServiceClient().GetAllRegionsByCorporationId(context.Background(), &proto2.IdRequest{
		Id:corporationResponse.Corporation.Id,
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(regionResponses.Regions)

	dristrictResponse, err := customerServiceClient().CreateDistrict(context.Background(), &proto.District{
		Name:"Jaipur",
		RegionId:regionResponse.Region.Id,
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(dristrictResponse.District)
	dristrictResponse,err = customerServiceClient().GetDistrictById(context.Background(), &proto2.IdRequest{
		Id:dristrictResponse.District.Id,
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(dristrictResponse.District)

	dristrictResponses,err := customerServiceClient().GetAllDistrictsByRegionId(context.Background(), &proto2.IdRequest{
		Id:regionResponse.Region.Id,
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(dristrictResponses.Districts)


	locationResponse, err := customerServiceClient().CreateLocation(context.Background(), &proto.Location{
		Name:"Jaipur",
		DistrictId:dristrictResponse.District.Id,
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(locationResponse.Location)
	locationResponse,err = customerServiceClient().GetLocationById(context.Background(), &proto2.IdRequest{
		Id:locationResponse.Location.Id,
	})
	then.AssertThat(t,err,is.Nil())
	t.Log(locationResponse.Location)

	locationResponses,err := customerServiceClient().GetAllLocationsByDistrictId(context.Background(), &proto2.IdRequest{
		Id:dristrictResponse.District.Id,
	})
	then.AssertThat(t,err,is.Nil())
	//then.AssertThat(t,err,has.Length(2))
	t.Log(locationResponses.Locations)

}

func customerServiceClient() proto.CustomerService{
	return proto.NewCustomerService("customer-service",client.NewClient(common.UseConsul))
}


