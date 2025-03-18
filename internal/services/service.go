package services

import (
	"context"

	"github.com/jeevangb/authservice/internal/grpc/proto"
	"github.com/jeevangb/authservice/internal/repository"
)

type GrpcImpl struct {
	proto.AuthServiceServer
	proto.ProjectServiceServer
	usrsrc repository.UserService
	ctx    context.Context
}

func NewService(ctx context.Context, usrc repository.UserService) GrpcImpl {
	return GrpcImpl{
		ctx:    ctx,
		usrsrc: usrc,
	}
}
