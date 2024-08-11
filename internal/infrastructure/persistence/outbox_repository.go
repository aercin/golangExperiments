package persistence

import (
	"context"
	"go-poc/internal/domain/abstractions"
	"go-poc/internal/domain/entities"

	"gorm.io/gorm"
)

type outboxRepository struct {
	db *gorm.DB
}

func NewOutboxRepository(db *gorm.DB) abstractions.OutboxRepository {
	return &outboxRepository{
		db: db,
	}
}

func (rep *outboxRepository) Find(ctx context.Context, query string) ([]entities.OutboxMessage, error) {

	var outboxMessages []entities.OutboxMessage

	if err := rep.db.WithContext(ctx).Raw(query).Find(&outboxMessages).Error; err != nil {
		return nil, err
	}

	return outboxMessages, nil
}

func (rep *outboxRepository) Delete(ctx context.Context, id int64) error {
	if err := rep.db.WithContext(ctx).Delete(&entities.OutboxMessage{}, id).Error; err != nil {
		return err
	}
	return nil
}
