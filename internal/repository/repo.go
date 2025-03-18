package repository

import (
	"context"

	"github.com/jeevangb/authservice/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

type UserService interface {
	Createuser(user models.User) error
	FindUserByEmail(email string) (*models.User, error)
	CreateProject(ctx context.Context, project models.Project) error
	GetProjectByTitle(ctx context.Context, title string, project *models.Project) (*models.Project, error)
	SaveProject(project *models.Project) (*models.Project, error)
}

func NewService(conn *gorm.DB) UserService {
	return &UserRepo{
		db: conn,
	}
}
