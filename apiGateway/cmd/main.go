package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

func main() {
	configLoad := config.MustLoad()

	logger, err := setupLogger(configLoad)
	if err != nil {
		slog.Error("Failed to setup logger: %v", err)
		os.Exit(1)
	}
	slog.SetDefault(logger)

	conn, err := setupGRPCConnection()
	if err != nil {
		slog.Error("Failed to setup gRPC connection: %v", err)
		os.Exit(1)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			slog.Error("Failed to close gRPC connection: %v", err)
		}
	}(conn)

	authClient := auth_v1.NewAuthV1Client(conn)
	listClient := list_v1.NewListV1Client(conn)
	itemClient := item_v1.NewItemV1Client(conn)

	handlers := handler.NewHandler(authClient, listClient, itemClient)
	srv := new(internal.Server)

	if err := runServer(srv, configLoad.Address, handlers.InitRoutes()); err != nil {
		slog.Error("Failed to run server: %v", err)
		os.Exit(1)
	}
	slog.Info("Server exited gracefully")
}

func setupLogger(configLoad *config.Config) (*slog.Logger, error) {
	logger := config.SetupLogger(configLoad.Env, configLoad.LogPath)
	slog.Info("Logger set up successfully", slog.String("env", configLoad.Env))
	return logger, nil
}

func setupGRPCConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	slog.Info("gRPC connection established")
	return conn, nil
}

func runServer(srv *internal.Server, address string, routes http.Handler) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := srv.Run(address, routes); err != nil {
			slog.Error("Something went wrong: %s", err)
			cancel()
		}
	}()

	<-quit
	slog.Info("Shutting down server...")

	ctxTimeout, cancelTimeout := context.WithTimeout(ctx, 5*time.Second)
	defer cancelTimeout()

	if err := srv.Shutdown(ctxTimeout); err != nil {
		return err
	}

	slog.Info("Server exiting")
	return nil
}
