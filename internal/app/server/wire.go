//go:build wireinject

package server

import (
	"context"

	"github.com/google/wire"
)

func Init(_ context.Context, _ *Config) (*App, func(), error) {
	panic(
		wire.Build(
			wire.Struct(new(App), "*"),
			initGRPCServer,
		),
	)
}
