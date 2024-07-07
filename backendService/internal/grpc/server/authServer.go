package server

import (
	"context"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/auth_v1"
	"github.com/yakuzzaa/GoDone/backendService/internal/grpc/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	auth_v1.UnimplementedAuthV1Server
	Service service.AuthServiceInterface
}

func NewAuthServer(service service.AuthServiceInterface) *AuthServer {
	return &AuthServer{
		Service: service,
	}
}

func (s *AuthServer) SignIn(ctx context.Context, req *auth_v1.SignInRequest) (*auth_v1.SignInResponse, error) {
	createUser, err := s.Service.CreateUser(req.Info)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	return &auth_v1.SignInResponse{Id: createUser}, nil
}

func (s *AuthServer) SignUp(ctx context.Context, req *auth_v1.SignUpRequest) (*auth_v1.SignUpResponse, error) {
	token, err := s.Service.Login(req.Info)
	if err != nil {
		return &auth_v1.SignUpResponse{}, status.Errorf(codes.Internal, "failed to sign up: %v", err)
	}
	return &auth_v1.SignUpResponse{Token: token}, nil
}
