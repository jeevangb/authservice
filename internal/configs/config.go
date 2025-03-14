package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	AuthPORT string `mapstructure:"GRPC_SERVER_PORT"`
	DbHost   string `mapstructure:"DB_HOST"`
	DbPort   string `mapstructure:"DB_PORT"`
	DbUser   string `mapstructure:"DB_USER"`
	DbPass   string `mapstructure:"DB_PASSWORD"`
	DbName   string `mapstructure:"DB_NAME"`
}

func LoadConfig(env *string) (Config, error) {
	v := viper.New()
	var cfg Config
	envs := []string{
		"GRPC_SERVER_PORT",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
	}
	// v.AddConfigPath("./")
	// v.SetConfigFile("internal/configs/dev/application.env")
	v.SetConfigFile("C:/Users/jeevangb/Documents/Projects/authservice/internal/configs/dev/application.env")
	v.ReadInConfig()
	for _, val := range envs {
		if err := v.BindEnv(val); err != nil {
			return cfg, err
		}
	}
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("error while unmarshalling data to struct")
		return cfg, err
	}
	if err := validator.New().Struct(&cfg); err != nil {
		log.Fatalf("error while unmarshalling data to struct")
		return cfg, err
	}
	return cfg, nil
}
