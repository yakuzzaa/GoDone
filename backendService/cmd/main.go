package main

import (
	"log"

	"github.com/yakuzzaa/GoDone/backendService/internal/config"
	"github.com/yakuzzaa/GoDone/backendService/internal/grpc"
	"github.com/yakuzzaa/GoDone/backendService/internal/storage"
)

func main() {
	cfg := config.MustLoad()
	db, err := storage.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err.Error())
	}
	log.Println("Database connected successfully")

	grpcServer, lis, err := grpc.SetupGRPCServer(cfg, db)
	log.Println("Listening on port ", cfg.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
