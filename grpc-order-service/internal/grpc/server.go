package grpcserver

import (
	"log/slog"
	"net"

	"google.golang.org/grpc"

	orderv1 "grpc-order-service/api/gen/order/v1"
)

func NewServer(
	handler *OrderHandler,
	logger *slog.Logger,
) *grpc.Server {

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			RecoveryInterceptor(logger),
			LoggingInterceptor(logger),
		),
	)

	orderv1.RegisterOrderServiceServer(
		server,
		handler,
	)

	return server
}

func ListenAndServe(
	server *grpc.Server,
	port string,
) error {

	listener, err := net.Listen(
		"tcp",
		":"+port,
	)

	if err != nil {
		return err
	}

	return server.Serve(listener)
}
