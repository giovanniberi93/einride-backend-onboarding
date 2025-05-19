//go:build wireinject

package server

import (
	"context"

	"github.com/google/wire"
  "github.com/giovanniberi93/einride-backend-onboarding/internal/svc/freightsvc"
)

func Init(_ context.Context, _ *Config) (*App, func(), error) {
	panic(
		wire.Build(
			wire.Struct(new(App), "*"),
			initGRPCServer,
			// einride.onboarding.v1.FreightService
			wire.Struct(new(freightsvc.Server), "*"),
		),
	)
}
