package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yakuzzaa/GoDone/backendService/internal/config"
	server "github.com/yakuzzaa/GoDone/backendService/internal/grpc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	cfg := config.MustLoad()

	db, err := setupDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	grpcServer, lis, err := setupGRPCServer(cfg, db)
	if err != nil {
		log.Fatalf("failed to setup gRPC server: %v", err)
	}
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			log.Fatalf("failed to close gRPC listener: %v", err)
		}
	}(lis)

	log.Println("Listening on port", cfg.Port)

	if err := runGRPCServer(grpcServer, lis); err != nil {
		log.Fatalf("failed to run gRPC server: %v", err)
	}

}

func setupDatabase(cfg *config.Config) (*gorm.DB, error) {
	db, err := config.Connect(cfg)
	if err != nil {
		return nil, err
	}
	log.Println("Database connected successfully")
	return db, nil
}

func setupGRPCServer(cfg *config.Config, db *gorm.DB) (*grpc.Server, net.Listener, error) {
	grpcServer, lis, err := server.SetupGRPCServer(cfg, db)
	if err != nil {
		return nil, nil, err
	}
	return grpcServer, lis, nil
}

func runGRPCServer(grpcServer *grpc.Server, lis net.Listener) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	serverStopped := make(chan struct{})

	go func() {
		if err := grpcServer.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			log.Printf("failed to serve: %v", err)
		}
		close(serverStopped)
	}()

	select {
	case sig := <-quit:
		log.Printf("Received signal %v, initiating shutdown...", sig)
	case <-serverStopped:
		log.Println("gRPC server stopped unexpectedly")
		return errors.New("gRPC server stopped unexpectedly")
	}

	grpcServer.GracefulStop()
	log.Println("gRPC server stopped, waiting for connections to close...")

	timeout := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct{})
	go func() {
		grpcServer.Stop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("gRPC server exited gracefully")
	case <-ctx.Done():
		log.Println("gRPC server shutdown timeout exceeded")
		return ctx.Err()
	}

	return nil
}
