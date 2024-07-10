package main

import (
	"fmt"
	"log"

	_ "github.com/yakuzzaa/GoDone/apiGateway/docs"
	"github.com/yakuzzaa/GoDone/apiGateway/internal"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/config"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/handler"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/auth_v1"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/item_v1"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/list_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title GoDone API
// @version 1.0
// @description Simple To-Do list backend
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	configLoad := config.MustLoad()
	fmt.Println(configLoad)

	// Connect to gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	authClient := auth_v1.NewAuthV1Client(conn)
	listClient := list_v1.NewListV1Client(conn)
	itemClient := item_v1.NewItemV1Client(conn)

	handlers := handler.NewHandler(authClient, listClient, itemClient)
	srv := new(internal.Server)
	if err := srv.Run(configLoad.Address, handlers.InitRoutes()); err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}
}
