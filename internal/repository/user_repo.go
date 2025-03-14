package repository

import (
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
