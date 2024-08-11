package persistence

import (
	"context"
	"go-poc/internal/domain/abstractions"
	"go-poc/internal/domain/entities"

	"gorm.io/gorm"
)

type inboxRepository struct {
	db *gorm.DB
}

func NewInboxRepository(db *gorm.DB) abstractions.InboxRepository {
	return &inboxRepository{
		db: db,
	}
}

func (rep *inboxRepository) Any(ctx context.Context, messageId string) bool {
	var inboxMessage entities.InboxMessage
	if result := rep.db.Where("message_id = ?", messageId).First(&inboxMessage); result.Error != nil {
		return false
	}
	return true
}
