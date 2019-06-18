package impl

import (
	"context"
	"fleet-backend/truck-service/proto"
	"github.com/google/uuid"
)

type Service struct {
	Repository *TruckServiceRepository
}


func (service Service) CreateTruck(ctx context.Context, request *proto.Truck) (*proto.Truck, error) {
	truck := &proto.Truck{
		Id:uuid.New().String(),
		LicensePlate:request.LicensePlate,
		ClockedInUser:request.ClockedInUser,
		Miles:request.Miles,
		FleetCompanyId:request.FleetCompanyId,
		CorporationId:request.CorporationId,
		RegionId:request.RegionId,
		DistrictId:request.DistrictId,
		LocationId:request.LocationId,
	}
	if err := service.Repository.AddTruck(ctx,truck); err !=nil {
		return nil,err
	} else {
		return truck,nil
	}
}

func (service Service) UpdateTruck(ctx context.Context, request *proto.Truck) (*proto.Truck, error) {
	if truck,err := service.Repository.UpdateTruck(ctx,request); err != nil{
		return nil,err
	} else {
		return truck,nil
	}
}

func (service Service) GetTruckById(ctx context.Context, truckId string) (*proto.Truck,error) {
	if truck,err := service.Repository.GetTruckById(ctx,truckId); err !=nil {
		return nil,err
	} else {
		return truck,nil
	}
}

func (service Service) GetAllTrucksByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([] *proto.Truck, error) {
	if trucks,err := service.Repository.GetAllTrucksByFleetCompanyId(ctx, fleetCompanyId); err != nil{
		return nil, err
	} else {
		return trucks, nil
	}
}

func (service Service) ClockIn(ctx context.Context, operation *proto.ClockOperation) (*proto.Truck, error) {
	if truck,err := service.Repository.ClockIn(ctx,operation); err != nil{
		return nil,err
	} else {
		return truck,nil
	}
}
func (service Service) ClockOut(ctx context.Context, operation *proto.ClockOperation) (*proto.Truck, error) {
	if truck,err := service.Repository.ClockOut(ctx,operation); err != nil{
		return nil,err
	} else {
		return truck,nil
	}
}
