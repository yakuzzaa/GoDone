package grpc

import (
	"fmt"
	"net"

	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/auth_v1"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/item_v1"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/list_v1"
	"github.com/yakuzzaa/GoDone/backendService/internal/config"
	"github.com/yakuzzaa/GoDone/backendService/internal/grpc/repository"
	"github.com/yakuzzaa/GoDone/backendService/internal/grpc/server"
	"github.com/yakuzzaa/GoDone/backendService/internal/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func SetupGRPCServer(cfg *config.Config, db *gorm.DB) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authServer := server.NewAuthServer(authService)

	listRepo := repository.NewListRepository(db)
	listService := service.NewListService(listRepo)
	listServer := server.NewListServer(listService)

	itemRepo := repository.NewItemRepository(db)
	itemService := service.NewItemRepository(itemRepo)
	itemServer := server.NewItemServer(itemService)

	auth_v1.RegisterAuthV1Server(grpcServer, authServer)
	list_v1.RegisterListV1Server(grpcServer, listServer)
	item_v1.RegisterItemV1Server(grpcServer, itemServer)

	return grpcServer, lis, nil
}
