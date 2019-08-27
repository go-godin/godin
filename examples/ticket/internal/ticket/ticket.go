package ticket

import (
	"context"
	"github.com/rs/xid"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Create(ctx context.Context, title string, description string) (ticket *Ticket, err error) {

	ticket = &Ticket{}
	ticket.Id = xid.New().String()
	ticket.Title = title
	ticket.Status = Status_OPEN

	return ticket, nil
}

func (s service) Get(ctx context.Context, id string) (ticket *Ticket, err error) {
	panic("implement me")
}
