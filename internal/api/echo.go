package api

import (
	"fmt"
	"go-poc/configs"
	"go-poc/internal/interactor"

	"github.com/labstack/echo/v4"
)

func NewHttpServer(c *configs.Config) {
	e := echo.New()

	//version-1 routes group
	eg_v1 := e.Group("api/v1.0")

	NewRouter(eg_v1, interactor.ResolveHandler(c))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", c.HttpServer.Port)))
}
