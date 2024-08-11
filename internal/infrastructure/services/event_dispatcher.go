package services

import (
	"context"
	app_abstraction "go-poc/internal/application/abstractions"
	repo_abstraction "go-poc/internal/domain/abstractions"
	"go-poc/pkg/rabbitMQ"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type eventDispatcher struct {
	outboxRepo repo_abstraction.OutboxRepository
	producer   rabbitMQ.Producer
}

func NewEventDispatcher(outboxRepo repo_abstraction.OutboxRepository, producer rabbitMQ.Producer) app_abstraction.EventDispatcher {
	return &eventDispatcher{
		outboxRepo: outboxRepo,
		producer:   producer,
	}
}

func (ed *eventDispatcher) DispatchEvents(ctx context.Context) error {

	query, _, err := goqu.From("outbox_messages").Order(goqu.I("created_on").Asc()).Limit(100).ToSQL()
	if err != nil {
		return err
	}

	outboxMessages, err := ed.outboxRepo.Find(ctx, query)
	if err != nil {
		return err
	}

	for _, outboxMessage := range outboxMessages {

		err := ed.producer.PublishMessage(ctx, []byte(outboxMessage.Message))

		if err == nil {
			ed.outboxRepo.Delete(ctx, outboxMessage.Id)
		}
	}

	return nil
}
