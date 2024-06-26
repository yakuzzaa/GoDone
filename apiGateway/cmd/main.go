package main

import (
	"github.com/yakuzzaa/GoDone/apiGateway"
	"github.com/yakuzzaa/GoDone/apiGateway/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(apiGateway.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}
}
