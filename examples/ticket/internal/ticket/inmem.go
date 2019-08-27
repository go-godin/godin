package ticket

import (
	"context"
)

type InMemoryDatastore struct {
}

func NewInMemory() *InMemoryDatastore {
	return &InMemoryDatastore{}
}

func (ds *InMemoryDatastore) Get(ctx context.Context, ticketID string) (*Ticket, error) {
	return &Ticket{}, nil
}
