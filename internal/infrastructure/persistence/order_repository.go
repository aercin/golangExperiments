package persistence

import (
	"context"
	"go-poc/internal/domain/abstractions"
	"go-poc/internal/domain/entities"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) abstractions.OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (rep *OrderRepository) Get(ctx context.Context, query string) (entities.Order, error) {
	var order entities.Order

	if err := rep.db.WithContext(ctx).Raw(query).First(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (rep *OrderRepository) Find(ctx context.Context, query string) ([]entities.Order, error) {

	var orders []entities.Order

	if err := rep.db.WithContext(ctx).Preload("OrderProducts").Raw(query).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (rep *OrderRepository) Create(ctx context.Context, order *entities.Order, outboxMsg *entities.OutboxMessage) error {

	err := rep.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(order).Error; err != nil {
			return err
		}

		if err := tx.Create(&outboxMsg).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (rep *OrderRepository) ChangeOrderStatus(ctx context.Context, order *entities.Order, inboxMsg *entities.InboxMessage) error {

	err := rep.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		if err := tx.Create(&inboxMsg).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
