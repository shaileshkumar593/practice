package grpcserver

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc/status"
)

func LoggingInterceptor(
	logger *slog.Logger,
) grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		start := time.Now()

		resp, err := handler(
			ctx,
			req,
		)

		logger.Info(
			"grpc request",
			"method", info.FullMethod,
			"duration", time.Since(start),
			"error", err,
		)

		return resp, err
	}
}

func RecoveryInterceptor(
	logger *slog.Logger,
) grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (
		resp any,
		err error,
	) {

		defer func() {

			if r := recover(); r != nil {

				logger.Error(
					"panic recovered",
					"method",
					info.FullMethod,
					"panic",
					r,
				)

				err = status.Error(
					codes.Internal,
					"internal server error",
				)
			}

		}()

		return handler(ctx, req)
	}
}

func AuthInterceptor(
	validate func(string) error,
) grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		md, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return nil, status.Error(
				codes.Unauthenticated,
				"missing metadata",
			)
		}

		values := md.Get("authorization")

		if len(values) == 0 {
			return nil, status.Error(
				codes.Unauthenticated,
				"missing authorization",
			)
		}

		if err := validate(values[0]); err != nil {
			return nil, status.Error(
				codes.Unauthenticated,
				"invalid token",
			)
		}

		return handler(ctx, req)
	}
}
