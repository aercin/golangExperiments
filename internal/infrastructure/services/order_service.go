package services

import (
	"context"
	"fmt"
	app_abstraction "go-poc/internal/application/abstractions"
	integrationevents "go-poc/internal/application/integration_events"
	"go-poc/internal/application/models/change_order_status"
	"go-poc/internal/application/models/get_order"
	"go-poc/internal/application/models/place_order"
	repo_abstraction "go-poc/internal/domain/abstractions"
	"go-poc/internal/domain/entities"
	"time"

	"go-poc/configs"

	"net/http"

	"go-poc/internal/domain/constants"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	jsoniter "github.com/json-iterator/go"
)

type OrderService struct {
	orderRepo repo_abstraction.OrderRepository
	inboxRepo repo_abstraction.InboxRepository
	config    *configs.Config
}

func NewOrderService(orderRepo repo_abstraction.OrderRepository, inboxRepo repo_abstraction.InboxRepository, cfg *configs.Config) app_abstraction.OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		inboxRepo: inboxRepo,
		config:    cfg,
	}
}

func (srv *OrderService) PlaceOrder(ctx context.Context, request place_order.Request) place_order.Response {

	var userBasket place_order.BasketResponse

	//TODO: belki buraya backoff paketi ile retry mekanizması konabilir.
	basketRes, err := resty.New().R().
		SetQueryParam("UserId", request.UserId).
		SetResult(&userBasket).
		Get(srv.config.BasketService.Address)
	if err != nil ||
		basketRes.StatusCode() != http.StatusOK ||
		!userBasket.IsSuccess {
		return place_order.Response{
			IsSuccess: false,
		}
	}

	var orderProducts []entities.OrderProduct
	copier.Copy(&orderProducts, &userBasket.Data)

	order := entities.Order{
		OrderNo:       request.OrderNo,
		UserId:        request.UserId,
		Status:        constants.Suspend,
		OrderProducts: orderProducts,
	}

	var orderPlacedEvent integrationevents.OrderPlacedEvent
	uuid := uuid.New().String()
	orderPlacedEvent.ConversationId = uuid
	orderPlacedEvent.MessageId = uuid
	orderPlacedEvent.MessageType = []string{
		integrationevents.IntegrationEventBaseMessageType,
		integrationevents.OrderPlacedEventMessageType,
	}
	orderPlacedEvent.Message.OrderNo = order.OrderNo
	copier.Copy(&orderPlacedEvent.Message.Items, &order.OrderProducts)

	serializedOrderPlacedEvent, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(orderPlacedEvent)

	outboxMessage := entities.OutboxMessage{
		Message:   string(serializedOrderPlacedEvent),
		CreatedOn: time.Now(),
	}

	if err := srv.orderRepo.Create(ctx, &order, &outboxMessage); err != nil {
		return place_order.Response{
			IsSuccess: false,
		}
	}

	return place_order.Response{
		IsSuccess: true,
		OrderId:   order.Id,
	}
}

func (srv *OrderService) GetOrder(ctx context.Context, request get_order.Request) get_order.Response {

	query, _, err := goqu.From("orders").Where(goqu.Ex{
		"order_no": request.OrderNo,
	}).ToSQL()

	if err != nil {
		panic(fmt.Sprintf("Raise an error when query an order: %v", err.Error()))
	}

	result, err := srv.orderRepo.Find(ctx, query)
	if err != nil {
		panic(fmt.Sprintf("Raise an error when query an order: %v", err.Error()))
	}

	if len(result) == 0 {
		return get_order.Response{}
	}

	var response get_order.Response

	copier.Copy(&response, &result[0])

	return response
}

func (srv *OrderService) ChangeOrderStatus(ctx context.Context, request change_order_status.Request) change_order_status.Response {

	if srv.inboxRepo.Any(ctx, request.MessageId) {
		return change_order_status.Response{
			IsSuccess: true,
		}
	}

	query, _, _ := goqu.From("orders").Where(goqu.Ex{
		"order_no": request.OrderNo,
	}).ToSQL()

	order, err := srv.orderRepo.Get(ctx, query)
	if err != nil {
		return change_order_status.Response{
			IsSuccess: false,
		}
	}
	order.Status = request.OrderStatus

	inboxMessage := entities.InboxMessage{
		MessageId: request.MessageId,
		CreatedOn: time.Now(),
	}

	if err := srv.orderRepo.ChangeOrderStatus(ctx, &order, &inboxMessage); err != nil {
		return change_order_status.Response{
			IsSuccess: false,
		}
	}

	return change_order_status.Response{
		IsSuccess: true,
	}
}
