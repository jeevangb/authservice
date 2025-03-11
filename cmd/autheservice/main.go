package main

import (
	"context"
	"flag"

	config "github.com/jeevangb/authservice/internal/configs"
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
	srv := services.NewService(context.Background())
	err = server.RegisterAuthGrpcServer(cfg.AuthPORT, srv)
	if err != nil {
		return
	}
}
