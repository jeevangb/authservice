package repository

import (
	"github.com/jeevangb/authservice/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

type UserService interface {
	Createuser(user models.User) error
	FindUserByEmail(email string) (*models.User, error)
}

func NewService(conn *gorm.DB) UserService {
	return &UserRepo{
		db: conn,
	}
}
