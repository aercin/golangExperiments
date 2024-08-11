package main

import (
	"context"
	"errors"
	"fmt"
	"go-poc/configs"
	integrationevents "go-poc/internal/application/integration_events"
	"go-poc/internal/application/models/change_order_status"
	"go-poc/internal/interactor"
	"go-poc/pkg/rabbitMQ"

	"go-poc/internal/domain/constants"

	jsoniter "github.com/json-iterator/go"
)

func main() {

	fmt.Println("Consumer service is preparing now..")

	cfg := configs.NewConfig()

	orderSvc := interactor.ResolveOrderService(cfg)

	ch, _ := rabbitMQ.InitRabbitMQ(cfg)

	consumer, _ := rabbitMQ.NewConsumer(ch, cfg)

	consumer.ConsumeMessages(context.Background(), func(msg []byte) bool {

		var consumedEvent integrationevents.StockReportedEvent

		if err := jsoniter.Unmarshal(msg, &consumedEvent); err != nil {
			panic(err.Error())
		}

		newOrderStatus, err := DefineOrderStatus(consumedEvent.MessageType)

		if err != nil {
			panic(err.Error())
		}

		changeOrderStatusReq := change_order_status.Request{
			MessageId:   consumedEvent.MessageId,
			OrderNo:     consumedEvent.Message.OrderNo,
			OrderStatus: newOrderStatus,
		}

		changeOrderStatusRes := orderSvc.ChangeOrderStatus(context.Background(), changeOrderStatusReq)

		return changeOrderStatusRes.IsSuccess
	})
}

func DefineOrderStatus(msgTypes []string) (int, error) {
	for _, msgType := range msgTypes {
		if msgType == integrationevents.StockNotDecreasedEventMessageType {
			return constants.Failed, nil
		} else if msgType == integrationevents.StockDecreasedEventMessageType {
			return constants.Successed, nil
		}
	}
	return 0, errors.New("an error occurred, order status did not define")
}
