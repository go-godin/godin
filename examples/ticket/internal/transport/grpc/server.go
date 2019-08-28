// Code generated by godin. DO NOT EDIT.
package grpc

import (
	"context"
	"errors"
	"github.com/go-kit/kit/transport/grpc"

	pb "github.com/go-godin/ticket-service/api"

	"github.com/go-godin/godin/examples/ticket/internal/endpoint"
)

type TicketServiceServer struct {
	CreateHandler grpc.Handler
	GetHandler    grpc.Handler
}

func NewTicketServiceServer(endpoints endpoint.Set, options ...grpc.ServerOption) pb.TicketServiceServer {
	return &TicketServiceServer{

		CreateHandler: grpc.NewServer(
			endpoints.CreateEndpoint,
			DecodeCreateRequest,
			EncodeCreateResponse,
			options...,
		),
		GetHandler: grpc.NewServer(
			endpoints.GetEndpoint,
			DecodeGetRequest,
			EncodeGetResponse,
			options...,
		),
	}
}

// Create is the actual handler implementation of the RPC defined in the protobuf.
func (srv *TicketServiceServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	_, response, err := srv.CreateHandler.ServeGRPC(ctx, request)
	if err != nil {
		// TODO: encode error
		return nil, err
	}
	return response.(*pb.CreateResponse), nil
}

// Get is the actual handler implementation of the RPC defined in the protobuf.
func (srv *TicketServiceServer) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	_, response, err := srv.GetHandler.ServeGRPC(ctx, request)
	if err != nil {
		// TODO: encode error
		return nil, err
	}
	return response.(*pb.GetResponse), nil
}

// DecodeCreateRequest is responsible of decoding an incoming protobuf CreateRequest message
// into a domain-specific struct.
func DecodeCreateRequest(ctx context.Context, pbRequest interface{}) (interface{}, error) {
	if pbRequest == nil {
		return nil, errors.New("nil CreateRequest")
	}
	req := pbRequest.(*pb.CreateRequest)
	request, err := CreateRequestDecoder(req)
	if err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeCreateResponse is responsible of encoding outgoing CreateResponse structs
// into a protobuf message.
func EncodeCreateResponse(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil CreateResponse")
	}
	res := response.(endpoint.CreateResponse)
	pbResponse, err := CreateResponseEncoder(res)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil
}

// DecodeGetRequest is responsible of decoding an incoming protobuf GetRequest message
// into a domain-specific struct.
func DecodeGetRequest(ctx context.Context, pbRequest interface{}) (interface{}, error) {
	if pbRequest == nil {
		return nil, errors.New("nil GetRequest")
	}
	req := pbRequest.(*pb.GetRequest)
	request, err := GetRequestDecoder(req)
	if err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeGetResponse is responsible of encoding outgoing GetResponse structs
// into a protobuf message.
func EncodeGetResponse(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetResponse")
	}
	res := response.(endpoint.GetResponse)
	pbResponse, err := GetResponseEncoder(res)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil
}