package rabbitMQ

import (
	"context"
	"fmt"
	"go-poc/configs"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	ConsumeMessages(ctx context.Context, cb func([]byte) bool) error
}

type consumer struct {
	ch         *amqp.Channel
	queue_name string
}

func NewConsumer(ch *amqp.Channel, cfg *configs.Config) (Consumer, error) {
	if err := initQueue(ch, cfg.RabbitMQ.ConsumeQueue); err != nil {
		return nil, err
	}

	if err := initExchange(ch, cfg.RabbitMQ.ConsumeQueue); err != nil {
		return nil, err
	}

	if err := bindExchangeToQueue(ch, cfg.RabbitMQ.ConsumeQueue, cfg.RabbitMQ.ConsumeQueue); err != nil {
		return nil, err
	}

	return &consumer{
		ch:         ch,
		queue_name: cfg.RabbitMQ.ConsumeQueue,
	}, nil
}

func (c *consumer) ConsumeMessages(ctx context.Context, cb func([]byte) bool) error {

	msgs, err := c.ch.Consume(
		c.queue_name, // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		return err
	}

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			isSuccessed := cb(msg.Body)
			if isSuccessed {
				msg.Ack(false)
			} else {
				msg.Nack(false, true)
			}
		}
	}()

	fmt.Println("Incoming messages are listening...")

	<-forever

	return nil
}
