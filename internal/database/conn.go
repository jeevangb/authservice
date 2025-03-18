package database

import (
	"fmt"

	"gorm.io/gorm"

	config "github.com/jeevangb/authservice/internal/configs"
	"github.com/jeevangb/authservice/internal/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
)

func GetConnection(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DbHost, cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error().Err(err).Msg("failed to connect to database")
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Error().Err(err).Msg("failed to auto-migraete user database")
		return nil, err
	}
	err = db.AutoMigrate(&models.Project{})
	if err != nil {
		log.Error().Err(err).Msg("failed to auto-migraete project database")
		return nil, err
	}
	return db, nil
}
