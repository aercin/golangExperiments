package v1

import (
	"context"
	"go-poc/internal/api/abstractions"
	application "go-poc/internal/application/abstractions"
	"go-poc/internal/application/models/get_order"
	"go-poc/internal/application/models/place_order"
	"net/http"

	"go-poc/pkg/logrus"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	orderService application.OrderService
	logger       logrus.Logger
}

func NewHandler(orderSrv application.OrderService, logger logrus.Logger) abstractions.Handlers {
	return &Handlers{
		orderService: orderSrv,
		logger:       logger,
	}
}

func (h *Handlers) PlaceOrder(c echo.Context) error {

	request := new(place_order.Request)

	c.Bind(&request)

	validate := validator.New() //todo: bunu middleware olarak yapabilirim

	if err := validate.Struct(request); err != nil {
		h.logger.Log(err.Error(), &logrus.Configs{
			Level: logrus.Error,
		})
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	response := h.orderService.PlaceOrder(ctx, *request)

	return c.JSON(http.StatusOK, response)
}

func (h *Handlers) GetOrder(c echo.Context) error {

	order_no := c.Param("order_no")

	if order_no == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "getOrder ep's route item must pass")
	}

	request := &get_order.Request{
		OrderNo: order_no,
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	response := h.orderService.GetOrder(ctx, *request)

	return c.JSON(http.StatusOK, response)
}
