package main

import (
	"go-poc/configs"
	"go-poc/internal/api"
)

func main() {
	c := configs.NewConfig()
	api.NewHttpServer(c)
}
