package main

import (
	"fmt"
	"github.com/yakuzzaa/GoDone/apiGateway"
	_ "github.com/yakuzzaa/GoDone/apiGateway/docs"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/config"
	"github.com/yakuzzaa/GoDone/apiGateway/pkg/handler"
	"log"
)

func main() {
	configLoad := config.MustLoad()
	fmt.Println(configLoad)
	handlers := new(handler.Handler)
	srv := new(apiGateway.Server)
	if err := srv.Run(configLoad.Address, handlers.InitRoutes()); err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}
}
