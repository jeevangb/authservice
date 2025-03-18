package repository

import (
	"fmt"

	"github.com/jeevangb/authservice/internal/models"
	"github.com/rs/zerolog/log"
)

func (r *UserRepo) Createuser(user models.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		log.Error().Err(err).Msg("error in creating a user")
		return err
	}
	return nil
}

func (r *UserRepo) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email=?", email).First(&user).Error; err != nil {
		log.Error().Err(err).Msg("user not found")
		return &models.User{}, err
	}
	return &user, nil
}

func (r *UserRepo) SaveProject(project *models.Project) (*models.Project, error) {
	if err := r.db.Save(project).Error; err != nil {
		return nil, fmt.Errorf("failed to update project: %v", err)
	}
	return project, nil
}
