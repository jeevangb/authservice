package services

import (
	"context"

	"github.com/jeevangb/authservice/internal/grpc/proto"
)

type GrpcImpl struct {
	proto.AuthServiceServer
	ctx context.Context
}

func NewService(ctx context.Context) GrpcImpl {
	return GrpcImpl{
		ctx: ctx,
	}
}
