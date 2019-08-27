package grpc

import (
	"github.com/go-godin/godin/examples/ticket/internal/endpoint"
	ticket_v1 "github.com/go-godin/ticket-service/api"
)

func CreateRequestDecoder(pbRequest *ticket_v1.CreateRequest) (request endpoint.CreateRequest, err error) {
	return endpoint.CreateRequest{
		Title:       pbRequest.Title,
		Description: pbRequest.Description,
	}, nil
}

func CreateResponseEncoder(response endpoint.CreateResponse) (pbResponse *ticket_v1.CreateResponse, err error) {
	return &ticket_v1.CreateResponse{
		Ticket: &ticket_v1.Ticket{
			Id:     response.Ticket.Id,
			Title:  response.Ticket.Title,
			Status: ticket_v1.Status(response.Ticket.Status),
		},
	}, nil
}

func GetRequestDecoder(pbRequest *ticket_v1.GetRequest) (request endpoint.GetRequest, err error) {
	return endpoint.GetRequest{}, nil
}

func GetResponseEncoder(response endpoint.GetResponse) (pbResponse *ticket_v1.GetResponse, err error) {
	return &ticket_v1.GetResponse{}, nil
}
