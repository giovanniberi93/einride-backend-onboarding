package server

import (
	"context"

	"github.com/giovanniberi93/einride-backend-onboarding/internal/svc/freightsvc"
	"go.einride.tech/cloudrunner"
	"google.golang.org/grpc"

	onboardingv1 "github.com/giovanniberi93/einride-backend-onboarding/proto/gen/go/einride/onboarding/v1"
)

type App struct {
	GRPCServer *grpc.Server
}

func initGRPCServer(
	ctx context.Context,
	freightServer *freightsvc.Server,
) *grpc.Server {
	grpcServer := cloudrunner.NewGRPCServer(ctx)
  onboardingv1.RegisterFreightServiceServer(grpcServer, freightServer)
	return grpcServer
}
