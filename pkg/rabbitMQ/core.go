package rabbitMQ

import (
	"go-poc/configs"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ(cfg *configs.Config) (*amqp.Channel, error) {
	conn, err := amqp.Dial(cfg.RabbitMQ.BrokerAddress)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func initQueue(ch *amqp.Channel, queue_name string) error {
	_, err := ch.QueueDeclare(
		queue_name,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	return nil
}

func initExchange(ch *amqp.Channel, exchange_name string) error {
	if err := ch.ExchangeDeclare(
		exchange_name, // name
		"fanout",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	); err != nil {
		return err
	}
	return nil
}

func bindExchangeToQueue(ch *amqp.Channel, exchange_name, queue_name string) error {
	if err := ch.QueueBind(
		queue_name,    // queue name
		"",            // routing key
		exchange_name, // exchange
		false,
		nil); err != nil {
		return err
	}
	return nil
}
