// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/service1/service1.proto

package go_micro_api_service1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Service1 service

func NewService1Endpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Service1 service

type Service1Service interface {
	Json(ctx context.Context, in *Empty, opts ...client.CallOption) (*Response, error)
}

type service1Service struct {
	c    client.Client
	name string
}

func NewService1Service(name string, c client.Client) Service1Service {
	return &service1Service{
		c:    c,
		name: name,
	}
}

func (c *service1Service) Json(ctx context.Context, in *Empty, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Service1.Json", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service1 service

type Service1Handler interface {
	Json(context.Context, *Empty, *Response) error
}

func RegisterService1Handler(s server.Server, hdlr Service1Handler, opts ...server.HandlerOption) error {
	type service1 interface {
		Json(ctx context.Context, in *Empty, out *Response) error
	}
	type Service1 struct {
		service1
	}
	h := &service1Handler{hdlr}
	return s.Handle(s.NewHandler(&Service1{h}, opts...))
}

type service1Handler struct {
	Service1Handler
}

func (h *service1Handler) Json(ctx context.Context, in *Empty, out *Response) error {
	return h.Service1Handler.Json(ctx, in, out)
}
