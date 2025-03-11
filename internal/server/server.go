package server

import (
	"net"

	"github.com/jeevangb/authservice/internal/services"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func RegisterAuthGrpcServer(port string, srv services.GrpcImpl) error {
	//Initialize the lister
	listner, err := net.Listen("tcp", port)
	if err != nil {
		log.Error().Err(err).Msg("error while listening port")
	}
	//create new grpc server instance
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(50 * 1024 * 1024))
	RegisterGrpcServer(grpcServer, &srv)
	log.Info().Msg("auth service listening on port: 8081")
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Error().Err(err).Msg("error while running grpc  server")
	}
	return nil
}
