package rabbitMQ

import (
	"context"
	"go-poc/configs"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer interface {
	PublishMessage(ctx context.Context, msg []byte) error
}

type producer struct {
	ch            *amqp.Channel
	sendTimeout   int
	exchange_name string
}

func NewProducer(ch *amqp.Channel, cfg *configs.Config) (Producer, error) {
	if err := initQueue(ch, cfg.RabbitMQ.ProduceQueue); err != nil { //idempotent
		return nil, err
	}

	if err := initExchange(ch, cfg.RabbitMQ.ProduceQueue); err != nil { //idempotent
		return nil, err
	}

	return &producer{
		ch:            ch,
		sendTimeout:   cfg.RabbitMQ.ProduceTimeout,
		exchange_name: cfg.RabbitMQ.ProduceQueue,
	}, nil
}

func (p *producer) PublishMessage(ctx context.Context, msg []byte) error {

	ctx, cancel := context.WithTimeout(ctx, time.Duration(p.sendTimeout)*time.Second)

	defer cancel()

	if err := p.ch.PublishWithContext(ctx,
		p.exchange_name, // exchange
		"",              // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         msg,
			Type:         "blabla",
		}); err != nil {
		return err
	}

	return nil
}
