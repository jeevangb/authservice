package main

import (
	"context"
	"flag"

	config "github.com/jeevangb/authservice/internal/configs"
	"github.com/jeevangb/authservice/internal/database"
	"github.com/jeevangb/authservice/internal/repository"
	"github.com/jeevangb/authservice/internal/server"
	"github.com/jeevangb/authservice/internal/services"
)

func main() {
	env := flag.String("env", "", "")
	flag.Parse()
	// If 'env' is not provided, set it to the default value
	if *env == "" {
		*env = "dev" // Set your default environment value here
	}
	cfg, err := config.LoadConfig(env)
	if err != nil {
		return
	}
	conn, err := database.GetConnection(cfg)
	if err != nil {
		return
	}
	usersrc := repository.NewService(conn)
	srv := services.NewService(context.Background(), usersrc)
	err = server.RegisterAuthGrpcServer(cfg.AuthPORT, srv)
	if err != nil {
		return
	}
}
