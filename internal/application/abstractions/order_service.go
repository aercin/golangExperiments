package abstractions

import (
	"context"
	"go-poc/internal/application/models/change_order_status"
	"go-poc/internal/application/models/get_order"
	"go-poc/internal/application/models/place_order"
)

type OrderService interface {
	PlaceOrder(ctx context.Context, request place_order.Request) place_order.Response
	GetOrder(ctx context.Context, request get_order.Request) get_order.Response
	ChangeOrderStatus(ctx context.Context, request change_order_status.Request) change_order_status.Response
}
