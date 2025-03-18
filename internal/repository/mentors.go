package repository

import (
	"context"
	"errors"

	"github.com/jeevangb/authservice/internal/models"
	"gorm.io/gorm"
)

func (r *UserRepo) CreateProject(ctx context.Context, project models.Project) error {
	if err := r.db.Create(&project).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) GetProjectByTitle(ctx context.Context, title string, project *models.Project) (*models.Project, error) {
	if err := r.db.WithContext(ctx).Where("title = ?", title).First(&project).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return project, nil
}
