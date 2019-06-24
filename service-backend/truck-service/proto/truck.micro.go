// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: truck.proto

package proto

import (
	proto1 "fleet-backend/common/proto"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TruckService service

type TruckService interface {
	CreateTruck(ctx context.Context, in *Truck, opts ...client.CallOption) (*TruckResponse, error)
	UpdateTruck(ctx context.Context, in *Truck, opts ...client.CallOption) (*TruckResponse, error)
	GetTruckById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*TruckResponse, error)
	GetAllTrucksByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*TrucksResponse, error)
	ClockIn(ctx context.Context, in *ClockOperation, opts ...client.CallOption) (*TruckResponse, error)
	ClockOut(ctx context.Context, in *ClockOperation, opts ...client.CallOption) (*TruckResponse, error)
}

type truckService struct {
	c    client.Client
	name string
}

func NewTruckService(name string, c client.Client) TruckService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &truckService{
		c:    c,
		name: name,
	}
}

func (c *truckService) CreateTruck(ctx context.Context, in *Truck, opts ...client.CallOption) (*TruckResponse, error) {
	req := c.c.NewRequest(c.name, "TruckService.CreateTruck", in)
	out := new(TruckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *truckService) UpdateTruck(ctx context.Context, in *Truck, opts ...client.CallOption) (*TruckResponse, error) {
	req := c.c.NewRequest(c.name, "TruckService.UpdateTruck", in)
	out := new(TruckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *truckService) GetTruckById(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*TruckResponse, error) {
	req := c.c.NewRequest(c.name, "TruckService.GetTruckById", in)
	out := new(TruckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *truckService) GetAllTrucksByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, opts ...client.CallOption) (*TrucksResponse, error) {
	req := c.c.NewRequest(c.name, "TruckService.GetAllTrucksByFleetCompanyId", in)
	out := new(TrucksResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *truckService) ClockIn(ctx context.Context, in *ClockOperation, opts ...client.CallOption) (*TruckResponse, error) {
	req := c.c.NewRequest(c.name, "TruckService.ClockIn", in)
	out := new(TruckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *truckService) ClockOut(ctx context.Context, in *ClockOperation, opts ...client.CallOption) (*TruckResponse, error) {
	req := c.c.NewRequest(c.name, "TruckService.ClockOut", in)
	out := new(TruckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TruckService service

type TruckServiceHandler interface {
	CreateTruck(context.Context, *Truck, *TruckResponse) error
	UpdateTruck(context.Context, *Truck, *TruckResponse) error
	GetTruckById(context.Context, *proto1.IdRequest, *TruckResponse) error
	GetAllTrucksByFleetCompanyId(context.Context, *proto1.IdRequest, *TrucksResponse) error
	ClockIn(context.Context, *ClockOperation, *TruckResponse) error
	ClockOut(context.Context, *ClockOperation, *TruckResponse) error
}

func RegisterTruckServiceHandler(s server.Server, hdlr TruckServiceHandler, opts ...server.HandlerOption) error {
	type truckService interface {
		CreateTruck(ctx context.Context, in *Truck, out *TruckResponse) error
		UpdateTruck(ctx context.Context, in *Truck, out *TruckResponse) error
		GetTruckById(ctx context.Context, in *proto1.IdRequest, out *TruckResponse) error
		GetAllTrucksByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, out *TrucksResponse) error
		ClockIn(ctx context.Context, in *ClockOperation, out *TruckResponse) error
		ClockOut(ctx context.Context, in *ClockOperation, out *TruckResponse) error
	}
	type TruckService struct {
		truckService
	}
	h := &truckServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TruckService{h}, opts...))
}

type truckServiceHandler struct {
	TruckServiceHandler
}

func (h *truckServiceHandler) CreateTruck(ctx context.Context, in *Truck, out *TruckResponse) error {
	return h.TruckServiceHandler.CreateTruck(ctx, in, out)
}

func (h *truckServiceHandler) UpdateTruck(ctx context.Context, in *Truck, out *TruckResponse) error {
	return h.TruckServiceHandler.UpdateTruck(ctx, in, out)
}

func (h *truckServiceHandler) GetTruckById(ctx context.Context, in *proto1.IdRequest, out *TruckResponse) error {
	return h.TruckServiceHandler.GetTruckById(ctx, in, out)
}

func (h *truckServiceHandler) GetAllTrucksByFleetCompanyId(ctx context.Context, in *proto1.IdRequest, out *TrucksResponse) error {
	return h.TruckServiceHandler.GetAllTrucksByFleetCompanyId(ctx, in, out)
}

func (h *truckServiceHandler) ClockIn(ctx context.Context, in *ClockOperation, out *TruckResponse) error {
	return h.TruckServiceHandler.ClockIn(ctx, in, out)
}

func (h *truckServiceHandler) ClockOut(ctx context.Context, in *ClockOperation, out *TruckResponse) error {
	return h.TruckServiceHandler.ClockOut(ctx, in, out)
}