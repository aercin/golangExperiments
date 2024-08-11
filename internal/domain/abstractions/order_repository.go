package abstractions

import (
	"context"
	"go-poc/internal/domain/entities"
)

type OrderRepository interface {
	Get(ctx context.Context, query string) (entities.Order, error)

	Find(ctx context.Context, query string) ([]entities.Order, error)

	Create(ctx context.Context, order *entities.Order, outboxMsg *entities.OutboxMessage) error

	ChangeOrderStatus(ctx context.Context, order *entities.Order, inboxMsg *entities.InboxMessage) error
}
