package main

import (
	"context"
	"log"

	"github.com/giovanniberi93/einride-backend-onboarding/internal/app/server"
	"go.einride.tech/cloudrunner"
)

func main() {
	var config server.Config
	if err := cloudrunner.Run(
		func(ctx context.Context) error {
			cloudrunner.Logger(ctx).Info("halloooo")
			app, cleanup, err := server.Init(ctx, &config)
			if err != nil {
				return err
			}
			defer cleanup()
			return cloudrunner.ListenGRPC(ctx, app.GRPCServer)
		},
		cloudrunner.WithConfig("server", &config),
	); err != nil {
		log.Fatal(err)
	}
}
