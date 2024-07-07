package repository

import (
	"errors"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/auth_v1"
	"github.com/yakuzzaa/GoDone/backendService/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	CreateUser(userInfo *auth_v1.SignInInfo) (uint64, error)
	GetUser(userInfo *auth_v1.SignUpInfo) (uint64, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepositoryInterface {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(userInfo *auth_v1.SignInInfo) (uint64, error) {
	user := models.User{
		Name:     userInfo.Name,
		Username: userInfo.Username,
		Password: userInfo.Password,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}

	return uint64(user.ID), nil
}

func (r *AuthRepository) GetUser(userInfo *auth_v1.SignUpInfo) (uint64, error) {
	var user models.User

	if err := r.db.Where("username = ?", userInfo.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("user not found")
		}
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password)); err != nil {
		return 0, errors.New("invalid password")
	}

	return uint64(user.ID), nil
}
