package server

import (
	"context"

	"go.einride.tech/cloudrunner"
	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpc.Server
}

func initGRPCServer(ctx context.Context) *grpc.Server {
	grpcServer := cloudrunner.NewGRPCServer(ctx)
	return grpcServer
}
