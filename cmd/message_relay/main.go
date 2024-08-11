package main

import (
	"context"
	"fmt"
	"go-poc/configs"
	"go-poc/internal/interactor"
	"time"
)

func main() {

	fmt.Println("Message relay service is preparing now..")

	cfg := configs.NewConfig()

	eventDispatcher := interactor.ResolveEventDispatcher(cfg)

	fmt.Printf("Message relay service is started at %v \n", time.Now())

	for {

		if err := eventDispatcher.DispatchEvents(context.Background()); err != nil {
			fmt.Printf("An error occurred thats details : %v\n", err)
		}

		time.Sleep(time.Duration(cfg.MessageRelay.CycleTime) * time.Second)
	}
}
