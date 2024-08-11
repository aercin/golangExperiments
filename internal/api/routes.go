package api

import (
	"go-poc/internal/api/abstractions"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Group, h abstractions.Handlers) {
	e.POST("/orders", h.PlaceOrder)
	e.GET("/orders/:order_no", h.GetOrder)
}
