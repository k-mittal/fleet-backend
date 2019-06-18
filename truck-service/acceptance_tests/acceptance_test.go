package acceptance_tests

import (
	"context"
	"fleet-backend/common"
	proto2 "fleet-backend/common/proto"
	proto3 "fleet-backend/customer-service/proto"
	"fleet-backend/truck-service/proto"
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/micro/go-micro/client"
	"testing"
)

func TestTruck(t *testing.T){

	customerResponse, err :=customerServiceClient().SignUp(context.Background(), &proto3.SignUpRequest{
		FleetCompanyName:"asd",
		Name:"komal",
		Email:"komal@allonblock.com",
		Password:"asdswqew",
	})
	then.AssertThat(t,err,is.Nil())
	driver := customerResponse.Driver
	then.AssertThat(t,driver.Id=="",is.False())
	then.AssertThat(t,driver.Name=="",is.False())
	then.AssertThat(t,driver.Email=="",is.False())
	then.AssertThat(t,driver.Password=="",is.False())

	corporationResponse,err := customerServiceClient().CreateCorporation(context.Background(), &proto3.Corporation{
		Name:"All On Block",
		FleetCompanyId:driver.FleetCompanyId,
	})
	corporation := corporationResponse.Corporation
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,corporation.Name=="",is.False())
	then.AssertThat(t,corporation.FleetCompanyId=="",is.False())

	regionResponse,err := customerServiceClient().CreateRegion(context.Background(), &proto3.Region{
		Name:"AOB Jaipur",
		CorporationId:corporationResponse.Corporation.Id,
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,regionResponse.Region.CorporationId=="",is.False())

	districtResponse, err := customerServiceClient().CreateDistrict(context.Background(), &proto3.District{
		Name:"Jaipur",
		RegionId:regionResponse.Region.Id,
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,districtResponse.District.RegionId=="",is.False())

	locationResponse, err := customerServiceClient().CreateLocation(context.Background(), &proto3.Location{
		Name:"Jaipur",
		DistrictId:districtResponse.District.Id,
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,locationResponse.Location.DistrictId=="",is.False())



	response,err := truckServiceClient().CreateTruck(context.Background(),&proto.Truck{
		LicensePlate:"weqeq2312",
		ClockedInUser:"",
		Miles:100.00,
		FleetCompanyId:driver.FleetCompanyId,
		CorporationId:corporationResponse.Corporation.Id,
		RegionId:regionResponse.Region.Id,
		DistrictId:districtResponse.District.Id,
		LocationId:locationResponse.Location.Id,
	})
	truck := response.Truck
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,truck.LicensePlate=="",is.False())
	then.AssertThat(t,truck.Miles==100.00,is.True())
	then.AssertThat(t,truck.FleetCompanyId=="",is.False())
	then.AssertThat(t,truck.CorporationId=="",is.False())
	then.AssertThat(t,truck.RegionId=="",is.False())
	then.AssertThat(t,truck.DistrictId=="",is.False())
	then.AssertThat(t,truck.LocationId=="",is.False())
	 
	response,err = truckServiceClient().CreateTruck(context.Background(),&proto.Truck{
		LicensePlate:"weqeq2312",
		ClockedInUser:"",
		Miles:123.23,
		FleetCompanyId:driver.FleetCompanyId,
		CorporationId:corporationResponse.Corporation.Id,
		RegionId:regionResponse.Region.Id,
		DistrictId:districtResponse.District.Id,
		LocationId:locationResponse.Location.Id,
	})
	truck = response.Truck
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,truck.LicensePlate=="",is.False())
	then.AssertThat(t,truck.Miles==123.23,is.True())
	then.AssertThat(t,truck.FleetCompanyId=="",is.False())
	then.AssertThat(t,truck.CorporationId=="",is.False())
	then.AssertThat(t,truck.RegionId=="",is.False())
	then.AssertThat(t,truck.DistrictId=="",is.False())
	then.AssertThat(t,truck.LocationId=="",is.False())

	response,err = truckServiceClient().GetTruckById(context.Background(),&proto2.IdRequest{
		Id:truck.Id,
	})
	then.AssertThat(t,err,is.Nil())

	responses,err := truckServiceClient().GetAllTrucksByFleetCompanyId(context.Background(),&proto2.IdRequest{
		Id:truck.FleetCompanyId,
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,responses.Trucks,has.Length(2))
	//t.Log(responses.Trucks)

	response,err = truckServiceClient().ClockIn(context.Background(),&proto.ClockOperation{
		DriverId:driver.Id,
		TruckId:truck.Id,
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,response.Truck.ClockedInUser,is.EqualTo(driver.Id))

	response,err = truckServiceClient().ClockOut(context.Background(),&proto.ClockOperation{
		DriverId:driver.Id,
		TruckId:truck.Id,
	})
	then.AssertThat(t,err,is.Nil())
	then.AssertThat(t,response.Truck.ClockedInUser,is.EqualTo(""))

}

func truckServiceClient() proto.TruckService{
	return proto.NewTruckService("truck-service",client.NewClient(common.UseConsul))
}

func customerServiceClient() proto3.CustomerService{
	return proto3.NewCustomerService("customer-service",client.NewClient(common.UseConsul))
}
