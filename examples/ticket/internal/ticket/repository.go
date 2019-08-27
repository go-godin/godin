package ticket

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, ticketID string) (*Ticket, error)
}
