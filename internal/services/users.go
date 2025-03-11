package services

import (
	"context"
	"sync"

	"github.com/jeevangb/authservice/internal/grpc/proto"
	"github.com/jeevangb/authservice/internal/utils"
)

var users sync.Map

func (s *GrpcImpl) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.UserResponse, error) {
	if _, exists := users.Load(req.Email); exists {
		return &proto.UserResponse{
			Message: "email already exists",
		}, nil
	}
	hashPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return &proto.UserResponse{
			Message: "internal server error",
		}, err
	}
	req.Password = hashPass
	users.Store(req.Email, req)
	return &proto.UserResponse{
		Message: "signup successfull",
	}, nil
}
