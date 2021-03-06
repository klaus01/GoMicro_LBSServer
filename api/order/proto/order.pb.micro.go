// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/order/proto/order.proto

package api_order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Client API for Order service

type OrderService interface {
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*OrderModel, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResult, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResult, error)
	SetDeliveryInfo(ctx context.Context, in *SetDeliveryInfoRequest, opts ...client.CallOption) (*empty.Empty, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*OrderModel, error) {
	req := c.c.NewRequest(c.name, "Order.Get", in)
	out := new(OrderModel)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResult, error) {
	req := c.c.NewRequest(c.name, "Order.Search", in)
	out := new(SearchResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResult, error) {
	req := c.c.NewRequest(c.name, "Order.Create", in)
	out := new(CreateResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) SetDeliveryInfo(ctx context.Context, in *SetDeliveryInfoRequest, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "Order.SetDeliveryInfo", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	Get(context.Context, *GetRequest, *OrderModel) error
	Search(context.Context, *SearchRequest, *SearchResult) error
	Create(context.Context, *CreateRequest, *CreateResult) error
	SetDeliveryInfo(context.Context, *SetDeliveryInfoRequest, *empty.Empty) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		Get(ctx context.Context, in *GetRequest, out *OrderModel) error
		Search(ctx context.Context, in *SearchRequest, out *SearchResult) error
		Create(ctx context.Context, in *CreateRequest, out *CreateResult) error
		SetDeliveryInfo(ctx context.Context, in *SetDeliveryInfoRequest, out *empty.Empty) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) Get(ctx context.Context, in *GetRequest, out *OrderModel) error {
	return h.OrderHandler.Get(ctx, in, out)
}

func (h *orderHandler) Search(ctx context.Context, in *SearchRequest, out *SearchResult) error {
	return h.OrderHandler.Search(ctx, in, out)
}

func (h *orderHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResult) error {
	return h.OrderHandler.Create(ctx, in, out)
}

func (h *orderHandler) SetDeliveryInfo(ctx context.Context, in *SetDeliveryInfoRequest, out *empty.Empty) error {
	return h.OrderHandler.SetDeliveryInfo(ctx, in, out)
}
