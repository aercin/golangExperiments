package abstractions

import "github.com/labstack/echo/v4"

type Handlers interface {
	PlaceOrder(e echo.Context) error
	GetOrder(e echo.Context) error
}
