package server

import (
	"github.com/jeevangb/authservice/internal/grpc/proto"
	"google.golang.org/grpc"
)

func RegisterGrpcServer(s grpc.ServiceRegistrar, srv any) {
	proto.RegisterAuthServiceServer(s, srv.(proto.AuthServiceServer))
}
