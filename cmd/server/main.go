package main

import (
	"context"
	"log"

	"go.einride.tech/cloudrunner"
)

func main() {
	if err := cloudrunner.Run(
		func(ctx context.Context) error {
			cloudrunner.Logger(ctx).Info("halloooo")
			return nil
		},
	); err != nil {
		log.Fatal(err)
	}
}
