package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/auth_v1"
	"github.com/yakuzzaa/GoDone/backendService/internal/grpc/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	signingKey = "d23ud#bGHK54hds#ci5c"
	tokenTTL   = 24 * time.Hour
)

type AuthServiceInterface interface {
	CreateUser(userInfo *auth_v1.SignInInfo) (uint64, error)
	Login(userInfo *auth_v1.SignUpInfo) (string, error)
}

type AuthService struct {
	repo repository.AuthRepositoryInterface
}

func NewAuthService(repo repository.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(userInfo *auth_v1.SignInInfo) (uint64, error) {
	password, err := generatePasswordHash(userInfo.Password)
	if err != nil {
		return 0, err
	}
	userInfo.Password = password
	registration, err := s.repo.CreateUser(userInfo)
	if err != nil {
		return 0, err
	}
	return registration, nil
}

func (s *AuthService) Login(userInfo *auth_v1.SignUpInfo) (string, error) {
	user, err := s.repo.GetUser(userInfo)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"exp":    time.Now().Add(tokenTTL).Unix(),
		"data":   time.Now().Unix(),
		"userId": user,
	})
	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
