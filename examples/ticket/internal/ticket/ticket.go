package ticket

import (
	"context"
	"fmt"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type service struct {
	repo Repository
	logger *zap.Logger
}

func NewService(repo Repository, logger *zap.Logger) Service {
	return &service{
		repo: repo,
		logger: logger,
	}
}

func (s service) Create(ctx context.Context, title string, description string) (ticket *Ticket, err error) {

	ticket = &Ticket{}
	ticket.Id = fmt.Sprintf( "ticket_%s", xid.New().String())
	ticket.Title = title
	ticket.Status = Status_OPEN

	s.logger.Info("ticket created",
		zap.String("ticket.Id", ticket.Id))

	return ticket, nil
}

func (s service) Get(ctx context.Context, id string) (ticket *Ticket, err error) {
	panic("implement me")
}
