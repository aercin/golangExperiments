package interactor

import (
	"go-poc/configs"
	api_abstraction "go-poc/internal/api/abstractions"
	v1 "go-poc/internal/api/v1"
	app_abstraction "go-poc/internal/application/abstractions"
	"go-poc/internal/infrastructure/persistence"
	"go-poc/internal/infrastructure/services"
	"go-poc/pkg/logrus"
	"go-poc/pkg/rabbitMQ"
)

func ResolveHandler(config *configs.Config) api_abstraction.Handlers {
	db := persistence.NewDbConnection(config)
	orderRepository := persistence.NewOrderRepository(db)
	inboxRepository := persistence.NewInboxRepository(db)
	orderService := services.NewOrderService(orderRepository, inboxRepository, config)
	logger, _ := logrus.NewLogger(logrus.Info, logrus.NewFileHook(config.Log.Path))
	return v1.NewHandler(orderService, logger)
}

func ResolveEventDispatcher(config *configs.Config) app_abstraction.EventDispatcher {
	db := persistence.NewDbConnection(config)
	outboxRepo := persistence.NewOutboxRepository(db)
	amqpChan, _ := rabbitMQ.InitRabbitMQ(config)
	producer, _ := rabbitMQ.NewProducer(amqpChan, config)
	return services.NewEventDispatcher(outboxRepo, producer)
}

func ResolveOrderService(config *configs.Config) app_abstraction.OrderService {
	db := persistence.NewDbConnection(config)
	orderRepo := persistence.NewOrderRepository(db)
	inboxRepo := persistence.NewInboxRepository(db)
	return services.NewOrderService(orderRepo, inboxRepo, config)
}
